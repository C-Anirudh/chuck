# chuck

[![Gitter](https://badges.gitter.im/chuck-project/community.svg)](https://gitter.im/chuck-project/community?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge) [![Maintenance](https://img.shields.io/badge/Maintained%3F-yes-green.svg)](https://github.com/C-Anirudh/chuck/graphs/commit-activity) [![GitHub license](https://img.shields.io/github/license/C-Anirudh/chuck.svg)](https://github.com/C-Anirudh/chuck/blob/master/LICENSE) [![GitHub contributors](https://img.shields.io/github/contributors/C-Anirudh/chuck.svg)](https://GitHub.com/C-Anirudh/chuck/graphs/contributors/) [![HitCount](http://hits.dwyl.io/C-Anirudh/chuck.svg)](http://hits.dwyl.io/C-Anirudh/chuck)

[![GitHub stars](https://img.shields.io/github/stars/C-Anirudh/chuck.svg?style=social&label=Star&maxAge=2592000)](https://GitHub.com/C-Anirudh/chuck/stargazers/) [![GitHub forks](https://img.shields.io/github/forks/C-Anirudh/chuck.svg?style=social&label=Fork&maxAge=2592000)](https://GitHub.com/C-Anirudh/chuck/network/) [![GitHub watchers](https://img.shields.io/github/watchers/C-Anirudh/chuck.svg?style=social&label=Watch&maxAge=2592000)](https://GitHub.com/C-Anirudh/chuck/watchers/)

To Do and personal planner.

## :minidisc: Installation instructions

You must have [Go](https://golang.org/) and [npm](https://www.npmjs.com/) installed in your computer. Then follow these steps:

```
go get github.com/C-Anirudh/chuck
```

Install all Go dependencies by running

```
go get github.com/gorilla/mux
```

Install `http-server` by running the command

```
npm install -g http-server
```

Run your http-server by executing the command

```
npm start
```

and your Go server by running the following command in your `service` folder

```
go run main.go
```

> Navigate to **http://localhost:8080/#!/** in your browser to see the application.

> **Tip**: Changes to the html, css and js files may not reflect immediately in the browser because of caching. It is advisable to install a browser extension that clears the cache for you.

## Setting up the Database

This project uses PostgreSQL as its backend database. You can install it by following the instructions provided in their official [site](https://www.postgresql.org/download/).

Once installed you can connect to PostgreSQL from the command line using the command

```
psql -U postgres
```

Command to create a database

```sql
CREATE DATABASE chuck;
```

You can connect to the database by typing

```
\c chuck
```

You can see the list of relations by typing

```
\d
```

Add the name of the database and your db password to the const variables provided at the start of `main.go`

## :wrench: Technology Stack

* **Backend** Go
* **Front-end** AngularJS

## Contributing

Refer [CONTRIBUTING.md](https://github.com/C-Anirudh/chuck/blob/master/CONTRIBUTING.md) for instructions on how to contribute to the project.