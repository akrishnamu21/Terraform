//variable "agility_userid" {}
//variable "agility_password" {}
provider "agility" {
  userid = "${var.userid}"
  password = "${var.password}"
}


#checkin

resource "agility_checkin" "checkinterraform"{
  containername="${var.container_name}"
  headversionallowed="${var.headversion_allowed}"
  assetname="${var.asset_name}"
  asset="${var.asset_type}"
  projectname="${var.project_name}"
}
