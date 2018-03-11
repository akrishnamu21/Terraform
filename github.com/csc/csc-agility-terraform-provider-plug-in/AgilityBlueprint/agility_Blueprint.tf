//variable "agility_userid" {}
//variable "agility_password" {}
provider "agility" {
  userid = "${var.userid}"
  password = "${var.password}"
}




#Assign Policy
/*
resource "agility_assignpolicy" "assignpolicyterraform"{
  projectname="${var.project_name}"
  policyname="${var.policy_name}"
 //depends_on = ["agility_firewall.terraformfirewall"]
}

*/

#Create Blueprint

resource "agility_blueprint" "createblueprintterraform" {
  projectname = "${var.project_name}"
  blueprintname = "${var.blueprint_name}"
  blueprintdesc = "${var.blueprint_desc}"
  stackname = "${var.stack_name}"
  packagename = "${var.package_name}"
  policyname = "${var.policy_name}"
  headversionallowed = "${var.headversion_allow}"
  workloadname = "${var.workload_name}"
  policyassignmentname = "${var.policyassignment_name}"
  //depends_on = ["agility_assignpolicy.assignpolicyterraform"]
}


#unassign
/*
resource "agility_unassignpolicy" "unassignpolicyterraform"{
  projectname="Agility Factory"
  policyname="CollectorFirewallTerraform"
  depends_on = ["agility_blueprint.createblueprintterraform"]
}
*/
