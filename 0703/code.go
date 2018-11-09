// 703. Kth Largest Element in a Stream

// Design a class to find the kth largest element in a stream. Note that it is the kth largest element in the sorted order, not the kth distinct element.
// Your KthLargest class will have a constructor which accepts an integer k and an integer array nums, which contains initial elements from the stream. For each call to the method KthLargest.add, return the element representing the kth largest element in the stream.

// Example:
// int k = 3;
// int[] arr = [4,5,8,2];
// KthLargest kthLargest = new KthLargest(3, arr);
// kthLargest.add(3);   // returns 4
// kthLargest.add(5);   // returns 5
// kthLargest.add(10);  // returns 5
// kthLargest.add(9);   // returns 8
// kthLargest.add(4);   // returns 8

// Note:
// You may assume that nums' length ≥ k-1 and k ≥ 1.

package leetcode

// KthLargest find the kth largest elements in a stream.
type KthLargest struct {
	heap []int
	k    int
}

// Constructor accepts an integer k and an integer array nums, which contains initial elements from the stream.
func Constructor(k int, nums []int) KthLargest {
	kl := KthLargest{
		heap: make([]int, 1, k+1),
		k:    k,
	}
	for _, num := range nums {
		kl.Add(num)
	}
	return kl
}

// Add return the element representing the kth largest element in the stream.
func (kl *KthLargest) Add(val int) int {
	if len(kl.heap)-1 < kl.k {
		kl.heap = append(kl.heap, val)
		kl.swim(len(kl.heap) - 1)
	} else if val > kl.heap[1] {
		kl.heap[1] = val
		kl.sink(1)
	}
	return kl.heap[1]
}

func (kl *KthLargest) swim(i int) {
	for i > 1 && kl.heap[i] < kl.heap[i/2] {
		kl.heap[i], kl.heap[i/2] = kl.heap[i/2], kl.heap[i]
		i /= 2
	}
}

func (kl *KthLargest) sink(i int) {
	for c := i + i; c <= kl.k; i, c = c, c+c {
		if c < kl.k && kl.heap[c+1] < kl.heap[c] {
			c++
		}
		if kl.heap[i] <= kl.heap[c] {
			break
		}
		kl.heap[i], kl.heap[c] = kl.heap[c], kl.heap[i]
	}
}
