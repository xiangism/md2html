# mk2html
markdown to html lib

没有依赖第三方工具，完全用golang实现。解析markdown，构建语法树，生成html结构。
# 

# 现支持的markdown标记
```
表格
| head1 | head2 | head3 |
| --    |  --   |  --   |
| a     | b     | c     |

* ul
1. ol

# head
## head
###### head
code area <pre><code>
```

## 未来支持
```
*em*
**strong**
~~del~~
`code`

```

#h

*
1. 123
2. 45

123*em*cb***_*** *1* **2** 

**stronge**

~~del~~

123`code`abc


# level
TODO: 解析出head层级 锚点
```
<a id="md_0"></a>
<h1></h1>
在h1-h6前面添加 <a></a>
然后用 # 即可直达
file:///D:/work/Go_app/GoPath/src/github.com/xiangism/md2html/out.html#md_0


h0
    h0_0
h1
    h1_0
h2
h3
```
