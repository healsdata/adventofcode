<?php

require_once 'functions.php';

$input = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))";


$generator = inputGenerator($input);
$generator = fgetsGenerator(__DIR__ . DIRECTORY_SEPARATOR . "day3.txt");

$total = 0;

function mul(int $a, int $b): int {
    return $a * $b;
}

$enabled = true;

foreach ($generator as $key => $value) {

    $matches = [];
    preg_match_all('/mul\(\d+,\d+\)|do(?:n\'t)?\(\)/i', $value, $matches);

    foreach ($matches[0] as $match) {

        if ($match == "don't()") {
            $enabled = false;
            continue;
        } elseif ( $match == "do()") {
            $enabled = true;
            continue;
        }

        if ($enabled) {
            $total += eval("return " . $match . ";");
        }
    }
}

print $total;