package repository

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptrace"
	"os"
	"time"
	"upsee/backend/db"
	"upsee/backend/utils"
	"upsee/models"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ServiceRepo interface {
    GetServiceByID(c *gin.Context)
    GetServiceByLimit(c *gin.Context)
    AddService(c *gin.Context)
    GetServices(c *gin.Context)
    SetServicesStatus()
    GetServiceStatusFromDB(id int) string
    CheckService(id int, url string) string
}

type serviceRepo struct {
	db *pgxpool.Pool
    logRepo LogRepo
    emailLogRepo EmailLogRepo
}

func NewServiceRepo(database *db.Database, logRepo LogRepo, emailLogRepo EmailLogRepo) ServiceRepo {
    return &serviceRepo {
        db: database.Context,
        logRepo: logRepo,
        emailLogRepo: emailLogRepo,
    }
}

// getServiceByID locates the service whose ID value matches the id
// parameter sent by the client, then returns that Service as a response.
func (s *serviceRepo) GetServiceByID(c *gin.Context) {
	// retrieve the id path parameter from the URL. 
	// When you map this handler to a path, you’ll include a placeholder for the parameter in the path.
    id := c.Param("id")

    sqlStatement := `SELECT * FROM services WHERE id = $1`
    _, err := s.db.Exec(context.Background(), sqlStatement, id)

    if err != nil {
        panic(err)
    }
	// return an HTTP 404 error with http.StatusNotFound if the Service isn’t found.
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "service not found"})
}

func (s *serviceRepo) GetServiceByLimit(c *gin.Context) {
	
	// retrieve the id path parameter from the URL. 
	// When you map this handler to a path, you’ll include a placeholder for the parameter in the path.
    limit := c.Param("limit")

    list := &models.ServiceList{}

    sqlStatement := `SELECT * FROM services LIMIT $1`
    rows, err := s.db.Query(context.Background(), sqlStatement, limit)

    if err != nil {
        panic(err)
    }

    for rows.Next() {
        var r models.Service
        err := rows.Scan(&r.ID, &r.Name, &r.URL, &r.CreatedAt, &r.Status)
        
        if err != nil {
            panic(err)
        }
        list.Services = append(list.Services, r)
    }

    c.IndentedJSON(http.StatusOK, list)
}

// addServices adds an service from JSON received in the request body.
func (s *serviceRepo) AddService(c *gin.Context) {
    var newService models.Service
    c.BindJSON(&newService)
    c.JSON(http.StatusOK, gin.H{"service": newService.Name})

    newService.Status = "online"

    sqlStatement := `INSERT INTO services (name, url, status) VALUES ($1, $2, $3) RETURNING id, created_at`
    _, err := s.db.Exec(context.Background(), sqlStatement, newService.Name, newService.URL, newService.Status)
    if err != nil {
        panic(err)
    }
}

func (s *serviceRepo) GetServices(c *gin.Context) {
    list := &models.ServiceList{}

    sqlStatement := `SELECT * FROM services`
    rows, err := s.db.Query(context.Background(), sqlStatement)

    if err != nil {
        panic(err)
    }

    
    for rows.Next() {
        var r models.Service
        err := rows.Scan(&r.ID, &r.Name, &r.URL, &r.CreatedAt, &r.Status)
        
        if err != nil {
            panic(err)
        }
        list.Services = append(list.Services, r)
    }

    c.IndentedJSON(http.StatusOK, list)
}

func (s *serviceRepo) SetServicesStatus() {
    sqlStatement := `SELECT * FROM services`
    rows, err := s.db.Query(context.Background(), sqlStatement)

    if err != nil {
        panic(err)
    }

    
    for rows.Next() {
        var r models.Service
        err := rows.Scan(&r.ID, &r.Name, &r.URL, &r.CreatedAt, &r.Status)
        
        if err != nil {
            panic(err)
        }

        currentStatus := s.CheckService(r.ID, r.URL)

        //fmt.Printf("row STATUS: ", r.Status)
        //fmt.Printf("Current STATUS: ", currentStatus)

        if (r.Status != currentStatus) {
            //fmt.Printf("entering update query logic")

            updateStatement := `UPDATE services SET status = $1 WHERE id = $2`
            _, err := s.db.Exec(context.Background(), updateStatement, currentStatus, r.ID)
            
            if err != nil {
                panic(err)
            }
        }
    }

    // Call this func every minute
    time.AfterFunc(time.Minute, s.SetServicesStatus)
}

func (s *serviceRepo) GetServiceStatusFromDB(id int) string {
    sqlStatement := `SELECT status FROM services WHERE id = $1`
    rows, err := s.db.Query(context.Background(), sqlStatement, id)
    
    var status string

    if err != nil {
        panic(err)
    }

    for rows.Next() {
        err := rows.Scan(&status)
        
        if err != nil {
            panic(err)
        }
    }

   return status
}

// TODO: This function doesn't belong in a repository. Added it here because of golang strict circular depency handling
func (s *serviceRepo) CheckService(id int, url string) string {
    var status = s.GetServiceStatusFromDB(id)
	
    // start timer for counting request reaction time
    start := time.Now()

    // GET request to requestet url
	req, _ := http.NewRequest("GET", url, nil)
    
    // trace is needed to check different states of the request 
    // (connection started, connection done, handshake done and so on)
	trace := &httptrace.ClientTrace{}
    
    // change the request context to "ctx"
    // the context controls the lifetime of a request + its response
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
    
    // Roundtrip executes a HTTP transaction
    // if there is no response, err will contain a error message (meaning that the service is down)
	if _, err := http.DefaultTransport.RoundTrip(req); err != nil {
		if(status != "offline") {
            if err := utils.SendMail(id, url); err != nil {
                panic(err)
            }

            s.emailLogRepo.CreateEmailLogEntry(id, os.Getenv("MAS_EMAIL_TO"))

            status = "offline"
        }
        s.logRepo.WriteLogEntry(id, err.Error())

    } else {
        if (status == "offline") {
            s.logRepo.WriteLogEntry(id, "service is back online")
            fmt.Println("Log entry written for" + url)
        }
        status = "online"
    }

	reaction_time := time.Since(start).Milliseconds()

    sqlStatement := `INSERT INTO request (service_id, reaction_time) VALUES ($1, $2) RETURNING id, created_at`
    _, err := s.db.Exec(context.Background(), sqlStatement, id, reaction_time)
    
    if err != nil {
        panic(err)
    }

    req.Close = true

    return status
}
