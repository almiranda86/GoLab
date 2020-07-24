
Well, let's star our tour throught the Go World!

Firt of all, you'll need to install Go and choose your favorite code editor. I choosed Visual Studio Code.

After installing Go and the Visual Studio Code, you'll need to install an extension inside Visual Studio Code.

I choosed the Go Team at Google, which has more then 3million downloads.

So, once you had all this set up... we're able to start.

==================================================================================================================

We'll start with the classic Hello World, but, with this, we'll explore some aspects from Go itself

So, create a directory, and a file named main.go, and fill it with:

package main

import "fmt"

func main() {
	fmt.Println("Hello there!")
}


PS: Maybe, when you intall the extension, and create the file, the extension could ask for install other components. Just choose the YES option, and be happy! ;)


Now, with this small code, we can start our path to understandin Go.

Firt thing is: How am I suposed to execute this code?
The answer is... In your terminal, or even inside the Visual Studio Code terminal, type Go
This should bring a list of all commands tha Go executes. If in Visual Studio Code it doesn't work, close and open again.

Well, inside the command list, you'll notice one that could sounds familiar... the run command
This command will Compile the code and Execute it.

So, you'll need to navigate to the directory where your code is... and run go run main.go
and with this, you should see Hello there! printed.
