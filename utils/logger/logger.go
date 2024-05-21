package logger

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func generateFileName() string {
	t := time.Now()
	randomDigits := generateRandomDigits(6)
	return fmt.Sprintf("%s_%s.log", t.Format("20060102_150405"), randomDigits)
}

func generateRandomDigits(n int) string {
	randGenerator := rand.New(rand.NewSource(time.Now().Unix()))
	digits := make([]byte, n)
	for i := 0; i < n; i++ {
		digits[i] = byte(randGenerator.Intn(10) + '0')
	}
	return string(digits)
}

func init() {
	fileName := generateFileName()
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	mw := io.MultiWriter(os.Stdout, file)
	mwError := io.MultiWriter(os.Stderr, file)
	InfoLogger = log.New(mw, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(mw, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(mwError, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
