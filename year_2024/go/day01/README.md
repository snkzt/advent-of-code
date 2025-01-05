## Process
1. Read the input data: Retrieving raw data from the text.
2. Parse the data: Converting raw data into a structured format.
  - Split each line into two columns.
  - Convert the column values to integers.
3. Process each line:
  - Sort each column with ASC order.
4. Solution for part 1
  - Calculate the difference between the two columns.
  - Add the difference into a variable and return.
5. Solution for part 2
  - Multiply appearance of the numbers from column 1 in column 2
  - Sum the result of the multiply of each number and return. 


## Note
##### Check [GO standard library](https://pkg.go.dev/std) for details about the answers below.
<br>

***Q1:*** What's happeing when opening `.text`file with `os.ReadFile`?
<br>
***A1:*** The output data will contain the entire content of the input file as a single `[]byte` slice.
<br>

***Example:***
<br>
1. Input file
```
12 5
30
   30
```
<br>

2. `[]byte`: The output data will be a slice of bytes that includes all characters in the file, including whitespace and newline characters `\n`.
```
[]byte{49, 50, 32, 53, 10, 51, 48, 10, 32, 32, 32, 51, 48, 10}
```
- 49, 50 corresponds to 1 and 2 ("12") <br>
- 32 corresponds to a space <br>
- 53 corresponds to 5 <br>
- 10 corresponds to a newline (\n) <br>
If we convert this to a string it will look like
```
"12 5\n30\n   30\n"
```
<br>

3. Processing: When splitting the output data into lines with `bytes.Split(inputData, []byte("\n"))`, each line becomes its own `[]byte` slice based on the separator(`\n` in this case) .
```
[]byte("12 5")
[]byte("30")
[]byte("   30")
```
Then `bytes.Fields`run through each line individually, and It splits the slice s around each instance of one or more consecutive white space characters. It will return a slice of subslices of the input or an empty slice if it contains only white space:
`func Fields(s []byte) [][]byte`
```
fields := bytes.Fields([]byte("12 5"))
fmt.Printf("%q\n", fields) // ["12" "5"]

fields := bytes.Fields([]byte("   30"))
fmt.Printf("%q\n", fields) // ["30"]
```
This time, we are using this `Split()` to separate each left and right location ID as a row, then use `Fields()` to make it possible to retrieve each ID then add it back to a slice to sort per column.
<br>
<br>

***Q2:*** How we deal with bigger data?
 <br>
***A2:*** If the input data will be GN+ then use `bufio.Scanner` to read the file. This will read the input file line-by-line and we can avoid high memory consumption and handle data more efficiently.
 <br>
 <br>

***Q3:*** `log.Fataf`'s error message always start with capital but `fmt.Errorf`'s in small letter...?
 <br>
***A3:*** `log.Fatalf` is usually treated as a standalone logging function. The message it logs is often capitalized because log entries are like sentences, starting with a capital letter for readability. On the other hand, `fmt.Errorf` is used to create and wrap errors. These error messages are often lowercase because they are typically appended or combined with other error messages.
