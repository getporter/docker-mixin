# A Docker Mixin for Porter

This is a Docker mixin for Porter. The mixin provides the Docker CLI. 

## Mixin Declaration

To use this mixin in a bundle, declare it like so:

```yaml
mixins:
- docker
```

## Mixin Commands
The commands available are docker pull, push, build, run, remove, and login. 

## Mixin Syntax
The same syntax applies for install, upgrade, and uninstall.

### Docker pull
```yaml
- docker:
    description: 
    pull:
      name:
      tag:
```

### Docker push
```yaml
- docker:
    description: 
    push:
      name:
      tag:
```

### Docker build
```yaml
- docker:
    description: 
    build:
      tag: 
      file: OPTIONAL
      path: #defaults to "." OPTIONAL
```

### Docker run
```yaml
- docker:
    description:
    run:
      image:
      name:
      detach: #defaults to false
      ports:
        - host:
          container:
      env:
        variable: 
      privileged: #defaults to false
      rm: #defaults to false
```

### Docker remove
```yaml
- docker:
    description: 
    remove:
      container: 
      force: #defaults to false
```

### Docker login
```yaml
- docker:
    description:
    login: 
      username: OPTIONAL
      password: OPTIONAL
```

## Examples

### Install
```yaml
install:
- docker:
    description: "Install Whalesay"
    pull:
      name: docker/whalesay
      tag: latest
- docker:
    description: "Build image"
    build:
      tag: "gmadhok/cookies:v1.0"
      file: Dockerfile-cookies
- docker:
    description: "Run Whalesay"
    run:
      name: mixinpractice
      image: "docker/whalesay:latest"
      command: cowsay
      arguments:
        - "Hello World"
```

### Upgrade
```yaml
upgrade:
- docker:
    description: "Upgrade image"
    pull:
      name: docker/whalesay
      tag: latest
```

### Uninstall
```yaml
uninstall:
- docker:
    description: "Login to docker"
    login:
- docker:
    description: "Push image"
    push:
      name: gmadhok/cookies
      tag: v1.0
- docker:
    description: "Remove mixinpractice"
    remove:
      container: mixinpractice
```

## Invocation

Use of this mixin requires opting-in to Docker host access via a Porter setting.  See the Porter [documentation](https://porter.sh/configuration/#allow-docker-host-access) for further details.

Here we opt-in via the CLI flag, `--allow-docker-host-access`:
```shell
$ porter install --allow-docker-host-access
```