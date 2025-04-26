package distributors

import (
	"context"

	"finworker/internal/models"
)

func (repo *repository) Create(ctx context.Context, distributor *models.Distributor) (*models.Distributor, error) {
	q := `
		INSERT INTO 
    	distributors (name, source_wallet_id, target_wallet_id, percent) 
		VALUES (:name, :source_wallet_id, :target_wallet_id, :percent) 
		returning id`

	namedStmt, err := repo.db.PrepareNamed(q)
	if err != nil {
		return nil, err
	}

	err = namedStmt.QueryRowxContext(ctx, distributor).Scan(&distributor.ID)

	return distributor, err
}
