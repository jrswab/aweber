package aweber

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
)

// Aweber contains the necessary data to interact with the Aweber API.
type Aweber struct {
	Client *http.Client

	baseURL string
	config  *oauth2.Config
}

// NewAweber creates and returns a new Aweber struct used for calls to the Aweber API
func NewAweber(accountID string) (*Aweber, error) {
	var a = new(Aweber)
	a.baseURL = fmt.Sprintf("https://api.aweber.com/1.0/accounts/%s/", accountID)

	err := a.getOauth2Config()
	if err != nil {
		return nil, err
	}

	a.getClient()

	return a, nil
}

func (a *Aweber) getOauth2Config() error {
	_, isFound := os.LookupEnv("CLIENT_ID")
	if isFound {
		return errors.New("environment variable CLIENT_ID is not found")
	}
	_, isFound = os.LookupEnv("CLIENT_SECRET")
	if isFound {
		return errors.New("environment variable CLIENT_SECRET is not found")
	}

	a.config = &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Scopes:       []string{"account.read", "list.read"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://auth.aweber.com/oauth2/authorize",
			TokenURL: "https://auth.aweber.com/oauth2/tokeen",
		},
	}

	return nil
}

func (a *Aweber) getClient() {
	tokeFile := "token.json"
	toke, err := tokenFromFile(tokeFile)
	if err != nil {
		toke = tokenFromWeb(a.config)
		saveToken(tokeFile, toke)
	}

	config := a.config
	a.Client = config.Client(context.Background(), toke)
}

func tokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-tokeen", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code here and press enter: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	toke, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve tokeen from web: %v", err)
	}
	return toke
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	toke := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(toke)
	return toke, err
}

func saveToken(path string, tokeen *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth tokeen: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(tokeen)
}
