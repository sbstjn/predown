# predown

> Preprocess Markdown files using Go templates and TOML configuration. See [privacy](https://github.com/sbstjn/privacy) as example …

## Installation

### Bash

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
$ > ./predown version
```

### Makefile

To add `predown` to your existing setup, you can update your `Makefile` (hopefully you have one!) to download the binary file.

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
    @ $(PREDOWN) version
```
