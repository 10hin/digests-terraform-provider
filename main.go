package main

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"

	"github.com/10hin/digests-terraform-provider/digests"
)

func main() {
	_ = providerserver.Serve(context.Background(), digests.New, providerserver.ServeOpts{
		Address: "example.com/10hin/digests",
	})
}
