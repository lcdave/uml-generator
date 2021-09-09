package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"upsee/models"

	"upsee/backend/db"
	"upsee/backend/repository"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
    database *db.Database
    serviceRepo repository.ServiceRepo
    logRepo repository.LogRepo
    emailLogRepo repository.EmailLogRepo
)

func main() {
    var err error
    database, err = db.Initialize()
    
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

	// Initialize a Gin router using Default.
    router := gin.Default()

    router.Use(cors.New(cors.Config{
        AllowOrigins: []string{"https://localhost:8080", "http://localhost:8080"},
        AllowMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
        ExposeHeaders:   []string{"Link"},
        AllowCredentials: false,
        MaxAge: 12 * time.Hour,
    }))

    logRepo = repository.NewLogRepo(database)
    serviceRepo = repository.NewServiceRepo(database, logRepo, emailLogRepo)
    emailLogRepo = repository.NewEmailLogRepo(database)

    serviceRepo.SetServicesStatus()

	// listen to /services and call getService function (GET Requests)
    router.GET("/services/", serviceRepo.GetServices)
    
    // listen to /services and call postService function (POST Requests)
	router.POST("/services/", serviceRepo.AddService)

	// listen to /services/:id and call getServiceById function (GET Requests)
	router.GET("/services/:id", serviceRepo.GetServiceByID)
	
    router.GET("/services/limited/:limit", serviceRepo.GetServiceByLimit)

    router.GET("/logs/", logRepo.GetLogs)
    router.GET("/logs/limited/:limit", logRepo.GetLogsByLimit)
    router.GET("/logs/count/", logRepo.CountLogs)

    router.GET("/emails/count/", countEmails)

    router.GET("/requests/:service_id", getRequestDataByService)

	//attach the router to an http.Server and start the server.
    router.Run(":8085")

}

/*
* Context is the most important part of gin.
* It carries request details, validates and serializes JSON, and more.
* https://pkg.go.dev/github.com/gin-gonic/gin#Context
*/


func getRequestDataByService(c *gin.Context) {
    // retrieve the id path parameter from the URL. 
	// When you map this handler to a path, you’ll include a placeholder for the parameter in the path.
    serviceId := c.Param("service_id")

    list := &models.RequestList{}

    sqlStatement := `SELECT * FROM request WHERE service_id = $1 ORDER BY created_at DESC LIMIT 10`
    rows, err := database.Context.Query(context.Background(), sqlStatement, serviceId)

    if err != nil {
        panic(err)
    }

    for rows.Next() {
        var r models.Request
        err := rows.Scan(&r.ID, &r.ServiceId, &r.ResponseTime, &r.CreatedAt)
        
        if err != nil {
            panic(err)
        }
        list.Requests = append(list.Requests, r)
    }

    c.IndentedJSON(http.StatusOK, list)
}


func countEmails(c *gin.Context) {
    var emailsCounter = 0;
    
    sqlStatement := `
    SELECT * FROM emails`
    rows, err := database.Context.Query(context.Background(), sqlStatement)

    if err != nil {
        panic(err)
    }

    for rows.Next() {
        emailsCounter++
    }

    c.IndentedJSON(http.StatusOK, emailsCounter)
}