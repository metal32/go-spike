# go-spike

Building a web application in Go using Gin webframework and GORM

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
GET : Get all todos
POST : Create a todo
GET : Get a todo
PUT : Update a todo
DELETE : Delete a todo
```

####Dependency Management
check go.mod to see the dependencies
```
go mod vendor
```
Now whenever you do go get <any package> the dependecy
will automatically add to go.sum file (It contains the dependency tree of the overall project and its dependencies)