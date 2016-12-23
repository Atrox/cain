# CAIN - Simple Media Management
[<img src="https://files.atrox.me/selif/scroll.png" align="right" width="100">](https://github.com/Atrox/cain)

Cain is using [FileBot][filebot] under the hood to organize your files and plays very well with your favorite **Torrent/Download** Client.

## Features
- Fully automated media management
- Simple to setup and easy to use
- Plays well with your Downloader, Torrent Client or whatever else you like to use
- [FileBot][filebot] under the hood - *No need to reinvent the wheel*

## Installation

### Requirements
[FileBot][filebot] needs to be installed and the `filebot` executable has to be globally available.

### Install
Installation packages are available for all major operating systems:
- [Windows][dl-win]
- [MacOS][dl-mac]
    - Brew: `brew install eqnxio/atrox/cain`
- [Linux][dl-linux]

[View all prebuilt binaries][dl-page]

### Setup
Before you can use Cain you need to configure some basic things first. Cain provides some sensitive defaults to get started with.

Run `cain setup` to interactively configure Cain. After you finished the setup process you are pretty much ready to go.

How to configure your Torrent/Download Client to automatically run Cain: Wiki Link here

## Configuration
- **DefaultRetrievePath**: Where to get the unsorted media from
    - If not specified, parameter `--path` is required.
- **AutoUpdate**: Enable/Disable automatic updates.

### Naming Schemes
Please see the [FileBot Documentation][filebot-naming] for more informations on this subject.

#### Default Naming Schemes
- **Movie**: `{n} ({y})/{n}`
- **Series/Anime**: `{n}/Season {s.pad(2)}/{n} - {s00e00} - {t}`

### Notifiers
Cain can also automatically notify specific apps for changes.

- **Kodi**: `host[:port]` Tell the given Kodi/XBMC instance to rescan it's library
- **Plex**: `host[:token]` Tell the given Plex instance to rescan it's library. Plex Home instances require an [authentication token](https://support.plex.tv/hc/en-us/articles/204059436-Finding-your-account-token-X-Plex-Token).
- **Emby**: `host:apikey` Tell the given Emby instance to rescan it's library.
- **Pushover**: `userkey` Send update notifications to your devices via Pushover.
- **PushBullet**: `apikey` Send full reports to all your PushBullet devices
- **Gmail**: `username:password` Use the following gmail account to send and receive full reports. You must use an [App Password](https://support.google.com/accounts/answer/185833?hl=en) for security reasons.
- **Mail**: `host:port:from[:username:password]` Send email via custom mail server

## Command line reference
- `cain setup`: Configure Cain interactively.
- `cain run`: Run Cain *(this should idealy be run by your Download/Torrent Client)*
    - `--path`: Where Cain should look for media to sort. This Parameter is required if `DefaultRetrievePath` is not set.

## Contributing

Everyone is encouraged to help improve this project. Here are a few ways you can help:

- [Report bugs](https://github.com/atrox/cain/issues)
- Fix bugs and [submit pull requests](https://github.com/atrox/cain/pulls)
- Write, clarify, or fix documentation
- Suggest or add new features

[dl-win]: https://bin.equinox.io/c/Dyvh1T2kPn/cain-stable-windows-amd64.msi
[dl-mac]: https://bin.equinox.io/c/Dyvh1T2kPn/cain-stable-darwin-amd64.pkg
[dl-linux]: https://dl.equinox.io/atrox/cain/stable
[dl-page]: https://dl.equinox.io/atrox/cain/stable

[filebot]: http://www.filebot.net/
[filebot-naming]: http://www.filebot.net/naming.html
