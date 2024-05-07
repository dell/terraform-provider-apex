TEST?=$$(go list ./... | grep -v 'vendor')
HOSTNAME=registry.terraform.io
NAMESPACE=dell
NAME=apex
BINARY=terraform-provider-${NAME}
VERSION=1.0.0-alpha
OS_ARCH=linux_amd64

default: install

clean:
	rm -rf apexclient
	rm -rf jobsclient
	rm -rf powerflexclient
	rm -rf powerscaleclient
	rm -f ${BINARY}

extract-client:
	unzip --qq -o 'goClientZip/apex-client-openapi-1.1.0.zip' -d ./apexclient/
	unzip --qq -o 'goClientZip/apex-jobs-client-openapi-1.0.0.zip' -d ./jobsclient/
	unzip --qq -o 'goClientZip/powerflex-auth-client-4.5.zip' -d ./powerflexclient/
	unzip --qq -o 'goClientZip/powerscale-auth.zip' -d ./powerscaleclient/

no-extract-build:
	go mod download
	go build -o ${BINARY}

build: extract-client
	go mod download
	go build -o ${BINARY}

release:
	GOOS=darwin GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_darwin_amd64
	GOOS=freebsd GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_freebsd_386
	GOOS=freebsd GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_freebsd_amd64
	GOOS=freebsd GOARCH=arm go build -o ./bin/${BINARY}_${VERSION}_freebsd_arm
	GOOS=linux GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_linux_386
	GOOS=linux GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_linux_amd64
	GOOS=linux GOARCH=arm go build -o ./bin/${BINARY}_${VERSION}_linux_arm
	GOOS=openbsd GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_openbsd_386
	GOOS=openbsd GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_openbsd_amd64
	GOOS=solaris GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_solaris_amd64
	GOOS=windows GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_windows_386
	GOOS=windows GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_windows_amd64


install: build
	rm -rfv ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	find examples -type d -name ".terraform" -exec rm -rfv "{}" +;
	find examples -type f -name "trace.*" -delete
	find examples -type f -name "*.tfstate" -delete
	find examples -type f -name "*.hcl" -delete
	find examples -type f -name "*.backup" -delete
	rm -rf trace.*
	
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

uninstall:
	rm -rfv ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	find examples -type d -name ".terraform" -exec rm -rfv "{}" +;
	find examples -type f -name "trace.*" -delete
	find examples -type f -name "*.tfstate" -delete
	find examples -type f -name "*.hcl" -delete
	find examples -type f -name "*.backup" -delete
	rm -rf trace.*


test: check
	go test -i $(TEST) || exit 1                                                   
	echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4                    

check:
	terraform fmt -recursive examples/
	gofmt -s -w .
	golangci-lint run --fix --timeout 5m
	go vet

gosec:
	go list ./... | grep -vE 'powerflexclient|apexclient|jobsclient|powerscaleclient' | xargs gosec -stdout -log gosec.log -out=gosecresults.csv -fmt=csv
testacc:
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m   

generate:
	go generate ./...

cover:
	rm -f coverage.*
	go test -coverprofile=coverage.out ./...
	go tool cover -html coverage.out -o coverage.html

all: test gosec testacc generate cover install
