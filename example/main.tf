terraform {
  required_providers {
    briaadmin = {
      source = "galoymoney/briaadmin"
      version = "0.1.0"
    }
    briaaccount = {
      source = "galoymoney/briaaccount"
      version = "0.1.0"
    }
  }
}

provider "briaadmin" {
  api_key = "bria_admin_SodrZAMNgGJWVoRRpPh1Uw3388Ku8VmbWYtPAeKB3wnKdy1aGjQjkrSLPAg78lMJ"
}

resource "briaadmin_account" "example" {
  name = "tf-example-2"
}

provider "briaaccount" {
  api_key = briaadmin_account.example.api_key
}

resource "briaaccount_xpub" "lnd" {
  name = "lnd"
  xpub = "tpubDDEGUyCLufbxAfQruPHkhUcu55UdhXy7otfcEQG4wqYNnMfq9DbHPxWCqpEQQAJUDi8Bq45DjcukdDAXasKJ2G27iLsvpdoEL5nTRy5TJ2B"
  derivation = "m/64h/1h/0"
}
