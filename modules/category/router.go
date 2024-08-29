package category

import (
	"github.com/gin-gonic/gin"
	"go-libr/middlewares"
	"go-libr/modules/category/pb"
	"go-libr/utils/common"
	"google.golang.org/grpc"
	"log"
	"strconv"
)

func Initiator(router *gin.Engine) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}

	grpcClient := pb.NewServiceClient(conn)
	categorySrv := NewService(grpcClient)

	api := router.Group("/api/categories")
	api.Use(middlewares.Logging())
	//api.Use(middlewares.JwtMiddleware())
	{
		api.POST("", func(c *gin.Context) {
			CreateCategory(c, categorySrv)
		})
		api.GET("", func(c *gin.Context) {
			GetListCategory(c, categorySrv)
		})
		api.PUT("/:id", func(c *gin.Context) {
			UpdateCategory(c, categorySrv)
		})
	}

}

// CreateCategory godoc
// @Tags Categories
// @Summary CreateCategory
// @Description	This endpoint is used for create category
// @Accept json
// @Produce json
// @Param login body InsertCategoryReq true "Request"
// @Success 200 {object} common.APIResponse{data=InsertCategoryRes} "Success"
// @Failure 500	{object} common.APIResponse "Failed"
// @Router /api/categories [post]
// @Security BearerAuth
func CreateCategory(ctx *gin.Context, service Service) {
	var (
		req InsertCategoryReq
	)

	// validation request section
	err := ctx.ShouldBind(&req)
	if err != nil {
		return
	}

	_, err = service.InsertCategorySrv(ctx, req)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully create category")
}

// GetListCategory godoc
// @Tags Categories
// @Summary GetListCategory
// @Description	This endpoint is used for create category
// @Accept json
// @Produce json
// @Param page query int false "Page number for pagination" default(1)
// @Param limit query int false "Page number for pagination" default(100)
// @Param search query string false "Search by name for pagination" default(arif)
// @Success 200 {object} common.APIResponse{data=GetListCategoryRes} "Success"
// @Failure 500	{object} common.APIResponse "Failed"
// @Router /api/categories [get]
// @Security BearerAuth
func GetListCategory(ctx *gin.Context, service Service) {
	var (
		req GetListCategoryReq
	)

	req.Page, _ = strconv.ParseInt(ctx.Query("page"), 10, 64)
	req.Limit, _ = strconv.ParseInt(ctx.Query("limit"), 10, 64)
	req.Search.Name = ctx.Query("search")

	resp, total, err := service.GetListCategorySrv(ctx, req)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithListData(ctx, "successfully get all data category", total, resp)
}

// UpdateCategory godoc
// @Tags Categories
// @Summary UpdateCategory
// @Description	This endpoint is used for create category
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param update body UpdateCategoryReq true "Request"
// @Success 200 {object} common.APIResponse{data=UpdateCategoryRes} "Success"
// @Failure 500	{object} common.APIResponse "Failed"
// @Router /api/categories/{id} [put]
// @Security BearerAuth
func UpdateCategory(ctx *gin.Context, service Service) {
	var (
		req UpdateCategoryReq
	)

	// validation request section
	err := ctx.ShouldBind(&req)
	if err != nil {
		return
	}

	paramId := ctx.Param("id")
	req.ID, _ = strconv.ParseInt(paramId, 10, 64)

	err = req.Validate()
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	err = service.UpdateCategorySrv(ctx, req)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully update category")
}
