# Packs Calculator
**how to use**

- To start, please install: docker, docker-compose and go

 - `make up` on root folder run it will start test-task docker
   container with all requied dependencies.
  - `make stop` stops all containers
  - `make rm` stops and removes containers
  - `make logs` displays logs
  - `make bash` to open container bash window
  - `make tests` run all available go tests
  - `localhost:3333` to access go container from localhost


## API

### Order of 1 Item
`curl -X POST -H "Content-Type: application/json" -d '{"items_ordered": 1}' http://localhost:3333/calculate-packs`
```json
{
    "items_ordered":1001,
     "packs_to_send":[
      {"pack_size":250,"count":1}
    ]
}
```

### Order 750 Items
`curl -X POST -H "Content-Type: application/json" -d '{"items_ordered": 750}' http://localhost:3333/calculate-packs`
```json
{
    "items_ordered":1001,
     "packs_to_send":[
      {"pack_size":250,"count":1},
      {"pack_size":500,"count":1}
    ]
}
```

### Order 1001 Items
`curl -X POST -H "Content-Type: application/json" -d '{"items_ordered": 1001}' http://localhost:3333/calculate-packs`
response:
```json
{
    "items_ordered":1001,
     "packs_to_send":[
      {"pack_size":250,"count":1},
      {"pack_size":1000,"count":1}
    ]
}
```
