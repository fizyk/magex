package http

import (
	"io"
	"net/http"
	"os"
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

	// Create the file, but give it a tmp file extension, this means we won't overwrite a
	// file until it's downloaded, but we'll remove the tmp extension once downloaded.
	out, err := os.Create(filename + ".tmp")
	if err != nil {
		return err
	}

	resp, err := Client.Get(uri)
	if err != nil {
		out.Close()
		return err
	}
	defer resp.Body.Close()
	// Create our progress reporter and pass it to be used alongside our writer
	counter := &WriteCounter{}
	if _, err = io.Copy(out, io.TeeReader(resp.Body, counter)); err != nil {
		out.Close()
		return err
	}

	out.Close()
	if err = os.Rename(filename+".tmp", filename); err != nil {
		return err
	}
	return nil
}
