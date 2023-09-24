package remote

import (
	"github.com/pulumi/pulumi-go-provider/infer"
)

type BaseOutputs struct {
	MD5Sum string `pulumi:"md5sum"`
}

// Implementing Annotate lets you provide descriptions and default values for fields and they will
// be visible in the provider's schema and the generated SDKs.
func (c *BaseOutputs) Annotate(a infer.Annotator) {
	a.Describe(&c.MD5Sum, "The md5sum of the uploaded file")
}
