package adviceslip

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/marcocharlie/advice-app/api/config"
)

type AdviceSlipResponse struct {
	//TotalResults int    `json:"total_results"`
	Query string `json:"query"`
	Slips []Slip `json:"slips"`
}

type Slip struct {
	ID     int    `json:"id"`
	Advice string `json:"advice"`
	Date   string `json:"date"`
}

// searchAdviceSlips searches for advices based on topic
func SearchAdviceSlips(topic string) (*[]Slip, error) {

	var sleep int = 1
	var content []byte
	var adviceSlipResponse AdviceSlipResponse

	searchURL := fmt.Sprintf("%s%s", config.AdviceSlipBaseURL, topic)
	log.Printf("searching advices for topic: %s", topic)

	for {
		res, err := http.Get(searchURL)
		if err != nil {
			log.Printf("searching advices failed, retrying after %d seconds\n", sleep)
			sleep++
			if sleep > 5 {
				return nil, err
			}
			time.Sleep(time.Duration(sleep) * time.Second)
		}

		content, err = ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			log.Println(err)
			return nil, err
		}
		break
	}

	err := json.Unmarshal(content, &adviceSlipResponse)
	if err != nil {
		return nil, err
	}

	return &adviceSlipResponse.Slips, nil
}
