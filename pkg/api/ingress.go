package api

import (
	v1 "k8s.io/api/networking/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"net/url"
	"strings"
)

func generateDefaultIngress() v1.Ingress {
	return v1.Ingress{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Ingress",
			APIVersion: "networking.k8s.io/v1beta1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        "name",
			Labels:      nil,
			Annotations: nil,
		},
		Spec: v1.IngressSpec{
			Rules: []v1.IngressRule{{}},
		},
	}
}

func CreateIngress(app Application) (v1.Ingress, error) {
	hostUrl, err := url.Parse(app.Url)
	if err != nil {
		return v1.Ingress{}, err
	}
	
	ingress := generateDefaultIngress()
	ingress.ObjectMeta.Namespace = app.Namespace
	
	ingress.ObjectMeta.Name = app.Name
	
	ingress.Spec.Rules = append(ingress.Spec.Rules, v1.IngressRule{
		Host: hostUrl.Host,
		IngressRuleValue: v1.IngressRuleValue{
			HTTP: &v1.HTTPIngressRuleValue{
				Paths: []v1.HTTPIngressPath{{
					Path:     "/",
					Backend:  v1.IngressBackend{
						ServiceName: app.Name,
						ServicePort: intstr.IntOrString{
							IntVal: 80,
						},
					},
				}},
			},
		},
	})
	
	if hostUrl.Scheme == "https" {
		ingress.Spec.TLS = []v1.IngressTLS{
			{
				Hosts: []string{
					hostUrl.Host,
				},
				SecretName: strings.Join([]string{
					app.Name,
					"tls",
				}, "-"),
			},
		}
	}

	return ingress, nil
}