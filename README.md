# go-spike

Building a web application in Go using Gin webframework and GORM for MySql

```
Go to Config/Database.go
Update DB config according to your MYSQL config
```

```
cd `Directory`
go build -o main
./main
```

Make sure you clone this repo inside GOPATH directory
```
go env
```
This command will give you the path of GOPATH

###Project Structure
```
Project --
 
 | ----- Config
 |          |
 |          | ---- Database.go
 | ----- Controllers
 |          |
 |          | ---- Todo.go
 | ----- Models
 |          |
 |          | ----- Model.go
 |          | ----- Todos.go
 | ----- Routes
 |          |
 |          | ---- Routes.go
 |
 | ------- main.go

 ```

 ###API

```
How Gin process a request
Request -> Route Parser -> [Optional Middleware] -> Route Handler -> [Optional Middleware] -> Response
```

 ```
GET : Get all todos
POST : Create a todo
GET : Get a todo
PUT : Update a todo
DELETE : Delete a todo
```

Post Params
```
{
	"name": "Invideo",
	"description": "Online video editing, colloboration, pop. templates",
}
```

####Dependency Management
check go.mod to see the dependencies
```
go mod vendor
```
Now whenever you do go get <any package> the dependecy
will automatically add to go.sum file (It contains the dependency tree of the overall project and its dependencies)

### Logging

Logging flags we can update according to our requirements
```
const (
    Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
    Ltime                         // the time in the local time zone: 01:23:23
    Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
    Llongfile                     // full file name and line number: /a/b/c/d.go:23
    Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
    LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
    Lmsgprefix                    // move the "prefix" from the beginning of the line to before the message
    LstdFlags     = Ldate | Ltime // initial values for the standard logger
)
```