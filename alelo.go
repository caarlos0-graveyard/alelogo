package alelogo

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

const url = "https://www.meualelo.com.br/meualelo.services/rest"

// ErrAuth happens on authentication failures
var ErrAuth = errors.New("Authentication failure")

// Login a user and return the cookies
func Login(cpf, pwd string) (cookies []*http.Cookie, err error) {
	pwd = base64.StdEncoding.EncodeToString([]byte(pwd))
	json := "{\"cpf\":\"" + cpf + "\",\"pwd\":\"" + pwd + "\",\"captchaResponse\":\"\"}"
	client := &http.Client{}
	req, err := http.NewRequest("POST", url+"/login/authenticate", strings.NewReader(json))
	if err != nil {
		return cookies, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return cookies, ErrAuth
	}
	return resp.Cookies(), err
}

// Balance get the user card's balances
func Balance(cookies []*http.Cookie) (cards []Card, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url+"/user/card/preference/list", nil)
	if err != nil {
		return cards, err
	}
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return cards, ErrAuth
	}
	var preferences preferencesJSON
	err = json.NewDecoder(resp.Body).Decode(&preferences)
	return preferences.List, err
}

// Card type
type Card struct {
	ID      string `json:"cardId"`
	Title   string `json:"title"`
	Balance string `json:"balance"`
}

type preferencesJSON struct {
	UID  string `json:"uid"`
	List []Card `json:"cardList"`
}
