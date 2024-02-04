package log

import (
	"log"
	"time"
)

const (
	// YYYY-MM-DD: 2022-03-23
	YYYYMMDD = "2006-01-02"
	// 24h hh:mm:ss: 14:23:20
	HHMMSS24h = "15:04:05"
	// 12h hh:mm:ss: 2:23:20 PM
	HHMMSS12h = "3:04:05 PM"
	// text date: March 23, 2022
	TextDate = "January 2, 2006"
	// text date with weekday: Wednesday, March 23, 2022
	TextDateWithWeekday = "Monday, January 2, 2006"
	// abbreviated text date: Mar 23 Wed
	AbbrTextDate = "Jan 2 Mon"
)

func Info(msg string) {
	flags := log.Lshortfile
	// Log without log level
	datetime := time.Now().UTC().Format(YYYYMMDD+" "+HHMMSS24h) + ": "
	log.SetFlags(flags)
	// INFO log level
	log.SetPrefix("INFO: " + datetime)
	log.Println(msg)
}

func Warn(msg string) {
	flags := log.Lshortfile
	// Log without log level
	datetime := time.Now().UTC().Format(YYYYMMDD+" "+HHMMSS24h) + ": "
	log.SetFlags(flags)
	// INFO log level
	log.SetPrefix("WARN: " + datetime)
	log.Println(msg)
}

func Error(msg string) {
	flags := log.Lshortfile
	// Log without log level
	datetime := time.Now().UTC().Format(YYYYMMDD+" "+HHMMSS24h) + ": "
	log.SetFlags(flags)
	// INFO log level
	log.SetPrefix("ERROR: " + datetime)
	log.Println(msg)
}

func Fatalf(msg string, err ...any) {
	log.Fatalf(msg, err...)
}
