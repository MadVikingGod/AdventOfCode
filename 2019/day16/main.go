package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"math"
)

func main() {
	fmt.Println(len(input))
	//in := []float64{1,2,3,4,5,6,7,8}
	//in := []float64{8,0,8,7,1,2,2,4,5,8,5,9,1,4,5,4,6,6,1,9,0,8,3,2,1,8,6,4,5,5,9,5} //24176176
	//input = []float64{1,9,6,1,7,8,0,4,2,0,7,2,0,2,2,0,9,1,4,4,9,1,6,0,4,4,1,8,9,9,1,7} //73745418

	x := mat.NewDense(len(input), 1, input)
	tmp := mat.DenseCopyOf(x)
	A := createMat(len(input))


	for i:=0;i<100;i++ {
		tmp.Mul(A,x)
		x.Apply(norm, tmp)
	}
	for i:=0;i<8;i++ {
		fmt.Print(x.At(i,0))
	}
	fmt.Println()
   	// part 2
	inputInt = []int{0,3,0,3,6,7,3,2,5,7,7,2,1,2,9,4,4,0,6,3,4,9,1,5,6,5,4,7,4,6,6,4,} //84462026
   	offset := 0
	for i:=0; i<7;i++ {
		offset += inputInt[i] * int(math.Pow10( 6-i))
	}
	n := (len(inputInt)*10000-offset)/len(inputInt)+1

	fmt.Println("offset", offset, "N", n)

	offset = offset - (10000-n)*len(inputInt)
	fmt.Println("New offset", offset)
	//6500000-5979191 = 520809  We have to move this far in our input.  But no digit depends on the digit before
	//520809/650 = 801.244615385  //We don't need the full 10,000 copies, just the last 802
	//(10000-802)*650=5978700  //If we start at copy 9198 (10000-802) we are at this offset
	//5979191-5978700=491  //we then just need to read the 8 digits from this position in our matrix.

	in := duplicate(inputInt, 802)
	in = in[offset-1:]

	temp := make([]int, len(in))
	for count:=0; count<100; count++ {
		sum := 0
		for i := len(in) - 1; i >= 0; i-- {
			sum += in[i]
			temp[i] = sum
		}
		for i, v := range temp {
			in[i] = v % 10
		}
		fmt.Println(in[0:8])
	}
	fmt.Println(in[0:8])

}
func sum(in []int) int {
	s := 0
	for _,i:=range in {
		s+=i
	}
	return s
}

func duplicate(a []int, n int) []int {
	x := make([]int, len(a)*n)
	l := len(a)
	for i:=0; i<n; i++ {
		copy(x[i*l:i*l+l], a)
	}
	return x
}


func createMat(len int) *mat.Dense {
	starter := []float64{0,1,0,-1}
	data := make([]float64, len*len)
	for i:=0; i<len; i++ {
		for j:=0; j<len; j++{
			data[i*len+j] = starter[(j+1)/(i+1)%4]
		}
	}
	return mat.NewDense(len,len, data)
}

func norm(_,_ int,v float64) float64 {
	return math.Mod(math.Abs(v),10.0)
}