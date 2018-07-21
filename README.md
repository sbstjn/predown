# predown

[![Current Release](https://badgen.now.sh/github/release/sbstjn/predown)](https://github.com/sbstjn/predown/releases)
[![MIT License](https://badgen.now.sh/badge/License/MIT/blue)](https://github.com/sbstjn/predown/blob/master/LICENSE.md)
[![CircleCI Build Status](https://badgen.now.sh/circleci/github/sbstjn/predown)](https://circleci.com/gh/sbstjn/predown)

Preprocess Markdown files using [Go templates](https://golang.org/pkg/text/template/) and [TOML](https://github.com/toml-lang/toml) configuration.

## Usage

```bash
# Show generated Markdown content in stdout

$ > predown template.md
$ > predown template.md --data data.toml
$ > predown template.md --data data.toml --wrap wrapper.frontmatter

# Pipe generated Markdown to file

$ > predown template.md > output.md

# Save generated Markdown in file

$ > predown template.md output.md
$ > predown template.md output.md --data data.toml
$ > predown template.md output.md --data data.toml --wrap wrapper.frontmatter

# Transform generated Markdown to HTML

$ > predown template.md output.html --data data.toml
```

### Example

Let's assume you have `example.md`, `data.toml`, and `wrapper.frontmatter` files like these:

##### example.md

```markdown
# Headline

Variable Example from `data.toml` is **{{ .Example }}**.

## Table

| Name | Age |
| ---- | --- |
{{- range .People }}
| {{ .Name }} | {{ .Age }} |
{{- end }}

Nice, isn't it? ;)
````

##### data.toml

```toml
Example = "Foo"

[[People]]

Name = "Frank"
Age = 21

[[People]]

Name = "Claude"
Age = 45

[[People]]

Name = "Natascha"
Age = 37
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

Table
-----

| Name     | Age |
|----------|-----|
| Frank    | 21  |
| Claude   | 45  |
| Natascha | 37  |

Nice, isn't it? ;)
```

As you can see, the `Markdown` file will be tidied up. This can be very handy if you work with tables in your Markdown files. After running `predown`, they nicely align …

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
