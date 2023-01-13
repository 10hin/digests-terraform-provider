package digests

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"golang.org/x/crypto/md4"
)

func NewMD4DataSource() datasource.DataSource {
	return &hashDataSourceBase{
		typeNameSuffix: "_md4",
		algorithmName:  "MD4",
		hash: func(input []byte) []byte {
			hash := md4.New()
			hash.Write(input)

			return hash.Sum([]byte{})

		},
	}
}
