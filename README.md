# `cocainate`
## Why This Name?
The program's functionality and name are inspired by [macOS's `caffeinate`](https://github.com/apple-oss-distributions/PowerManagement/blob/PowerManagement-1132.141.1/caffeinate) utility that prevents the system from entering sleep mode.

This name is simply a stupid ~~pun~~, therefore **I do not condone and do not promote drug use**, for more information: [Wikipedia](https://en.wikipedia.org/wiki/Cocaine_(song)).

## Installation (working progress)
### Arch Linux Distributions
* [`yay`](https://github.com/Jguer/yay):
```bash
yay -S cocainate-bin
```
* [`paru`](https://github.com/morganamilo/paru):
```bash
paru -S cocainate-bin
```
### macOS
* [Homebrew Tap](https://github.com/AppleGamer22/homebrew-cocainate):
	1. Add tap
	```bash
	brew tap AppleGamer22/cocainate
	```
	2. Install:
	```bash
	brew install cocainate
	```
### Windows (working progress)
* [`winget`](https://github.com/microsoft/winget-cli):
```bash
winget install AppleGamer22.cocainate
```
### Other
* `go`:
```
go install github.com/AppleGamer22/cocainate
```

## Functionality
### Global `-d`/`--duration` Flag
This is an optional flag that accepts a duration string (see [Go's `time.ParseDuration`](https://pkg.go.dev/time#ParseDuration) for more details). If this flag is not provided, the program will run until manually stopped.

#### Acceptable Time Units
* nanoseconds: `ns`
* microseconds: `us` (or `µs`)
* milliseconds: `ms`
* seconds: `s`
* minutes: `m`
* hours: `h`

#### Examples
* 10 hours: `-d 10h`
* 1 hour, 10 minutes and 10 seconds: `-d 1h10m10s`
* 1 microsecond: `-d 1us`

<!-- ### Global `--pid` Flag
This is an optional flag that accepts a process ID (PID). If a valid PID is provided, the program will wait until that process is terminated. -->

### `version` Sub-command
#### `-v`/`--verbose` Flag
* If this flag is provided, the following details are printed to the screen:
	1. semantic version number
	2. commit hash
	3. Go compiler version
	4. processor architecture & operating system
* Otherwise, only the semantic version number is printed.

## Dependencies
### Linux
* [D-Bus](https://www.freedesktop.org/wiki/Software/dbus/)
* One of the following desktop environments:
	* [KDE](https://kde.org) 4 or later
	* [GNOME](https://gnome.org) 3.10 or later
	<!-- * [MATE](https://mate-desktop.org) -->
### macOS
* [D-Bus](https://www.freedesktop.org/wiki/Software/dbus/) (optional)
### Windows
