// List specific Ipv6Filteroptions using filters
list "infoblox_ipv6filteroption" "list_ipv6filteroption_using_filters" {
  provider = infoblox
  config {
    filters = {
      comment = "Created by Terraform"
    }
  }
  limit = 10
}

// List specific Ipv6Filteroptions using Tags
list "infoblox_ipv6filteroption" "list_ipv6filteroption_using_tags" {
  provider = infoblox
  config {
    tag_filters = {
      Site = "location-1"
    }
  }
}

// List Ipv6Filteroptions with resource details included
list "infoblox_ipv6filteroption" "list_ipv6filteroption_with_resource" {
  provider         = infoblox
  include_resource = true
}
