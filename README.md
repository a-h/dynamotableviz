# dynamotableviz - DynamoDB Table Visualizer

While working on a talk, I got frustrated with how long it was taking to create diagrams, and make screenshots.

The screenshots were relatively slow to load in presentations, sometimes fuzzy, and didn't work well on my blog either.

So, I thought I'd create a CLI tool.

## Workflow

Create a file containing your data in key/value format.

* Key names cannot contain equals signs.
* Quotes are optional, but required if you want to have newlines or commas in your data.

```
pk=users/1,sk=details,name=Albert Einstein,occupation=Scientist
pk=users/2,sk=details,name=Bill Evans,occupation=Musician
pk=users/2,sk=discography/1956/new_jazz_conceptions,title=New Jazz Conceptions,year=1956
```

Execute dynamotableviz, passing:

* The name of the `pk` (defaults to `pk`).
* Optional `sk` name if your table has one.
* Optional comma seprated list of named attributes (e.g. `gsi1`, `ttl`).
* The name of the input file (can also pipe input).

```
./dynamotableviz -pk=pk -sk=sk -attrs=occupation -file ./example.txt
```

It will write HTML to stdout, so depending on where you're using it, you might want to redirect it to a file.

```
./dynamotableviz -pk=pk -sk=sk -attrs=occupation -file ./example.txt > index.html
```

![Web browser rendering of output](screenshot.png)

## Usage

```
Usage of ./dynamotableviz:
  -attrs string
        Defines named attributes, which are then shown as a column (default "gsi1,gsi2,gsi3,ttl")
  -file string
        Load the data from the file instead of stdin
  -omit-css
        Set to true to disable the output of CSS
  -pk string
        Name of the partition key attribute (default "pk")
  -sk string
        Name of the sort key attribute
```

## Custom styling

By default CSS is output. Use the `omit-css` CLI flag to disable it, and write your own.

You can view the default CSS in the `table.templ` file.

## Getting images

Once you've generated your HTML file, you can use Firefox or Chrome in Headless mode to automate the creation of screenshots.

```bash
firefox --headless --screenshot file:///home/user/github.com/a-h/dynamotableviz/index.html
```

```bash
/opt/google/chrome/chrome --headless --window-size=1600,900 --screenshot=screenshot.png --screenshot file:///path/to/file/index.html
```

If you want to customise how it looks, you can use a site generator like Hugo, or automate the merging of the CLI output with custom CSS and HTML.

