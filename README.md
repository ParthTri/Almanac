# Overview

## TODO Features

## Format
An overview of the format:
```
2023-02-13
	09:00 - 09:10 Meditate
	17:00 - 17:30 Accounting Meeting +work
		* Talk about balance sheet
```

The almanac format utilizes the [ISO 8061](https://en.wikipedia.org/wiki/ISO_8601) standard, although slightly modified. ISO 8061 date times generally look like `YYYY-MM-DDTHH:MM:SS`. Almanac uses the format but foregoes the `T` character to seperate the date and time. The date and time are seperated by a single space, as well as a the EVENT name.

The example above uses the multiline format, for when there are more than one events scheduled in a day.

- Every entry is tabbed in, followed by the time and the name of the event.
- Events can be given tags by prefixing them with the `+` or plus operator.
- Events can be given extra details using the `*` or astrix symbol.

Another way of writing it would be the following:
```
2023-02-14 14:00 - 16:00 Computer Science lecture +school
```

This is a more concise way of showing scheduled items. This format could be used to schedule multiple items of the same day using the same date time stamp. This is entirely up to you as the user.

And that is it. Almanac aims to be simple to write and simple to read, making very fast to interact with.

To learn more about the different features including in the markup look at the docs here. 

