package seq

import (
	"net/http"
)

type seqClient struct{
	baseUrl string
	apiKey string
}

func (sc *seqClient) send(log seqLog, client *http.Client) error {
	return nil
}