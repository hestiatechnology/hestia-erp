# Setup your dev environment
This guide will walk you through setting up your development environment for the Hestia ERP Backend.

## Tools
### Coding Tools
- [Visual Studio Code](https://code.visualstudio.com/)
    - This is the recommended IDE for this project, however you can use any IDE you like.
    - [Go Extension](https://marketplace.visualstudio.com/items?itemName=golang.go)
    - [gRPC Tools](https://marketplace.visualstudio.com/items?itemName=zxh404.vscode-proto3)
- [Golang](https://golang.org/)
    - This is the language the backend is written in.
    - Install the latest version of Go.
- [Protocol Buffers](https://github.com/protocolbuffers/protobuf/releases/latest)
    - Download and install the latest version of Protoc.
    - Add the Protoc to your PATH.
- [Docker](https://www.docker.com/products/docker-desktop)
    - This is used to run the database and other services locally.
    - Install the latest version of Docker.

#### Setup Protoc for Golang
1. Install the Protoc plugin for Golang.
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

### Database Tools
We use PostgreSQL for the database. You will need to install the following tools to work with the database. 

- [PGModeler](https://pgmodeler.io/)
    - This is used to design and see the database schema.
    - You will need to compile PGModeler or request a copy.
- [PGAmin](https://www.pgadmin.org/)
    - This is used to manage the database.
    - You will use this to test queries.
    - Install the latest version of PGAdmin.

### Setup database
#### Using Docker
We recommend using Docker to run the database locally. You can use the following command to start a PostgreSQL instance.
```bash
docker run --name hestia-erp-db -e POSTGRES_PASSWORD=YOUR_PASSWORD_HERE -p 5432:5432 -d postgres:16.1-bullseye
```
You will need to replace `YOUR_PASSWORD_HERE` with your own password.

#### Importing test data to the local DB
Using PGAdmin, connect to the database and run the `test_data.sql` file to import the test data.
This will create the tables and import the test data. It may take a few minutes to run.

#### Using the Dev Database
You can request access to the dev database to work with test data. 
Not available yet.

### Setup the ENV variables
You will need to add the following to your environment variables.
```bash
export PGHOST=localhost
export PGUSER=YOUR_USERNAME_HERE
export PGPASSWORD=YOUR_PASSWORD_HERE
export PGDATABASE=erp
```
For **Windows** you can use the command `set` instead of `export`, however we recommend just setting the variables via the settings.  
Replace `YOUR_USERNAME_HERE` and `YOUR_PASSWORD_HERE` with your own username and password.
If you are using Docker to run the database, you can use the default `postgres` user and the password you set when you started the container.
