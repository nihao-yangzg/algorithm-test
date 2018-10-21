##递归

def sum(arr):
    if len(arr) == 1:
        return arr[0]
    else:
        return arr[0] + sum(arr[1:])


def count(arr):
    if len(arr) < 1:
        return len(arr)
    else:
        return 1 + count(arr[1:])

def findMax(arr):
    if len(arr) == 1:
        return arr[0]
    else:
        leftMax = findMax(arr[1:])
        return (arr[0] if arr[0] > leftMax else leftMax)

##二分法查找元素
def binarySearch(arr, target):
    if len(arr) == 0:
        return False
    elif len(arr) == 1:
        return arr[0] == target
    else:
        mid = int(len(arr) / 2)
        if target < arr[mid]:
            return binarySearch(arr[0:mid], target)
        elif target > arr[mid]:
            return binarySearch(arr[mid:len(arr)], target)
        else:
            return True

##快排，效率较低，动态分配内存
def quickSort(arr):
    if len(arr) <= 1:
        return arr
    else:
        selected = arr[0];
        left = []
        right = []
        for item in arr[1:]:
            if item <= selected:
                left.append(item)
            else:
                right.append(item)
        return quickSort(left) + [selected] + quickSort(right)

#快排，内存大小不变
def quickSort0(arr, left,right):
    if right - left <= 0:
        return
    else:
        selected = arr[left]
        index = left
        L = left + 1
        R = right
        while L < R:
            while arr[L] <= selected:
                L = L + 1
            while arr[R] > selected:
                R = R - 1

            if L < R:
                tmp = arr[L]
                arr[L] = arr[R]
                arr[R] = tmp

        if arr[L] > selected:
            index = L - 1
        else:
            index = L       
        arr[left] = arr[index]
        arr[index] = selected
        
        if left < index - 1:
            quickSort0(arr, left, index-1)
        if index + 1 < right:
            quickSort0(arr,  index+1, right)
        
def quickSort1(arr):
    quickSort0(arr, 0, len(arr)-1)

if __name__ == '__main__':
    arr = [1,2,3,4,6,8]
    print(count(arr))
    print(sum(arr))
    print(findMax(arr))
    print(binarySearch(arr, 6))
    arr = [9,3,5,2,4,7]
    print(quickSort(arr))
    arr = [2,4,1,6,10,3]
    quickSort0(arr, 0, len(arr)-1)
    print(arr)
    arr = [2,2]
    arr = []
    import random
    for i in range(10):
        arr.append(random.random())
    
    quickSort1(arr)
    print(arr)
