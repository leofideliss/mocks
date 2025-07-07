package handlers

import (
    "github.com/gin-gonic/gin"
    "net/http"
)
func UpdateDynamicChargeRendimento(c *gin.Context) {
    response :=map[string]interface{}{
        "data": map[string]interface{}{
            "value": map[string]interface{}{
                "qrCode": map[string]interface{}{
                    "formato": map[string]interface{}{
                        "id":        "BRCode",
                        "descricao": "BRCode",
                },
                    "valor": "00020101021226850014BR.GOV.BCB.PIX2563apisandbox.rendimento.com.br/q/b216d81a0aa5451d828cef0fd18688fa520400005303986540510.105802BR5914CLIENTE EVP PJ6009São Paulo62400536670d4691-ea32-4bde-ba2d-6b3a758bdb5d63049A0",
                },
            },
            "transactionId": "ead6d862-198c-447d-ae3e-44be284ebd0d",
            "erroMessage": map[string]interface{}{
                "statusCode": 0,
                "message":    nil,
                "errors":     []interface{}{},
            },
            "isSuccess": true,
            "isFailure": false,
        },
    }
    c.JSON(http.StatusOK, response)
}
