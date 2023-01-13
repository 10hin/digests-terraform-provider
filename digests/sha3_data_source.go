package digests

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"golang.org/x/crypto/sha3"
)

func NewSHA3_224DataSource() datasource.DataSource {
	return &hashDataSourceBase{
		typeNameSuffix: "_sha3_224",
		algorithmName:  "sha3_224",
		hash: func(input []byte) []byte {
			hash := sha3.New224()
			hash.Write(input)

			return hash.Sum([]byte{})

		},
	}
}

func NewSHA3_256DataSource() datasource.DataSource {
	return &hashDataSourceBase{
		typeNameSuffix: "_sha3_256",
		algorithmName:  "sha3_256",
		hash: func(input []byte) []byte {
			hash := sha3.New256()
			hash.Write(input)

			return hash.Sum([]byte{})

		},
	}
}

func NewSHA3_384DataSource() datasource.DataSource {
	return &hashDataSourceBase{
		typeNameSuffix: "_sha3_384",
		algorithmName:  "sha3_384",
		hash: func(input []byte) []byte {
			hash := sha3.New384()
			hash.Write(input)

			return hash.Sum([]byte{})

		},
	}
}

func NewSHA3_512DataSource() datasource.DataSource {
	return &hashDataSourceBase{
		typeNameSuffix: "_sha3_512",
		algorithmName:  "sha3_512",
		hash: func(input []byte) []byte {
			hash := sha3.New512()
			hash.Write(input)

			return hash.Sum([]byte{})

		},
	}
}
