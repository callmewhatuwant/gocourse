package main

import (
	"fmt"
	"maps"
)

func main() {

	// var mapVariable map[key]valueType

	// mapVariable = make (map[keyType]valueType)

	//  using Map Literal

	// mapVaribale = map[keyType]valueType{
	// 	key1: value1,
	// 	key2: value2
	// }

	myMap := make(map[string]int)

	fmt.Println(myMap)

	myMap["key1"] = 9
	myMap["code"] = 18

	fmt.Println(myMap)
	fmt.Println(myMap["key1"])
	fmt.Println(myMap["key"])
	myMap["code"] = 21
	fmt.Println(myMap)

	delete(myMap, "key1")
	fmt.Println(myMap)

	myMap["key1"] = 9
	myMap["key2"] = 10
	myMap["key3"] = 11
	fmt.Println(myMap)

	// clear(myMap)
	// fmt.Println(myMap)

	// value, unkownvalue := myMap ["key1"]
	_, unkownvalue := myMap ["key1"]
	// fmt.Println(value)
	fmt.Println(unkownvalue)


	myMap2 := map[string]int{"a":1, "b":2}
	myMap3 := map[string]int{"a":1, "b":2}
	fmt.Println(myMap2)

	if maps.Equal(myMap2,myMap3){
		fmt.Println("map 1 and 2 are equal")
	}else{
		fmt.Println("maps are not equal")
	}

	for k,v := range myMap3{
		fmt.Println(k,v)
	}


}
