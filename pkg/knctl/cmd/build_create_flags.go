/*
Copyright 2018 The Knative Authors

Licensed under the Apache License, Open 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	ctlbuild "github.com/cppforlife/knctl/pkg/knctl/build"
	"github.com/spf13/cobra"
)

type BuildCreateFlags struct {
	ctlbuild.BuildSpecOpts
}

type BuildCreateFlagsOpts struct {
	Optional bool
	NoImage  bool
}

func (s *BuildCreateFlags) Set(cmd *cobra.Command, opts BuildCreateFlagsOpts) {
	cmd.Flags().StringVar(&s.GitURL, "git-url", "", "Set Git URL")
	if !opts.Optional {
		cmd.MarkFlagRequired("git-url")
	}

	cmd.Flags().StringVar(&s.GitRevision, "git-revision", "", "Set Git revision (Examples: https://git-scm.com/docs/gitrevisions#_specifying_revisions)")
	if !opts.Optional {
		cmd.MarkFlagRequired("git-revision")
	}

	cmd.Flags().StringVar(&s.ServiceAccountName, "service-account", "", "Set service account name for building") // TODO separate

	if !opts.NoImage {
		cmd.Flags().StringVarP(&s.Image, "image", "i", "", "Set image URL")
		cmd.MarkFlagRequired("image")
	}
}
