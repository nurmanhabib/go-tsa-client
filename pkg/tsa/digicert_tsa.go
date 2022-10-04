package tsa

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/nurmanhabib/go-tsa-client/domain/entity"
)

type DigiCertClient struct {
	client *http.Client
}

func NewDigiCertClient() *DigiCertClient {
	return &DigiCertClient{
		client: &http.Client{},
	}
}

func (c *DigiCertClient) TSARequest(tsq []byte) (reply *entity.TSReply, err error) {
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

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	tsr := &entity.TSReply{
		Data:          data,
		ContentLength: res.ContentLength,
		Date:          res.Header.Get("Date"),
	}

	return tsr, nil
}
