https://medium.com/@amsokol.com/tutorial-how-to-develop-go-grpc-microservice-with-http-rest-endpoint-middleware-kubernetes-daebb36a97e9


Create table with following

```bash
CREATE TABLE `grpc-todo`.`ToDo` (
  `ID` INT NOT NULL AUTO_INCREMENT,
  `Title` VARCHAR(45) NULL,
  `Description` VARCHAR(100) NULL,
  `Remainder` DATETIME NULL,
  PRIMARY KEY (`ID`));

```

Run 

```bash
go mod tidy
```


Run
```bash
$ cd cmd/server
$ go build .
$ ./server -host=localhost:3306 -user=root -password=root -scheme=grpc-tod
```

From another terminal run

```bash
$ cd cmd/client
$ go build .
$ ./client -server=localhost:808
```



Tutorials:

1. https://grpc.io/docs/tutorials/basic/go/


REST:


Run
```bash
$ cd cmd/server
$ go build .
$ ./server -host=localhost:3306 -user=root -password=root -scheme=grpc-todo http-port=9090 -grpc-port=8080
```

From another terminal run

```bash
$ cd cmd/client
$ go build .
$ ./client-rest -server=http://localhost:9090
```
