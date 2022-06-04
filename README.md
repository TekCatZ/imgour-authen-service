[![Contributors][contributors-shield]][contributors-url]
[![AGPLv3 License][license-shield]][license-url]

# Imgour Authenticate Service

## About the project

This is the service handling all authentication/authorization process of Imgour.

### Status

The project is in developing status.

## Getting started

### Layout

```tree
├── .gitignore
├── CHANGELOG.md
├── README.md
├── LICENSE
├── go.mod
├── go.sum
├── .dockerignore
├── Dockerfile
├── api
├── cmd
│   └── imgour-authen
│       └── main.go
├── configs
├── pkg
├── intertal
│   └── imgour-authen
│       ├── controllers
│       │   ├── models
│       │   │   ├── dto
│       │   │   └── param
│       │   └── route
│       ├── entities
│       │   ├── models
│       │   ├── repositories
│       │   └── services
│       │       └── db
│       ├── usecases
│       │   ├── handlers
│       │   └── middleware
│       ├── app.go
│       └── configs.go
├── deployments
└── test
```

A brief description of the layout:

* `.gitignore` includes files that should not be committed to git
* `README.md` is a detailed description of the project, which is this file.
* `cmd` contains main packages, each subdirectory of `cmd` is a main package.
* `pkg` contains sharable code between projects, use common senses.
* `api` holds api related file like swagger, proto, etc.
* `configs` contains config, config files' name convention: **\<env\>.config.yaml**, by default, env is dev, env is set when running the service
* `internal` holds internal logics of each service. These logics should not be shared between service if any, same with `pkg`, use common senses. The "root" of services' logics should match the structure of `cmd`
* `internal/<service>/controllers` contains logics and models related to the controller part, in this project, it's routing, param validation logic
* `internal/<service>/usecases` contains logics and models related to the activity of the service which should not be related to the core logic, include logging, mapping, etc. However, if putting or not putting any logic here makes senses, feel free to do it.
* `internal/<service>/entities` contains logics and models related to the core activity, business-related logics of the service. External services call, db-related activities, calculating, aggregating etc. should be put here. However, same with `usecases`, if putting or not putting any logic here makes senses, feel free to do it.
* `test` holds all tests (except unit tests), e.g. integration, e2e tests.

### Config file structure

```yaml
db:
  host: <host url with options if any>
  username: <username>
  password: <password>
```

[contributors-shield]: https://img.shields.io/github/contributors/TekCatZ/imgour-authen-service.svg?style=for-the-badge
[contributors-url]: https://github.com/TekCatZ/imgour-authen-service/graphs/contributors
[license-shield]: https://img.shields.io/github/license/TekCatZ/imgour-authen-service.svg?style=for-the-badge
[license-url]: https://github.com/TekCatZ/imgour-authen-service/blob/master/LICENSE
