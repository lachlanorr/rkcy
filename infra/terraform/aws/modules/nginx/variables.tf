variable "stack" {
  type = string
}

variable "cluster" {
  type = string
}

variable "vpc" {
  type = object({
    id = string
    cidr_block = string
  })
}

variable "subnets" {
  type = list(object({
    id = string
    cidr_block = string
  }))
}

variable "dns_zone" {
  type = object({
    name = string
    zone_id = string
  })
}

variable "bastion_ip" {
  type = string
}

variable "inbound_cidr" {
  type = string
}

variable "routes" {
  type = list(object({
    name = string
    hosts = list(string)
    port = number
    grpc = bool
  }))
}

variable "ssh_key_path" {
  type = string
  default = "~/.ssh/rkcy_id_rsa"
}

variable "nginx_count" {
  type = number
  default = 1
}

variable "public" {
  type = bool
}
