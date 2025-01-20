package getlogstream

import (
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/schema"
)

func New() schema.Plugin {
	return schema.Plugin{
		Name: "getlogstream",
		Platform: schema.PlatformInfo{
			Name:     "GetLogstream",
			Homepage: sdk.URL("https://aws.amazon.com/"),
		},
		Credentials: []schema.CredentialType{
			AccessKey(),
		},
		Executables: []schema.Executable{
			GetLogstreamCLI(),
		},
	}
}
