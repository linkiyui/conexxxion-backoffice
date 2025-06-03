package logger_v2

import (
	"io"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
)

type LoggerOutput struct {
	console io.Writer
	lumberj *lumberjack.Logger
}

func NewLoggerOutput(verbose bool, filename string) *LoggerOutput {
	l := &LoggerOutput{
		console: nil,
		lumberj: nil,
	}
	l.SetLumberjackLogger(filename)
	if verbose {
		l.SetConsoleOutput()
	}
	return l
}

func (lo *LoggerOutput) SetLumberjackLogger(filename string) {
	if filename != "" {
		lo.lumberj = &lumberjack.Logger{
			Filename: filename,
			MaxSize:  500,
			Compress: true,
			MaxAge:   30,
		}
	}
}

func (lo *LoggerOutput) SetConsoleOutput() {
	lo.console = os.Stdout
}

func (lo *LoggerOutput) Write(b []byte) (int, error) {
	var (
		err   error
		wrttn int
	)
	if lo.console != nil {
		if wrttn, err = lo.console.Write(b); err != nil {
			return wrttn, err
		}
	}
	if lo.lumberj != nil {
		if wrttn, err = lo.lumberj.Write(b); err != nil {
			return wrttn, err
		}
	}
	return wrttn, nil
}

func (lo *LoggerOutput) Sync() error {
	return lo.lumberj.Close()
}
