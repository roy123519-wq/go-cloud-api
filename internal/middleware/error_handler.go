package middleware

import (
	"net/http"

	"go-cloud-api/internal/response"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// 沒有任何 error，直接放過
		if len(c.Errors) == 0 {
			return
		}

		last := c.Errors.Last()
		if last == nil || last.Err == nil {
			return
		}

		err := last.Err

		// AppError：照內容回
		if appErr, ok := err.(*response.AppError); ok {
			c.JSON(appErr.Status, response.Fail(appErr.Code, appErr.Message))
			return
		}

		// 其他未知錯誤：統一 500
		c.JSON(http.StatusInternalServerError, response.Fail("INTERNAL_ERROR", "internal error"))
	}
}
