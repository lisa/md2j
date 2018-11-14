# md2j

Very simplistic program to convert a specific style of Markdown link to a format Jira can consume.

# Usage

Fetch the source, and `go build`. Then `./md2j -inputFile /path/to/markdownsource.md`. The Jira-friendly version will be rendered on STDOUT.

# Link Styles

The only style of link style supported is in this sample Markdown document:

    # Link Example
    Visit the [home page of this project][1]

    [1]: https://github.com/lisa/md2j

Other styles, such as `visit the [home page][https://github.com/lisa/md2j]` are not supported at this time.
