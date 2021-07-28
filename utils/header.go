package utils

import (
	"net/http"
	"strings"
)

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	if len(bearToken) > 0 {
		//normally Authorization the_token_xxx
		strArr := strings.Split(bearToken, " ")
		if len(strArr) == 2 {
			return strArr[1]
		}
		return ""
	}
	return ""
}
