/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

package int64planmodifier_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"

	"github.com/hashicorp/terraform-plugin-go/tftypes"

	"github.com/orange-cloudavenue/terraform-plugin-framework-planmodifiers/int64planmodifier"
)

func Test_requireReplaceIfBool(t *testing.T) {
	t.Parallel()

	testSchema := schema.Schema{
		Attributes: map[string]schema.Attribute{
			"testattr": schema.Int64Attribute{},
			"testbool": schema.BoolAttribute{},
		},
	}

	nullPlan := tfsdk.Plan{
		Schema: testSchema,
		Raw: tftypes.NewValue(
			testSchema.Type().TerraformType(context.Background()),
			nil,
		),
	}

	nullState := tfsdk.State{
		Schema: testSchema,
		Raw: tftypes.NewValue(
			testSchema.Type().TerraformType(context.Background()),
			nil,
		),
	}

	testPlan := func(value types.Int64) tfsdk.Plan {
		tfValue, err := value.ToTerraformValue(context.Background())
		if err != nil {
			panic("ToTerraformValue error: " + err.Error())
		}

		return tfsdk.Plan{
			Schema: testSchema,
			Raw: tftypes.NewValue(
				testSchema.Type().TerraformType(context.Background()),
				map[string]tftypes.Value{
					"testattr": tfValue,
					"testbool": tftypes.NewValue(tftypes.Bool, true),
				},
			),
		}
	}

	testState := func(value types.Int64) tfsdk.State {
		tfValue, err := value.ToTerraformValue(context.Background())
		if err != nil {
			panic("ToTerraformValue error: " + err.Error())
		}

		return tfsdk.State{
			Schema: testSchema,
			Raw: tftypes.NewValue(
				testSchema.Type().TerraformType(context.Background()),
				map[string]tftypes.Value{
					"testattr": tfValue,
					"testbool": tftypes.NewValue(tftypes.Bool, true),
				},
			),
		}
	}

	testCases := map[string]struct {
		request  planmodifier.Int64Request
		expected *planmodifier.Int64Response
	}{
		"state-null": {
			// resource creation
			request: planmodifier.Int64Request{
				Plan:       testPlan(types.Int64Unknown()),
				PlanValue:  types.Int64Unknown(),
				State:      nullState,
				StateValue: types.Int64Null(),
			},
			expected: &planmodifier.Int64Response{
				PlanValue:       types.Int64Unknown(),
				RequiresReplace: false,
			},
		},
		"plan-null": {
			// resource destroy
			request: planmodifier.Int64Request{
				Plan:       nullPlan,
				PlanValue:  types.Int64Null(),
				State:      testState(types.Int64Value(10)),
				StateValue: types.Int64Value(10),
			},
			expected: &planmodifier.Int64Response{
				PlanValue:       types.Int64Null(),
				RequiresReplace: false,
			},
		},
		"planvalue-statevalue-different": {
			request: planmodifier.Int64Request{
				Plan:       testPlan(types.Int64Value(20)),
				PlanValue:  types.Int64Value(20),
				State:      testState(types.Int64Value(10)),
				StateValue: types.Int64Value(10),
			},
			expected: &planmodifier.Int64Response{
				PlanValue:       types.Int64Value(20),
				RequiresReplace: true,
			},
		},
		"planvalue-statevalue-equal": {
			request: planmodifier.Int64Request{
				Plan:       testPlan(types.Int64Value(10)),
				PlanValue:  types.Int64Value(10),
				State:      testState(types.Int64Value(10)),
				StateValue: types.Int64Value(10),
			},
			expected: &planmodifier.Int64Response{
				PlanValue:       types.Int64Value(10),
				RequiresReplace: false,
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			resp := &planmodifier.Int64Response{
				PlanValue: testCase.request.PlanValue,
			}

			int64planmodifier.RequireReplaceIfBool(path.Root("testbool"), true).PlanModifyInt64(context.Background(), testCase.request, resp)

			if diff := cmp.Diff(testCase.expected, resp); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
