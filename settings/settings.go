package settings

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"
)

// Place your API keys at your main.go
var Key string
var Secret string
var DebugMode bool

const (
	PublicApiLink = "https://yobit.net/api/3/"
	TradeApiLink  = "https://yobit.io/tapi/"
)

// Check is a shared function for error checking
func Check(msg string, err error) {
	if err != nil {
		log.Println(msg, err, string(debug.Stack()))
		if DebugMode {
			writeToLog(msg, err)
		}
	}
}

// writeToLog writes errors info to logfile
func writeToLog(msg string, err error) {
	f, err := os.OpenFile("errors.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte(fmt.Sprintln(msg, err, string(debug.Stack())))); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
