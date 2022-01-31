package user

import "time"

type User struct {
	ID           int64  `json:"id" deepcopier:"field:ID"`
	UserName     string `json:"user_name" deepcopier:"UserName"`
	Phone        string `json:"phone" deepcopier:"Phone"`
	RegisteredAt string `json:"registered_at" deepcopier:"RegisteredAt"`
	Role         string `json:"role" deepcopier:"Role"`
}

type UserLogin struct {
	Phone    string `json:"phone" deepcopier:"Phone"`
	Password string `json:"password" deepcopier:"Password"`
}

type UserLoginRes struct {
	AccessToken string    `json:"access_token"`
	ExpiredAt   time.Time `json:"expired_at"`
}

type ClaimRes struct {
	UserName     string `json:"user_name" deepcopier:"UserName"`
	Phone        string `json:"phone" deepcopier:"Phone"`
	RegisteredAt string `json:"registered_at" deepcopier:"RegisteredAt"`
	Role         string `json:"role" deepcopier:"Role"`
}
