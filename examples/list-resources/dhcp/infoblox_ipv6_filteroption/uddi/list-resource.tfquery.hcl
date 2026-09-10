// List specific Ipv6 Filteroptions using filters
list "infoblox_ipv6_filteroption" "list_ipv6_filteroption_using_filters" {
  provider = infoblox
  config {
    filters = {
      comment = "Created by Terraform"
    }
  }
  limit = 10
}

// List specific Ipv6 Filteroptions using Tags
list "infoblox_ipv6_filteroption" "list_ipv6_filteroption_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List Ipv6 Filteroptions with resource details included
list "infoblox_ipv6_filteroption" "list_ipv6_filteroption_with_resource" {
  provider         = infoblox
  include_resource = true
}
