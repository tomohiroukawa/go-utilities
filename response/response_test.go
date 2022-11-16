package response

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"

	validation "github.com/go-ozzo/ozzo-validation"
)

func TestOK(t *testing.T) {

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	OK(c, gin.H{
		"message": "ok",
	})

	if w.Code != http.StatusOK {
		t.Fatalf("status code should be 200. %d returned", w.Code)
	}

	if w.Body.String() != `{"message":"ok"}` {
		t.Fatalf("result did not match")
	}
}

func TestError(t *testing.T) {

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	message := "error occurred"

	Error(c, http.StatusInternalServerError, message)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("status code should be 500. %d returned", w.Code)
	}

	if w.Body.String() != fmt.Sprintf(`{"message":"%s"}`, message) {
		t.Fatalf("result did not match")
	}
}

func TestBadRequest(t *testing.T) {

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	message := "message string"

	BadRequest(c, message)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("status code should be 400. %d returned", w.Code)
	}

	if w.Body.String() != fmt.Sprintf(`{"message":"%s"}`, message) {
		t.Fatalf("result did not match")
	}
}

func TestNotFound(t *testing.T) {

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	message := "message string"

	NotFound(c, message)

	if w.Code != http.StatusNotFound {
		t.Fatalf("status code should be 404. %d returned", w.Code)
	}

	if w.Body.String() != fmt.Sprintf(`{"message":"%s"}`, message) {
		t.Fatalf("result did not match")
	}
}

func TestUnauthorized(t *testing.T) {

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	message := "message string"

	Unauthorized(c, message)

	if w.Code != http.StatusUnauthorized {
		t.Fatalf("status code should be 401. %d returned", w.Code)
	}

	if w.Body.String() != fmt.Sprintf(`{"message":"%s"}`, message) {
		t.Fatalf("result did not match")
	}
}

func TestForbidden(t *testing.T) {

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	message := "message string"

	Forbidden(c, message)

	if w.Code != http.StatusForbidden {
		t.Fatalf("status code should be 403. %d returned", w.Code)
	}

	if w.Body.String() != fmt.Sprintf(`{"message":"%s"}`, message) {
		t.Fatalf("result did not match")
	}
}

func TestInternalServerError(t *testing.T) {

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	message := "message string"

	InternalServerError(c, message)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("status code should be 500. %d returned", w.Code)
	}

	if w.Body.String() != fmt.Sprintf(`{"message":"%s"}`, message) {
		t.Fatalf("result did not match")
	}
}

func TestValidationError(t *testing.T) {

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	validationErrors := validation.Errors{
		"email":    fmt.Errorf("invalid format"),
		"nickname": fmt.Errorf("required"),
	}

	ValidationError(c, validationErrors)

	if w.Code != http.StatusNotAcceptable {
		t.Fatalf("status code should be 406. %d returned", w.Code)
	}

	if w.Body.String() != `{"email":"invalid format","nickname":"required"}` {
		t.Fatalf("result did not match")
	}
}
