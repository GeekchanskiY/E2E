package users

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"finworker/internal/models"
	"finworker/internal/models/requests/users"
	"finworker/internal/repositories/banks"
	"finworker/internal/repositories/operationGroups"
	"finworker/internal/repositories/operations"
	"finworker/internal/repositories/permission_groups"
	"finworker/internal/repositories/user_permissions"
	userRepo "finworker/internal/repositories/users"
	"finworker/internal/repositories/wallets"
	"finworker/internal/storage"

	"github.com/GeekchanskiY/migratigo"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.uber.org/zap"
)

func TestController_RegisterUser(t *testing.T) {
	ctx := context.Background()

	pgContainer, err := postgres.Run(ctx,
		"postgres:15.3-alpine",
		postgres.WithDatabase("test-db"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		t.Fatal(err)
	}

	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	assert.NoError(t, err)

	db, err := sqlx.Connect("postgres", connStr)
	assert.NoError(t, err)

	connector, err := migratigo.New(db.DB, storage.Migrations, "migrations", zap.NewNop())
	assert.NoError(t, err)

	err = connector.RunMigrations(false)
	assert.NoError(t, err)

	t.Cleanup(func() {
		if err := pgContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate pgContainer: %s", err)
		}
	})

	logger := zap.NewNop()

	controller := New(
		logger,
		userRepo.New(db, logger),
		permission_groups.New(db, logger),
		user_permissions.New(db, logger),
		wallets.New(db, logger),
		banks.New(db, logger),
		operationGroups.New(db, logger),
		operations.New(db, logger),
	)

	const newUserName = "newUser"

	t.Run("success", func(t *testing.T) {

		reps, err := controller.RegisterUser(ctx, users.RegisterRequest{
			Username:          newUserName,
			Password:          "testPassword1234!",
			Name:              "test",
			Gender:            "male",
			Birthday:          time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
			PreferredBankName: "priorbank",
			Salary:            2000,
			SalaryCurrency:    "USD",
			SalaryDate:        time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
		})

		require.NoError(t, err)
		assert.NotNil(t, reps)

		// Check new user creation
		var newID int64
		q := `select id from  users where username = $1`
		err = db.QueryRow(q, newUserName).Scan(&newID)
		require.NoError(t, err)

		assert.Equal(t, int64(1), newID)

		// Check new permission group creation
		q = `select id from permission_groups where name = $1`
		err = db.QueryRow(q, newUserName).Scan(&newID)
		require.NoError(t, err)

		assert.Equal(t, int64(1), newID)

		// Check new user permission creation
		var newLevel string
		q = `select id, level from user_permission where user_id = 1 and permission_group_id = 1`
		err = db.QueryRow(q).Scan(&newID, &newLevel)
		require.NoError(t, err)

		assert.Equal(t, int64(1), newID)
		assert.Equal(t, newLevel, string(models.AccessLevelOwner))

		// Check new wallet creation
		var newWalletName string
		var newWalletIsSalary bool
		q = `select id, name, is_salary from wallets where permission_group_id = 1`
		err = db.QueryRow(q).Scan(&newID, &newWalletName, &newWalletIsSalary)
		require.NoError(t, err)

		assert.Equal(t, true, newWalletIsSalary)
		assert.Equal(t, int64(1), newID)
		assert.Equal(t, newUserName+"_salary", newWalletName)

		q = `select id, name from operation_groups where wallet_id = 1`
		err = db.QueryRow(q).Scan(&newID, &newWalletName)
		require.NoError(t, err)

		assert.Equal(t, int64(1), newID)
		assert.Equal(t, newUserName+"_salary", newWalletName)

		var operation models.Operation
		q = `select id, amount, time, is_monthly from operations where operation_group_id = 1`
		err = db.QueryRow(q).Scan(&operation.ID, &operation.Amount, &operation.Time, &operation.IsMonthly)
		require.NoError(t, err)

		assert.Equal(t, int64(1), operation.ID)
		assert.Equal(t, float64(2000), operation.Amount)
		assert.Equal(t, true, operation.IsMonthly)
		assert.Equal(t, time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC), operation.Time.UTC())

	})

	t.Run("invalid user", func(t *testing.T) {
		reps, err := controller.RegisterUser(ctx, users.RegisterRequest{
			Username:          newUserName,
			Password:          "testPassword1234!",
			Name:              "test",
			Gender:            "male",
			Birthday:          time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
			PreferredBankName: "priorbank",
		})
		require.Error(t, err)

		// Check that error is user already exists
		var pgErr *pq.Error
		if errors.As(err, &pgErr) {
			assert.Equal(t, "23505", string(pgErr.Code))
		}

		assert.Nil(t, reps)
	})
}
