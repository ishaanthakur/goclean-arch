package author

import (
	"context"

	"github.com/ishaanthakur/goclean-arch/models"
)

// Repository represent the author's repository contract
type Repository interface {
	GetByID(ctx context.Context, id int64) (*models.Author, error)
}
