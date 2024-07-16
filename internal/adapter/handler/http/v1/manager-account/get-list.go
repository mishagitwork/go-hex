package manager_account

import (
	"go-hex-arch/internal/adapter/handler/http/v1/contracts"
	errs "go-hex-arch/internal/adapter/handler/http/v1/errors"
	"go-hex-arch/internal/core/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (c *ManagerAccountHandler) GetList(ctx *gin.Context) {
	var managerQuery dto.QueryManagerAccount

	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "20"))
	search := ctx.DefaultQuery("search", "")

	res, err := c.svc.GetList(ctx, dto.Query{
		Page:   page,
		Limit:  limit,
		Search: search,
	}, managerQuery)

	if err != nil {
		errs.RespondError(ctx, err)
		return
	}

	resp := dto.List[contracts.ManagerAccountResponseContract]{}
	resp.Meta = res.Meta

	for _, value := range res.Data {
		resp.Data = append(resp.Data, contracts.ManagerAccountResponseContract{
			ID: value.ID,
			ManagerAccountContract: contracts.ManagerAccountContract{
				Image:         value.Image,
				Name:          value.Name,
				Email:         value.Email,
				MessengerLink: value.MessengerLink,
				Phone:         value.Phone,
			},
		})
	}

	ctx.JSON(http.StatusOK, contracts.JsonResponse[dto.List[contracts.ManagerAccountResponseContract]]{
		Result: resp,
	})
}
