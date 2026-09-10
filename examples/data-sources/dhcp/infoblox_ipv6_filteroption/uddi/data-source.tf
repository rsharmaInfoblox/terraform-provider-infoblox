// Retrieve a specific DHCP IPv6 Option Filter by filters
data "infoblox_ipv6_filteroption" "get_ipv6_filteroption_using_filters" {
  filters = {
    name = "ipv6_filteroption_example"
  }
}

// Retrieve specific DHCP IPv6 Option Filters using Tags
data "infoblox_ipv6_filteroption" "get_ipv6_filteroption_using_tag_filters" {
  tag_filters = {
    location = "site1"
  }
}

// Retrieve all DHCP IPv6 Option Filters
data "infoblox_ipv6_filteroption" "get_all_ipv6_filteroptions" {}
