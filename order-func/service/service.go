package service

import (
	"encoding/json"
	"io"

	"github.com/google/uuid"
	"gofr.dev/pkg/gofr"

	"order-func/model"
)

type service struct{}

func New() Order {
	return service{}
}

type list struct {
	Data []model.Order `json:"data"`
}

type response struct {
	Data *model.Order `json:"data"`
}

func (s service) Create(ctx *gofr.Context, order *model.Order) (*model.Order, error) {
	body, err := json.Marshal(order)
	if err != nil {
		return nil, err
	}

	resp, err := ctx.GetHTTPService("order-data").PostWithHeaders(ctx, "orders", nil, body,
		map[string]string{"Content-Type": "application/json"})
	if err != nil {
		return nil, err
	}

	var o response

	body, _ = io.ReadAll(resp.Body)
	err = json.Unmarshal(body, &o)
	if err != nil {
		return nil, err
	}

	return o.Data, nil
}

func (s service) GetAll(ctx *gofr.Context) ([]model.Order, error) {
	resp, err := ctx.GetHTTPService("order-data").Get(ctx, "orders", nil)
	if err != nil {
		return nil, err
	}

	var l list

	body, _ := io.ReadAll(resp.Body)

	err = json.Unmarshal(body, &l)
	if err != nil {
		return nil, err
	}

	return l.Data, err
}

func (s service) GetByID(ctx *gofr.Context, id uuid.UUID) (*model.Order, error) {
	resp, err := ctx.GetHTTPService("order-data").Get(ctx, "orders"+"/"+id.String(), nil)
	if err != nil {
		return nil, err
	}

	var order response

	body, _ := io.ReadAll(resp.Body)

	err = json.Unmarshal(body, &order)
	if err != nil {
		return nil, err
	}

	return order.Data, nil
}

func (s service) Update(ctx *gofr.Context, order *model.Order) (*model.Order, error) {
	body, err := json.Marshal(order)
	if err != nil {
		return nil, err
	}

	resp, err := ctx.GetHTTPService("order-data").PutWithHeaders(ctx, "orders"+"/"+order.ID.String(), nil, body,
		map[string]string{"Content-Type": "application/json"})
	if err != nil {
		return nil, err
	}

	var o response

	body, _ = io.ReadAll(resp.Body)

	err = json.Unmarshal(body, &o)
	if err != nil {
		return nil, err
	}

	return o.Data, nil
}

func (s service) Delete(ctx *gofr.Context, id uuid.UUID) error {
	_, err := ctx.GetHTTPService("order-data").Delete(ctx, "orders"+"/"+id.String(), nil)

	return err
}
