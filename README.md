[![GitHub release](https://img.shields.io/github/release/sgaunet/docker-auth.svg)](https://github.com/sgaunet/docker-auth/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/sgaunet/docker-auth)](https://goreportcard.com/report/github.com/sgaunet/docker-auth)
![GitHub Downloads](https://img.shields.io/github/downloads/sgaunet/docker-auth/total)
[![License](https://img.shields.io/github/license/sgaunet/docker-auth.svg)](LICENSE)

# docker-auth

Tool to manage the authentifications in the docker configuration $HOME/.docker/config.json

## Install

Download binary from the releases page.

A docker image is available. Use it to install docker-auth in your docker image. Example:

```Dockerfile
FROM sgaunet/docker-auth:v0.1.0 AS auth-image

FROM alpine:3.20.3
COPY --from=auth-image /usr/bin/docker-auth /usr/bin/docker-auth
...
````

## Run

```bash
docker-auth add -l login -p password -r registry.example.com
# authentication will be added to ~/.docker/config.json
# use the argument -c if you want to initialize a different configuration file
```

## Project Status

🟨 **Maintenance Mode**: This project is in maintenance mode.

While we are committed to keeping the project's dependencies up-to-date and secure, please note the following:

- New features are unlikely to be added
- Bug fixes will be addressed, but not necessarily promptly
- Security updates will be prioritized

## Issues and Bug Reports

We still encourage you to use our issue tracker for:

- 🐛 Reporting critical bugs
- 🔒 Reporting security vulnerabilities
- 🔍 Asking questions about the project

Please check existing issues before creating a new one to avoid duplicates.

## Contributions

🤝 Limited contributions are still welcome.

While we're not actively developing new features, we appreciate contributions that:

- Fix bugs
- Update dependencies
- Improve documentation
- Enhance performance or security

~~If you're interested in contributing, please read our [CONTRIBUTING.md](link-to-contributing-file) guide for more information on how to get started.~~

## Support

As this project is in maintenance mode, support may be limited. We appreciate your understanding and patience.

Thank you for your interest in our project!
