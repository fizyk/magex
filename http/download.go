package http

import (
	"io/ioutil"
	"net/http"
)

// HTTPClient interface
type HTTPClient interface {
	Get(url string) (*http.Response, error)
}

var (
	Client HTTPClient
)

func init() {
	Client = &http.Client{}
}

func DownloadFile(uri, filename string) error {
	resp, err := Client.Get(uri)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if body, err := ioutil.ReadAll(resp.Body); err != nil {
		return err
		// Write downloaded http locally
	} else if err := ioutil.WriteFile(filename, body, 0644); err != nil {
		return err
	}
	return nil
}
