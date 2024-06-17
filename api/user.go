package api

import (
	"laptopsShop/engine"
)

func (h *Handler) GetUser(ctx *engine.Context) {
	ctx.Response.Write([]byte("Я юзер номер 5"))
}

func (h *Handler) GetAllUsers(ctx *engine.Context) {
	ctx.Response.Write([]byte("Здесь все юзеры"))
}

func (h *Handler) PostIndex(ctx *engine.Context) {
	ctx.Response.Write([]byte("Authorized"))
}
