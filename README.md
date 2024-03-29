# AwkP(review)

A proof of concept for live previewing output of an awk command.

The application is only about 100 lines of code.

[![asciicast](https://asciinema.org/a/439432.svg)](https://asciinema.org/a/439432)

# Usage:

After building the binary with `go build`:

```
./awkp path_to_file
```

You should have `awk` in your path already for this to work correctly.

# Installation

```
go install github.com/ngynkvn/awkp@latest
```

# Motivation

There's a lot of friction for new users when learning command line tools like
awk, sed, and grep.

The motivation was to see how a quick feedback loop could potentially be useful for iteratively building out a command.

`awk` is one of those commands I quickly iterate on until I get the appropriate output.

TUI's are just quicker to pull up than an online playground, and awk was a
simple target for experimenting ideas with how people interact with computers.


If you're interested in some cool TUI libraries, I recommend checking some of
these:

- _Go_, the language and library this tool was written in,
  https://github.com/rivo/tview
  - This one has come to my attention recently: https://github.com/charmbracelet/bubbletea
- _Rust_, https://github.com/fdehau/tui-rs
- _Python_,
  https://github.com/willmcgugan/textual

# Nice to haves

Some extensions I think would be nice to have:

- Scrolling the output if it's longer than the text preview.
- Truncating the number of records returned in the preview (like
  `awk {print} | head -n 5`)
- Descriptions for the different flags that awk supports, and interative buttons
  / fields for setting them.

# Dependencies

- https://github.com/rivo/tview
- https://github.com/kballard/go-shellquote
- https://gopkg.in/alecthomas/kingpin.v2
