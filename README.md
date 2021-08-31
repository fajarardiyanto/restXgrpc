#REST - Microservice with GRPC 

### INSTALATION
###
##### <I>CREATE .env</I>
```dotenv
API_SECRET=
DbHost=
DbPort=
DbUser=
DbName=
DbPassword=
BIND_ADDR_CLIENT=
BIND_ADDR_SERVER=
```
###
##### <I>INSTALL PACKAGE</I>
```go
go mod tidy
```
###
##### <I>RUN SERVER</I>
```go
go run cmd/main.go
```
###
##### <I>RUN CLIENT</I>
```go
go run client/main.go
```