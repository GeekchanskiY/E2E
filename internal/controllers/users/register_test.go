package users

import (
	"context"
	"errors"
	"testing"
	"time"

	"finworker/internal/models"
	"finworker/internal/models/requests/users"
	"finworker/internal/repositories"
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

	repos := repositories.NewRepositories(db)

	controller := New(
		zap.NewNop(),
		repos.GetUsers(),
		repos.GetPermissionGroups(),
		repos.GetUserPermissions(),
		repos.GetWallets(),
		repos.GetBanks(),
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
		})

		assert.NoError(t, err)
		assert.NotNil(t, reps)

		// Check new user creation
		var newId int64
		q := `select id from  users where username = $1`
		err = db.QueryRow(q, newUserName).Scan(&newId)
		assert.NoError(t, err)
		assert.Equal(t, int64(1), newId)

		// Check new permission group creation
		q = `select id from permission_groups where name = $1`
		err = db.QueryRow(q, newUserName).Scan(&newId)
		assert.NoError(t, err)
		assert.Equal(t, int64(1), newId)

		// Check new user permission creation
		var newLevel string
		q = `select id, level from user_permission where user_id = 1 and permission_group_id = 1`
		err = db.QueryRow(q).Scan(&newId, &newLevel)
		assert.NoError(t, err)
		assert.Equal(t, int64(1), newId)
		assert.Equal(t, newLevel, string(models.AccessLevelOwner))

		// Check new wallet creation
		var newWalletName string
		var newWalletIsSalary bool
		q = `select id, name, is_salary from wallets where permission_group_id = 1`
		err = db.QueryRow(q).Scan(&newId, &newWalletName, &newWalletIsSalary)
		assert.NoError(t, err)
		assert.Equal(t, true, newWalletIsSalary)
		assert.Equal(t, int64(1), newId)
		assert.Equal(t, newUserName+"_salary", newWalletName)

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
		assert.Error(t, err)

		// Check that error is user already exists
		var pgErr *pq.Error
		if errors.As(err, &pgErr) {
			assert.Equal(t, "23505", string(pgErr.Code))
		}

		assert.Nil(t, reps)
	})
}
