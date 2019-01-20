package Debug

import  (
	"os"
)

//日志文件路径
var logFilePath = "./"

//日志文件名
var logFileName = "serverlog.log"

//日志文件
var logFullName = "./serverlog.log"

//osFlag 日志文件的打开模式
var osFlag = 0

//os.FileMode
var osPerm os.FileMode = 0

//日志打印级别
var logRecordLevel = LOG_LEVEL_BUTT

//日志记录模式
var logRecordMode = L_BUTT

//日志
var logModeSimple = L_DATE | L_LEVEL |L_FUNC
var logModeNormal = L_DATE | L_LEVEL |L_FILE | L_LINE | L_FUNC

//日志大小 默认10M
var logMaxFileSize int64 = 5

//日志超大时的操作，默认按照日期重命名
var logOverPolicy = L_RENAME_DATE

//index
var logIndex = 1

const (
	LOG_LEVEL_DEBUG 	= 0
	LOG_LEVEL_INFO 		= 1
	LOG_LEVEL_WARNING 	= 2
	LOG_LEVEL_ERROR 	= 3
	LOG_LEVEL_EMERGENCY = 4
	LOG_LEVEL_BUTT 		= 5
)

const (
	L_NONE 	= 0x00000 //不加参数
	L_DATE 	= 0x00001  //日期
	L_LEVEL = 0x00002 //日志级别
	L_FILE 	= 0x00004 //文件名
	L_PATH 	= 0x00008 //文件路径
	L_LINE 	= 0x00010  //文件行数
	L_FUNC 	= 0x00020 //函数名
	L_ALL	= 0x00040 //全部
	L_BUTT 	= 0xfffff //非法
)

var log_level_string = []string{
	"DEBUG",
	"INFORMATION",
	"WARNING",
	"ERROR",
	"EMERGENCY",
	"UNKNOW",
}

const (
	L_DELETE 	= 0
	L_RENAME_INDEX 		= 1
	L_RENAME_DATE 	= 2
)