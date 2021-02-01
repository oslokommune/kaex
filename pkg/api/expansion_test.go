package api_test

import (
	"bytes"
	"github.com/oslokommune/kaex/pkg/api"
	"github.com/sebdah/goldie/v2"
	"testing"
)

func TestExpand(t *testing.T) {
	testCases := []struct {
		name string

		withApp api.Application
	}{
		{
			name: "Should produce expected output when expanding a simple app",

			withApp: api.Application{
				Name:    "dummy-app",
				Image:   "dummygres",
				Version: "0.1.0",
			},
		},
		{
			name: "Should produce expected output when expanding app with optionals",

			withApp: api.Application{
				Name:            "dummy-app",
				Namespace:       "dummyns",
				Image:           "dummygres",
				Version:         "8.0.1",
				ImagePullSecret: "so-secret",
				Replicas:        3,
				Environment: map[string]string{
					"DUMMY_VAR":  "avalue",
					"DUMMY_HOST": "somehost",
				},
			},
		},
		{
			name: "Should produce expected output when expanding app with service",

			withApp: api.Application{
				Name:    "dummy-app",
				Image:   "dummyredis",
				Version: "8.2.1",
				Port:    3000,
			},
		},
		{
			name: "Should produce expected output when expanding app with service and ingress",

			withApp: api.Application{
				Name:    "dummy-app",
				Image:   "dummyredis",
				Version: "8.2.1",
				Port:    3000,
				Url:     "http://dummy.io",
			},
		},
		{
			name: "Should produce expected output when expanding app with annotated ingress",

			withApp: api.Application{
				Name:    "dummy-app",
				Image:   "dummyredis",
				Version: "8.2.1",
				Port:    3000,
				Url:     "http://dummy.io",
				Ingress: api.IngressConfig{
					Annotations: map[string]string{
						"cert-manager.io/cluster-issuer": "letsencrypt-production",
					},
				},
			},
		},
		{
			name: "Should produce expected output when expanding app with service and tls enabled ingress",

			withApp: api.Application{
				Name:    "dummy-app",
				Image:   "dummyredis",
				Version: "8.2.1",
				Port:    3000,
				Url:     "https://dummy.io",
			},
		},
		{
			name: "Should produce expected output when expanding app with volumes",

			withApp: api.Application{
				Name:    "dummy-app",
				Image:   "dummygres",
				Version: "0.1.0",
				Volumes: []map[string]string{
					{"/etc/config": "4Gi"},
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			buf := bytes.NewBufferString("")

			err := api.Expand(buf, tc.withApp, false)
			if err != nil {
				t.Fatal(err)
			}

			g := goldie.New(t)
			g.Assert(t, t.Name(), buf.Bytes())
		})
	}
}
