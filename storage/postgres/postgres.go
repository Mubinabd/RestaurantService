package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Mubinabd/RestaurantService/config"
	"github.com/Mubinabd/RestaurantService/storage"
)

type Storage struct {
	db          *sql.DB
	RestaurantS storage.RestaurantI
}

func ConnectDB() (*Storage, error) {
	cfg := config.Load()
	dbConn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase)
	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	restaurantS := NewRestaurantStorage(db)
	return &Storage{
		db:    db,
		RestaurantS: restaurantS,
	}, nil
}
func (s *Storage) Restaurant() storage.RestaurantI {
	if s.RestaurantS == nil {
		s.RestaurantS = NewRestaurantStorage(s.db)
	}
	return s.RestaurantS
}
