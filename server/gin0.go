package server

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name       string    `form:"name"`
	Address    string    `form:"address"`
	Birthday   time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
	CreateTime time.Time `form:"createTime" time_format:"unixNano"`
	UnixTime   time.Time `form:"unixTime" time_format:"unix"`
}

func RUN_GIN_01() {

	// Creates a gin route with default middleware:
	// logger and recovery (crash-free) middleware
	route := gin.Default()

	route.GET("/testing", startPage)

	route.GET("/ping", func(ctx *gin.Context) {

		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})

	})

	// This handler will match /user/john but will not match /user/ or /user
	route.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/john/
	route.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	// For each matched request Context will hold the route definition
	route.POST("/user/:name/*action", func(c *gin.Context) {
		b := c.FullPath() == "/user/:name/*action" // true
		c.String(http.StatusOK, "%t", b)
	})

	// This handler will add a new router for /user/groups.
	// Exact routes are resolved before param routes, regardless of the order they were defined.
	// Routes starting with /user/groups are never interpreted as /user/:name/... routes
	route.GET("/user/groups", func(c *gin.Context) {
		c.String(http.StatusOK, "The available groups are [...]")
	})

	// Query string parameters are parsed using the existing underlying request object.
	// The request responds to an url matching:  /welcome?firstName=Jane&lastName=Doe
	route.GET("/welcome", func(c *gin.Context) {
		firstName := c.DefaultQuery("firstName", "Guest")
		lastName := c.Query("lastName") // shortcut for c.Request.URL.Query().Get("lastName")

		c.String(http.StatusOK, "Hello %s %s", firstName, lastName)
	})

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	route.Run()
	// router.Run(":3001") for a hard coded port
}

func startPage(ctx *gin.Context) {

	var person Person

	// If `GET`, only `Form` binding engine (`query`) used.
	// If `POST`, first checks the `content-type` for `JSON` or `XML`, then uses `Form` (`form-data`).
	// See more at https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L88
	if ctx.ShouldBind(&person) == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Birthday)
		log.Println(person.CreateTime)
		log.Println(person.UnixTime)
	}

	ctx.String(http.StatusOK, "Success")

	// test with:
	// curl -X GET "localhost:3000/testing?name=appleboy&address=xyz&birthday=1992-03-15&createTime=1562400033000000123&unixTime=1562400033"
}
