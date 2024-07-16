package http

import (
	manager_account "go-hex-arch/internal/adapter/handler/http/v1/manager-account"
	"go-hex-arch/internal/core/port"
)

type ManagerAccountHandler struct {
	*manager_account.ManagerAccountHandler
}

// NewCategoryHandler creates a new ManagerAccountHandler instance
func NewManagerAccountHandler(svc port.ManagerAccountService) *ManagerAccountHandler {
	return &ManagerAccountHandler{manager_account.New(svc)}
}

// func (c *ManagerAccountHandler) getManagerAccount(ctx *gin.Context) {
// 	id, err := uuid.Parse(ctx.Param("id"))
// 	if err != nil {
// 		respondError(ctx, err)
// 		return
// 	}

// 	res, err := c.svc.Get(ctx, id)
// 	if err != nil {
// 		respondError(ctx, err)
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, JsonResponse[contracts.ManagerAccountResponseContract]{
// 		Result: *serializers.ManagerAccountResponse(res),
// 	})
// }

// func (c *ManagerAccountHandler) getAdminManagerAccountList(ctx *gin.Context) {
// 	var managerQuery dto.QueryManagerAccount

// 	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
// 	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "20"))
// 	search := ctx.DefaultQuery("search", "")

// 	res, err := c.svc.GetList(ctx, dto.Query{
// 		Page:   page,
// 		Limit:  limit,
// 		Search: search,
// 	}, managerQuery)

// 	if err != nil {
// 		respondError(ctx, err)
// 		return
// 	}

// 	resp := dto.List[contracts.ManagerAccountResponseContract]{}
// 	resp.Meta = res.Meta

// 	for _, value := range res.Data {
// 		resp.Data = append(resp.Data, contracts.ManagerAccountResponseContract{
// 			ID: value.ID,
// 			ManagerAccountContract: contracts.ManagerAccountContract{
// 				Image:         value.Image,
// 				Name:          value.Name,
// 				Email:         value.Email,
// 				MessengerLink: value.MessengerLink,
// 				Phone:         value.Phone,
// 			},
// 		})
// 	}

// 	ctx.JSON(http.StatusOK, JsonResponse[dto.List[contracts.ManagerAccountResponseContract]]{
// 		Result: resp,
// 	})
// }

// func (c *ManagerAccountHandler) createManagerAccount(ctx *gin.Context) {

// 	var data contracts.ManagerAccountContract

// 	if err := ctx.ShouldBindJSON(&data); err != nil {
// 		respondError(ctx, err)
// 		return
// 	}

// 	domainData := domain.ManagerAccount{Image: data.Image, Name: data.Name, Email: data.Email, Phone: data.Phone}

// 	res, err := c.svc.Create(ctx, domainData)
// 	if err != nil {
// 		respondError(ctx, err)
// 		return
// 	}

// 	response := contracts.ManagerAccountResponseContract{
// 		ID: res.ID,
// 		ManagerAccountContract: contracts.ManagerAccountContract{
// 			Name:          res.Name,
// 			Image:         res.Image,
// 			Email:         res.Email,
// 			MessengerLink: res.MessengerLink,
// 			Phone:         res.Phone,
// 		}}

// 	ctx.JSON(http.StatusCreated, JsonResponse[contracts.ManagerAccountResponseContract]{
// 		Result: response,
// 	})
// }
