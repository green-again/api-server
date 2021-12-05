package api

import (
	apipkg "api-server/internal/pkg/api"
)

const (
	InvalidRequest apipkg.ErrorCode = "article-01"
	NotFound       apipkg.ErrorCode = "article-02"
	Unknown        apipkg.ErrorCode = "article-03"
)
