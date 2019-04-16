package logger

import (
	"fmt"
	"time"
)

var timeFormat = "2006/01/02 15:04:05.999"

//Info 级别
func Info(format string, args ...interface{}) {
	output(fmt.Sprintf(format, args...))
}

//Error 级别
func Error(format string, args ...interface{}) {
	output(fmt.Sprintf(format, args...))
}

//Fatal 级别
func Fatal(format string, args ...interface{}) {
	output(fmt.Sprintf(format, args...))
	panic(fmt.Sprintf(format, args...))
}

func output(s string) {
	fmt.Println(time.Now().Format(timeFormat) + " " + s)
}
