import sys
from ruamel.yaml import YAML

from .resource import Resource

class Service(Resource):
    def __init__(self, app):
        self.apiVersion = 'v1'
        self.kind = 'Service'

        self.metadata = {
            'name': app['name']
        }

        self.ports = [ { 'port': 80, 'targetPort': app.port } ]
        self.selector = { 'app': app['name'] }
        self.type = 'ClusterIP'
