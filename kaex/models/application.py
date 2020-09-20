def extractImage(raw):
    extractedImage = {
        'uri': raw['image'],
        'version': raw['version']
    }

    if 'imagePullSecret' in raw:
        extractedImage['imagePullSecret'] = raw['imagePullSecret']

    return extractedImage


def extractIngress(raw):
    url = raw.get('url', None)

    if not url:
        return None

    parts = raw['url'].split('://', 1)

    return {
        'tls': True if parts[0] == 'https' else False,
        'url': parts[1]
    }


def extractService(raw):
    port = raw.get('port', None)

    if not port:
        return None

    return {
        'port': 80,
        'targetPort': port
    }


def expandEnvironment(environment):
    expandedEnvironment = list()

    if not environment:
        return expandedEnvironment

    for key, value in environment.items():
        expandedEnvironment.append({
            'name': key,
            'value': value
        })

    return expandedEnvironment


def expandService(service):
    if not service:
        return None

    return {
        'port': 80,
        'targetPort': service['port']
    }


def expandVolumes(volumes):
    expandedVolumes = list()

    if not volumes:
        return expandedVolumes

    for volume in volumes:
        expandedVolume = dict({'size': '1Gi'})

        if type(volume) == str:
            expandedVolume['path'] = volume
        else:
            expandedVolume['path'] = next(iter(volume))
            expandedVolume['size'] = volume[expandedVolume['path']]

        expandedVolumes.append(expandedVolume)

    return expandedVolumes


class Application:
    def __init__(self, raw):
        self.name = raw['name']
        self.port = raw['port']

        self.image = extractImage(raw)
        self.ingress = extractIngress(raw)
        self.service = extractService(raw)
        self.env = expandEnvironment(raw.get('environment', None))
        self.volumes = expandVolumes(raw.get('volumes', None))

        self.replicas = raw.get('replicas', 1)
