package middleware

import "net/http"

func GetUserID(r *http.Request) uint {
	val := r.Context().Value(UserKey)
	if val == nil {
		return 0
	}
	return val.(uint)
}