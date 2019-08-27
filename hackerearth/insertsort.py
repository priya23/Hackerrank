#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the insertionSort1 function below.
def insertionSort1(n, arr):
    n=int(n)
    temp=arr[n-1]
    for i in reversed(range(n-1)):
        j=i-1
        if arr[i] > temp :
            arr[i+1] = arr[i]
            print(*arr,sep=' ')
            if j < 0:
                arr[i]=temp
                print(*arr,sep=' ')
                break
            j = j-1
        else:
            arr[i+1]=temp
            print(*arr,sep=' ')
            break
        
if __name__ == '__main__':
    n = int(input())

    arr = list(map(int, input().rstrip().split()))

    insertionSort1(n, arr)
