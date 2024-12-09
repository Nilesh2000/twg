## What is a test?

A test is a repeatable process that verifies whether or not something is working or not.
You can have either manual or automated tests.

We will focus on automated software testing, which means we will write code that runs our tests.

## Why do we even need tests?

1. Tests help find, fix, and prevent mistakes (bugs, side effects, edge cases, etc).
2. Tests document expected behaviour.
3. Test encourage us to write better code. (Side effect of writing tests)
4. Tests can speed up development. Really helps when you want to go back to better projects.

## How to write great tests?

1. Test With a Purpose.
2. Don't overdo it. When you learn tests first time, you want to test everything.
3. Testing is a skill, skills evolve.
4. Not too many, not too few, just enough. Avoid write tests that literally never fail.

All of the above are arbitrary points and may vary from team to team.
You get paid for code that works, not for tests. Test as little as possible to reach a given level of confidence.

Use code coverage to measure what percentage of your code is tested.
Use it cautiously since it might sometimes be deceptive.

Tests are just go code.

Test file names must be of the format *_test.go
This tells go to not use these files for builds that are not tests.

The test function name should be Test*(t *testing.T). * refers to the name of the function you want to test.
It should start with uppercase letter.

testing.T must be passed in every function that is a test.
This allows you tell Go when to pass a test, skip a test, or fail a test.

t.Errorf allows you to tell Go that you have failed a test case.

When you run `go test`, go is building a binary using all the _test.go files.
This binary is then executed. At the end of the tool, it cleans up the binary and related artifacts.

File Naming Conventions
File name must be *_test.go.
When you do `go build`, the go compiler does not include *_test.go files in the executable binary.
Typically, you have a source and a source_test.go file.
This is not mandatory, but its a convention.

Some Caveats:-
- `export_test.go` to access unexported variables in external tests.
- `xxx_internal_test.go` for internal tests.
- `example_xxx_test.go` for examples in an isolated files.

Function name should preferably be of the format func Test(t *testing.T) {}
You could also write separate test functions for different if-statements in our code.
We will learn table-driven tests in the future as another approach towards such situations.

Variable name should usually be `got` and `want` inside test functions.
Some people also use `arg` to specify the function arguments.

> %q formats a string with quotes. "hello" is printed as "hello" in printf statements.
