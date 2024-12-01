<?php

require_once 'functions.php';

$input = "3   4
4   3
2   5
1   3
3   9
3   3";


$generator = inputGenerator($input);
$generator = fgetsGenerator(__DIR__ . DIRECTORY_SEPARATOR . "day1.txt");

$list1 = [];
$list2 = [];

$totalDistance = 0;

foreach ($generator as $line) {
    $values = preg_split("/\s+/", $line);

    $list1[] = $values[0];
    $list2[] = $values[1];
}

sort($list1);
sort($list2);

foreach ($list1 as $key => $value1) {
    // $value2 = $list2[$key];
    // $distance = abs($value1 - $value2);
    // $totalDistance += $distance;

    $matches = sizeof(array_keys($list2, $value1));
    $totalDistance += ($value1 * $matches);
}

print $totalDistance;