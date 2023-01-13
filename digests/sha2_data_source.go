package digests

import (
	"crypto/sha256"
	"crypto/sha512"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

func NewSHA224DataSource() datasource.DataSource {
	return &hashDataSourceBase{
		typeNameSuffix: "_sha224",
		algorithmName:  "SHA224",
		hash: func(input []byte) []byte {
			hash := sha256.New224()
			hash.Write(input)

			return hash.Sum([]byte{})

		},
	}
}

func NewSHA256DataSource() datasource.DataSource {
	return &hashDataSourceBase{
		typeNameSuffix: "_sha256",
		algorithmName:  "SHA256",
		hash: func(input []byte) []byte {
			hash := sha256.New()
			hash.Write(input)

			return hash.Sum([]byte{})

		},
	}
}

func NewSHA384DataSource() datasource.DataSource {
	return &hashDataSourceBase{
		typeNameSuffix: "_sha384",
		algorithmName:  "SHA384",
		hash: func(input []byte) []byte {
			hash := sha512.New384()
			hash.Write(input)

			return hash.Sum([]byte{})

		},
	}
}

func NewSHA512DataSource() datasource.DataSource {
	return &hashDataSourceBase{
		typeNameSuffix: "_sha512",
		algorithmName:  "SHA512",
		hash: func(input []byte) []byte {
			hash := sha512.New()
			hash.Write(input)

			return hash.Sum([]byte{})

		},
	}
}

func NewSHA512_224DataSource() datasource.DataSource {
	return &hashDataSourceBase{
		typeNameSuffix: "_sha512_224",
		algorithmName:  "SHA512_224",
		hash: func(input []byte) []byte {
			hash := sha512.New512_224()
			hash.Write(input)

			return hash.Sum([]byte{})

		},
	}
}

func NewSHA512_256DataSource() datasource.DataSource {
	return &hashDataSourceBase{
		typeNameSuffix: "_sha512_256",
		algorithmName:  "SHA512_256",
		hash: func(input []byte) []byte {
			hash := sha512.New512_256()
			hash.Write(input)

			return hash.Sum([]byte{})

		},
	}
}
