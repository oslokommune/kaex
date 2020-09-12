from .resource import Resource

class Ingress(Resource):
    def __init__(self, app):
        self.apiVersion = 'extensions/v1beta1'
        self.kind = 'Ingress'

        self.metadata = {
            'name': app.name
        }

        self.rules = [
            {
                'host': app.ingress['url'],
                'http': {
                    'backend': {
                        'service': {
                            'name': app.name,
                            'port': app.service['targetPort']
                        }
                    }
                }
            }
        ]

        if app.ingress['tls']:
            self.tls = [
                { hosts: [ app.ingress['url'] ] }
            ]
