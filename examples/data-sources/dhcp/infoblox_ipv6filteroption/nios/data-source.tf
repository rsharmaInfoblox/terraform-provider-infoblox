// Retrieve a specific DHCP IPv6 Filter Option by filters
data "infoblox_ipv6filteroption" "get_dhcp_ipv6filteroptions_using_filters" {
  filters = {
    name = "example_ipv6_filter_option_1"
  }
}

// Retrieve specific DHCP IPv6 Option Filters using Extensible Attributes
data "infoblox_ipv6filteroption" "get_dhcp_ipv6filteroption_using_extensible_attributes" {
  ext_attr_filters = {
    Site = "location-1"
  }
}

// Retrieve all DHCP IPv6 Option Filters
data "infoblox_ipv6filteroption" "get_all_dhcp_ipv6filteroptions" {}
