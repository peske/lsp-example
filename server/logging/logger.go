package logging

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func init() {
	cfg := &zap.Config{
		Level:             zap.NewAtomicLevelAt(zapcore.DebugLevel),
		Development:       true,
		DisableCaller:     true,
		DisableStacktrace: true,
		Encoding:          "json",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",
			LevelKey:   "level",
		},
		OutputPaths: []string{"/tmp/lsp-example.log"},
	}
	if lgr, err := cfg.Build(); err == nil {
		Logger = lgr
	} else {
		log.Fatal(err)
	}
}
