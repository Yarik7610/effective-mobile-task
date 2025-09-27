package query

type ListSubscriptions struct {
	Page  uint   `form:"page" binding:"required,min=1"`
	Count uint   `form:"count" binding:"required,min=1,max=50"`
	Sort  string `form:"sort" binding:"oneof service_name price start_date end_date"`
	Order string `form:"order" binding:"oneof=asc desc"`
}
