package logging

import (
	"log"
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func init() {
	lf := filepath.Join(os.TempDir(), "lsp-example.log")
	cfg := &zap.Config{
		Level:             zap.NewAtomicLevelAt(zapcore.InfoLevel),
		Development:       true,
		DisableCaller:     true,
		DisableStacktrace: true,
		Encoding:          "json",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",
			LevelKey:   "level",
		},
		OutputPaths:      []string{lf},
		ErrorOutputPaths: []string{"stderr", lf},
	}
	if lgr, err := cfg.Build(); err == nil {
		Logger = lgr
	} else {
		log.Fatal(err)
	}
}
