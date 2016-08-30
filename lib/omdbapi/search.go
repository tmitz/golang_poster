package omdbapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

func SearchMovies(title string, year int) (*Movie, error) {
	q := "t=" + url.QueryEscape(title)
	if year != 0 {
		q += "&y=" + url.QueryEscape(strconv.Itoa(year))
	}
	resp, err := http.Get(omdbURL + "?" + q)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result Movie
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
