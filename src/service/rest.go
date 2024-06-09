package service

import (
	"encoding/json"
	"net/http"

	"github.com/derarken/vlr-api/src/api"
)

func StartRest() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /match/{id}", func(w http.ResponseWriter, r *http.Request) {
		match, err := api.GetMatch(r.PathValue("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(match)
	})

	// mux.HandleFunc("GET /matchIds?{status}", func(w http.ResponseWriter, r *http.Request) {
	// 	status := proto.Status_value[r.PathValue("status")]
	// 	matchIds, err := api.GetMatchIds(proto.Status(status), time.Now(), time.Now().Add(time.Hour*24))
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.Header().Set("Access-Control-Allow-Origin", "*")
	// 	json.NewEncoder(w).Encode(matchIds)
	// })

	err := http.ListenAndServe(":8081", mux)
	if err != nil {
		panic(err)
	}
}
