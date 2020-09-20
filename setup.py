from setuptools import setup, find_packages

setup(
    name='kaex',
    version='0.0.2',
    author='Julius Pedersen',
    author_email='deifyed@ctemplar.com',
    description='Kubernetes Application.yaml EXpander',
    project_urls={
        'Bug Tracker': 'https://github.com/deifyed/kaex/issues',
        'Source Code': 'https://github.com/deifyed/kaex'
    },
    license='https://github.com/deifyed/kaex/blob/master/LICENSE',
    packages=find_packages(),
    install_requires=[
        'ruamel.yaml==0.16.12',
        'ruamel.yaml.clib==0.2.2'
    ],
    entry_points={
        'console_scripts': [
            'kaex=kaex.entry:main'
        ]
    }
)
