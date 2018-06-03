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
mashup -n=1 http://www.sacred-texts.com/tao/taote.htm http://vimdoc.sourceforge.net/htmldoc/usr_02.html
Processing http://www.sacred-texts.com/tao/taote.htm...
Processing http://vimdoc.sourceforge.net/htmldoc/usr_02.html...
 Earth takes its claws, nor robbers and yet we see its law of it, and far-reaching is so? By 
 this command: :help pattern<Ctrl-D> See |29.1| for any of Heaven, it is hidden, and so doing 
 so there is that of action, it could renounce our employment of arms enfold, From this model and weak
```
