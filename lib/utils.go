package lib

import (
	"math/rand"
	"net/http"
	"net/url"
	"os"
)

// Generates a random state sting
func GenerateState () string {
  var state string
  for i := 0; i < 32; i++ {
    state += string(rune(rand.Intn(26) + 97))
  }
  return state
}

/*
  Constructs a URL to the OAuth2 authorization endpoint
  
  @function Parameters:

    - @param {*http.Request} r - The request object
    - @param {*url.URL} redirect - The redirect URL

  URL Parameters that are set:
    - clientID: The client ID
    - redirectURI: The redirect URI
    - scope: The scope of the access request
    - state: The state to be returned to the client

  Returns:
    - The URL to redirect the user to

*/
func BuildQueryParams(r *http.Request, redirUri *url.URL) string {
	queryValues := r.URL.Query()
	  queryValues.Add("response_type", "code")
	  queryValues.Add("client_id", os.Getenv("SPOTIFY_CLIENT_ID"))
	  queryValues.Add("scopes", "user-read-private user-read-email")
	  queryValues.Add("redirect_uri", redirUri.String())
	  queryValues.Add("state", GenerateState())
	  query := queryValues.Encode()
	return query
}
