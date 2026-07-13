package router

import (
	"github.com/gin-gonic/gin"
	"iotatfan.com/airline-voucher/internal/voucher/handler"
)

func SetVoucherRoutes(g *gin.Engine, h handler.VoucherHandler) {
	g.POST("/api/check", h.CheckFlight)
	g.POST("/api/generate", h.GenerateRandomSeats)
}
