# CAIN - Simple Media Management
### Simple and Easy Media Management with CAIN
![CAIN Preview GIF](https://github.com/atrox/cain/raw/master/preview.gif)
Cain is using __[FileBot][filebot]__ under the hood to organize your files and plays very well with your favorite __Torrent/Download__ Client.

# Features
- Fully automated media management
- Simple to setup and easy to use
- Plays well with your Downloader, Torrent Client or whatever else you like to use
- __[FileBot][filebot]__ under the hood. *No need to reinvent the wheel*

# Installation

## Requirements
- __[FileBot][filebot]__ installed and `filebot` executable in `$PATH`

## Install
__Cain__ is one single executable, that you put in your `$PATH`.
After that you can run `cain` by yourself or automate everything and let your Torrent Client handle the rest.

## Setup
Before you can use Cain you need to setup the configuration. Cain provides some sensitive defaults to get started with.
Run `cain setup` to interactively configure Cain.


# Configuration
- **RetrievePath**: Placeholder
- **AutoUpdate**: Enable/Disable automatic updates.

## Destinations

## Naming Schemes

## Notifiers
Cain can also automatically notify specific apps for changes.

- **Kodi**: `host[:port]` Tell the given Kodi/XBMC instance to rescan it's library
- **Plex**: `host[:token]` Tell the given Plex instance to rescan it's library. Plex Home instances require an [authentication token](https://support.plex.tv/hc/en-us/articles/204059436-Finding-your-account-token-X-Plex-Token).
- **Emby**: `host:apikey` Tell the given Emby instance to rescan it's library.
- **Pushover**: `userkey` Send update notifications to your devices via Pushover.
- **PushBullet**: `apikey` Send full reports to all your PushBullet devices
- **Gmail**: `username:password` Use the following gmail account to send and receive full reports. You must use an [App Password](https://support.google.com/accounts/answer/185833?hl=en) for security reasons.
- **Mail**: `host:port:from[:username:password]` Send email via custom mail server

# Command line reference

[filebot]: http://www.filebot.net/