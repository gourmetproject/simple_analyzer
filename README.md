# Simple Analyzer
Simple Analyzer is an example analyzer for the Gourmet project. It is meant to be analagous to a
"Hello, World!" example program often found in software projects.

## Usage
To use Simple Analyzer, there are two ways:

1. Add the Github URL to your Gourmet `config.yml` file. Gourmet's [example config file]()
   demonstrates how to do this.
2. Pass the URL as command-line option like this:
   `gourmet -a github.com/gourmetproject/simple_analyzer`

## How It Works

Gourmet Analyzers are Go Plugins that implement the Gourmet `Analyzer` interface. In order to 
create a Gourmet Analyzer, your repository must:

1. Contain a `main.go` file at the root of the repository

2. The `main.go` file must be part in the `main` package. Analyzers are not standard Go libraries, 
   but are instead Go plugins.
   [Read the Go plugin docs for more information](https://golang.org/pkg/plugin/).

3. The code in `main.go` must implement the `gourmet.Analyzer` interface. Simple Analyzer does this
   by creating a new public type called SimpleAnalyzer, which is just an empty struct. It then
   declares two methods, Filter and Analyze, with SimpleAnalyzer as the pointer receiver. These two
   methods' signatures implement the `gourmet.Analyzer` interface.

    > NOTE: If you are not familiar with methods and pointer receivers,
    > [check this out](https://tour.golang.org/methods/4).

    > NOTE: Go enforces public/private visibility through the name of the identifier itself. If the
    > identifier starts in a capital letter, it is exported (public). If it starts with a lowercase
    > letter, it is private. You can
    > [read more here](https://golang.org/doc/effective_go.html#names).

4. The code in `main.go` must also implement the `gourmet.Result` interface. Simple Analyzer does
   this by creating a new public type called SimpleResult, which is just an integer. It then
   declares the Key method with SimpleResult as the value receiver (not pointer receiver, since we
   aren't modifying the underlying SimpleResult object in the Key method).

And that's it! We now have a simple Gourmet Analyzer.

# Contact Us

<a
href="https://join.slack.com/t/gourmetproject/shared_invite/enQtNzczMjQ4MzgzMTg5LTRjOTllNjc2MzNhMDQyNDdiMWQwZjQ5OTEwZDEyYjhiNWEwZjI3M2Y2MzExMGQ1ZjNkZjlkMjlkYTc3ZDZmN2Y">
	<img
		src="https://cdn.appstorm.net/web.appstorm.net/web/files/2013/10/slack_icon.png"
		alt="Slack icon"
		width="40"
	>
</a>

# Support Us

[![Patreon][patreon-badge]][patreon-link]

[patreon-badge]: https://img.shields.io/endpoint.svg?url=https%3A%2F%2Fshieldsio-patreon.herokuapp.com%2Fkvasirlabs&style=flat-round
[patreon-link]: https://patreon.com/kvasirlabs