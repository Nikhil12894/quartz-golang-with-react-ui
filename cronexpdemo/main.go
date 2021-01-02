package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/lnquy/cron"
	"github.com/nikhil12894/cronexpr"
)

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":9091")

}

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.New()
	gin.SetMode(gin.ReleaseMode)
	// Serve frontend static files
	r.Use(static.Serve("/", static.LocalFile("./public", true)))
	api := r.Group("/api")
	// Ping test
	api.GET("/description", getDescription)
	// Get user value
	api.GET("/next/:n", nextN)
	return r
}

func describe(exp string) (string, error) {
	exprDesc, err := cron.NewDescriptor(
		cron.Use24HourTimeFormat(true),
		cron.DayOfWeekStartsAtOne(false),
		cron.Verbose(true),
		cron.SetLogger(log.New(os.Stdout, "cron: ", 0)),
		cron.SetLocales(cron.Locale_en),
	)
	if err != nil {
		return "", fmt.Errorf("failed to create CRON expression descriptor: %s", err)
	}

	desc, err := exprDesc.ToDescription(exp, cron.Locale_en)
	if err != nil {
		return "", fmt.Errorf("failed to create CRON expression descriptor: %s", err)
	}
	return desc, nil
}

func nextNScheduledTime(exp string, n uint) []time.Time {
	// get the current time
	now := time.Now()
	// 1. Define two cronJob
	expr1 := cronexpr.MustParse(exp) // parse cron expression will be successful
	return expr1.NextN(now, n)

}

func getDescription(c *gin.Context) {
	expration, ok := c.GetQuery("expration")
	if !ok {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	description, err := describe(expration)
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
	description := nextNScheduledTime(expration, n)
	c.JSON(http.StatusOK, gin.H{"data": description})
}
