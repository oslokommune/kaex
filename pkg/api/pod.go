package api

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	podTemplate = v1.Pod{
		TypeMeta:   metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{},
		Spec:       v1.PodSpec{
			Volumes:                       nil,
		},
	}
)

func CreatePod(app Application) (v1.Pod, error) {
	pod := podTemplate
	
	pod.ObjectMeta.Name = app.Name
	pod.ObjectMeta.Namespace = app.Namespace

	pod.Spec.Containers = CreateContainers(app)
	
	if app.ImagePullSecret != "" {
		pod.Spec.ImagePullSecrets = []v1.LocalObjectReference{{Name: app.ImagePullSecret}}
	}
	
	return pod, nil
}