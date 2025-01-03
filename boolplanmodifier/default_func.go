/*
 * SPDX-FileCopyrightText: Copyright (c) Orange Business Services SA
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at <https://www.mozilla.org/en-US/MPL/2.0/>
 * or see the "LICENSE" file for more details.
 */

// Package boolplanmodifier provides a plan modifier for boolean values.
package boolplanmodifier

import "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"

// SetDefaultFunc returns a plan modifier that conditionally requires
// resource replacement if:
//
//   - The resource is planned for update.
//   - The plan and state values are not equal.
//   - The plan or state values are not null or known
func SetDefaultFunc(f DefaultFunc) planmodifier.Bool {
	return setDefaultFunc(
		f,
		"Set default value",
		"Set default value",
	)
}
