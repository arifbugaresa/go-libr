package category

import (
	"github.com/gin-gonic/gin"
	"go-libr/middlewares"
	"go-libr/modules/category/pb"
	"go-libr/utils/common"
	"google.golang.org/grpc"
	"log"
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
	{
		api.POST("", func(c *gin.Context) {
			CreateCategory(c, categorySrv)
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

	common.GenerateSuccessResponse(ctx, "successfully login")
}
