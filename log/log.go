package log

import (
	"log"
	"os"
	"time"
)

// Log 在指定路径写入指定日志
func Log(path, strContent string) error {
	// 追加模式打开文件，并写入目标字符串
	logFile, _ := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	defer logFile.Close()

	//创建一个Logger
	//参数1：日志写入目的地
	//参数2：每条日志的前缀
	//参数3：日志属性
	loger := log.New(logFile, "", 0)

	loger.Println(strContent)

	return nil
}

// WriteLoginLog s
func WriteLoginLog(token string) {
	// log 实例
	// 23/Aug/2010:17:26:44 +0800|GET /222 HTTP/1.1|200|Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MzYxMjc4NjgsImlzcyI6IkFwaUF1dGhvcml6YXRpb24iLCJuYW1lIjoib3BlcmF0b3J0d28iLCJ0eXBlIjoyfQ.zyvCMvIoQlPkqbyjA6LQFnv_o0p_jtO9aiV4wppClpM

	// 写入nginx日志
	logPath := "/Users/cch/go/src/github.com/my-go-test/test122.log"
	logStr := time.Now().Format("02/Jan/2006:15:04:05 -0700") +
		"|POST /api/v1/authz/login|200|" +
		token

	Log(logPath, logStr)
}
