package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

var logTime string
var loggger *os.File

// LoggerWrite
// 重写一个接口日志 输出到控制台并输出到日志文件中
func LoggerWrite() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestTime := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method
		ip := c.ClientIP()
		if logTime != requestTime.Format("2006-01-02") {
			loggger.Close()
			logTime = getLogTime()
			loggger = ioWrite("./log/", logTime)
		}
		s := requestTime.Format("2006-01-02 15:04:05")
		sprintf := fmt.Sprintf("[%s][%s][%s][%s]", s, method, ip, path)
		_, err := loggger.WriteString(sprintf + "\n")
		if err != nil {
			panic(err)
		}
		c.Next()
	}
}

// getLogTime 返回日志文件名
func getLogTime() string {
	return time.Now().Format("2006-01-02")
}

// ioWrite 返回一个io事件
func ioWrite(filepath, filename string) *os.File {
	name := filepath + filename + ".log"
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY|os.O_SYNC|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return file
}
