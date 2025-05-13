/*
 * SPDX-FileCopyrightText: Copyright (c) 2025 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

// Package int32planmodifier provides a plan modifier for int32 values.
package int32planmodifier

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

// SetDefault
//
// SetDefault returns a plan modifier that sets the plan value to the
// provided value if the following conditions are met:
//
//   - The resource is planned for update.
//   - The plan and state values are not equal.
//   - The plan or state values are not null or known
func SetDefault(i int32) planmodifier.Int32 {
	return setDefaultFunc(
		func(_ context.Context, _ planmodifier.Int32Request, resp *DefaultFuncResponse) {
			resp.Value = i
		},
		"Set default value",
		"Set default value",
	)
}
