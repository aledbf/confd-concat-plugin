package concat

import (
	"testing"

	"github.com/kelseyhightower/confd/resource/template"
	"github.com/kelseyhightower/memkv"
)

var pluginTests = []template.PluginTest{
	template.PluginTest{
		Desc: "contains, should return true",
		Toml: `
[template]
src = "test.conf.tmpl"
dest = "./tmp/test.conf"
keys = [
    "/test/key",
]
`,
		Tmpl: `
key: {{contains "example string A" (getv "/test/key")}}
`,
		Expected: `
key: true
`,
		UpdateStore: func(tr memkv.Store) {
			tr.Set("/test/key", "A")
		},
	},
	template.PluginTest{
		Desc: "contains, should return false",
		Toml: `
[template]
src = "test.conf.tmpl"
dest = "./tmp/test.conf"
keys = [
    "/test/key",
]
`,
		Tmpl: `
key: {{contains "example string B" "/test/key"}}
`,
		Expected: `
key: false
`,
		UpdateStore: func(tr memkv.Store) {
			tr.Set("/test/key", "A")
		},
	},
}

// TestPlugin runs all tests in plugin contains
func TestPlugin(t *testing.T) {
	template.TestPlugin(pluginTests, t)
}
