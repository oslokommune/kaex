#!/usr/bin/python
import argparse
import os
import sys
import urllib.request
from ruamel.yaml import YAML

from kaex.models import Application, Deployment, Ingress, Service, PersistentVolumeClaim


def saveResources(path, data):
    resources = data.split('---')

    for resource in resources:
        yaml = YAML(typ='safe')

        data = yaml.load(resource)
        file_name = data['kind'].lower() + '.yaml'

        with open(os.path.join(path, file_name), 'w') as f:
            f.write(resource.strip())


def save(path, data):
    if data.find('---') != -1:
        saveResources(path, data)
    else:
        with open(os.path.join(path, 'application.yaml'), 'w') as f:
            f.write(data)


def initializeApplication():
    result = ''

    url = 'https://raw.githubusercontent.com/deifyed/kaex/master/examples/application-minimal.yaml'
    with urllib.request.urlopen(url) as response:
        result = response.read().decode('utf-8')

    return result


def generateYAML():
    yaml = YAML(typ='safe')

    application = Application(yaml.load(sys.stdin.read()))

    resources = list()

    if application.volumes:
        for volume in application.volumes:
            resources.append(PersistentVolumeClaim(application, volume))

    resources.append(Deployment(application))

    if application.service:
        resources.append(Service(application))
    if application.ingress:
        resources.append(Ingress(application))

    output = '\n---\n'.join([resource.toYAML() for resource in resources])

    return output


def main():
    parser = argparse.ArgumentParser(description='Creates Kubernetes resources')
    parser.add_argument('--save', '-s', **{
        'metavar': 'path',
        'type': str,
        'help': 'save output to file(s)'
    })

    subparsers = parser.add_subparsers(help='available actions. Default is expand')
    init_parser = subparsers.add_parser('initialize', **{
        'help': 'initiates a application.yaml file',
        'aliases': ['init']
    })
    init_parser.set_defaults(action=initializeApplication)

    expand_parser = subparsers.add_parser('expand', **{
        'help': 'expands an application.yaml file into kubernetes resources',
    })
    expand_parser.set_defaults(action=generateYAML)

    args = parser.parse_args()

    action = args.action if 'action' in args else generateYAML

    result = action()

    if args.save:
        save(args.save, result)

    print(result)


if __name__ == '__main__':
    main()
