# api-server
green-again api server repository

### Prerequisites
Install the softwares below to run and develop the service:
- [go](https://golang.org/) (1.16)

### Building the project
Use `Makefile` to build the project. The command uses `docker-compose` to build the project.
```bash
make build
```

#### Running the Server
The command below run the server.
```bash
make up
```
### Updating dependencies
You can add dependencies in the code and just build to download dependencies. Do not
forget to run the command below after making changes.
```bash
make tidy
```

### Running the tests
You can run all tests with the command below.
```bash
make test
```
