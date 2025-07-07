package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/leofideliss/mocks/handlers"
)

func SetupRoutes(r *gin.Engine) {
    //JD
    r.PUT("/jdpi/qrcode/api/v1/dinamico/:id", handlers.UpdateDynamicChargeJD) 
    r.POST("/jdpi/qrcode/api/v1/dinamico/gerar", handlers.CreateDynamicChargeJD)
    r.POST("/jdpi/qrcode/api/v1/dinamico/cobv/gerar", handlers.CreateDynamicDueChargeJD)
    r.PUT("/jdpi/qrcode/api/v1/dinamico/cobv/:idDcocumento", handlers.UpdateDynamicDueChargeJD)


    // Rendimento
    r.PUT("/spi/qrcode/updateDynamicDue", handlers.UpdateDynamicChargeRendimento)

}
