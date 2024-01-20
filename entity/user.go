package entity

type UserRegisterReq struct {
	Name string `json:"name" binding:"required"`
}
