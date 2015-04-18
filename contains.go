package concat

import (
	"strings"

	"github.com/kelseyhightower/confd/resource/template"
)

func init() {
	template.RegisterPlugin("contains", new(ConcatPlugin))
}

type ConcatPlugin struct{}

func (p *ConcatPlugin) Function() interface{} {
	return strings.Contains
}
