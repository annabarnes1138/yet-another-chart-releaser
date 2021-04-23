<h1 align="center">
  <img src="https://github.com/stecky/yet-another-chart-releaser/blob/main/.github/yacr_logo.png" alt="yacr - Chart Releaser">
  <br>
  Yet Another Helm Chart Releaser
</h1>
<h4 align="center">Be sure to :star:  this repo so that you can keep up to date on any progress!</h4>
<div align="center">
  <h4>
    <a href="https://opensource.org/licenses/Apache-2.0">
    <img src="https://img.shields.io/badge/License-Apache%202.0-blue.svg"
         alt="License"/></a>
    <a href="https://goreportcard.com/report/github.com/stecky/yet-another-chart-releaser">
    <img src="https://goreportcard.com/badge/github.com/stecky/yet-another-chart-releaser"
         alt="Go Report Card"/></a>
    <a href="https://github.com/stecky/yet-another-chart-releaser/commits/main">
    <img src="https://img.shields.io/github/last-commit/stecky/yet-another-chart-releaser.svg?style=plasticr"
         alt="Last commit"/></a>
    <a href="https://github.com/stecky/yet-another-chart-releaser/issues">
    <img src="https://img.shields.io/github/issues-raw/stecky/yet-another-chart-releaser.svg?style=plasticr"
         alt="Issues"/></a>
    <a href="https://github.com/stecky/yet-another-chart-releaser/pulls">
    <img src="https://img.shields.io/github/issues-pr-raw/stecky/yet-another-chart-releaser.svg?style=plasticr"
         alt="Pull Requests"/></a>
    <a href="https://github.com/stecky/yet-another-chart-releaser/stargazers">
    <img src="https://img.shields.io/github/stars/stecky/yet-another-chart-releaser.svg?style=plasticr"
        alt="Stars"/></a>
    <a href="https://www.buymeacoffee.com/stecky">
    <img src="https://img.shields.io/badge/Donate-Buy%20me%20a%20coffee-orange?style=plasticr"
        alt="Buy Me a Coffee"/></a>
  </h4>
</div>
<p><font size="3">
<strong>yacr</strong> (pronounced like <em><a href="https://en.wiktionary.org/wiki/yacker">yacker</a></em>) is a tool designed to help GitHub repos self-host their own chart repos by adding Helm chart artifacts to GitHub Releases named for the chart version and then creating an `index.yaml` file for those releases that can be hosted on GitHub Pages (or elsewhere!).</font></p>

<p><font size="2">This repo is for the <a href="https://en.wikipedia.org/wiki/Command-line_interface">cli</a> tool itself. If you are looking for a <a href="https://docs.github.com/en/actions">GitHub Action</a> to add this tool to your workflow, it can be found at <a href="https://github.com/stecky/yet-another-chart-releaser-action">stecky/yet-another-chart-releaser-action</a> or on the <a href="https://github.com/marketplace/actions/helm-chart-releaser">GitHub Marketplace</a></font></p>

<div align="center">
  <h2>
  <a href="#installation">Installation</a> <span> | </span>
  <a href="#usage">Usage</a> <span> | </span>
  <a href="#configuration">Configuration</a> <span> | </span>
  <a href="#goals">Goals</a>
  </h2>
</div>

<p><font size="2">This repo is a clone of <a href="https://github.com/helm/chart-releaser">helm/chart-releaser</a>. 99% of the credit for this project goes to their team for the great work they did on chart-releaser. I simply needed some additional functionality and decided to create my own clone of it. If you want to see a list of enhancements that I plan to make, check out the <a href="#goals">goals</a> section of this readme.</font></p>
<hr>

## Installation

### Binaries (recommended)

Download your preferred asset from the [releases page](https://github.com/stecky/yet-another-chart-releaser/releases) and install manually.

### Homebrew

```console
$ brew tap stecky/tap
$ brew install yet-another-chart-releaser
```

### Go get (for contributing)

```console
$ # clone repo to some directory outside GOPATH
$ git clone https://github.com/stecky/yet-another-chart-releaser
$ cd yet-another-chart-releaser
$ go mod download
$ go install ./...
```

### Docker (for Continuous Integration)

Docker images are pushed to the [stecky/yet-another-chart-releaser](https://hub.docker.com/r/stecky/yet-another-chart-releaser/tags?page=1&ordering=last_updated) dockerhub registry. The Docker image is built on top of Alpine and its default entry-point is `yacr`. See the [Dockerfile](./Dockerfile) for more details.

## Usage

Currently, `yacr` can create GitHub Releases from a set of charts packaged up into a directory and create an `index.yaml` file for the chart repository from GitHub Releases.

```console
$ yacr --help
Create Helm chart repositories on GitHub Pages by uploading Chart packages
and Chart metadata to GitHub Releases and creating a suitable index file

Usage:
  yacr [command]

Available Commands:
  help        Help about any command
  index       Update Helm repo index.yaml for the given GitHub repo
  upload      Upload Helm chart packages to GitHub Releases
  package     Package Helm charts
  version     Print version information

Flags:
      --config string   Config file (default is $HOME/.yacr.yaml)
  -h, --help            help for yacr

Use "yacr [command] --help" for more information about a command.
```

### Create Helm Chart Packages from Helm Charts

Scans a path for Helm charts and creates packages.

```console
$ yacr package --help
This command packages a chart into a versioned chart archive file. If a path
is given, this will look at that path for a chart (which must contain a
Chart.yaml file) and then package that directory.


If you wish to use advanced packaging options such as creating signed
packages or updating chart dependencies please use "helm package" instead.

Usage:
  yacr package [CHART_PATH] [...] [flags]

Flags:
  -h, --help                     help for package
      --key string               Name of the key to use when signing
      --keyring string           Location of a public keyring (default "/Users/steven.barnes/.gnupg/pubring.gpg")
  -p, --package-path string      Path to directory with chart packages (default ".yacr-release-packages")
      --passphrase-file string   Location of a file which contains the passphrase for the signing key. Use '-' in order to read from stdin
      --sign                     Use a PGP private key to sign this package

Global Flags:
      --config string   Config file (default is $HOME/.yacr.yaml)
```

### Create GitHub Releases from Helm Chart Packages

Scans a path for Helm chart packages and creates releases in the specified GitHub repo uploading the packages.

```console
$ yacr upload --help
Upload Helm chart packages to GitHub Releases

Usage:
  yacr upload [flags]

Flags:
  -c, --commit string                  Target commit for release
  -b, --git-base-url string            GitHub Base URL (only needed for private GitHub) (default "https://api.github.com/")
  -r, --git-repo string                GitHub repository
  -u, --git-upload-url string          GitHub Upload URL (only needed for private GitHub) (default "https://uploads.github.com/")
  -h, --help                           help for upload
  -o, --owner string                   GitHub username or organization
  -p, --package-path string            Path to directory with chart packages (default ".yacr-release-packages")
      --release-name-template string   Go template for computing release names, using chart metadata (default "{{ .Name }}-{{ .Version }}")
  -t, --token string                   GitHub Auth Token

Global Flags:
      --config string   Config file (default is $HOME/.yacr.yaml)
```

### Create the Repository Index from GitHub Releases

Once uploaded you can create an `index.yaml` file that can be hosted on GitHub Pages (or elsewhere).

```console
$ yacr index --help
Update a Helm chart repository index.yaml file based on a the
given GitHub repository's releases.

Usage:
  yacr index [flags]

Flags:
  -c, --charts-repo string             The URL to the charts repository
  -b, --git-base-url string            GitHub Base URL (only needed for private GitHub) (default "https://api.github.com/")
  -r, --git-repo string                GitHub repository
  -u, --git-upload-url string          GitHub Upload URL (only needed for private GitHub) (default "https://uploads.github.com/")
  -h, --help                           help for index
  -i, --index-path string              Path to index file (default ".yacr-index/index.yaml")
  -o, --owner string                   GitHub username or organization
  -p, --package-path string            Path to directory with chart packages (default ".yacr-release-packages")
      --release-name-template string   Go template for computing release names, using chart metadata (default "{{ .Name }}-{{ .Version }}")
  -t, --token string                   GitHub Auth Token (only needed for private repos)
  --packages-with-index                Save a copy of the package files to the GitHub pages branch and reference them in the index

Global Flags:
      --config string   Config file (default is $HOME/.yacr.yaml)
```

## Configuration

`yacr` is a command-line application.
All command-line flags can also be set via environment variables or config file.
Environment variables must be prefixed with `YACR_`.
Underscores must be used instead of hyphens.

CLI flags, environment variables, and a config file can be mixed.
The following order of precedence applies:

1. CLI flags
1. Environment variables
1. Config file

### Examples

The following example show various ways of configuring the same thing:

#### CLI

    yacr upload --owner myaccount --git-repo helm-charts --package-path .deploy --token 123456789

#### Environment Variables

    export YACR_OWNER=myaccount
    export YACR_GIT_REPO=helm-charts
    export YACR_PACKAGE_PATH=.deploy
    export YACR_TOKEN="123456789"
    export YACR_GIT_BASE_URL="https://api.github.com/"
    export YACR_GIT_UPLOAD_URL="https://uploads.github.com/"

    yacr upload

#### Config File

`config.yaml`:

```yaml
owner: myaccount
git-repo: helm-charts
package-path: .deploy
token: 123456789
git-base-url: https://api.github.com/
git-upload-url: https://uploads.github.com/
```

#### Config Usage

    yacr upload --config config.yaml


`yacr` supports any format [Viper](https://github.com/spf13/viper) can read, i. e. JSON, TOML, YAML, HCL, and Java properties files.

Notice that if no config file is specified, `yacr.yaml` (or any of the supported formats) is loaded from the current directory, `$HOME/.yacr`, or `/etc/yacr`, in that order, if found.

#### Notes for Github Enterprise Users

For Github Enterprise, `yacr` users need to set `git-base-url` and `git-upload-url` correctly, but the correct values are not always obvious to endusers.

By default they are often along these lines:

```
https://ghe.example.com/api/v3/
https://ghe.example.com/api/uploads/
```

If you are trying to figure out what your `upload_url` is try to use a curl command like this:
`curl -u username:token https://example.com/api/v3/repos/org/repo/releases`
and then look for `upload_url`. You need the part of the URL that appears before `repos/` in the path.

## Goals

- [ ] [kubeval](https://github.com/instrumenta/kubeval) style schema validation
- [ ] [helm lint](https://helm.sh/docs/helm/helm_lint/)
- [ ] [Conftest](https://github.com/open-policy-agent/conftest/)
- [ ] Only package a chart if it has changed
- [ ] Add ability to update dependencies before packaging
- [ ] Add option to insert release notes similar to [buchanae/github-release-notes](https://github.com/buchanae/github-release-notes)

- Copyright © [stecky](https://github.com/stecky)