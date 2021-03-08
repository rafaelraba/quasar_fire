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
        response.Position = resolveLocation(req.Satellites, dataInterpreter)
        response.Message = resolveMessage(req.Satellites, dataInterpreter)
        validateMessageResponse(response, ctx)
    }
}

func resolveMessage(satellites []SatelliteRequest, interpreter quasar.DataInterpreter) string {
    var messages = make([][]string, len(satellites))
    for index, satellite := range satellites {
        messages [index] = append(satellite.Message)
    }
    return interpreter.GetMessage(messages...)
}

func validateMessageResponse(response *TopSecretResponse, ctx *gin.Context) {
    if response.Message == "indeterminate" {
        info := buildMessageInfo("ERROR","message cannot be determined")
        ctx.JSON(http.StatusBadRequest,info)
        return
    }
    ctx.JSON(http.StatusOK, response)
}

func resolveLocation(Satellites []SatelliteRequest, interpreter quasar.DataInterpreter) Position {
    x, y := interpreter.GetLocation(
        float32(Satellites[0].Distance),
        float32(Satellites[1].Distance),
        float32(Satellites[2].Distance),
    )
    return Position{
        X: x,
        Y: y,
    }
}

func buildMessageInfo(status, description string) *ResponseInfo {
    info := new (ResponseInfo)
    info.Description = description
    info.Status = status
    return info
}