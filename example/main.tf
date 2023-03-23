terraform {
  required_providers {
    briaadmin = {
      source = "galoymoney/briaadmin"
      version = "0.1.0"
    }
  }
}

provider "briaadmin" {
  # You can add any required configuration for the admin provider here
}

resource "briaadmin_dummy" "example" {
  # You can add any required attributes for the bria_admin_dummy resource here
}
