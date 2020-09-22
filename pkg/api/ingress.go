package api

import (
	v1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/url"
	"strings"
)

var (
	ingressTemplate = v1.Ingress{
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
			Rules: []v1.IngressRule{},
		},
	}
)

func CreateIngress(app Application) (v1.Ingress, error) {
	url, err := url.Parse(app.Url)
	if err != nil {
		return v1.Ingress{}, err
	}
	
	ingress := ingressTemplate
	
	ingress.ObjectMeta.Name = app.Name

	ingress.Spec.Rules = append(ingress.Spec.Rules, v1.IngressRule{
		Host: url.Host,
	})
	
	if url.Scheme == "https" {
		ingress.Spec.TLS = []v1.IngressTLS{
			{
				Hosts: []string{
					url.Host,
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