package Debug

import  (
	"os"
	"fmt"
	"runtime"
)

type Log struct {
	filePath 		string	//日志文件路径
	fileName 		string	//日志文件名
	fileFullName 	string	//日志文件路径+文件名
	openFlag 		int		//osFlag 日志文件的打开模式
	openPerm 		os.FileMode	//os.FileMode
	recLevel 		int		//日志打印级别 5个级别
	recModel 		int		//日志记录模式
	logSize			int64		//单日志文件大小 单位M
	logMaxOption	int		//日志文件超出大小时的操作 删除、按日期重命名、序号重命名
}

func Init(def  Log){
	LogSetFilePath(def.filePath)
	LogSetFileName(def.fileName)
	LogSetFileFullName(def.fileFullName)
	LogSetOsFlag(def.openFlag)
	LogSetOsPerm(def.openPerm)
	LogSetLogLevel(def.recLevel)
	LogSetLogMode(def.recModel)
	LogSetMaxSize(def.logSize)
	LogSetOverPolicy(def.logMaxOption)
}

func LOG(level int, format string, a ...interface{}){
	var strLog,fmtInfo string

	//日志级别
	if logRecordLevel > level {
		return
	}

	//日志大小
	fsize := getFileSize(getLogFile());
	if fsize > logMaxFileSize {
		logFileOverProcess()
	}

	fmtInfo = fmt.Sprintf(format,a...)

	//函数信息
	pc,filepath,line,ok := runtime.Caller(1)
	if !ok {
		strLog = logModeDate() + logModeLevel(level) + fmtInfo + "\n"
		writeWithGolog(getLogFile(),strLog)
		return
	}

	//日志信息
	strLog = logModeDate() + logModeLevel(level) + logModePath(filepath) + logModeFile(filepath) + logModeLine(line) + logModeFunc(pc) + fmtInfo + "\n"
	writeWithGolog(getLogFile(),strLog)
}

//设置日志文件信息 默认是./serverlog.log
func LogSetFileFullName(fullname string) {
	logFullName = fullname
}

//设置日志文件路径
func LogSetFilePath(filepath string) {
	logFilePath = filepath
}

//设置日志文件名称
func LogSetFileName(filename string){
	logFileName = filename
}

//设置文件打开参数，默认是os.O_RDWR|os.O_CREATE|os.O_APPEND，读写，创建，增量写
func LogSetOsFlag(flag int) {
	osFlag = flag
}

//设置 FileMode，默认是0644
func LogSetOsPerm(perm os.FileMode) {
	osPerm = perm
}

//设置日志打印模式
//提供的有date，level，funcname，filename，filepath，line，默认 date+level+file+line+func
func LogSetLogMode(flag int) {
	logRecordMode = flag
}

//设置日志打印级别，默认ERR
func LogSetLogLevel(level int) {
	logRecordLevel = level
}

//设置日志文件大小,单位M
func LogSetMaxSize(logsize int64) {
	logMaxFileSize = logsize/1024 + 1
}

//设置日志文件超过最大内存的策略
func LogSetOverPolicy(logop int) {
	logOverPolicy = logop
}
