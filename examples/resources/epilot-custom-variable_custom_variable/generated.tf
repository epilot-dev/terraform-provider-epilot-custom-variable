# __generated__ by Terraform
# Please review these resources and move them into your main configuration files.

# __generated__ by Terraform from "367e96ab-8d6f-49ee-b449-74f1ef4abf93"
resource "epilot-custom-variable_custom_variable" "test_update" {
  lifecycle {
    prevent_destroy = true
  }
  config = jsonencode({
    body = {
      amount_tax = {
        enable = true
        style = {
          border         = "none !important"
          color          = "#222"
          font-family    = "Arial,sans-serif"
          font-size      = "14px"
          font-style     = "normal"
          overflow-wrap  = "break-word"
          padding-bottom = "8px"
          padding-top    = "16px"
          text-align     = "left"
          vertical-align = "top"
          width          = "10%"
          word-wrap      = "break-word"
        }
      }
      composite_component = {
        enable        = true
        parent_column = "item"
      }
      composite_component_price = {
        enable        = true
        parent_column = "item"
      }
      gross_total = {
        enable = true
        style = {
          border         = "none !important"
          color          = "#222"
          font-family    = "Arial,sans-serif"
          font-size      = "14px"
          font-style     = "normal"
          overflow-wrap  = "break-word"
          padding-bottom = "8px"
          padding-right  = "12px"
          padding-top    = "16px"
          text-align     = "right"
          vertical-align = "top"
          width          = "20%"
          word-wrap      = "break-word"
        }
      }
      net_total = {
        enable = false
        style = {
          border         = "none !important"
          color          = "#222"
          font-family    = "Arial,sans-serif"
          font-size      = "14px"
          font-style     = "normal"
          overflow-wrap  = "break-word"
          padding-bottom = "8px"
          padding-right  = "18px"
          padding-top    = "16px"
          text-align     = "right"
          vertical-align = "top"
          width          = "15%"
          word-wrap      = "break-word"
        }
      }
      payment_type = {
        enable = true
        style = {
          color       = "#222"
          font-family = "Arial,sans-serif"
          font-size   = "12px"
          font-style  = "normal"
          opacity     = "50%"
        }
      }
      position = {
        enable        = true
        parent_column = "item"
      }
      price_description = {
        enable        = false
        parent_column = "item"
        style = {
          color         = "#222"
          font-family   = "Arial,sans-serif"
          font-size     = "12px"
          font-style    = "normal"
          overflow-wrap = "break-word"
          text-align    = "left"
          white-space   = "pre-line"
          width         = "24%"
          word-wrap     = "break-word"
        }
      }
      price_long_description = {
        enable        = false
        parent_column = "item"
        style = {
          color         = "#222"
          font-family   = "Arial,sans-serif"
          font-size     = "12px"
          font-style    = "normal"
          font-weight   = "normal"
          overflow-wrap = "break-word"
          text-align    = "left"
          white-space   = "pre-line"
          width         = "24%"
          word-wrap     = "break-word"
        }
      }
      product_description = {
        enable        = true
        parent_column = "item"
        style = {
          color         = "#222"
          font-family   = "Arial,sans-serif"
          font-size     = "12px"
          font-style    = "normal"
          font-weight   = "normal"
          overflow-wrap = "break-word"
          text-align    = "left"
          white-space   = "pre-line"
          width         = "24%"
          word-wrap     = "break-word"
        }
      }
      product_name = {
        enable        = true
        parent_column = "item"
        style = {
          border         = "none !important"
          color          = "#222"
          font-family    = "Arial,sans-serif"
          font-size      = "14px"
          font-style     = "normal"
          overflow-wrap  = "break-word"
          padding-bottom = "8px"
          padding-left   = "8px"
          padding-top    = "16px"
          text-align     = "left"
          vertical-align = "top"
          width          = "24%"
          word-wrap      = "break-word"
        }
      }
      quantity = {
        enable = true
        style = {
          border         = "none !important"
          color          = "#222"
          font-family    = "Arial,sans-serif"
          font-size      = "14px"
          font-style     = "normal"
          min-width      = "20px"
          overflow-wrap  = "break-word"
          padding-bottom = "8px"
          padding-left   = "8px"
          padding-right  = "14px"
          padding-top    = "16px"
          text-align     = "right"
          vertical-align = "top"
          width          = "6%"
          word-wrap      = "break-word"
        }
      }
      tax = {
        enable = true
        style = {
          border         = "none !important"
          color          = "#222"
          font-family    = "Arial,sans-serif"
          font-size      = "14px"
          font-style     = "normal"
          overflow-wrap  = "break-word"
          padding-bottom = "8px"
          padding-right  = "38px"
          padding-top    = "16px"
          text-align     = "right"
          vertical-align = "top"
          width          = "13%"
          word-wrap      = "break-word"
        }
      }
      unit_amount = {
        enable = true
        style = {
          border         = "none !important"
          color          = "#222"
          font-family    = "Arial,sans-serif"
          font-size      = "14px"
          font-style     = "normal"
          overflow-wrap  = "break-word"
          padding-bottom = "8px"
          padding-top    = "16px"
          text-align     = "left"
          vertical-align = "top"
          width          = "14%"
          word-wrap      = "break-word"
        }
      }
    }
    footer = {
      amount_tax = {
        enable = true
        style = {
          border         = "none !important"
          color          = "#222"
          font-family    = "Arial,sans-serif"
          font-size      = "14px"
          font-style     = "normal"
          font-weight    = "bold"
          padding-bottom = "8px"
          padding-right  = "12px"
          padding-top    = "12px"
          text-align     = "right"
          vertical-align = "top"
          white-space    = "nowrap"
        }
      }
      gross_total = {
        enable = true
        style = {
          border         = "none !important"
          color          = "#222"
          font-family    = "Arial,sans-serif"
          font-size      = "14px"
          font-style     = "normal"
          font-weight    = "bold"
          padding-bottom = "8px"
          padding-right  = "12px"
          padding-top    = "12px"
          text-align     = "right"
          vertical-align = "top"
          white-space    = "nowrap"
        }
      }
      net_total = {
        enable = false
        style = {
          border         = "none !important"
          color          = "#222"
          font-family    = "Arial,sans-serif"
          font-size      = "14px"
          font-style     = "normal"
          font-weight    = "bold"
          padding-bottom = "8px"
          padding-right  = "12px"
          padding-top    = "12px"
          text-align     = "right"
          vertical-align = "top"
          white-space    = "nowrap"
        }
      }
      payment_type = {
        enable = true
        style = {
          border         = "none !important"
          color          = "#222"
          font-family    = "Arial,sans-serif"
          font-size      = "14px"
          font-style     = "normal"
          font-weight    = "bold"
          padding-bottom = "8px"
          padding-right  = "12px"
          padding-top    = "12px"
          text-align     = "left"
          vertical-align = "none !important"
          white-space    = "nowrap"
        }
      }
    }
    header = {
      columns = [{
        draggable = true
        enable    = true
        id        = "item"
        index     = 0
        label = {
          de = "Produkt"
          en = "Item"
        }
        style = {
          padding-left = "8px"
          width        = "34%"
        }
        }, {
        draggable = true
        enable    = true
        id        = "quantity"
        index     = 1
        label = {
          de = "Menge"
          en = "Quantity"
        }
        style = {
          overflow-wrap = "break-word"
          padding-left  = "8px"
          width         = "5%"
          word-wrap     = "break-word"
        }
        }, {
        draggable = true
        enable    = false
        id        = "unit_amount"
        index     = 2
        label = {
          de = "Einzelpreis"
          en = "Unit Price"
        }
        style = {
          overflow-wrap = "break-word"
          width         = "8%"
          word-wrap     = "break-word"
        }
        }, {
        draggable = false
        enable    = true
        id        = "net_total"
        index     = 3
        label = {
          de = "Nettosumme"
          en = "Net total"
        }
        style = {
          overflow-wrap = "break-word"
          width         = "10%"
          word-wrap     = "break-word"
        }
        }, {
        draggable = true
        enable    = true
        id        = "tax"
        index     = 4
        label = {
          de = "Steuerrate"
          en = "Tax Rate"
        }
        style = {
          overflow-wrap = "break-word"
          width         = "10%"
          word-wrap     = "break-word"
        }
        }, {
        draggable = false
        enable    = false
        id        = "amount_tax"
        index     = 5
        label = {
          de = "Steuer"
          en = "Tax"
        }
        style = {
          overflow-wrap = "break-word"
          width         = "18%"
          word-wrap     = "break-word"
        }
        }, {
        draggable = false
        enable    = true
        id        = "gross_total"
        index     = 6
        label = {
          de = "Bruttosumme"
          en = "Total"
        }
        style = {
          overflow-wrap = "break-word"
          width         = "10%"
          word-wrap     = "break-word"
        }
      }]
      style = {
        background     = "none"
        border         = "none !important"
        color          = "#222"
        font-family    = "Arial,sans-serif"
        font-size      = "14px"
        font-weight    = "bold"
        padding-bottom = "4px"
        text-align     = "left"
      }
    }
  })
  helper_logic  = null
  helper_params = []
  key           = "Bestelluebersicht_mit_Positionen"
  name          = "Bestelluebersicht_mit_Positionen"
  tags          = []
  template      = "\n\nasdads<table style=\"table-layout: fixed;width: 100%;max-width: 1000px;border-collapse: collapse;\">\n<thead>\n  <tr style=\"height: 48px;border-bottom: none;\">\n    {{#each table_config.header.columns as |column|}}\n      {{#if column.enable}}\n        <th style=\"{{makeStyle @root.table_config.header.style}};{{makeStyle column.style}};\">{{column._label}}</th>\n      {{/if}}\n    {{/each}}\n  </tr>\n</thead>\n<tbody style=\"vertical-align: baseline  !important;font-weight: 400;font-size: 12px;position: relative;\">\n  <!-- Start rendering products -->\n  {{#each order.products as |product|}}\n    {{#if @last}}\n      <tr style=\"height: 48px;font-size:14px;border-bottom: none;\">\n    {{else}}\n      <tr style=\"height: 48px;font-size:14px;\">\n    {{/if}}\n      {{#each @root.table_config.header.columns as |column|}}\n        {{#if column.enable}}\n          {{#if (eq column.id 'item')}}\n            <!-- Item -->\n            {{#if product.is_composite_component}}\n              <td style=\"{{makeStyle @root.table_config.body.product_name.style}}; padding-left:20px;\">\n                {{#if @root.table_config.body.position.enable}}\n                  {{product._position}}\n                {{/if}}\n                {{product.price.description}}\n                <br>\n                {{#if @root.table_config.body.price_long_description.enable}}\n                  <span style=\"{{makeStyle @root.table_config.body.price_long_description.style}}\">{{product.price.long_description}}</span>\n                {{/if}}\n            {{else}}\n              <td style=\"{{makeStyle @root.table_config.body.product_name.style}}; font-weight: bold;\">\n                {{#if @root.table_config.body.product_name.enable}}\n                  {{#if @root.table_config.body.position.enable}}\n                    {{product._position}}\n                  {{/if}}\n                  {{product.name}}\n                  {{#if product.name}}\n                    <br>\n                  {{/if}}\n                {{/if}}\n                {{#if @root.table_config.body.product_description.enable}}\n                  {{#unless (eq product.price.description product.description)}}\n                    {{#if product.name}}\n                      <span style=\"{{makeStyle @root.table_config.body.product_description.style}}\">{{product.description}}</span>\n                    {{else}}\n                      {{product.description}}\n                    {{/if}}\n                    <br>\n                  {{/unless}}\n                {{/if}}\n                {{#if @root.table_config.body.price_description.enable}}\n                  <span style=\"{{makeStyle @root.table_config.body.price_description.style}}\">{{product.price.description}}</span>\n                  {{#if product.price.description}}\n                    <br>\n                  {{/if}}\n                {{/if}}\n                {{#if @root.table_config.body.price_long_description.enable}}\n                  <span style=\"{{makeStyle @root.table_config.body.price_long_description.style}}\">{{product.price.long_description}}</span>\n                {{/if}}\n\n            {{/if}}\n            </td>\n          {{/if}}\n          {{#if (eq column.id 'quantity')}}\n            <!-- Quantity -->\n            {{#if product.is_composite_component}}\n              <td style=\"{{makeStyle @root.table_config.body.quantity.style}}\">\n                {{product.price.quantity}}\n                <br>\n                <span style=\"{{makeStyle @root.table_config.body.payment_type.style}}\">{{product.price.quantity_billing_period}}</span>\n            {{else}}\n              <td style=\"{{makeStyle @root.table_config.body.quantity.style}}; font-weight: bold;\">\n                {{product.price.quantity}}\n                <br>\n                <span style=\"{{makeStyle @root.table_config.body.payment_type.style}}; font-weight: bold;\">{{product.price.quantity_billing_period}}</span>\n            {{/if}}\n              </td>\n          {{/if}}\n          {{#if (eq column.id 'tax')}}\n            <!-- Tax -->\n            {{#if product.is_composite_component}}\n              <td style=\"{{makeStyle @root.table_config.body.tax.style}};\">\n            {{else}}\n              <td style=\"{{makeStyle @root.table_config.body.tax.style}}; font-weight: bold;\">\n            {{/if}}\n                {{#if product.is_composite_component}}\n                  {{#if @root.table_config.body.composite_component_price.enable}}\n                    {{product.price.tax_rate}}\n                  {{/if}}\n                {{else}}\n                  {{#unless product.is_composite_price}}\n                    {{product.price.tax_rate}}\n                  {{/unless}}\n                {{/if}}\n              </td>\n          {{/if}}\n          {{#if (eq column.id 'unit_amount')}}\n            <!-- Unit amount -->\n            {{#if product.is_composite_component}}\n              <td style=\"{{makeStyle @root.table_config.body.unit_amount.style}};\">\n            {{else}}\n              <td style=\"{{makeStyle @root.table_config.body.unit_amount.style}}; font-weight: bold;\">\n            {{/if}}\n                {{#if product.is_composite_component}}\n                  {{#if @root.table_config.body.composite_component_price.enable}}\n                    {{product.price.unit_amount}}<span style=\"{{makeStyle @root.table_config.body.payment_type.style}}\">{{product.price.display_unit}}</span>\n                  {{/if}}\n                {{else}}\n                  {{product.price.unit_amount}}<span style=\"{{makeStyle @root.table_config.body.payment_type.style}}\">{{product.price.display_unit}}</span>\n                {{/if}}\n              </td>\n          {{/if}}\n          {{#if (eq column.id 'net_total')}}\n            <!-- Amount Subtotal -->\n            {{#if product.is_composite_component}}\n              <td style=\"{{makeStyle @root.table_config.body.net_total.style}};\">\n            {{else}}\n              <td style=\"{{makeStyle @root.table_config.body.net_total.style}}; font-weight: bold;\">\n            {{/if}}\n                {{#if product.is_composite_component}}\n                  {{#if @root.table_config.body.composite_component_price.enable}}\n                    {{product.price.amount_subtotal}}\n                  {{/if}}\n                {{else}}\n                  {{#unless product.is_composite_price}}\n                    {{product.price.amount_subtotal}}\n                  {{/unless}}\n                {{/if}}\n              </td>\n          {{/if}}\n          {{#if (eq column.id 'amount_tax')}}\n            <!-- Tax amount-->\n            {{#if product.is_composite_component}}\n              <td style=\"{{makeStyle @root.table_config.body.amount_tax.style}}\">\n            {{else}}\n              <td style=\"{{makeStyle @root.table_config.body.amount_tax.style}}; font-weight: bold;\">\n            {{/if}}\n                {{#if product.is_composite_component}}\n                  {{#if @root.table_config.body.composite_component_price.enable}}\n                    {{product.price.amount_tax}}\n                  {{/if}}\n                {{else}}\n                  {{#unless product.is_composite_price}}\n                    {{product.price.amount_tax}}\n                  {{/unless}}\n                {{/if}}\n              </td>\n          {{/if}}\n          {{#if (eq column.id 'gross_total')}}\n            <!-- Gross total -->\n            {{#if product.is_composite_component}}\n              <td style=\"{{makeStyle @root.table_config.body.gross_total.style}}\">\n            {{else}}\n              <td style=\"{{makeStyle @root.table_config.body.gross_total.style}}; font-weight: bold;;\">\n            {{/if}}\n                {{#if product.is_composite_price}}\n                  {{#each product.total_details.breakdown.recurrences as |item|}}\n                    {{item.amount_total}}\n                    {{#if @root.table_config.body.payment_type.enable}}\n                      {{#if (eq item.type 'recurring')}}\n                        <br>\n                        <span style=\"{{makeStyle @root.table_config.body.payment_type.style}}\">{{item.billing_period}}</span>\n                      {{/if}}\n                    {{/if}}\n                    <br>\n                  {{/each}}\n                {{else}}\n                  {{#if product.is_composite_component}}\n                    {{#if @root.table_config.body.composite_component_price.enable}}\n                      {{product.price.amount_total}}\n                      {{#if @root.table_config.body.payment_type.enable}}\n                        {{#if (eq product.price.type 'recurring')}}\n                          <br>\n                          <span style=\"{{makeStyle @root.table_config.body.payment_type.style}};\">{{product.price.billing_period}}</span>\n                        {{/if}}\n                      {{/if}}\n                    {{/if}}\n                  {{else}}\n                    {{product.price.amount_total}}\n                    {{#if @root.table_config.body.payment_type.enable}}\n                      {{#if (eq product.price.type 'recurring')}}\n                        <br>\n                        <span style=\"{{makeStyle @root.table_config.body.payment_type.style}}; font-weight: bold;\">{{product.price.billing_period}}</span>\n                      {{/if}}\n                    {{/if}}\n                  {{/if}}\n                {{/if}}\n              </td>\n          {{/if}}\n        {{/if}}\n      {{/each}}\n      </tr>\n  {{/each}}\n</tbody>\n</table>\n<table style=\"table-layout: fixed;width: 100%;max-width: 1000px;border-collapse: collapse;\">\n<thead>\n  <tr>\n  </tr>\n</thead>\n<tbody style=\"vertical-align: baseline  !important;font-weight: 400;font-size: 12px;position: relative;\">\n  <!-- Finish rendering products -->\n    <tr style=\"font-size: 16px;\">\n    <td></td>\n    <td></td>\n    <td></td>\n    <td></td>\n    <td style=\"{{makeStyle @root.table_config.footer.payment_type.style}} ; font-weight: bold; padding-top: 5px\">Gesamt Netto</td>\n    <td style=\"{{makeStyle @root.table_config.footer.net_total.style}}; font-weight: bold; padding-top: 5px\">{{order.total_details.one_time.amount_subtotal}}</td>\n  </tr>  \n  <tr style=\"font-size: 16px;\">\n    <td></td>\n    <td></td>\n    <td></td>\n    <td></td>\n    <td style=\"{{makeStyle @root.table_config.footer.payment_type.style}} ; font-weight: normal; padding-top: 5px\">19% MwSt.</td>\n    <td style=\"{{makeStyle @root.table_config.footer.gross_total.style}} ; font-weight: normal; padding-top: 5px\">{{order.total_details.amount_tax}}</td>\n  </tr>\n  <tr style=\"font-size: 16px;\">\n    <td></td>\n    <td></td>\n    <td></td>\n    <td></td>\n    <td style=\"{{makeStyle @root.table_config.footer.payment_type.style}} ; font-weight: bold; padding-top: 5px\">Gesamt Brutto</td>\n    <td style=\"{{makeStyle @root.table_config.footer.gross_total.style}} ; font-weight: bold; padding-top: 5px\">{{order.total_details.one_time.amount_total}}</td>\n  </tr>\n  <tr style=\"height:16px !important;\"></tr>\n  </tbody>\n  </table>\n"
  type          = "order_table"
}
