from kaex.models.resource import Resource
from kaex.models.pvc import generatePVCName


def generateContainer(app):
    container = dict()

    volumeMounts = list()
    for volume in app.volumes:
        volumeMounts.append({
            'mountPath': volume['path'],
            'name': generatePVCName(app.name, volume['path'])
        })

    return {
        'name': 'app',
        'image': f'{app.image["uri"]}:{app.image["version"]}',
        'env': app.env,
        'volumeMounts': volumeMounts
    }


def generateTemplateSpec(app):
    imagePullSecrets = list()

    if 'imagePullSecret' in app.image:
        imagePullSecrets.append({'name': app.image['imagePullSecret']})

    volumes = list()
    for volume in app.volumes:
        volumes.append({
            'name': generatePVCName(app.name, volume['path']),
            'persistentVolumeClaim': {
                'claimName': generatePVCName(app.name, volume['path'])
            }
        })

    return {
        'containers': [generateContainer(app)],
        'imagePullSecrets': imagePullSecrets,
        'volumes': volumes
    }


def generateTemplate(app):
    labels = {
        'app': app.name
    }

    return {
        'metadata': {'labels': labels},
        'spec': generateTemplateSpec(app)
    }


def generateDeploymentSpec(app):
    return {
        'replicas': app.replicas,
        'template': generateTemplate(app),
        'selector': {
            'matchLabels': {
                'app': app.name
            }
        }
    }


class Deployment(Resource):
    def __init__(self, app):
        self.apiVersion = 'apps/v1'
        self.kind = 'Deployment'

        self.metadata = {
            'name': app.name
        }

        self.spec = generateDeploymentSpec(app)
