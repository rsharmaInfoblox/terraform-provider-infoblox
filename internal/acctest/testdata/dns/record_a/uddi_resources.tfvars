# Auto-generated resource acceptance-test cases for RecordA (uddi).
case "basic" {
  # basic — generated from terraform-provider-uddi
  backend = "uddi"
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_zone_auth" "test" {
  #   uddi = {
  #     fqdn = "{{random}}.com."
  #     primary_type = "cloud"
  #   }
  # }
  # PREREQ

  step {
    uddi {
      rdata = { address = "{{random_ip}}" }
      zone  = "dns/auth_zone/113e8a4d-440c-488f-aaf0-1acea9437ff9"
    }
    check = {
      "uddi.rdata.address" = "{{random_ip}}"
    }
  }

}

case "disappears" {
  # disappears — generated from terraform-provider-uddi
  backend               = "uddi"
  disappears            = true
  expect_non_empty_plan = true
  skip                  = true
  skip_reason           = "Test Skipped due to inconsistent error codes returned by the API [NORTHSTAR-12575]"
  # prerequisites_hcl     = <<-PREREQ
  # resource "infoblox_zone_auth" "test" {
  #   uddi = {
  #     fqdn = "{{random}}.com."
  #     primary_type = "cloud"
  #   }
  # }
  # PREREQ

  step {
    uddi {
      rdata = { address = "{{random_ip}}" }
      zone  = "dns/auth_zone/113e8a4d-440c-488f-aaf0-1acea9437ff9"
    }
  }

}

case "comment" {
  # comment — generated from terraform-provider-uddi
  backend = "uddi"
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_zone_auth" "test" {
  #   uddi = {
  #     fqdn = "{{random}}.com."
  #     primary_type = "cloud"
  #   }
  # }
  # PREREQ

  step {
    uddi {
      rdata   = { address = "{{random_ip}}" }
      zone    = "dns/auth_zone/113e8a4d-440c-488f-aaf0-1acea9437ff9"
      comment = "some comment"
    }
    check = {
      "uddi.comment" = "some comment"
    }
  }

  step {
    uddi {
      rdata   = { address = "{{random_ip}}" }
      zone    = "dns/auth_zone/113e8a4d-440c-488f-aaf0-1acea9437ff9"
      comment = "updated comment"
    }
    check = {
      "uddi.comment" = "updated comment"
    }
  }

}

case "disabled" {
  # disabled — generated from terraform-provider-uddi
  backend = "uddi"
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_zone_auth" "test" {
  #   uddi = {
  #     fqdn = "{{random}}.com."
  #     primary_type = "cloud"
  #   }
  # }
  # PREREQ

  step {
    uddi {
      rdata    = { address = "{{random_ip}}" }
      zone     = "dns/auth_zone/113e8a4d-440c-488f-aaf0-1acea9437ff9"
      disabled = true
    }
    check = {
      "uddi.disabled" = "true"
    }
  }

  step {
    uddi {
      rdata    = { address = "{{random_ip}}" }
      zone     = "dns/auth_zone/113e8a4d-440c-488f-aaf0-1acea9437ff9"
      disabled = false
    }
    check = {
      "uddi.disabled" = "false"
    }
  }

}

case "inheritance_sources" {
  # inheritance_sources — generated from terraform-provider-uddi
  backend = "uddi"
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_zone_auth" "test" {
  #   uddi = {
  #     fqdn = "{{random}}.com."
  #     primary_type = "cloud"
  #   }
  # }
  # PREREQ

  step {
    uddi {
      rdata               = { address = "{{random_ip}}" }
      zone                = "dns/auth_zone/113e8a4d-440c-488f-aaf0-1acea9437ff9"
      inheritance_sources = { ttl = { action = "inherit" } }
    }
    check = {
      "uddi.inheritance_sources.ttl.action" = "inherit"
    }
  }

  step {
    uddi {
      rdata               = { address = "{{random_ip}}" }
      zone                = "dns/auth_zone/113e8a4d-440c-488f-aaf0-1acea9437ff9"
      inheritance_sources = { ttl = { action = "override" } }
    }
    check = {
      "uddi.inheritance_sources.ttl.action" = "override"
    }
  }

}

case "name_in_zone" {
  # name_in_zone — generated from terraform-provider-uddi
  backend = "uddi"
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_zone_auth" "test" {
  #   uddi = {
  #     fqdn = "{{random}}.com."
  #     primary_type = "cloud"
  #   }
  # }
  # PREREQ

  step {
    uddi {
      rdata        = { address = "{{random_ip}}" }
      zone         = "dns/auth_zone/113e8a4d-440c-488f-aaf0-1acea9437ff9"
      name_in_zone = "xyz"
    }
    check = {
      "uddi.name_in_zone" = "xyz"
    }
  }

  step {
    uddi {
      rdata        = { address = "{{random_ip}}" }
      zone         = "dns/auth_zone/113e8a4d-440c-488f-aaf0-1acea9437ff9"
      name_in_zone = "abc"
    }
    check = {
      "uddi.name_in_zone" = "abc"
    }
  }

}

case "rdata" {
  # rdata — generated from terraform-provider-uddi
  backend = "uddi"
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_zone_auth" "test" {
  #   uddi = {
  #     fqdn = "{{random}}.com."
  #     primary_type = "cloud"
  #   }
  # }
  # PREREQ

  step {
    uddi {
      rdata = { address = "{{random_ip}}" }
      zone  = "dns/auth_zone/113e8a4d-440c-488f-aaf0-1acea9437ff9"
    }
    check = {
      "uddi.rdata.address" = "{{random_ip}}"
    }
  }

  step {
    uddi {
      rdata = { address = "{{random_ip2}}" }
      zone  = "dns/auth_zone/113e8a4d-440c-488f-aaf0-1acea9437ff9"
    }
    check = {
      "uddi.rdata.address" = "{{random_ip2}}"
    }
  }

}

case "tags" {
  # tags — generated from terraform-provider-uddi
  backend = "uddi"
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_zone_auth" "test" {
  #   uddi = {
  #     fqdn = "{{random}}.com."
  #     primary_type = "cloud"
  #   }
  # }
  # PREREQ

  step {
    uddi {
      rdata = { address = "{{random_ip}}" }
      zone  = "dns/auth_zone/113e8a4d-440c-488f-aaf0-1acea9437ff9"
      tags  = { tag1 = "value1" }
    }
    check = {
      "uddi.tags.tag1" = "value1"
    }
  }

  step {
    uddi {
      rdata = { address = "{{random_ip}}" }
      zone  = "dns/auth_zone/113e8a4d-440c-488f-aaf0-1acea9437ff9"
      tags  = { tag1 = "value2" }
    }
    check = {
      "uddi.tags.tag1" = "value2"
    }
  }

}

case "ttl" {
  # ttl — generated from terraform-provider-uddi
  backend = "uddi"
  # prerequisites_hcl = <<-PREREQ
  # resource "infoblox_zone_auth" "test" {
  #   uddi = {
  #     fqdn = "{{random}}.com."
  #     primary_type = "cloud"
  #   }
  # }
  # PREREQ

  step {
    uddi {
      rdata = { address = "{{random_ip}}" }
      zone  = "dns/auth_zone/113e8a4d-440c-488f-aaf0-1acea9437ff9"
      ttl   = 60
    }
    check = {
      "uddi.ttl" = "60"
    }
  }

  step {
    uddi {
      rdata = { address = "{{random_ip}}" }
      zone  = "dns/auth_zone/113e8a4d-440c-488f-aaf0-1acea9437ff9"
      ttl   = 90
    }
    check = {
      "uddi.ttl" = "90"
    }
  }

}

case "view" {
  backend = "uddi"
  step {
    uddi {
      rdata              = { address = "{{random_ip}}" }
      absolute_name_spec = "10.in-addr.arpa."
      view               = "dns/view/28b9c115-8d5f-416e-979f-e7e71d80a3a3"
    }
    check = {
      "uddi.view" = "dns/view/28b9c115-8d5f-416e-979f-e7e71d80a3a3"
    }
  }

  step {
    uddi {
      rdata              = { address = "{{random_ip}}" }
      absolute_name_spec = "zone-xoeivh.com."
      view               = "dns/view/ce528cc5-7482-4278-835f-801fb4f884fe"
    }
    check = {
      "uddi.view" = "dns/view/ce528cc5-7482-4278-835f-801fb4f884fe"
    }
  }

}

case "zone" {
  # zone — generated from terraform-provider-uddi
  backend = "uddi"
  step {
    uddi {
      rdata = { address = "{{random_ip}}" }
      zone  = "dns/auth_zone/c75d3700-05b5-4ff8-a413-dfa0bcb5b020"
    }
    check = {
      "uddi.zone" = "dns/auth_zone/c75d3700-05b5-4ff8-a413-dfa0bcb5b020"
    }
  }
  step {
    uddi {
      rdata = { address = "{{random_ip}}" }
      zone  = "dns/auth_zone/113e8a4d-440c-488f-aaf0-1acea9437ff9"
      ttl   = 90
    }
    check = {
      "uddi.zone" = "dns/auth_zone/113e8a4d-440c-488f-aaf0-1acea9437ff9"
    }
  }
}

case "options" {
  backend = "uddi"

  step {
    uddi {
      rdata   = { address = "{{random_ip}}" }
      zone    = "dns/auth_zone/113e8a4d-440c-488f-aaf0-1acea9437ff9"
      options = { create_ptr = true, check_rmz = true }
    }
    check = {
      "uddi.options.create_ptr" = "true"
      "uddi.options.check_rmz"  = "true"
    }
  }

  step {
    uddi {
      rdata   = { address = "{{random_ip}}" }
      zone    = "dns/auth_zone/113e8a4d-440c-488f-aaf0-1acea9437ff9"
      options = { create_ptr = true, check_rmz = false }
    }
    check = {
      "uddi.options.create_ptr" = "true"
      "uddi.options.check_rmz"  = "false"
    }
  }

  step {
    uddi {
      rdata   = { address = "{{random_ip}}" }
      zone    = "dns/auth_zone/113e8a4d-440c-488f-aaf0-1acea9437ff9"
      options = { create_ptr = false, check_rmz = false }
    }
    check = {
      "uddi.options.create_ptr" = "false"
      "uddi.options.check_rmz"  = "false"
    }
  }

}
