Almanac Usage Docs

# Table of Contents

- [Overview](#overview)
- [Adding Entries](#adding-entries)
- [Editing Events](#editing-events)
- [Cancelling Events](#cancelling-events)
- [Searching](#searching)
  - [Dates](#dates)
  - [Time](#times)
  - [Tags](#tags)
  - [Event Names](#event-name)
  - [Event Descriptions](#event-descriptions)
- [Views](#views)
- [Exports](#exports)
  - [Filtered Exports](#filtered-exports)

# Overview

The Almanac companion tool can be used in conjuction with the Almanac
format to supercharge your workflow.

By default Almanac will look for your calendar file path at environment
variable `$ALMANAC`{.verbatim}. Alternatively you can pass in the flag
`-f`{.verbatim} or `--file`{.verbatim} followed by the path to the
Almanac file.

# Adding Entries

New entries can be seemlessly added using the `-a`{.verbatim} or
`--add`{.verbatim} flag.

``` example
$ almanac -a 2023-02-15 10:00-12:00 Dentist appointment +health * Routine check
```

New entries must contain the date of the event. If the time is ommited
the event is treated as an all-day event.

When adding tags and descriptions, almanac will autodetect them and add
them accordingly.

# Editing Events

Entries can be cancelled using the `-e`{.verbatim} or
`--edit`{.verbatim} flag, followed by the date timestamp of the event.
Everything following the event\'s datetime will be used to overwrite the
previous event details.

You can change the time:

``` shell
$ almanac -e 2023-02-15 10:00-12:00 12:00-13:00
```

Add or remove tags:

``` shell
$ almanac -e 2023-02-15 10:00-12:00 +routine -health
```

Change the description:

``` shell
$ almanac -e 2023-02-15 10:00-12:00 * Ask dentist about cleaning
```

However sometimes we cannot remember the exact time of things scheduled,
but do remeber the day. What you can do instead is enter the date of the
event you want editted, and you will be presented with all the events on
that day you could edit.

Upon selecting one you will be able to edit it interactively.

# Cancelling Events

Entries can be cancelled using the `-c`{.verbatim} or
`--cancel`{.verbatim} flag, followed by the date timestamp of the event.

``` shell
$ almanac -c 2023-02-15 10:00-12:00
```

Cancelled events will be prepened with double front slashes or
`//`{.verbatim} to *\"comment\"* them out. They will continue to exist
in the file if the need arises to reschedule them, but will not show up
in [Views](###Views) and [Exports](###Exports) unless explicitly stated
so.

Much like editing events, you can supply the date and will be prompted
with the options for that day.

# Searching

Events can be searched using the companion tool or other CLI tools like
grep and fzf. This section will focus on inbuilt searching tool.

To search for events use the `-s`{.verbatim} or `--search`{.verbatim}
flags. This are to be followed up by: - Dates or Date Ranges - Times
(this could be either start or end times) - Tags - Event Names - Event
descriptions (fuzzy search)

This can all be combined to get more accurate results.

``` shell
$ almanac -s [2025-02-01 2025-02-07]
```

This will show all the events from the 1st of February 2025 to the 7th
of February 2025.

## Dates

Events can be searched for through the date syntax. You can pass in the
date to show all the events on a particular day.

``` shell
$ almanac -s 2025-02-01
```

This will reveal all events on the 1st of February 2025.

You can also use the date syntax as described [here](./Foramt.md#dates),
to filter out events in a specific date range.

``` shell
$ almanac -s [2025-02-01 2025-02-07]
```

## Times

You can also use times to search for events. To do so you use the search
flag and pass in the time it starts or finishes.

``` shell
$ almanac -s 09:00
```

This will find all the events that start or end at 09:00.

You can also combine this with ranges to find events that fall in a
particular range. To learn more about time ranges see
[here](./Foramt.md#time)

``` shell
$ almanac -s 09:00-10:00
```

This will find all events that start or finish between 09:00 to 10:00.

**Note that this doesn\'t include that the event elapses over this
time.**

``` shell
...
2025-02-16 12:00-14:00 Lunch Meeting
...
```

If we have the above event scheduled in our Almanac file and then run
the following command:

``` shell
$ almanac -s 13:00-13:00
```

The Lunch Meeting event won\'t be found as the event itself doesn\'t
start or end during that time.

## Tags

Tags can also be used to search for events. You can pass the tag as an
argumet to find all the matching results.

``` shell
$ almanac -s +work
```

This would search for everything that has a `+work`{.verbatim} tag.

Much like with everything else you can chain tags to get more precise
results. For example:

``` shell
$ almanac -s +work +volunteer
```

This will return all events that have both work AND volunter.

## Event Name

To search for event names you can simply pass the event name as an
argument after the search flag and it will exact match the event name.

``` shell
$ almanac -s Accounting Meeting
```

This will search for all your scheduled events that are called
Accounting Meeting.

## Event Descriptions

You can also fuzzy search through event descriptions to find
descriptions that contain text you pass through. To run this however you
need to first use the description identifier `*`{.verbatim} before
passing the arguments.

For example:

``` shell
$ almanac -s * meeting
```

This will search through all the descriptions that contain the word
meeting.

# Views

The views command provides a ASCII drawing representation of your
schedule. This can be filtered to show your day, week, or month. This
can be produced using the `-v`{.verbatim} or `--view`{.verbatim}
command, followed by the `day`{.verbatim}, `week`{.verbatim}, or
`month`{.verbatim} keywords.

For example:

- `$ almanac -v day`{.verbatim}
- `$ almanac -v week`{.verbatim}
- `$ almanac -v month`{.verbatim}

Further options include

- Hiding days and Times that don\'t have events
- Showing block view (diplaying a day in proper format)
- Outline view (showing all the events and their time stamps)

To use them use the corrosponding flag:

- `-h`{.verbatim} `--hide`{.verbatim}, exmaple
  `$ almanac -v -h day`{.verbatim}
- `-b`{.verbatim} `--block`{.verbatim}, example
  `$ almanac -v -b day`{.verbatim}
- `--o`{.verbatim} `--outline,`{.verbatim} example
  `$ almanac -v -o day`{.verbatim}

# Exports

Due to Almanac lacking a proper application some users may find it
harder to use with the lack of notifications. To support this Almanac
provides ways to export your file into an iCalendar file or
`.ics`{.verbatim}.

This can be achieved using the `-E`{.verbatim} or `--export`{.verbatim}
flags. An optional file name can be provided to change the name of the
output file.

``` shell
$ almanac -E personal.ics
```

## Filtered Exports

Filtered exports works by exporting only certain entries that match the
provided tag. See below for an example:

``` shell
$ almanac -E -f +work
```

This example would export only events with `+work`{.verbatim} tag. You
can also chain tags by appending them to the end of the line. For
example:

``` shell
$ almanac -E -f +work +volunteer 
```

This would export all your work and volunteer events into a
`.ics`{.verbatim} file.

This can also be used for filtering out a tag by using the
`-`{.verbatim} operator infront of the tag name.

``` shell
$ almanac -E -f -work
```

This will export all entries EXCEPT events that have the tag work.
