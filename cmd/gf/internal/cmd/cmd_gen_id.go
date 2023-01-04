package cmd

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/yitter/idgenerator-go/idgen"
)

type (
	cGenId      struct{}
	cGenIdInput struct {
		g.Meta `name:"id" brief:"Generate unique ID" usage:"gf gen id" `
	}
)

type CGenIdOutput struct{}

func (c cGenId) id(_ context.Context, _ cGenIdInput) (out *CGenIdOutput, err error) {
	fmt.Printf("Generate unique ID: %v", idgen.NextId())
	return
}
