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

module "hello_world" {
  source = "../"
}

module "hello_john" {
  source = "../"
  name   = "John Doe"
}
