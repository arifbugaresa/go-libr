package category

import "errors"

type InsertCategoryReq struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (req *InsertCategoryReq) Validate() error {
	if req.Name == "" {
		return errors.New("name is required")
	}

	if req.Description == "" {
		return errors.New("description is required")
	}

	return nil
}

type InsertCategoryRes struct{}

type GetListCategoryReq struct {
	Page   int64    `json:"page"`
	Limit  int64    `json:"limit"`
	Search SearchBy `json:"search"`
}

type SearchBy struct {
	Name string `json:"name"`
}

type GetListCategoryRes struct {
	ID          int64  `json:"ID"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateCategoryReq struct {
	ID          int64  `json:"ID"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (req *UpdateCategoryReq) Validate() error {
	if req.ID == 0 {
		return errors.New("ID is required")
	}

	if req.Name == "" {
		return errors.New("name is required")
	}

	if req.Description == "" {
		return errors.New("description is required")
	}

	return nil
}

type UpdateCategoryRes struct{}
