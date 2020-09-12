from kaex.models.resource import Resource

class Ingress(Resource):
    def __init__(self, app):
        self.apiVersion = 'extensions/v1beta1'
        self.kind = 'Ingress'

        self.metadata = {
            'name': app.name
        }

        rules = [
            {
                'host': app.ingress['url'],
                'http': {
                    'paths': [
                        {
                            'backend': {
                                'serviceName': app.name,
                                'servicePort': app.service['port']
                            }
                        }
                    ]
                }
            }
        ]

        self.spec = { 'rules': rules }

        if app.ingress['tls']:
            self.spec['tls'] = [
                { 'hosts': [ app.ingress['url'] ] }
            ]
