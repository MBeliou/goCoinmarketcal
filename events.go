// Package gocoinmarketcal - Methods relative to the Event endpoint
package gocoinmarketcal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type Event struct {
	ID                int
	Title             string
	CoinsConcerned    Coins
	DateEvent         string
	CreatedDate       string
	Description       string
	Proof             string
	Source            string
	IsHot             string
	VoteCount         int
	PositiveVoteCount int
	Percentage        int
	Categories        Categories
	TipSymbol         string
	TipAddress        string
	TwitterAccount    string
	CanOccurBefore    bool
}

type Events []Event

func (s *Session) QueryEvents(params map[string]interface{}) (Events, error) {
	// Take care of the parameters
	urlVal := url.Values{}
	for k, v := range params {
		switch t := v.(type) {
		case string:
			urlVal.Add(k, t)
		case int:
			urlVal.Add(k, strconv.Itoa(t))
		default:
			log.Fatalf("Error: Bad entry %s", k)
		}
	}
	/*
		page number integer
		page max integer
		dateRangeStart string
		dateRangeEnd string
		coins string - coins id
		categories strings - see /api/categories
		sortBy - string (created_desc, hot_events)
		showOnly - string (hot_events)
		showMetaData - string (true) - useless
	*/
	endpoint := "v1/events?"
	urlQuery := fmt.Sprintf("%s%saccess_token=%s&%s", apiurl, endpoint, s.accessToken, urlVal.Encode())
	fmt.Printf("urlQuery: %s\n", urlQuery)
	resp, err := http.Get(urlQuery)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	log.Printf("Resp: %s", body)
	events := &Events{}

	err = json.Unmarshal(body, events)
	if err != nil {
		return *events, err
	}

	return *events, nil
}
