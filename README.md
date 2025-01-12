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

When you write tests, you need a way to signal test failure.

The `testing.T` type has a `Log` and `Logf` package. These work similar to `Print` and `Printf` in the `fmt` package.
`Logf` allows you to pass formatted messages.
However, `Log` and `Logf` and only printed when a test fails or when you run tests via `go test -v`.
Not recommented to use `fmt.Println` becuase it prints at the top and hence does not help in debugging.

These are two basic ways to signal that a test has failed:-
- `t.Fail()` = fail, but keep running
- `t.FailNow()` = fail now and stop test

You will rarely see the above being called:-
- `Fatal()` - `Log()` + `FailNow()`
- `Fatalf()` - `Logf()` + `FailNow()`
- `Error()` - `Log()` + `Fail()`
- `Errorf()` - `LogF()` + `Fail()`

Which you use?
- If you can keep test running, use `Error()`/`Errorf()`
- If a test is completely over and running won't help at all, use `Fatal()`/`Fatalf()`

If not using subtests, Fatal will other tests in the function from running.

httptest.NewRecorder() is a fake response writer that we can pass into our method to see if the handler is working correctly or not.

The purpose of a test is to debug an issue quickly.
You want to see what went wrong and then how to fix it.

Standard to check variables during a test
if got != want
if actual != expected
Also, while logging errors, print what you got first and then what you expected.

Make it simple to recreate your failures to help you fix code that causes your tests to fail.

We can use Fatal more frequently when using subtests becuase ending a subtest immeditely doesn't actually stop the rest of the subtests from running.
This gives you freedom to exit the test whenever you want while simultaneously allowing us more control on what tests continue to run.

You can write example as test cases and can them view them on the documentation.
This is extremely powerful becuase if your examples fail, you will see them on your terminal as a test failure.
They are implemented exactly the same way as tests, but just have a Example prefix instead. Also, you have to add an // Output command

Lets say you have a struct Demo. You write a receiver method func (d Demo) Hello() on it.
Then, the test will be defined as func TestDemo_Hello()

Similar for example to show up on documentation, use
func Example()
func ExampleTest()
func ExampleDemo_Text()
func Example_text() - Provide specific names to the examples

You can add a comment at the top of a file (before package) to see it under overview on the Go documentation

When you want to show an example in an unordered fashion.
Use // Unordered Output instead of // Output in the Example method.

If your example has multiple imports and stuff that is too large for a test files, move it to a separate Go file altogether.
The file name should be example_<func>_test.go. You should probably use example_test as the package name to express isolation.
To ensure it shows up on go doc along with its imports, your example file should have some sort of type or variable defined.
Examples help make things very clear in Golang.

Look at the standard library on how to write examples.

Table Driven Tests are a design pattern that allow you to have multiple tests for the same method without repeating a lot of code.
Use tt / tc for iterating over table-driven tests.
You can also setup Table Driven Tests for a single test. This makes it extensible and makes it simple to add more tests in the near future.

t.Run() is how you initiate a subtest.
[gotests](https://github.com/cweill/gotests) is the package used by code editor to generate table driven tests in go.


You can also do nested subtests.
Each subtest can run a smaller subtests while nesting.

You don't have to use just anonymous structs for Table Driven Tests.
You can also use map[string]struct.

Using t.Fatalf in subtests allows you to fail a particular row in a table driven tests but continue with the others seamlessly.
This is very useful when you have multiple checks for a single row in a given table driven test.
This is a massive benefit of subtests since you have fine grained control of what tests you can run.

Use setup() and teardown() when you have to setup some common stuff everytime you run a test in a package.
An extremly common example is running tests on a database, where you have to connect to a database every time.

Table Driven Tests + SubTests + Setup and Teardown combination is extremely powerful.

TestMain() is a special case that will be called before all your functions are called.
It serves purposes like:-
1. Reading custom flags from the console.
2. Setup anything you need.
3. You need to explicity tell tests to run if you add TestMain() to your program

os.Exit() stops your functions immediately. It does not allow even defer() methods to run.
You should you use a shared setup and teardown if your tests take time to run (database).
Else you should run setup() and teardown() on each test.
