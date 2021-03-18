/**
 * @Author: "lyupeng"
 * @Description: 二分查找
 * @File:  Binnay_test
 * @Version: 1.0.0
 * @Date: 2021/3/18 9:46
 */

package test

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	nums:=[]int{0,2,3,4,5,16,77,88,99,100}
	targetNum:=99
	r:=DoBinarySearch(nums,targetNum)
	if r>0 {
		fmt.Println("查询所得的下标:",r)
	}else {
		fmt.Println("提示:","数据不存在")
	}
}
func DoBinarySearch(nums []int,targetNum int)int  {
	low:=0
	high:=len(nums)-1
	for low<=high{
		midIndex:=(high-low)/2+low
		midValue:=nums[midIndex]
		if targetNum==midValue {
			return  midIndex
		}else if targetNum<midValue {
			high=high-1
		}else if targetNum>midValue{
           low=low+1
		}
	}
	return -1
}