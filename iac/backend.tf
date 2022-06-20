terraform {
  backend "gcs" {
    bucket  = "banksimulation-tfstate"
    prefix  = "terraform.tfstate"
  }
}