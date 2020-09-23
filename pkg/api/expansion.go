package api

import (
	"fmt"
	"io"
	"sigs.k8s.io/yaml"
)

func Expand(w io.Writer, app Application, podonly bool) error {
	if app.Port != 0 {
		service, err := CreateService(app)
		if err != nil {
			return err
		}
		err = WriteResource(w, service)
		if err != nil {
			return err
		}
	}

	if app.Url != "" {
		ingress, err := CreateIngress(app)
		if err != nil {
			return err
		}
		err = WriteResource(w, ingress)
		if err != nil {
			return err
		}
	}

	if podonly == false {
		deployment, err := CreateDeployment(app)
		if err != nil {
			return err
		}
		err = WriteResource(w, deployment)
		if err != nil {
			return err
		}
	} else {
		pod, err := CreatePod(app)
		if err != nil {
			return err
		}
		err = WriteResource(w, pod)
		if err != nil {
			return err
		}
	}

	return nil
}

func WriteResource(w io.Writer, resource interface{}) error {
	serializedResource, err := yaml.Marshal(resource)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(w, "%s\n---\n", serializedResource)
	if err != nil {
		return err
	}

	return nil
}

