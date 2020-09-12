
class Application:
    def __init__(self, raw):
        self.name = raw['name']

        self.image = {
            'uri': raw['image'],
            'version': raw['version']
        }

        if 'imagePullSecret' in raw:
            self.image['imagePullSecret'] = raw['imagePullSecret']

        self.ingress = None
        if 'url' in raw:
            self.ingress = { 'url': raw['url'] }

        self.service = None
        if 'port' in raw:
            self.service = { 'port': raw['port'] }

        self.env = list()
        if 'environment' in raw:
            for key, value in raw['environment'].items():
                self.env.append({
                    'name': key,
                    'value': value
                })
