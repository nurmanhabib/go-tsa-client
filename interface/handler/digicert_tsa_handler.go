package handler

import (
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/nurmanhabib/go-tsa-client/pkg/tsa"
)

func DigicertTSAHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}

	tsq, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	client := tsa.NewDigiCertClient()
	tsr, err := client.TSARequest(tsq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/timestamp-reply")
	w.Header().Set("Content-Length", strconv.Itoa(int(tsr.ContentLength)))
	w.Header().Set("Date", tsr.Date)
	w.WriteHeader(http.StatusOK)

	w.Write(tsr.Data)
}
