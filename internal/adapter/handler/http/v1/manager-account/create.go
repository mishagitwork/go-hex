package manager_account

import (
	"go-hex-arch/internal/adapter/handler/http/v1/contracts"
	errs "go-hex-arch/internal/adapter/handler/http/v1/errors"
	"go-hex-arch/internal/adapter/handler/http/v1/serializers"
	"go-hex-arch/internal/core/domain"

	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *ManagerAccountHandler) Create(ctx *gin.Context) {

	var data contracts.ManagerAccountContract

	if err := ctx.ShouldBindJSON(&data); err != nil {
		errs.RespondError(ctx, err)
		return
	}

	domainData := domain.ManagerAccount{Image: data.Image, Name: data.Name, Email: data.Email, Phone: data.Phone}

	res, err := c.svc.Create(ctx, domainData)
	if err != nil {
		errs.RespondError(ctx, err)
		return
	}

	response := *serializers.ManagerAccountResponse(res)

	ctx.JSON(http.StatusCreated, contracts.JsonResponse[contracts.ManagerAccountResponseContract]{
		Result: response,
	})
}
