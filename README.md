# tqp

tqp (tiny query parser) is a bare bones Go package for parsing (potentially) structured search query strings. It offers a basic ability to extract key/value attributes from the input string, and returns unparsed "noise" back to the caller so the application can further handle the input as it desires.

## Schema

A `key` is any A-Z value (case-insensitive) string that is followed by a `:` character. It is followed by a non-zero length string, known as a `value`, which may be quoted or unquoted. Quoting the value impacts the characters it may contain. _Unquoted values_ are limited to `[a-z0-9-_]+` (ASCII letters, numbers, hyphens and underscores). _Quoted values_ are limited to `[\w()\-._ ]+` (ASCII letters, numbers, hyphens, underscores, spaces, periods, and parentheses). Any strings with disallowed characters will cause the attribute to be considered _noise_ (aka an unparsed user string). Any strings that are not key/value attributes will automatically be appended to the _noise_ output as well.

## Technical Specifics

- Keys are converted to lowercase when stored in the attributes map
- A key can be used multiple times in a single query as each key is associated with a `[]string` of values
- An empty value (e.g. "") will be discarded and not inserted in the attributes map
- Insertion order to value and noise slices follows the left-to-right processing of the string

## Examples

```
Input: 'hello world name:nick'
Output:
    attributes: map[string][]string{"name": []string{"nick"}}
    noise: []string{"hello", "world"}
```

```
Input: 'today is randomattr:"long value example" monday`
Output:
    attributes: map[string][]string{"randomattr": []string{"long value example"}}
    noise: []string{"today", "is", "monday"}
```

```
Input: 'blue pants with:pockets price:"less than $20"`
Output:
    attributes: map[string][]string{"with": []string{"pockets"}, "price": []string{"less than $20"}}
    noise: []string{"blue", "pants"}
```

```
Input: 'year:2023 year:2024'
Output:
    attributes: map[string][]string{'year': []string{"2023", "2024"}}
    noise: []string{}
```

## Usage

Install using `go get github.com/Cryptkeeper/tqp`

```go
package main

import "github.com/Cryptkeeper/tqp"

func main() {
    myQuery := "hello world name:nick"
    attrs, noise := tqp.Find(myQuery)
    // attrs is map[string][]string{"name": []string{"nick"}}
    // noise is []string{"hello", "world"}
}
```

