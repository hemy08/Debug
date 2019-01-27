package main

import (
	"WebServer/Debug"
	"fmt"
	"os"
)

const (
	LOG_DEBUG = Debug.LOG_LEVEL_DEBUG  	//调试打印
	LOG_INFO = Debug.LOG_LEVEL_INFO
	LOG_WARNING = Debug.LOG_LEVEL_WARNING
	LOG_ERROR = Debug.LOG_LEVEL_ERROR
	LOG_EMERGENCY = Debug.LOG_LEVEL_EMERGENCY
	LOG_BUTT = Debug.LOG_LEVEL_BUTT
)

func logInit(){
	Debug.LogSetFileFullName("./testlog.log")
	Debug.LogSetOsFlag(os.O_RDWR|os.O_CREATE|os.O_APPEND)//读写、没有则创建、增量写日志
	Debug.LogSetLogMode(Debug.L_DATE|Debug.L_LEVEL|Debug.L_FILE|Debug.L_LINE|Debug.L_FUNC)
	Debug.LogSetLogLevel(LOG_INFO)//INFO级别以上的日志
	Debug.LogSetMaxSize(10)//日志文件大小10M
	Debug.LogSetOverPolicy(Debug.L_RENAME_DATE)//安装日期重命名
}

func LogRecord(level int, format string, a ...interface{}){
	Debug.LOG(level,format,a...)
}

func main(){
	fmt.Println("This is a Web Server test")
	logInit()
	LogRecord(LOG_DEBUG,"Test function log  01")
	LogRecord(LOG_INFO,"Test function log  02")
	LogRecord(LOG_WARNING,"Test function log  03")
	LogRecord(LOG_ERROR,"Test function log  04")
	LogRecord(LOG_EMERGENCY,"Test function log  05")
	LogRecord(LOG_BUTT,"Test function log  06")
	LogRecord(LOG_INFO,"Test function log  07")
}