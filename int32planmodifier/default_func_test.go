/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

// Package int32planmodifier provides a plan modifier for int32 values.
package int32planmodifier_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"

	"github.com/orange-cloudavenue/terraform-plugin-framework-planmodifiers/int32planmodifier"
)

func TestDefaultFuncModifierPlanModifyInt32(t *testing.T) {
	const expectedValue = 123

	testCases := map[string]struct {
		request  planmodifier.Int32Request
		expected *planmodifier.Int32Response
	}{
		"null-state": {
			// when we first create the resource, use the unknown
			// value
			request: planmodifier.Int32Request{
				StateValue:  types.Int32Null(),
				PlanValue:   types.Int32Unknown(),
				ConfigValue: types.Int32Null(),
			},
			expected: &planmodifier.Int32Response{
				PlanValue: types.Int32Value(expectedValue),
			},
		},
		"known-plan": {
			// this would really only happen if we had a plan
			// modifier setting the value before this plan modifier
			// got to it
			//
			// but we still want to preserve that value, in this
			// case
			request: planmodifier.Int32Request{
				StateValue:  types.Int32Value(10),
				PlanValue:   types.Int32Value(11),
				ConfigValue: types.Int32Null(),
			},
			expected: &planmodifier.Int32Response{
				PlanValue: types.Int32Value(11),
			},
		},
		"non-null-state-unknown-plan": {
			// this is the situation we want to preserve the state
			// in
			request: planmodifier.Int32Request{
				StateValue:  types.Int32Value(10),
				PlanValue:   types.Int32Unknown(),
				ConfigValue: types.Int32Null(),
			},
			expected: &planmodifier.Int32Response{
				PlanValue: types.Int32Value(expectedValue),
			},
		},
		"unknown-config": {
			// this is the situation in which a user is
			// interpolating into a field. We want that to still
			// show up as unknown, otherwise they'll get apply-time
			// errors for changing the value even though we knew it
			// was legitimately possible for it to change and the
			// provider can't prevent this from happening
			request: planmodifier.Int32Request{
				StateValue:  types.Int32Value(10),
				PlanValue:   types.Int32Unknown(),
				ConfigValue: types.Int32Unknown(),
			},
			expected: &planmodifier.Int32Response{
				PlanValue: types.Int32Value(expectedValue),
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			// t.Parallel()

			resp := &planmodifier.Int32Response{
				PlanValue: testCase.request.PlanValue,
			}

			x := int32planmodifier.DefaultFunc(func(_ context.Context, _ planmodifier.Int32Request, resp *int32planmodifier.DefaultFuncResponse) {
				resp.Value = expectedValue
			})

			int32planmodifier.SetDefaultFunc(x).PlanModifyInt32(context.Background(), testCase.request, resp)

			if diff := cmp.Diff(testCase.expected, resp); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
