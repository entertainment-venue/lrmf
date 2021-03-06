package lrmf

import (
	"io/ioutil"
	"log"
)

// StdLogger is used to log error messages.
type StdLogger interface {
	Print(v ...interface{})
	Printf(format string, v ...interface{})
	Println(v ...interface{})
}

var Logger StdLogger = log.New(ioutil.Discard, "[lrmf] ", log.LstdFlags)
