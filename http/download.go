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

// DownloadFile downloads file from the internet, based on the uri.
// based on the https://golangcode.com/download-a-file-with-progress/ and comments
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
	counter := NewWriteCounter(int(resp.ContentLength))
	counter.Start()
	if _, err = io.Copy(out, io.TeeReader(resp.Body, counter)); err != nil {
		out.Close()
		return err
	}
	counter.Finish()

	out.Close()
	if err = os.Rename(filename+".tmp", filename); err != nil {
		return err
	}
	return nil
}
