variable "do_token" {}

variable "probe_count" {}

variable "ssh_key_path" {
    default = "~/.ssh/id_rsa.pub"
}

variable "ssh_private_key_path" {
    default = "~/.ssh/id_rsa"
}

variable "probe_image" {
    default = "ubuntu-18-04-x64"
}

variable "probe_size" {
    default = "s-1vcpu-1gb"
}

variable "prom_image" {
    default = "ubuntu-18-04-x64"
}

variable "prom_size" {
    default = "s-1vcpu-1gb"
}