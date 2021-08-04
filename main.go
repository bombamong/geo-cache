package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bombamong/geo-cache/pkg/cache"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: r,
	}
	ca := cache.Cache{
		Data: make([]cache.Row, 0),
	}

	fmt.Println("start reader")
	file, err := os.Open("./tmp/geo-api_v3_polygon.csv")
	if err != nil {
		fmt.Println(err.Error())
	}
	br := bufio.NewReader(file)

	var line string
	_, err = br.ReadString('\n')

	for err == nil {
		line, err = br.ReadString('\n')
		r := new(cache.Row)
		r.ParseLine(line)
		ca.Data = append(ca.Data, *r)
	}

	r.GET("/:merchant", func(c *gin.Context) {
		p, ok := c.Params.Get("merchant")
		if !ok {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "bad request"})
			return
		}
		md := ca.Query(p)
		c.JSON(http.StatusOK, map[string]interface{}{
			p: map[string]int{
				"purchaser_count":   len(md.Purchasers),
				"transaction_count": md.Transactions,
				"transaction_sum":   md.TransactionSum,
			},
		})
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "up and running..."})
	})

	fmt.Println("Start server")
	err = server.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}
