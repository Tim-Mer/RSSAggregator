RSSAggregator
A simple and efficient RSS feed aggregator built with Go.

[![Go Reference](https://pkg.go.dev/badge/github.com/Tim-Mer/RSSAggregator.svg)](https://pkg.go.dev/github.com/Tim-Mer/RSSAggregator)

RSSAggregator is a Go program designed to fetch, parse, and aggregate content from multiple RSS feeds. It's ideal for developers and users who want to centralize their news consumption from various sources into a single, manageable stream. The project focuses on simplicity, performance, and ease of integration.

*   Fetch content from multiple RSS/Atom feeds.
*   Parse feed entries and extract relevant information.
*   Aggregate entries into a unified list.
*   Configurable feed sources.

## Installation

```bash
go install github.com/Tim-Mer/RSSAggregator@latest
```

## Usage

To start the aggregator and begin fetching feeds, run:

```bash
RSSAggregator
```

## Configuration

Configuration can be managed through a `config.yaml` file in the root directory of the project.

Example `config.yaml`:

```yaml
feeds:
  - url: "http://example.com/rss.xml"
    name: "Example Feed 1"
  - url: "https://anothersite.org/atom.xml"
    name: "Another Feed"
```

## Contributing

Contributions are welcome! Please refer to the [CONTRIBUTING.md](CONTRIBUTING.md) file for details on how to contribute.

License
MIT