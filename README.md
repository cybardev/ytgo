# yt.go

## A Go program to play media from YouTube without needing API keys

<img height="128px" width="128px" src="./ytgo.png" alt="ytgo logo"><img height="128px" src="./ytgo-banner.png" alt="ytgo banner">

> [!NOTE]
> This is a `Python -> Go` translation of [cybardev/ytpy][ytpy] which itself was a `Shell -> Python` translation of [pystardust/ytfzf][ytfzf] _(before I rewrote it from scratch)_

[![Go Reference][reference_badge]][reference_link]
[![Go Report Card][go_report_badge]][go_report_link]
[![Test Coverage][coveralls_badge]][coveralls_link]
[![Workflow Status][workflows_badge]][workflows_link]

### Table of Contents

Click to navigate.

- [Dependencies](#dependencies)
- [Installation](#installation)
  - [Manual](#manual)
  - [With Go](#with-go)
- [Usage](#usage)
  - [Examples](#examples)
- [Credits](#credits)

### Dependencies

- [mpv][mpv]
- [yt-dlp][ytdl]
- [ffmpeg][ffmpeg]

### Installation

#### Manual

- Download the file from the Releases page: [ytgo-{os}-{arch}][release]

  - **PS**: Make sure to choose the right binary for your OS and architecture

- Place it on your `$PATH` and make it executable.

#### With Go

> **Link to package**: [pkg.go.dev/github.com/cybardev/ytgo/v3][gopkg]

- Run the following command:

    ```sh
    go install github.com/cybardev/ytgo/v3/cmd/ytgo@latest
    ```

- Ensure `$GOPATH/bin` is added to `$PATH`. An easy way is to add this line to `~/.profile`:

    ```sh
    export PATH="$(go env GOPATH)/bin:$PATH"
    ```

> [!IMPORTANT]
> The [pkg.go.dev][gopkg] registry may have an outdated version. If you encounter bugs or feature disparity, please replace `latest` after `@` with `main`, e.g. `github.com/cybardev/ytgo/v3/cmd/ytgo@main`, **OR** try the [Manual installation](#manual) method for the latest updates.

### Usage

Output of `ytgo -h`:

```sh
Usage of ytgo:
  -d	Display URL only
  -f string
    	Play from playlist file
  -i	Interactive selection
  -m	Play music only
  -n int
    	Play nth media (default 1)
  -p	Prompt mode
  -u	Play from URL
  -v	Display version
```

**HINT**: [Here][mpv_hotkeys]'s a list of mpv keyboard shortcuts for your convenience.

#### Examples

- Play a video:

    `ytgo rickroll`

- Play an audio:

    `ytgo -m gurenge band cover`

- Play the third search result:

    `ytgo -n 3 racing into the night`

- Play an audio from URL:

    `ytgo -u -m "https://www.youtube.com/watch?v=y6120QOlsfU"`

  - **PS**: The URL must be quoted to avoid parsing by the shell

- Find the URL of a video:

    `ytgo -d hotaru maiko fujita`

- Interactive selection mode:

    `ytgo -i marmot scream meme`

- Playlist mode:

    `ytgo -f playlist.txt`

  - If file exists, items will be played; else file is created and user is prompted to add items.
  - File is a plaintext list of valid Video IDs, one per line. Generate one and view it for an example.

### Credits

- [pystardust][pystardust]'s [ytfzf][ytfzf]
- [This article][article] I found during my quest to implement a simplified version of ytfzf in Python3
- [StackOverflow answer][regex] used for the regex `var ytInitialData = ({.*?});`

<!-- Links -->

[ytpy]: https://github.com/cybardev/ytpy
[gopkg]: https://pkg.go.dev/github.com/cybardev/ytgo/v3
[release]: https://github.com/cybardev/ytgo/releases/tag/latest
[mpv]: https://github.com/mpv-player/mpv
[ytdl]: https://github.com/yt-dlp/yt-dlp
[ffmpeg]: https://github.com/FFmpeg/FFmpeg
[mpv_hotkeys]: https://defkey.com/mpv-media-player-shortcuts
[pystardust]: https://github.com/pystardust
[ytfzf]: https://github.com/pystardust/ytfzf
[article]: https://www.codeproject.com/articles/873060/python-search-youtube-for-video
[regex]: https://stackoverflow.com/a/68262735
[reference_link]: https://pkg.go.dev/github.com/cybardev/ytgo/v3/cmd/ytgo
[go_report_link]: https://goreportcard.com/report/github.com/cybardev/ytgo/v3
[coveralls_link]: https://coveralls.io/github/cybardev/ytgo
[workflows_link]: https://github.com/cybardev/ytgo/actions/workflows/latest.yml
[reference_badge]: https://pkg.go.dev/badge/github.com/cybardev/ytgo/v3/cmd/ytgo.svg
[go_report_badge]: https://goreportcard.com/badge/github.com/cybardev/ytgo/v3?style=flat-square
[coveralls_badge]: https://img.shields.io/coveralls/github/cybardev/ytgo?style=flat-square
[workflows_badge]: https://img.shields.io/github/actions/workflow/status/cybardev/ytgo/latest.yml?style=flat-square
