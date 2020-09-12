from kaex.models.resource import Resource

def generateContainer(app):
    container = dict()

    volumeMounts = list()

    return {
        'name': 'app',
        'image': f'{app.image["uri"]}:{app.image["version"]}',
        'env': app.env,
        'volumeMounts': volumeMounts
    }

def generateTemplateSpec(app):
    imagePullSecrets = list()

    if 'imagePullSecret' in app.image:
        imagePullSecrets.append({ 'name': app.image['imagePullSecret'] })

    return {
        'containers': [generateContainer(app)],
        'imagePullSecrets': imagePullSecrets
    }

def generateTemplate(app):
    labels = {
        'app': app.name
    }

    return {
        'metadata': { 'labels': labels },
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
