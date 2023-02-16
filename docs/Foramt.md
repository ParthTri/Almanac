<h1 style="text-align:center">Almanac Format Docs</h1>
----------

This file will contain all the nitty gritty details about the Almanac foramt. 

# Table of Contents

- [Entries](#entries) 
    - [Multi-Line](#multi-line-entries)
- [Tags](#tags)
- [Description](#description)
- [Ranges and Repetition](#ranges-and-repetition)
    - [Dates](#dates)
    - [Time](#time)

# Entries

Almanac events are created using the following format:
```almanac
YYYY-MM-DD HH:MM EVENT_NAME 
```

An example of this would be: 
```almanac
2025-02-13 09:00-09:10 Meditate
```

This event would be scheduled for 13th of Feb 2025, from 9:00 AM to 9:10AM (see [Time Ranges](#time) for more info), with the name of Meditate.

By default entries used the [ISO 8061](https://en.wikipedia.org/wiki/ISO_8601) standard, although slightly modified. ISO 8061 date times generally look like `YYYY-MM-DDTHH:MM:SS`. Almanac uses the format but foregoes the `T` character to seperate the date and time.

The time is also represented in 24-hour time.

## Multi-Line Entries

For days that you have more than one event scheduled you can list out the date entry, and tab in each event on a new line, beginning with the TIME or TIME_RANGE followed by the EVENT_NAME. See example below.

```almanac
2025-02-13
	09:00-09:10 Meditate
	17:00-17:30 Accounting Meeting
```

Here we have multiple events scheduled.

1) The date selected is 13 of February 2025
2) From 09:00 to 09:10, Meditate as been scheduled
3) From 17:00-17:30, an Accounting Meeting has been scheduled.

# Tags

Tags provide an easier way to group similar events together. This could be achieved by keeping seperate Almanac files for different purposes, for example, `personal` and `work`. However we also provide the distinction of adding tags to your events.

See below for an example:

```almanac
2025-02-13 09:00-09:10 Meditate +health
```

Tags are denoted using the `+` (plus operator) at the start of the keyword.

You can also add multiple tags to an event to further categorize them. For example: 

```almanac
2025-02-13 09:00-09:10 Meditate +health +daily +habit
```

Notice how for each extra tag, another plus operator must be used. This is to provide much better clarity when reading.

**Note** 

Tags with spaces in them are not recommended if you wish to use the componian CLI app. Instead it is recommended to use `_` (underscores) to use instead of a space.

For example:
```almanac
2025-02-13 09:00-09:10 Meditate +health +daily +habit +somthing_else
```

# Description

Sometimes when creating events we like to associate them with extra details. Almanac also supports this through the use of the `*` (asterisk or star) operator. When this identifier is used any text following it will be considered apart of the description.

For example:
```almanac
2025-02-13 09:00-09:10 Meditate * Go to the beach to meditate
```

This would be valid entry. We have the event of meditate with the description of *Go to the beach to meditate*. This is fine if you are adding a event without any tags however things fall apart once we add them.

The following description will kept being read up until the new line, meaning if using the companion tool the tags will be considered part of the event description. Also this kinda hard to read.
```almanac
2025-02-13 09:00-09:10 Meditate * Go to the beach to meditate +daily +health +habit
```

Instead the description should be on new line and indented in, like so (following is recommended).
```almanac
2025-02-13 09:00-09:10 Meditate +daily +health +habit
    * Go to the beach to meditate 
```

The overall syntax remains the same, but becomes more readable. This also remains true for multiple events in a day.

```almanac
2025-02-13
	09:00-09:10 Meditate +daily +health +habit
        * Go to the beach to meditate
	17:00-17:30 Accounting Meeting
```

You indent the description one more level in and write the description like normal.

# Ranges and Repetition

## Dates

Sometimes you have events that last over a couple of days. To signify this you can use the following method.

The overall syntax for events remains the same, however the date portion is slighlt modified. To include the date range you must surrond the dates in a `[ ]` (square brackets) and seperate them with a space.

For example:

```almanac
[2025-02-13 2025-02-28] 09:00-09:10 Meditate +daily +health +habit
    * Go to the beach to meditate 
```

This means that the *Mediate* event is repeated from the 13 of Feb to the 28 of Feb, at the same time everyday.

You could specify this more by including what day or days it should occur on in the square brackets. For example:

```almanac
[2025-02-13 2025-02-28 monday] 09:00-09:10 Meditate +daily +health +habit
    * Go to the beach to meditate 
```

The above would show that *Meditate* should show from 13 of Feb to 28 of Feb, only on Mondays.

This can be taken a step further by:
```almanac
[2025-02-13 2025-02-28 monday-friday] 09:00-09:10 Meditate +daily +health +habit
    * Go to the beach to meditate 
```

The above would show that *Meditate* should show from 13 of Feb to 28 of Feb, only from Mondays to Fridays (Monday, Tuesday, Wednesday, Thursday, and Friday).

## Time

Much like date ranges an event's time can be set to occur as a range. Going back to an earlier example, this would mean the event *Meditate* would occur between 09:00 to 09:10.
```almanac
2025-02-13 09:00-09:10 Meditate +daily +health +habit
    * Go to the beach to meditate 
```

In this case the timestamps are seperated by a dash `-` to signify a range.

However the syntax changes when you want to involve reptition. Using the same square bracket system from earlier, we surrond our time range in it preserving the dash, but follow it up with repeat interval.
```almanac
2025-02-13 [09:00-09:10 1h] Rest +daily +health +habit
    * Go to the beach to meditate 
```

This example shows the *Rest* event, from 09:00 to 09:10 and repeat it every hour from the start point. I.e. from 10-10:10, 11-11:10, 12-12:10, so on and so forth.

**Note this would only be for that day, and to repeat it over different days see [Dates](#dates).**

The repetition syntax follows this format `[interval time][unit]`. The units used are
- `h` for hour
- `m` for minute
- `s` for second (if for some reason you want it)

Another example would be:
```almanac
2025-02-13 [09:00-09:10 90m] Eat +health
```

In this example the event *Eat* is scheduled from 09:00-09:10 and repeated every *90 minutes* from 09:00. Meaning the next event is scheduled at 10:30. 

It is important to note that this interval could also be written as `1.5h`, as Almanac supports floating point intervals or intervals that exceed the typical conversion point, i.e. 90 minutes instead of 1.5 hours. These two are interchangable. The same goes for 0.2h being the same as 12 minutes or 720s (for whatever reason should you choose this).

