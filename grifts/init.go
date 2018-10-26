package grifts

import (
	"github.com/cileonard/lrn/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
