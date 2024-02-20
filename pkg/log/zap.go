package log

import (
	"go.uber.org/zap"
)

type ZapLog struct {
}

func NewZapLog() Logger {
	return &ZapLog{}
}

func (z *ZapLog) Debug(msg string, args ...Field) {
	//zap.L().Debug(fmt.Sprintf("错误%s,%s", "张三", "李四"))
	fields := z.fieldCover(args...)
	zap.L().Debug(msg, fields...)
}

func (z *ZapLog) Info(msg string, args ...Field) {
	fields := z.fieldCover(args...)
	zap.L().Info(msg, fields...)
}

func (z *ZapLog) Warn(msg string, args ...Field) {
	fields := z.fieldCover(args...)
	zap.L().Warn(msg, fields...)
}

func (z *ZapLog) Error(msg string, args ...Field) {
	fields := z.fieldCover(args...)
	zap.L().Error(msg, fields...)
}
func (z *ZapLog) fieldCover(logField ...Field) []zap.Field {
	fields := make([]zap.Field, 0)
	for _, field := range logField {
		zv := zap.Field{
			Key:    field.Key,
			String: field.Value,
		}
		fields = append(fields, zv)
	}
	return fields
}
