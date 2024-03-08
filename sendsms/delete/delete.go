package delete

import "net/http"

var url string = "https://api.sms.ir/v1/send/scheduled/"

func Delete(apikey, packid string) (*http.Response, error) {
	url = url + packid

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-API-KEY", apikey)
	client := new(http.Client)
	res, err := client.Do(req)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}
	return res, nil

}
