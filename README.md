# chuck
To Do and personal planner.

## :minidisc: Installation instructions
You must have [Go](https://golang.org/) and [npm](https://www.npmjs.com/) installed in your computer. Then follow these steps:

```
go get github.com/TechFlyte/shiken
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

## :wrench: Technology Stack
* **Backend** Go
* **Front-end** AngularJS
