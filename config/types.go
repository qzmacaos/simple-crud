package config

type Config struct {
	CfgDB CfgDB    `yaml:"db"`
	Http  CfgHTTP  `yaml:"http"`
	Other CfgOther `yaml:"other"`
}

type CfgDB struct {
	Host     string `yaml:"host""`
	DBName   string `yaml:"dbname"`
	SSL      string `yaml:"sslmode"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type CfgHTTP struct {
	Port string `yaml:"port"`
}

type CfgOther struct {
	PageSize uint32 `yaml:"pagesize"`
}
