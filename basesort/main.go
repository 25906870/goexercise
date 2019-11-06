package main

import(
	. "basesort/popsort"
	. "basesort/selectsort"
	. "basesort/insertsort"
	. "basesort/quicksort"
	"fmt"
)

func main(){

	nums := []int{3,2,1,6,7,8,9,5}

	Popsort(nums)

	stnums := []int{3,2,1,6,7,8,9,5}

	Selectsort(stnums)

	itnums := []int{3,2,1,6,7,8,9,5}

	Insertsort(itnums)

	qknums := []int{3,2,1,6,7,8,9,5}

	res := QuicksortRecusive(qknums)

	fmt.Println("Popsort result ",nums)
	fmt.Println("Selectsort result ",stnums)
	fmt.Println("Insertsort result ",itnums)
	fmt.Println("Quicksort result ",res)
}