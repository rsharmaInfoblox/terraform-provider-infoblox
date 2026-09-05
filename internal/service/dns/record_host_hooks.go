package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateRecordHost validates the RecordHost configuration.
func ValidateRecordHost(ctx context.Context, data RecordHostModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSRecordHostModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateRecordHostNIOSConfig(ctx, nios, resp)
	}
}

func validateRecordHostNIOSConfig(ctx context.Context, m *NIOSRecordHostModel, resp *resource.ValidateConfigResponse) {
}
