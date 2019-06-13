# _Asana_ OMG Microservice

[![Open Microservice Guide](https://img.shields.io/badge/OMG%20Enabled-👍-green.svg?)](https://microservice.guide)
[![Build Status](https://travis-ci.com/heaptracetechnology/microservice-asana.svg?branch=master)](https://travis-ci.com/heaptracetechnology/microservice-asana)
[![codecov](https://codecov.io/gh/heaptracetechnology/microservice-asana/branch/master/graph/badge.svg)](https://codecov.io/gh/heaptracetechnology/microservice-asana)


An OMG service for asana, it designed to help teams organize, track, and manage their work.

## Usage in [Storyscript](https://storyscript.io/)

##### Create Project
```coffee
>>> asana createProject name:'name' notes:'notes' color:'color' workspace:'workspace' 

```

Curious to [learn more](https://docs.storyscript.io/)?

✨🍰✨

## Usage with [OMG CLI](https://www.npmjs.com/package/omg)
##### Create Project
```shell
$ omg run createProject -a name=<PROJECT_NAME> -a notes=<NOTES> -a color=<COLOR> -a workspace=<WORKSPACE_ID> -a public=<PUBLIC_TO_ORGANIZATION> -e ACCESS_TOKEN=<ACCESS_TOKEN>
```
##### Create Task
```shell
$ omg run createTask -a name=<TASK_NAME> -a notes=<NOTES> -a projectId=<PROJECT_ID> -a assignee=<ASSIGNEE_EMAIL_ADDRESS> -a workspace=<WORKSPACE_ID> -a followers=<LIST_OF_EMAIL_ADDRESS> -a hearted=<BOOLEAN> -e ACCESS_TOKEN=<ACCESS_TOKEN>
```
##### Delete Project
```shell
$ omg run deleteProject -a projectId=<PROJECT_ID> -e ACCESS_TOKEN=<ACCESS_TOKEN>
```
##### Delete Task
```shell
$ omg run deleteTask -a taskId=<TASK_ID> -e ACCESS_TOKEN=<ACCESS_TOKEN>
```
##### List Task
```shell
$ omg run listTask -a workspace=<WORKSPACE_ID> -e ACCESS_TOKEN=<ACCESS_TOKEN>
```
##### List Workspace
```shell
$ omg run listWorkspace -e ACCESS_TOKEN=<ACCESS_TOKEN>
```
##### Find Task
```shell
$ omg run findTask -a taskId=<TASK_ID> -e ACCESS_TOKEN=<ACCESS_TOKEN>
```
##### Find Project
```shell
$ omg run findProject -a projectId=<PROJECT_ID> -e ACCESS_TOKEN=<ACCESS_TOKEN>
```
##### Update Project
```shell
$ omg run updateProject -a id=<PROJECT_ID> -a name=<PROJECT_NAME> -a notes=<NOTES> -a color=<COLOR> -a public=<PUBLIC_TO_ORGANIZATION> -e ACCESS_TOKEN=<ACCESS_TOKEN>
```
##### List Tasks Form Project
```shell
$ omg run listProjectTasks -a projectId=<PROJECT_ID> -e ACCESS_TOKEN=<ACCESS_TOKEN>
```
##### SubscribeTasks
```shell
$ omg subscribe receive task -a workspaceId=<WORKSPACE_ID> -a projectId=<PROJECT_ID> -a existing=<TRUE/FALSE> -e ACCESS_TOKEN=<ACCESS_TOKEN>
```

**Note**: the OMG CLI requires [Docker](https://docs.docker.com/install/) to be installed.

## License
[MIT License](https://choosealicense.com/licenses/mit/).
```
