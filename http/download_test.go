package http

import (
	"bytes"
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

type httpSuite struct {
	suite.Suite
}

// Custom type that allows setting the func that our Mock Do func will run instead
type MockGetType func(url string) (*http.Response, error)

// MockClient is the mock client
type MockClient struct {
	MockGet MockGetType
}

// Overriding what the Do function should "get" in our MockClient
// Source: https://levelup.gitconnected.com/mocking-outbound-http-calls-in-golang-9e5a044c2555
func (m *MockClient) Get(url string) (*http.Response, error) {
	return m.MockGet(url)
}

func TestHTTPSuite(t *testing.T) {
	suite.Run(t, new(httpSuite))
}

func removeFile(fileName string) error {
	if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
		return nil
	}
	return os.Remove(fileName)
}

func (s *httpSuite) TestCreateFile() {
	// create a new reader with that JSON
	var output string = uuid.NewString()
	var fileName string = "http.test"
	r := ioutil.NopCloser(bytes.NewReader([]byte(output)))
	Client = &MockClient{
		MockGet: func(uri string) (*http.Response, error) {
			return &http.Response{
				StatusCode:    200,
				Body:          r,
				ContentLength: int64(len([]byte(output))),
			}, nil
		},
	}
	err := DownloadFile("http://example.com", fileName)
	defer removeFile(fileName)
	s.NoError(err)
	outputBytes, err := ioutil.ReadFile(fileName)
	s.NoError(err)
	s.Equal(string(outputBytes), output)

}
