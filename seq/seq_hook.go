package seq

import (
	"net/http"
	"github.com/fatih/structs"
	"time"
)

type SeqHook struct{
	BaseUrl string
	ApiKey string
}

type event struct{
	Timestamp       time.Time
	Level           string
	MessageTemplate string
	Properties      map[string]interface{}
	Exception       string
}

type seqEvent []event

func (seqhook *SeqHook) Fatal(msg string, s interface{}){
	seqhook.Error(msg, s)
	panic(msg)
}

func (seqhook *SeqHook) Error(msg string, s interface{}){
	var m map[string]interface{}
	if s != nil {
		m = structs.Map(s)
	}

	event := event{
		Timestamp:       time.Now().UTC(),
		Level:           "Error",
		MessageTemplate: msg,
		Properties:      m,
		Exception:       "",
	}
	seqhook.log(event)
}

func (seqhook *SeqHook) log(ev event){
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