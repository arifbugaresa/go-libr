package category

type InsertCategoryReq struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type InsertCategoryRes struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
