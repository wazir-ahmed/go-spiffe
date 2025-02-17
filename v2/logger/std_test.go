package logger_test

import (
	"bytes"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vishnusomank/go-spiffe/v2/logger"
)

func TestStd(t *testing.T) {
	buf := new(bytes.Buffer)
	log.SetOutput(buf)
	log.SetFlags(0)

	logger.Std.Debugf("%s", "debug")
	logger.Std.Warnf("%s", "warn")
	logger.Std.Infof("%s", "info")
	logger.Std.Errorf("%s", "error")

	require.Equal(t, `[DEBUG] debug
[WARN] warn
[INFO] info
[ERROR] error
`, buf.String())
}
