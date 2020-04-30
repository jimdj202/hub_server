package hub

import (
	"github.com/gin-gonic/gin"
	"hub/src/app/models"
	"hub/src/app/pkg/app"
	"hub/src/app/pkg/e"
	"net/http"
)

type AllTypesRet struct {
	TypeDomainID string `json:"typeDomainId"`
	TypeDomain string `json:"typeDomain"`
}

func GetAllTypes(c *gin.Context){
	appG := app.Gin{C: c}
	data ,err:= models.GetAllTypes()
	var allTypes []AllTypesRet
	for _,v := range data{
		allTypes = append(allTypes, AllTypesRet{TypeDomainID: v.TypeDomainID,TypeDomain: v.TypeDomain})
	}
	if err != nil {
		appG.Response(http.StatusBadRequest,e.ERROR,allTypes)
	}else {
		appG.Response(http.StatusOK,e.SUCCESS,allTypes)
	}
}

