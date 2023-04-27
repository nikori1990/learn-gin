package core

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"learn-gin/global"
	"os"
	"path/filepath"
	"time"
)

func InitLogger() *zap.SugaredLogger {
	encoder := getEncoder()
	writeSyncer := getWriterSyncer()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	return zap.New(core).Sugar()
}

func getEncoder() zapcore.Encoder {
	config := zap.NewProductionEncoderConfig()
	config.TimeKey = "time"
	config.EncodeLevel = zapcore.CapitalLevelEncoder
	config.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(time.Local().Format("2006-01-02 15:04:05"))
	}
	return zapcore.NewJSONEncoder(config)
}

func getWriterSyncer() zapcore.WriteSyncer {
	stSeparator := string(filepath.Separator)
	stRoot, _ := os.Getwd()
	stLogFilePath := stRoot + stSeparator + "log" + stSeparator + time.Now().Format("2006-01-02")
	fmt.Println(stLogFilePath)

	lumberjackSyncer := &lumberjack.Logger{
		Filename:   stLogFilePath,
		MaxSize:    global.CONFIG.Zap.MaxSize, // megabytes
		MaxBackups: global.CONFIG.Zap.MaxBackups,
		MaxAge:     global.CONFIG.Zap.MaxAge, //days
		Compress:   false,                    // disabled by default
	}

	return zapcore.AddSync(lumberjackSyncer)
}
