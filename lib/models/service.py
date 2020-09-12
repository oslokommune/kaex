from .resource import Resource

class Service(Resource):
    def __init__(self, app):
        self.apiVersion = 'v1'
        self.kind = 'Service'

        self.metadata = {
            'name': app.name
        }

        ports = [
            { 'port': app.service['port'], 'targetPort': app.service['targetPort'] }
        ]
        selector = { 'app': app.name }
        service_type = 'ClusterIP'

        self.spec = {
            'ports': ports,
            'selector': selector,
            'type': service_type
        }
