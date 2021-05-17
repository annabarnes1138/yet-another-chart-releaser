// Copyright © 2021 Steven Barnes <stecky@users.noreply.github.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stecky/yet-another-chart-releaser/pkg/config"
	"github.com/stecky/yet-another-chart-releaser/pkg/git"
	"github.com/stecky/yet-another-chart-releaser/pkg/github"
	"github.com/stecky/yet-another-chart-releaser/pkg/releaser"
)

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload Helm chart packages to GitHub Releases",
	Long:  `Upload Helm chart packages to GitHub Releases`,
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := config.LoadConfiguration(cfgFile, cmd, getRequiredUploadArgs())
		if err != nil {
			return err
		}
		ghc := github.NewClient(config.Owner, config.GitRepo, config.Token, config.GitBaseURL, config.GitUploadURL)
		releaser := releaser.NewReleaser(config, ghc, &git.Git{})
		return releaser.CreateReleases()
	},
}

func getRequiredUploadArgs() []string {
	return []string{"owner", "git-repo", "token"}
}

func init() {
	rootCmd.AddCommand(uploadCmd)
	uploadCmd.Flags().StringP("owner", "o", "", "GitHub username or organization")
	uploadCmd.Flags().StringP("git-repo", "r", "", "GitHub repository")
	uploadCmd.Flags().StringP("package-path", "p", ".yacr-release-packages", "Path to directory with chart packages")
	uploadCmd.Flags().StringP("token", "t", "", "GitHub Auth Token")
	uploadCmd.Flags().StringP("git-base-url", "b", "https://api.github.com/", "GitHub Base URL (only needed for private GitHub)")
	uploadCmd.Flags().StringP("git-upload-url", "u", "https://uploads.github.com/", "GitHub Upload URL (only needed for private GitHub)")
	uploadCmd.Flags().StringP("commit", "c", "", "Target commit for release")
	uploadCmd.Flags().Bool("skip-existing", false, "Skip upload if release exists")
	uploadCmd.Flags().String("release-name-template", "{{ .Name }}-{{ .Version }}", "Go template for computing release names, using chart metadata")
	uploadCmd.Flags().String("pages-branch", "gh-pages", "The GitHub pages branch")
	uploadCmd.Flags().String("remote", "origin", "The Git remote used when creating a local worktree for the GitHub Pages branch")
	uploadCmd.Flags().Bool("push", false, "Push index.yaml to the GitHub Pages branch (must not be set if --pr is set)")
	uploadCmd.Flags().Bool("pr", false, "Create a pull request for index.yaml against the GitHub Pages branch (must not be set if --push is set)")
	uploadCmd.Flags().Bool("packages-with-index", false, "Save a copy of the package files to the GitHub Pages branch")
}
