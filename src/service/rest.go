package service

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/derarken/vlr-api/proto"
	"github.com/derarken/vlr-api/src/api"
)

func StartRest() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /match", func(w http.ResponseWriter, r *http.Request) {
		const PARAM_ID = "id"
		id := r.URL.Query().Get(PARAM_ID)
		if id == "" {
			http.Error(w, "missing id parameter", http.StatusBadRequest)
			return
		}

		match, err := api.GetMatch(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(match)
	})

	// froma and to are in RFC3339 format
	mux.HandleFunc("GET /matchIds", func(w http.ResponseWriter, r *http.Request) {
		const PARAM_STATUS = "status"
		const PARAM_FROM = "from"
		const PARAM_TO = "to"

		statusParam := r.URL.Query().Get(PARAM_STATUS)
		fromParam := r.URL.Query().Get(PARAM_FROM)
		toParam := r.URL.Query().Get(PARAM_TO)

		status := proto.Status(proto.Status_value[statusParam])

		var from time.Time
		var to time.Time
		var err error
		if fromParam != "" {
			from, err = time.Parse(time.RFC3339, fromParam)
		}
		if toParam != "" {
			to, err = time.Parse(time.RFC3339, toParam)
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		matchIds, err := api.GetMatchIds(status, from, to)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(matchIds)
	})

	err := http.ListenAndServe(":8081", mux)
	if err != nil {
		panic(err)
	}
}
