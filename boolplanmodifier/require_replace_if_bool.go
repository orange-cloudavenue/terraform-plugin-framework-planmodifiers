/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

package boolplanmodifier

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

/*
RequireReplaceIfBool

returns a plan modifier that requires replacement
if the attribute value is equal to the excepted value.
*/
func RequireReplaceIfBool(path path.Path, exceptedValue bool) planmodifier.Bool {
	description := fmt.Sprintf("Attribute require replacement if `%s` is `%v`", path.String(), exceptedValue)
	return boolplanmodifier.RequiresReplaceIf(boolplanmodifier.RequiresReplaceIfFunc(func(ctx context.Context, req planmodifier.BoolRequest, resp *boolplanmodifier.RequiresReplaceIfFuncResponse) {
		boolValue := &types.Bool{}

		resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, path, boolValue)...)
		if resp.Diagnostics.HasError() {
			return
		}

		if boolValue.ValueBool() == exceptedValue {
			resp.RequiresReplace = true
		}
	}), description, description)
}
