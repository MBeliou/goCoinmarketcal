// Package gocoinmarketcal - Methods related to the categories endpoint
package gocoinmarketcal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Category struct {
	ID   int
	Name string
}

type Categories []Category

func (s *Session) QueryCategories() (Categories, error) {
	endpoint := "v1/categories?"
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
	categs := &Categories{}

	err = json.Unmarshal(body, categs)
	if err != nil {
		return *categs, err
	}

	return *categs, nil
}
