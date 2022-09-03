// go mod init example/simple_go_api
// go get github.com/gin-gonic/gin
package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	// "errors"
)

type beer struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Origin_Place string `json:"origin_place"`
	Variety      string `json:"variety"`
	Qty			 int    `json:"qty"`
}

var beers = []beer{
	{Id: "1", Name: "Ayinger Celebrator Doppelbock", Origin_Place: "Germany", Variety: "weissbeer",Qty:3},
	{Id: "2", Name: "Bitburger Premium Pils", Origin_Place: "Germany", Variety: "weissbeer",Qty:3},
	{Id: "3", Name: "Erdinger Hefe Weissbier", Origin_Place: "Germany", Variety: "weissbeer",Qty:3},
	{Id: "4", Name: "Erdinger Hefe Weissbier Dark", Origin_Place: "Germany", Variety: "helles",Qty:3},
	{Id: "5", Name: "Franziskaner Hefe Weissbier", Origin_Place: "Germany", Variety: "weissbeer",Qty:3},
	{Id: "6", Name: "Gaffel Kölsch", Origin_Place: "Germany", Variety: "weissbeer",Qty:3},
	{Id: "7", Name: "Hofbräu Dunkel", Origin_Place: "Germany", Variety: "pils",Qty:3},
	{Id: "8", Name: "Hofbräu Münchner Kindl Weissbier", Origin_Place: "Germany", Variety: "weissbeer",Qty:3},
	{Id: "9", Name: "Hofbräu Original Premium Lager", Origin_Place: "Germany", Variety: "helles",Qty:3},
	{Id: "10", Name: "Klosterbrauerei Ettaler Dunkel", Origin_Place: "Germany", Variety: "pils",Qty:3},
	{Id: "11", Name: "Köstritzer Schwarzbier", Origin_Place: "Germany", Variety: "helles",Qty:3},
	{Id: "12", Name: "Paulaner Salvator", Origin_Place: "Germany", Variety: "pils",Qty:3},
	{Id: "13", Name: "Schlenkerla Helles Lagerbier", Origin_Place: "Germany", Variety: "pils",Qty:3},
	{Id: "14", Name: "Schlenkerla Rauchbier", Origin_Place: "Germany", Variety: "weissbeer",Qty:3},
	{Id: "15", Name: "Schneider Aventinus", Origin_Place: "Germany", Variety: "helles",Qty:3},
	{Id: "16", Name: "Schneider Aventinus Eisbock", Origin_Place: "Germany", Variety: "weissbeer",Qty:3},
	{Id: "17", Name: "Schneider Weisse", Origin_Place: "Germany", Variety: "pils",Qty:3},
	{Id: "18", Name: "Weihenstephaner Kristall Weissbier", Origin_Place: "Germany", Variety: "helles",Qty:3},}

func getBeers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, beers)
}

func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBeerById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "beer not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func getBeerById(id string) (*beer, error) {
	for i, b := range beers {
		if b.Id == id {
			return &beers[i], nil
		}
	}

	return nil, errors.New("beer not found")
}

func addBeer(c *gin.Context) {
	var newBeer beer

	if err := c.BindJSON(&newBeer); err != nil {
		return
	}

	beers = append(beers, newBeer)
	c.IndentedJSON(http.StatusCreated, newBeer)
}

func stockCountUpdate(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"'id' parameter missing"})
		return
	}

	beer, err := getBeerById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message":"beer not found"})
	}

	if beer.Qty <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"beer not available"})
		return
	}

	beer.Qty -= 1
	c.IndentedJSON(http.StatusOK, beer)
}

func main() {
	router := gin.Default()
	router.GET("/beers", getBeers)
	router.GET("/beers/:id", bookById)
	router.POST("/beers", addBeer)
	router.PATCH("/stock", stockCountUpdate)
	router.Run("localhost:3692")
}
