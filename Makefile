.PHONY: all clean terraform check-env passwords
all: passwords build terraform

build:
	docker run \
		--rm \
		-v $(GOPATH):$(GOPATH) \
		-w ${PWD} \
		-e GOPATH=$(GOPATH) \
		-u $$(id -u) \
		golang:1.11 \
		go build -o ./bin/configserver ./cmd/configserver/

	docker run \
		--rm \
		-v $(GOPATH):$(GOPATH) \
		-w ${PWD} \
		-e GOPATH=$(GOPATH) \
		-u $$(id -u) \
		golang:1.11 \
		go build -o ./bin/configproxy ./cmd/configproxy/

clean:
	rm -rf ./bin/
	rm -f ${PWD}/config/terraform/authproxy/passwords

terraform:
	docker run \
	   	-i \
	   	-t \
		-v ${PWD}:${PWD} \
		-w ${PWD}/config/terraform \
	   	hashicorp/terraform:light init

	docker run \
	   	-i \
	   	-t \
		-v ~/.ssh:/root/.ssh \
		-v ${PWD}:${PWD} \
		-w ${PWD}/config/terraform \
	   	hashicorp/terraform:light apply

passwords: check-env
	docker run --rm -it xmartlabs/htpasswd serinus $(SERINUS_PW) > ${PWD}/config/terraform/authproxy/passwords

check-env:
	@if test -z "$(SERINUS_PW)"; then echo "SERINUS_PW must be set"; exit 1; fi
