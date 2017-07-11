package rest

import (
	"github.com/valyala/fasthttp"
	"encoding/json"
	"fmt"
)

var (
	error	errorMessage
)

type errorMessage struct {
	Status		int		`json:"status"`
	Message		string	`json:"message"`
}

func ErrorHandler(ctx *fasthttp.RequestCtx, status int, err string) *fasthttp.RequestCtx{
	error.Status = status
	switch status {
		case fasthttp.StatusNoContent:
			ctx.SetStatusCode(fasthttp.StatusNoContent)
			break
		default:
			error.Message = "An error has been produced."
			errorJson, _ := json.Marshal(&error)
			fmt.Fprintf(ctx, string(errorJson))
			ctx.SetStatusCode(fasthttp.StatusOK)
			break
	}
	return  ctx
}
