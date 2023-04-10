package response

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorResponse はエラー時のレスポンス型
type ErrorResponse struct {
	Message string `json:"message"`
}

// Error はエラー発生時の汎用レスポンス
func Error(c *gin.Context, status int, message string) {

	var res ErrorResponse
	res.Message = message

	c.JSON(status, res)
}

// BadRequest ステータスのレスポンス
func BadRequest(c *gin.Context, message string) {
	Error(c, http.StatusBadRequest, message)
}

// NotFound ステータスのレスポンス
func NotFound(c *gin.Context, message string) {
	Error(c, http.StatusNotFound, message)
}

// Forbidden ステータスのレスポンス
func Forbidden(c *gin.Context, message string) {
	Error(c, http.StatusForbidden, message)
}

// Unauthorized ステータスのレスポンス
func Unauthorized(c *gin.Context, message string) {
	Error(c, http.StatusUnauthorized, message)
}

// InternalServerError ステータスのレスポンス
func InternalServerError(c *gin.Context, message string) {
	Error(c, http.StatusInternalServerError, message)
}

// ValidationError はバリデーションエラー時のレスポンス
func ValidationError(c *gin.Context, validationError error) {

	var errs map[string]string

	b, err := json.Marshal(validationError)

	if err != nil {
		log.Println(err.Error())
		return
	}

	if err := json.Unmarshal(b, &errs); err != nil {
		log.Println(err.Error())
		return
	}

	c.JSON(http.StatusNotAcceptable, errs)
}

// OK は正常時のレスポンス
func OK(c *gin.Context, res interface{}) {
	c.JSON(http.StatusOK, res)
}
