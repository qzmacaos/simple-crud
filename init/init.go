package initialization

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"gopkg.in/yaml.v2"
	"os"
	"simple-crud/config"
)

//reading config
func InitCfg() (config.Config, error)  {
	var (
		configFileData []byte
		cfg config.Config
	)
	file, err := os.Open("config/config.yml")
	defer file.Close()
	if err != nil {
		return cfg, err
	}

	_, err = file.Read(configFileData)
	if err != nil {
		return cfg, err
	}

	err = yaml.Unmarshal([]byte(configFileData), &cfg)
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}

//InitDB - init db connection
func InitDB(cfgDB config.CfgDB) (*sql.DB, error) {

	dataSource := fmt.Sprintf(
		"host=%v dbname=%v sslmode=%v user=%v password=%v",
		cfgDB.Host,
		cfgDB.DBName,
		cfgDB.SSL,
		cfgDB.User,
		cfgDB.Password,
		)

	db, err:= sql.Open("postgres", dataSource)
	if err != nil {
		return nil, err
	}

	return db, nil
}