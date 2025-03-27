# 🔩 mend.html

Mend is a simple HTML template processor designed to, but not limited to be used to generate static websites.

Refer to the [Wiki](https://github.com/bbfh-dev/mend.html/wiki) for in-depth tutorial.

<!-- vim-markdown-toc GFM -->

* [📥 Installation](#-installation)
* [⚙️ Usage](#-usage)
* [📌 Developer notes](#-developer-notes)

<!-- vim-markdown-toc -->

- [Wiki / Documentation](https://github.com/bbfh-dev/mend.html/wiki)

# 📥 Installation

Download the [latest release](https://github.com/bbfh-dev/mend.html/releases/latest) or install via the command line with:

```bash
go install github.com/bbfh-dev/mend.html
# (optional) change binary name to "mend" instead of "mend.html"
mv ~/go/bin/mend.html ~/go/bin/mend
```

# ⚙️ Usage

Run `mend --help` to display usage information.

```bash
Usage:
        mend <options> [html files...]

Commands:

Options:
        --help, -h             Print help and exit
        --version, -V          Print version and exit
        --input, -i <value>    Set global input parameters
        --indent <value>       Set amount of spaces to indent with. Gets ignored if --tabs is used
        --tabs, -t             Use tabs instead of spaces
        --decomment            Strips away any comments
```

# 📌 Developer notes

These are some important development notes, informing about parts of the project that need to be polished out.

> **Expressions are very clunky.**
>
> 1. In code they require every node to implement its own processing while referencing a global function that handles them. Basically, it's just a big bowl of spaghetti. There's gotta be a better way of doing them.
> 1. The way expressions are parsed is very primitive, it could cause unexpected errors/behavior when using bad syntax.
>
> — [@bbfh-dev](https://github.com/bbfh-dev/)
