package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
)

func CreateErrorLog(errLogName string, nrmLog *log.Logger) *os.File {
	file, err := os.Create(errLogName)

	if err != nil {
		nrmLog.Println(fmt.Scanf("Cannot Create Error Log File (Default : %v used)", defaultErrorLog.Name()))
		return defaultErrorLog
	}

	return file
}

func CreateFile(fileName string, errLog *log.Logger, nrmLog *log.Logger) *os.File {
	file, err := os.Create(fileName)

	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	fileName, line := f.FileLine(pc[0])

	if err != nil {
		nrmLog.Println(fmt.Scanf("Error occured in Line %v : Func : %v", line+5, fileName))
		nrmLog.Println("See error log for more details !")
		errLog.Fatalln(err.Error())
	}

	return file
}

// function return the data read from the external file if there is no error in reading it
// if there cencounter's error , then the func creates it's own file and places some default data inside
// then return's the data inside the func
func CheckExternalResource(exSrcName string, defData string, nrmLog *log.Logger, errLog *log.Logger) string {

	file, err := os.Open(exSrcName)

	if os.IsNotExist(err) {
		nrmLog.Println(fmt.Scanf("Requested file %v doesn't exist (default is created with default strings)", exSrcName))
		file = CreateFile(exSrcName, errLog, nrmLog)

		// now we need to write the default data inside the newly created file
		fileBuf := bufio.NewWriter(file)
		fileBuf.WriteString(defData)
		fileBuf.Flush()

		nrmLog.Println("Requested file is created automatically with default data")
		return defData
	}
	if err != nil {
		// Now the error is different from the recoverable one , so we need to exit
		nrmLog.Println("Non Recoverable error occured")
		nrmLog.Println("See error log for more details !")
		errLog.Fatalln(err.Error())
	}

	// We have file which is not automatically generated one , so we shoulod read data from it
	data, err := io.ReadAll(file)

	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	fileName, line := f.FileLine(pc[0])

	if err != nil {
		nrmLog.Println(fmt.Scanf("Error occured in Line %v : Func : %v", line+5, fileName))
		nrmLog.Println("See error log for more details !")
		errLog.Fatalln(err.Error())
	}

	return string(data)
}
