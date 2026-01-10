package handlers

import (
	"net/http"
	"strconv"
	"strings"
)

func parseIntParam(r *http.Request, key string, defaultVal int) int {
	val := r.URL.Query().Get(key)
	if val == "" {
		return defaultVal
	}
	i, err := strconv.Atoi(val)
	if err != nil {
		return defaultVal
	}
	return i
}

func parseArrayParam(r *http.Request, key string) []string {
	val := r.URL.Query().Get(key)
	if val == "" {
		return []string{}
	}
	return strings.Split(val, ",")
}
