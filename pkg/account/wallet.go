package account

import (
	"fmt"

	"github.com/GaloyMoney/terraform-provider-bria/bria/account"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceBriaAccountWallet() *schema.Resource {
	return &schema.Resource{
		Create: resourceBriaAccountWalletCreate,
		Update: resourceBriaAccountWalletUpdate,
		Read:   resourceBriaAccountWalletRead,
		Delete: resourceBriaAccountWalletSoftDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the wallet.",
			},
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of the wallet.",
			},
			"xpubs": {
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Required:    true,
				Description: "A list of xpub reference IDs associated with the wallet.",
			},
		},
	}
}

func resourceBriaAccountWalletCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*account.AccountClient)

	name := d.Get("name").(string)
	xpubRefsRaw := d.Get("xpubs").([]interface{})
	xpubRefs := make([]string, len(xpubRefsRaw))
	for i, v := range xpubRefsRaw {
		xpubRefs[i] = v.(string)
	}

	resp, err := client.CreateWallet(name, xpubRefs)
	if err != nil {
		return fmt.Errorf("error creating Bria wallet: %w", err)
	}

	d.SetId(resp.Id)

	return resourceBriaAccountWalletRead(d, meta)
}

func resourceBriaAccountWalletRead(d *schema.ResourceData, meta interface{}) error {
	// Implement the read function for the bria_account_wallet resource
	// This can be a no-op if there is no way to read a wallet from the API
	return nil
}

func resourceBriaAccountWalletUpdate(d *schema.ResourceData, meta interface{}) error {
	return fmt.Errorf("briaaccount_wallet resource does not support updates")
}

func resourceBriaAccountWalletSoftDelete(d *schema.ResourceData, meta interface{}) error {
	// Implement the soft delete function for the bria_account_wallet resource
	// If the API does not provide a delete functionality, you can set the ID to an empty string
	d.SetId("")
	return nil
}
