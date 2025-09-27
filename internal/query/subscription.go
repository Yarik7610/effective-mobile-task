package query

type ListSubscriptions struct {
	Page  uint   `form:"page" binding:"required,min=1"`
	Count uint   `form:"count" binding:"omitempty,max=50"`
	Sort  string `form:"sort" binding:"omitempty,oneof=service_name price start_date end_date"`
	Order string `form:"order" binding:"omitempty,oneof=asc desc"`
}

type TotalSubscriptionsPrice struct {
	StartDate   string  `form:"start_date" binding:"required"`
	EndDate     string  `form:"end_date" binding:"required"`
	UserID      *string `form:"user_id"`
	ServiceName *string `form:"service_name"`
}
