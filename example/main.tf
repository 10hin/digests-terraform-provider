terraform {
  required_providers {
    digests = {
      source = "local/10hin/digests"
    }
    local = {
      source = "hashicorp/local"
    }
  }
}

provider "digests" {
  attr_string = "AttributeWithTypeString"
  num_attr = 10
  sensitive_attr = "sensitive"
}

provider "local" {}

data "digests_md5" "hoge" {
  input = "hoge"
}

resource "local_file" "output" {
  content = <<-EOT
  base64: ${data.digests_md5.hoge.base64}
  hex: ${data.digests_md5.hoge.hex}
  EOT
  filename = "${path.module}/output.txt"
}
