# Releasing new versions of Git Town

This guide is for maintainers who make releases of Git Town.

### update release notes

- file `RELEASE_NOTES.md`
- commit to `main`

### bump the version

- search-and-replace the old version with the new version
- if bumping the major version, also update `github.com/git-town/git-town/v8/`
  everywhere in this repo

### create a GitHub release

On a Linux machine:

- install [hub](https://github.com/github/hub#installation)
- install [goreleaser](https://goreleaser.com/install)
- create and push a new Git Tag for the release: `git tag v8.0.0`
- `env GITHUB_TOKEN=<your Github token> VERSION=8.0.0 make release-linux`
  - or omit the Github token and enter your credentials when asked
- this opens a release in draft mode the browser
- delete the empty release that the script has created
- copy the release notes into the good release
- leave the release as a draft for now

On a Windows machine, in Git Bash:

- install [hub](https://github.com/github/hub#installation)
- install [go-msi](https://github.com/mh-cbon/go-msi#install)
- install [wix](https://wixtoolset.org/releases)
- optionally install
  [.NET 3.5](https://dotnet.microsoft.com/download/dotnet-framework)
- `env VERSION=8.0.0 make msi` to create the Windows installer
- test the created Windows installer in the `dist` directory
- `env GITHUB_TOKEN=<your Github token> VERSION=8.0.0 make release-win`
- this opens the release in the browser
- verify that it added the `.msi` file
- publish the release
- merge the `main` branch into the `public` branch

### create a Homebrew release

TODO: try the new `brew bump-formula-pr` command next time.

- fork [Homebrew](https://github.com/Homebrew/homebrew-core)
- update `Library/Formula/git-town.rb`
  - get the sha256 by downloading the release (`.tar.gz`) and using
    `shasum -a 256 /path/to/file`
  - ignore the `bottle` block, the homebrew maintainers update it
- create a pull request and get it merged

### Arch Linux

Flag the package out of date on the right hand side menu of
[Git Town's AUR page](https://aur.archlinux.org/packages/git-town/).
[allonsy](https://github.com/allonsy) will update the package.

### debugging

To test the goreleaser setup:

```
goreleaser --snapshot --skip-publish --rm-dist
```
