package web

import (
	"fmt"
	"net/http"
)

func (apiCfg *apiConfig) setHeaders(r *http.Request) {
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %v", apiCfg.APIKey))
}
