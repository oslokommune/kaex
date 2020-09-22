package api

import (
	v1 "k8s.io/api/apps/v1"
	v12 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	deploymentTemplate = v1.Deployment{
		TypeMeta:   metav1.TypeMeta{
			Kind:       "Deployment",
			APIVersion: "apps/v1",
		},
		ObjectMeta: metav1.ObjectMeta{},
		Spec:       v1.DeploymentSpec{
			Replicas:                nil,
			Selector:                &metav1.LabelSelector{},
			Template:                v12.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name:                       "",
					Annotations:                nil,
				},
				Spec:       v12.PodSpec{
					Volumes:                       nil,
				},
			},
		},
	}
)

func CreateContainers(app Application) []v12.Container {
	var envVars []v12.EnvVar
	for key, value := range app.Environment {
		envVars = append(envVars,v12.EnvVar{ Name: key, Value: value })
	}

	containers := []v12.Container{{
		Name: app.Name,
		Image: app.Image + ":" + app.Version,
		Env: envVars,
		VolumeMounts: nil,
	}}
	
	return containers
}

func CreateDeployment(app Application) (v1.Deployment, error) {
	deployment := deploymentTemplate
	
	deployment.ObjectMeta.Name = app.Name
	
	if app.Replicas == 0 {
		app.Replicas = 1
	}
	deployment.Spec.Replicas = &app.Replicas

	deployment.Spec.Selector.MatchLabels = map[string]string{
		"app": app.Name,
	}
	
	if app.ImagePullSecret != "" {
		deployment.Spec.Template.Spec.ImagePullSecrets = []v12.LocalObjectReference{
			{Name: app.ImagePullSecret},
		}
	}

	deployment.Spec.Template.ObjectMeta.Labels = map[string]string{
		"app": app.Name,
	}

	deployment.Spec.Template.Spec.Containers = CreateContainers(app)

	return deployment, nil
}