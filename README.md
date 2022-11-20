## Installation

Download project:

```git clone github.com/Jaimera/poc-services```

```docker compose up``` 


## Tests

After clone, open the project file and run 
```go test ./...```

## Running 

While containers up, run a: 

```curl -X POST -H 'Content-Type: application/json' -d @ports.json http://localhost:8080/api/v1/ports```