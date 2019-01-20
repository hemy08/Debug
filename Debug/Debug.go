package Debug

import  (
	"fmt"
	"runtime"
)

func log(level int, format string, a ...interface{}){
	var strLog,fmtInfo string

	//日志级别
	if logRecordLevel > level {
		return
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
