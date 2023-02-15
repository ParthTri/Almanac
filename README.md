# Summary

Almanac is an open source plain text calendar format, designed to quickly capture, search and share calendar events.

Almanac is created with the goal of replacing your cloud based calendar apps, i.e. Google Calenadar, Apple Icalendar, etc, etc. 

This repository contains the Alamanc format but also comes with the companion cli tool with the same name. This README will be focused on both, however in the docs you will find greater details for both the format and the cli tool.

The CLI companion app aims to provide a better command line experience, to quickly add, remove, search, filter and view your calendar events. It also provides a way to export your `.almanac` files to other formats such as `.ical`.

**NOTE: The almanac companion app is still in development however the format can be used without an dependencies.** 


# Features

Almanac provides a framework for scheduling calendar events, that are:

- Human readable
- CLI tools accessible (i.e. grep)
- In a plain text format 
- Portable

# Instalation/Quickstart

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


# Usage 

## Foramt 
An overview of the format:
```
2023-02-13
	09:00-09:10 Meditate
	17:00-17:30 Accounting Meeting +work
		* Talk about balance sheet
```

The almanac format utilizes the [ISO 8061](https://en.wikipedia.org/wiki/ISO_8601) standard, although slightly modified. ISO 8061 date times generally look like `YYYY-MM-DDTHH:MM:SS`. Almanac uses the format but foregoes the `T` character to seperate the date and time. The date and time are seperated by a single space, as well as a the EVENT name.

The example above uses the multiline format, for when there are more than one events scheduled in a day.

- Every entry is tabbed in, followed by the time and the name of the event.
- Events can be given tags by prefixing them with the `+` or plus operator.
- Events can be given extra details using the `*` or astrix symbol.

Another way of writing it would be the following:
```
2023-02-14 14:00-16:00 Computer Science lecture +school
```

This is a more concise way of showing scheduled items. This format could be used to schedule multiple items of the same day using the same date time stamp. This is entirely up to you as the user.

And that is it. Almanac aims to be simple to write and simple to read, making very fast to interact with.

To learn more about the different features including in the markup look at the docs here. 

## Companion Tool

The Almanac companion tool can be used in conjuction with the Almanac format to supercharge your workflow.

By default Almanac will look for your calendar file path at environment variable `$ALMANAC`. Alternatively you can pass in the flag `-f` or `--file` followed by the path to the Almanac file.

### Add

New entries can be seemlessly added using the `-a` or `--add` flag. 

```shell
$ almanac -a 2023-02-15 10:00-12:00 Dentist appointment +health * Routine check
```

New entries must contain the date of the event. If the time is ommited the event is treated as an all-day event.

When adding tags and descriptions, almanac will autodetect them and add them accordingly.

### Edit

Entries can be cancelled using the `-e` or `--edit` flag, followed by the date timestamp of the event. Everything following the event's datetime will be used to overwrite the previous event details.

You can change the time:
```shell
$ almanac -e 2023-02-15 10:00-12:00 12:00-13:00
```

Add or remove tags:
```shell
$ almanac -e 2023-02-15 10:00-12:00 +routine -health
```

Change the description:
```shell
$ almanac -e 2023-02-15 10:00-12:00 * Ask dentist about cleaning
```

However sometimes we cannot remember the exact time of things scheduled, but do remeber the day. What you can do instead is enter the date of the event you want editted, and you will be presented with all the events on that day you could edit. 

Upon selecting one you will be able to edit it interactively.

### Cancel

Entries can be cancelled using the `-c` or `--cancel` flag, followed by the date timestamp of the event.

```shell
$ almanac -c 2023-02-15 10:00-12:00
```

Cancelled events will be prepened with double front slashes or `//` to *"comment"* them out. They will continue to exist in the file if the need arises to reschedule them, but will not show up in [Views](###Views) and [Exports](###Exports) unless explicitly stated so.

Much like editing events, you can supply the date and will be prompted with the options for that day.

### Searching
Events can be searched using the companion tool or other CLI tools like grep and fzf. This section will focus on inbuilt searching tool.

To search for events use the `-s` or `--search` flags. This are to be followed up by:
- Dates or Date Ranges
- Times (this could be either start or end times)
- Tags
- Event Names
- Event descriptions (fuzzy search)

This can all be combined to get more accurate results.
```shell
$ almanac -s 2023-02-01 - 2023-02-07
```
This will show all the events from the 1st of February 2023 to the 7th of February 2023.

For more detailed examples check out the docs.

### Views

The views command provides a ASCII drawing representation of your schedule. This can be filtered to show your day, week, or month.  This can be produced using the `-v` or `--view` command, followed by the `day`, `week`, or `month` keywords.

For example:
- `$ almanac -v day`
- `$ almanac -v week`
- `$ almanac -v month`

Further options include 
- Hiding days and Times that don't have events
- Showing block view (diplaying a day in proper format)
- Outline view (showing all the events and their time stamps)

To use them use the corrosponding flag:
- `-h` `--hide`, example `$ almanac -v -h day`
- `-b` `--block`, example `$ almanac -v -b day`
- `-o` `--outline`, example `$ almanac -v -o day`

### Exports

Due to Almanac lacking a proper application some users may find it harder to use with the lack of notifications. To support this Almanac provides ways to export your file into an iCalendar file or `.ics`.

This can be achieved using the `-E` or `--export` flags. An optional file name can be provided to change the name of the output file.

```shell
$ almanac -E personal.ics
```

Almanac also supports filtered exporting, to learn more about this check out the docs here.
