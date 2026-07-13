package handler

import (
	"github.com/gin-gonic/gin"
	"iotatfan.com/airline-voucher/internal/voucher/models"
	"iotatfan.com/airline-voucher/internal/voucher/service"
)

type JSONResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	IsError bool   `json:"isError"`
}

type VoucherHandler struct {
	voucherService service.VoucherService
}

func NewVoucherHandler(voucherService service.VoucherService) *VoucherHandler {
	return &VoucherHandler{
		voucherService: voucherService,
	}
}

func (h *VoucherHandler) CheckFlight(c *gin.Context) {
	req := models.CheckFlightRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		returnJSONResponse(c, err.Error(), nil, 400, true)
		return
	}

	result, err := h.voucherService.CheckFlight(req.FlightNumber, req.Date)
	if err != nil {
		returnJSONResponse(c, err.Error(), nil, 500, true)
		return
	}

	c.JSON(200, result)
}

func (h *VoucherHandler) GenerateRandomSeats(c *gin.Context) {
	req := models.GenerateRandomSeatsRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		returnJSONResponse(c, err.Error(), nil, 400, true)
	}

	result, err := h.voucherService.GenerateRandomSeats(req.Name, req.ID, req.FlightNumber, req.Date, req.Aircraft)
	if err != nil {
		returnJSONResponse(c, err.Error(), nil, 500, true)
	}

	c.JSON(200, result)
}

func returnJSONResponse(c *gin.Context, message string, data any, statusCode int, isError bool) {
	response := JSONResponse{
		Message: message,
		Data:    data,
		IsError: isError,
		Code:    statusCode,
	}
	c.JSON(statusCode, response)
}
