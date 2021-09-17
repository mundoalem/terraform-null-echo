// This file is part of terraform-null-echo.
//
// terraform-null-echo is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// terraform-null-echo is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with terraform-null-echo. If not, see <https://www.gnu.org/licenses/>.

package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformModule(t *testing.T) {
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples",
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	outputEcho := terraform.Output(t, terraformOptions, "echo")
	outputEchoDefault := terraform.Output(t, terraformOptions, "echo_default")

	assert.Equal(t, outputEcho, "Hello darkness, my old friend.")
	assert.Equal(t, outputEchoDefault, "Hello, World!")
}
