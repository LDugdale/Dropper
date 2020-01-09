package logger

import (
    //"io/ioutil"
    "log"
    "os"
)

type ILogger interface {
	LogTrace(v ...interface{})	
	LogInfo(message string)	
	LogWarning(message string)	
	LogError(message string) 
}

type Logger struct {
    Trace   *log.Logger
    Info    *log.Logger
    Warning *log.Logger
    Error   *log.Logger
}

func NewLogger() *Logger {

	traceHandle := os.Stdout
	infoHandle := os.Stdout
	warningHandle := os.Stdout
	errorHandle := os.Stderr

	l := &Logger {
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

func (l Logger) LogTrace(v ...interface{}) {
	l.Trace.Println(v)
}

func (l Logger) LogInfo(message string) {
	l.Info.Println(message)
}

func (l Logger) LogWarning(message string) {
	l.Warning.Println(message)
}

func (l Logger) LogError(message string) {
	l.Error.Println(message)
}
