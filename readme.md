## 统计脚本
* 统计leetcode刷题的数目
```shell
ls leetcode/src|egrep -v '.*_test.go'|wc -l
```
* 统计已看clrs的章节
```shell
ls -l exercise/src|grep "^d"|awk '{print $9}'|awk '{print substr($1,8)}'|sort -r |head -1
```

## chapter1 The role of Algorithms in computing
* what are algo

```shell
a tool for solving a well-specified computational problem
```

* why study algo
```
with a good backgroud in algorithms, you can do much, much more.
```



### concepts
A **data structures** is a way to store and organize data in order to facilitate access and modifications


