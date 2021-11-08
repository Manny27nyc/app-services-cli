package serviceaccount

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/serviceaccount/create"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/serviceaccount/delete"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/serviceaccount/describe"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/serviceaccount/list"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/serviceaccount/resetcredentials"
	"github.com/spf13/cobra"
)

func NewConfigCommand(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Generates and manages services configurations",
		Long: heredoc.Doc(`
		Generates and manages services configurations that can be used 
		to configure your application deployments, CI/CD instances to be able to
		connect with all managed services.
		`),
		Example: ``,
		Args:    cobra.ExactArgs(1),
	}

	cmd.AddCommand(
		create.NewCreateCommand(f),
		list.NewListCommand(f),
		delete.NewDeleteCommand(f),
		resetcredentials.NewResetCredentialsCommand(f),
		describe.NewDescribeCommand(f),
	)

	return cmd
}
