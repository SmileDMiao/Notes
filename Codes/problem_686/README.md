### 686. Repeated String Match
> question

给两个字符串 A B，重复A直到B是A的字串, 求A重复的最小次数

> example

Input: a = "abcd", b = "cdabcdab", Output: 3

> 思路

要使得字符串B成为字符串A的子串，必须使字符串A的长度大于或等于字符串B的长度
最多重复次数为: len(b)/len(a) + 1
在字符串A的长度小于字符串B的长度时，重复叠加字符串A,并记录叠加次数；若字符串B仍不能成为叠加后的字符串A的子串，再一次叠加字符串A,再次判断，若仍不能，则B并不是其子串。
