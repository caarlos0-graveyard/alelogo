package alelogo_test

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/caarlos0/alelogo"
)

func TestSuccessLoginAndBalance(t *testing.T) {
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.URL.Path)
			if r.URL.Path == "/login/authenticate" {
				w.Write([]byte("OK"))
			}
			if r.URL.Path == "/user/card/preference/list" {
				f, _ := os.Open("./fakes/preferences_list.json")
				bts, _ := ioutil.ReadAll(f)
				w.Header().Add("Content-Type", "application/json")
				w.Write(bts)
			}
			if r.URL.Path == "/user/card/balance" {
				f, _ := os.Open("./fakes/balance.json")
				bts, _ := ioutil.ReadAll(f)
				w.Header().Add("Content-Type", "application/json")
				w.Write(bts)
			}
		}),
	)
	defer ts.Close()

	client, _ := alelogo.New("a", "b", alelogo.Config{BaseURL: ts.URL})
	cards, err := client.Cards()
	if err != nil {
		t.Error("Error getting cards", err)
	}
	if len(cards) != 1 {
		t.Error("cards len = ", len(cards))
	}
	details, err := client.Details(cards[0])
	if err != nil {
		t.Error("Error getting balance", err)
	}
	if details.Balance != "R$ 123,45" {
		t.Error("Card balance = ", details.Balance)
	}
}

func TestConnectionIssue(t *testing.T) {
	client, err := alelogo.New("a", "b", alelogo.Config{
		BaseURL: "http://google.com",
	})
	if err == nil {
		t.Error("Should have errored")
	}
	cards, err := client.Cards()
	if err == nil {
		t.Error("Should have errored")
	}
	if len(cards) != 0 {
		t.Error("cards len = ", len(cards))
	}
}
