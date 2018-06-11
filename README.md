# predown

> Preprocess Markdown files using Go templates and TOML configuration.

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

### Bash

To download the latest binary for your system (`linux` or `darwin`), just use the following `curl` request:

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

### Makefile

To add `predown` to your existing setup, you can update your `Makefile` to download the binary file.

```make
PREDOWN= ./predown

$(PREDOWN):
  @ curl -sSLfo $(PREDOWN) \
    ` \
      curl -sSLf https://api.github.com/repos/sbstjn/predown/releases/latest \
      | grep browser_download_url \
      | cut -d '"' -f 4 \
      | grep $(shell uname -s | tr A-Z a-z) \
    `
    @ chmod +x $(PREDOWN)
    @ $(PREDOWN) --version
```
