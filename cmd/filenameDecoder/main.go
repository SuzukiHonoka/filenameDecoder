package filenameDecoder

import (
	"filenameDecoder/pkg/config"
	"filenameDecoder/pkg/rename"
)

func Main() {
	cfg := config.NewConfigFromFlags()
	if err := rename.BatchRename(cfg); err != nil {
		panic(err)
	}
}
