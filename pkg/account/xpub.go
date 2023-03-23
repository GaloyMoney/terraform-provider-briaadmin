package account

import (
	"fmt"

	"github.com/GaloyMoney/terraform-provider-bria/bria/account"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceBriaAccountXpub() *schema.Resource {
	return &schema.Resource{
		Create: resourceBriaAccountXpubCreate,
		Read:   resourceBriaAccountXpubRead,
		Update: resourceBriaAccountXpubUpdate,
		Delete: resourceBriaAccountXpubSoftDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the xpub.",
			},
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of the xpub.",
			},
			"xpub": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The extended public key.",
			},
			"derivation": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The derivation path for the xpub.",
			},
		},
	}
}

func resourceBriaAccountXpubCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*account.AccountClient)

	name := d.Get("name").(string)
	xpub := d.Get("xpub").(string)
	derivation := d.Get("derivation").(string)

	resp, err := client.ImportXpub(name, xpub, derivation)
	if err != nil {
		return fmt.Errorf("error creating Bria xpub: %w", err)
	}

	d.SetId(resp.Id)

	return resourceBriaAccountXpubRead(d, meta)
}

func resourceBriaAccountXpubRead(d *schema.ResourceData, meta interface{}) error {
	// Implement the read function for the bria_account_xpub resource
	// This can be a no-op if there is no way to read an xpub from the API
	return nil
}

func resourceBriaAccountXpubSoftDelete(d *schema.ResourceData, meta interface{}) error {
	// Implement the soft delete function for the bria_account_xpub resource
	// If the API does not provide a delete functionality, you can set the ID to an empty string
	d.SetId("")
	return nil
}

func resourceBriaAccountXpubUpdate(d *schema.ResourceData, meta interface{}) error {
	return fmt.Errorf("briaaccount_xpub resource does not support updates")
}
