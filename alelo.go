package alelogo

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/cookiejar"
	"strings"
	"time"
)

// Config of the client
type Config struct {
	BaseURL string
	Timeout int
}

// ErrAuth happens on authentication failures
var ErrAuth = errors.New("Authentication failure")

// ErrDumbass happens when random shit happens
var ErrDumbass = errors.New("Random shit happened within Alelo API, try again")

// DefaultConfig for the client
var DefaultConfig = Config{
	BaseURL: "https://www.meualelo.com.br/meualelo.services/rest",
	Timeout: 30,
}

// Client for Alelo API
type Client struct {
	http.Client
	BaseURL string
}

// New client with defaults
func New(cpf, pwd string, configs ...Config) (*Client, error) {
	var config = DefaultConfig
	if len(configs) > 0 {
		config = configs[0]
	}
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	client := &Client{
		http.Client{
			Timeout: time.Duration(config.Timeout) * time.Second,
			Jar:     jar,
		}, config.BaseURL,
	}
	return client, client.login(cpf, pwd)
}

func (client *Client) login(cpf, pwd string) (err error) {
	pwd = base64.StdEncoding.EncodeToString([]byte(pwd))
	json := "{\"cpf\":\"" + cpf +
		"\",\"pwd\":\"" + pwd +
		"\",\"captchaResponse\":\"\"}"
	req, err := http.NewRequest(
		"POST",
		client.BaseURL+"/login/authenticate",
		strings.NewReader(json),
	)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return ErrAuth
	}
	return err
}

// Cards get the user card's balances
func (client *Client) Cards() (cards []Card, err error) {
	req, err := http.NewRequest(
		"GET",
		client.BaseURL+"/user/card/preference/list",
		nil,
	)
	if err != nil {
		return cards, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return cards, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		// this should never happen
		return cards, ErrDumbass
	}
	var preferences preferencesJSON
	err = json.NewDecoder(resp.Body).Decode(&preferences)
	return preferences.List, err
}

func (client *Client) Details(card Card) (CardDetails, error) {
	var details CardDetails
	req, err := http.NewRequest(
		"GET",
		client.BaseURL+"/user/card/balance?selectedCardNumberId="+card.ID,
		nil,
	)
	if err != nil {
		return details, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return details, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		// this should never happen
		return details, ErrDumbass
	}
	err = json.NewDecoder(resp.Body).Decode(&details)
	return details, err
}

// Card type
type Card struct {
	ID    string `json:"cardId"`
	Title string `json:"title"`
}

type preferencesJSON struct {
	UID  string `json:"uid"`
	List []Card `json:"cardList"`
}

// CardDetails type
type CardDetails struct {
	Balance string `json:"balance"`
	Name    string `json:"productName"`
	Type    string `json:"cardType"`
}
