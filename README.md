# predown

[![GitHub release](https://img.shields.io/github/release/sbstjn/predown.svg?maxAge=600)](https://github.com/sbstjn/predown/releases)
[![MIT License](https://img.shields.io/github/license/sbstjn/predown.svg?maxAge=3600)](https://github.com/sbstjn/predown/blob/master/LICENSE.md)
[![Go Report Card](https://goreportcard.com/badge/github.com/sbstjn/predown)](https://goreportcard.com/report/github.com/sbstjn/predown)
[![Build Status](https://img.shields.io/circleci/project/sbstjn/predown.svg?maxAge=600)](https://circleci.com/gh/sbstjn/appsync-router)

Preprocess Markdown files using Go templates and TOML configuration.

## Usage

```bash
# Print merged data to stdout

$ > predown template.md
$ > predown template.md --data data.toml
$ > predown template.md --data data.toml --wrap wrapper.frontmatter

# Pipe merged data to file

$ > predown template.md > output.md

# Create file and store data

$ > predown template.md output.md
$ > predown template.md output.md --data data.toml
$ > predown template.md output.md --data data.toml --wrap wrapper.frontmatter
```

### Example

Let's assume you have `example.md`, `data.toml`, and `wrapper.frontmatter` files like these:

##### example.md

```markdown
# Headline

Variable Example from `data.toml` is **{{ .Example }}**.
````

##### data.toml

```toml
Example = "Foo"
```

##### wrapper.frontmatter

```frontmatter
---
Example: {{ .Data.Example }}
---

{{ .Content }}
```

#### Command

Using `predown`, you can merge the data from `data.toml` with `example.md`, and wrap everything in `wrapper.frontmatter` afterwords.

```bash
$ > predown template.md output.md \
    --data data.toml \
    --wrap wrapper.frontmatter
```

This will create a file called `output.md` with the following content:

```markdown
---
Example: Foo
---

Headline
========

Variable Example from `data.toml` is **Foo**.
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

## License

Feel free to use the code, it's released using the [MIT license](LICENSE.md).

## Contribution

You are welcome to contribute to this project! 😘 

To make sure you have a pleasant experience, please read the [code of conduct](CODE_OF_CONDUCT.md). It outlines core values and beliefs and will make working together a happier experience.
