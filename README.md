[![Build](https://github.com/Jira-Analyzer/backend-services/actions/workflows/go.yaml/badge.svg)](https://github.com/Jira-Analyzer/backend-services/actions/workflows/go.yaml)
![Coverage](https://img.shields.io/badge/Coverage-62.2%25-yellow)
<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
<!--   *generated with [DocToc](https://github.com/thlorenz/doctoc)* -->

- [Jira-Analyzer](#jira-analyzer)
  - [About the project](#about-the-project)
    - [API docs](#api-docs)
    - [Status](#status)
    - [See also](#see-also)
  - [Getting started](#getting-started)
    - [Download project](#download-project)
    - [Docker deployment](#docker-deployment)
    - [Run unit test and update coverage bage](#run-unit-test-and-update-coverage-bage)
    - [Build project](#build-project)
    - [Update swagger documentation](#update-swagger-documentation)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->


# Jira-Analyzer

## About the project

This projects contains service that provides API to get information about Apache Jira issues and project statistics.

### API docs

Project has configured **swagger API documentation**, that can be accessed by endpoint `GET /swagger/`
<details>
  <summary>API screenshots</summary>
  <img src="https://github.com/Jira-Analyzer/backend-services/assets/113100755/1eae634d-a10c-484f-b94f-5dd2257ce00a">
  <img src="https://github.com/Jira-Analyzer/backend-services/assets/113100755/323ccfa6-83cb-4701-bd77-2e26bf75f739">
</details>

### Status

The project is in ready for deployment.

### See also

* [Jira Analyzer web app project](https://github.com/Jira-Analyzer/frontend-services)

## Getting started

### Download project

To download use:
```bash
git clone https://github.com/Jira-Analyzer/backend-services.git
```

### Docker deployment

You can use docker deployment, to prepare docker images use:
> [!IMPORTANT]
> You should have installed **Docker** and **docker-compose** on your machine


> [!CAUTION]
> This will prune all unused images in the end of build to free up space after multistage Docker image build
```bash
make build-images
```

Then you can simply run
```bash
make start-dev
```
to run docker container using docker-compose, and

```bash
make stop-dev
```
to stop docker container

### Run unit test and update coverage bage

This will **run unit tests** and **update link** for coverage bage in README
```bash
make unit-test
```
You can regenerate mocks with
```bash
make gogen
```

### Build project

You can build executable files with
```bash
make build-backend
```
or
```bash
make build-connector
```
for every node.
Executables will be in **/bin** folder

### Update swagger documentation

> [!IMPORTANT]
> Project uses [swaggo](https://github.com/swaggo/swag), so you should install it

To update swagger documentation after adding new endpoints use:
```bash
make swag-backend
```
or
```bash
make swag-connector
```
This will generate separate documentation for **connector** and **backend** node
