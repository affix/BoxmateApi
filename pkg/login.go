package boxmate

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Response struct {
	Success bool   `json:"success"`
	ApiKey  string `json:"member_token"`
}

// BoxmateLogin Log in to boxmate and return the user api token
// returned token does not expire
func BoxmateLogin(email string, password string) (Response, error) {
	client := http.Client{}
	resp, err := client.PostForm("https://api.boxmateapp.co.uk/member/authenticate", url.Values{
		"Member_Email":    {email},
		"Member_Password": {password},
	})
	if err != nil {
		return Response{}, err
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Response{}, err
	}
	var result Response
	json.Unmarshal(data, &result)
	return result, nil
}
