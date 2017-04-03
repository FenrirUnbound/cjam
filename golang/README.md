## Usage

```bash
$ go build -o run
$ ./run -i input.txt -o output.txt
# Note:
#  -i: optional. Defaults to 'input.txt'
#  -o: optional. Defaults to 'output.txt'
```

### func Solve

```
func Solve(cases []string) []string
```

**cases** - Input from the problem set input. Typically has *n* number of cases as the 0th element, followed by *n* cases.

**output** - Each element in the array should map to the case it answers.
