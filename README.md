# predown

> Preprocess Markdown files using Go templates and TOML configuration

## Usage

```
$ > predown template.md
$ > predown template.md --data data.toml
$ > predown template.md --data data.toml --wrap wrapper.frontmatter

$ > predown template.md > output.md

$ > predown template.md output.md
$ > predown template.md output.md --data data.toml
$ > predown template.md output.md --data data.toml --wrap wrapper.frontmatter
```

## Install

To download the latest binary for your system (`linux` or `darwin`), just use the `curl` request from below. Depending on your internet connection, GitHub might refuse the connection due to API rate limitations.

```bash
$ > curl -sSLfo predown \
    ` \
      curl -sSLf \
      https://api.github.com/repos/sbstjn/predown/releases/latest \
      | grep browser_download_url \
      | cut -d '"' -f 4 \
      | grep \`uname -s | tr A-Z a-z\` \
    `
$ > chmod +x predown
$ > ./predown --version
```