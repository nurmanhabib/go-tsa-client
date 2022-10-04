package tsa

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/nurmanhabib/go-tsa-client/domain/entity"
)

type TecxoftTSA struct {
	username string
	password string
	client   *http.Client
}

func NewTecxoftTSA(username, password string) *TecxoftTSA {
	return &TecxoftTSA{
		username: username,
		password: password,
		client:   &http.Client{},
	}
}

func (t *TecxoftTSA) TSARequest(tsq []byte) (*entity.TSReply, error) {
	url := "http://tsa.tecxoft.com/test"
	method := "POST"

	payload := bytes.NewReader(tsq)

	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, err
	}

	auth := fmt.Sprintf("%s:%s", t.username, t.password)
	encodedAuth := base64.StdEncoding.EncodeToString([]byte(auth))

	req.Header.Add("Content-Type", "application/timestamp-query")
	req.Header.Add("Accept", "application/timestamp-reply, application/timestamp-response")
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", encodedAuth))

	res, err := t.client.Do(req)
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
