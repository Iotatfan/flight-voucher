package voucher

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"iotatfan.com/airline-voucher/internal/router"
	"iotatfan.com/airline-voucher/internal/voucher/handler"
	"iotatfan.com/airline-voucher/internal/voucher/repository"
	"iotatfan.com/airline-voucher/internal/voucher/service"
)

func Register(g *gin.Engine, db *gorm.DB) {
	repo := repository.NewVoucherRepository(db)
	service := service.NewVoucherService(repo)
	handler := handler.NewVoucherHandler(service)
	router.SetVoucherRoutes(g, *handler)
}
