package im_log

import "fmt"

func Info(v ...interface{}) {
	fmt.Printf(v[0].(string), v[1:])
}

func Warn(v ...interface{}) {
	fmt.Printf(v[0].(string), v[1:])
}

func Error(v ...interface{}) {
	fmt.Printf(v[0].(string), v[1:])
}
