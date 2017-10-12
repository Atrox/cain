# CAIN - Simple Media Management
[<img src="https://files.atrox.me/selif/scroll.png" align="right" width="100">][repo]

Cain *automatically* organises your movies and TV-shows according to your configuration.
Your favourite torrent/download client should call Cain after the download has finished and Cain will do the rest.

> Cain is still WIP and may very well contain bugs. Please feel free to report those [here][issues]

## Features
- Fully automated media management
- Simple to setup and easy to use
- Plays well with your [Downloader][download-client], [Torrent][torrent-client] Client or whatever else you like to use
- [FileBot][filebot] under the hood - *No need to reinvent the wheel*

## Installation

### Requirements
[FileBot][filebot] needs to be installed and the `filebot` executable has to be globally available.

### Install
[Latest Release][latest-release] is available for all major operating systems as a prebuilt binary.

1. Download and extract the [archive][latest-release]
2. Put the binary (`cain`) somewhere save and add the location to your `PATH`
3. Run `cain setup`

### Setup
Before you can use Cain you need to configure some basic things first. Cain provides some sensitive defaults to get started with.

Run `cain setup` to interactively configure Cain. After you finished the setup process you are pretty much ready to go.
Now you can configure your [Downloader][download-client] and/or [Torrent][torrent-client] Client to automatically run Cain after finishing downloading.

## Configuration
> Location: `~/.config/cain/config.yaml`

- **defaultRetrievePath**: Where to get the unsorted media from
    - If not specified, parameter `--path` is required.
- **autoUpdate**: Enable/Disable automatic updates.
- **language**: [2-letter language code][language-iso-codes] (default: en)
- **nonStrictMatching**: Enable to non strictly match for movies/tv-shows. Be aware that this could result in wrong matches.
- **cleanupAfterwards**: Cain will automatically clean the remaining unused/unneeded files after moving the matched files.
- **hideBanner**: If `true` Cain will no longer show the ascii banner before every command

### Naming Schemes
Please see the [FileBot Documentation][filebot-naming] for more informations on this subject.

#### Default Naming Schemes
- **movie**: `{n} ({y})/{n}`
- **series/anime**: `{n}/Season {s.pad(2)}/{n} - {s00e00} - {t}`
- **music**: `{n}/{album}{pi.pad(2)}{artist} - {t}`

### Notifiers
Cain can also automatically notify specific apps for changes.

- **kodi**: `host[:port]` Tell the given Kodi/XBMC instance to rescan it's library
- **plex**: `host[:token]` Tell the given Plex instance to rescan it's library. Plex Home instances require an [authentication token][plex-token].
- **emby**: `host:apikey` Tell the given Emby instance to rescan it's library.
- **pushover**: `userkey` Send update notifications to your devices via Pushover.
- **pushBullet**: `apikey` Send full reports to all your PushBullet devices
- **gmail**: `username:password` Use the following gmail account to send and receive full reports. You must use an [App Password][gmail-app-password] for security reasons.
- **mail**: `host:port:from[:username:password]` Send email via custom mail server

## Command line reference
- `cain setup`: Configure Cain interactively.
- `cain run`: Run Cain *(this should idealy be run by your Download/Torrent Client)*
    - `--path`: Where Cain should look for media to sort. This Parameter is required if `DefaultRetrievePath` is not set.
    - `--path-env`: Get path from specified environment variable.
    - `--non-strict`: Enable to non strictly match for movies/tv-shows. Be aware that this could result in wrong matches.

## Contributing

Everyone is encouraged to help improve this project. Here are a few ways you can help:

- [Report bugs][issues]
- Fix bugs and [submit pull requests][pulls]
- Write, clarify, or fix documentation
- Suggest or add new features

[filebot]: http://www.filebot.net/
[filebot-naming]: http://www.filebot.net/naming.html

[download-client]: https://github.com/Atrox/cain/wiki/Download-Clients
[torrent-client]: https://github.com/Atrox/cain/wiki/Torrent-Clients

[repo]: https://github.com/atrox/cain
[issues]: https://github.com/atrox/cain/issues
[pulls]: https://github.com/atrox/cain/pulls
[latest-release]: https://github.com/atrox/cain/releases/latest

[language-iso-codes]: http://en.wikipedia.org/wiki/List_of_ISO_639-1_codes
[plex-token]: https://support.plex.tv/hc/en-us/articles/204059436-Finding-your-account-token-X-Plex-Token
[gmail-app-password]: https://support.google.com/accounts/answer/185833?hl=en
