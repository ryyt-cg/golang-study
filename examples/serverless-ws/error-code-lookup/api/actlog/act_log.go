package actlog

import (
	"time"

	"github.com/gofrs/uuid"
)

type ActivityLog struct {
	ID          uuid.UUID
	Description string
	Initial     string
	CreatedAt   time.Time
}
