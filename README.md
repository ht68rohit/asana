# Asana as a microservice
An OMG service for asana, it designed to help teams organize, track, and manage their work. 

[![Open Microservice Guide](https://img.shields.io/badge/OMG-enabled-brightgreen.svg?style=for-the-badge)](https://microservice.guide)


## [OMG](hhttps://microservice.guide) CLI

### OMG

* omg validate
```
omg validate
```
* omg build
```
omg build
```
### Test Service

* Test the service by following OMG commands

### CLI

##### Create Project
```sh
$ omg run create_project -a name=<PROJECT_NAME> -a notes=<NOTES> -a color=<COLOR> -a workspace=<WORKSPACE_ID> -a public=<PUBLIC_TO_ORGANIZATION> -e ACCESS_TOKEN=<ACCESS_TOKEN>
```
##### Create Task
```sh
$ omg run create_task -a name=<TASK_NAME> -a notes=<NOTES> -a project=<PROJECT_ID> -a assignee=<ASSIGNEE_EMAIL_ADDRESS> -a workspace=<WORKSPACE_ID> -a followers=<LIST_OF_EMAIL_ADDRESS> -a hearted=<BOOLEAN> -e ACCESS_TOKEN=<ACCESS_TOKEN>
```
##### Delete Project
```sh
$ omg run delete_project -a project_id=<PROJECT_ID> -e ACCESS_TOKEN=<ACCESS_TOKEN>
```

## License
### [MIT](https://choosealicense.com/licenses/mit/)

## Docker
### Build
```
docker build -t microservice-asana .
```
### RUN
```
docker run -p 3000:3000 microservice-asana
```
