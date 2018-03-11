package agility

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {

	// The username and password are retrieved by terraform from the environment variables TF_VAR_agility_userid and TF_VAR_agility_password
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"userid": &schema.Schema{
				Type:     schema.TypeString,
				Required: 	true,
				Computed: 	false,
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Required: 	true,
				Computed: 	false,
			},
		},

		//define the supported resources and point to their respective .go classes
		ResourcesMap: map[string]*schema.Resource{
			//"agility_compute":				resourceAgilityCompute(),
			//"agility_createcontainer":		resourceCreateSubContainer(),
			//"agility_createproject":		resourceCreateSubProject(),
			//"agility_createenvironment":	resourceCreateEnvironments(),
			//"agility_addcloudprovider":		resourceAddCloudProvider(),
			//"agility_synccloudprovider":	resourceSyncCloudProvider(),
			//"agility_license":  			resourceLicenseUpload(),
			//"agility_createstack":			resourceCreateStack(),
			//"agility_createscript":				resourceCreateScript(),
			//"agility_attachment":					resourceAttachment(),
			//"agility_createpackage":				resourceCreatePackage(),
			"agility_firewall":						resourceFirewall(),
			//"agility_checkin":					resourceCheckIn(),
			//"agility_approve":							resourceApprove(),
			//"agility_assignpolicy":					resourceAssignPolicy(),
			//"agility_blueprint":					resourceCreateBlueprint(),
			//"agility_unassignpolicy":				resourceUnassignPolicy(),
			//"agility_publish":						resourceAgilityPublish(),
			//"agility_scriptId":						resourceWriteScriptId(),
		},

		ConfigureFunc: providerConfigure,
	}
}

//var descriptions map[string]string

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	creds := ProvCredentials{
		UserName:        d.Get("userid").(string),
		Password:        d.Get("password").(string),
	}
	return creds, nil
}
