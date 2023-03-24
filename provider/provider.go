package provider

import (
	"fmt"

	"github.com/GaloyMoney/terraform-provider-bria-admin/bria"
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
			"briaadmin_account": resourceBriaAdminAccount(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	endpoint := d.Get("endpoint").(string)
	apiKey := d.Get("api_key").(string)

	return bria.NewAdminClient(endpoint, apiKey)
}

func resourceBriaAdminAccount() *schema.Resource {
	return &schema.Resource{
		Create: resourceBriaAdminAccountCreate,
		Read:   resourceBriaAdminAccountRead,
		Update: resourceBriaAdminAccountUpdate,
		Delete: resourceBriaAdminAccountDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the account.",
			},
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of the account.",
			},
			"api_key_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of the API key associated with the account.",
			},
			"api_key": {
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
				Description: "The API key associated with the account.",
			},
		},
	}
}

func resourceBriaAdminAccountCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*bria.AdminClient)

	accountName := d.Get("name").(string)

	resp, err := client.CreateAccount(accountName)
	if err != nil {
		return fmt.Errorf("error creating Bria admin account: %w", err)
	}

	d.SetId(resp.Key.AccountId)
	d.Set("api_key_id", resp.Key.Id)
	d.Set("api_key", resp.Key.Key)

	return resourceBriaAdminAccountRead(d, meta)
}

func resourceBriaAdminAccountRead(d *schema.ResourceData, meta interface{}) error {
	// Implement the read function for the bria_admin_dummy resource
	return nil
}

func resourceBriaAdminAccountUpdate(d *schema.ResourceData, meta interface{}) error {
	return fmt.Errorf("briaadmin_account resource does not support updates")
}

func resourceBriaAdminAccountDelete(d *schema.ResourceData, meta interface{}) error {
	// Soft delete: just remove the account from the Terraform state
	d.SetId("")
	return nil
}
