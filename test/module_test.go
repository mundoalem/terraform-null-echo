// This file is part of template-terraform-module.
//
// template-terraform-module is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// template-terraform-module is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with template-terraform-module. If not, see <https://www.gnu.org/licenses/>.

package test

import (
	"regexp"
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

	outputHelloJohnDoe := terraform.Output(t, terraformOptions, "hello_john")
	outputHelloWorld := terraform.Output(t, terraformOptions, "hello_world")
	outputNow := terraform.Output(t, terraformOptions, "now")
	re, _ := regexp.Compile("^[0-9]{4}-[0-9]{2}-[0-9]{2}T[0-9]{2}:[0-9]{2}:[0-9]{2}Z$")

	assert.Equal(t, "Hello, John Doe!", outputHelloJohnDoe)
	assert.Equal(t, "Hello, World!", outputHelloWorld)
	assert.Equal(t, re.MatchString(outputNow), true)
}
