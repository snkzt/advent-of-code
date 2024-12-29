## Python standard library information
https://docs.python.org/3/library/os.path.html
<br>
<br>


## Note
- ***Where is the start point of execution*** <br>
In Python, the interpreter executes code from the top to the bottom of a script, starting at the first line of the file. If no specific entry point like main() is defined, Python will simply run any code that is not inside a function or class. This means the order of execution is determined by the structure of the file.
<br>

- ***How Python Knows Where to Start*** <br>
When you run a Python script (python script.py):
The Python interpreter starts reading the script from the first line.
It executes all top-level statements sequentially.
Functions or classes defined in the script are not executed when they are defined. Instead, they are stored in memory for use when explicitly called.
If there is an `if __name__ == "__main__":` block, Python checks if the script is being run directly, and if so, executes the code inside this block.
<br>

- ***Read a file*** <br>
`with open()` is suitable for reading small- to medium-sized files. For large files, it should be read in chunks or line by line to avoid memory overload: 
```
with open(file_path, "r", encoding="utf-8") as file:
    for line in file:
        print(line.strip())
```
<br>

- ***Ternary Operator Syntax*** <br>
The general format of Python's ternary operator is:
`<value_if_condition_true> if <condition> else <value_if_condition_false>`
<br>

- ***List comprehension*** <br>
  - `fields` = [field for field in line.split() if field]`: A compact way to create a list by iterating over an iterable.<br>
  - `line.split()`: This splits the string line into a list of substrings separated by whitespace.<br>
  - `for field in ...`: This iterates over each element (substring) in the list produced by line.split().<br>
  - `field`: This is the value of each substring during the iteration. It is directly added to the resulting list.<br>
  - `if field`: This checks if field is not empty. Only non-empty fields are included in the resulting list.<br>

***Example***

```
Input: line = " 12   5 "

# Without the filter:
fields = [field for field in line.split()]
# Output: ['12', '', '', '5']

# With the filter (if field):
fields = [field for field in line.split() if field]
# Output: ['12', '5']
```
