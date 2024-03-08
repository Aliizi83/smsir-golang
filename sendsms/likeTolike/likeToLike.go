package likeTolike

import (
	"bytes"
	"encoding/json"
	"net/http"
)

var url string = "https://api.sms.ir/v1/send/likeToLike"

type LikeToLike struct {
	ApiKey       string
	LineNumber   string
	MessageTexts []string
	Mobiles      []string
}

func New(key, line string, messageTexts, mobiles []string) *LikeToLike {
	return &LikeToLike{
		key,
		line,
		messageTexts,
		mobiles,
	}
}

func (l *LikeToLike) Send() (*http.Response, error) {
	body, err := json.Marshal(l)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-KEY", l.ApiKey)

	client := new(http.Client)
	res, err := client.Do(req)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}
	return res, nil
}
