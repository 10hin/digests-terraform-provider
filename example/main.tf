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

provider "digests" {}

provider "local" {}

locals {
  input = "example"
}

data "digests_md4" "example" {
  input = local.input
}

data "digests_md5" "example" {
  input = local.input
}

data "digests_sha1" "example" {
  input = local.input
}

data "digests_sha224" "example" {
  input = local.input
}

data "digests_sha256" "example" {
  input = local.input
}

data "digests_sha384" "example" {
  input = local.input
}

data "digests_sha512" "example" {
  input = local.input
}

data "digests_sha512_224" "example" {
  input = local.input
}

data "digests_sha512_256" "example" {
  input = local.input
}

data "digests_sha3_224" "example" {
  input = local.input
}

data "digests_sha3_256" "example" {
  input = local.input
}

data "digests_sha3_384" "example" {
  input = local.input
}

data "digests_sha3_512" "example" {
  input = local.input
}

data "digests_crc32_ieee" "example" {
  input = local.input
}

data "digests_crc32_castagnoli" "example" {
  input = local.input
}

data "digests_crc32_koopman" "example" {
  input = local.input
}

data "digests_crc32_php" "example" {
  input = local.input
}

data "digests_adler32" "example" {
  input = local.input
}

resource "local_file" "output" {
  content = <<-EOT
  Input: ${local.input}
  Algorithms:
    MD4:              ${data.digests_md4.example.hex}
    MD5:              ${data.digests_md5.example.hex}
    SHA1:             ${data.digests_sha1.example.hex}
    SHA224:           ${data.digests_sha224.example.hex}
    SHA256:           ${data.digests_sha256.example.hex}
    SHA384:           ${data.digests_sha384.example.hex}
    SHA512:           ${data.digests_sha512.example.hex}
    SHA512_224:       ${data.digests_sha512_224.example.hex}
    SHA512_256:       ${data.digests_sha512_256.example.hex}
    SHA3_224:         ${data.digests_sha3_224.example.hex}
    SHA3_256:         ${data.digests_sha3_256.example.hex}
    SHA3_384:         ${data.digests_sha3_384.example.hex}
    SHA3_512:         ${data.digests_sha3_512.example.hex}
    CRC32_IEEE:       ${data.digests_crc32_ieee.example.hex}
    CRC32_Castagnoli: ${data.digests_crc32_castagnoli.example.hex}
    CRC32_Koopman:    ${data.digests_crc32_koopman.example.hex}
    CRC32_PHP:        ${data.digests_crc32_php.example.hex}
    Adler32:          ${data.digests_adler32.example.hex}
  Output variants(MD5):
    base64: ${data.digests_md5.example.base64}
    hex: ${data.digests_md5.example.hex}
  EOT
  filename = "${path.module}/output.txt"
}
