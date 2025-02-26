resource "null_resource" "jaeger_collector_provisioner" {
  #---------------------------------------------------------
  # node_exporter
  #---------------------------------------------------------
  provisioner "remote-exec" {
    inline = ["sudo hostnamectl set-hostname ${var.hostname}"]
  }
  provisioner "file" {
    content = templatefile("${path.module}/../../node_exporter_install.sh", {})
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
    content = templatefile("${path.module}/jaeger-collector.service.tpl", {
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
    bastion_host        = var.bastion_ip
    bastion_private_key = file(var.ssh_key_path)

    user        = "ubuntu"
    host        = var.jaeger_collector_ip
    private_key = file(var.ssh_key_path)
  }
}
