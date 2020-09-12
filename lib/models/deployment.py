import sys
from ruamel.yaml import YAML

from .resource import Resource

yaml = YAML(typ='safe')

def generateContainer(app):
    container = dict()

    image = ''
    env = list()
    volumeMounts = list()

    for key, value in app["environment"].items():
        env.append({
            'name': key,
            'value': value
        })

    return {
        'name': 'app',
        'image': f'{app["image"]}:{app["version"]}',
        'env': env,
        'volumeMounts': volumeMounts
    }

def generateTemplateSpec(app):
    return {
        'containers': [generateContainer(app)]
    }

def generateTemplate(app):
    metadata = dict()

    return {
        'metadata': metadata,
        'spec': generateTemplateSpec(app)
    }

def generateDeploymentSpec(app):
    return {
        'replicas': app.get('replicas', 1),
        'template': generateTemplate(app)
    }

class Deployment(Resource):
    def __init__(self, app):
        self.apiVersion = 'extensions/v1beta1'
        self.kind = 'Deployment'

        self.metadata = {
            'name': app['name']
        }

        self.spec = generateDeploymentSpec(app)

yaml.register_class(Deployment)
