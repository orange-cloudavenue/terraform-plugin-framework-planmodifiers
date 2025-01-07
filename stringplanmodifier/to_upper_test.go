/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

// Package stringplanmodifier provides a plan modifier for string values.
package stringplanmodifier_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"

	"github.com/orange-cloudavenue/terraform-plugin-framework-planmodifiers/stringplanmodifier"
)

func TestToUpperPlanModifyString(t *testing.T) {
	t.Parallel()

	type testCase struct {
		val         types.String
		exceptedVal types.String
		expectError bool
	}

	tests := map[string]testCase{
		"unknown String": {
			val:         types.StringUnknown(),
			exceptedVal: types.StringNull(),
		},
		"null String": {
			val:         types.StringNull(),
			exceptedVal: types.StringNull(),
		},
		"valid String": {
			val:         types.StringValue("test"),
			exceptedVal: types.StringValue("TEST"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			request := planmodifier.StringRequest{
				Path:           path.Root("test"),
				PathExpression: path.MatchRoot("test"),
				ConfigValue:    test.val,
			}

			resp := &planmodifier.StringResponse{}
			stringplanmodifier.ToUpper().PlanModifyString(context.Background(), request, resp)

			if diff := cmp.Diff(test.exceptedVal, resp.PlanValue); diff != "" && !test.expectError {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
