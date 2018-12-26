provider "digitalocean" {
    token = "${var.do_token}"
}

# Create droplet for blackbox exporter
resource "digitalocean_droplet" "blackbox_node" {
    count = "${var.probe_count}"
    image = "${var.probe_image}"
    name = "blackbox-${var.region}-${count.index}"
    region = "${var.region}"
    size =  "${var.probe_size}"
    monitoring = true
    ssh_keys = ["${var.ssh_fingerprint}"]
    tags = ["healthcheck", "terraform", "${var.region}_probes"]

    provisioner "file" {
        source = "${path.module}/../../blackbox_exporter/blackbox_settings.yml"
        destination = "blackbox_settings.yml"

        connection {
            type = "ssh"
            user = "root"
            private_key = "${file("${var.ssh_private_key_path}")}"
        }
    }

    provisioner "file" {
        source = "${path.module}/../../blackbox_exporter/Dockerfile"
        destination = "Dockerfile"

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
            "cp /etc/ssl/certs/ca-certificates.crt .",
            "docker build -t blackbox_do .",
            "docker run -d -p 9115:9115 blackbox_do"
        ]

        connection {
            type = "ssh"
            user = "root"
            private_key = "${file("${var.ssh_private_key_path}")}"
        }
    }
}

resource "digitalocean_loadbalancer" "probe_balance" {
    name = "lb-${var.region}-health-probes"
    region = "${var.region}"

    forwarding_rule {
        entry_port = 9115
        entry_protocol = "http"

        target_port = 9115
        target_protocol = "http"
    }

    healthcheck {
        port = 9115
        protocol = "http"
        path = "/metrics"
    }

    droplet_tag = "${var.region}_probes"
}

data "template_file" "prom_config" {
    template = "${file("${path.module}/../../prometheus/prom.yml")}"
    
    vars {
        blackbox_addr = "${digitalocean_loadbalancer.probe_balance.ip}"
    }
}

# Create prometheus sink
resource "digitalocean_droplet" "prometheus_node" {
    image = "${var.prom_image}"
    name = "prometheus-${var.region}"
    region = "${var.region}"
    size = "${var.prom_size}"
    monitoring = true
    ssh_keys = ["${var.ssh_fingerprint}"]
    tags = ["healthcheck", "terraform", "${var.region}_prometheus"]

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
        source = "${path.module}/../../../bin/configserver"
        destination = "/configserver"

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
            "mkdir /configs",
            "mv /configserver /configs",
            "docker pull prom/prometheus",
            "docker run -d -p 9090:9090 -v /configs:/configs -v /etc/prom.yml:/etc/prom.yml prom/prometheus --config.file=/etc/prom.yml",
            "cd /configs",
            "chmod +x configserver",
            "nohup ./configserver &",
            "sleep 1",
        ]

        connection {
            type = "ssh"
            user = "root"
            private_key = "${file("${var.ssh_private_key_path}")}"
        }
    }
}

output "prometheus_addr" {
    value = "${digitalocean_droplet.prometheus_node.ipv4_address}"
}

output "loadbalancer_id" {
    value = "${digitalocean_loadbalancer.probe_balance.id}"
}
