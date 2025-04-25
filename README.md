# 🔩 mend

Mend is a simple HTML template processor designed to, but not limited to be used to generate static websites.

The produced HTML is always consistently formatted and sorted.

<!-- vim-markdown-toc GFM -->

* [📥 Installation](#-installation)
* [⚙️ Usage](#-usage)

<!-- vim-markdown-toc -->

- [Wiki / Documentation](https://github.com/bbfh-dev/mend/wiki)

# 📥 Installation

Download the [latest release](https://github.com/bbfh-dev/mend/releases/latest) or install via the command line with:

```bash
go install github.com/bbfh-dev/mend@latest
```

# ⚙️ Usage

Check out the [Wiki](https://github.com/bbfh-dev/mend/wiki) for detailed documentation.

Run `mend --help` to display CLI usage information.

```bash
mend v1.0.0-alpha

HTML template processor designed to, but not limited to be used to generate static websites

Usage:
    mend [options] <html files...>

Options:
    --help
        # Print this help message
    --version
        # Print the program version
    --tabs, -t
        # Use tabs for indentation
    --indent <int> (default: 4)
        # The amount of spaces to be used for indentation (overwriten by --tabs)
    --strip-comments
        # Strips away HTML comments from the output
    --input <string>
        # Provide input to the provided files in the following format: 'attr1=value1,attr2.a.b.c=value2,...'
    --output <string>
        # (Optional) output path. Use '.' to substitute the same filename (e.g. './out/.' -> './out/input.html')
```
