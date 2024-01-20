package logic

import (
	"free-community/db/model"
	"gorm.io/gorm"
)

type Handler struct {
	tx *gorm.DB
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) DB() *gorm.DB {
	if h.tx == nil {
		return model.DB
	}
	return h.tx
}
