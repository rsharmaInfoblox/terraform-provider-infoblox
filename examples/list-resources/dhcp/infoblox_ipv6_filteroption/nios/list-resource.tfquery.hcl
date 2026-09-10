// List specific DHCP IPv6 Option Filters using filters
list "infoblox_ipv6_filteroption" "list_ipv6_filteroptions_using_filters" {
  provider = infoblox
  config {
    filters = {
      name = "example_ipv6_filter_option_1"
    }
  }
}

// List specific DHCP IPv6 Option Filters using Extensible Attributes
list "infoblox_ipv6_filteroption" "list_ipv6_filteroptions_using_extensible_attributes" {
  provider = infoblox
  config {
    ext_attr_filters = {
      Site = "location-1"
    }
  }
}

// List DHCP IPv6 Option Filters with resource details included
list "infoblox_ipv6_filteroption" "list_ipv6_filteroptions_with_resource" {
  provider         = infoblox
  include_resource = true
}
