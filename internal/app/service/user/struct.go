package user

type User struct {
	ID           int64  `json:"id" deepcopier:"field:ID"`
	UserName     string `json:"user_name" deepcopier:"UserName"`
	Phone        string `json:"phone" deepcopier:"Phone"`
	RegisteredAt string `json:"registered_at" deepcopier:"RegisteredAt"`
	Role         string `json:"role" deepcopier:"Role"`
}
