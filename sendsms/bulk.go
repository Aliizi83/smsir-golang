package sendsms

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
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	return resp, err
}
