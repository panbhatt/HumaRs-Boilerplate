package response

type StatusResponse struct {
	Body struct {
		Message string `json:"message" example:"System is UP" doc:"Status Message from the System`
	}
}
