<?php
declare(strict_types=1);

function parseLine(string $line): array {
    $fields = preg_split('/\s+/', $line);
    $num1 = isset($fields[0]) ? intval($fields[0]) : 0;
    $num2 = isset($fields[1]) ? intval($fields[1]) : 0;

    if (!is_numeric($fields[0]) && !empty($fields[0])) {
        throw new Exception("Invalid number in column 1: " . $fields[0]);
    }

    if (!is_numeric($fields[1]) && !empty($fields[1])) {
        throw new Exception("Invalid number in column 2: " . $fields[1]);
    }

    return [$num1, $num2];
}

function sortColumns(array &$colmn1, array &$colmn2): void {
    sort($colmn1);
    sort($colmn2);
}

function calculateTotalDistance(string $inputData): int {
    $colmn1 = [];
    $colmn2 = [];
    $lines = explode("\n", $inputData);

    foreach($lines as $line) {
        $line = trim($line); // Remove start/end \n, \r, \t, and spaces.
        if ($line === '') continue; // Skip empty lines.

        try {
            list($num1, $num2) = parseLine($line);
        } catch (Exception $e) {
            throw new Exception("Error parsing line: " . $e->getMessage());
        }

        // Need [] to add values to arrays.
        $colmn1[] = $num1;
        $colmn2[] = $num2;
    }

    sortColumns($colmn1, $colmn2);

    $totalDistance = 0;
    for ($i = 0; $i < count($colmn1); $i++) {
        $totalDistance += Abs($colmn1[$i]-$colmn2[$i]);
    }

    return $totalDistance;
}

function readFileContent(string $filePath): string {
    if (!file_exists($filePath)) {
        throw new Exception("File does not exist: ". $filePath);
    }

    $data = file_get_contents($filePath);
    if ($data === false) {
        throw new Exception("Failed to read file: ". $filePath);
    }

    return $data;
}

try {
    $filePath = "../../shared/inputs/day01/input.txt";
    $inputData = readFileContent($filePath);
    $totalDistance = calculateTotalDistance($inputData);

    echo "Total Distance: " . $totalDistance . PHP_EOL;
} catch (EXCEPTION $e) {
    echo "Error: " . $e->getMessage() . PHP_EOL;
}
