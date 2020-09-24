package api

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

var (
	serviceTemplate = v1.Service{
		TypeMeta:   metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{},
		Spec: v1.ServiceSpec{
			Ports: []v1.ServicePort{{ Port: 80 }},
			Type: "ClusterIP",
		},
	}
)

func CreateService(app Application) (v1.Service, error) {
	service := serviceTemplate
	
	service.ObjectMeta.Name = app.Name
	service.ObjectMeta.Namespace = app.Namespace

	service.Spec.Selector = map[string]string{
		"app": app.Name,
	}
	
	service.Spec.Ports[0].TargetPort = intstr.IntOrString{
		IntVal: app.Port,
	}

	return service, nil
}