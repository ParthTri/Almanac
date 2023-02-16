<h1 style="text-align:center">Almanac Documentation</h1>

You will find all the details about the Almanac format and command line tool right here.

# Table of Contents
- [Installation](#Installation)
    - [Binary Release](#binary-release)
    - [Build from source](#build-from-source)
    - [Package Managers](#in-progress-package-managers)
- [Format](./Foramt.md)
- [Usage](./Usage.md)

# Installation
## Binary Release

For Windows, Mac OS or Linux, you can download a binary release here.

## Build from source

To build Almanac from source, make sure you have `go` installed with version >= 1.18.1.

First clone this repo.
```shell
git clone https://github.com/ParthTri/Almanac.git
```

Next `cd` in to the directory.
```shell
cd Almanac
```

Finally run the build script:
```shell
./build.sh
```

This will create a binary in the `bin` directory, from there you should move it to somewhere on your `$PATH`.

## IN PROGRESS Package Managers

This is a work in progress feature, where you should be able to add it with your choice of package manager.

Intended Support:

- AUR
- Homebrew 
- Scoop (Windows)
- Debian
- Void Linux
- Nix 

