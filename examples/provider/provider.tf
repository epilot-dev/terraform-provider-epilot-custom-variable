terraform {
  required_providers {
    epilot-custom-variable = {
      source  = "epilot-dev/epilot-custom-variable"
      version = "1.3.1"
    }
  }
}

provider "epilot-custom-variable" {
  server_url = "..." # Optional
}