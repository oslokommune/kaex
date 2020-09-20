from ruamel.yaml import YAML

from kaex.models.application import Application

from kaex.models.resource import Resource
from kaex.models.deployment import Deployment
from kaex.models.service import Service
from kaex.models.ingress import Ingress
from kaex.models.pvc import PersistentVolumeClaim

yaml = YAML(typ='safe')

yaml.register_class(Deployment)
yaml.register_class(Service)
yaml.register_class(Ingress)
yaml.register_class(PersistentVolumeClaim)
