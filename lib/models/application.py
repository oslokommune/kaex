
class Application:
    def __init__(self, raw):
        self.name = raw['name']
        self.port = raw['port']

        self.image = {
            'uri': raw['image'],
            'version': raw['version']
        }

        if 'imagePullSecret' in raw:
            self.image['imagePullSecret'] = raw['imagePullSecret']

        self.ingress = None
        if 'url' in raw:
            parts = raw['url'].split('://', 1)

            self.ingress = {
                'tls': True if parts[0] == 'https' else False,
                'url': parts[1]
            }

        self.service = None
        if 'port' in raw:
            self.service = {
                'port': 80,
                'targetPort': raw['port']
            }

        self.env = list()
        if 'environment' in raw:
            for key, value in raw['environment'].items():
                self.env.append({
                    'name': key,
                    'value': value
                })

        self.replicas = raw.get('replicas', 1)
