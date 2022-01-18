terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
      version = "~> 3.27"
    }
  }

  required_version = ">= 0.14.9"
}

provider "aws" {
  profile = "default"
  region = "us-east-2"
}

variable "stack" {
  type = string
}

variable "dns_zone" {
  type = any
}

variable "vpc" {
  type = any
}

variable "subnet_app" {
  type = any
}

variable "bastion_ips" {
  type = list
}

variable "nginx_hosts" {
  type = any
}

variable "elasticsearch_urls" {
  type = list
}

variable "ssh_key_path" {
  type = string
  default = "~/.ssh/rkcy_id_rsa"
}

variable "jaeger_collector_count" {
  type = number
  default = 1
}

variable "jaeger_query_count" {
  type = number
  default = 1
}

variable "jaeger_collector_port" {
  type = number
  default = 14250
}

variable "jaeger_query_port" {
  type = number
  default = 16686
}

variable "otelcol_count" {
  type = number
  default = 1
}

variable "otelcol_port" {
  type = number
  default = 4317
}

locals {
  sn_ids   = "${values(zipmap(var.subnet_app.*.cidr_block, var.subnet_app.*.id))}"
  sn_cidrs = "${values(zipmap(var.subnet_app.*.cidr_block, var.subnet_app.*.cidr_block))}"
}

data "aws_ami" "telemetry" {
  most_recent      = true
  name_regex       = "^rkcy/telem-[0-9]{8}-[0-9]{6}$"
  owners           = ["self"]
}

resource "aws_key_pair" "telemetry" {
  key_name = "rkcy-${var.stack}-telemetry"
  public_key = file("${var.ssh_key_path}.pub")
}

#-------------------------------------------------------------------------------
# jaeger_collector
#-------------------------------------------------------------------------------
locals {
  jaeger_collector_ips = [for i in range(var.jaeger_collector_count) : "${cidrhost(local.sn_cidrs[i], 14)}"]
}

resource "aws_security_group" "rkcy_jaeger_collector" {
  name        = "rkcy_${var.stack}_jaeger_collector"
  description = "Allow SSH and jaeger_collector inbound traffic"
  vpc_id      = var.vpc.id

  ingress = [
    {
      cidr_blocks      = [ var.vpc.cidr_block ]
      description      = ""
      from_port        = 22
      to_port          = 22
      ipv6_cidr_blocks = []
      prefix_list_ids  = []
      protocol         = "tcp"
      security_groups  = []
      self             = false
    },
    {
      cidr_blocks      = [ var.vpc.cidr_block ]
      description      = "jaeger-collector gRPC server"
      from_port        = var.jaeger_collector_port
      to_port          = var.jaeger_collector_port
      ipv6_cidr_blocks = []
      prefix_list_ids  = []
      protocol         = "tcp"
      security_groups  = []
      self             = false
    },
    {
      cidr_blocks      = [ var.vpc.cidr_block ]
      description      = "jaeger-collector HTTP server"
      from_port        = 14268
      to_port          = 14268
      ipv6_cidr_blocks = []
      prefix_list_ids  = []
      protocol         = "tcp"
      security_groups  = []
      self             = false
    },
    {
      cidr_blocks      = [ var.vpc.cidr_block ]
      description      = "jaeger-collector Admin server"
      from_port        = 14269
      to_port          = 14269
      ipv6_cidr_blocks = []
      prefix_list_ids  = []
      protocol         = "tcp"
      security_groups  = []
      self             = false
    },
    {
      cidr_blocks      = [ var.vpc.cidr_block ]
      description      = "node_exporter"
      from_port        = 9100
      to_port          = 9100
      ipv6_cidr_blocks = []
      prefix_list_ids  = []
      protocol         = "tcp"
      security_groups  = []
      self             = false
    },
  ]

  egress = [
    {
      cidr_blocks      = [ var.vpc.cidr_block ]
      description      = ""
      from_port        = 0
      ipv6_cidr_blocks = []
      prefix_list_ids  = []
      protocol         = "-1"
      security_groups  = []
      self             = false
      to_port          = 0
    }
  ]
}

resource "aws_network_interface" "jaeger_collector" {
  count = var.jaeger_collector_count
  subnet_id   = local.sn_ids[count.index]
  private_ips = [local.jaeger_collector_ips[count.index]]

  security_groups = [aws_security_group.rkcy_jaeger_collector.id]
}

resource "aws_placement_group" "jaeger_collector" {
  name     = "rkcy_${var.stack}_jaeger_collector_pc"
  strategy = "spread"
}

resource "aws_instance" "jaeger_collector" {
  count = var.jaeger_collector_count
  ami = data.aws_ami.telemetry.id
  instance_type = "m4.large"
  placement_group = aws_placement_group.jaeger_collector.name

  key_name = aws_key_pair.telemetry.key_name

  network_interface {
    network_interface_id = aws_network_interface.jaeger_collector[count.index].id
    device_index = 0
  }

  credit_specification {
    cpu_credits = "unlimited"
  }

  tags = {
    Name = "rkcy_${var.stack}_inst_jaeger_collector_${count.index}"
  }
}

resource "aws_route53_record" "jaeger_collector_private" {
  count = var.jaeger_collector_count
  zone_id = var.dns_zone.zone_id
  name    = "jaeger-collector-${count.index}.${var.stack}.local.${var.dns_zone.name}"
  type    = "A"
  ttl     = "300"
  records = [local.jaeger_collector_ips[count.index]]
}

resource "null_resource" "jaeger_collector_provisioner" {
  count = var.jaeger_collector_count
  depends_on = [
    aws_instance.jaeger_collector
  ]

  #---------------------------------------------------------
  # node_exporter
  #---------------------------------------------------------
  provisioner "file" {
    content = templatefile("${path.module}/../../../shared/node_exporter_install.sh", {})
    destination = "/home/ubuntu/node_exporter_install.sh"
  }
  provisioner "remote-exec" {
    inline = [
      <<EOF
sudo bash /home/ubuntu/node_exporter_install.sh
rm /home/ubuntu/node_exporter_install.sh
EOF
    ]
  }
  #---------------------------------------------------------
  # node_exporter (END)
  #---------------------------------------------------------

  provisioner "file" {
    content = templatefile("${path.module}/../../../shared/telemetry/jaeger-collector.service.tpl", {
      elasticsearch_urls = var.elasticsearch_urls
    })
    destination = "/home/ubuntu/jaeger-collector.service"
  }

  provisioner "remote-exec" {
    inline = [
      <<EOF
sudo mv /home/ubuntu/jaeger-collector.service /etc/systemd/system/jaeger-collector.service
sudo systemctl daemon-reload

%{for url in var.elasticsearch_urls}
RET=1
while [ $RET -ne 0 ]; do
  echo 'Trying elasticsearch ${url}'
  curl -s -f '${url}'
  RET=$?
  sleep 2
done
echo 'Connected elasticsearch ${url}'
%{endfor}

sudo systemctl start jaeger-collector
sudo systemctl enable jaeger-collector
EOF
    ]
  }

  connection {
    type     = "ssh"

    bastion_user        = "ubuntu"
    bastion_host        = var.bastion_ips[0]
    bastion_private_key = file(var.ssh_key_path)

    user        = "ubuntu"
    host        = local.jaeger_collector_ips[count.index]
    private_key = file(var.ssh_key_path)
  }
}
#-------------------------------------------------------------------------------
# jaeger_collector (END)
#-------------------------------------------------------------------------------


#-------------------------------------------------------------------------------
# jaeger_query
#-------------------------------------------------------------------------------
locals {
  jaeger_query_ips = [for i in range(var.jaeger_query_count) : "${cidrhost(local.sn_cidrs[i], 16)}"]
}

resource "aws_security_group" "rkcy_jaeger_query" {
  name        = "rkcy_${var.stack}_jaeger_query"
  description = "Allow SSH and jaeger_query inbound traffic"
  vpc_id      = var.vpc.id

  ingress = [
    {
      cidr_blocks      = [ var.vpc.cidr_block ]
      description      = ""
      from_port        = 22
      to_port          = 22
      ipv6_cidr_blocks = []
      prefix_list_ids  = []
      protocol         = "tcp"
      security_groups  = []
      self             = false
    },
    {
      cidr_blocks      = [ var.vpc.cidr_block ]
      description      = "jaeger-query gRPC server"
      from_port        = 16685
      to_port          = 16685
      ipv6_cidr_blocks = []
      prefix_list_ids  = []
      protocol         = "tcp"
      security_groups  = []
      self             = false
    },
    {
      cidr_blocks      = [ var.vpc.cidr_block ]
      description      = "jaeger-query HTTP server"
      from_port        = var.jaeger_query_port
      to_port          = var.jaeger_query_port
      ipv6_cidr_blocks = []
      prefix_list_ids  = []
      protocol         = "tcp"
      security_groups  = []
      self             = false
    },
    {
      cidr_blocks      = [ var.vpc.cidr_block ]
      description      = "jaeger-query Admin server"
      from_port        = 16687
      to_port          = 16687
      ipv6_cidr_blocks = []
      prefix_list_ids  = []
      protocol         = "tcp"
      security_groups  = []
      self             = false
    },
    {
      cidr_blocks      = [ var.vpc.cidr_block ]
      description      = "node_exporter"
      from_port        = 9100
      to_port          = 9100
      ipv6_cidr_blocks = []
      prefix_list_ids  = []
      protocol         = "tcp"
      security_groups  = []
      self             = false
    },
  ]

  egress = [
    {
      cidr_blocks      = [ var.vpc.cidr_block ]
      description      = ""
      from_port        = 0
      ipv6_cidr_blocks = []
      prefix_list_ids  = []
      protocol         = "-1"
      security_groups  = []
      self             = false
      to_port          = 0
    }
  ]
}

resource "aws_network_interface" "jaeger_query" {
  count = var.jaeger_query_count
  subnet_id   = local.sn_ids[count.index]
  private_ips = [local.jaeger_query_ips[count.index]]

  security_groups = [aws_security_group.rkcy_jaeger_query.id]
}

resource "aws_placement_group" "jaeger_query" {
  name     = "rkcy_${var.stack}_jaeger_query_pc"
  strategy = "spread"
}

resource "aws_instance" "jaeger_query" {
  count = var.jaeger_query_count
  ami = data.aws_ami.telemetry.id
  instance_type = "m4.large"
  placement_group = aws_placement_group.jaeger_query.name

  key_name = aws_key_pair.telemetry.key_name

  network_interface {
    network_interface_id = aws_network_interface.jaeger_query[count.index].id
    device_index = 0
  }

  credit_specification {
    cpu_credits = "unlimited"
  }

  tags = {
    Name = "rkcy_${var.stack}_inst_jaeger_query_${count.index}"
  }
}

resource "aws_route53_record" "jaeger_query_private" {
  count = var.jaeger_query_count
  zone_id = var.dns_zone.zone_id
  name    = "jaeger-query-${count.index}.${var.stack}.local.${var.dns_zone.name}"
  type    = "A"
  ttl     = "300"
  records = [local.jaeger_query_ips[count.index]]
}

resource "null_resource" "jaeger_query_provisioner" {
  count = var.jaeger_query_count
  depends_on = [
    aws_instance.jaeger_query
  ]

  #---------------------------------------------------------
  # node_exporter
  #---------------------------------------------------------
  provisioner "file" {
    content = templatefile("${path.module}/../../../shared/node_exporter_install.sh", {})
    destination = "/home/ubuntu/node_exporter_install.sh"
  }
  provisioner "remote-exec" {
    inline = [
      <<EOF
sudo bash /home/ubuntu/node_exporter_install.sh
rm /home/ubuntu/node_exporter_install.sh
EOF
    ]
  }
  #---------------------------------------------------------
  # node_exporter (END)
  #---------------------------------------------------------

  provisioner "file" {
    content = templatefile("${path.module}/../../../shared/telemetry/jaeger-query.service.tpl", {
      elasticsearch_urls = var.elasticsearch_urls
    })
    destination = "/home/ubuntu/jaeger-query.service"
  }

  provisioner "remote-exec" {
    inline = [
      <<EOF
sudo mv /home/ubuntu/jaeger-query.service /etc/systemd/system/jaeger-query.service
sudo systemctl daemon-reload

%{for url in var.elasticsearch_urls}
RET=1
while [ $RET -ne 0 ]; do
  echo 'Trying elasticsearch ${url}'
  curl -s -f '${url}'
  RET=$?
  sleep 2
done
echo 'Connected elasticsearch ${url}'
%{endfor}

sudo systemctl start jaeger-query
sudo systemctl enable jaeger-query
EOF
    ]
  }

  connection {
    type     = "ssh"

    bastion_user        = "ubuntu"
    bastion_host        = var.bastion_ips[0]
    bastion_private_key = file(var.ssh_key_path)

    user        = "ubuntu"
    host        = local.jaeger_query_ips[count.index]
    private_key = file(var.ssh_key_path)
  }
}
#-------------------------------------------------------------------------------
# jaeger_query (END)
#-------------------------------------------------------------------------------

#-------------------------------------------------------------------------------
# otelcol
#-------------------------------------------------------------------------------
locals {
  otelcol_ips = [for i in range(var.otelcol_count) : "${cidrhost(local.sn_cidrs[i], 43)}"]
}

resource "aws_security_group" "rkcy_otelcol" {
  name        = "rkcy_${var.stack}_otelcol"
  description = "Allow SSH and otelcol inbound traffic"
  vpc_id      = var.vpc.id

  ingress = [
    {
      cidr_blocks      = [ var.vpc.cidr_block ]
      description      = ""
      from_port        = 22
      to_port          = 22
      ipv6_cidr_blocks = []
      prefix_list_ids  = []
      protocol         = "tcp"
      security_groups  = []
      self             = false
    },
    {
      cidr_blocks      = [ var.vpc.cidr_block ]
      description      = "otelcol gRPC listener"
      from_port        = var.otelcol_port
      to_port          = var.otelcol_port
      ipv6_cidr_blocks = []
      prefix_list_ids  = []
      protocol         = "tcp"
      security_groups  = []
      self             = false
    },
    {
      cidr_blocks      = [ var.vpc.cidr_block ]
      description      = "otelcol HTTP listener"
      from_port        = 4318
      to_port          = 4318
      ipv6_cidr_blocks = []
      prefix_list_ids  = []
      protocol         = "tcp"
      security_groups  = []
      self             = false
    },
    {
      cidr_blocks      = [ var.vpc.cidr_block ]
      description      = "otelcol Prometheus metrics"
      from_port        = 8888
      to_port          = 8888
      ipv6_cidr_blocks = []
      prefix_list_ids  = []
      protocol         = "tcp"
      security_groups  = []
      self             = false
    },
    {
      cidr_blocks      = [ var.vpc.cidr_block ]
      description      = "otelcol Prometheus metrics"
      from_port        = 9999
      to_port          = 9999
      ipv6_cidr_blocks = []
      prefix_list_ids  = []
      protocol         = "tcp"
      security_groups  = []
      self             = false
    },
    {
      cidr_blocks      = [ var.vpc.cidr_block ]
      description      = "otelcol legacy HTTP listener"
      from_port        = 55681
      to_port          = 55681
      ipv6_cidr_blocks = []
      prefix_list_ids  = []
      protocol         = "tcp"
      security_groups  = []
      self             = false
    },
    {
      cidr_blocks      = [ var.vpc.cidr_block ]
      description      = "node_exporter"
      from_port        = 9100
      to_port          = 9100
      ipv6_cidr_blocks = []
      prefix_list_ids  = []
      protocol         = "tcp"
      security_groups  = []
      self             = false
    },
  ]

  egress = [
    {
      cidr_blocks      = [ var.vpc.cidr_block ]
      description      = ""
      from_port        = 0
      ipv6_cidr_blocks = []
      prefix_list_ids  = []
      protocol         = "-1"
      security_groups  = []
      self             = false
      to_port          = 0
    }
  ]
}

resource "aws_network_interface" "otelcol" {
  count = var.otelcol_count
  subnet_id   = local.sn_ids[count.index]
  private_ips = [local.otelcol_ips[count.index]]

  security_groups = [aws_security_group.rkcy_otelcol.id]
}

resource "aws_placement_group" "otelcol" {
  name     = "rkcy_${var.stack}_otelcol_pc"
  strategy = "spread"
}

resource "aws_instance" "otelcol" {
  count = var.otelcol_count
  ami = data.aws_ami.telemetry.id
  instance_type = "m4.large"
  placement_group = aws_placement_group.otelcol.name

  key_name = aws_key_pair.telemetry.key_name

  network_interface {
    network_interface_id = aws_network_interface.otelcol[count.index].id
    device_index = 0
  }

  credit_specification {
    cpu_credits = "unlimited"
  }

  tags = {
    Name = "rkcy_${var.stack}_inst_otelcol_${count.index}"
  }
}

resource "aws_route53_record" "otelcol_private" {
  count = var.otelcol_count
  zone_id = var.dns_zone.zone_id
  name    = "otelcol-${count.index}.${var.stack}.local.${var.dns_zone.name}"
  type    = "A"
  ttl     = "300"
  records = [local.otelcol_ips[count.index]]
}

resource "null_resource" "otelcol_provisioner" {
  count = var.otelcol_count
  depends_on = [
    aws_instance.otelcol
  ]

  #---------------------------------------------------------
  # node_exporter
  #---------------------------------------------------------
  provisioner "file" {
    content = templatefile("${path.module}/../../../shared/node_exporter_install.sh", {})
    destination = "/home/ubuntu/node_exporter_install.sh"
  }
  provisioner "remote-exec" {
    inline = [
      <<EOF
sudo bash /home/ubuntu/node_exporter_install.sh
rm /home/ubuntu/node_exporter_install.sh
EOF
    ]
  }
  #---------------------------------------------------------
  # node_exporter (END)
  #---------------------------------------------------------

  provisioner "file" {
    content = templatefile("${path.module}/../../../shared/telemetry/otelcol.service.tpl", {
      elasticsearch_urls = var.elasticsearch_urls
    })
    destination = "/home/ubuntu/otelcol.service"
  }

  provisioner "file" {
    content = templatefile("${path.module}/../../../shared/telemetry/otelcol.yaml.tpl", {
      jaeger_collector = "${ var.nginx_hosts.app[0] }:${ var.jaeger_collector_port }"
    })
    destination = "/home/ubuntu/otelcol.yaml"
  }

  provisioner "remote-exec" {
    inline = [
      <<EOF
sudo mv /home/ubuntu/otelcol.service /etc/systemd/system/otelcol.service
sudo mv /home/ubuntu/otelcol.yaml /etc/otelcol.yaml
sudo chown telem:telem /etc/otelcol.yaml
sudo chmod 660 /etc/otelcol.yaml

sudo systemctl daemon-reload
sudo systemctl start otelcol
sudo systemctl enable otelcol
EOF
    ]
  }

  connection {
    type     = "ssh"

    bastion_user        = "ubuntu"
    bastion_host        = var.bastion_ips[0]
    bastion_private_key = file(var.ssh_key_path)

    user        = "ubuntu"
    host        = local.otelcol_ips[count.index]
    private_key = file(var.ssh_key_path)
  }
}
#-------------------------------------------------------------------------------
# otelcol (END)
#-------------------------------------------------------------------------------

output "jaeger_collector_hosts" {
  value = sort(aws_route53_record.jaeger_collector_private.*.name)
}
output "jaeger_collector_port" {
  value = var.jaeger_collector_port
}

output "jaeger_query_hosts" {
  value = sort(aws_route53_record.jaeger_query_private.*.name)
}
output "jaeger_query_port" {
  value = var.jaeger_query_port
}

output "otelcol_hosts" {
  value = sort(aws_route53_record.otelcol_private.*.name)
}
output "otelcol_port" {
  value = var.otelcol_port
}
