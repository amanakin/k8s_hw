package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type TimeResponse struct {
	Time string `json:"time"`
}

func (h *Handler) HandleTime(w http.ResponseWriter, r *http.Request) {
	h.service.IncRequestsCount()

	w.Header().Set("Content-Type", "application/json")
	t, err := h.service.GetTime(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var resp = TimeResponse{
		Time: t.Format(time.RFC3339),
	}

	var b bytes.Buffer
	err = json.NewEncoder(&b).Encode(resp)
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
