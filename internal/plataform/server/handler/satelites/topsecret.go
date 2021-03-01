package satellites

import (
    "github.com/gin-gonic/gin"
    quasar "github.com/rafaelraba/quasar_fire/internal"
    "net/http"
)

func TopsecretHandler(dataInterpreter quasar.DataInterpreter) gin.HandlerFunc {
    return func(ctx *gin.Context) {
        var req TopSecretRequest
        if err := ctx.BindJSON(&req); err != nil {
            ctx.JSON(http.StatusBadRequest, err.Error())
            return
        }
        response := new(TopSecretResponse)
        response.Position.X = -100.00
        response.Position.Y = 100.00
        response.Message = resolveMessage(req.Satellites, dataInterpreter)
        ctx.JSON(http.StatusOK, response)
    }
}

func resolveMessage(satellites []SatelliteRequest, interpreter quasar.DataInterpreter) string {
    var messages = make([][]string, len(satellites))
    for index, satellite := range satellites {
        messages [index] = append(satellite.Message)
    }
    return interpreter.GetMessage(messages...)
    
}