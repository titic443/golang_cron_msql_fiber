package logs

import (
	"encoding/json"
	"fmt"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	var err error

	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.StacktraceKey = ""
	log, err = config.Build(zap.AddCallerSkip(1))

	if err != nil {
		panic(err)
	}
}

func Debug(msg string, fields ...zapcore.Field) {
	log.Debug(msg, fields...)
}
func Info(msg string, fields ...zapcore.Field) {
	log.Info(msg, fields...)
}
func Error(msg interface{}, fields ...zapcore.Field) {
	switch v := msg.(type) {
	case error:
		log.Error(v.Error(), fields...)
	case string:
		log.Error(v, fields...)
	case []byte:
		var msg map[string]interface{}
		json.Unmarshal(v, &msg)
		v, _ = json.Marshal(msg["_server_messages"])
		vs := string(v)
		rfm := strings.ReplaceAll(vs, `\\\`, "")
		log.Error(rfm, fields...)
	case map[string]interface{}:
		err := ""
		for k, value := range v {
			if len(err) > 0 {
				err += k
				err += ":"
				if stringValue, ok := value.(string); ok {
					err += stringValue
				}
			} else {

				err = fmt.Sprintf("%v: %v,", k, value)

			}
		}
		rfm := strings.ReplaceAll(err, `\`, "")
		log.Error(rfm, fields...)
	}

}
