package manager_account

import (
	"go-hex-arch/internal/adapter/handler/http/v1/contracts"
	errs "go-hex-arch/internal/adapter/handler/http/v1/errors"
	"go-hex-arch/internal/adapter/handler/http/v1/serializers"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c *ManagerAccountHandler) Get(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		errs.RespondError(ctx, err)
		return
	}

	res, err := c.svc.Get(ctx, id)
	if err != nil {
		errs.RespondError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, contracts.JsonResponse[contracts.ManagerAccountResponseContract]{
		Result: *serializers.ManagerAccountResponse(res),
	})
}
