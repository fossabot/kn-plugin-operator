// Copyright 2022 The Knative Authors
//
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

package configure

import (
	"github.com/spf13/cobra"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc" // from https://github.com/kubernetes/client-go/issues/345
	"knative.dev/kn-plugin-operator/pkg"
)

// NewConfigureCommand represents the configure commands for Knative Serving or eventing
func NewConfigureCommand(p *pkg.OperatorParams) *cobra.Command {
	var configureCmd = &cobra.Command{
		Use:   "configure",
		Short: "Configure the Knative Serving or Eventing",
		Example: `
  # Configure the Knative Serving or Eventing
  kn operation configure resource --component serving --requestMemory 999M --namespace knative-serving`,
	}

	configureCmd.AddCommand(newResourcesCommand(p))
	return configureCmd
}