package errs

import (
	"errors"
	"go-hex-arch/internal/adapter/handler/http/v1/contracts"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func RespondError(ctx *gin.Context, err error) {
	_ = ctx.Error(err)
	var status int
	var response string

	switch true {
	case errors.As(err, &validator.ValidationErrors{}):
		response = err.Error()
		status = http.StatusUnprocessableEntity

	default:
		response = err.Error()
		status = http.StatusBadRequest
	}

	ctx.JSON(status, contracts.JsonResponse[struct{}]{
		Error: response,
	})
}
