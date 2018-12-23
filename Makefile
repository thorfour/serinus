all:
	docker run \
		--rm \
		-v $(GOPATH):$(GOPATH) \
		-w ${PWD} \
		-e GOPATH=$(GOPATH) \
		golang:1.11 \
		go build -o ./bin/configserver ./cmd/configserver/

	docker run \
		--rm \
		-v $(GOPATH):$(GOPATH) \
		-w ${PWD} \
		-e GOPATH=$(GOPATH) \
		golang:1.11 \
		go build -o ./bin/configproxy ./cmd/configproxy/

clean:
	rm -R ./bin/
	rm ${PWD}/config/terraform/authproxy/passwords

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

passwords:
	docker run --rm -it xmartlabs/htpasswd serinus $(SERINUS_PW) > ${PWD}/config/terraform/authproxy/passwords
