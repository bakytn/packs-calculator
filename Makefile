build:
	./docker/build.sh

.PHONY: vendor
vendor:
	docker run --rm -e GO111MODULE=on -v "$(PWD)":/go/src/repartners repartners/golang-dev sh -c 'go mod tidy && go mod vendor'

up:
	@make build
	@make vendor
	docker-compose -f docker-compose.yml up -d
logs:
	docker-compose logs -f
rm:
	docker-compose rm  -sfv
start:
	docker-compose start
stop:
	docker-compose stop
bash:
	docker-compose exec packs-calculator bash
tests:
	docker-compose exec packs-calculator go test ./...
