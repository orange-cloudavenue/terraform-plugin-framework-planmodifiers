/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

// Package stringplanmodifier provides a plan modifier for string values.
package stringplanmodifier

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

func ToLower() planmodifier.String {
	return setChangeStringFunc(
		func(_ context.Context, req planmodifier.StringRequest, resp *StringChangeFuncResponse) {
			resp.Value = types.StringValue(strings.ToLower(req.ConfigValue.ValueString()))
		},
		"Force to lower case",
		"Force to lower case",
	)
}
