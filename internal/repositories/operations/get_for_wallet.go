package operations

import (
	"context"

	"finworker/internal/models"
)

func (r *repository) GetForWallet(ctx context.Context, id int64) ([]*models.OperationExtended, error) {
	q := `select 
    	o.id, o.operation_group_id, o.amount, o.time, o.is_monthly, o.is_confirmed, o.initiator_id, 
		og.name as operation_group_name,
		u.username as initiator_name
	from 
	    operations o
	join 
	    operation_groups og on og.id = o.operation_group_id
	join
	    users u on u.id = o.initiator_id
	where 
	    og.wallet_id = $1`

	operations := make([]*models.OperationExtended, 0)

	err := r.db.SelectContext(ctx, &operations, q, id)
	if err != nil {
		return nil, err
	}

	return operations, nil
}
