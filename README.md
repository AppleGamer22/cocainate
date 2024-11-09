# `cocainate`
[![Go Reference](https://pkg.go.dev/badge/github.com/AppleGamer22/cocainate.svg)](https://pkg.go.dev/github.com/AppleGamer22/cocainate) [![Test](https://github.com/AppleGamer22/cocainate/actions/workflows/test.yml/badge.svg)](https://github.com/AppleGamer22/cocainate/actions/workflows/test.yml) [![CodeQL](https://github.com/AppleGamer22/cocainate/actions/workflows/codeql.yml/badge.svg)](https://github.com/AppleGamer22/cocainate/actions/workflows/codeql.yml) [![Release](https://github.com/AppleGamer22/cocainate/actions/workflows/release.yml/badge.svg)](https://github.com/AppleGamer22/cocainate/actions/workflows/release.yml) [![Update Documentation](https://github.com/AppleGamer22/cocainate/actions/workflows/tag.yml/badge.svg)](https://github.com/AppleGamer22/cocainate/actions/workflows/tag.yml)

## Description
`cocainate` is a cross-platform CLI utility for keeping the screen awake until stopped, or for a specified duration.

## Why This Name?
The program's functionality and name are inspired by [macOS's `caffeinate`](https://github.com/apple-oss-distributions/PowerManagement/blob/main/caffeinate) utility that prevents the system from entering sleep mode.

This name is simply a stupid ~~pun~~, therefore **I do not condone and do not promote drug use**, for more information: [Wikipedia](https://en.wikipedia.org/wiki/Cocaine_(song)).

## Installation
### Nix Flakes
```nix
{
  inputs = {
    # or your preferred NixOS channel
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    applegamer22.url = "github:AppleGamer22/nur";
  };
  outputs = { nixpkgs }: {
    nixosConfigurations.nixos = nixpkgs.lib.nixosSystem {
      specialArgs = {
        pkgs = import nixpkgs {
          # ...
          overlays = [
            (final: prev: {
              # ...
              ag22 = applegamer22.packages."<your_system>";
            })
          ];
        };
      };
      modules = [
        # or in a separate Nix file
        ({ pkgs, ... }: {
          programs.nix-ld.enable = true;
          environment.systemPackages = with pkgs; [
            ag22.cocainate
          ];
        })
        # ...
      ];
    };
  };
}
```
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
* [Homebrew Tap](https://github.com/AppleGamer22/homebrew-tap):
```bash
brew install AppleGamer22/tap/cocainate
```

### Windows (working progress)
* [`winget`](https://github.com/microsoft/winget-cli):
```bash
winget install AppleGamer22.cocainate
```
### Other
* `go`:
	* Does not ship with:
		* a manual page
		* pre-built shell completion scripts
```
go install github.com/AppleGamer22/cocainate
```

## Functionality
### Root `-d`/`--duration` Flag
This is an optional flag that accepts a duration string (see [Go's `time.ParseDuration`](https://pkg.go.dev/time#ParseDuration) for more details). If this flag is not provided, the program will run until manually stopped.

#### Acceptable Time Units
* nanoseconds: *`ns`*
* microseconds: *`us`* or *`µs`*
* milliseconds: *`ms`*
* seconds: *`s`*
* minutes: *`m`*
* hours: *`h`*

#### Examples
* 10 hours: `-d 10h`
* 1 hour, 10 minutes and 10 seconds: `-d 1h10m10s`
* 1 microsecond: `-d 1us`

If the `-p` flag is provided, the `-d` flag's value is used as process polling interval.

### Root `-p`/`--pid` Flag
This is an optional flag that accepts a process ID (PID). If a valid PID is provided, the program will wait until that process is terminated. The delay between the termination of the provided process and the termination of screensaver inhibitation depends on the `-d` flag (which must be provided).

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
	* Any other desktop environment that implements [`org.freedesktop.ScreenSaver`](https://people.freedesktop.org/~hadess/idle-inhibition-spec/re01.html)
	<!-- * [MATE](https://mate-desktop.org) -->
<!-- ### macOS
* [D-Bus](https://www.freedesktop.org/wiki/Software/dbus/) (optional)
### Windows -->

## Common Contributor Routines
### Testing
Running the following command will run `go test` on the commands and session sub-modules:
```bash
make test
```
### Building From Source
#### Development
* Using the following `make` command will save a `cocainate` binary with the last version tag and the latest git commit hash:
```bash
make debug
```

#### Release
* Using the following [GoReleaser](https://github.com/goreleaser/goreleaser) command with a version `git` tag and a clean `git` state:
```bash
goreleaser build --clean
```
* All release artificats will stored in the `dist` child directory in the codebase's root directory:
	* compressed package archives with:
		* a `cocainate` binary
		* manual page
		* shell completion scripts
	* checksums
	* change log

## Copyright
`cocainate` is free software; you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation; either version 3, or (at your option) any later version.

`cocainate` is distributed in the hope that it will be useful, but **WITHOUT ANY WARRANTY**; without even the implied warranty of **MERCHANTABILITY** or **FITNESS FOR A PARTICULAR PURPOSE**.  See the GNU General Public License for more details.