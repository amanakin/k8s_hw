package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Statistics struct {
	Requests int `json:"requests"`
}

func (h *Handler) HandleStatistics(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var resp = Statistics{
		Requests: h.service.GetRequestsCount(),
	}
	var b bytes.Buffer

	err := json.NewEncoder(&b).Encode(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(b.Bytes())
	if err != nil {
		return
	}
}
