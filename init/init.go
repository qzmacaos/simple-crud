package initialization

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"simple-crud/config"
)

//reading config
func InitCfg(pathConfig string) (config.Config, error)  {
	var (
		cfg config.Config
	)
	configFileData, err := ioutil.ReadFile(pathConfig)
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