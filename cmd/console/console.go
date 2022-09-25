package console

import (
	"battleship-go/cmd/contracts"
	"fmt"
)

type printerBuilder struct {
	level         int
	timestampFlag bool
	background    contracts.Color
	foreground    contracts.Color
}

func (b *printerBuilder) Background(color contracts.Color) PrinterBuilder {
	b.background = color
	return b
}

func (b *printerBuilder) Foreground(color contracts.Color) PrinterBuilder {
	b.foreground = color
	return b
}

type printer struct {
	background contracts.Color
	foreground contracts.Color
}

func (b *printerBuilder) Build() Printer {
	return &printer{
		background: b.background,
		foreground: b.foreground,
	}
}

type Printer interface {
	SetForegroundColor(contracts.Color)
	Println(string string)
	Print(string string)
	Printf(string string, args ...interface{})
}

type PrinterBuilder interface {
	Background(contracts.Color) PrinterBuilder
	Foreground(contracts.Color) PrinterBuilder
	Build() Printer
}

func ColoredPrinter(level int, tsFlag bool) *printerBuilder {
	return &printerBuilder{
		level:         level,
		timestampFlag: tsFlag,
	}
}

func (p *printer) SetForegroundColor(color contracts.Color) {
	fmt.Print(color)
}

func (p *printer) ResetForegroundColor() {
	fmt.Print(contracts.DEFAULT_GREY)
}

func (p *printer) Println(text string) {
	fmt.Println(text)
}

func (p *printer) Print(text string) {
	fmt.Print(text)
}

func (p *printer) Printf(pattern string, args ...interface{}) {
	fmt.Printf(pattern, args...)
}
