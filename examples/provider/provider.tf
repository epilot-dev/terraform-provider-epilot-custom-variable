terraform {
  required_providers {
    epilot-custom-variable = {
      source  = "epilot-dev/epilot-custom-variable"
      version = "0.10.6"
    }
  }
}

variable "epilot_auth" {
  type = string
}
variable "custom_variables_api_url" {
  type    = string
  default = "https://template-variables-api.dev.sls.epilot.io"
}

provider "epilot-custom-variable" {
  epilot_auth = var.epilot_auth
  server_url = var.custom_variables_api_url
}

resource "epilot-variable_custom_variable" "my_variable" {
  # (resource arguments)
}