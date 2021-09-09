package repository

import (
	"context"
	"upsee/backend/db"
	"upsee/models"

	"github.com/jackc/pgx/v4/pgxpool"
)

type EmailLogRepo interface {
    CreateEmailLogEntry(serviceId int, recipient string)
}

type emailLogRepo struct {
	db *pgxpool.Pool
}

func NewEmailLogRepo(database *db.Database) EmailLogRepo {
    return &emailLogRepo {
        db: database.Context,
    }
}
func (e *emailLogRepo) CreateEmailLogEntry(serviceId int, recipient string) {
    var mailLogEntry models.MailLogEntry

    mailLogEntry.ServiceId = serviceId
    mailLogEntry.Recipient = recipient


    sqlStatement := `INSERT INTO emails (service_id, recipient) VALUES ($1, $2) RETURNING id, created_at`
    _, err := e.db.Exec(context.Background(), sqlStatement, mailLogEntry.ServiceId, mailLogEntry.Recipient)
    
	if err != nil {
        panic(err)
    }
}