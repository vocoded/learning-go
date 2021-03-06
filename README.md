# learning-go

The traditional approach to programming language learning usually focuses on a bottom-up, comprehensive study of internals.  For example, you not only see how to add and multiply numbers, you are buried in the detail of how to arithmetic left and right shift them too.  This can be valuable for folks new to programming.  It doesn't serve those with development proficiency nearly as well.  

I've been curious whether a more practical progression of concepts and constructs would be better suited to the experienced audience.  This project outlines one such course for a language generating quite a bit of press these days - Go.

## Basics

1. Program structure - basic syntax and layout (files, braces, main(), import, headers)
1. Control flow - if/for/while, true/false, nullity
1. Variable declaration and scope
1. Basic I/O - console in/out, formatted printing
1. Debugging/repl - how to compile, run and step
1. Function declaration and usage - parameter passing, return values
1. Strings and string manipulation - formatting, concat, substring, indexOf, regex
1. Arrays, arraylists, and dictionaries - declaration and mutation, operations (map, reduce, filter, find)
1. Intermediate I/O - file read/write
1. Namespacing/packaging
1. Basic classes (if supported) - declaration, inheritance/mixins, dispatch/polymorphism, visibility, static vs instance
1. Date and time

## Exercises

Exercise 1: Write an application that loops the natural numbers to 100 and prints only the values evenly divisible by 5.

Exercise 2: Take the previous application and split the loop logic into a method that accepts an arbitrary upper limit, the required divisor, and returns a comma-separated string of the form “x, y, z, …”

Exercise 3: Write an application that reads a file containing some paragraphs of text.  Read a word from the console and identify the matches in the file (case insensitive).  Print out the matches with context; in other words, print the match and the word immediately before it.  Write to the console the total number of matches found.

Exercise 4: Factor the previous application so that matching is handled by a dedicated class implementing a specific interface.

Exercise 5: Modify the previous application so that it counts the number of unique words in the given file. Words are defined as a non-zero length sequence of alphabetic characters.

Exercise 6: Use the previous application to print a sorted list of the unique words in the given file.

Exercise 7: Write an application that downloads a URL entered from the command line and times how long it takes.

Exercise 8: Create a package (or equivalent) with a utility that reads input from the command line.  Call this method from the main program and print the reverse of the string read from the command line.

Exercise 9: Use the new utility method to get a command line string, then write a file containing all the unique substrings of length 3.

Exercise 10: Write an application that iterates random numbers to find a multiple of 99 less than 1000, on a background thread, and print the result.
