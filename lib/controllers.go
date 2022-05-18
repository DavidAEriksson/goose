package lib

import (
	"fmt"
	"net/http"
	"os"
)

/*
  Controller:
  - RedirectUser: redirects request to SPOTIFY_AUTH_URL
  
  @function Parameters:
  - w: http.ResponseWriter - response writer
  - r: *http.Request - request

  @function Returns:
  - nil
*/
func RedirectUser(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  redirUri, _ := r.URL.Parse("http://" + os.Getenv("GOOSE_SERVER_HOST") + ":" + os.Getenv("GOOSE_SERVER_PORT") + "/callback")
  query := BuildQueryParams(r, redirUri)
  fmt.Printf("Redirecting to %s\n", redirUri)
  http.Redirect(w, r, os.Getenv("SPOTIFY_AUTH_URL") + query, http.StatusTemporaryRedirect)
}
