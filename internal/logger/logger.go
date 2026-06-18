package logger

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	Info  *log.Logger
	Error *log.Logger
	SIP   *log.Logger
}

func New() *Logger {

	file, err := os.OpenFile(
		"client.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666,
	)

	if err != nil {
		log.Fatal(err)
	}

	writer := io.MultiWriter(
		os.Stdout,
		file,
	)

	return &Logger{
		Info: log.New(
			writer,
			"[INFO] ",
			log.Ldate|log.Ltime|log.Lshortfile,
		),

		Error: log.New(
			writer,
			"[ERROR] ",
			log.Ldate|log.Ltime|log.Lshortfile,
		),

		SIP: log.New(
			writer,
			"[SIP] ",
			log.Ldate|log.Ltime,
		),
	}
}
