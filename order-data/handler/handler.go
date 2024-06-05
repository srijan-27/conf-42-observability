package handler

import (
	"strings"

	"github.com/google/uuid"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/http"

	"order-data/model"
	"order-data/store"
)

type handler struct {
	store store.Order
}

func New(s store.Order) handler {
	return handler{store: s}
}

func (h handler) Create(ctx *gofr.Context) (interface{}, error) {
	var order model.Order

	if err := ctx.Bind(&order); err != nil {
		return nil, http.ErrorInvalidParam{Params: []string{"body"}}
	}

	res, err := h.store.Create(ctx, &order)
	if err != nil {
		if _, ok := err.(http.ErrorEntityAlreadyExist); ok {
			return res, err
		}

		return nil, err
	}

	return res, nil
}

func (h handler) GetAll(ctx *gofr.Context) (interface{}, error) {
	res, err := h.store.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	if res == nil {
		return []model.Order{}, nil
	}

	return res, nil
}

func (h handler) GetByID(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")
	if strings.TrimSpace(id) == "" {
		return nil, http.ErrorMissingParam{Params: []string{"id"}}
	}

	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, http.ErrorInvalidParam{Params: []string{"id"}}
	}

	res, err := h.store.GetByID(ctx, uid)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (h handler) Update(ctx *gofr.Context) (interface{}, error) {
	var orders model.Order

	id := ctx.PathParam("id")
	if strings.TrimSpace(id) == "" {
		return nil, http.ErrorMissingParam{Params: []string{"id"}}
	}

	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, http.ErrorInvalidParam{Params: []string{"id"}}
	}

	orders.ID = uid

	if err = ctx.Bind(&orders); err != nil {
		return nil, http.ErrorInvalidParam{Params: []string{"body"}}
	}

	res, err := h.store.Update(ctx, &orders)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (h handler) Delete(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")
	if strings.TrimSpace(id) == "" {
		return nil, http.ErrorMissingParam{Params: []string{"id"}}
	}

	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, http.ErrorInvalidParam{Params: []string{"id"}}
	}

	err = h.store.Delete(ctx, uid)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
