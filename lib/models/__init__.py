from ruamel.yaml import YAML

from .application import Application
from .deployment import Deployment
from .service import Service

yaml = YAML(typ='safe')

yaml.register_class(Deployment)
yaml.register_class(Service)
