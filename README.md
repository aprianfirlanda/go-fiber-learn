## how is this project installed 

init project
```shell
go mod init go-fiber-learn
```

install testify for unit test
```shell
go get github.com/stretchr/testify
```

install framework go fiber v2
```shell
go get github.com/gofiber/fiber/v2
```

install google wire for dependency injection
```shell
go get github.com/google/wire
```

install logrus for logger
```shell
go get github.com/sirupsen/logrus
```

install golang validation
```shell
go get github.com/go-playground/validator/v10
```

install golang viper for manage configuration env file
```shell
go get github.com/spf13/viper
```

install gorm postgres
```shell
go get gorm.io/gorm
go get gorm.io/driver/postgres
```

install raw postgres
```shell
go get github.com/lib/pq
```

## requirement

### Dependency Injection Generator
```shell
go install github.com/google/wire/cmd/wire@latest
```
add to the path --> $GOPATH/bin/wire

create file 'injector.go' with this code on first line
```
// go:build wireinject
// +build wireinject
```
if you complete create injector. then run
```shell
wire
```
it will generate file wire_gen.go that you can use.

### DB Migration

```shell
go install -tags ‘postgres’ github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

how to create file migration
```shell
migrate create -ext sql -dir db/migrations nama_file_migration
```

how to run db migration (all)
````shell
migrate -database "koneksidatabase" -path folder up
````

how to rollback db migration (1 by 1)
````shell
migrate -database "koneksidatabase" -path folder down
````

if there is dirty state. fix it manually and use this command to choose version the before migration is error.
```shell
migrate -database "koneksi_database" -path folder force versi
```
