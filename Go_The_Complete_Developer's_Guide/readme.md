<p>Well, let's star our tour through the Go World!</p>
<p>First of all, you'll need to install Go and choose your favorite code editor. I chose Visual Studio Code.</p>
<p>After installing Go and the Visual Studio Code, you'll need to install an extension inside Visual Studio Code.</p>
<p>I chose the Go Team at Google, which has more than 3million downloads.</p>
<p>So, once you had all this set up... we're able to start.</p>

*** 

<p>We'll start with the classic Hello World, but, with this, we'll explore some aspects from Go itself</p>
<p>So, create a directory, and a file named main.go, and fill it with:</p>

    package main
    import "fmt"
    func main() {
	    fmt.Println("Hello there!")
    }

> PS: Maybe, when you install the extension, and create the file, the
> extension could ask for install other components. Just choose the YES
> option, and be happy! ;)

<p>Now, with this small code, we can start our path to understanding Go.</p>
<p><em>First thing is: How am I supposed to execute this code?</em></p>
<p>The answer is... In your terminal, or even inside the Visual Studio Code terminal, type Go.</p>
<p>This should bring a list of all commands that Go executes. If in Visual Studio Code it doesn't work, close and open again.</p>
<p>Well, inside the command list, you'll notice one that could sounds familiar... the run command
<p>This command will Compile the code and Execute it.</p>
<p>So, you'll need to navigate to the directory where your code is... and run go run main.go
and with this, you should see Hello there! printed.</p> 
<p>You probably will notice too, the command build. Well, this command also could execute the code, but instead of Compile and Execute, it'll only Compile it. So, if you run it, by typing go build main.go, once in your terminal, if you check out the content of your directory, you'll now notice that there's an executable file there, that was generated through this command with our simple program.</p>

***
<p>Now, talking about packages.</p>
<p>In Go, a Package is a collection of common source files.</p>
<p>A package could contain any number of files inside of it, the conditions for this are, they need to be named as fileName.go and inside of each file, you'll need to inform at the first line, which package this files belongs to.</p>
<p>In our code, we are using the: </p>

    package main.
    
<p>But, why we are using 'package main' in our code?</p>
<p>Well, this is because in Go, we have two types of packages:</p>
<p>The Executable Type and The Reusable Type.</p>

#### The Executable Type:
<p>The executable package, when compiled generate a executable file, like mentioned before, when we saw the build command.</p>
  
#### The Reusable Type:
<p>Basically, this package type, could be like a common library with helpers and that kind of stuff...</p>
<p>This type of package will help us by, making us reuse it's code</p>
<p><em>And, when we're coding, how do we know if we are creating a Executable or a Reusable type?</em></p>
<p>Well, it'll depends on the first line of code, inside our file.</p>
<p>Basically, if we write 'package main', we are declaring that, the file will generate an Executable. </p>
<p>Otherwise, should be a Reusable one.</p>
<p>And, last but not least, when we write a package main, we'll need a function main, that will be executed automatically when the program runs.</p>

***

<p>Now, we're kind wondering that the line:</p>

    import "fmt" 

<p>This one, will bring to our file some functionality that we'll need, in order to properly execute our code?</p>
<p>And the answer is that, we're right about this guessing!</p>
<p>The line 'import "fmt" ', will bring to us, the Go one of the standard package libraries. The fmt stands for, Format, which will help us to print things at the console, and help us to debug.</p>
<p>Go has a huge number of standard libraries.</p>
<p><b>We can find them at: golang.org/pkg</b></p>
<p>This import statement is how we could bring to our code, stuff from other files, that were written as Reusable Packages</p>
<p>And, if in some situation, we need to import more than one package into our code, we should do this, this way:</p>
    
    import ("fmt"
    "strings")
<p>We'll inform our packages, inside a parenthesis, and break them into new lines, without coma.</p>

***
<p>In order to make our code run properly, we'll basically need to follow the structure:</p>

    package main
    import "fmt"
    func main() {
	    fmt.Println("Hello there!")
    }
<p>With:</p>
<p>our package declaration</p>
<p>our imports</p>
<p>our function</p>

***
 

> Go, is a Static Type Language.

<p>Which means that, once you create a variable, you need to define a type to it, and the type of the variable will be important.</p>
<p>In Go, when we're creating a variable, we can do this way:</p>

    var card string = "Ace of Spades"
<p>or</p>

    card := "Ace of Spades"
<p>The second way, relies on Go compiler to figure out on what type of value we're trying to work with.</p>
<p>So, this means that, we're creating a variable, named card.</p>
<p>And, the string word, after the variable name, tells the compiler that this variable will accept only string values.</p>

***
<p>In Go, when we create a function, we need to explicit the type of the function, when we define a return of any type.</p>
<p>So, if we create a function like this:</p>

    func newCard() {
	    return "Five of Diamonds"
    }

<p>Go will throw an error, saying that there's too many arguments to return. This happens because we don't informed the return type to the compiler.</p>
<p>So, we need to write the function this way, and make the type of the function explicit:</p>

    func newCard() string {
	    return "Five of Diamonds"
    }

***
  

<p>In Go, we have two types to work with list of data.</p>
<p><b>The Array and The Slice.</b></p>
<p>The Array, has a defined size, that it's immutable and the Slice no, it can grow or shrink at will.</p><p>But, the Slice by definition, is an Array.</p>
<p>Both Array and Slice must be defined with a data type. This way, if we create one of them of the type string, it only will accept strings.</p>
<p>So, in order to create a Slice, we could do in the following way:

    cards := []string{"Ace of Diamonds", newCard()}
<p>This way, we're saying that this Slice, is a Slice of string.</p>
<p>And, in order to add values to it, we could do like this:</p>

    cards = append(cards, "Six of Spades")
<p>This, will return to us a new state of the cards Slice.</p>
<p>And, if we want to iterate through the Slice, we could use a for loop:</p>

    for i, card := range cards {
    fmt.Println(i, card)
    }
<p>Where <b>i</b> is the index of the elements;</p>
<p>The <b>card</b> is the current object that we're iterating over;</p>
<p>The <b>range cards</b> indicates that we want to iterate through;</p>
<p>The <b>fmt.Println(i, card)</b> is the action that we want to execute on each iteration</p>
<p>The <b>:=</b> syntax works fine here, because for each iteration we're kind throwing out the old variable, and redeclaring it.</p>
<p>Slice is zero index based.</p>
<p>When working with Slices, we can retrieve it's value with the following ways:</p>

    cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
    cardSuits[0] => "Spades"

<p>Here, we're saying that we want the value at index zero of the Slice</p>
<p>But, we could retrieve the information in another way, like this:</p>

    cardSuits[0:2] => 0:"Spades", 1:"Diamonds"
<p>This way, will return to us the values of index 0 and index 1, cause it works with value - 1.</p>
<p>Other way to user this approach is:</p>

    cardSuits[:2] => 0:"Spades", 1:"Diamonds"
<p>This way, the Go compiler will infer that we want from index zero until index 2 -1</p>
<p>And, we still can do this way:</p>

    cardSuits[2:] => 2:"Hearts", 3:"Clubs"
<p>This means that, I want from index 2 until the end</p>
<p>Basically, this approach could be explanied like this:</p>

    slice[startIndexIncluding:upToNotIncluding]

> Note: if we declare a variable like 'i', for the index, we must use it
> somewhere into the for loop, otherwise, it'll cause a compilation
> problem.

<p>In go, if you declare a variable, you need to use it.</p>
<p>And, if for some reason you create a for loop that won't need the index value, you can replace it for un <b> _ </b></p>
<p>This underscore means that, we know that should have a variable there, but we don't care about.</p>
<p>This would be like this:</p>

    for _, card := range cards {
	    fmt.Println(i, card)
    }
***
<p>Go is not an OO language, but, it doesn't mean that you can't use this approach with it.</p>
<p>In order to work more or less like and OO approach, we can create a custom type.</p>
<p>For example, we can create a type for our Deck, which will still be a slice of string.</p>
<p>This would be like this:</p>

    type deck []string
<p>Here, we're saying that we're creating a new type, named 'deck', which is a slice of string.</p>
<p>And, we can go further, and create a function for our new type, to help the user to retrieve and print the values. This could be like this:</p>

    func (d deck) print() {
        for i, card := range d {
		    fmt.Println(i, card)
	    }
    }
<p>Where, <b>d deck</b> is called Receiver, where the <b>d</b> is the working variable, a reference to the type who called the function, and <b>deck</b> is a reference of the type which we would attach this method to.</p>
<p>By convention, the receiver should be called with one or two letters, like and abbreviation of the type, but, you're not forbidden to use a name you want.</p>
<p>The receiver works more or less like an extension in C#, when we create a class, with custom methods to extend an object</p>
<p>If we want, we could create a function that receive arguments to work with, and it should be like this:</p>

    func funcName(x type, y type){
    }
<p>And going further, we could return values from our function:</p>

    func funcName(x type, y type) returType {
        return value
    }

<p>And going a bit deep, we could return something like a Tuple, like this:</p>

    func funcName(x type, y type) (returTypeA, returTypeB) {
	    return valueA, valueB
    }
***
<p>In Go, if we someday need to convert one type to another, we could this, this way: type we want (value to convert)</p>
<p>For example, convert a string to a byte slice:</p>

    []byte("Hello there")


***
<p>In order to create a Test in Go, we could create a new file, for exemple deck.go, but, we must name this Test file with a _test, so this way, Go will automatically see this Test file.</p>
<p>This way, our teste file should be named as deck_test.go</p>
<p>And, in order to create a new test function, we should do like this:</p>

    func TestNameFunction(t *testing.T)
    
<p>With the name Test, the compiler will recognize this as test function, and another important thing, is that we need usa a capital letter after the Test, otherwise, we'll have a compilation problem, of malformed name.</p>

***

<p>With Go, we can work with a type know as struct.</p>
<p>Struct are collections of fields</p>
<p>We can create a struct, this way:</p>

    type person struct {
        firstName string
        lastName string
    }

<p>And, we can inform the values to it, in many ways.</p>
<p>One of them is this:</p>

    func main() {
	    person := person{"Andre", "Miranda"}
    }

<p>This way, Go will infer that the first value is related to the first property, and the second value, to the second property.</p>
<p>But,it could be a problem because, we'll always need to inform the values in the right order, and double our attention when creating our struct.</p>
<p>For example, if this happens:</p>

    func main() {
	    person := person{"Miranda", "Andre"}
    }

<p>the values will be changed, same way if we do this:</p>

    type person struct {
	    lastName string
	    firstName string
    }
    
    func main() {
	    person := person{"Andre", "Miranda"}
    }

<p>But, we can inform the values in another way, a way where the order of the values or properties doesn't matter.</p>
<p>This would be like this:</p>

    person := person{firstName: "Andre", lastName: "Miranda"}

<p>This way, no matter the order, we're informing which property we want to fill with the respective value.</p>
<p>And, there's another way that we can initialize a strut.</p>
<p>We could do like this:</p>

    func main() {
	    var p person
    }

<p>This way, Go will initialize the properties inside of the struct with the base value for each type that exists inside of it.</p>
<p>And, doing this way, if we want to fill the properties, we can do this way:</p>

    func main() {
	    var p person
	    p.firstName = "Andre"
	    p.lastName = "Miranda"
    }

<p>We could also, create another struct if we want or need, and embed this second struct inside this fisrt struct.</p>
<p>If we do this, it could be like this:</p>

    type contactInfo struct {
	    email string
	    zipCode int
    }
    
    type person struct {
	    firstName string
	    lastName string
	    contact contactInfo
    }

<p>And, for initialize it, we could do like this:</p>

func main() {

    person := person{
	    firstName: "Andre",
	    lastName: "Miranda",
	    contact: contactInfo{
		    email: "myemail@email.com",
		    zipCode: 000000,
		    },
	    }
    }

<p>or like this:</p>

    func main() {
	    var p person
	    p.firstName = "Andre"
	    p.lastName = "Miranda"
	    p.contact.email = "myemail@email.com"
	    p.contact.zipCode = 000000
    }

<p>And, in cases when we have structs embedded, we can suppress the name for the type.</p>
<p>For example:</p>

    type contactInfo struct {
	    email string
	    zipCode int
    }
    
    type person struct {
	    firstName string
	    lastName string
	    contactInfo
    }

<p>This way, Go will infer that we have a property named contactInfo inside our person struct, that is of the type contactInfo.</p>
<p>And, if we do this way, in order to fill the information, we must do like this:</p>

    func main() {
	    person := person{
	    firstName: "Andre",
	    lastName: "Miranda",
	    contactInfo: contactInfo{
		    email: "myemail@email.com",
		    zipCode: 000000,
		    },
	    }
    }

<p>or like this:</p>

    func main() {
	    var p person
	    p.firstName = "Andre"
	    p.lastName = "Miranda"
	    p.contactInfo.email = "myemail@email.com"
	    p.contactInfo.zipCode = 000000
    }

***
<p>As Go work with structs, when we want to manipulate it's value, we need to access it first.</p>
<p>So, this way, if we do something like this:</p>

    func main() {
	    var p person
	    p.firstName = "Andre"
	    p.lastName = "Miranda"
	    p.contactInfo.email = "myemail@email.com"
	    p.contactInfo.zipCode = 000000
	    p.updateName("Adao")
	    p.print()
    }

    func (p person) updateName(newFirstName string) {
	    p.firstName = newFirstName
    }

    func (p person) print() {
	    fmt.Printf("%+v", p)
    }
<p>We won't see any change, because when we do this, Go fist write the person struct in a place into the memory.</p>
<p>Then, when asked to do the update, Go will do this, but, it will place the new value, in a new memory place.</p>
<p>This way, we'll never see the changes.</p>
<p>So, as said before, if we want to change, and see it happens, we need to work with pointers.</p>
<p>Pointers that work mostly like pointers in C or C++</p>
<p>So in order to create a pointer, we could do this way:</p>

    func main() {
	    var p person
	    p.firstName = "Andre"
	    p.lastName = "Miranda"
	    p.contactInfo.email = "myemail@email.com"
	    p.contactInfo.zipCode = 000000
	    pPointer := &p
	    pPointer.updateName("Adao")
	    p.print()
    }

	func (pointerToPerson *person) updateName(newFirstName string) {
		(*pointerToPerson).firstName = newFirstName
	}

	func (p person) print() {
		fmt.Printf("%+v", p)
	}

<p>Here, we notice two things:</p>
<p><b>pPointer := &p</b></p>
<p><b>pPointer.updateName("Name")</b></p>

	func (pointerToPerson *person) updateName(newFirstName string) {
		(*pointerToPerson).firstName = newFirstName
	}

<p>So, a bit of explanation...</p>
<p><b>pPointer := &variable => this means, give me access to the memory address of this variable
*pPointer => this means, give me the value of this memory address</b></p>

<p>As for the function, the definition of it changed a bit, but it still make sense:</p>
<p>func (pointerToPerson *person) updateName(newFirstName string)</p>
<p>Here, with this function, we're saying that this updateName, could only be called by a pointer, of the type person.</p>
<p>So, when we did this:</p>

	pPointer := &p

<p>We created a pointer, with the person type...</p>
<p>and when we did this:</p>

	pPointer.updateName("Name")

<p>we were saying that a pointer of the type person was calling the updateName function</p>
<p>and when this is executed:</p>

	func (pointerToPerson *person) updateName(newFirstName string) {
		(*pointerToPerson).firstName = newFirstName
	}

<p>we are now receiving kind a reference to the person, inside our function: </p>
<p><b>pointerToPerson *person</b></p>
<p>and, with this, we are now able to access it's value:</p>

	(*pointerToPerson) == var p person

<p>and with all this, this line:</p>
	
	(*pointerToPerson).firstName = newFirstName

<p>became something like this in memory:</p>

	p.firstName.firstName = newFirstName

<p>We can say that:</p>
<p>If you have a memory address, and you want to turn it into a value, you can do this with *memoryAddress</p>

	(*pointerToPerson).firstName = newFirstName

<p>If you have a value, and you want to turn it into a memoryAddress, you can do this with &value</p>

	pPointer := &p

<p>And, Go can also do some inference about our call.</p>
<p>What I mean by this?</p>
<p>Well, I mean that if we do this way:</p>

	p.updateName("Adao")

	func (pointerToPerson *person) updateName(newFirstName string) {
		(*pointerToPerson).firstName = newFirstName
	}

<p>Go will understand that our variable, of the type person, is able to call a function, which has a pointer receiver of the same type, in thi case, type person</p>