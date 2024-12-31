## Process
1. Read the input data: Retrieving raw data from the text.
2. Parse the data: Converting raw data into a structured format.
  - Split each line into a slice.
  - Convert the column values to integers.
3. Process and calculate each line:
  - Loop each slice and compare the elements if
    - The elements are either all increasing or all decreasing.
    - Any two adjacent elements differ by at least one and at most three.
    - Add up if a report is safe based on the loop result.
  - Return the safe count result.
