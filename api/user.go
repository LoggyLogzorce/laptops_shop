package api

import (
	"vapingShop/engine"
)

func (h *Handler) GetUser(ctx *engine.Context) {
	ctx.Response.Write([]byte("Я юзер номер 5"))
}

func (h *Handler) GetAllUsers(ctx *engine.Context) {
	ctx.Response.Write([]byte("Здесь все юзеры"))
}
