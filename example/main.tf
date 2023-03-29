terraform {
  required_providers {
    briaadmin = {
      source  = "galoymoney/briaadmin"
      version = "0.1.0"
    }
  }
}

resource "random_string" "postfix" {
  length  = 6
  special = false
  upper   = false
  numeric  = false
}

resource "briaadmin_account" "example" {
  name = "tf-example-${random_string.postfix.result}"
}
