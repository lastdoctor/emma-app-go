package repository

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

//import (
//	"database/sql"
//	"errors"
//	"github.com/lastdoctor/emma-app-go/internal/domain"
//)
type PostgresConfig struct {
	HOST      string `mapstructure:"PG_HOST"`
	User      string `mapstructure:"PG_USER"`
	Password  string `mapstructure:"PG_PASSWORD"`
	Database  string `mapstructure:"PG_DATABASE"`
	Port      int    `mapstructure:"PG_PORT"`
	OpenConns int    `mapstructure:"PG_CONNS"`
	IdleConns int    `mapstructure:"PG_IDLE"`
}

func pgConfig() (*PostgresConfig, error) {
	var pgCfg PostgresConfig
	const path = "."
	viper.AddConfigPath(path)
	viper.SetConfigName("development")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		zap.L().Error("reading .env is failed", zap.Error(err))
		return nil, err
	}

	err = viper.Unmarshal(&pgCfg)
	return &pgCfg, nil
}

//var (
//	ErrRecordNotFound = errors.New("record not found")
//)
//
//type Users interface {
//	Create(user *Users) error
//	GetById(id int64) (*domain.User, error)
//}
//
//type UsersRepos struct {
//	DB *sql.DB
//}
//
//type Merchants interface {
//	Create(merchant *domain.Merchants) error
//	GetById(id int64) (*domain.Merchants, error)
//}
//
//type MerchantsRepos struct {
//	DB *sql.DB
//}
//
//type Transactions interface {
//	Create()
//}
//
//type Repositories struct {
//	Users        Users
//	Merchants    Merchants
//	Transactions Transactions
//}
//
//func InitRepositories(db *sql.DB) *Repositories {
//return &Repositories{
//	Users:        UsersRepos{DB: db},
//	Merchants:    MerchantsRepos{DB: db},
//	Transactions: nil,
//}
//}
