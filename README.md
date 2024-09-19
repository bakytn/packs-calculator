# Packs Calculator
**how to use**

- To start, please install: docker, docker-compose and go

 - `make up` on root folder run it will start test-task docker
   container with all requied dependencies.
  - `make bash` to open container bash window
  - `localhost:3333` to access go container from localhost


## API

### Order of 1 Item
`curl -X POST -H "Content-Type: application/json" -d '{"items_ordered": 1}' http://localhost:3333/calculate-packs`

### Custom Pack Sizes
`curl -X POST -H "Content-Type: application/json" -d '{"items_ordered": 750}' http://localhost:3333/calculate-packs`

### Invalid Pack Sizes
`curl -X POST -H "Content-Type: application/json" -d '{"items_ordered": 100}' http://localhost:3333/calculate-packs`
