# CLRS

## 统计脚本
* 统计leetcode刷题的数目
```shell
ls yan/leetcode/src/ --ignore="*_test.go"|wc -l
```
* 统计已看clrs的章节
```shell
ls -l yan/exercise/src/| grep "^d"|awk '{print $9}'|sed 's/\///'|awk '{print substr($1,8)}'|sort -r|head -1

```