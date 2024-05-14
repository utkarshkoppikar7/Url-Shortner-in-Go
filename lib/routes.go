package lib

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"net/http"
	"os"
	"strconv"
)

//const short_url = "uko.kr/"

var CurrentDns = os.Getenv("HOSTNAME_DNS")
var shortUrl = CurrentDns + "/url?q="

var ctx = context.Background()

var rdb = redis.NewClient(&redis.Options{
	Addr:     os.Getenv("REDIS_URL"),
	Password: "", // no password set
	DB:       0,  // use default DB
})

func ShortenNewUrl(c *gin.Context) {
	urlId := rdb.DBSize(ctx).Val()
	fmt.Println(urlId)
	urlId++

	var urlBody struct {
		Url string `json:"url"`
	}

	err := c.Bind(&urlBody)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	shortenedUrl := shortUrl + strconv.Itoa(int(urlId))

	fmt.Print(urlBody)

	// Prepare statement for inserting data

	err = rdb.Set(ctx, strconv.Itoa(int(urlId)), urlBody.Url, 0).Err()
	if err != nil {
		panic(err)
	}

	fmt.Println("Data inserted successfully")

	c.JSON(200, gin.H{
		"longUrl":  urlBody.Url,
		"shortUrl": shortenedUrl,
	})

}

func GetLongUrl(c *gin.Context) {
	req := c.Request
	err := req.ParseForm()

	r := req.Form
	for k, v := range r {
		fmt.Println(fmt.Sprintf("%s=%v", k, v))
	}

	if err != nil {
		return
	}
	urlString := req.FormValue("q")
	fmt.Println(urlString)

	longUrl, err := rdb.Get(ctx, urlString).Result()

	fmt.Println(longUrl)

	if errors.Is(err, redis.Nil) {
		fmt.Println("Url does not exist in DB. Please insert it first")
	} else if err != nil {
		fmt.Println(err)
		panic(err)
	}

	c.Redirect(http.StatusFound, longUrl)
}


func Hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"response":  "Hello Welcome to URL Shortner"
	})
}