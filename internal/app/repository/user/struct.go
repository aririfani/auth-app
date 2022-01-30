package user

type User struct {
	ID           int64  `json:"id" deepcopier:"ID"`
	UserName     string `json:"user_name" deepcopier:"UserName"`
	Phone        string `json:"phone" deepcopier:"Phone"`
	RegisteredAt string `json:"registered_at" deepcopier:"RegisteredAt"`
	Password     string `json:"password" deepcopier:"Password"`
	Role         string `json:"role" deepcopier:"Role"`
}

type CreateRes struct {
	ID       int64  `json:"id" deepcopier:"ID"`
	UserName string `json:"user_name" deepcopier:"UserName"`
	Phone    string `json:"phone" deepcopier:"Phone"`
	Role     string `json:"role" deepcopier:"Role"`
}

type Res struct {
	ID           int64  `json:"id" deepcopier:"ID"`
	UserName     string `json:"user_name" deepcopier:"UserName"`
	Phone        string `json:"phone" deepcopier:"Phone"`
	RegisteredAt string `json:"registered_at" deepcopier:"RegisteredAt"`
	Role         string `json:"role" deepcopier:"Role"`
	Password     string `json:"password"`
}
