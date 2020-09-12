#!/usr/bin/python
import sys
from ruamel.yaml import YAML

from lib.models import Application, Deployment, Ingress, Service

yaml = YAML(typ='safe')

application = Application(yaml.load(sys.stdin))

deployment = Deployment(application)
service = Service(application)
ingress = Ingress(application)

output = '---\n'.join([
    deployment.toYAML(),
    service.toYAML()
])

print(output)
