package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/niubir/area-service/http/docs"
	_ "github.com/niubir/area-service/models"
	"github.com/niubir/area-service/services"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initRoutes(engine *gin.Engine) {
	api := engine.Group("")

	docs.SwaggerInfo.Title = "area service"
	docs.SwaggerInfo.Version = "v0.0.1"

	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	api.POST("", freshAreas)
	api.GET("", getArea)
	api.GET("/provinces", getProvinces)
	api.GET("/cities", getCities)
	api.GET("/districts", getDistricts)
	api.GET("/streets", getStreets)
}

// @Summary      刷新区域
// @Tags         相关接口
// @success      200 "结果"
// @Router       / [post]
func freshAreas(c *gin.Context) {
	err := services.FreshAreas()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}

// @Summary      获取区域
// @Tags         相关接口
// @Param        code  query  string  true  "区域编码"
// @success      200 {object} models.Area "结果"
// @Router       / [get]
func getArea(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, "parameter code must")
		return
	}
	areas, err := services.GetArea(code)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, areas)
}

// @Summary      获取省份
// @Tags         相关接口
// @success      200 {object} models.Areas "结果"
// @Router       /provinces [get]
func getProvinces(c *gin.Context) {
	provinces, err := services.GetProvinces()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, provinces)
}

// @Summary      获取城市
// @Tags         相关接口
// @Param        provinceCode  query  string  true  "省份区域编码"
// @success      200 {object} models.Areas "结果"
// @Router       /cities [get]
func getCities(c *gin.Context) {
	provinceCode := c.Query("provinceCode")
	if provinceCode == "" {
		c.JSON(http.StatusBadRequest, "parameter provinceCode must")
		return
	}
	cities, err := services.GetCities(provinceCode)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, cities)
}

// @Summary      获取区县
// @Tags         相关接口
// @Param        cityCode  query  string  true  "城市区域编码"
// @success      200 {object} models.Areas "结果"
// @Router       /districts [get]
func getDistricts(c *gin.Context) {
	cityCode := c.Query("cityCode")
	if cityCode == "" {
		c.JSON(http.StatusBadRequest, "parameter cityCode must")
		return
	}
	districts, err := services.GetDistricts(cityCode)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, districts)
}

// @Summary      获取街道
// @Tags         相关接口
// @Param        districtCode  query  string  true  "区县区域编码"
// @success      200 {object} models.Areas "结果"
// @Router       /streets [get]
func getStreets(c *gin.Context) {
	districtCode := c.Query("districtCode")
	if districtCode == "" {
		c.JSON(http.StatusBadRequest, "parameter districtCode must")
		return
	}
	streets, err := services.GetStreet(districtCode)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, streets)
}
