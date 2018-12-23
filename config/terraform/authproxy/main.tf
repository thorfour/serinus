provider "digitalocean" {
    token = "${var.do_token}"
}

data "template_file" "authconf" {
    template = "${file("${path.module}/authproxy.conf")}"

    vars {
        server = "${var.endpoint}"
    }
}

resource "digitalocean_droplet" "authproxy" {
    image = "${var.auth_image}"
    name = "authproxy"
    region = "${var.region}"
    size =  "${var.auth_size}"
    monitoring = true
    ssh_keys = ["${var.ssh_fingerprint}"]
    tags = ["terraform", "${var.authtag}"]

    provisioner "file" {
        content = "${data.template_file.authconf.rendered}"
        destination = "authproxy.conf"

        connection {
            type = "ssh"
            user = "root"
            private_key = "${file("${var.ssh_private_key_path}")}"
        }
    }

    provisioner "file" {
        source = "${path.module}/Dockerfile"
        destination = "Dockerfile"

        connection {
            type = "ssh"
            user = "root"
            private_key = "${file("${var.ssh_private_key_path}")}"
        }
    }

    provisioner "file" {
        source = "${path.module}/passwords"
        destination = "passwords"

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
            "docker build -t authproxy .",
            "docker run -d -p 9090:9090 -p 9091:9091 authproxy",
        ]

        connection {
            type = "ssh"
            user = "root"
            private_key = "${file("${var.ssh_private_key_path}")}"
        }
    }

}

output "authproxy_ip" {
    value = "${digitalocean_droplet.authproxy.ipv4_address}"
}
