package digests

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"hash/adler32"
)

func NewAdler32DataSource() datasource.DataSource {
	return &hashDataSourceBase{
		typeNameSuffix: "_adler32",
		algorithmName:  "Adler32",
		hash: func(input []byte) []byte {
			hash := adler32.New()
			hash.Write(input)

			return hash.Sum([]byte{})

		},
	}
}
