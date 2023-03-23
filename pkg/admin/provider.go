package admin

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"endpoint": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("BRIA_ADMIN_API_URL", "localhost:2743"),
				Description: "The API endpoint for the Bria admin service.",
			},
			"api_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("BRIA_ADMIN_API_KEY", ""),
				Description: "The API key for the Bria admin service.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"briaadmin_dummy": resourceBriaAdminDummy(),
		},

		ConfigureFunc: providerConfigure,
	}
}

type AdminClient struct {
	Endpoint string
	APIKey   string
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	endpoint := d.Get("endpoint").(string)
	apiKey := d.Get("api_key").(string)

	return &AdminClient{
		Endpoint: endpoint,
		APIKey:   apiKey,
	}, nil
}

func resourceBriaAdminDummy() *schema.Resource {
	return &schema.Resource{
		Create: resourceBriaAdminDummyCreate,
		Read:   resourceBriaAdminDummyRead,
		Update: resourceBriaAdminDummyUpdate,
		Delete: resourceBriaAdminDummyDelete,

		Schema: map[string]*schema.Schema{
			// Add any required attributes for the bria_admin_dummy resource here
		},
	}
}

func resourceBriaAdminDummyCreate(d *schema.ResourceData, meta interface{}) error {
	// Implement the create function for the bria_admin_dummy resource
	d.SetId("example_id")
	return resourceBriaAdminDummyRead(d, meta)
}

func resourceBriaAdminDummyRead(d *schema.ResourceData, meta interface{}) error {
	// Implement the read function for the bria_admin_dummy resource
	return nil
}

func resourceBriaAdminDummyUpdate(d *schema.ResourceData, meta interface{}) error {
	// Implement the update function for the bria_admin_dummy resource
	return resourceBriaAdminDummyRead(d, meta)
}

func resourceBriaAdminDummyDelete(d *schema.ResourceData, meta interface{}) error {
	// Implement the delete function for the bria_admin_dummy resource
	d.SetId("")
	return nil
}
