package dto

type CreateSubscription struct {
	ServiceName string  `json:"service_name" binding:"required"`
	Price       int     `json:"price" binding:"required"`
	UserID      string  `json:"user_id" binding:"required,uuid"`
	StartDate   string  `json:"start_date" binding:"required"` //MM-YYYY
	EndDate     *string `json:"end_date,omitempty"`            //MM-YYYY
}

type UpdateSubscription struct {
	ServiceName *string `json:"service_name,omitempty"`
	Price       *int    `json:"price,omitempty"`
	StartDate   *string `json:"start_date,omitempty"` //MM-YYYY
	EndDate     *string `json:"end_date,omitempty"`   //MM-YYYY
}
