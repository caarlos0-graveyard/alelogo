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
	f, _ := os.Open("./fakes/balance.json")
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.URL.Path)
			if r.URL.Path == "/login/authenticate" {
				w.Write([]byte("OK"))
			}
			if r.URL.Path == "/user/card/preference/list" {
				bts, _ := ioutil.ReadAll(f)
				w.Header().Add("Content-Type", "application/json")
				w.Write(bts)
			}
		}),
	)
	defer ts.Close()

	client, _ := alelogo.New("a", "b", alelogo.Config{BaseURL: ts.URL})
	cards, err := client.Balance()
	if err != nil {
		t.Error("Error getting balance", err)
	}
	if len(cards) != 1 {
		t.Error("cards len = ", len(cards))
	}
	if cards[0].Balance != "R$ 123,45" {
		t.Error("Card balance = ", cards[0].Balance)
	}
}

func TestConnectionIssue(t *testing.T) {
	client, err := alelogo.New("a", "b", alelogo.Config{
		BaseURL: "http://google.com",
	})
	if err == nil {
		t.Error("Should have errored")
	}
	cards, err := client.Balance()
	if err == nil {
		t.Error("Should have errored")
	}
	if len(cards) != 0 {
		t.Error("cards len = ", len(cards))
	}
}
