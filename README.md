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

By default serinus deplotes to multiple regions (nyc3,lon1)

This will build the required binaries, and deploy federated [prometheus](https://prometheus.io/) that point to [blackbox exporters](https://github.com/prometheus/blackbox_exporter)
