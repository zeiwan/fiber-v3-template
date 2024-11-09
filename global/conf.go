package global

type config struct {
	Server server    `yaml:"server"`
	Log    log       `yaml:"log"`
	Mysql  mysql     `yaml:"mysql"`
	Redis  redisConf `yaml:"redis"`
	JWT    jwt       `yaml:"jwt"`
}

type log struct {
	Path       string `yaml:"path"`
	MaxSize    int    `yaml:"maxSize"`
	MaxAge     int    `yaml:"maxAge"`
	MaxBackups int    `yaml:"maxBackups"`
	Compress   bool   `yaml:"compress"`
	Level      int    `yaml:"level"`
}

type server struct {
	Name          string `yaml:"name"`
	URLPrefix     string `yaml:"urlPrefix"`
	Port          int    `yaml:"port"`
	BaseAPI       string `yaml:"baseApi"`
	ResourcePath  string `yaml:"resourcePath"`
	EnablePrefork bool   `yaml:"enablePrefork"`
}

type mysql struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Query    string `yaml:"query"`
}

type redisConf struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Port     int    `yaml:"port"`
	Host     string `yaml:"host"`
	Database int    `yaml:"database"`
}

type jwt struct {
	Timeout    int `yaml:"timeout"`
	MaxRefresh int `yaml:"maxRefresh"`
}
