# obspruner

obspruner is a simple go utility designed to prune dangingling attachments from an Obsidian vault. When run it will look for any attachments that are not referenced in any of your Obsidian files and ask if you would like to delete them.

## Limitations

- MacOS only 
  - Presumably will compline on other Unix operating systems but this has not been tested
- Assumes all attachments are stored in a single folder in your vault
- Currently only detects embedded attachments with the `![[{attachment name}]]` syntax

## Usage

`obspruner <path to vault> <path to attachments directory>`

## Installation

### Option 1

If you have Go installed you can install with:

```bash
go install github.com/jmarkIT/obspruner@latest
```

### Option 2

Download the the executable from the Releases section and place it anywhere in your path.

## About

