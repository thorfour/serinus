provider "digitalocean" {
    token = "${var.do_token}"
}

resource "digitalocean_ssh_key" "default" {
    name = "default"
    public_key = "${file("${var.ssh_key_path}")}"
}

module "healthcheck_nyc3" {
    source = "./module"

    do_token = "${var.do_token}"
    region = "nyc3"

    # probe settings
    probe_count = "${var.probe_count}"
    probe_image = "${var.probe_image}"
    probe_size = "${var.probe_size}"

    # prometheus settings
    prom_image = "${var.prom_image}"
    prom_size = "${var.prom_size}"

    # ssh settings
    ssh_private_key_path = "${var.ssh_private_key_path}"
    ssh_fingerprint = "${digitalocean_ssh_key.default.fingerprint}"
}

module "healthcheck_lon1" {
    source = "./module"

    do_token = "${var.do_token}"
    region = "lon1"

    # probe settings
    probe_count = "${var.probe_count}"
    probe_image = "${var.probe_image}"
    probe_size = "${var.probe_size}"

    # prometheus settings
    prom_image = "${var.prom_image}"
    prom_size = "${var.prom_size}"

    # ssh settings
    ssh_private_key_path = "${var.ssh_private_key_path}"
    ssh_fingerprint = "${digitalocean_ssh_key.default.fingerprint}"
}

data "template_file" "prom_config" {
    template = "${file("../prometheus/master_prom.yml")}"

    vars {
        prom_addr1 = "${module.healthcheck_nyc3.prometheus_addr}"
        region1 = "nyc3"
        prom_addr2 = "${module.healthcheck_lon1.prometheus_addr}"
        region2 = "lon1"
    }
}

# Main prometheus to form federation
# Create prometheus sink
# Add configproxy to config other prometheus
resource "digitalocean_droplet" "prometheus_main" {
    image = "${var.prom_image}"
    name = "prometheus-main"
    region = "nyc3"
    size = "${var.prom_size}"
    monitoring = true
    ssh_keys = ["${digitalocean_ssh_key.default.fingerprint}"]
    tags = ["healthcheck", "terraform", "main_prometheus"]

    provisioner "file" {
        content = "${data.template_file.prom_config.rendered}"
        destination = "/etc/prom.yml"

        connection {
            type = "ssh"
            user = "root"
            private_key = "${file("${var.ssh_private_key_path}")}"
        }
    }

    provisioner "file" {
        source = "../../bin/configproxy"
        destination = "/configproxy"

        connection {
            type = "ssh"
            user = "root"
            private_key = "${file("${var.ssh_private_key_path}")}"
        }
    }

    provisioner "remote-exec" {
        script = "docker.sh"

        connection {
            type = "ssh"
            user = "root"
            private_key = "${file("${var.ssh_private_key_path}")}"
        }
    }

    provisioner "remote-exec" {
        inline = [
            "docker pull prom/prometheus",
            "docker run -d -p 9090:9090 -v /etc/prom.yml:/etc/prom.yml prom/prometheus --config.file=/etc/prom.yml",
            "chmod +x /configproxy",
            "/configproxy -p 9091 -r 'nyc3,${module.healthcheck_nyc3.prometheus_addr}:9091,lon1,${module.healthcheck_lon1.prometheus_addr}:9091' &",
        ]

        connection {
            type = "ssh"
            user = "root"
            private_key = "${file("${var.ssh_private_key_path}")}"
        }
    }
}

module "authproxy" {
    source = "./authproxy"

    do_token = "${var.do_token}"
    region = "nyc3"
    authtag = "authproxy"
    endpoint = "${digitalocean_droplet.prometheus_main.ipv4_address}"

    # ssh settings
    ssh_private_key_path = "${var.ssh_private_key_path}"
    ssh_fingerprint = "${digitalocean_ssh_key.default.fingerprint}"
}

# Firewall all resources
resource "digitalocean_firewall" "healthcheck" {
    name = "healthcheck-auth-proxy-only"

    tags = ["healthcheck"]

    inbound_rule = [
    {
        protocol = "tcp"
        port_range = "22"
        source_addresses = ["0.0.0.0/0", "::/0"]
    },
    {
        protocol = "tcp"
        port_range = "9090"
        source_tags = ["healthcheck","authproxy"]
    },
    {
        protocol = "tcp"
        port_range = "9091"
        source_tags = ["healthcheck","authproxy"]
    },
    {
        protocol = "tcp"
        port_range = "9115"
        source_tags = ["healthcheck","authproxy"]
        source_load_balancer_uids = [
            "${module.healthcheck_nyc3.loadbalancer_id}",
            "${module.healthcheck_lon1.loadbalancer_id}",
        ]
    },
    ]

    outbound_rule = [
    {
        protocol = "tcp"
        port_range = "1-65535"
        destination_addresses = ["0.0.0.0/0", "::/0"]
    },
    {
        protocol = "udp"
        port_range = "1-65535"
        destination_addresses = ["0.0.0.0/0", "::/0"]
    },
    {
        protocol = "icmp"
        destination_addresses = ["0.0.0.0/0", "::/0"]
    },
    ]
}

output "proxy_ip" {
    value = "${module.authproxy.authproxy_ip}"
}
