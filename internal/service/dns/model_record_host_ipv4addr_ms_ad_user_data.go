package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// RecordHostIpv4addrMsAdUserDataModel is the Terraform model for RecordHostIpv4addrMsAdUserData
type RecordHostIpv4addrMsAdUserDataModel struct {
	ActiveUsersCount types.Int64 `tfsdk:"active_users_count"`
}

// RecordHostIpv4addrMsAdUserDataAttrTypes contains the attribute types for RecordHostIpv4addrMsAdUserDataModel
var RecordHostIpv4addrMsAdUserDataAttrTypes = map[string]attr.Type{
	"active_users_count": types.Int64Type,
}

// RecordHostIpv4addrMsAdUserDataResourceSchemaAttributes contains the schema attributes for RecordHostIpv4addrMsAdUserDataModel
var RecordHostIpv4addrMsAdUserDataResourceSchemaAttributes = map[string]schema.Attribute{
	"active_users_count": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The number of active users.",
	},
}

// ExpandRecordHostIpv4addrMsAdUserData converts a Terraform Object to SDK type
func ExpandRecordHostIpv4addrMsAdUserData(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdns.RecordHostIpv4addrMsAdUserData {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RecordHostIpv4addrMsAdUserDataModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *RecordHostIpv4addrMsAdUserDataModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdns.RecordHostIpv4addrMsAdUserData {
	if m == nil {
		return nil
	}
	to := &niosdns.RecordHostIpv4addrMsAdUserData{
		ActiveUsersCount: flex.ExpandInt64Pointer(m.ActiveUsersCount),
	}
	return to
}

// FlattenRecordHostIpv4addrMsAdUserData converts an SDK type to Terraform Object
func FlattenRecordHostIpv4addrMsAdUserData(ctx context.Context, from *niosdns.RecordHostIpv4addrMsAdUserData, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RecordHostIpv4addrMsAdUserDataAttrTypes)
	}
	m := &RecordHostIpv4addrMsAdUserDataModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RecordHostIpv4addrMsAdUserDataAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *RecordHostIpv4addrMsAdUserDataModel) Flatten(ctx context.Context, from *niosdns.RecordHostIpv4addrMsAdUserData, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.ActiveUsersCount = flex.FlattenInt64Pointer(from.ActiveUsersCount)
}
