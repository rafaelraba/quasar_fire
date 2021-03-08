package satellites

import (
	"fmt"
	"github.com/gin-gonic/gin"
	quasar "github.com/rafaelraba/quasar_fire/internal"
	"net/http"
)

var satellitesGroup = make([]SatelliteRequest, 0)

func TopSecretSaveHandler(dataInterpreter quasar.DataInterpreter) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req SatelliteRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		req.Name = ctx.Param("satellite_name")
		if !satelliteExist(req.Name) {
			fmt.Println("No Existe")
			fmt.Println(addMessage(req))
		}
		ctx.Status(http.StatusCreated)
	}

}

func TopSecretGetHandler(dataInterpreter quasar.DataInterpreter) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if len(satellitesGroup) < 3 {
			info := buildMessageInfo("ERROR", "Insufficient information")
			ctx.JSON(http.StatusBadRequest, info)
			return
		}
		response := new(TopSecretResponse)
		response.Message = resolveMessage(satellitesGroup, dataInterpreter)
		response.Position = resolveLocation(satellitesGroup, dataInterpreter)
		validateMessageResponse(response, ctx)

	}
}

func satelliteExist(name string) bool {
	for _, satellite := range satellitesGroup {
		if satellite.Name == name {
			return true
		}
	}
	return false
}

func addMessage(satellite SatelliteRequest) []SatelliteRequest {
	satellitesGroup = append(satellitesGroup, satellite)
	return satellitesGroup
}
