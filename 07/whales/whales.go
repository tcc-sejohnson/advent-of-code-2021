package whales

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func PartOne(str string) string {
	cost, err := EfficientAlignmentFromString(str, Median, MedianCost)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("Part One: The crabs have aligned! It cost them... dearly. %v fuel consumed.", cost)
}

func PartTwo(str string) string {
	cost, err := EfficientAlignmentFromString(str, Mean, MeanCost)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("Part Two: The crabs have aligned! It has cost them... most dearly. %v fuel consumed.", cost)
}

func EfficientAlignmentFromString(str string, alignmentFunc func([]int) (int, int), costFunc func(alignmentPosition, currentPosition int) int) (int, error) {
	strs := strings.Split(str, ",")
	ints := make([]int, len(strs))
	for i, num := range strs {
		parsed, err := strconv.Atoi(num)
		if err != nil {
			return 0, err
		}
		ints[i] = parsed
	}
	return EfficientAlignment(ints, alignmentFunc, costFunc), nil
}

func EfficientAlignment(ints []int, alignmentFunc func([]int) (int, int), costFunc func(alignmentPosition, currentPos int) int) int {
	alignmentOne, alignmentTwo := alignmentFunc(ints)
	if alignmentOne == alignmentTwo {
		return CostToMove(ints, alignmentOne, costFunc)
	}
	costOne, costTwo := CostToMove(ints, alignmentOne, costFunc), CostToMove(ints, alignmentTwo, costFunc)
	if costOne < costTwo {
		return costOne
	}
	return costTwo
}

func CostToMove(ints []int, alignmentPosition int, costFunc func(alignmentPosition, currentPos int) int) int {
	cost := 0
	for _, num := range ints {
		cost += costFunc(alignmentPosition, num)
	}
	return cost
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

// Median gets the median in a slice of ints. If the slice has an even length,
// this will be the two middle numbers, which could be different. If it has an odd length, it
// will return the actual middle number twice. This doesn't handle edge cases.
func Median(ints []int) (first, second int) {
	sort.Ints(ints)
	lenInts := len(ints)
	isOdd := lenInts%2 == 1
	if isOdd {
		medianIndex := (lenInts / 2) + 1 // "floor"
		return ints[medianIndex], ints[medianIndex]
	}
	medianIndex := lenInts / 2
	return ints[medianIndex], ints[medianIndex+1]
}

func MedianCost(alignmentPosition, currentPosition int) int {
	return abs(alignmentPosition - currentPosition)
}

func Mean(ints []int) (first, second int) {
	sum := 0
	for _, num := range ints {
		sum += num
	}
	mean := sum / len(ints)
	return mean, mean + 1
}

func MeanCost(alignmentPosition, currentPosition int) int {
	distanceFromAlignment := abs(alignmentPosition - currentPosition)
	costToAlignment := (distanceFromAlignment * (distanceFromAlignment + 1)) / 2
	return costToAlignment
}

var ChallengeInput string = "1101,1,29,67,1102,0,1,65,1008,65,35,66,1005,66,28,1,67,65,20,4,0,1001,65,1,65,1106,0,8,99,35,67,101,99,105,32,110,39,101,115,116,32,112,97,115,32,117,110,101,32,105,110,116,99,111,100,101,32,112,114,111,103,114,97,109,10,90,350,371,573,395,1345,2,660,190,88,16,88,168,1148,336,190,546,531,734,686,87,502,375,722,69,639,936,592,1084,264,299,287,603,1109,485,1081,1481,981,356,437,879,259,6,142,194,1428,1264,543,590,167,43,63,155,114,1061,594,1823,710,1607,305,457,135,277,302,162,75,95,1334,320,892,55,1080,5,390,1275,301,827,597,385,208,79,83,344,954,426,811,79,130,227,14,448,98,136,1000,408,858,263,144,860,552,1025,63,319,178,474,1709,234,1452,664,966,295,321,62,132,427,179,1705,110,34,373,367,1110,80,1842,659,268,10,791,378,1390,395,42,364,404,481,127,243,332,254,55,513,335,14,167,787,242,176,65,923,750,1281,583,124,139,146,453,228,418,860,975,75,84,132,945,1257,988,1179,92,48,1631,1267,689,243,950,1389,417,23,865,54,365,1253,476,952,1337,113,1226,383,353,33,90,249,4,512,206,501,465,821,257,1668,330,203,16,817,224,339,1389,1331,1230,538,221,1124,48,365,582,206,191,228,0,653,198,32,661,386,669,213,810,842,192,1451,51,21,609,203,557,8,124,339,597,273,299,187,1753,329,335,767,1404,306,138,192,580,1069,298,7,166,358,968,838,985,267,372,631,1597,113,803,523,567,405,50,30,1254,965,125,438,879,723,4,81,454,239,1495,394,677,424,519,224,1307,744,44,134,480,949,535,603,257,388,350,479,293,471,95,94,312,8,179,7,154,383,112,26,1694,15,245,435,364,148,594,778,316,1471,670,3,1021,64,142,97,500,58,124,311,392,489,277,863,859,1549,64,1759,116,258,245,595,108,800,89,29,1171,318,36,1529,691,238,622,191,130,1016,408,35,0,1078,186,95,83,287,188,275,1385,198,4,697,553,583,98,1506,1351,166,330,925,230,3,147,748,640,733,355,330,33,1084,753,53,690,245,436,1028,343,533,361,779,328,409,744,414,669,568,235,76,244,843,165,197,1693,6,18,110,48,279,832,702,32,1599,685,245,212,24,124,300,177,20,6,1035,1721,767,138,1116,3,296,1042,1335,347,215,377,1028,192,220,475,1323,9,663,738,88,367,187,56,263,19,80,4,466,1,696,128,571,1215,981,58,368,693,333,40,149,46,252,532,12,526,1171,302,112,1017,1262,807,332,16,715,569,1184,158,570,277,31,414,572,848,1633,784,357,529,286,1510,109,455,902,43,203,443,175,298,484,607,961,1114,295,9,781,487,183,846,336,577,6,723,1369,1484,301,1366,240,641,1937,14,354,458,22,1567,1169,683,544,708,538,12,683,872,7,209,65,18,936,40,511,297,860,732,748,449,549,861,110,106,21,58,201,665,26,604,140,1188,275,132,21,1079,20,648,120,1079,480,751,465,22,744,374,894,1628,3,367,945,19,373,185,28,711,251,0,1488,756,761,1424,757,69,15,649,280,116,102,122,71,79,18,934,596,463,779,216,1183,1354,18,1147,247,113,1379,686,524,45,1007,108,408,965,64,718,44,104,1510,143,161,46,34,871,329,992,559,503,1497,229,358,469,421,124,54,941,407,385,460,967,470,25,1552,109,992,6,331,345,225,23,48,712,483,1109,970,379,513,825,91,65,25,515,41,332,84,671,1318,505,772,21,1463,517,143,238,31,529,532,833,1671,110,766,44,257,36,458,1358,603,7,366,968,579,303,469,399,387,296,681,117,1089,931,1602,37,79,71,132,147,744,264,206,132,214,99,104,177,547,102,1550,771,1517,785,106,245,8,1602,298,15,533,451,339,351,448,241,199,128,1059,12,1,126,9,943,81,342,931,1007,499,1034,81,1483,98,782,1096,1050,952,185,1043,461,896,58,309,56,6,409,855,576,243,825,991,547,93,1721,604,580,355,11,854,0,222,496,671,935,148,1202,7,346,795,1409,1499,834,34,450,44,126,1203,18,779,1084,535,386,320,575,14,670,81,1036,1336,223,1054,1631,339,18,130,1002,131,36,998,462,559,1322,472,491,215,1402,1611,113,127,484,78,277,19,1104,71,220,40,586,555,489,611,267,507,632,47,786,59,1352,350,58,281,770,267,147,293,827,273,103,56,50,6,1224,1065,752,99,956,441,316,687,127,494,136,1336,1065,250,547,513,563,90,45,637,715,203,388,477,555,460,360,39,14,641,585,22,801,157,202,1152,505,227,572,599,659,361,1600,986,108,945,75,341,51,264,449,199,41,141,1258,870,1648,257,895,267,62,8,1208,166,142,618,375,1403"
