package getlogstream

import (
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/needsauth"
	"github.com/1Password/shell-plugins/sdk/schema"
	"github.com/1Password/shell-plugins/sdk/schema/credname"
)

func GetLogstreamCLI() schema.Executable {
	return schema.Executable{
		Name:    "GetLogstream CLI",
		Runs:    []string{"getlogstream"},
		DocsURL: sdk.URL("https://aws.amazon.com/cli/"),
		NeedsAuth: needsauth.IfAll(
			needsauth.NotForHelpOrVersion(),
			needsauth.NotWithoutArgs(),
		),
		Uses: []schema.CredentialUsage{
			{
				Name:        credname.AccessKey,
				Provisioner: CLIProvisioner{},
			},
		},
	}
}
