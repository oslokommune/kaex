from ruamel.yaml import YAML
from io import StringIO

yaml = YAML(typ='safe')

CLASSES = ['Application', 'Deployment', 'Service', 'Ingress']

def cleanClassNames(string):
    for item in CLASSES:
        string = string.replace(f'!{item}', '')

    return string

class Resource():
    def toYAML(self):
        stream = StringIO()

        yaml.dump(self, stream)

        result = stream.getvalue()
        stream.close

        result = cleanClassNames(result)

        return result
