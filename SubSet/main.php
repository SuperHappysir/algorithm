<?php
/**
 * 求给定集合获取所有子集
 *
 * @param array $array
 *
 * @return array
 *
 * @usage php main.php
 */
function subsets($array) : array
{
    $n        = count($array);
    $subN     = 2 ** $n;
    $subArray = [];

    $subArray[] = [];
    for ($i = 1; $i < $subN; $i++) {
        $m    = sprintf('%0+' . $n . 'b', $i);
        $tArr = [];
        for ($j = 0; $j < $n; $j++) {
            if ($m{$j} == 1) {
                $tArr[] = $array[$j];
            }
        }
        $subArray[] = $tArr;
    }

    return $subArray;
}

var_dump(subsets([1, 2, 3]));