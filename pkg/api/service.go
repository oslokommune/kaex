package api

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func CreateService(app Application) (v1.Service, error) {
	service := v1.Service{
		TypeMeta:   metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: app.Name,
		},
		Spec:       v1.ServiceSpec{
			Ports:                    []v1.ServicePort{
				{
					Port: 80,
					TargetPort: intstr.IntOrString{
						IntVal: app.Port,
					},
				},
			},
			Selector:                 map[string]string{
				"app": app.Name,
			},
			Type:                     "ClusterIP",
		},
	}
	
	return service, nil
}