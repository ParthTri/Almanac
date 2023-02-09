---
title: Almanac
description: A plain text calendar
---

# Overview

## TODO Features

## Format
An overview of the format:
```
2023-02-10 11:31 Meeting
```

The almanac format utilizes the [ISO 8061](https://en.wikipedia.org/wiki/ISO_8601) standard, although slightly modified. ISO 8061 date times generally look like `YYYY-MM-DDTHH:MM:SS`. Almanac uses the format but foregoes the `T` character to seperate the date and time. 

The date and time are seperated by a single space, as well as a the EVENT name.

The above example shows a single date entry for an event. For multiple items in a day the following is the preferred way of listing events.
```
2023-02-10
    11:31 Meeting
    13:40 Lunch with mum
```

Every subsequent entry is tabbed in, followed by the time and the name of the event.

And that is it. Almanac aims to be simple to write and simple to read, making very fast to interact with.

To learn more about the different features including in the markup look at the docs here. 

