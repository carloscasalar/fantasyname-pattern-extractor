# Fantasy Name Pattern Extractor

This is a simple tool to extract patterns from a given name using the rules of [fantasyname](https://github.com/skeeto/fantasyname).
[Here](https://github.com/s0rg/fantasyname) you can find an explanation of the patterns (and a golang implementation).

## Build

You need [golang](https://golang.org/dl/) installed in your machine. 

To build the extractor binary you can run:

```bash
make build
```

This will create a binary in `out/extract-pattern`.

## Usage

To run the extractor you can build it with the previous command and execute, or you can run it using the following command:

```bash 
make run -- -s "Tanis" 
```

## Tests

You can run the tests with the following command:

```bash
make test
```