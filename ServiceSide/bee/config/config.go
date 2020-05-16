package config

type Server struct {
	Mysql   Mysql    `mapstructure:"mysql" json:"mysql"`
	JWT     JWT      `mapstructure:"jwt" json:"jwt"`
	System 	System	 `mapstructure:"system" json:"system"`
	Casbin  Casbin  `mapstructure:"casbin" json:"casbin"`

}

type Mysql struct {
	Username     string `mapstructure:"username" json:"username"`
	Password     string `mapstructure:"password" json:"password"`
	Path         string `mapstructure:"path" json:"path"`
	Dbname       string `mapstructure:"db-name" json:"dbname"`
	Config       string `mapstructure:"config" json:"config"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns"`
	LogMode      bool   `mapstructure:"log-mode" json:"logMode"`
}

type JWT struct {
	SigningKey string `mapstructure:"signing-key" json:"signingKey"`
}

type Casbin struct {
	ModelPath string `mapstructure:"model-path" json:"modelPath"`
}

type System struct {
	UserMultiPoint bool 	`mapstructure:"use-multipoint" json:"useMultipoint"`
	Env 		   string 	`mapstructure:"env" json:"env"`
}