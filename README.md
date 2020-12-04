# mk2html
markdown to html

没有依赖第三方工具，完全用golang实现。解析markdown，构建语法树，生成html结构。

# 现支持的markdown标记
```
# head
## head
* lo
1. lu
*em*
**stronge**
~~del~~
`code`

```

## 未来支持
```表格
| head1 | head2 | head3 |
| --    |  --   |  --   |
| a     | b     | c     |
```
