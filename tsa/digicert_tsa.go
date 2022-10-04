package tsa

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type DigiCertClient struct {
	client *http.Client
}

func NewDigiCertClient() *DigiCertClient {
	return &DigiCertClient{
		client: &http.Client{},
	}
}

func (c *DigiCertClient) TSARequest(tsq []byte) (tsr []byte, err error) {
	url := "http://timestamp.digicert.com"
	method := "POST"

	payload := bytes.NewReader(tsq)

	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/timestamp-query")
	req.Header.Add("Accept", "application/timestamp-reply, application/timestamp-response")
	req.Header.Add("Pragma", "no-cache")

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	tsr, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return tsr, nil
}
