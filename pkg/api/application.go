package api

type Application struct {
	Name string
	Image string
	Version string
	
	Url string
	Port int32
	
	Environment map[string]string
}