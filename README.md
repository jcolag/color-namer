# color-namer

_color-namer_ attempts to find the nearest named color to an input.

## Concept

In a recent conversation, there was the possibility that someone was going to need arbitrary color inputs, but present the used colors in a human-readable (i.e., non-RGB) way.  That led to a side discussion around color spaces and distances.

That seemed worth writing up, and it seemed worth giving [Go](https://golang.org/) a shot.  As a disclaimer, what I've seen of Go hasn't impressed me, so this is as much to see if I can buy into the popular language or if it doesn't feel worth continuing on.

## Usage

Run:

    colorname RRGGBB

...where `RRGGBB` is a hex color code, such as `00FF00` for Electric green.

I would _like_ for input to be able to be formatted the way people would expect (`#00FF00`, in the example), but it turns out that the hash mark is _always_ used by the Linux command line as an end-of-line comment marker.  Oops.

## Notes

The list of colors is generated from [Wikipedia's compact list of colors](https://en.wikipedia.org/wiki/List_of_colors_%28compact%29) with the included `wikicolor.sh` script, which has a high enough signal-to-noise ratio that the HTML parts can be stripped out without much trouble.

