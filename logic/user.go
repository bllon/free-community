package logic

import (
	"free-community/db/model"
	"free-community/entity"
	"log"
)

func (h *Handler) UserRegister(req entity.UserRegisterReq) (string, error) {
	if err := h.DB().Create(&model.User{
		Name: req.Name,
	}).Error; err != nil {
		log.Panicln(err)
	}
	return "ok", nil
}

func (h *Handler) UserInfo() (string, error) {
	return "hello", nil
}
