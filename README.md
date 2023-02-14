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

