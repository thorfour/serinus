# serinus
terraform modules for a federated prometheus uptime application

# Prerequisites 
- git
- [docker](https://www.docker.com/)
- [digitalocean](https://www.digitalocean.com/) account and an [access token](https://www.digitalocean.com/docs/api/create-personal-access-token/)

# Install
`git clone git@github.com:thorfour/serinus.git`

# Deploy 
`SERINUS_PW=<your password> make`

By default serinus deploys to multiple regions (nyc3,lon1)

This will build the required binaries, and deploy federated [prometheus](https://prometheus.io/) that point to [blackbox exporters](https://github.com/prometheus/blackbox_exporter)

# Usage

After serinus deploys all the resources, terraform will printout `proxy_ip` the address of the federated prometheus instance to view metrics from all regions, as well as the ip address of the configuration proxy. 

The proxy endpoint will require a login which will be username `serinus` password being the environment variable `SERINUS_PW` you set at the beginning. 

## View metrics

open up <proxy_ip>:9090 in a web browser

## Add new endpoint to a region

`curl -X PATCH <proxy_ip>:9091/api/v1/add?target=<your_url>&module=http_2xx&region=<given_region>`

## Remove endpoint from a region

`curl -X DELETE <proxy_ip>:9091/api/v1/del?target=<your_url>&module=http_2xx&region=<given_region>`

## List all targets for a region

`curl -X GET <proxy_ip>:9091/api/v1/targets?module=http_2xx&region=<given_region>`
