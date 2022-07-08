package awscmd

import (
	commands "github.com/SlalomBuild/fusion/internal/commands"
	aws "github.com/SlalomBuild/fusion/templates/aws"
)

// NewAPIGatewayv2Cmd creates a new api_gatewayv2
type NewAPIGatewayCmd struct {
	Globals
	*aws.APIGateway `embed:""`
}

func (cmd *NewAPIGatewayCmd) Run(ctx *commands.Context) error {
	return cmd.Render(ctx.Output, ctx.NoColor.Bool())
}
