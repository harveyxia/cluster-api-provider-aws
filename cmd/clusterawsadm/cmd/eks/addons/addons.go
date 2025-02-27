/*
Copyright 2021 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package addons

import "github.com/spf13/cobra"

// RootCmd is EKS addons root CLI command.
func RootCmd() *cobra.Command {
	newCmd := &cobra.Command{
		Use:   "addons",
		Short: "Commands related to EKS addons",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
	newCmd.AddCommand(listAvailableCmd())
	newCmd.AddCommand(listInstalledCmd())

	return newCmd
}
