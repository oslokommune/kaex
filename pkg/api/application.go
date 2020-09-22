package api

type Application struct {
	Name string
	Image string
	Version string
	ImagePullSecret string
	
	Url string
	Port int32
	
	Replicas int32
	
	Environment map[string]string
}