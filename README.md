jsonfilter
==================

Pipe json into the app to get formatted stream out.

Example
```bash
go build main.go
echo "{\"foo\": {\"bar\": [\"baz\"]}}" | ./main
{
  "foo": {
    "bar": [
      "baz"
    ]
  }
}
```
