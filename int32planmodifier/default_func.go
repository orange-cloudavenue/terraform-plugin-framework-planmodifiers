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

import "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"

// SetDefaultFunc returns a plan modifier that conditionally requires
// resource replacement if:
//
//   - The resource is planned for update.
//   - The plan and state values are not equal.
//   - The plan or state values are not null or known
func SetDefaultFunc(f DefaultFunc) planmodifier.Int32 {
	return setDefaultFunc(
		f,
		"Set default value",
		"Set default value",
	)
}
