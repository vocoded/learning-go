# learning-go

Many standard approaches for language learning focus on a bottom-up, comprehensive treatment of internals.  This is especially valuable for folks new to programming, but can be a trial for others with development proficiency.  

I've long been curious whether a more practical progression of concepts and constructs would be better suited to the experienced audience.  This project outlines one such course for a language generating quite a bit of press these days - Go.

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

Exercise 1: Write an application that loops the natural numbers to 100 and prints only the values evenly divisible by 5.

Exercise 2: Take the previous application and split the loop logic into a method that accepts an arbitrary upper limit, the required divisor, and returns a comma-separated string of the form “x, y, z, …”

Exercise 3: Write an application that reads a file containing some paragraphs of text.  Read a word from the console and identify the matches in the file (case insensitive).  Print out the matches with context; in other words, print the match and the word immediately before it.  Write to the console the total number of matches found.

Exercise 4: Factor the previous application so that matching is handled by a dedicated class implementing a specific interface.

Exercise 5: Modify the previous application so that it counts the number of unique words in the given file. Words are defined as a non-zero length sequence of alphabetic characters.

Exercise 6: Use the previous application to print a sorted list of the unique words in the given file.

Exercise 7: Write an application that downloads a URL entered from the command line and times how long it takes