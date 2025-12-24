# `txtable`

> Make your CSVs look fancy.

`txtable` transforms your CSV files into clean, readable, and styled tables right in your terminal. Perfect for CLI output and copying tables as Unicode (or ASCII) into documents or reports.

## Features

- Read CSV from file or stdin
- Choose table themes: `classic`, `compact`, `heavy`, `light`, `rounded`
- Align table cells and header: `left`, `center`, `right`
- Optional wrapping of cell content with configurable width
- Output ready for the terminal and Unicode/ASCII friendly for copy/paste

## Usage

```text
txtable [flags] [file]
```

### Examples

**Pipe CSV from stdin**

```bash
cat data.csv | txtable -t classic -w 20
```

**Read CSV from file**

```bash
txtable -t rounded -c center -x center data.csv
```

**Custom header alignment and wrapping**

```bash
txtable -x left -w 30 heroes.csv
```

**Copying ASCII-ready table for a document**

```bash
txtable -t heavy data.csv > table.txt
# Now you can paste table.txt into Markdown, LaTeX, or emails
```

## Themes

`txtable` supports multiple table styles:

* `classic` – minimal
* `compact` – borderless
* `heavy` – bold Unicode borders
* `light` – light Unicode borders
* `rounded` – rounded Unicode borders (default)


## Support

All tools are completely free to use, with every feature fully unlocked and accessible.

If you find one or more of these tool helpful, please consider supporting its development with a donation.

Your contribution, no matter the amount, helps cover the time and effort dedicated to creating and maintaining these tools, ensuring they remain free and receive continuous improvements.

Every bit of support makes a meaningful difference and allows me to focus on building more tools that solve real-world challenges.

Thank you for your generosity and for being part of this journey!

[![Donate with PayPal](https://img.shields.io/badge/💸-Tip%20me%20on%20PayPal-0070ba?style=for-the-badge&logo=paypal&logoColor=white)](https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=FV575PVWGXZBY&source=url)

## How To Install

### Using the _install.sh_ script (macOS & linux)

Simply run the following command in your terminal:

```sh
curl -sL https://raw.githubusercontent.com/lucasepe/txtable/main/install.sh | bash
```

This script will:

- Detect your operating system and architecture
- Download the latest release binary
- Install it into _/usr/local/bin_ (requires sudo)
  - otherwise fallback to _$HOME/.local/bin_ 
- Make sure the install directory is in your _PATH_ environment variable


### Manually download the latest binaries from the [releases page](https://github.com/lucasepe/txtable/releases/latest):

- [macOS](https://github.com/lucasepe/txtable/releases/latest)
- [Windows](https://github.com/lucasepe/txtable/releases/latest)
- [Linux (arm64)](https://github.com/lucasepe/txtable/releases/latest)
- [Linux (amd64)](https://github.com/lucasepe/txtable/releases/latest)

Unpack the binary into any directory that is part of your _PATH_.

## If you have [Go](https://go.dev/dl/) installed

You can also install it using:

```bash
go install github.com/lucasepe/txtable@latest
```

Make sure your `$GOPATH/bin` is in your PATH to run `txtable` from anywhere.

---

### Credits

Table rendering inspired by [`simpletable`](https://github.com/alexeyco/simpletable). Big thanks for making tables less boring!  
