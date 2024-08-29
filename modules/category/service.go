package category

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-libr/modules/category/pb"
	"go-libr/utils/logger"
)

type service struct {
	categorySrvClient pb.ServiceClient
}

type Service interface {
	InsertCategorySrv(ctx *gin.Context, req InsertCategoryReq) (res InsertCategoryRes, err error)
	GetListCategorySrv(ctx *gin.Context, req GetListCategoryReq) (res []GetListCategoryRes, totalData int64, err error)
	UpdateCategorySrv(ctx *gin.Context, req UpdateCategoryReq) (err error)
}

func NewService(categorySrvClient pb.ServiceClient) Service {
	return &service{
		categorySrvClient: categorySrvClient,
	}
}

func (s *service) InsertCategorySrv(ctx *gin.Context, req InsertCategoryReq) (res InsertCategoryRes, err error) {
	// call gRPC method
	resp, err := s.categorySrvClient.InsertCategory(context.Background(), &pb.InsertCategoryRequest{
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

func (s *service) GetListCategorySrv(ctx *gin.Context, req GetListCategoryReq) (res []GetListCategoryRes, totalData int64, err error) {
	resp, err := s.categorySrvClient.ListCategory(context.Background(), &pb.ListCategoryRequest{
		Name:  req.Search.Name,
		Page:  req.Page,
		Limit: req.Limit,
	})
	if err != nil {
		logger.ErrorWithCtx(ctx, nil, err)
		err = errors.New("failed get list category")
		return
	}

	for _, category := range resp.Categories {
		res = append(res, GetListCategoryRes{
			ID:          category.Id,
			Name:        category.Name,
			Description: category.Description,
		})
	}

	totalData = resp.TotalData

	return
}

func (s *service) UpdateCategorySrv(ctx *gin.Context, req UpdateCategoryReq) (err error) {
	_, err = s.categorySrvClient.UpdateCategory(context.Background(), &pb.UpdateCategoryRequest{
		Name:        req.Name,
		Description: req.Description,
		ModifiedBy:  "system",
		Id:          req.ID,
	})
	if err != nil {
		logger.ErrorWithCtx(ctx, nil, err)
		err = errors.New("failed update category")
		return
	}

	return
}
