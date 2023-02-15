package controller

import (
	"crud_go/model"
	"crud_go/service"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

const rest string = "restaurantes"
const stor string = "tiendas"

func GetLocations(ctx *gin.Context) {
	var locations []model.Location
	model.Database.Find(&locations)
	ctx.JSON(http.StatusOK, gin.H{"data": locations})
}

func CreateLocation(ctx *gin.Context) {
	var pld model.LocationPayloadDto
	if err := ctx.ShouldBindJSON(&pld); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	loc := model.Location{Ent: pld.Ent, Nombre: pld.Nombre, Lat: pld.Lat, Lon: pld.Lon}
	model.Database.Create(&loc)
	ctx.JSON(http.StatusCreated, gin.H{"data": loc})
}

func GetLocation(ctx *gin.Context) {
	var loc model.Location
	if err := model.Database.Where("id = ?", ctx.Param("id")).First(&loc).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Location not found."})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": loc})
}

func UpdateLocation(ctx *gin.Context) {
	var loc model.Location
	if err := model.Database.Where("id = ?", ctx.Param("id")).First(&loc).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Location not found."})
		return
	}
	var pld model.LocationUpdateDto

	if err := ctx.ShouldBindJSON(&pld); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedLoc := model.Location{Nombre: pld.Nombre, Ent: pld.Ent, Lat: pld.Lat, Lon: pld.Lon}

	model.Database.Model(&loc).Updates(&updatedLoc)
	ctx.JSON(http.StatusOK, gin.H{"data": loc})

}

func DeleteLocation(ctx *gin.Context) {
	var loc model.Location
	if err := model.Database.Where("id = ?", ctx.Param("id")).First(&loc).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Location not found."})
		return
	}
	model.Database.Delete(&loc)
	msg := fmt.Sprint("Deleted location with id ", loc.ID)
	ctx.JSON(http.StatusOK, gin.H{"data": msg})

}

func GetDenue(ctx *gin.Context) {
	var loc model.Location
	if err := model.Database.Where("id = ?", ctx.Param("id")).First(&loc).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Location not found."})
		return
	}

	restchan := make(chan *http.Response)
	storchan := make(chan *http.Response)

	go service.GetDenueData(rest, loc.Lat, loc.Lon, restchan)
	go service.GetDenueData(stor, loc.Lat, loc.Lon, storchan)

	restResp := <-restchan
	defer restResp.Body.Close()

	storResp := <-storchan
	defer storResp.Body.Close()

	restBytes, _ := io.ReadAll(restResp.Body)
	storBytes, _ := io.ReadAll(storResp.Body)

	rests := json.RawMessage(string(restBytes))
	stores := json.RawMessage(string(storBytes))

	ctx.JSON(200, gin.H{"restaurant_data": rests, "store_data": stores})

}
