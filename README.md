# yt.go

## A Go program to play media from YouTube without needing API keys

**PS**: this is a `Python -> Go` translation of [cybardev/ytpy][ytpy] which itself was a `Shell -> Python` translation of [pystardust/ytfzf][ytfzf] _(before I rewrote it from scratch)_

> **WARNING**: Windows is likely to not run. I'm assuming because of how `mpv` is executed by the program. Feel free to fork and make a PR if you would like to fix this. Meanwhile, WSL can be used to run the Linux binaries.

### Table of Contents

Click to navigate.

-   [Dependencies](#Dependencies)
-   [Installation](#Installation)
-   [Usage](#Usage)
    -   [Examples](#Examples)
-   [Credits](#Credits)
-   [Extras](#Extras)

### Dependencies

-   [mpv][mpv]
-   [yt-dlp][ytdl]
-   [ffmpeg][ffmpeg]

### Installation

-   Download the file from the Releases page: [ytgo-{os}-{arch}][release]
    -   **PS**: Make sure to choose the right binary for your OS and architecture
-   Place it on your `$PATH` and make it executable.

### Usage

**PS**: For simplicity, I will refer to the binary as `ytgo`.

```sh
Usage of ytgo:
  -f	Play from URL
  -l	Select from list
  -m	Play music only
  -n int
    	Play nth media (default 1)
  -u	Display URL only
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

    `ytgo -f -m "https://www.youtube.com/watch?v=y6120QOlsfU"`

    -   **PS**: The URL must be quoted to avoid parsing by the shell

-   Find the URL of a video:

    `ytgo -u hotaru maiko fujita`

-   Interactive selection mode:

    `ytgo -l marmot scream meme`

### Credits

-   [pystardust][pystardust]'s [ytfzf][ytfzf]
-   [This article][article] I found during my quest to implement a simplified version of ytfzf in Python3

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
