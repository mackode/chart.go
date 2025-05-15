package main

import (
	"fmt"
	"github.com/mschilli/go-murmur"
	"github.com/tidwall/gjson"
	"io/ioutil"

	"net/http"
	"net/url"
)

type MovieMeta struct {
	Title string
	Rating string
}

func omdbFetch(title string) (MovieMeta, error) {
	mmeta := MovieMeta{}
	baseURL := "http://www.omdbapi.com/"
	apiURL, _ := url.Parse(baseURL)

	apiKey, err := murmur.NewMurmur().Lookup("omdb-api-key")
	if err != nil {
		return mmeta, err
	}

	q := url.Values{}
	q.Add("t", title)
	q.Add("apikey", apiKey)
	apiURL.RawQuery = q.Encode()
	resp, err := http.Get(apiURL.String())
	if err != nil {
		return mmeta, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return mmeta, fmt.Errorf("HTTP status: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return mmeta, err
	}

	jsonErr := gjson.GetBytes(body, "Error")
	if jsonErr.Exists() {
		return mmeta, fmt.Errorf("%s not found in omdb", title)
	}

	mmeta.Title = title
	mmeta.Rating = gjson.Get(string(body), `Ratings.#(Source=="Internet Movie Database").Value`).String()

	return mmeta, nil
}