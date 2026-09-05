package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// RecordHostIpv6addrDiscoveredDataModel is the Terraform model for RecordHostIpv6addrDiscoveredData
type RecordHostIpv6addrDiscoveredDataModel struct {
	DeviceModel                     types.String `tfsdk:"device_model"`
	DevicePortName                  types.String `tfsdk:"device_port_name"`
	DevicePortType                  types.String `tfsdk:"device_port_type"`
	DeviceType                      types.String `tfsdk:"device_type"`
	DeviceVendor                    types.String `tfsdk:"device_vendor"`
	DiscoveredName                  types.String `tfsdk:"discovered_name"`
	Discoverer                      types.String `tfsdk:"discoverer"`
	Duid                            types.String `tfsdk:"duid"`
	FirstDiscovered                 types.Int64  `tfsdk:"first_discovered"`
	IprgNo                          types.Int64  `tfsdk:"iprg_no"`
	IprgState                       types.String `tfsdk:"iprg_state"`
	IprgType                        types.String `tfsdk:"iprg_type"`
	LastDiscovered                  types.Int64  `tfsdk:"last_discovered"`
	MacAddress                      types.String `tfsdk:"mac_address"`
	MgmtIpAddress                   types.String `tfsdk:"mgmt_ip_address"`
	NetbiosName                     types.String `tfsdk:"netbios_name"`
	NetworkComponentDescription     types.String `tfsdk:"network_component_description"`
	NetworkComponentIp              types.String `tfsdk:"network_component_ip"`
	NetworkComponentModel           types.String `tfsdk:"network_component_model"`
	NetworkComponentName            types.String `tfsdk:"network_component_name"`
	NetworkComponentPortDescription types.String `tfsdk:"network_component_port_description"`
	NetworkComponentPortName        types.String `tfsdk:"network_component_port_name"`
	NetworkComponentPortNumber      types.String `tfsdk:"network_component_port_number"`
	NetworkComponentType            types.String `tfsdk:"network_component_type"`
	NetworkComponentVendor          types.String `tfsdk:"network_component_vendor"`
	OpenPorts                       types.String `tfsdk:"open_ports"`
	Os                              types.String `tfsdk:"os"`
	PortDuplex                      types.String `tfsdk:"port_duplex"`
	PortLinkStatus                  types.String `tfsdk:"port_link_status"`
	PortSpeed                       types.String `tfsdk:"port_speed"`
	PortStatus                      types.String `tfsdk:"port_status"`
	PortType                        types.String `tfsdk:"port_type"`
	PortVlanDescription             types.String `tfsdk:"port_vlan_description"`
	PortVlanName                    types.String `tfsdk:"port_vlan_name"`
	PortVlanNumber                  types.String `tfsdk:"port_vlan_number"`
	VAdapter                        types.String `tfsdk:"v_adapter"`
	VCluster                        types.String `tfsdk:"v_cluster"`
	VDatacenter                     types.String `tfsdk:"v_datacenter"`
	VEntityName                     types.String `tfsdk:"v_entity_name"`
	VEntityType                     types.String `tfsdk:"v_entity_type"`
	VHost                           types.String `tfsdk:"v_host"`
	VSwitch                         types.String `tfsdk:"v_switch"`
	VmiName                         types.String `tfsdk:"vmi_name"`
	VmiId                           types.String `tfsdk:"vmi_id"`
	VlanPortGroup                   types.String `tfsdk:"vlan_port_group"`
	VswitchName                     types.String `tfsdk:"vswitch_name"`
	VswitchId                       types.String `tfsdk:"vswitch_id"`
	VswitchType                     types.String `tfsdk:"vswitch_type"`
	VswitchIpv6Enabled              types.Bool   `tfsdk:"vswitch_ipv6_enabled"`
	VportName                       types.String `tfsdk:"vport_name"`
	VportMacAddress                 types.String `tfsdk:"vport_mac_address"`
	VportLinkStatus                 types.String `tfsdk:"vport_link_status"`
	VportConfSpeed                  types.String `tfsdk:"vport_conf_speed"`
	VportConfMode                   types.String `tfsdk:"vport_conf_mode"`
	VportSpeed                      types.String `tfsdk:"vport_speed"`
	VportMode                       types.String `tfsdk:"vport_mode"`
	VswitchSegmentType              types.String `tfsdk:"vswitch_segment_type"`
	VswitchSegmentName              types.String `tfsdk:"vswitch_segment_name"`
	VswitchSegmentId                types.String `tfsdk:"vswitch_segment_id"`
	VswitchSegmentPortGroup         types.String `tfsdk:"vswitch_segment_port_group"`
	VswitchAvailablePortsCount      types.Int64  `tfsdk:"vswitch_available_ports_count"`
	VswitchTepType                  types.String `tfsdk:"vswitch_tep_type"`
	VswitchTepIp                    types.String `tfsdk:"vswitch_tep_ip"`
	VswitchTepPortGroup             types.String `tfsdk:"vswitch_tep_port_group"`
	VswitchTepVlan                  types.String `tfsdk:"vswitch_tep_vlan"`
	VswitchTepDhcpServer            types.String `tfsdk:"vswitch_tep_dhcp_server"`
	VswitchTepMulticast             types.String `tfsdk:"vswitch_tep_multicast"`
	VmhostIpAddress                 types.String `tfsdk:"vmhost_ip_address"`
	VmhostName                      types.String `tfsdk:"vmhost_name"`
	VmhostMacAddress                types.String `tfsdk:"vmhost_mac_address"`
	VmhostSubnetCidr                types.Int64  `tfsdk:"vmhost_subnet_cidr"`
	VmhostNicNames                  types.String `tfsdk:"vmhost_nic_names"`
	VmiTenantId                     types.String `tfsdk:"vmi_tenant_id"`
	CmpType                         types.String `tfsdk:"cmp_type"`
	VmiIpType                       types.String `tfsdk:"vmi_ip_type"`
	VmiPrivateAddress               types.String `tfsdk:"vmi_private_address"`
	VmiIsPublicAddress              types.Bool   `tfsdk:"vmi_is_public_address"`
	CiscoIseSsid                    types.String `tfsdk:"cisco_ise_ssid"`
	CiscoIseEndpointProfile         types.String `tfsdk:"cisco_ise_endpoint_profile"`
	CiscoIseSessionState            types.String `tfsdk:"cisco_ise_session_state"`
	CiscoIseSecurityGroup           types.String `tfsdk:"cisco_ise_security_group"`
	TaskName                        types.String `tfsdk:"task_name"`
	NetworkComponentLocation        types.String `tfsdk:"network_component_location"`
	NetworkComponentContact         types.String `tfsdk:"network_component_contact"`
	DeviceLocation                  types.String `tfsdk:"device_location"`
	DeviceContact                   types.String `tfsdk:"device_contact"`
	ApName                          types.String `tfsdk:"ap_name"`
	ApIpAddress                     types.String `tfsdk:"ap_ip_address"`
	ApSsid                          types.String `tfsdk:"ap_ssid"`
	BridgeDomain                    types.String `tfsdk:"bridge_domain"`
	EndpointGroups                  types.String `tfsdk:"endpoint_groups"`
	Tenant                          types.String `tfsdk:"tenant"`
	VrfName                         types.String `tfsdk:"vrf_name"`
	VrfDescription                  types.String `tfsdk:"vrf_description"`
	VrfRd                           types.String `tfsdk:"vrf_rd"`
	BgpAs                           types.Int64  `tfsdk:"bgp_as"`
}

// RecordHostIpv6addrDiscoveredDataAttrTypes contains the attribute types for RecordHostIpv6addrDiscoveredDataModel
var RecordHostIpv6addrDiscoveredDataAttrTypes = map[string]attr.Type{
	"device_model":                       types.StringType,
	"device_port_name":                   types.StringType,
	"device_port_type":                   types.StringType,
	"device_type":                        types.StringType,
	"device_vendor":                      types.StringType,
	"discovered_name":                    types.StringType,
	"discoverer":                         types.StringType,
	"duid":                               types.StringType,
	"first_discovered":                   types.Int64Type,
	"iprg_no":                            types.Int64Type,
	"iprg_state":                         types.StringType,
	"iprg_type":                          types.StringType,
	"last_discovered":                    types.Int64Type,
	"mac_address":                        types.StringType,
	"mgmt_ip_address":                    types.StringType,
	"netbios_name":                       types.StringType,
	"network_component_description":      types.StringType,
	"network_component_ip":               types.StringType,
	"network_component_model":            types.StringType,
	"network_component_name":             types.StringType,
	"network_component_port_description": types.StringType,
	"network_component_port_name":        types.StringType,
	"network_component_port_number":      types.StringType,
	"network_component_type":             types.StringType,
	"network_component_vendor":           types.StringType,
	"open_ports":                         types.StringType,
	"os":                                 types.StringType,
	"port_duplex":                        types.StringType,
	"port_link_status":                   types.StringType,
	"port_speed":                         types.StringType,
	"port_status":                        types.StringType,
	"port_type":                          types.StringType,
	"port_vlan_description":              types.StringType,
	"port_vlan_name":                     types.StringType,
	"port_vlan_number":                   types.StringType,
	"v_adapter":                          types.StringType,
	"v_cluster":                          types.StringType,
	"v_datacenter":                       types.StringType,
	"v_entity_name":                      types.StringType,
	"v_entity_type":                      types.StringType,
	"v_host":                             types.StringType,
	"v_switch":                           types.StringType,
	"vmi_name":                           types.StringType,
	"vmi_id":                             types.StringType,
	"vlan_port_group":                    types.StringType,
	"vswitch_name":                       types.StringType,
	"vswitch_id":                         types.StringType,
	"vswitch_type":                       types.StringType,
	"vswitch_ipv6_enabled":               types.BoolType,
	"vport_name":                         types.StringType,
	"vport_mac_address":                  types.StringType,
	"vport_link_status":                  types.StringType,
	"vport_conf_speed":                   types.StringType,
	"vport_conf_mode":                    types.StringType,
	"vport_speed":                        types.StringType,
	"vport_mode":                         types.StringType,
	"vswitch_segment_type":               types.StringType,
	"vswitch_segment_name":               types.StringType,
	"vswitch_segment_id":                 types.StringType,
	"vswitch_segment_port_group":         types.StringType,
	"vswitch_available_ports_count":      types.Int64Type,
	"vswitch_tep_type":                   types.StringType,
	"vswitch_tep_ip":                     types.StringType,
	"vswitch_tep_port_group":             types.StringType,
	"vswitch_tep_vlan":                   types.StringType,
	"vswitch_tep_dhcp_server":            types.StringType,
	"vswitch_tep_multicast":              types.StringType,
	"vmhost_ip_address":                  types.StringType,
	"vmhost_name":                        types.StringType,
	"vmhost_mac_address":                 types.StringType,
	"vmhost_subnet_cidr":                 types.Int64Type,
	"vmhost_nic_names":                   types.StringType,
	"vmi_tenant_id":                      types.StringType,
	"cmp_type":                           types.StringType,
	"vmi_ip_type":                        types.StringType,
	"vmi_private_address":                types.StringType,
	"vmi_is_public_address":              types.BoolType,
	"cisco_ise_ssid":                     types.StringType,
	"cisco_ise_endpoint_profile":         types.StringType,
	"cisco_ise_session_state":            types.StringType,
	"cisco_ise_security_group":           types.StringType,
	"task_name":                          types.StringType,
	"network_component_location":         types.StringType,
	"network_component_contact":          types.StringType,
	"device_location":                    types.StringType,
	"device_contact":                     types.StringType,
	"ap_name":                            types.StringType,
	"ap_ip_address":                      types.StringType,
	"ap_ssid":                            types.StringType,
	"bridge_domain":                      types.StringType,
	"endpoint_groups":                    types.StringType,
	"tenant":                             types.StringType,
	"vrf_name":                           types.StringType,
	"vrf_description":                    types.StringType,
	"vrf_rd":                             types.StringType,
	"bgp_as":                             types.Int64Type,
}

// RecordHostIpv6addrDiscoveredDataResourceSchemaAttributes contains the schema attributes for RecordHostIpv6addrDiscoveredDataModel
var RecordHostIpv6addrDiscoveredDataResourceSchemaAttributes = map[string]schema.Attribute{
	"device_model": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The model name of the end device in the vendor terminology.",
	},
	"device_port_name": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The system name of the interface associated with the discovered IP address.",
	},
	"device_port_type": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The hardware type of the interface associated with the discovered IP address.",
	},
	"device_type": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The type of end host in vendor terminology.",
	},
	"device_vendor": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The vendor name of the end host.",
	},
	"discovered_name": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The name of the network device associated with the discovered IP address.",
	},
	"discoverer": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Specifies whether the IP address was discovered by a NetMRI or NIOS discovery process.",
	},
	"duid": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "For IPv6 address only. The DHCP unique identifier of the discovered host. This is an optional field, and data might not be included.",
	},
	"first_discovered": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The date and time the IP address was first discovered in Epoch seconds format.",
	},
	"iprg_no": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The port redundant group number.",
	},
	"iprg_state": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("VIP", "ACTIVE", "STANDBY", "NEGOTIATION"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The status for the IP address within port redundant group.",
	},
	"iprg_type": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("HSRP", "VRRP"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The port redundant group type.",
	},
	"last_discovered": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The date and time the IP address was last discovered in Epoch seconds format.",
	},
	"mac_address": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The discovered MAC address for the host. This is the unique identifier of a network device. The discovery acquires the MAC address for hosts that are located on the same network as the Grid member that is running the discovery. This can also be the MAC address of a virtual entity on a specified vSphere server.",
	},
	"mgmt_ip_address": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The management IP address of the end host that has more than one IP.",
	},
	"netbios_name": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The name returned in the NetBIOS reply or the name you manually register for the discovered host.",
	},
	"network_component_description": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "A textual description of the switch that is connected to the end device.",
	},
	"network_component_ip": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The IPv4 Address or IPv6 Address of the switch that is connected to the end device.",
	},
	"network_component_model": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Model name of the switch port connected to the end host in vendor terminology.",
	},
	"network_component_name": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "If a reverse lookup was successful for the IP address associated with this switch, the host name is displayed in this field.",
	},
	"network_component_port_description": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "A textual description of the switch port that is connected to the end device.",
	},
	"network_component_port_name": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The name of the switch port connected to the end device.",
	},
	"network_component_port_number": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The number of the switch port connected to the end device.",
	},
	"network_component_type": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Identifies the switch that is connected to the end device.",
	},
	"network_component_vendor": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The vendor name of the switch port connected to the end host.",
	},
	"open_ports": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The list of opened ports on the IP address, represented as: \"TCP: 21,22,23 UDP: 137,139\". Limited to max total 1000 ports.",
	},
	"os": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The operating system of the detected host or virtual entity. The OS can be one of the following: * Microsoft for all discovered hosts that have a non-null value in the MAC addresses using the NetBIOS discovery method. * A value that a TCP discovery returns. * The OS of a virtual entity on a vSphere server.",
	},
	"port_duplex": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The negotiated or operational duplex setting of the switch port connected to the end device.",
	},
	"port_link_status": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The link status of the switch port connected to the end device. Indicates whether it is connected.",
	},
	"port_speed": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The interface speed, in Mbps, of the switch port.",
	},
	"port_status": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The operational status of the switch port. Indicates whether the port is up or down.",
	},
	"port_type": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The type of switch port.",
	},
	"port_vlan_description": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The description of the VLAN of the switch port that is connected to the end device.",
	},
	"port_vlan_name": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The name of the VLAN of the switch port.",
	},
	"port_vlan_number": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The ID of the VLAN of the switch port.",
	},
	"v_adapter": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The name of the physical network adapter through which the virtual entity is connected to the appliance.",
	},
	"v_cluster": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The name of the VMware cluster to which the virtual entity belongs.",
	},
	"v_datacenter": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The name of the vSphere datacenter or container to which the virtual entity belongs.",
	},
	"v_entity_name": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The name of the virtual entity.",
	},
	"v_entity_type": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The virtual entity type. This can be blank or one of the following: Virtual Machine, Virtual Host, or Virtual Center. Virtual Center represents a VMware vCenter server.",
	},
	"v_host": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The name of the VMware server on which the virtual entity was discovered.",
	},
	"v_switch": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The name of the switch to which the virtual entity is connected.",
	},
	"vmi_name": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Name of the virtual machine.",
	},
	"vmi_id": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "ID of the virtual machine.",
	},
	"vlan_port_group": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Port group which the virtual machine belongs to.",
	},
	"vswitch_name": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Name of the virtual switch.",
	},
	"vswitch_id": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "ID of the virtual switch.",
	},
	"vswitch_type": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("Unknown", "Standard", "Distributed"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Type of the virtual switch: standard or distributed.",
	},
	"vswitch_ipv6_enabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Indicates the virtual switch has IPV6 enabled.",
	},
	"vport_name": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Name of the network adapter on the virtual switch connected with the virtual machine.",
	},
	"vport_mac_address": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "MAC address of the network adapter on the virtual switch where the virtual machine connected to.",
	},
	"vport_link_status": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Link status of the network adapter on the virtual switch where the virtual machine connected to.",
	},
	"vport_conf_speed": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Configured speed of the network adapter on the virtual switch where the virtual machine connected to. Unit is kb.",
	},
	"vport_conf_mode": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("Unknown", "Full-duplex", "Half-duplex"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Configured mode of the network adapter on the virtual switch where the virtual machine connected to.",
	},
	"vport_speed": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Actual speed of the network adapter on the virtual switch where the virtual machine connected to. Unit is kb.",
	},
	"vport_mode": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("Unknown", "Full-duplex", "Half-duplex"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Actual mode of the network adapter on the virtual switch where the virtual machine connected to.",
	},
	"vswitch_segment_type": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Type of the network segment on which the current virtual machine/vport connected to.",
	},
	"vswitch_segment_name": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Name of the network segment on which the current virtual machine/vport connected to.",
	},
	"vswitch_segment_id": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "ID of the network segment on which the current virtual machine/vport connected to.",
	},
	"vswitch_segment_port_group": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Port group of the network segment on which the current virtual machine/vport connected to.",
	},
	"vswitch_available_ports_count": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Numer of available ports reported by the virtual switch on which the virtual machine/vport connected to.",
	},
	"vswitch_tep_type": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Type of virtual tunnel endpoint (VTEP) in the virtual switch.",
	},
	"vswitch_tep_ip": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "IP address of the virtual tunnel endpoint (VTEP) in the virtual switch.",
	},
	"vswitch_tep_port_group": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Port group of the virtual tunnel endpoint (VTEP) in the virtual switch.",
	},
	"vswitch_tep_vlan": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "VLAN of the virtual tunnel endpoint (VTEP) in the virtual switch.",
	},
	"vswitch_tep_dhcp_server": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "DHCP server of the virtual tunnel endpoint (VTEP) in the virtual switch.",
	},
	"vswitch_tep_multicast": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Muticast address of the virtual tunnel endpoint (VTEP) in the virtual swtich.",
	},
	"vmhost_ip_address": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "IP address of the physical node on which the virtual machine is hosted.",
	},
	"vmhost_name": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Name of the physical node on which the virtual machine is hosted.",
	},
	"vmhost_mac_address": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "MAC address of the physical node on which the virtual machine is hosted.",
	},
	"vmhost_subnet_cidr": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "CIDR subnet of the physical node on which the virtual machine is hosted.",
	},
	"vmhost_nic_names": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "List of all physical port names used by the virtual switch on the physical node on which the virtual machine is hosted. Represented as: \"eth1,eth2,eth3\".",
	},
	"vmi_tenant_id": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "ID of the tenant which virtual machine belongs to.",
	},
	"cmp_type": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "If the IP is coming from a Cloud environment, the Cloud Management Platform type.",
	},
	"vmi_ip_type": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Discovered IP address type.",
	},
	"vmi_private_address": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Private IP address of the virtual machine.",
	},
	"vmi_is_public_address": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Indicates whether the IP address is a public address.",
	},
	"cisco_ise_ssid": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The Cisco ISE SSID.",
	},
	"cisco_ise_endpoint_profile": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The Endpoint Profile created in Cisco ISE.",
	},
	"cisco_ise_session_state": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("AUTHENTICATED", "AUTHENTICATING", "DISCONNECTED", "POSTURED", "STARTED"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The Cisco ISE connection session state.",
	},
	"cisco_ise_security_group": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The Cisco ISE security group name.",
	},
	"task_name": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The name of the discovery task.",
	},
	"network_component_location": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Location of the network component on which the IP address was discovered.",
	},
	"network_component_contact": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Contact information from the network component on which the IP address was discovered.",
	},
	"device_location": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Location of device on which the IP address was discovered.",
	},
	"device_contact": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Contact information from device on which the IP address was discovered.",
	},
	"ap_name": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Discovered name of Wireless Access Point.",
	},
	"ap_ip_address": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Discovered IP address of Wireless Access Point.",
	},
	"ap_ssid": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Service set identifier (SSID) associated with Wireless Access Point.",
	},
	"bridge_domain": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Discovered bridge domain.",
	},
	"endpoint_groups": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "A comma-separated list of the discovered endpoint groups.",
	},
	"tenant": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Discovered tenant.",
	},
	"vrf_name": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The name of the VRF.",
	},
	"vrf_description": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Description of the VRF.",
	},
	"vrf_rd": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Route distinguisher of the VRF.",
	},
	"bgp_as": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The BGP autonomous system number.",
	},
}

// ExpandRecordHostIpv6addrDiscoveredData converts a Terraform Object to SDK type
func ExpandRecordHostIpv6addrDiscoveredData(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdns.RecordHostIpv6addrDiscoveredData {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RecordHostIpv6addrDiscoveredDataModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *RecordHostIpv6addrDiscoveredDataModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdns.RecordHostIpv6addrDiscoveredData {
	if m == nil {
		return nil
	}
	to := &niosdns.RecordHostIpv6addrDiscoveredData{
		DeviceModel:                     flex.ExpandStringPointerNullAsEmpty(m.DeviceModel),
		DevicePortName:                  flex.ExpandStringPointerNullAsEmpty(m.DevicePortName),
		DevicePortType:                  flex.ExpandStringPointerNullAsEmpty(m.DevicePortType),
		DeviceType:                      flex.ExpandStringPointerNullAsEmpty(m.DeviceType),
		DeviceVendor:                    flex.ExpandStringPointerNullAsEmpty(m.DeviceVendor),
		DiscoveredName:                  flex.ExpandStringPointerNullAsEmpty(m.DiscoveredName),
		Discoverer:                      flex.ExpandStringPointerNullAsEmpty(m.Discoverer),
		Duid:                            flex.ExpandStringPointerNullAsEmpty(m.Duid),
		FirstDiscovered:                 flex.ExpandInt64Pointer(m.FirstDiscovered),
		IprgNo:                          flex.ExpandInt64Pointer(m.IprgNo),
		IprgState:                       flex.ExpandStringPointerNullAsEmpty(m.IprgState),
		IprgType:                        flex.ExpandStringPointerNullAsEmpty(m.IprgType),
		LastDiscovered:                  flex.ExpandInt64Pointer(m.LastDiscovered),
		MacAddress:                      flex.ExpandStringPointerNullAsEmpty(m.MacAddress),
		MgmtIpAddress:                   flex.ExpandStringPointerNullAsEmpty(m.MgmtIpAddress),
		NetbiosName:                     flex.ExpandStringPointerNullAsEmpty(m.NetbiosName),
		NetworkComponentDescription:     flex.ExpandStringPointerNullAsEmpty(m.NetworkComponentDescription),
		NetworkComponentIp:              flex.ExpandStringPointerNullAsEmpty(m.NetworkComponentIp),
		NetworkComponentModel:           flex.ExpandStringPointerNullAsEmpty(m.NetworkComponentModel),
		NetworkComponentName:            flex.ExpandStringPointerNullAsEmpty(m.NetworkComponentName),
		NetworkComponentPortDescription: flex.ExpandStringPointerNullAsEmpty(m.NetworkComponentPortDescription),
		NetworkComponentPortName:        flex.ExpandStringPointerNullAsEmpty(m.NetworkComponentPortName),
		NetworkComponentPortNumber:      flex.ExpandStringPointerNullAsEmpty(m.NetworkComponentPortNumber),
		NetworkComponentType:            flex.ExpandStringPointerNullAsEmpty(m.NetworkComponentType),
		NetworkComponentVendor:          flex.ExpandStringPointerNullAsEmpty(m.NetworkComponentVendor),
		OpenPorts:                       flex.ExpandStringPointerNullAsEmpty(m.OpenPorts),
		Os:                              flex.ExpandStringPointerNullAsEmpty(m.Os),
		PortDuplex:                      flex.ExpandStringPointerNullAsEmpty(m.PortDuplex),
		PortLinkStatus:                  flex.ExpandStringPointerNullAsEmpty(m.PortLinkStatus),
		PortSpeed:                       flex.ExpandStringPointerNullAsEmpty(m.PortSpeed),
		PortStatus:                      flex.ExpandStringPointerNullAsEmpty(m.PortStatus),
		PortType:                        flex.ExpandStringPointerNullAsEmpty(m.PortType),
		PortVlanDescription:             flex.ExpandStringPointerNullAsEmpty(m.PortVlanDescription),
		PortVlanName:                    flex.ExpandStringPointerNullAsEmpty(m.PortVlanName),
		PortVlanNumber:                  flex.ExpandStringPointerNullAsEmpty(m.PortVlanNumber),
		VAdapter:                        flex.ExpandStringPointerNullAsEmpty(m.VAdapter),
		VCluster:                        flex.ExpandStringPointerNullAsEmpty(m.VCluster),
		VDatacenter:                     flex.ExpandStringPointerNullAsEmpty(m.VDatacenter),
		VEntityName:                     flex.ExpandStringPointerNullAsEmpty(m.VEntityName),
		VEntityType:                     flex.ExpandStringPointerNullAsEmpty(m.VEntityType),
		VHost:                           flex.ExpandStringPointerNullAsEmpty(m.VHost),
		VSwitch:                         flex.ExpandStringPointerNullAsEmpty(m.VSwitch),
		VmiName:                         flex.ExpandStringPointerNullAsEmpty(m.VmiName),
		VmiId:                           flex.ExpandStringPointerNullAsEmpty(m.VmiId),
		VlanPortGroup:                   flex.ExpandStringPointerNullAsEmpty(m.VlanPortGroup),
		VswitchName:                     flex.ExpandStringPointerNullAsEmpty(m.VswitchName),
		VswitchId:                       flex.ExpandStringPointerNullAsEmpty(m.VswitchId),
		VswitchType:                     flex.ExpandStringPointerNullAsEmpty(m.VswitchType),
		VswitchIpv6Enabled:              flex.ExpandBoolPointer(m.VswitchIpv6Enabled),
		VportName:                       flex.ExpandStringPointerNullAsEmpty(m.VportName),
		VportMacAddress:                 flex.ExpandStringPointerNullAsEmpty(m.VportMacAddress),
		VportLinkStatus:                 flex.ExpandStringPointerNullAsEmpty(m.VportLinkStatus),
		VportConfSpeed:                  flex.ExpandStringPointerNullAsEmpty(m.VportConfSpeed),
		VportConfMode:                   flex.ExpandStringPointerNullAsEmpty(m.VportConfMode),
		VportSpeed:                      flex.ExpandStringPointerNullAsEmpty(m.VportSpeed),
		VportMode:                       flex.ExpandStringPointerNullAsEmpty(m.VportMode),
		VswitchSegmentType:              flex.ExpandStringPointerNullAsEmpty(m.VswitchSegmentType),
		VswitchSegmentName:              flex.ExpandStringPointerNullAsEmpty(m.VswitchSegmentName),
		VswitchSegmentId:                flex.ExpandStringPointerNullAsEmpty(m.VswitchSegmentId),
		VswitchSegmentPortGroup:         flex.ExpandStringPointerNullAsEmpty(m.VswitchSegmentPortGroup),
		VswitchAvailablePortsCount:      flex.ExpandInt64Pointer(m.VswitchAvailablePortsCount),
		VswitchTepType:                  flex.ExpandStringPointerNullAsEmpty(m.VswitchTepType),
		VswitchTepIp:                    flex.ExpandStringPointerNullAsEmpty(m.VswitchTepIp),
		VswitchTepPortGroup:             flex.ExpandStringPointerNullAsEmpty(m.VswitchTepPortGroup),
		VswitchTepVlan:                  flex.ExpandStringPointerNullAsEmpty(m.VswitchTepVlan),
		VswitchTepDhcpServer:            flex.ExpandStringPointerNullAsEmpty(m.VswitchTepDhcpServer),
		VswitchTepMulticast:             flex.ExpandStringPointerNullAsEmpty(m.VswitchTepMulticast),
		VmhostIpAddress:                 flex.ExpandStringPointerNullAsEmpty(m.VmhostIpAddress),
		VmhostName:                      flex.ExpandStringPointerNullAsEmpty(m.VmhostName),
		VmhostMacAddress:                flex.ExpandStringPointerNullAsEmpty(m.VmhostMacAddress),
		VmhostSubnetCidr:                flex.ExpandInt64Pointer(m.VmhostSubnetCidr),
		VmhostNicNames:                  flex.ExpandStringPointerNullAsEmpty(m.VmhostNicNames),
		VmiTenantId:                     flex.ExpandStringPointerNullAsEmpty(m.VmiTenantId),
		CmpType:                         flex.ExpandStringPointerNullAsEmpty(m.CmpType),
		VmiIpType:                       flex.ExpandStringPointerNullAsEmpty(m.VmiIpType),
		VmiPrivateAddress:               flex.ExpandStringPointerNullAsEmpty(m.VmiPrivateAddress),
		VmiIsPublicAddress:              flex.ExpandBoolPointer(m.VmiIsPublicAddress),
		CiscoIseSsid:                    flex.ExpandStringPointerNullAsEmpty(m.CiscoIseSsid),
		CiscoIseEndpointProfile:         flex.ExpandStringPointerNullAsEmpty(m.CiscoIseEndpointProfile),
		CiscoIseSessionState:            flex.ExpandStringPointerNullAsEmpty(m.CiscoIseSessionState),
		CiscoIseSecurityGroup:           flex.ExpandStringPointerNullAsEmpty(m.CiscoIseSecurityGroup),
		TaskName:                        flex.ExpandStringPointerNullAsEmpty(m.TaskName),
		NetworkComponentLocation:        flex.ExpandStringPointerNullAsEmpty(m.NetworkComponentLocation),
		NetworkComponentContact:         flex.ExpandStringPointerNullAsEmpty(m.NetworkComponentContact),
		DeviceLocation:                  flex.ExpandStringPointerNullAsEmpty(m.DeviceLocation),
		DeviceContact:                   flex.ExpandStringPointerNullAsEmpty(m.DeviceContact),
		ApName:                          flex.ExpandStringPointerNullAsEmpty(m.ApName),
		ApIpAddress:                     flex.ExpandStringPointerNullAsEmpty(m.ApIpAddress),
		ApSsid:                          flex.ExpandStringPointerNullAsEmpty(m.ApSsid),
		BridgeDomain:                    flex.ExpandStringPointerNullAsEmpty(m.BridgeDomain),
		EndpointGroups:                  flex.ExpandStringPointerNullAsEmpty(m.EndpointGroups),
		Tenant:                          flex.ExpandStringPointerNullAsEmpty(m.Tenant),
		VrfName:                         flex.ExpandStringPointerNullAsEmpty(m.VrfName),
		VrfDescription:                  flex.ExpandStringPointerNullAsEmpty(m.VrfDescription),
		VrfRd:                           flex.ExpandStringPointerNullAsEmpty(m.VrfRd),
		BgpAs:                           flex.ExpandInt64Pointer(m.BgpAs),
	}
	return to
}

// FlattenRecordHostIpv6addrDiscoveredData converts an SDK type to Terraform Object
func FlattenRecordHostIpv6addrDiscoveredData(ctx context.Context, from *niosdns.RecordHostIpv6addrDiscoveredData, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RecordHostIpv6addrDiscoveredDataAttrTypes)
	}
	m := &RecordHostIpv6addrDiscoveredDataModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RecordHostIpv6addrDiscoveredDataAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *RecordHostIpv6addrDiscoveredDataModel) Flatten(ctx context.Context, from *niosdns.RecordHostIpv6addrDiscoveredData, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.DeviceModel = flex.FlattenStringPointerEmptyAsNull(from.DeviceModel)
	m.DevicePortName = flex.FlattenStringPointerEmptyAsNull(from.DevicePortName)
	m.DevicePortType = flex.FlattenStringPointerEmptyAsNull(from.DevicePortType)
	m.DeviceType = flex.FlattenStringPointerEmptyAsNull(from.DeviceType)
	m.DeviceVendor = flex.FlattenStringPointerEmptyAsNull(from.DeviceVendor)
	m.DiscoveredName = flex.FlattenStringPointerEmptyAsNull(from.DiscoveredName)
	m.Discoverer = flex.FlattenStringPointerEmptyAsNull(from.Discoverer)
	m.Duid = flex.FlattenStringPointerEmptyAsNull(from.Duid)
	m.FirstDiscovered = flex.FlattenInt64Pointer(from.FirstDiscovered)
	m.IprgNo = flex.FlattenInt64Pointer(from.IprgNo)
	m.IprgState = flex.FlattenStringPointerEmptyAsNull(from.IprgState)
	m.IprgType = flex.FlattenStringPointerEmptyAsNull(from.IprgType)
	m.LastDiscovered = flex.FlattenInt64Pointer(from.LastDiscovered)
	m.MacAddress = flex.FlattenStringPointerEmptyAsNull(from.MacAddress)
	m.MgmtIpAddress = flex.FlattenStringPointerEmptyAsNull(from.MgmtIpAddress)
	m.NetbiosName = flex.FlattenStringPointerEmptyAsNull(from.NetbiosName)
	m.NetworkComponentDescription = flex.FlattenStringPointerEmptyAsNull(from.NetworkComponentDescription)
	m.NetworkComponentIp = flex.FlattenStringPointerEmptyAsNull(from.NetworkComponentIp)
	m.NetworkComponentModel = flex.FlattenStringPointerEmptyAsNull(from.NetworkComponentModel)
	m.NetworkComponentName = flex.FlattenStringPointerEmptyAsNull(from.NetworkComponentName)
	m.NetworkComponentPortDescription = flex.FlattenStringPointerEmptyAsNull(from.NetworkComponentPortDescription)
	m.NetworkComponentPortName = flex.FlattenStringPointerEmptyAsNull(from.NetworkComponentPortName)
	m.NetworkComponentPortNumber = flex.FlattenStringPointerEmptyAsNull(from.NetworkComponentPortNumber)
	m.NetworkComponentType = flex.FlattenStringPointerEmptyAsNull(from.NetworkComponentType)
	m.NetworkComponentVendor = flex.FlattenStringPointerEmptyAsNull(from.NetworkComponentVendor)
	m.OpenPorts = flex.FlattenStringPointerEmptyAsNull(from.OpenPorts)
	m.Os = flex.FlattenStringPointerEmptyAsNull(from.Os)
	m.PortDuplex = flex.FlattenStringPointerEmptyAsNull(from.PortDuplex)
	m.PortLinkStatus = flex.FlattenStringPointerEmptyAsNull(from.PortLinkStatus)
	m.PortSpeed = flex.FlattenStringPointerEmptyAsNull(from.PortSpeed)
	m.PortStatus = flex.FlattenStringPointerEmptyAsNull(from.PortStatus)
	m.PortType = flex.FlattenStringPointerEmptyAsNull(from.PortType)
	m.PortVlanDescription = flex.FlattenStringPointerEmptyAsNull(from.PortVlanDescription)
	m.PortVlanName = flex.FlattenStringPointerEmptyAsNull(from.PortVlanName)
	m.PortVlanNumber = flex.FlattenStringPointerEmptyAsNull(from.PortVlanNumber)
	m.VAdapter = flex.FlattenStringPointerEmptyAsNull(from.VAdapter)
	m.VCluster = flex.FlattenStringPointerEmptyAsNull(from.VCluster)
	m.VDatacenter = flex.FlattenStringPointerEmptyAsNull(from.VDatacenter)
	m.VEntityName = flex.FlattenStringPointerEmptyAsNull(from.VEntityName)
	m.VEntityType = flex.FlattenStringPointerEmptyAsNull(from.VEntityType)
	m.VHost = flex.FlattenStringPointerEmptyAsNull(from.VHost)
	m.VSwitch = flex.FlattenStringPointerEmptyAsNull(from.VSwitch)
	m.VmiName = flex.FlattenStringPointerEmptyAsNull(from.VmiName)
	m.VmiId = flex.FlattenStringPointerEmptyAsNull(from.VmiId)
	m.VlanPortGroup = flex.FlattenStringPointerEmptyAsNull(from.VlanPortGroup)
	m.VswitchName = flex.FlattenStringPointerEmptyAsNull(from.VswitchName)
	m.VswitchId = flex.FlattenStringPointerEmptyAsNull(from.VswitchId)
	m.VswitchType = flex.FlattenStringPointerEmptyAsNull(from.VswitchType)
	m.VswitchIpv6Enabled = flex.FlattenBoolPointer(from.VswitchIpv6Enabled)
	m.VportName = flex.FlattenStringPointerEmptyAsNull(from.VportName)
	m.VportMacAddress = flex.FlattenStringPointerEmptyAsNull(from.VportMacAddress)
	m.VportLinkStatus = flex.FlattenStringPointerEmptyAsNull(from.VportLinkStatus)
	m.VportConfSpeed = flex.FlattenStringPointerEmptyAsNull(from.VportConfSpeed)
	m.VportConfMode = flex.FlattenStringPointerEmptyAsNull(from.VportConfMode)
	m.VportSpeed = flex.FlattenStringPointerEmptyAsNull(from.VportSpeed)
	m.VportMode = flex.FlattenStringPointerEmptyAsNull(from.VportMode)
	m.VswitchSegmentType = flex.FlattenStringPointerEmptyAsNull(from.VswitchSegmentType)
	m.VswitchSegmentName = flex.FlattenStringPointerEmptyAsNull(from.VswitchSegmentName)
	m.VswitchSegmentId = flex.FlattenStringPointerEmptyAsNull(from.VswitchSegmentId)
	m.VswitchSegmentPortGroup = flex.FlattenStringPointerEmptyAsNull(from.VswitchSegmentPortGroup)
	m.VswitchAvailablePortsCount = flex.FlattenInt64Pointer(from.VswitchAvailablePortsCount)
	m.VswitchTepType = flex.FlattenStringPointerEmptyAsNull(from.VswitchTepType)
	m.VswitchTepIp = flex.FlattenStringPointerEmptyAsNull(from.VswitchTepIp)
	m.VswitchTepPortGroup = flex.FlattenStringPointerEmptyAsNull(from.VswitchTepPortGroup)
	m.VswitchTepVlan = flex.FlattenStringPointerEmptyAsNull(from.VswitchTepVlan)
	m.VswitchTepDhcpServer = flex.FlattenStringPointerEmptyAsNull(from.VswitchTepDhcpServer)
	m.VswitchTepMulticast = flex.FlattenStringPointerEmptyAsNull(from.VswitchTepMulticast)
	m.VmhostIpAddress = flex.FlattenStringPointerEmptyAsNull(from.VmhostIpAddress)
	m.VmhostName = flex.FlattenStringPointerEmptyAsNull(from.VmhostName)
	m.VmhostMacAddress = flex.FlattenStringPointerEmptyAsNull(from.VmhostMacAddress)
	m.VmhostSubnetCidr = flex.FlattenInt64Pointer(from.VmhostSubnetCidr)
	m.VmhostNicNames = flex.FlattenStringPointerEmptyAsNull(from.VmhostNicNames)
	m.VmiTenantId = flex.FlattenStringPointerEmptyAsNull(from.VmiTenantId)
	m.CmpType = flex.FlattenStringPointerEmptyAsNull(from.CmpType)
	m.VmiIpType = flex.FlattenStringPointerEmptyAsNull(from.VmiIpType)
	m.VmiPrivateAddress = flex.FlattenStringPointerEmptyAsNull(from.VmiPrivateAddress)
	m.VmiIsPublicAddress = flex.FlattenBoolPointer(from.VmiIsPublicAddress)
	m.CiscoIseSsid = flex.FlattenStringPointerEmptyAsNull(from.CiscoIseSsid)
	m.CiscoIseEndpointProfile = flex.FlattenStringPointerEmptyAsNull(from.CiscoIseEndpointProfile)
	m.CiscoIseSessionState = flex.FlattenStringPointerEmptyAsNull(from.CiscoIseSessionState)
	m.CiscoIseSecurityGroup = flex.FlattenStringPointerEmptyAsNull(from.CiscoIseSecurityGroup)
	m.TaskName = flex.FlattenStringPointerEmptyAsNull(from.TaskName)
	m.NetworkComponentLocation = flex.FlattenStringPointerEmptyAsNull(from.NetworkComponentLocation)
	m.NetworkComponentContact = flex.FlattenStringPointerEmptyAsNull(from.NetworkComponentContact)
	m.DeviceLocation = flex.FlattenStringPointerEmptyAsNull(from.DeviceLocation)
	m.DeviceContact = flex.FlattenStringPointerEmptyAsNull(from.DeviceContact)
	m.ApName = flex.FlattenStringPointerEmptyAsNull(from.ApName)
	m.ApIpAddress = flex.FlattenStringPointerEmptyAsNull(from.ApIpAddress)
	m.ApSsid = flex.FlattenStringPointerEmptyAsNull(from.ApSsid)
	m.BridgeDomain = flex.FlattenStringPointerEmptyAsNull(from.BridgeDomain)
	m.EndpointGroups = flex.FlattenStringPointerEmptyAsNull(from.EndpointGroups)
	m.Tenant = flex.FlattenStringPointerEmptyAsNull(from.Tenant)
	m.VrfName = flex.FlattenStringPointerEmptyAsNull(from.VrfName)
	m.VrfDescription = flex.FlattenStringPointerEmptyAsNull(from.VrfDescription)
	m.VrfRd = flex.FlattenStringPointerEmptyAsNull(from.VrfRd)
	m.BgpAs = flex.FlattenInt64Pointer(from.BgpAs)
}
