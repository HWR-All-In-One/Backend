package hwr

import (
	"net/http"
	"net/url"
)

const (
	login = "https://webmail.stud.hwr-berlin.de/appsuite/api/login?"
)

func ValidateUser(username, password string) (bool, error) {
	form := url.Values{}
	form.Add("action", "login")
	form.Add("name", username)
	form.Add("password", password)
	resp, err := http.PostForm(login, form)

	if err != nil {
		return false, err
	}

	return resp.StatusCode == 200, nil
}
