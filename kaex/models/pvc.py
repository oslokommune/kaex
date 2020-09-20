from kaex.models.resource import Resource


def generatePVCName(appName, path):
    sanitized_path = path.replace('/', '')

    return f'{appName}-{sanitized_path}'


class PersistentVolumeClaim(Resource):
    def __init__(self, app, volume):
        self.apiVersion = 'v1'
        self.kind = 'PersistentVolumeClaim'

        self.metadata = {
            'name': generatePVCName(app.name, volume['path'])
        }

        accessModes = ['ReadWriteMany']
        resources = {
            'requests': {
                'storage': volume['size']
            }
        }

        self.spec = {
            'accessModes': accessModes,
            'resources': resources
        }
