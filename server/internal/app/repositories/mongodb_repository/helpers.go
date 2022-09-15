package mongodb_repository

import (
	"server/internal/app"
	"strconv"
)

func ParseCtxTimeoutEnv() (int, error) {
	ctxTimeout, err := strconv.Atoi(app.CtxTimeout)
	return ctxTimeout, err
}
