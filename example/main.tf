terraform {
  required_providers {
    admin = {
      source = "GaloyMoney/bria-admin"
      version = "0.1.0"
    }
  }
}

provider "bria_admin" {}
