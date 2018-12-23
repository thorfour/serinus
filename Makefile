all:
	go build -o ./bin/configserver ./cmd/configserver/
	go build -o ./bin/configproxy ./cmd/configproxy
clean:
	rm ./bin/configserver
	rm ./bin/configproxy

terraform:
	cd ./config/terraform
	terraform apply
