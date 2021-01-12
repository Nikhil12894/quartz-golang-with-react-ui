package main

import (
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/nikhil12894/quartz-golang-with-react-ui/server/scheduling"
)

func main() {
	var port = os.Getenv("PORT") //Getenv("PORT")
	if port == "" {
		port = "9091"
	}
	gin.SetMode(gin.ReleaseMode)
	r := setupRouter()
	// Listen and Server in 0.0.0.0:$PORT
	r.Run(":" + port)

}

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.New()

	// Serve frontend static files
	r.Use(static.Serve("/", static.LocalFile("./public/build", true)))
	api := r.Group("/api")
	// Ping test
	api.GET("/description", getDescription)
	// Get user value
	api.GET("/next/:n", nextN)
	return r
}

func getDescription(c *gin.Context) {
	expration, ok := c.GetQuery("expration")
	if !ok {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	description, err := scheduling.Describe(expration)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, gin.H{"data": description})

}

func nextN(c *gin.Context) {
	expration, ok := c.GetQuery("expration")
	if !ok {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	num, err := strconv.ParseUint(c.Param("n"), 10, 32)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	n := uint(num)
	description := scheduling.NextNScheduledTime(expration, n)
	c.JSON(http.StatusOK, gin.H{"data": description})
}
