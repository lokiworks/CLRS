package main

/**
Select(int *A,int p,int r,int i)算法思想：
1.将输入数组的n个元素划分为n/5组，每组5个元素，且至多只有一组由剩下的n mod 5个元素组成。
2.首先对每组元素进行插入排序，然后确定每组有序元素的中位数。
3.递归调用Select(),找出上述所有组中位数的中位数，若为偶数个，可以约定用哪一个。
4.调用Partition(int *A,int p,int r,int x)，x为上一步中的中位数，k低区元素数目多1，因此x为第k小的元素。
5.若i == k ,返回 x ,如果i < k,则在低区递归调用Select()，若 i > k ，则在高区递归查找第i -k小的元素。
————————————————
版权声明：本文为CSDN博主「勿在浮砂筑高台」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/luoshixian099/article/details/45286303
*/
func select1(a []int, p, r, i int) int {
	for (r-p)%5 != 0 {
		for j := p + 1; j < r+1; j++ {
			if a[p] > a[j] {
				a[p], a[j] = a[j], a[p]
			}
		}
		if i == 0 {
			return a[p]
		}
		p++
		i--
	}
	g := (r - p) / 5
	for j := p; j < p+g; j++ {

		a2 := []int{a[j], a[j+g], a[j+2*g], a[j+3*g], a[j+4*g]}
		insertSort(a2)
		a[j], a[j+g], a[j+2*g], a[j+3*g], a[j+4*g] = a2[0], a2[1], a2[2], a2[3], a2[4]

	}
	x := select1(a, p+2*g, p+3*g-1, g/2)
	q := partition2(a, p, r, x)
	k := q - p + 1
	if i == k {
		return a[q]
	} else if i < k {
		return select1(a, p, q-1, i)
	}
	return select1(a, q+1, r, i-k)
}

func partition2(a []int, p, r, x int) int {
	i := p - 1               // the highest index into the low side
	for j := p; j < r; j++ { // process each element other than the pivot
		if a[j] <= x { // does this element belong on the low side?
			i++ //index of a new slot in the low side
			a[i], a[j] = a[j], a[i]

		}

	}
	a[i+1], a[r] = a[r], a[i+1]
	return i + 1
}

func insertSort(a []int) []int {
	len := len(a)
	for i := 1; i < len; i++ {
		k := a[i]
		// insert k into the sorted subarray A[0:i-1]
		j := i - 1
		for j >= 0 && a[j] > k {
			a[j+1] = a[j]
			j--
		}
		a[j+1] = k
	}

	return a
}
