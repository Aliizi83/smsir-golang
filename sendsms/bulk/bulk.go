package bulk

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

var (
	url string = "https://api.sms.ir/v1/send/bulk"
)

type Bulk struct {
	ApiKey       string
	LineNumber   string
	MessageText  string
	Mobiles      []string
	sendDateTime time.Time
}

func (b *Bulk) New(key, line, message string, mobiles []string, send time.Time) *Bulk {
	return &Bulk{
		key,
		line,
		message,
		mobiles,
		send,
	}
}

func (b *Bulk) Send() (*http.Response, error) {
	body, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-KEY", b.ApiKey)

	client := new(http.Client)
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	return resp, nil

}
