package hwr

import (
	"encoding/json"
	"net/http"
	"net/url"
)

const (
	login = "https://webmail.stud.hwr-berlin.de/appsuite/api/login?"
)

type HWRError struct {
	Error string `json:"error"`
}

func ValidateUser(email, password string) (bool, error) {
	hwrError := &HWRError{}
	form := url.Values{}
	form.Add("action", "login")
	form.Add("name", email)
	form.Add("password", password)
	resp, err := http.PostForm(login, form)

	if err != nil {
		return false, err
	}

	dec := json.NewDecoder(resp.Body)

	err = dec.Decode(hwrError)

	if err != nil {
		return false, err
	}

	return hwrError.Error == "", nil
}
