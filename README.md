# Blog

A project to better my golang abilities.
This project is a simple web application with these features:

- smart routing
- html template
- personal modules
- repository layer
- service layer
- model layer
- connect to db
- modern cli
- flexible configs
- session management
- etc

---

## Prerequisite packages

The following packages are used to develop the project:

- go v1.21
- gin v1.9.1
- uuid v1.6.0
- sessions v1.2.1
- cobra v1.8.0
- viper v1.18.2
- gorm v1.25.7

> database is Mysql

## Run

To run the project just type:

```
go run main.go migrate
go run main.go seed
go run main.go serve
```

> with migrate you can create database tables.
> with seed will generate sample data on DB.

Also you can see other command with:

```
go run main.go help
```

> this project by default run on http://localhost:8081, you can change this on config.yml file

## Run Mysql docker container

If you are not running a Mysql on your machine, you can use docker compose file on /deployment:

```
cd deployment
docker-compose up -d
```

## Todo List

- [x] Article background image
- [] User create with picture
- [] Delete article
- [] Edit article
- [] Display user info page
- [] Edit user info
- [] create Dockerfile
- [] Deploy on K8S
