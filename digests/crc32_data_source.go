package digests

import (
	crc32var "github.com/10hin/crcvariants/crc32"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"hash/crc32"
)

func NewCRC32IEEEDataSource() datasource.DataSource {
	return &hashDataSourceBase{
		typeNameSuffix: "_crc32_ieee",
		algorithmName:  "CRC32_IEEE",
		hash: func(input []byte) []byte {
			hash := crc32.NewIEEE()
			hash.Write(input)

			return hash.Sum([]byte{})

		},
	}
}

func NewCRC32CastagnoliDataSource() datasource.DataSource {
	return &hashDataSourceBase{
		typeNameSuffix: "_crc32_castagnoli",
		algorithmName:  "CRC32_Castagnoli",
		hash: func(input []byte) []byte {
			hash := crc32.New(crc32.MakeTable(crc32.Castagnoli))
			hash.Write(input)

			return hash.Sum([]byte{})

		},
	}
}

func NewCRC32KoopmanDataSource() datasource.DataSource {
	return &hashDataSourceBase{
		typeNameSuffix: "_crc32_koopman",
		algorithmName:  "CRC32_Koopman",
		hash: func(input []byte) []byte {
			hash := crc32.New(crc32.MakeTable(crc32.Koopman))
			hash.Write(input)

			return hash.Sum([]byte{})

		},
	}
}

func NewCRC32PHPDataSource() datasource.DataSource {
	return &hashDataSourceBase{
		typeNameSuffix: "_crc32_php",
		algorithmName:  "CRC32_PHP",
		hash: func(input []byte) []byte {
			hash := crc32var.NewPHP()
			hash.Write(input)

			return hash.Sum([]byte{})

		},
	}
}
