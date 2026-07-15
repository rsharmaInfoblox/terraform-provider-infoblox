package dns

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/dynamicallocation"
)

// ValidateRecordA validates the RecordA configuration.
func ValidateRecordA(ctx context.Context, niosBlock, uddiBlock types.Object, resp *resource.ValidateConfigResponse) {
	if !niosBlock.IsNull() && !niosBlock.IsUnknown() {
		validateRecordANIOSConfig(ctx, niosBlock, resp)
	}
	if !uddiBlock.IsNull() && !uddiBlock.IsUnknown() {
		validateRecordAUDDIConfig(ctx, uddiBlock, resp)
	}
}

func validateRecordANIOSConfig(ctx context.Context, data types.Object, resp *resource.ValidateConfigResponse) {
}

func validateRecordAUDDIConfig(ctx context.Context, data types.Object, resp *resource.ValidateConfigResponse) {
	attrs := data.Attributes()

	// rdata: the address subfield is required for an A record.
	if rdataMap, ok := attrs["rdata"].(types.Map); ok && !rdataMap.IsNull() && !rdataMap.IsUnknown() {
		rdataPath := path.Root("uddi").AtName("rdata")
		address, present := rdataMap.Elements()["address"]
		if !present {
			resp.Diagnostics.AddAttributeError(
				rdataPath,
				"Missing Required Subfield",
				"The `address` subfield is required in `rdata` for an A record.",
			)
		} else if addrStr, ok := address.(types.String); ok && !addrStr.IsUnknown() {
			if addrStr.IsNull() || strings.TrimSpace(addrStr.ValueString()) == "" {
				resp.Diagnostics.AddAttributeError(
					rdataPath.AtMapKey("address"),
					"Invalid Subfield Value",
					"The `address` subfield in `rdata` must be a non-empty IPv4 address for an A record.",
				)
			}
		}
	}

	// options: only the boolean keys `create_ptr` and `check_rmz` are valid.
	if optionsMap, ok := attrs["options"].(types.Map); ok && !optionsMap.IsNull() && !optionsMap.IsUnknown() {
		optionsPath := path.Root("uddi").AtName("options")
		allowed := map[string]struct{}{"create_ptr": {}, "check_rmz": {}}
		for key, val := range optionsMap.Elements() {
			if _, valid := allowed[key]; !valid {
				resp.Diagnostics.AddAttributeError(
					optionsPath.AtMapKey(key),
					"Invalid Option",
					fmt.Sprintf("`%s` is not a valid option for an A record. Valid options are: create_ptr, check_rmz.", key),
				)
				continue
			}
			if valStr, ok := val.(types.String); ok && !valStr.IsUnknown() && !valStr.IsNull() {
				if _, err := strconv.ParseBool(valStr.ValueString()); err != nil {
					resp.Diagnostics.AddAttributeError(
						optionsPath.AtMapKey(key),
						"Invalid Option Value",
						fmt.Sprintf("`%s` must be a boolean value (\"true\" or \"false\"), got %q.", key, valStr.ValueString()),
					)
				}
			}
		}
	}
}

var boolRecordAOptionKeys = []string{"create_ptr", "check_rmz"}

func BuildRecordAFuncCall(ctx context.Context, data types.Object, diags *diag.Diagnostics) *niosdns.FuncCall {
	if data.IsNull() || data.IsUnknown() {
		return nil
	}

	var m dynamicallocation.NextAvailableIpModel
	diags.Append(data.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.FuncCall(ctx, "Ipv4addr", "network", diags)
}

func PostExpandRecordAUDDI(ctx context.Context, ext *coremodel.UDDIRecordAExt, diags *diag.Diagnostics) *coremodel.UDDIRecordAExt {
	if ext == nil {
		return ext
	}
	if ext.Options != nil {
		for _, k := range boolRecordAOptionKeys {
			if v, ok := ext.Options[k].(string); ok {
				if b, err := strconv.ParseBool(v); err == nil {
					ext.Options[k] = b
				}
			}
		}
	}
	return ext
}

func PostFlattenRecordAUDDI(ctx context.Context, plannedUDDI, uddiObj types.Object, diags *diag.Diagnostics) types.Object {
	if uddiObj.IsNull() || uddiObj.IsUnknown() {
		return uddiObj
	}

	attrs := uddiObj.Attributes()

	// Preserve the planned options value.
	if !plannedUDDI.IsNull() && !plannedUDDI.IsUnknown() {
		if plannedOptions, ok := plannedUDDI.Attributes()["options"].(types.Map); ok && !plannedOptions.IsNull() && !plannedOptions.IsUnknown() {
			attrs["options"] = plannedOptions
		} else {
			attrs["options"] = types.MapNull(types.StringType)
		}
	} else {
		attrs["options"] = types.MapNull(types.StringType)
	}

	updated, d := types.ObjectValue(uddiObj.AttributeTypes(ctx), attrs)
	diags.Append(d...)
	if diags.HasError() {
		return uddiObj
	}
	return updated
}
