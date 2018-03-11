
provider "agility" {
    userid = "${var.userid}"
    password = "${var.password}"
}

resource "agility_addcloudprovider" "AWS" {
  cloudname = "${var.cloud_name}"
  description = "Cloud provider created from terraform"
  cloudtype = "${var.cloud_type}"
  hostname = "${var.hostname}"
  publickey = "${var.aws_accesskey}"
  privatekey = "${var.aws_secretkey}"
  awsaccountnumber = "${var.aws_accountnumber}"

}
resource "agility_synccloudprovider" "SyncAWS"{
  cloudname ="${var.cloud_name}"
  depends_on = ["agility_addcloudprovider.AWS"]
}