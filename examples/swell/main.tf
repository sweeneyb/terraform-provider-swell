terraform {
  required_providers {
    swell = {
      version = "0.2"
      source  = "briansweeney-dev/edu/swell"
    }
  }
}

data "swell_products" "all" {}

# Returns all products
output "all_products" {
  value = data.swell_products.all.products
}

variable "item_name" {
  type    = string
  default = "Moccamaster KBS Black"
}


# Only returns packer spiced latte
output "specific_product" {
  value = {
    for product in data.swell_products.all.products :
    product.name => product
    if product.name == var.item_name
  }
}

data "swell_product" "AeroPress" {
  name = "AeroPress Paper Filters (350 pack)"
}

output "AeroPress" {
  value = data.swell_product.AeroPress
}

resource "swell_category" "coffee" {
  name = "coffee items"
  description = "items for coffee"
  active = true
}