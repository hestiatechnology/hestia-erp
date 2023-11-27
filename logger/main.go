package logger

import (
	"log"
	"os"
	"sync"
)

var (
	once    sync.Once
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func initializeLogger() {
	file := os.Stdout
	// file, err := os.OpenFile("hestia.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// if err != nil {
	// 	log.Fatal("Unable to open log file, error: ", err)
	// }

	Info = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// InitLogger initializes the logger once
func InitLogger() {
	once.Do(initializeLogger)
}
