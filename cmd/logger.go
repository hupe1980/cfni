package cmd

import (
	"fmt"
	"os"

	"github.com/hupe1980/golog"
)

type logger struct{}

func (l *logger) Print(level golog.Level, v ...interface{}) {
	l.print(level, fmt.Sprint(v...))
}

func (l *logger) Println(level golog.Level, v ...interface{}) {
	l.print(level, fmt.Sprintln(v...))
}

func (l *logger) Printf(level golog.Level, format string, v ...interface{}) {
	l.print(level, fmt.Sprintf(format, v...))
}

func (l *logger) print(level golog.Level, msg string) {
	if level < golog.WARNING {
		fmt.Fprintf(os.Stderr, "[i] %s", msg)
	}

	fmt.Fprintf(os.Stderr, "[!] %s", msg)
}
