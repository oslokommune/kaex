#!/usr/bin/python
import sys
from ruamel.yaml import YAML

from lib.models import Application, Deployment

yaml = YAML(typ='safe')

application = Application(yaml.load(sys.stdin))

deployment = Deployment(application)

output = '---\n'.join([
    deployment.toYAML()
])

output = output.replace('!Deployment', '')

print(output)
