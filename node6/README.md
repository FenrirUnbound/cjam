## Usage

```bash
$ node main.js -i input.txt -o output.txt
# Note:
#  -i: optional. Defaults to 'input.txt'
#  -o: optional. Defaults to 'output.txt'
```

### Solve module.exports function

```
function (problemCases, <callback>) => [answers]
```

#### Parameters

**problemCases** - Array of strings. Input from the problem set input. Typically has *n* number of cases as the 0th element, followed by *n* cases.

**[callback]** - Optional. Takes on the signature `(err, data)`. `err` represents any errors encountered; `data` is an array of strings that map to the problem case it answers

#### Expected Output

**output** - Either an array of strings, or a Promise that resolves to an array of strings. Each element in the array should map to the case it answers.
