package controllers


import (
	"strconv"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"

	"github.com/gin-gonic/gin"
	"github.com/gonzalonegre/sandbox1/lib/services/clima"
	"github.com/gonzalonegre/sandbox1/planet"
)

// ResponseObject is a simple mapping object
type ResponseObject map[string]interface{}

// CreateDias creates pronostico para dias conocidos de 0 a 359 solamente y cuando no sean unknown.
func CreateDias(c *gin.Context) {
	var err error
	ctx := appengine.NewContext(c.Request)


	P1 := planet.Planet{ 500, 0, -1}
	P2 := planet.Planet{2000, 0, -3}
	P3 := planet.Planet{1000, 0,  5}

	for i:=0 ; i<planet.TOTAL_ANGLE ; i++ {
		currentWeather, _ := planet.WeatherForDay(&P1, &P2, &P3, i)
		var clm clima.Clima
		clm.Dia = i
		clm.Estado = currentWeather.String()
		if _, err = clima.Create(ctx, &clm); err != nil {
			break;
		}
	}

	if err != nil {
		log.Errorf(ctx, "ERROR: %v", err.Error())
		c.JSON(http.StatusPreconditionFailed, ResponseObject{"error": err.Error()})
	}else{
		c.JSON(http.StatusOK, nil)
	}
}

// GetDia based on its Dia
func GetDia(c *gin.Context) {
	var err error
	var output *clima.Clima
	reqDia := c.Param("dia")
	ctx := appengine.NewContext(c.Request)
        value,_ :=strconv.Atoi(reqDia)


	if output, err = clima.GetByDia(ctx, value); err == nil {
		c.JSON(http.StatusOK, output)
	}
	if err != nil {
		c.JSON(http.StatusPreconditionFailed, ResponseObject{"error": err.Error()})
	}
}

// GetDias Fectch all users
func GetDias(c *gin.Context) {
	var err error
	ctx := appengine.NewContext(c.Request)
	var output []clima.Clima

	if output, err = clima.GetDias(ctx); err == nil {
		c.JSON(http.StatusOK, output)
	}

	if err != nil {
		c.JSON(http.StatusPreconditionFailed, ResponseObject{"error": err.Error()})
	}
}

// UpdateDia Updates an clima
func UpdateDia(c *gin.Context) {
	var clm *clima.Clima
	var err error
	var output *clima.Clima
	ctx := appengine.NewContext(c.Request)

	if err = c.BindJSON(&clm); err == nil {
		if output, err = clima.Update(ctx, clm); err == nil {
			c.JSON(http.StatusOK, ResponseObject{"clima": output})
		}
	}

	if err != nil {
		log.Errorf(ctx, "ERROR: %v", err.Error())
		c.JSON(http.StatusPreconditionFailed, ResponseObject{"error": err.Error()})
	}
}

// DeleteDias deletes an clima based on its dia
func DeleteDias(c *gin.Context) {
	reqDia := c.Param("dia")
	ctx := appengine.NewContext(c.Request)
        value,_ :=strconv.Atoi(reqDia)

	err := clima.Delete(ctx, value)

	if err != nil {
		log.Errorf(ctx, "ERROR: %v", err.Error())
		c.JSON(http.StatusPreconditionFailed, ResponseObject{"error": err.Error()})
	}
	c.JSON(http.StatusOK, ResponseObject{"result": "ok"})
}