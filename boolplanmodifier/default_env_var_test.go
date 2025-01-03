/*
 * SPDX-FileCopyrightText: Copyright (c) Orange Business Services SA
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at <https://www.mozilla.org/en-US/MPL/2.0/>
 * or see the "LICENSE" file for more details.
 */

// Package boolplanmodifier provides a plan modifier for boolean values.
package boolplanmodifier_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"

	"github.com/orange-cloudavenue/terraform-plugin-framework-planmodifiers/boolplanmodifier"
)

func TestDefaultEnvVarModifierPlanModifyBool(t *testing.T) {
	const (
		envVarName = "TEST_VAR"
		envValue   = true
	)

	testCases := map[string]struct {
		request  planmodifier.BoolRequest
		expected *planmodifier.BoolResponse
	}{
		"null-state": {
			// when we first create the resource, use the unknown
			// value
			request: planmodifier.BoolRequest{
				StateValue:  types.BoolNull(),
				PlanValue:   types.BoolUnknown(),
				ConfigValue: types.BoolNull(),
			},
			expected: &planmodifier.BoolResponse{
				PlanValue: types.BoolValue(envValue),
			},
		},
		"known-plan": {
			// this would really only happen if we had a plan
			// modifier setting the value before this plan modifier
			// got to it
			//
			// but we still want to preserve that value, in this
			// case
			request: planmodifier.BoolRequest{
				StateValue:  types.BoolValue(false),
				PlanValue:   types.BoolValue(true),
				ConfigValue: types.BoolNull(),
			},
			expected: &planmodifier.BoolResponse{
				PlanValue: types.BoolValue(true),
			},
		},
		"non-null-state-unknown-plan": {
			// this is the situation we want to preserve the state
			// in
			request: planmodifier.BoolRequest{
				StateValue:  types.BoolValue(true),
				PlanValue:   types.BoolUnknown(),
				ConfigValue: types.BoolNull(),
			},
			expected: &planmodifier.BoolResponse{
				PlanValue: types.BoolValue(envValue),
			},
		},
		"unknown-config": {
			// this is the situation in which a user is
			// interpolating into a field. We want that to still
			// show up as unknown, otherwise they'll get apply-time
			// errors for changing the value even though we knew it
			// was legitimately possible for it to change and the
			// provider can't prevent this from happening
			request: planmodifier.BoolRequest{
				StateValue:  types.BoolValue(true),
				PlanValue:   types.BoolUnknown(),
				ConfigValue: types.BoolUnknown(),
			},
			expected: &planmodifier.BoolResponse{
				PlanValue: types.BoolValue(envValue),
			},
		},
	}

	for name, testCase := range testCases {
		// set environnement variable
		t.Setenv(envVarName, "true")

		t.Run(name, func(t *testing.T) {
			resp := &planmodifier.BoolResponse{
				PlanValue: testCase.request.PlanValue,
			}

			boolplanmodifier.SetDefaultEnvVar(envVarName).PlanModifyBool(context.Background(), testCase.request, resp)

			if diff := cmp.Diff(testCase.expected, resp); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
