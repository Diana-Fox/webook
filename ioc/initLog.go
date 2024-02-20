package ioc

import "go.uber.org/zap"

func InitLog() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return
	}
	zap.ReplaceGlobals(logger) //要放回全局，就可以全局使用了
}
func InitLogV1() {
	//logger, _ := zap.NewProduction()
	//defer logger.Sync() // flushes buffer, if any
	//sugar := logger.Sugar()
	//sugar.Infow("failed to fetch URL",
	//	// Structured context as loosely typed key-value pairs.
	//	"url", url,
	//	"attempt", 3,
	//	"backoff", time.Second,
	//)
	//sugar.Infof("Failed to fetch URL: %s", url)
}
