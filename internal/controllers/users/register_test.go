package users

import (
	"context"
	"testing"
	"time"

	"finworker/internal/models/requests/users"
	"finworker/internal/repositories"
	"finworker/internal/storage"
	"github.com/GeekchanskiY/migratigo"
	"github.com/jmoiron/sqlx"
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

	t.Run("success", func(t *testing.T) {
		reps, err := controller.RegisterUser(ctx, users.RegisterRequest{
			Username:          "test",
			Password:          "testPassword1234!",
			Name:              "test",
			Gender:            "male",
			Birthday:          time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
			PreferredBankName: "priorbank",
		})

		assert.NoError(t, err)
		assert.NotNil(t, reps)
	})

	t.Run("invalid user", func(t *testing.T) {
		reps, err := controller.RegisterUser(ctx, users.RegisterRequest{
			Username:          "test",
			Password:          "testPassword1234!",
			Name:              "test",
			Gender:            "male",
			Birthday:          time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
			PreferredBankName: "priorbank",
		})
		assert.Error(t, err)
		assert.Nil(t, reps)
	})
}
