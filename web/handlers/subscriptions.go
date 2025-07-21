package handlers

type SubscriptionResponse struct {
	ID          int    `json:"id" example:"123"`
	ServiceName string `json:"service_name" example:"Yandex Plus"`
	Price       int    `json:"price" example:"1899"`
	UserID      string `json:"user_id" example:"6ba7b811-9dad-11d1-80b4-00c04fd430c8"`
	StartDate   string `json:"start_date" example:"05-2025"`
} //@name SubscriptionResponse
