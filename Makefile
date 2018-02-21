.PHONY: all run-ci clean lint deps generate

all: clean lint generate

deps:
		@echo "Install dependencies..."
		dep ensure

clean:
		@echo "Cleanup..."
		find . -maxdepth 2 -type f -name "*_generated.go" -delete

lint:
		@echo "Run linters..."
		go fmt $$(go list ./... | grep -v /vendor/)
		go vet $$(go list ./... | grep -v /vendor/)
		golint $$(go list ./... | grep -v /vendor/) | grep -v _generated.go; test $$? -eq 1

generate:
		$$GOPATH/src/github.com/artemnikitin/flatdata/generator/app.py -v -g go \
            -s coappearances.flatdata \
            -O coappearances/coappearances_generated.go
