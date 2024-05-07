package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type Time struct {
	Seconds int64 `json:"seconds"`
}

func (h *Handler) HandleTime(w http.ResponseWriter, _ *http.Request) {
	h.service.IncRequestsCount()

	w.Header().Set("Content-Type", "application/json")
	var resp = Time{
		Seconds: time.Now().Unix(),
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
