package account

import (
	"fmt"

	"github.com/GaloyMoney/terraform-provider-bria/bria/account"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceBriaAccountSignerConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceBriaAccountSignerConfigCreate,
		Update: resourceBriaAccountSignerConfigUpdate,
		Read:   resourceBriaAccountSignerConfigRead,
		Delete: resourceBriaAccountSignerConfigDelete,

		Schema: map[string]*schema.Schema{
			"xpub": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the xpub.",
			},
			"lnd": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem:        lndConfigElem(),
				Description: "LND signer configuration.",
			},
		},
	}
}

func lndConfigElem() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"endpoint": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The LND endpoint.",
			},
			"macaroon_base64": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The base64 encoded macaroon.",
			},
			"cert": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The LND certificate.",
			},
		},
	}
}

func resourceBriaAccountSignerConfigCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*account.AccountClient)

	xpub := d.Get("xpub").(string)

	var err error
	if lndConfig, ok := d.GetOk("lnd"); ok {
		err = client.CreateLndSignerConfig(xpub, lndConfig.([]interface{}))
	} else {
		return fmt.Errorf("lnd block must be provided")
	}

	if err != nil {
		return fmt.Errorf("error creating Bria signer config: %w", err)
	}

	d.SetId(xpub)

	return resourceBriaAccountSignerConfigRead(d, meta)
}

func resourceBriaAccountSignerConfigRead(d *schema.ResourceData, meta interface{}) error {
	// Implement the read function for the bria_account_signer_config resource
	// This can be a no-op if there is no way to read a signer config from the API
	return nil
}

func resourceBriaAccountSignerConfigDelete(d *schema.ResourceData, meta interface{}) error {
	// Implement the delete function for the bria_account_signer_config resource
	// If the API does not provide a delete functionality, you can set the ID to an empty string
	d.SetId("")
	return nil
}

func resourceBriaAccountSignerConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	return fmt.Errorf("briaaccount_update resource does not support updates")
}
