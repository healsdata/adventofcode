<?php

require_once 'functions.php';

$input = "7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9";


$generator = inputGenerator($input);
$generator = fgetsGenerator(__DIR__ . DIRECTORY_SEPARATOR . "day2.txt");

$totalSafe = 0;

function isSafe(array $levels): bool
{
    $asc = $levels;

    sort($asc);

    if (!($asc == $levels || array_reverse($asc) == $levels)) {
        return false;
    }

    $lastLevel = null;
    foreach ($levels as $level) {
        if (is_null($lastLevel)) {
            $lastLevel = $level;
            continue;
        }

        $diff = abs($lastLevel - $level);

        if ($diff < 1 || $diff > 3) {
            return false;
        }

        $lastLevel = $level;
    }

    return true;
}


foreach ($generator as $report) {
    $levels = explode(" ", $report);

    if (isSafe($levels)) {
        $totalSafe++;
        continue;
    }

    foreach ($levels as $i => $_) {
        $dampened = $levels;
        unset($dampened[$i]);

        if (isSafe(array_values($dampened))) {
            $totalSafe++;
            continue 2;
        }
    }


}

print $totalSafe;;