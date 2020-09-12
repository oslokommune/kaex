import sys
from ruamel.yaml import YAML
from io import StringIO

yaml = YAML(typ='safe')

class Resource():
    def toYAML(self):
        stream = StringIO()

        yaml.dump(self, stream)

        result = stream.getvalue()
        stream.close

        return result
