#HOW TO DEFINE A FLAG
defining a flag in go using the flag package

## one

var usernamePtr = flag.StringVar("username", "guest","description of the guest")

## two

flag.Parse()
fmt.Println("Username:", \*username)
