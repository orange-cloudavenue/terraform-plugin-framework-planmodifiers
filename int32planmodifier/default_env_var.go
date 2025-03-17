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
	"fmt"
	"os"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

// SetDefaultEnvVar returns a plan modifier that conditionally requires
// resource replacement if:
//
//   - The resource is planned for update.
//   - The plan and state values are not equal.
//   - The plan or state values are not null or known
func SetDefaultEnvVar(envVar string) planmodifier.Int32 {
	return setDefaultFunc(
		func(_ context.Context, _ planmodifier.Int32Request, resp *DefaultFuncResponse) {
			v := os.Getenv(envVar)
			if v != "" {
				// string to int32
				parsedInt, err := strconv.ParseInt(v, 10, 32)
				if err != nil {
					resp.Diagnostics.AddError("Environment variable set but is not Int32", fmt.Sprintf("The environment variable %s is set but is not a Int32", envVar))
					return
				}
				resp.Value = int32(parsedInt)
			} else {
				resp.Diagnostics.AddError("Environment variable not set", fmt.Sprintf("The environment variable %s is not set", envVar))
			}
		},
		"Set default value from environment variable",
		"Set default value from environment variable",
	)
}
