
#### Add New Language

1. runner implementation for the language, [sample](https://github.com/itsjwala/locus/blob/master/runner/languages/python/runner.go)

2. language Dockerfile, [sample](https://github.com/itsjwala/locus/blob/master/runner/languages/python/Dockerfile)

3. source the runner implementation in [languages](https://github.com/itsjwala/locus/blob/master/runner/languages/languages.go)


#### Nice to have

* support for user input(stdin)

* timeouts if infinite loop

*  more secure docker files (currently running everything as root 😛 )

    . . . .
