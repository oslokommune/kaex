package api

import "sigs.k8s.io/yaml"

type IngressConfig struct {
	Annotations map[string]string
}

type Application struct {
	Name string
	Namespace string
	
	Image string
	Version string
	ImagePullSecret string
	
	Url string
	Port int32
	
	Replicas int32
	
	Environment map[string]string
	Volumes []map[string]string
	
	Ingress IngressConfig
}

func ParseApplication(raw string) (Application, error) {
	var app Application

	err := yaml.Unmarshal([]byte(raw), &app)
	if err != nil {
		return Application{}, err
	}

	return app, nil
}
