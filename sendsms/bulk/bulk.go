package bulk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var (
	url string = "https://api.sms.ir/v1/send/bulk"
)

type Bulk struct {
	ApiKey       string     `json:"-"`
	LineNumber   string     `json:"lineNumber"`
	MessageText  string     `json:"messageText"`
	Mobiles      []string   `json:"mobiles"`
	SendDateTime *time.Time `json:"sendDateTime"`
}

func NewBulk(key, line, message string, mobiles []string, send *time.Time) *Bulk {
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

	fmt.Println(string(body))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-KEY", b.ApiKey)

	client := new(http.Client)
	res, err := client.Do(req)
	defer res.Body.Close()

	if err != nil {
		return nil, err
	}
	return res, nil

}
