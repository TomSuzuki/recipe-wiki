package service

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

// QueryInt ...get query
func QueryInt(c *gin.Context, key string) (int, error) {
	if s, ok := c.GetQuery(key); ok {
		n, err := strconv.Atoi(s)
		if err != nil {
			return 0, err
		}
		return n, nil
	}
	return 0, errors.New("query not found")
}

// QueryString ...get query
func QueryString(c *gin.Context, key string) (string, error) {
	if s, ok := c.GetQuery(key); ok {
		return s, nil
	}
	return "", errors.New("query not found")
}
