# Takeaways & Points to Revise
These are the key takeaways and point to revise/reiterate of the learning from [Go Tutorial](https://go.dev/doc/tutorial/).

### Packages
- fmt --> String format
    - Sprintf --> Parse string with variables
    - Println --> Print to console

- log --> Print console logs
    - SetPrefix --> set prefix for log messages
    - SetFlags --> Timestamps?

- errors --> Generate error messages + ?
    - New --> create new error with a message

- math/rand --> Generate random numbers?

- time --> Time operations like Now()

- testing --> Builtin testing package
    - *testing.T pointer

- regexp --> 
    - want how is this used?
    - MustCompile()
    - Back ticks - purpose?
    - MatchString() -- Check if a string that matches regex is found in
      the argument passed

### Tips
#### 1. Short Declaration
:= can be used to declare and assign a variable

#### 2. nil
Can be used to validate if a variable is empty?

#### 3. Slices -- Arrays --> Check
- Declaration:

    ```go
    []string{
            "item1",
            "item2",
            "item3",
    }
    ```

    Note: Comma after last item in array

#### 4. init function
- init is called after all the variable declarations in the package
  have evaluated their initializers, and those are evaluated only
  after all the imported packages have been initialized. 

    ```go
    func init() {
        if user == "" {
            log.Fatal("$USER not set")
        }
        if home == "" {
            home = "/home/" + user
        }
        if gopath == "" {
            gopath = home + "/go"
        }
        // gopath may be overridden by --gopath flag on command line.
        flag.StringVar(&gopath, "gopath", gopath, "override default GOPATH")
    }
    ```

5. for Loop --> refer effective go
_ blank identifier
range of slices or maps
use ```make()``` to initialize map/slice/chan

6. Test Cases should cover both corerct values and incorrect values
    - It should expect errors are also handled properly, meaning a test
        case should include test functions to check if the errors are
        returned properly.
    - For a the test case for the wrong argument, if the function does
        not return proper error, it means the function does not handle the
        error and the test case failed.   