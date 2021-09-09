package repository

import (
	"context"
	"net/http"
	"upsee/backend/db"
	"upsee/models"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

type LogRepo interface {
    WriteLogEntry(serviceId int, description string)
    GetLogs(c *gin.Context)
    CountLogs(c *gin.Context)
    GetLogsByLimit(c *gin.Context)
}

type logRepo struct {
	db *pgxpool.Pool
}

func NewLogRepo(database *db.Database) LogRepo {
    return &logRepo {
        db: database.Context,
    }
}

func (l *logRepo) WriteLogEntry(serviceId int, description string) {
    var logEntry models.LogEntry

    logEntry.ServiceId = serviceId
    logEntry.Description = description

    sqlStatement := `INSERT INTO logs (service_id, description) VALUES ($1, $2) RETURNING id, created_at`
    _, err := l.db.Exec(context.Background(), sqlStatement, logEntry.ServiceId, logEntry.Description)
    
    if err != nil {
        panic(err)
    }
}

func (l *logRepo) GetLogs(c *gin.Context) {
    list := &models.LogsOutput{}

    sqlStatement := `
    SELECT s.url as url, l.description as description, l.created_at as time 
    FROM services s
    RIGHT JOIN logs l ON s.id = l.service_id`
    rows, err := l.db.Query(context.Background(), sqlStatement)

    if err != nil {
        panic(err)
    }

    for rows.Next() {
        var r models.LogEntryOutput
        err := rows.Scan(&r.URL, &r.Description, &r.Time)
        
        if err != nil {
            panic(err)
        }
        list.LogsOutput = append(list.LogsOutput, r)
    }
    c.IndentedJSON(http.StatusOK, list)
}

func (l *logRepo) CountLogs(c *gin.Context) {
    var logsCounter = 0;
    
    sqlStatement := `
    SELECT * FROM logs`
    rows, err := l.db.Query(context.Background(), sqlStatement)

    if err != nil {
        panic(err)
    }

    for rows.Next() {
        logsCounter++
    }

    c.IndentedJSON(http.StatusOK, logsCounter)
}

func (l *logRepo) GetLogsByLimit(c *gin.Context) {
	
	// retrieve the id path parameter from the URL. 
	// When you map this handler to a path, you’ll include a placeholder for the parameter in the path.
    limit := c.Param("limit")

    list := &models.LogsOutput{}

    sqlStatement := `
    SELECT s.url as url, l.description as description, l.created_at as time 
    FROM services s
    RIGHT JOIN logs l ON s.id = l.service_id
    LIMIT $1;`
    rows, err := l.db.Query(context.Background(), sqlStatement, limit)

    if err != nil {
        panic(err)
    }

    for rows.Next() {
        var r models.LogEntryOutput
        err := rows.Scan(&r.URL, &r.Description, &r.Time)
        
        if err != nil {
            panic(err)
        }
        list.LogsOutput = append(list.LogsOutput, r)
    }

    c.IndentedJSON(http.StatusOK, list)
}