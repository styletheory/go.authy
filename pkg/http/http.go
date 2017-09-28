package http

import (
    "net/http"
    "time"
    "io"
    // "log"
)

type Http interface {
  Get(url string, contentType string, headersMap map[string]string, queryParams map[string]string) (*http.Response, error)
  Post(url string, contentType string, body io.Reader, headersMap map[string]string, queryParams map[string]string) (*http.Response, error)
}

type DefaultHttpClient struct {
    Client  *http.Client
}

func NewDefaultHttpClient() *DefaultHttpClient {
    return &DefaultHttpClient {
        Client: &http.Client {
            Timeout: time.Second * 10,
        },
    }
}

func (h *DefaultHttpClient) Get(url string, contentType string, headersMap map[string]string, queryParams map[string]string) (*http.Response, error) {
    if len(queryParams) > 0 {
        url = url + "?"
    }
    for k, v := range queryParams {
        url = url + k + "=" + v + "&"
    }
    url = url[:len(url) - 1]

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, err
    }
    req.Header.Add("Content-Type", contentType)
    for k,v := range headersMap {
        req.Header.Add(k, v)
    }

    return h.Client.Do(req)
}

func (h *DefaultHttpClient) Post(url string, contentType string, body io.Reader, headersMap map[string]string, queryParams map[string]string) (*http.Response, error) {
    // req, err := http.NewRequest("POST", url, strings.NewReader(body.Encode()))
    if len(queryParams) > 0 {
        url = url + "?"
    }
    for k, v := range queryParams {
        url = url + k + "=" + v + "&"
    }
    url = url[:len(url) - 1]

    // req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
    req, err := http.NewRequest("POST", url, body)
    if err != nil {
        return nil, err
    }
    req.Header.Add("Content-Type", contentType)
    for k,v := range headersMap {
        req.Header.Add(k, v)
    }

    return h.Client.Do(req)
}