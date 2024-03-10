The Year Progress Indicator is a Go application that visually represents how much of the current year has elapsed. 
It calculates the percentage of the year that has passed based on today's date and displays this information through 
a progress bar composed of filled and unfilled segments. This utility is designed for those who would like to visualise 
the passage of time throughout the year.

# Features

- **Automatic Date Calculation**: Dynamically computes the current date to determine the year's elapsed portion.
- **Leap Year Handling**: Accurately accounts for leap years in day counts.
- **Visual Progress Display**: Utilises `█` for the elapsed part of the year and `▁` for what remains, complemented by a percentage display.

# Use Cases

- **Terminal Welcome Message**: Add the `yp` command to your shell's profile file (e.g., `.bashrc` or `.zshrc`) to see the year's progress every time you open a new terminal window.
- **GitHub README**: Include the progress bar in your project's README.md to show the year's progress in a creative way.
- **Daily Logs/Journals**: Include the year's progress in digital or printed daily logs to visualize how much of the year has passed.
- **Project Milestones**: Use the progress indicator to track the passage of time relative to project deadlines or milestones.
- **Personal Websites**: Incorporate the tool's output into a personal or professional website to share the year's progression with visitors.

# Installation

For convenience, pre-built binaries are available for various platforms. Download the appropriate binary for your system from the releases page.

## Homebrew

To install the Year Progress Indicator using Homebrew on macOS or Linux, you can follow these steps:

```shell
brew tap martinsirbe/clinkclank
brew install martinsirbe/clinkclank/yp
```

This will add the custom tap and install the `yp` CLI, making it readily accessible from any terminal.

## Build from Source
Make sure you have Go installed. You can then install the Year Progress Indicator globally via the following command:

```shell
go install github.com/martinsirbe/go-year-progress/cmd/yp@v0.1.0
```

This command compiles and installs the binary to your Go bin directory, making it accessible from any terminal provided the directory is in your system's PATH.

# Usage

Run the tool from any terminal with `yp`.

# Configuration

You can configure the application using CLI options or by setting environment variables. This flexibility allows you to choose the most convenient method for your workflow.

## CLI Options

You can specify the following options directly through the command line:

* `-e`, `--empty-char` - Character for empty sections of the progress bar (default "▁")
* `-f`, `--filled-char` - Character for filled sections of the progress bar (default "█")
* `-l`, `--location` - Timezone location for accurate time display (default "UTC")
* `-t`, `--total-blocks` - Total number of blocks in the progress bar (default 50)

## Environment Variables

As an alternative to CLI options, you can set the following environment variables. These are particularly useful for persistent configurations or for contexts where command-line arguments may not be ideal:

* `YP_LOCATION` - Overrides the default timezone location.
* `YP_TOTAL_BLOCKS` - Sets the total number of blocks in the progress bar.
* `YP_FILLED_CHAR` - Specifies the character for filled sections of the progress bar.
* `YP_EMPTY_CHAR` - Specifies the character for empty sections of the progress bar.

## Example Outputs

Without any configuration, the output will look like this:
```
███████▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁ 15.18%
```

### Using CLI Options

To customise the progress bar with CLI options, you might use the following command:
```shell
yp --location "America/New_York" --total-blocks 30 --filled-char "x" --empty-char "-"
```

### Using Environment Variables
Alternatively, you can achieve the same customisation by setting environment variables:
```shell
YP_LOCATION="America/New_York" YP_TOTAL_BLOCKS=30 YP_FILLED_CHAR="x" YP_EMPTY_CHAR="-" yp
```

Both of the above configurations adjust the progress bar to reflect a timezone of "America/New_York", a total of 30 blocks, with "x" representing filled sections and "-" for empty ones. The resulting output, assuming a similar progress of 15.18% UTC, would be:
```shell
xxxx-------------------------- 15.13%
```

# Contributing

Feel free to contribute to the project by submitting pull requests or creating issues for bugs and feature requests.

# License

This project is licensed under the MIT License. See [LICENSE.md](LICENSE.md).
