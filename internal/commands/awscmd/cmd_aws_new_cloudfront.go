package awscmd

import (
	commands "github.com/SlalomBuild/fusion/internal/commands"
	aws "github.com/SlalomBuild/fusion/templates/aws"
)

// NewCloudfrontCmd creates a new CloudFront
type NewCloudfrontCmd struct {
	Globals
	*aws.CloudfrontDistribution `embed:""`
}

func (cmd *NewCloudfrontCmd) Run(ctx *commands.Context) error {
	return cmd.Render(ctx.Output, ctx.NoColor.Bool())
}
