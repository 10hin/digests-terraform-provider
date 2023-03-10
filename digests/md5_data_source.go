package digests

import (
	"crypto/md5"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

func NewMD5DataSource() datasource.DataSource {
	return &hashDataSourceBase{
		typeNameSuffix: "_md5",
		algorithmName:  "MD5",
		hash: func(input []byte) []byte {
			hash := md5.New()
			hash.Write(input)

			return hash.Sum([]byte{})

		},
	}
}
