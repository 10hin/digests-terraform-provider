package digests

import (
	"crypto/sha1"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

func NewSHA1DataSource() datasource.DataSource {
	return &hashDataSourceBase{
		typeNameSuffix: "_sha1",
		algorithmName:  "SHA1",
		hash: func(input []byte) []byte {
			hash := sha1.New()
			hash.Write(input)

			return hash.Sum([]byte{})

		},
	}
}
