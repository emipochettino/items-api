package logger

import (
	"bytes"
	"fmt"
	"log"
	"time"

	"github.com/emipochettino/items-api-go/context"
)

const timeFormat = "2006-01-02T15:04:05.000"

func Info(message string) {
	t := time.Now().UTC().Format(timeFormat)
	log.Printf("%v - %s - %s", t, createKeyValuePairs(context.Context.ContextMap), message)
}

func Error(message string) {
	t := time.Now().UTC().Format(timeFormat)
	log.Panicf(t, createKeyValuePairs(context.Context.ContextMap), message)
}

func createKeyValuePairs(m map[string]string) string {
	b := new(bytes.Buffer)
	for key, value := range m {
		fmt.Fprintf(b, "[%s:%s]", key, value)

	}

	return b.String()
}
