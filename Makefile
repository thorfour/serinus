all:
	go build -o ./bin/configserver ./cmd/configserver/
	go build -o ./bin/configproxy ./cmd/configproxy
clean:
	rm -R ./bin/

terraform:
	cd ./config/terraform
	terraform apply

passwords:
	docker run --rm -it xmartlabs/htpasswd serinus $(SERINUS_PW) > ${PWD}/config/terraform/authproxy/passwords
