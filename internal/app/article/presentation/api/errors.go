package api

import (
	"api-server/internal/pkg/http"
)

const (
	InvalidRequest http.ErrorCode = "article-01"
	NotFound       http.ErrorCode = "article-02"
	Unknown        http.ErrorCode = "article-03"
)
