// Package gocoinmarketcal is the unofficial client to the Coinmarketcal API
package gocoinmarketcal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	apiurl = "https://api.coinmarketcal.com/"
)

// Session manages the communication with the coinmarketcal API
type Session struct {
	accessToken string
}

// NewSession uses the apiKey and apiSecret to query an authentication token
// Returns a Session struct
func NewSession(apiKey string, apiSecret string) (*Session, error) {
	endpoint := "oauth/v2/token?"
	grantType := "grant_type=client_credentials"
	clientID := fmt.Sprintf("client_id=%s", apiKey)
	clientSecret := fmt.Sprintf("client_secret=%s", apiSecret)

	urlQuery := fmt.Sprintf("%s%s%s&%s&%s", apiurl, endpoint, grantType, clientID, clientSecret)
	log.Printf("UrlQuery: %s", urlQuery)
	resp, err := http.Get(urlQuery)
	if err != nil {
		log.Printf("NewSession: %s", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("New Session: Read body - %s", err)
		return nil, err
	}
	fmt.Printf("Body: %s\n", body)
	session := &Session{}

	var rep map[string]interface{}
	err = json.Unmarshal(body, &rep)
	if err != nil {
		log.Printf("New Session: Unmarshal - %s", err)
	}

	session.accessToken = rep["access_token"].(string)

	return session, nil
}
