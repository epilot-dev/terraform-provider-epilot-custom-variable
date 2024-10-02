terraform {
  required_providers {
    epilot-custom-variable = {
      source  = "epilot-dev/epilot-custom-variable"
      version = "1.0.0"
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
  server_url  = var.custom_variables_api_url
}

# import {
#   to = epilot-custom-variable_custom_variable.my_variable
#   id = "25bad807-bf36-4ce4-8aca-27b6420597ac"
# }