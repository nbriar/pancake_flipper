CYCLO_COMPLEXITY ?= 25
PACKAGES ?= $$(go list ./... | grep -v /vendor/ | grep -v /tmp )
.PHONY: run cover cyclo tests vet build

export

run:
	go run main.go $(stack)

cover:
	bash ./scripts/coverage.sh

cyclo:
	find . -type f -name '*.go' -not -path './vendor/*' | grep -v health | xargs -I {} gocyclo -over $(CYCLO_COMPLEXITY) {}

tests:
	go test -v $(PACKAGES)

vet:
	go vet $(PACKAGES)

build:
	go build -v

