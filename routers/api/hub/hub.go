package hub

import (
	"github.com/gin-gonic/gin"
	"hub/src/app/models"
	"hub/src/app/pkg/app"
	"hub/src/app/pkg/e"
	"hub/src/app/pkg/setting"
	"hub/src/app/pkg/util"
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

type GetTypeReqParam struct {
	DomainType string `form:"type" binding:"required"`

}

type TypeRet struct {
	Index int `json:"index"`
	Title string `json:"title"`
	Url string `json:"url"`
	ImageUrl string `json:"imageUrl"`
	TypeDomainID string `json:"typeDomainId"`
	TypeDomain string `json:"typeDomain"`
	TypeFilter string `json:"typeFilter"`
	CommontNum int `json:"commentNum"`
	Desc string `json:"desc"`
	Extra string `json:"extra"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}

func GetType(c *gin.Context){
	appG := app.Gin{C: c}
	var allTypes []TypeRet
	domainType := c.Query("type")
	pageOffset := util.GetPage(c)
	var reqParam GetTypeReqParam
	if c.ShouldBind(&reqParam) == nil {
		data ,err:= models.GetType(domainType,pageOffset,setting.AppSetting.PageSize)
		if err != nil {
			appG.Response(http.StatusBadRequest,e.ERROR,allTypes)
		}
		for _,v := range data{
			allTypes = append(allTypes, TypeRet{
				Index: v.Index,
				Title: v.Title,
				Url: v.Url,
				ImageUrl: v.ImageUrl,
				TypeDomainID: v.TypeDomainID,
				TypeDomain: v.TypeDomain,
				TypeFilter: v.TypeFilter,
				Desc: v.Desc,
				Extra: v.Extra,
				CreateTime: v.CreatedAt.String(),//Format("2006-01-02 15:04:05")
				UpdateTime: v.UpdatedAt.String(),
				})
		}
		appG.Response(http.StatusOK,e.SUCCESS,allTypes)
	}else{
		appG.Response(http.StatusBadRequest,e.INVALID_PARAMS,allTypes)
	}



}
