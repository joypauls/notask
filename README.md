<!-- ![Build Status](https://img.shields.io/github/workflow/status/joypauls/notask/Build) -->
![Latest Release](https://img.shields.io/github/v/release/joypauls/notask?include_prereleases)
[![Go Report](https://goreportcard.com/badge/github.com/joypauls/notask)](https://goreportcard.com/badge/github.com/joypauls/notask)
<!-- ![Code Coverage](https://storage.googleapis.com/notask-build/code-coverage.svg) -->
![Go Version](https://img.shields.io/github/go-mod/go-version/joypauls/notask)
[![go.dev Reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/joypauls/notask)

> :warning: **WIP**: This is just a prototype, use at own risk!

# Notask: an oversimplified Notion client

Stay focused with **notask** - a distraction-free interface to manage tasks at the command line using a Notion database (such as a [Board](https://www.notion.so/help/boards)) as a backend through [Notion's API](https://developers.notion.com/docs/getting-started). 

Features:
- View a database
- Add a new page

If you're interested in hacking on this project, make sure to check [this](#developer-stuff) out.

## Current Limitations
- Doesn't play nice with page/card titles that contain emoji in the title
  - However, using the assigned page icon is *not* an issue
- Time only in UTC

## Future Additions
- Built-in timer

## Developer Stuff

### Compatibility/Environment

- Notion
  - Desktop app version **2.1.15** to verify actions in dev
  - Web app in Google Chrome to verify actions in dev
- Hardware
  - Dev machine is an M1 MacBook Air running MacOS **Ventura 13.3.1**

`cp .notask.yaml-template .notask.yaml`

### Release a New Version

1. Create new tag with git like so: `git tag -a v<NUMBER> -m "<MESSAGE>"`
2. Push the new tag `git push --tags`
