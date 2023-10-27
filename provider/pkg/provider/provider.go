// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package provider

import (
	"strings"

	"github.com/blang/semver"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi-go-provider/middleware/schema"

	"github.com/pulumi/pulumi-go-provider/integration"
	"github.com/spigell/pulumi-file/provider/pkg/provider/remote"
)

const (
	Name = "file"
)

func NewProvider() p.Provider {
	return infer.Provider(infer.Options{
		// This is the metadata for the provider
		Metadata: schema.Metadata{
			DisplayName: "File",
			Description: "The Pulumi File Provider enables you to create files either locally or remotely as part of the Pulumi resource model.",
			Keywords: []string{
				"pulumi",
				"file",
				"category/utility",
				"kind/native",
			},
			Homepage:   "https://pulumi.com",
			License:    "Apache-2.0",
			Repository: "https://github.com/spigell/pulumi-file",
			Publisher:  "spigell",
                        PluginDownloadURL: "github://api.github.com/spigell/pulumi-file",
			LanguageMap: map[string]any{
				"csharp": map[string]any{
					"packageReferences": map[string]string{
						"Pulumi": "3.*",
					},
				},
				"go": map[string]any{
					"generateResourceContainerTypes": true,
					"importBasePath":                 "github.com/spigell/pulumi-file/sdk/go/file",
				},
				"nodejs": map[string]any{
					"dependencies": map[string]string{
						"@pulumi/pulumi": "^3.0.0",
					},
				},
				"python": map[string]any{
					"requires": map[string]string{
						"pulumi": ">=3.0.0,<4.0.0",
					},
					"pyproject": map[string]bool{
						"enabled": true,
					},
				},
				"java": map[string]any{
					"buildFiles":                      "gradle",
					"gradleNexusPublishPluginVersion": "1.1.0",
					"dependencies": map[string]any{
						"com.pulumi:pulumi":               "0.6.0",
						"com.google.code.gson:gson":       "2.8.9",
						"com.google.code.findbugs:jsr305": "3.0.2",
					},
				},
			},
		},
		// A list of `infer.Resource` that are provided by the provider.
		Resources: []infer.InferredResource{
			// The Command resource implementation is commented extensively for new pulumi-go-provider developers.
			infer.Resource[
				// 1. This type is an interface that implements the logic for the Resource
				//    these methods include `Create`, `Update`, `Delete`, and `WireDependencies`.
				//    `WireDependencies` should be implemented to preserve the secretness of an input
				*remote.File,
				// 2. The type of the Inputs/Arguments to supply to the Resource.
				remote.FileInputs,
				// 3. The type of the Output/Properties/Fields of a created Resource.
				remote.FileOutputs,
			](),
		},
	})
}

func Schema(version string) (string, error) {
	version = strings.TrimPrefix(version, "v")
	s, err := integration.NewServer(Name, semver.MustParse(version), NewProvider()).
		GetSchema(p.GetSchemaRequest{})
	return s.Schema, err
}
