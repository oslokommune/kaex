from ruamel.yaml import YAML

from .application import Application

from .deployment import Deployment
from .service import Service
from .ingress import Ingress

yaml = YAML(typ='safe')

yaml.register_class(Application)
yaml.register_class(Deployment)
yaml.register_class(Service)
yaml.register_class(Ingress)
