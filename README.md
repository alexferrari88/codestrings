# codestrings 🔤

[strings](<https://en.wikipedia.org/wiki/Strings_(Unix)>) but for your source code.

Extracts all the strings from your code.

## Use Cases 🧑‍💻

- Facilitate translation of hard-coded text strings
- Quick glance at what a code is "saying" by looking at the strings within it (?)
- NLP (?)

## Features ⚡

- Extract to json or csv
- Multiple files at once
- Single executable

_So far, this has been only tested with Javascript._

## Usage ⚙️

```bash
codestrings [parameters] file1,file2,file3,...
```

### Parameters

```bash
  --delimiters
        delimiters to use for string extraction (comma separated and escaped) (default "\",',`")
  --output
        output type: json, csv (default "csv")
```

## Examples 🧪

Let's say you have this javascript file called `example.js`.

```javascript
function greet(name) {
  console.log(`Hello ${name}`);
}

var fname = "John";
const lname = "Doe";
let age = 42;
let address = {
  street: "123 Main St",
  city: "New York",
  state: "NY",
};
```

You can extract all the strings from this file with:

```bash
codestrings example.js
```

The output will be:

```bash
example.js,Hello ${name},John,Doe,123 Main St,New York,NY
```

## Roadmap 🛣️

- Add more tests especially with more programming languages

- Improve algorithm performance

## License 📃

[MIT](https://choosealicense.com/licenses/mit/)
