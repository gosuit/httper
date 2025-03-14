package httper

// Cookie represents an HTTP cookie with various attributes.
type Cookie struct {
	Name     string `yaml:"name"      env:"COOKIE_NAME"`
	Age      int    `yaml:"age"       env:"COOKIE_AGE"`
	Path     string `yaml:"path"      env:"COOKIE_PATH"`
	Host     string `yaml:"host"      env:"COOKIE_HOST"`
	Secure   bool   `yaml:"secure"    env:"COOKIE_SECURE"`
	HttpOnly bool   `yaml:"http_only" env:"COOKIE_HTTP_ONLY"`
}
