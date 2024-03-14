/*
Copyright (c) 2024 Dell Inc., or its subsidiaries. All Rights Reserved.

Licensed under the Mozilla Public License Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://mozilla.org/MPL/2.0/


Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
terraform {
  required_providers {
    apex = {
      #version = "~> 0.0.0"
      source  = "dell/apex"
    }
  }
}

variable "HOST" {
  type = string
  default = "https://cirrus-premier-dev8.apex-dev2.us.dell.com"
}

variable "JWT_TOKEN" {
  type = string
  default = "eyJhbGciOiJSUzI1NiIsImtpZCI6InN0LTE1NDYzMDA4MDAiLCJ0eXAiOiJKV1QifQ.eyJDQlN0YXR1cyI6IlN1Y2Nlc3MiLCJsbSI6IiIsImF1dGhzcmMiOiJQRklESFVCIiwiUFlBQ0lEIjoiMjMxMjIxODAxIiwiUFlBTElEIjoiMjMxMjc2NjI4IiwiUFlBTE4iOiJjaXJydXMtZGV2LXVzZXItMkBwcmVtaWVydGVzdC5jb20iLCJQWUFMVFAiOiJEZWxsV2ViQWNjb3VudCIsIlBZU1IiOiJQRklESFVCIiwiUFlTVCI6IkF1dGhlbnRpY2F0ZWQiLCJQWUFUWSI6Ikdsb2JhbCIsIlBZRUlEIjoiY2lycnVzLWRldi11c2VyLTJAcHJlbWllcnRlc3QuY29tIiwiUFlFTVZGIjoiVHJ1ZSIsIlBZRk4iOiJkZXYiLCJQWUxOIjoidXNlcjIiLCJQWVBJRCI6Ijk2NzM5YzM4LWZhMmMtNDAxNS1hMTg3LTBkMjJjNzU3ZTkwNyIsIlBZU0xUIjoiMzAiLCJQWUlJRCI6IjJjODAxZGNiLWE3NWMtNGUzNS1hZGFjLTQ2ZTdiYjA1MDdiMiIsIlBZU0lEIjoiM2QwOWEwNmYtOTRmMy00NTE5LTg3ZDItNDJiMDA4NWNiMWFhIiwiUFlVTiI6ImNpcnJ1cy1kZXYtdXNlci0yQHByZW1pZXJ0ZXN0LmNvbSIsIm1lbWJlcm9mIjoiY2lycnVzX3Rlc3RfZ3JvdXAiLCJlbnYiOiJHMi1TUjIiLCJjYyI6IlVTIiwiZ3JwRmlsdGVyIjoiMmZhYjI0ZWQtMmI0MC00MGQ2LWFmMjItMjdkZmZlYTg1NTI1IiwiZXh0SWRwSWQiOiI3NTVjYmU4ZjExMWI0ZmZmYWMxYzU5ODY3ZWNiMzUwNSIsIklIIjoiUEZJREhVQiIsIkNIIjoiZTA4MWMwMDEtNDNiMy00ODEzLThjOTUtYmQzMzA3MGQwMzIyK0EiLCJQSCI6Ik1hbmFnZWRTZXNzaW9uIiwiaWF0IjoxNzA5MzI5Mjk0LCJzc2lkIjoiMGM0YmViYWYtM2QxNS00OTg0LTkzZjEtYjY2ZDA1MjgzMGQwIiwiZXhwIjoxNzA5MzMxMDk0LCJpc3MiOiJodHRwOi8vd3d3LmRlbGwuY29tL2lkZW50aXR5IiwibmJmIjoxNzA5MzI5Mjk0LCJwZXIiOiJ0cnVlIiwiYXVkIjoid3d3LmRlbGwuY29tIn0.Mx7Ncjg0zKpMs3FogaCfIRuec7kaukHUHW_f-8IbHh57ieb30-vDHuUwsbj7HZi6IYIecmS0MFal08oaPCb7d034NaHWvOHtPDoF37iMz4ViFwxOgRZJP6bXaFXdIwbMRmYMcu8qAa3jKT-Wdc8egH-V86EdDh9rAq3eCRShlvLh-mtgTLFo7Tyhfe-7z21ZUi1T-1CjLHhWV_ZDCaaSc9iew6tc1QFby5E7AXdpwQW_HlZ20LUbGaGSdccs68rMviAEP0J3s9ZFKdgm-_-zFibr2N5rWwaWZUDOJxukSORnjwh-IGlmVLZZZu_8Mh40GZ103z_tOotL5VWHoNJonA"
}

provider "apex" {
  host  = var.HOST
  token = var.JWT_TOKEN
}

data "apex_navigator_volumes" "example" {}

output "examples_volumes" {
  value = data.apex_navigator_volumes.example
}
