package console

import "fmt"

type printerBuilder struct {
	level         int
	timestampFlag bool
	background    Color
	foreground    Color
}

func (b *printerBuilder) Background(color Color) PrinterBuilder {
	b.background = color
	return b
}

func (b *printerBuilder) Foreground(color Color) PrinterBuilder {
	b.foreground = color
	return b
}

type printer struct {
	background Color
	foreground Color
}

func (b *printerBuilder) Build() Printer {
	return &printer{
		background: b.background,
		foreground: b.foreground,
	}
}

type Printer interface {
	SetForegroundColor(Color)
	Println(string string)
	Print(string string)
	Printf(string string, args ...interface{})
}

type PrinterBuilder interface {
	Background(Color) PrinterBuilder
	Foreground(Color) PrinterBuilder
	Build() Printer
}

func ColoredPrinter(level int, tsFlag bool) *printerBuilder {
	return &printerBuilder{
		level:         level,
		timestampFlag: tsFlag,
	}
}

func (p *printer) SetForegroundColor(color Color) {
	fmt.Print(color)
}

func (p *printer) ResetForegroundColor() {
	fmt.Print(DEFAULT_GREY)
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
