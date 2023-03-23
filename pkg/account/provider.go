package account

import (
	"github.com/GaloyMoney/terraform-provider-bria/bria/account"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"endpoint": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("BRIA_API_URL", "localhost:2742"),
				Description: "The API endpoint for Bria.",
			},
			"api_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("BRIA_API_KEY", ""),
				Description: "The API key for Bria.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"briaaccount_xpub":   resourceBriaAccountXpub(),
			"briaaccount_wallet": resourceBriaAccountWallet(),
			// "briaaccount_batch_group": resourceBriaAccountBatchGroup(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	endpoint := d.Get("endpoint").(string)
	apiKey := d.Get("api_key").(string)

	return account.NewAccountClient(endpoint, apiKey)
}
