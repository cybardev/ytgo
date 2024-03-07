# yt.go

## A Go program to play media from YouTube without needing API keys

**PS**: this is a `Python -> Go` translation of [cybardev/ytpy][ytpy] which itself was a `Shell -> Python` translation of [pystardust/ytfzf][ytfzf] _(before I rewrote it from scratch)_

### Table of Contents

Click to navigate.

-   [Dependencies](#dependencies)
-   [Installation](#installation)
    -   [Manual](#manual)
    -   [With Go](#with-go)
-   [Usage](#usage)
    -   [Examples](#examples)
-   [Credits](#credits)

### Dependencies

-   [mpv][mpv]
-   [yt-dlp][ytdl]
-   [ffmpeg][ffmpeg]

### Installation

#### Manual

-   Download the file from the Releases page: [ytgo-{os}-{arch}][release]

    -   **PS**: Make sure to choose the right binary for your OS and architecture

-   Place it on your `$PATH` and make it executable.

#### With Go

> **PS**: the pkg.go.dev registry may have an outdated version. If you encounter bugs or feature disparity, please replace `main` after `@` with the first part of the commit hash, e.g. `github.com/cybardev/ytgo/cmd/ytgo@e819a79`, **OR** try the [Manual installation](#manual) method for the latest updates.

-   Run the following command:

    ```sh
    go install github.com/cybardev/ytgo/cmd/ytgo@main
    ```

-   Ensure `$GOPATH/bin` is added to `$PATH`. An easy way is to add this line to `~/.profile`:

    ```sh
    export PATH="$(go env GOPATH)/bin:$PATH"
    ```

### Usage

**PS**: For simplicity, I will refer to the binary as `ytgo`.

```sh
Usage of ytgo:
  -d	Display URL only
  -i	Interactive selection
  -m	Play music only
  -n int
    	Play nth media (default 1)
  -u	Play from URL
  -v	Display version
```

**HINT**: [Here][mpv_hotkeys]'s a list of mpv keyboard shortcuts for your convenience.

#### Examples

-   Play a video:

    `ytgo rickroll`

-   Play an audio:

    `ytgo -m gurenge band cover`

-   Play the third search result:

    `ytgo -n 3 racing into the night`

-   Play an audio from URL:

    `ytgo -u -m "https://www.youtube.com/watch?v=y6120QOlsfU"`

    -   **PS**: The URL must be quoted to avoid parsing by the shell

-   Find the URL of a video:

    `ytgo -d hotaru maiko fujita`

-   Interactive selection mode:

    `ytgo -i marmot scream meme`

### Credits

-   [pystardust][pystardust]'s [ytfzf][ytfzf]
-   [This article][article] I found during my quest to implement a simplified version of ytfzf in Python3
-   [StackOverflow answer][regex] used for the regex `var ytInitialData = ({.*?});`

<!-- Links -->

[ytpy]: https://github.com/cybardev/ytpy
[release]: https://github.com/cybardev/ytgo/releases/tag/latest
[mpv]: https://github.com/mpv-player/mpv
[ytdl]: https://github.com/yt-dlp/yt-dlp
[ffmpeg]: https://github.com/FFmpeg/FFmpeg
[mpv_hotkeys]: https://defkey.com/mpv-media-player-shortcuts
[pystardust]: https://github.com/pystardust
[ytfzf]: https://github.com/pystardust/ytfzf
[article]: https://www.codeproject.com/articles/873060/python-search-youtube-for-video
[regex]: https://stackoverflow.com/a/68262735
