package global

type config struct {
	Server `yaml:"server"`
	Log    `yaml:"log"`
}
type Log struct {
	Path       string `yaml:"path"`
	MaxSize    int    `yaml:"maxSize"`
	MaxAge     int    `yaml:"maxAge"`
	MaxBackups int    `yaml:"maxBackups"`
	Compress   bool   `yaml:"compress"`
	Level      int    `yaml:"level"`
}
type Server struct {
	Name         string `yaml:"name"`
	URLPrefix    string `yaml:"urlPrefix"`
	Port         int    `yaml:"port"`
	BaseAPI      string `yaml:"baseApi"`
	ResourcePath string `yaml:"resourcePath"`
}
