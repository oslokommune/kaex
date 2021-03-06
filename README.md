# Kaex

## Pronouncation
Kayx, like "cakes"

## Installation
For now;
1. `git clone git@github.com:oslokommune/kaex.git`
2. `cd kaex`
4. `make install`

### P.S.
This will install kaex into ~/.local/bin. Ensure that folder is in the path
If ~/.local/bin is not in the path, run 

`export PATH=${PATH}:~/.local/bin`

to add it.

To use a different folder for installation, configure the installation directory with `INSTALL_DIR`

`make install INSTALL_DIR=preferred/dir`

## Usage
1. `kaex init > application.yaml`
2. Edit application.yaml to suit your app
3. `cat application.yaml | kaex expand | kubectl apply -f`

## Erm. What was that three-step instructions above for?
Sorry! I assumed you.. Never mind. I'll explain.

Based on an application.yaml, Kaex produces relevant yaml to spawn Kubernetes
resources.

An application.yaml looks somewhat like this:
```yaml
name: my-app
image: docker.pkg.github.com/my-org/my-repo/my-package
version: 0.0.1

url: https://my-domain.io
port: 3000

environment:
  MY_VARIABLE: my value
```

## My eyes!!! Why would you do such a thing?
Tired of populating overpopulated resource templates. I wanted something with
some nifty defaults.

## Sheesh ok. Why not a CRD?
Seemed like more work.

## Right. And what about Helm?
I don't really know. I prefer starting with a minimalistic starting point and
adding stuff as I go.

## So a minimalistic Helm starter?
I don't really need everything Helm offers at all times. Like every other
abstraction layer it adds complexity.


## Roadmap

- [x] --save to file(s)
- [x] Make installable
- [x] Support for volumes
- [x] Support for annotations
- [ ] Support for secrets
