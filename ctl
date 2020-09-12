#!/usr/bin/python
import sys
from ruamel.yaml import YAML

from lib.models import Deployment

yaml = YAML(typ='safe')

application_yaml = yaml.load(sys.stdin)

deployment = Deployment(application_yaml)

output = '---\n'.join([
    deployment.toYAML()
])

output = output.replace('!Deployment', '')

print(output)
