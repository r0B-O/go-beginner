### Packages
fmt --> String format
    - Sprintf --> Parse string with variables
    - Println --> Print to console

log --> Print console logs
    - SetPrefix --> set prefix for log messages
    - SetFlags --> Timestamps?

errors --> Generate error messages + ?
    - New --> create new error with a message

math/rand --> Generate random numbers?

time --> Time operations like Now()

### Tips
1. Short Declaration
:= can be used to declare and assign a variable

2. nil
Can be used to validate if a variable is empty?

3. Slices -- Arrays --> Check
Declaration - []string{
    "item1",
    "item2",
    "item3",
}

Note: Comma after last item in array

4. init function

init is called after all the variable declarations in the package have evaluated their initializers, and those are evaluated only after all the imported packages have been initialized. 

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

5. for Loop --> refer effective go
_ blank identifier
range of slices or maps
use make() to initialize map/slice/chan