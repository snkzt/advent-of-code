
import os
import sys

def parse_line(line):
    """
    Parses a line of input and extracts two integer values.
    Returns a tuple of two integers (num1, num2).
    Raises ValueError if a field is invalid or not numeric.
    """
    fields = [field for field in line.split() if field]

    num1 = int(fields[0]) if len(fields) > 0 else 0
    num2 = int(fields[1]) if len(fields) > 1 else 0

    if len(fields) > 0 and not fields[0].isdigit() and fields[0] != "":
        raise ValueError(f"Invalid value in column 1: {fields[0]}")
    if len(fields) > 1 and not fields[1].isdigit() and fields[1] != "":
        raise ValueError(f"Invalid value in column 2: {fields[1]}")

    return num1, num2

def parse_input_data(input_data):
    column1 = []
    column2 = []

    # Split input data into lines and process each line.
    for line in input_data.splitlines():
        line = line.strip() # Removing leading/trailing whitespace.
        if not line: # Skip empty lines.
            continue

        try:
            num1, num2 = parse_line(line)
        except ValueError as e:
            raise ValueError(f"Error parsing line: {str(e)}")
        
        column1.append(num1)
        column2.append(num2)

    sort_columns(column1, column2)

    return column1, column2


def sort_columns(column1, column2):
    """
    Sort two ists in-place, ascending order.
    """
    column1.sort()
    column2.sort()

# Puzzle part 2
def calculate_similarity_score(column1, column2):
    """
    Calculate similarity score by multiplying the number in column1
    by the frequency of its appearance in column2.
    Then add up the result and return an integer.
    """
    count_map = {}

    for num in column2:
        if num not in count_map:
            # Initialise the key with 0 if it doesn't exist
            count_map[num] = 0
        
        count_map[num] += 1
    
    similarity_score = 0
    for num in column1:
        if num in count_map:
            similarity_score += count_map[num]*num
    
    return similarity_score

# Puzzle part 1
def calculate_total_distance(column1, column2):
    """
    Calculate the total distance as the absolute difference between
    the two values in a row from the two sorted columns retrieved from
    the input data.
    Return the total distance as an integer.
    """
    total_distance = sum(abs(col1-col2) for col1, col2 in zip(column1, column2))
    return total_distance

def read_file_content(file_path):
    """
    Reads the content of a file and returns it as a string.
    Raises an exception if the file does not exist or cannt be read. 
    """
    if not os.path.exists(file_path):
        raise FileNotFoundError(f"File does not exist: {file_path}")
    
    with open(file_path, "r", encoding="utf-8") as file:
        data = file.read()

    return data


def main():
    file_path = "../../shared/inputs/day01/input.txt"

    try:
        input_data = read_file_content(file_path)
        sorted_array1, sorted_array2 = parse_input_data(input_data)
        total_distance = calculate_total_distance(sorted_array1, sorted_array2)
        similarity_score = calculate_similarity_score(sorted_array1, sorted_array2)
        print(f"Total distance: {total_distance}")
        print(f"Similarity Score: {similarity_score}")
    except Exception as e:
        print(f"Error: {str(e)}")


if __name__ == "__main__":
    main()