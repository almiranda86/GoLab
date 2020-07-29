
Well, let's star our tour through the Go World!

First of all, you'll need to install Go and choose your favorite code editor. I chose Visual Studio Code.

After installing Go and the Visual Studio Code, you'll need to install an extension inside Visual Studio Code.

I chose the Go Team at Google, which has more than 3million downloads.

So, once you had all this set up... we're able to start.

==================================================================================================================

We'll start with the classic Hello World, but, with this, we'll explore some aspects from Go itself

So, create a directory, and a file named main.go, and fill it with:

package main

import "fmt"

func main() {
	fmt.Println("Hello there!")
}


PS: Maybe, when you install the extension, and create the file, the extension could ask for install other components. Just choose the YES option, and be happy! ;)

Now, with this small code, we can start our path to understanding Go.

First thing is: How am I supposed to execute this code?
The answer is... In your terminal, or even inside the Visual Studio Code terminal, type Go
This should bring a list of all commands that Go executes. If in Visual Studio Code it doesn't work, close and open again.

Well, inside the command list, you'll notice one that could sounds familiar... the run command
This command will Compile the code and Execute it.

So, you'll need to navigate to the directory where your code is... and run go run main.go
and with this, you should see Hello there! printed.

You probably will notice too, the command build. Well, this command also could execute the code, but instead of Compile and Execute, it'll only Compile it. So, if you run it, by typing go build main.go, once in your terminal, if you check out the content of your directory, you'll now notice that there's an executable file there, that was generated through this command with our simple program.

==================================================================================================================

Now, talking about packages.

In Go, a Package is a collection of common source files.

A package could contain any number of files inside of it, the conditions for this are, they need to be named as fileName.go and inside of each file, you'll need to inform at the first line, which package this files belongs to.

In our code, we are using the 'package main'.

But, why we are using 'package main' in our code?

Well, this is because in Go, we have two types of packages:
The Executable Type and The Reusable Type.

The Executable Type:
The executable package, when compiled generate a executable file, like mentioned before, when we saw the build command.

The Reusable Type:
Basically, this package type, could be like a common library with helpers and that kind of stuff...
This type of package will help us by, making us reuse it's code

And, when we're coding, how do we know if we are creating a Executable or a Reusable type?

Well, it'll depends on the first line of code, inside our file.
Basically, if we write 'package main', we are declaring that, the file will generate an Executable. Otherwise, should be a Reusable one.


And, last but not least, when we write a package main, we'll need a function main, that will be executed automatically when the program runs.

==================================================================================================================

Now, we're kind wondering that, the line 'import "fmt" ', will bring to our file some functionality that we'll need, in order to properly execute our code...
And, we're right about this guessing!

The line 'import "fmt" ', will bring to us, the Go one of the standard package libraries. The fmt stands for, Format, which will help us to print things at the console, and help us to debug.

Go has a huge number of standard libraries.
We can find them at: golang.org/pkg

This import statement is how we could bring to our code, stuff from other files, that were written as Reusable Packages

==================================================================================================================

In order to make our code run properly, we'll basically need to follow the structure:

package main

import "fmt"

func main() {
	fmt.Println("Hello there!")
}


With:
our package declaration
our imports
our function

==================================================================================================================

Go, is a Static Type Language.
Which means that, once you create a variable, you need to define a type to it, and the type of the variable will be important.

In Go, when we're creating a variable, we can do this way:
var card string = "Ace of Spades"
or
card := "Ace of Spades"

The second way, relies on Go compiler to figure out on what type of value we're trying to work with.

So, this means that, we're creating a variable, named card.
And, the string word, after the variable name, tells the compiler that this variable will accept only string values.

==================================================================================================================

In Go, when we create a function, we need to explicit the type of the function, when we define a return of any type.

So, if we create a function like this:
func newCard() {
	return "Five of Diamonds"
}

Go will throw an error, saying that there's too many arguments to return. This happens because we don't informed the return type to the compiler.


So, we need to write the function this way, and make the type of the function explicit:
func newCard() string {
	return "Five of Diamonds"
}

==================================================================================================================

In Go, we have two types to work with list of data.

The Array and The Slice.

The Array, has a defined size, that it's immutable and the Slice no, it can grow or shrink at will. But, the Slice by definition, is an Array.

Both Array and Slice must be defined with a data type. This way, if we create one of them of the type string, it only will accept strings. 

So, in order to create a Slice, we could do in the following way:
cards := []string{"Ace of Diamonds", newCard()}

This way, we're saying that this Slice, is a Slice of string.

And, in order to add values to it, we could do like this:
cards = append(cards, "Six of Spades")

This, will return to us a new state of the cards Slice.

And, if we want to iterate through the Slice, we could use a for loop:

for i, card := range cards {
 fmt.Println(i, card)
}

Where 'i' is the index of the elements;
The 'card' is the current object that we're iterating over;
The 'range cards' indicates that we want to iterate through
The 'fmt.Println(i, card)' is the action that we want to execute on each iteration

The ':=' syntax works fine here, because for each iteration we're kind throwing out the old variable, and redeclaring it.

Note: if we declare a variable like 'i', for the index, we must use it somewhere into the for loop, otherwise, it'll cause a compilation problem.
In go, if you declare a variable, you need to use it.

==================================================================================================================

Go is not an OO language, but, it doesn't mean that you can't use this approach with it.

In order to work more or less like and OO approach, we can create a custom type.

For example, we can create a type for our Deck, which will still be a slice of string.

This would be like this:
type deck []string

Here, we're saying that we're creating a new type, named 'deck', which is a slice of string

And, we can go further, and create a function for our new type, to help the user to retrieve and print the values. This could be like this:
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

Where, 'd deck' is called Receiver, where the 'd' is the working variable, a reference to the type who called the function, and 'deck' is a reference of the type which we would attach this method to.

By convention, the receiver should be called with one or two letters, like and abbreviation of the type, but, you're not forbidden to use a name you want.