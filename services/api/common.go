package api

import (
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/labstack/echo/v4"
	v "gopkg.in/go-playground/validator.v9"
)

type Param interface {
	SetAddress(address string)
}

type Response struct {
	Desc string      `json:"desc"`
	Data interface{} `json:"data,omitempty"`
}

func commonHandler(params interface{}, fn func(interface{}) (interface{}, error)) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var req Param
		if params != nil {
			req = reflect.New(reflect.TypeOf(params).Elem()).Interface().(Param)
			address, ok := ctx.Get(addressContextKey).(string)
			if ok {
				req.SetAddress(address)
			}

			if err := ctx.Bind(req); err != nil {
				return err
			}

			typ := reflect.TypeOf(req).Elem()
			val := reflect.ValueOf(req).Elem()

			for i := 0; i < typ.NumField(); i++ {
				structField := val.Field(i)
				if !structField.CanSet() {
					continue
				}

				typeField := typ.Field(i)
				inputFieldName := typeField.Tag.Get("param")
				if inputFieldName != "" {
					structField.SetString(ctx.Param(inputFieldName))
				}
			}

			if err := v.New().Struct(req); err != nil {
				return err
			}
		}

		resp, err := fn(req)
		if err != nil {
			return err
		}

		jsonBytes, err := json.Marshal(Response{
			Desc: "success",
			Data: resp,
		})
		if err != nil {
			return err
		}

		return ctx.String(http.StatusOK, string(jsonBytes))
	}
}
