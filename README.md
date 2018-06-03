# Mashup

Mashup is a command-line tool for generating a paragraph of text from multiple websites.

## Installation

With the Go toolchain:

`go get github.com/sgoedecke/mashup`

Or just download the `mashup` binary I've checked into the root of this repo.

## Usage

`mashup -n=[num-sentences] http://www.example.site/one http://www.site.example/two`

For example, to generate the Tao te Vim:
```
$ mashup -n=1 http://www.sacred-texts.com/tao/taote.htm http://vimdoc.sourceforge.net/htmldoc/usr_02.html
Processing http://www.sacred-texts.com/tao/taote.htm...
Processing http://vimdoc.sourceforge.net/htmldoc/usr_02.html...
   For those other words, if it in it that account to be liberal; shrinking from his mouth closed, and so 
   rich a vessel; and is to avoid it) will be seen throughout the world are running Microsoft Windows, 
   the world dashes against it; he will be kept full of men most valuable thing has no license; he has not 
   from the land (or place) under Microsoft Windows, the cursor is.
```
