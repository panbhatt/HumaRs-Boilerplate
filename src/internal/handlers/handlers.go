package handlers

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/danielgtaylor/huma/v2"

	"github.com/panbhatt/HumaRs-Boilerplate/src/internal/models/response"
)

func RegisterHandlers(api huma.API) {

	slog.Info("Registering Handlers ... ")
	registerStatusUrl(api)
}

func registerStatusUrl(api huma.API) {

	huma.Register(api, huma.Operation{
		OperationID:   "Get_System_Status",
		Method:        http.MethodGet,
		Path:          "/status",
		Summary:       "Get System Status",
		Description:   "/status will give you the System Status",
		Tags:          []string{"Status"},
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, input *struct{}) (*response.StatusResponse, error) {

		statusRes := &response.StatusResponse{}
		statusRes.Body.Message = "System is UP and functionaling till " + time.Now().String()
		return statusRes, nil

	})

}
