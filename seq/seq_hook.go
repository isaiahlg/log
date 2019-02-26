package seq

import (
	"net/http"
	"runtime"
	"time"

	"github.com/fatih/structs"
)

type SeqHook struct {
	BaseUrl string
	ApiKey  string
}

type event struct {
	Timestamp       time.Time
	Level           string
	MessageTemplate string
	Properties      map[string]interface{}
	Exception       string
}

type seqEvent []event

func (seqhook *SeqHook) Info(msg string, s interface{}) {
	m := mapProps(s)

	event := event{
		Timestamp:       time.Now().UTC(),
		Level:           "Information",
		MessageTemplate: msg,
		Properties:      m,
	}
	seqhook.log(event)
}

func (seqhook *SeqHook) Warning(msg string, s interface{}) {
	m := mapProps(s)

	event := event{
		Timestamp:       time.Now().UTC(),
		Level:           "Warning",
		MessageTemplate: msg,
		Properties:      m,
	}
	seqhook.log(event)
}

func (seqhook *SeqHook) Fatal(msg string, s interface{}) {
	seqhook.Error(msg, s)
	panic(msg)
}

func (seqhook *SeqHook) Error(msg string, s interface{}) {
	m := mapProps(s)

	trace := make([]byte, 1024)
	runtime.Stack(trace, true)

	event := event{
		Timestamp:       time.Now().UTC(),
		Level:           "Error",
		MessageTemplate: msg,
		Properties:      m,
		Exception:       string(trace),
	}
	seqhook.log(event)
}

func (seqhook *SeqHook) log(ev event) {
	var httpClient = &http.Client{
		Transport: &http.Transport{
			TLSHandshakeTimeout: 30 * time.Second,
		},
	}

	sc := seqClient{
		baseUrl: seqhook.BaseUrl,
		apiKey:  seqhook.ApiKey,
	}

	var ee seqEvent
	ee = append(ee, ev)

	sc.send(&ee, httpClient)
}

func mapProps(s interface{}) map[string]interface{} {
	var m map[string]interface{}
	if s != nil {
		switch t := s.(type) {
		case string:
			m = make(map[string]interface{})
			v := s.(string)
			m["key"] = v

		default:
			println(t)
			m = structs.Map(s)
		}
	}
	return m
}
