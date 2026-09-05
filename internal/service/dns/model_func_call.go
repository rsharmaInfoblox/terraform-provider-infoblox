package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mapplanmodifier "github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	stringplanmodifier "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// FuncCallModel is the Terraform model for FuncCall
type FuncCallModel struct {
	AttributeName    types.String `tfsdk:"attribute_name"`
	ObjectFunction   types.String `tfsdk:"object_function"`
	Parameters       types.Map    `tfsdk:"parameters"`
	ResultField      types.String `tfsdk:"result_field"`
	Object           types.String `tfsdk:"object"`
	ObjectParameters types.Map    `tfsdk:"object_parameters"`
}

// FuncCallAttrTypes contains the attribute types for FuncCallModel
var FuncCallAttrTypes = map[string]attr.Type{
	"attribute_name":    types.StringType,
	"object_function":   types.StringType,
	"parameters":        types.MapType{ElemType: types.StringType},
	"result_field":      types.StringType,
	"object":            types.StringType,
	"object_parameters": types.MapType{ElemType: types.StringType},
}

// FuncCallResourceSchemaAttributes contains the schema attributes for FuncCallModel
var FuncCallResourceSchemaAttributes = map[string]schema.Attribute{
	"attribute_name": schema.StringAttribute{
		Required: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplace(),
		},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The attribute to be called.",
	},
	"object_function": schema.StringAttribute{
		Optional: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplace(),
		},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The function to be called.",
	},
	"parameters": schema.MapAttribute{
		ElementType: types.StringType,
		Optional:    true,
		PlanModifiers: []planmodifier.Map{
			mapplanmodifier.RequiresReplace(),
		},
		MarkdownDescription: "The parameters for the function.",
	},
	"result_field": schema.StringAttribute{
		Optional: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplace(),
		},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The result field of the function.",
	},
	"object": schema.StringAttribute{
		Optional: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplace(),
		},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The object to be called.",
	},
	"object_parameters": schema.MapAttribute{
		ElementType: types.StringType,
		Optional:    true,
		PlanModifiers: []planmodifier.Map{
			mapplanmodifier.RequiresReplace(),
		},
		MarkdownDescription: "The parameters for the object.",
	},
}

// ExpandFuncCall converts a Terraform Object to SDK type
func ExpandFuncCall(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdns.FuncCall {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m FuncCallModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *FuncCallModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdns.FuncCall {
	if m == nil {
		return nil
	}
	to := &niosdns.FuncCall{
		AttributeName:    flex.ExpandString(m.AttributeName),
		ObjectFunction:   flex.ExpandStringPointerNullAsEmpty(m.ObjectFunction),
		Parameters:       flex.ExpandMapStringAny(ctx, m.Parameters, diags),
		ResultField:      flex.ExpandStringPointerNullAsEmpty(m.ResultField),
		Object:           flex.ExpandStringPointerNullAsEmpty(m.Object),
		ObjectParameters: flex.ExpandMapStringAny(ctx, m.ObjectParameters, diags),
	}
	return to
}

// FlattenFuncCall converts an SDK type to Terraform Object
func FlattenFuncCall(ctx context.Context, from *niosdns.FuncCall, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(FuncCallAttrTypes)
	}
	m := &FuncCallModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, FuncCallAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *FuncCallModel) Flatten(ctx context.Context, from *niosdns.FuncCall, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.AttributeName = flex.FlattenString(from.AttributeName)
	m.ObjectFunction = flex.FlattenStringPointerEmptyAsNull(from.ObjectFunction)
	m.Parameters = flex.FlattenMapStringAny(ctx, from.Parameters, diags)
	m.ResultField = flex.FlattenStringPointerEmptyAsNull(from.ResultField)
	m.Object = flex.FlattenStringPointerEmptyAsNull(from.Object)
	m.ObjectParameters = flex.FlattenMapStringAny(ctx, from.ObjectParameters, diags)
}
