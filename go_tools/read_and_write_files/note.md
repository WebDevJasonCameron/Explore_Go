# Notes

Instructions found at the following site

https://www.geeksforgeeks.org/how-to-read-and-write-the-files-in-golang/

Golang offers a vast inbuilt library that can be used to perform read and write operations on files. In order to read from files on the local system, the io/ioutil module is put to use. The io/ioutil module is also used to write content to the file.
The fmt module implements formatted I/O with functions to read input from the stdin and print output to the stdout. The log module implements simple logging package. It defines a type, Logger, with methods for formatting the output. The os module provides the ability to access native operating-system features. The bufio module implements buffered I/O which helps to improve the CPU performance.


os.Create() : The os.Create() method is used to creates a file with the desired name. If a file with the same name already exists, then the create function truncates the file.
ioutil.ReadFile() : The ioutil.ReadFile() method takes the path to the file to be read as it’s the only parameter. This method returns either the data of the file or an error.
ioutil.WriteFile() : The ioutil.WriteFile() is used to write data to a file. The WriteFile() method takes in 3 different parameters, the first is the location of the file we wish to write to, the second is the data object, and the third is the FileMode, which represents the file’s mode and permission bits.
log.Fatalf : Fatalf will cause the program to terminate after printing the log message. It is equivalent to Printf() followed by a call to os.Exit(1).
log.Panicf : Panic is just like an exception that may arise at runtime. Panicln is equivalent to Println() followed by a call to panic(). The argument passed to panic() will be printed when the program terminates.
bufio.NewReader(os.Stdin) : This method returns a new Reader whose buffer has the default size(4096 bytes).
inputReader.ReadString(‘\n’) : This method is used to read user input from stdin and reads until the first occurrence of delimiter in the input, returning a string containing the data up to and including the delimiter. If an error is encountered before finding a delimiter, it returns the data read before the error and the error itself.
Example 1: Use the offline compiler for better results. Save the file with .go extension. Use the command below command to execute the program.


```go run filename.go```


###
