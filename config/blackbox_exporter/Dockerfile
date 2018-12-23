FROM ubuntu

ADD blackbox_settings.yml blackbox_settings.yml
ADD ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

ADD https://github.com/prometheus/blackbox_exporter/releases/download/v0.12.0/blackbox_exporter-0.12.0.linux-amd64.tar.gz blackbox.tar.gz
RUN tar -zvxf blackbox.tar.gz blackbox_exporter-0.12.0.linux-amd64/blackbox_exporter
RUN mv blackbox_exporter-0.12.0.linux-amd64/blackbox_exporter blackbox_exporter

EXPOSE 9115

CMD ["/blackbox_exporter", "--config.file=blackbox_settings.yml"]
