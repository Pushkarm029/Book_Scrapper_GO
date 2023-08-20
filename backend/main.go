package main

import (
	"ebook-finder-app/handlers"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	initRoutes(r)
	r.Run(":8080")
}

func amazonRequestHandler(c *gin.Context) {
	searchQuery := c.Param("search")
	amazonHandlerPackets, err := handlers.AmazonHandler(searchQuery)
	if err != nil {
		log.Print("Error Fetching Data From Amazon", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed To Get Books From Amazon"})
		return
	}
	c.JSON(http.StatusOK, amazonHandlerPackets)
}

func gutenbergRequestHandler(c *gin.Context) {
	searchQuery := c.Param("search")
	HandlerPackets, err := handlers.GutenbergHandler(searchQuery)
	if err != nil {
		log.Print("Error Fetching Data From Gutenberg", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed To Get Books From Gutenberg"})
		return
	}
	c.JSON(http.StatusOK, HandlerPackets)
}

func oceanofPDFRequestHandler(c *gin.Context) {
	searchQuery := c.Param("search")
	HandlerPackets, err := handlers.OceanOfPDFHandler(searchQuery)
	if err != nil {
		log.Print("Error Fetching Data From OceanOfPDF", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed To Get Books From Ocean of PDF"})
		return
	}
	c.JSON(http.StatusOK, HandlerPackets)
}

func archiveRequestHandler(c *gin.Context) {
	searchQuery := c.Param("search")
	HandlerPackets, err := handlers.ArchiveHandler(searchQuery)
	if err != nil {
		log.Print("Error Fetching Data From Archives", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed To Get Books From Archives"})
		return
	}
	c.JSON(http.StatusOK, HandlerPackets)
}

func initRoutes(r *gin.Engine) {
	r.GET("/api/amazon/:search", amazonRequestHandler)
	r.GET("/api/gutenberg/:search", gutenbergRequestHandler)
	r.GET("/api/oceanofpdf/:search", oceanofPDFRequestHandler)
	r.GET("/api/archive/:search", archiveRequestHandler)
}
