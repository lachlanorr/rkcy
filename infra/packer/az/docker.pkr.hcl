variable "az_subscription_id" {}
variable "az_client_id" {}
variable "az_client_secret" {}
variable "az_tenant_id" {}
variable "az_image_resource_group" {}

source "azure-arm" "ubuntu" {
  subscription_id = "${var.az_subscription_id}"
  client_id =       "${var.az_client_id}"
  client_secret =   "${var.az_client_secret}"
  tenant_id =       "${var.az_tenant_id}"

  managed_image_resource_group_name = "${var.az_image_resource_group}"
  managed_image_name = "rkcy-docker-{{isotime `20060102-150405`}}"

  os_type = "Linux"
  image_publisher = "Canonical"
  image_offer = "0001-com-ubuntu-server-focal"
  image_sku = "20_04-lts"

  location = "Central US"
  vm_size = "Standard_DS2_v2"

  ssh_username = "ubuntu"
}

build {
  sources = ["source.azure-arm.ubuntu" ]

  provisioner "shell" {
    script = "../shared/install_baseline.sh"
  }
  provisioner "shell" {
    script = "../shared/install_hashi.sh"
  }
#  provisioner "shell" {
#    script = "../shared/install_docker.sh"
#  }
}
