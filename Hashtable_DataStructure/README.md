[![Open in Visual Studio Code](https://classroom.github.com/assets/open-in-vscode-f059dc9a6f8d3a56e377f745f24479a46679e63a5d9fe6f495e02850cd0d8118.svg)](https://classroom.github.com/online_ide?assignment_repo_id=5536843&assignment_repo_type=AssignmentRepo)
# Hashtable
For this assignment, you will implement a basic hash table that associates strings with ints in Go.

## Requirements
You will need to implement the type and methods in the `hashtable` package. 
- Don't modify the signatures of the exported methods, but do create your own unexported (start with a lowercase letter) helper functions.
- You're welcome to implement the `Hashtable` struct however you like, you shouldn't need to export any fields you place into the struct (all the fields in the struct should have a lowercase name to not be exported).
- Our implementation only uses strings as keys and ints as values, don't change that.
- There are some very basic sanity check functions in the `main.go` file, your implementation should build and then run and not produce any output, output is only generated when something is wrong. Passing those "tests" doesn't necessarily mean your implementation is correct though so you might want to write your own tests that cover additional cases.

### For 5010 Students Only
Your additional task for this assignment will be to implement EITHER [_separate chaining_](https://en.wikipedia.org/wiki/Hash_table#Separate_chaining) OR [_open addressing_](https://en.wikipedia.org/wiki/Hash_table#Open_addressing). For separate chaining, you don't need to use a specific data structure, feel free to pick what you'd like. For open addressing, you're welcome to use any strategy, linear probing is totally fine.
