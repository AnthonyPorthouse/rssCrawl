package http

import "net/http"

var client = &http.Client{}

func GetFeed(URL string) (*http.Response, error) {
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "RSSCrawl v0.0.1")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
