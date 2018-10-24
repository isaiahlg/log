package seq

import (
	"net/http"
)

type SeqHook struct{
	BaseUrl string
	ApiKey string
}

type event struct{

}

type seqLog []event

func (seqhook *SeqHook) Fatal(msg string, s interface{}){

}

func (seqhook *SeqHook) Error(msg string, s interface{}){

}

func (seqhook *SeqHook) log(ev event){
	var sl seqLog
	sc := seqClient{
		baseUrl: seqhook.BaseUrl,
		apiKey: seqhook.ApiKey,
	}
	sc.send(sl, &http.Client{})
}