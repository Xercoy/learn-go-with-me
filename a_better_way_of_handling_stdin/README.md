# A better way of handling stdin

I usually pipe content to programs that depend on some kind of input from `stdin`. Annoyingly enough, when running the program manually and forgetting to pipe content, it hangs.

To get past this, I'll `Ctrl-c` and correct the command. In terms of automatically handling this, I would skip relevant code by setting a flag. Simple, right?

What if I could have my code detect when nothing is piped to `stdin` and properly handle both situations? This would drastically reduce the likelihood of both human and automated error, not to mention increase code quality.

The key to this whole blog post is determining the file mode of `stdin` which is represented by the [`os.FileMode`](https://golang.org/src/os/types.go?s=1044:1064#L20) type. The next two lines of code are very useful:

```
// Retrieve file information of stdin (os.FileInfo)
stdinFileInfo, _ := os.Stdin.Stat()

// Print string representation of stdin's file mode.
fmt.Println(stdinFileInfo.Mode().String())
```

If we print the string representation of `stdin`'s file mode without piping anything, we get:
```
Dcrw--w----
```
the notation of the file mode, `Dc`, describes two things: `D`, that the file is a device, and more specifically, `c`, that the device is a character device. The remaining characters represent the file's permission bits.

Now, if we observe the file mode of `stdin` when content is being piped, we get:
```
prw-rw----
```
 this time, the file mode describes a pipe, denoted by `p`.

View the example in the repo [here](https://github.com/Xercoy/learn-go-with-me/tree/master/be_smart_with_stdin/example1).

######The Hard Work is Done

We could identify the the two distinct file modes by performing a binary AND operation against constants defined in the `os` package which represent them (view them [here](https://golang.org/pkg/os/#FileMode)). The constants can be used as a mask to determine which file mode is present.

For pipes (true):
```
// File Mode of os.Stdin :  prw-rw----
// os.ModeNamedPipe      :  p---------
// Binary AND Result     :  p---------

if (stdinFileInfo.Mode() & os.ModeNamedPipe != 0) {
    fmt.Printf("Input From Pipe!")
    }
    ```

Without piped content (true):
```
//File Mode of os.Stdin :  Dcrw--w----
//os.ModeCharDevice     :   c---------
//Binary AND Result     :   c---------

if (stdinFileInfo.Mode() & os.ModeCharDevice != 0) {
    fmt.Printf("Input From Terminal!")
    }
    ```

Run the example via [playground](https://play.golang.org/p/Jk_8UoKLhX) (hardcoded to not have content piped) or view it in the [repo](https://github.com/Xercoy/learn-go-with-me/tree/master/be_smart_with_stdin/example2).

######In Practice

The ability to identify when content is being piped gives rightful control back to the developer who can handle each situation accordingly.

This simplified example will output a greeting with a random name if one isn't provided via `stdin`. Run the full, more detailed example via [playground](https://play.golang.org/p/G5jI8s9LCd) or view it in the [repo](https://github.com/Xercoy/learn-go-with-me/tree/master/be_smart_with_stdin/example3). Note that the [randomization in playground isn't possible](https://blog.golang.org/playground).
```
var name string

// Get a name from stdin if possible
if (stdinFileInfo.Mode() & os.ModeNamedPipe != 0) {

    stdinContent, _ := ioutil.ReadAll(os.Stdin)
        name = string(stdinContent)
	} else {

    defaultNames := []string{"user","creeper"}

    // Pseudorandom magic
        rand.Seed(time.Now().UnixNano())
	    randomIndex := rand.Intn(len(defaultNames))

    name = defaultNames[randomIndex]
    }

fmt.Printf("Hello, %s!", name)
```

###### With piped content:
```
$ printf "Corey" | go run main.go
> Hello, Corey!
```

###### Without piped content:
```
$ go run main.go
> Hello, creeper!
```

Feel free to improve or correct any content, or head to the [repo](https://github.com/Xercoy/learn-go-with-me) to open an issue and suggest a topic. Happy coding!