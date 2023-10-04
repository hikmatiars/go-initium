package server

import (
	"context"
	"fmt"
	"go-initium/internal/config"
	"gorm.io/gorm/schema"
	"log"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Contract struct {
	DB *gorm.DB
}

func New(ctx context.Context, cfg config.Config) *Contract {
	// connect to redis
	cache := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.Db,
	})
	if _, err := cache.Ping(ctx).Result(); err != nil {
		log.Fatalf("failed connect to redis: %s", err)
	}

	//Connect to DB
	db := SetupPostgres(cfg)

	return &Contract{DB: db}
}

func SetupPostgres(cfg config.Config) *gorm.DB {
	dsn := fmt.
		Sprintf(`
			host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta`,
			cfg.Postgres.DbHost,
			cfg.Postgres.DbPort,
			cfg.Postgres.DbUser,
			cfg.Postgres.DbPass,
			cfg.Postgres.DbName,
		)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Successfully connecting DB")

	return db
}
