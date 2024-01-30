package idp

import "net/http"

func HelloFromIDP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello from IDP!"))
}