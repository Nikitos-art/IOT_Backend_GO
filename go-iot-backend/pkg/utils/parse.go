package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) error {
	if r.Body == nil {
		return fmt.Errorf("empty request body")
	}

	return json.NewDecoder(r.Body).Decode(x)
}