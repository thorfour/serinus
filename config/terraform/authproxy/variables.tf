variable "do_token" {}

variable "auth_image" {
    default = "ubuntu-18-04-x64"
}

variable "region" {
    default = "nyc3"
}

variable "auth_size" {
    default = "s-1vcpu-1gb"
}

variable "ssh_fingerprint" {}

variable "authtag" {
    default = "authproxy"
}

variable "ssh_private_key_path" {
    default = "~/.ssh/id_rsa"
}

variable "endpoint" {}
