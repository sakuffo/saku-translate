package cli

import (
	"log"
	"net/http"
	"sync"
)

type RequestBody struct {
	SourceLang string
	TargetLang string
	SourceText string
}

const translateUrl = "https://translate.googleapis.com/translate_a/single"

func RequestTranslate(reqBody *RequestBody, str chan string, wg *sync.WaitGroup) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", translateUrl, nil)

	query := req.URL.Query()
	query.Add("client", "gtx")
	query.Add("sl", reqBody.SourceLang)
	query.Add("tl", reqBody.TargetLang)
	query.Add("dt", "t")
	query.Add("q", reqBody.SourceText)

	req.URL.RawQuery = query.Encode()
	if err != nil {
		log.Fatal("1. There was a problem: %s", err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal("2. There was a problem: %s", err)
	}

	defer res.Body.Close()

}
