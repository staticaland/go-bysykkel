- [About](#orgcfdef8a)
- [Usage](#org6dd7e97)
- [Installation](#org60d42ac)
  - [Go](#orge86087f)
  - [Homebrew](#orgdbcd921)
- [Examples](#orgd3b8213)
  - [Route that does not exist](#org8c8f157)
- [Logging](#orga3ca0b5)
- [Build](#orgeb18cbe)


<a id="orgcfdef8a"></a>

# About

<div class="html" id="org0cf8607">
<p>
&lt;p&gt;
    &lt;a href=&ldquo;<a href="https://github.com/staticaland/go-bysykkel/actions/workflows/build.yml">https://github.com/staticaland/go-bysykkel/actions/workflows/build.yml</a>&rdquo;&gt;
        &lt;img alt=&ldquo;Build&rdquo; src=&ldquo;<img src="https://github.com/staticaland/go-bysykkel/actions/workflows/build.yml/badge.svg" alt="badge.svg" class="org-svg" />&rdquo; /&gt;
    &lt;/a&gt;
    &lt;a href=&ldquo;<a href="https://github.com/staticaland/go-bysykkel/actions/workflows/release.yml">https://github.com/staticaland/go-bysykkel/actions/workflows/release.yml</a>&rdquo;&gt;
        &lt;img alt=&ldquo;Release&rdquo; src=&ldquo;<img src="https://github.com/staticaland/go-bysykkel/actions/workflows/release.yml/badge.svg" alt="badge.svg" class="org-svg" />&rdquo; /&gt;
    &lt;/a&gt;
    &lt;a href=&ldquo;<a href="https://github.com/staticaland/go-bysykkel/actions/workflows/run.yml">https://github.com/staticaland/go-bysykkel/actions/workflows/run.yml</a>&rdquo;&gt;
        &lt;img alt=&ldquo;Run&rdquo; src=&ldquo;<img src="https://github.com/staticaland/go-bysykkel/actions/workflows/run.yml/badge.svg" alt="badge.svg" class="org-svg" />&rdquo; /&gt;
    &lt;/a&gt;
&lt;/p&gt;
</p>

</div>

Simple CLI for Oslo City Bike stations.


<a id="org6dd7e97"></a>

# Usage

Follow the [Unix philosophy](https://en.wikipedia.org/wiki/Unix_philosophy) and use `fzf` to filter the output:

```sh
bysykkel | fzf --reverse --header-lines=3 --bind 'ctrl-r:reload(bysykkel)'
```

This also lets you reload the data with Ctrl-R.

You may also consider combining `grep` with `watch` to always know the status of your nearest bike station.

```sh
watch --interval 5 --no-title 'bysykkel | grep Birkelunden'
```

You could also add it to your status bar (like `i3status` or `polybar`).


<a id="org60d42ac"></a>

# Installation

[Check out releases for pre-built binaries](https://github.com/staticaland/go-bysykkel/releases).


<a id="orge86087f"></a>

## Go

```sh
go install bysykkel.go
```


<a id="orgdbcd921"></a>

## Homebrew

A Homebrew formula is included at [Formula/bysykkel.rb](./Formula/bysykkel.rb).

```sh
brew tap staticaland/go-bysykkel https://github.com/staticaland/go-bysykkel
brew install bysykkel
```

If you watch the project (Watch → Custom → Releases) you can easily upgrade to the latest version when notified:

```sh
brew upgrade bysykkel
```

To uninstall:

```sh
brew uninstall bysykkel
brew untap staticaland/go-bysykkel
```


<a id="orgd3b8213"></a>

# Examples


<a id="org8c8f157"></a>

## Route that does not exist

```sh
curl localhost:4000/foo | prettier --parser=json
```

```json
{
  "error": "The requested resource could not be found"
}
```


<a id="orga3ca0b5"></a>

# Logging

Decided on [Zerolog](https://github.com/rs/zerolog) after evaluating a few options<sup><a id="fnr.1" class="footref" href="#fn.1" role="doc-backlink">1</a></sup>.


<a id="orgeb18cbe"></a>

# Build

```sh
go build bysykkel.go
```

## Footnotes

<sup><a id="fn.1" class="footnum" href="#fnr.1">1</a></sup> [Structured logging - notes](https://notes.garden/Cards/%F0%9F%8C%B2+Notes/Structured+logging)
