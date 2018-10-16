package nks

import (
	"time"

	"github.com/StackPointCloud/stackpoint-sdk-go/stackpointio"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a schema.Provider for StackPoint
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SPC_API_TOKEN", nil),
				Description: "The token key for API operations.",
			},
			"endpoint": {
				Type:     schema.TypeString,
				Required: true,
				DefaultFunc: schema.EnvDefaultFunc("SPC_BASE_API_URL",
					"https://api.stackpoint.io/"),
				Description: "The endpoint URL for API operations.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"nks_cluster":     resourceNKSCluster(),
			"nks_master_node": resourceNKSMasterNode(),
			"nks_nodepool":    resourceNKSNodePool(),
			"nks_solution":    resourceNKSSolution(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"nks_instance_specs": dataSourceNKSInstanceSpecs(),
			"nks_keysets":        dataSourceNKSKeysets(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Token:    d.Get("token").(string),
		EndPoint: d.Get("endpoint").(string),
		Client:   stackpointio.NewClient(d.Get("token").(string), d.Get("endpoint").(string)),
	}
	return &config, nil
}

var resourceDefaultTimeouts = schema.ResourceTimeout{
	Create:  schema.DefaultTimeout(40 * time.Minute),
	Update:  schema.DefaultTimeout(40 * time.Minute),
	Delete:  schema.DefaultTimeout(40 * time.Minute),
	Default: schema.DefaultTimeout(40 * time.Minute),
}
