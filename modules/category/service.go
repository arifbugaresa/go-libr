package category

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-libr/modules/category/pb"
	"go-libr/utils/logger"
)

type service struct {
	grpcClient pb.ServiceClient
}

type Service interface {
	InsertCategorySrv(ctx *gin.Context, req InsertCategoryReq) (res InsertCategoryRes, err error)
}

func NewService(grpcClient pb.ServiceClient) Service {
	return &service{
		grpcClient: grpcClient,
	}
}

func (s *service) InsertCategorySrv(ctx *gin.Context, req InsertCategoryReq) (res InsertCategoryRes, err error) {
	// call gRPC method
	resp, err := s.grpcClient.InsertCategory(context.Background(), &pb.InsertCategoryRequest{
		Name:        req.Name,
		Description: req.Description,
		ModifiedBy:  "system",
	})
	if err != nil {
		logger.ErrorWithCtx(ctx, nil, err)
		return
	}

	fmt.Println(resp)

	return
}
