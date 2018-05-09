// Package gocoinmarketcal - Methods related to the coins endpoint
package gocoinmarketcal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Coin struct {
	ID     string
	Name   string
	Symbol string
}

type Coins []Coin

func (s *Session) QueryCoins() (Coins, error) {
	endpoint := "v1/coins?"
	access := fmt.Sprintf("access_token=%s", s.accessToken)
	urlQuery := fmt.Sprintf("%s%s%s", apiurl, endpoint, access)

	resp, err := http.Get(urlQuery)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	coins := &Coins{}

	err = json.Unmarshal(body, coins)
	if err != nil {
		return *coins, err
	}

	return *coins, nil
}
