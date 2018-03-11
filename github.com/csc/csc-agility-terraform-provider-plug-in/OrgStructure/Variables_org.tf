variable "userid" {
    description="The Username of the Agility Platform"
    default = ""
}
variable "password" {
    description="The Username of the Agility Platform"
    default = ""
}
variable "headcontainername" {
    description="The name of the Parent Container",
    default ="Root"
}

variable "parentcontainer3" {
    default="Test"
}


variable "project2" {
    default = "TEST"
}

variable "environment1" {
    default = "Dev"
}

variable "environment1_type" {
    default = "DEV"
}

variable "cloud_name" {
    default = "varcloud_name1"
}
variable "cloud_type" {
    default = "varcloud_type1"
}
variable "hostname" {
    default = "varhostname1"
}
variable "aws_accesskey" {
    default = "varaws_accesskey1"
}
variable "aws_secretkey" {
    default = "varaws_secretkey1"
}
variable "aws_accountnumber" {
    default = "varaws_accountnumber1"
}



variable "project_name" {
  default = "varproject_name"
}
variable "image_name" {
  default = "varimage_name"
}
variable "stack_name" {
  default = "varstack_name"
}
variable "stack_desc" {
  default = "varstack_desc"
}
variable "os_name" {
  default = "varos_name"
}

