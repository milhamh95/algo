# class Solution:
#     def reverseint(num):
#         return 0

def reverseint(num):
    tmpNum = num
    if tmpNum < 0:
        tmpNum *= -1

    n = 0
    while(tmpNum > 0):
        remainder = tmpNum % 10
        tmpNum = tmpNum // 10
        n = (n * 10) + remainder

    if num < 0:
        n *= -1

    return n