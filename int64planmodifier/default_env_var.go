/*
 * SPDX-FileCopyrightText: Copyright (c) Orange Business Services SA
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at <https://www.mozilla.org/en-US/MPL/2.0/>
 * or see the "LICENSE" file for more details.
 */

// Package int64planmodifier provides a plan modifier for int64 values.
package int64planmodifier

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
func SetDefaultEnvVar(envVar string) planmodifier.Int64 {
	return setDefaultFunc(
		func(_ context.Context, _ planmodifier.Int64Request, resp *DefaultFuncResponse) {
			v := os.Getenv(envVar)
			if v != "" {
				// string to int64
				i, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					resp.Diagnostics.AddError("Environment variable set but is not Int64", fmt.Sprintf("The environment variable %s is set but is not a Int64", envVar))
					return
				}
				resp.Value = i
			} else {
				resp.Diagnostics.AddError("Environment variable not set", fmt.Sprintf("The environment variable %s is not set", envVar))
			}
		},
		"Set default value from environment variable",
		"Set default value from environment variable",
	)
}
