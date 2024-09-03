package main

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"os"
	"testing"
)

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f:\n", r)
		}
	}()
	fmt.Println("Calling main()")
	main()
	fmt.Println("Returned normally from main()")
}

func TestPlayGameShotHits(t *testing.T) {
	fs, err := New("a1\na2\na3\na4\na5\nb1\nb2\nb3\nb4\nc1\nc2\nc3\nd1\nd2\nd3\ne1\ne2\nb4\na\n")
	if err != nil {
		log.Fatal(err)
	}

	f()

	b, err := fs.ReadAndRestore()
	if err != nil {
		log.Fatal(err)
	}

	assert.Contains(t, string(b), "Welcome to Battleship")
	assert.Contains(t, string(b), "[fail]Yeah ! Nice hit !")
}

func TestPlayGameShotMisses(t *testing.T) {
	fs, err := New("a1\na2\na3\na4\na5\nb1\nb2\nb3\nb4\nc1\nc2\nc3\nd1\nd2\nd3\ne1\ne2\ne4\na\n")
	if err != nil {
		log.Fatal(err)
	}

	f()

	b, err := fs.ReadAndRestore()
	if err != nil {
		log.Fatal(err)
	}

	assert.Contains(t, string(b), "Welcome to Battleship")
	assert.Contains(t, string(b), "Miss")
}

type FakeStdio struct {
	origStdout   *os.File
	stdoutReader *os.File

	outCh chan []byte

	origStdin   *os.File
	stdinWriter *os.File
}

func New(stdinText string) (*FakeStdio, error) {
	// Pipe for stdin.
	//
	//                 ======
	//  w ------------->||||------> r
	// (stdinWriter)   ======      (os.Stdin)
	stdinReader, stdinWriter, err := os.Pipe()
	if err != nil {
		return nil, err
	}

	// Pipe for stdout.
	//
	//               ======
	//  w ----------->||||------> r
	// (os.Stdout)   ======      (stdoutReader)
	stdoutReader, stdoutWriter, err := os.Pipe()
	if err != nil {
		return nil, err
	}

	origStdin := os.Stdin
	os.Stdin = stdinReader

	_, err = stdinWriter.Write([]byte(stdinText))
	if err != nil {
		stdinWriter.Close()
		os.Stdin = origStdin
		return nil, err
	}

	origStdout := os.Stdout
	os.Stdout = stdoutWriter

	outCh := make(chan []byte)

	//This goroutine reads stdout into a buffer in the background.
	go func() {
		var b bytes.Buffer
		if _, err := io.Copy(&b, stdoutReader); err != nil {
			log.Println(err)
		}
		outCh <- b.Bytes()
	}()

	return &FakeStdio{
		origStdout:   origStdout,
		stdoutReader: stdoutReader,
		outCh:        outCh,
		origStdin:    origStdin,
		stdinWriter:  stdinWriter,
	}, nil
}

func (sf *FakeStdio) ReadAndRestore() ([]byte, error) {
	if sf.stdoutReader == nil {
		return nil, fmt.Errorf("ReadAndRestore from closed FakeStdio")
	}

	// Close the writer side of the faked stdout pipe. This signals to the
	// background goroutine that it should exit.
	os.Stdout.Close()
	out := <-sf.outCh

	os.Stdout = sf.origStdout
	os.Stdin = sf.origStdin

	if sf.stdoutReader != nil {
		sf.stdoutReader.Close()
		sf.stdoutReader = nil
	}

	if sf.stdinWriter != nil {
		sf.stdinWriter.Close()
		sf.stdinWriter = nil
	}

	return out, nil
}
