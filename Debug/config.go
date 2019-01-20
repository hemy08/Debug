package Debug

import  (
	"fmt"
	"time"
	"runtime"
	"path/filepath"
	"strings"
	"os"
	"io"
	"path"
)


//CurrentTime 当前时间，2018-12-22 14:41:21.4728403
func currentTime() string {
	t := time.Now()
	//fmt.Println(time.Now().Format("2006-01-02 15:04:05")) // 这是个奇葩,必须是这个时间点, 据说是go诞生之日, 记忆方法:6-1-2-3-4-5
	cur := fmt.Sprintf("%s.%2d", time.Now().Format("2006-01-02 15:04:05"), t.Nanosecond())
	return cur
}

func logModeDate() string {
	logmode := getLogMode()
	if logmode&L_DATE != 0 {
		return currentTime()
	}
	return ""
}

func logModeLevel(level int) string {
	logmode := getLogMode()
	if logmode&L_LEVEL != 0 {
		return "\t[" + log_level_string[level] + "] "
	}
	return ""
}

func logModeFile(file string) string {
	logmode := getLogMode()
	if logmode&L_FILE != 0{
		return fmt.Sprintf("file:%s ",getFileName(file))
	}
	return ""
}

func logModePath(path string) string {
	logmode := getLogMode()
	if logmode&L_PATH != 0{
		return fmt.Sprintf("path:%s ",getFilePath(path))
	}
	return ""
}

func logModeLine(line int) string {
	logmode := getLogMode()
	if logmode&L_LINE != 0{
		return fmt.Sprintf("line:%d ",line)
	}
	return ""
}

func logModeFunc(pc uintptr) string {
	logmode := getLogMode()
	if logmode&L_FUNC != 0{
		return fmt.Sprintf("func:%s ",getFuncName(pc))
	}
	return ""
}

//写日志文件
func writeWithGolog(filename,content string) error {
	f, err := os.OpenFile(filename, getOsFlag(), getOsPerm())
	if err != nil {
		return err
	}

	data :=[]byte(content)
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}

	return err
}

//获取运行文件路径
func getFilePath(files string) string {
	filepath, _ := filepath.Split(files)
	return filepath
}

//获取运行文件名
func getFileName(files string) string {
	_, fileName := filepath.Split(files)
	return fileName
}

//获取运行函数名
func getFuncName(pc uintptr) string{
	var funcname  string
	funcname = runtime.FuncForPC(pc).Name()       // main.(*MyStruct).foo
	funcname = filepath.Ext(funcname)             // .foo
	funcname = strings.TrimPrefix(funcname, ".")  // foo
	return funcname
}

//获取设置的日志文件绝对路径
func getLogFilePath(name string) string{
	paths,_ := filepath.Split(name)
	return  paths
}

//获取设置的文件名
func getLogFileName() string{
	return  logFileName
}


//日志文件名
func getLogFile() string{
	if "" != logFullName {
		return logFullName
	}

	return  logFilePath+logFileName
}

//文件打开参数，默认是os.O_RDWR|os.O_CREATE|os.O_APPEND，读写，创建，增量写
func getOsFlag()  int {
	if 0 != osFlag {
		return osFlag
	}
	return os.O_RDWR|os.O_CREATE|os.O_APPEND
}

//FileMode，默认是0644
func getOsPerm() os.FileMode {
	if 0 != osPerm {
		return osPerm
	}

	return 0644
}

//日志打印模式
func getLogMode() int{
	if L_BUTT != logRecordMode {
		return logRecordMode
	}

	return logModeNormal
}

//日志打印级别，默认ERR
func getLogLevel() int{
	if LOG_LEVEL_BUTT != logRecordLevel {
		return logRecordLevel
	}

	return LOG_LEVEL_ERROR
}

//文件是否 存在
func checkFileExist(name string) bool {
	_, err := os.Stat(name)
	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}

	return true
}
//日志文件大小
func getFileSize(filename string) int64 {
	var result int64

	if true != checkFileExist(filename) {
		return 0
	}

	filepath.Walk(filename, func(path string, f os.FileInfo, err error) error {
		result = f.Size()
		return nil
	})

	return result/1024 +1
}

//日志文件超过最大内存的策略
func getLogOverPolicy() int {
	if logOverPolicy != L_RENAME_DATE {
		return logOverPolicy
	}

	return L_RENAME_DATE
}


func logFileOverProcess() {
	var newFileName string
	//全路径
	fileFull := getLogFile() // ./testlog.log
    //带后缀的文件名
	fileName := path.Base(fileFull) // testlog.log
	//文件后缀
	fileNameSuffix := path.Ext(fileName) // .log
	//纯文件名
	fileNameOnly :=strings.TrimSuffix(fileName, fileNameSuffix) // testlog
	//当前系统时间

	if logOverPolicy == L_RENAME_DATE {
		//t := time.Now()
		//fmt.Println(time.Now().Format("2006-01-02 15:04:05")) // 这是个奇葩,必须是这个时间点, 据说是go诞生之日, 记忆方法:6-1-2-3-4-5
		cur := fmt.Sprintf("%s", time.Now().Format("20060102150405"))
		newFileName = fmt.Sprintf("/%s_%s%s",fileNameOnly,cur,fileNameSuffix)

	}else if logOverPolicy == L_RENAME_INDEX {
		newFileName = fmt.Sprintf("/%s.%d%s",fileNameOnly,logIndex,fileNameSuffix)
		logIndex ++
	}else if logOverPolicy  ==  L_DELETE {
		err := os.Remove(fileFull)
		if err != nil {
			fmt.Println("remove file Error", err)
		}
		return
	}

	//文件重命名
	err := os.Rename(fileFull, getLogFilePath(fileFull) + newFileName)
	if err != nil {
		fmt.Println("reName Error", err)
	}

	return
}
