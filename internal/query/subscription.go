package query

type ListSubscriptions struct {
	Page  uint   `form:"page" binding:"required,min=1"`
	Count uint   `form:"count" binding:"omitempty,max=50"`
	Sort  string `form:"sort" binding:"omitempty,oneof=service_name price start_date end_date"`
	Order string `form:"order" binding:"omitempty,oneof=asc desc"`
}
