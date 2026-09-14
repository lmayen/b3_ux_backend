package logx

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func InitLogger() error {
	// --- Encoder configuration
	encCfg := zap.NewProductionEncoderConfig()
	encCfg.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	encCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encCfg.EncodeCaller = zapcore.ShortCallerEncoder

	// --- Log level
	level := zap.DebugLevel

	consoleEncoder := zapcore.NewConsoleEncoder(encCfg)
	consoleWriter := zapcore.AddSync(os.Stdout)

	core := zapcore.NewCore(
		consoleEncoder,
		consoleWriter,
		level,
	)

	// --- Build the logger
	Logger = zap.New(
		core,
		zap.AddCaller(),
		zap.AddStacktrace(zap.ErrorLevel),
	)

	return nil
}

func PrettyPrintJSON(data interface{}) string {
	pretty, err := json.MarshalIndent(data, "    ", "    ")
	if err != nil {
		return fmt.Sprintf("%+v", data)
	}
	return string(pretty)
}
