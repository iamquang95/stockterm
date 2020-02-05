package crawler

import (
	"io/ioutil"
	"net/http"
)

// GetHTML gets HTML content of a given url
func GetHTML(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)
	return html, err
}
