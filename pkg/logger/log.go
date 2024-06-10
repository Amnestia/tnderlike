package logger

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

var Logger zerolog.Logger

type leveledWriter struct {
	io.Writer
	level zerolog.Level
}

func (lw *leveledWriter) WriteLevel(lv zerolog.Level, p []byte) (n int, err error) {
	if lv == zerolog.ErrorLevel {
		return lw.Writer.Write(p)
	}
	return len(p), nil
}

func InitLogger(serviceName, infoLog, errorLog string) (err error) {
	fileInfo, err := os.OpenFile(filepath.Clean(infoLog), os.O_APPEND|os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		return fmt.Errorf("failed to open %s : %s", infoLog, err)
	}
	fileError, err := os.OpenFile(filepath.Clean(errorLog), os.O_APPEND|os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		return fmt.Errorf("failed to open %s : %s", errorLog, err)
	}
	w := zerolog.MultiLevelWriter(
		zerolog.ConsoleWriter{
			Out: os.Stdout,
		},
		leveledWriter{
			zerolog.MultiLevelWriter(fileInfo),
			zerolog.DebugLevel,
		},
		leveledWriter{
			zerolog.MultiLevelWriter(fileInfo),
			zerolog.InfoLevel,
		},
		leveledWriter{
			zerolog.MultiLevelWriter(fileError),
			zerolog.WarnLevel,
		},
		leveledWriter{
			zerolog.MultiLevelWriter(fileError),
			zerolog.ErrorLevel,
		},
		leveledWriter{
			zerolog.MultiLevelWriter(fileError),
			zerolog.PanicLevel,
		},
		leveledWriter{
			zerolog.MultiLevelWriter(fileError),
			zerolog.FatalLevel,
		},
	)
	Logger = zerolog.New(w)
	return
}

// Error log error helper to show context
// ex: [context] msg : error
func ErrorWrap(err error, context string, msg ...string) error {
	context = strings.TrimSpace(context)
	if len(context) > 0 {
		context = fmt.Sprintf("[%s]", context)
	}
	message := fmt.Sprintf("%v %v : %v", context, strings.Join(msg, ", "), err)
	return errors.Wrap(err, message)
}
