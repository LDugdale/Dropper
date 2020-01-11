package log

import (
    //"io/ioutil"
    "log"
    "os"
)

type Logger interface {
	LogTrace(v ...interface{})	
	LogInfo(v ...interface{})	
	LogWarning(v ...interface{})	
	LogError(v ...interface{}) 
}

type GoLogger struct {
    Trace   *log.Logger
    Info    *log.Logger
    Warning *log.Logger
    Error   *log.Logger
}

func NewLogger() Logger {

	traceHandle := os.Stdout
	infoHandle := os.Stdout
	warningHandle := os.Stdout
	errorHandle := os.Stderr

	l := &GoLogger {
	Trace: log.New(traceHandle,
        "TRACE: ",
        log.Ldate|log.Ltime|log.Lshortfile),

    Info: log.New(infoHandle,
        "INFO: ",
        log.Ldate|log.Ltime|log.Lshortfile),

    Warning: log.New(warningHandle,
        "WARNING: ",
        log.Ldate|log.Ltime|log.Lshortfile),

    Error: log.New(errorHandle,
        "ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile),
	}

	return l
}

func (l GoLogger) LogTrace(v ...interface{}) {
	l.Trace.Println(v)
}

func (l GoLogger) LogInfo(v ...interface{}) {
	l.Info.Println(v)
}

func (l GoLogger) LogWarning(v ...interface{}) {
	l.Warning.Println(v)
}

func (l GoLogger) LogError(v ...interface{}) {
	l.Error.Println(v)
}
