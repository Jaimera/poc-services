## Installation

Requirements:
- Docker
- Go 1.19

Download project:

```git clone github.com/Jaimera/poc-services```

## Tests

After clone, open the project folder and run 
```go test ./...```

## Running 

Build and deploy containers with: 

```docker compose up```

While containers up, inside the project folder, open a terminal and run: 

```curl -X POST -H 'Content-Type: application/json' -d @ports.json http://localhost:8080/api/v1/ports```

## Validation

Check mysql by doing a local connection or with: 

```docker exec -i mysql mysql -uroot -proot  <<< "select * from db_poc.tb_port where slug = 'BRFNO';"```

Check by api, open a browser tab and execute: 

```http://localhost:8080/api/v1/ports/BRFNO```