all: tidy vendor build run
all-gorm: tidy vendor build-gorm run-gorm

build:
	go build -o healthcheck_sgad cmd/app/main.go
run:
	./healthcheck_sgad

build-gorm:
	go build -o appgorm cmd/appgorm/main.go
run-gorm:
	./appgorm

clean:
	rm ./app
	rm ./healthcheck_sgad
	rm -rf vendor

tidy:
	go mod tidy

.PHONY: vendor
vendor:
	go mod vendor

.PHONY: coverage
# coverage:
# 	go test \
# 		-race -covermode=atomic -timeout 30s \
# 		-coverprofile=coverage/coverage.out \
# 		github.com/captain-corgi/golang-os-example/pkg/iptrans
