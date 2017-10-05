# counter
String counter in go.

## Usage
```go
urls := []string{
  "google.com",
  "google.com",
  "yahoo.com",
  "hotmail.com",
}
c := &counter.String{}
for _, url := range urls {
  c.Add(url)
}
sorted := c.SortedValues()
max := sorted[0].String

c.Print(os.Stdout)
```