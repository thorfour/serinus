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
	cd ./config/terraform
	terraform init
	terraform apply

passwords:
	docker run --rm -it xmartlabs/htpasswd serinus $(SERINUS_PW) > ${PWD}/config/terraform/authproxy/passwords
