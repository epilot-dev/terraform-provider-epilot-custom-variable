terraform {
  required_providers {
    epilot-variable = {
      source  = "epilot-dev/epilot-variable"
      version = "0.10.5"
    }
  }
}

variable "epilot_auth" {
  type = string
}
variable "template_variables_api_url" {
  type    = string
  default = "https://template-variables-api.dev.sls.epilot.io"
}

provider "epilot-variable" {
  epilot_auth = var.epilot_auth
  server_url = var.template_variables_api_url
}

resource "epilot-variable_custom_variable" "my_variable" {
  # (resource arguments)
}