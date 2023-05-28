# goext
Go 1.18+ 基础类型扩展库：在go内置map、slice、channal等基础类型上基于泛型做一层扩展。

## 使用
- 安装
    ```go
    go get github.com/hellolib/goext
    ```
- 引入并调用
    ```go
    import github.com/hellolib/goext
    res := goext.Filter[int]([]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
       func(index int, item int) bool {
          return item >= 5
       },
    )
    fmt.Println(res) // [5 6 7 8 9]
    ```

## feature

### slice
- **Filter**: 切片过滤器
  ```go
    res := goext.Filter[int]([]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
       func(index int, item int) bool {
          return item >= 5
       },
    )
    fmt.Println(res) // [5 6 7 8 9]
  ```
- **Map**: 切片元素加工(返回切片元素类型可以不同)
  ```go
    res := goext.Map[int, int]([]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
       func(index int, item int) int {
          return item * 2
       },
    )
    fmt.Println(res) // [2 4 6 8 10 12 14 16 18]
  ```
- **Reduce**: 对切片的所有元素聚合

  ```go
    res := goext.Reduce[int, int]([]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
       func(index int, item int, initial int) int {
          return initial + item
       }, 0,
    )
    fmt.Println(res) // 45
  ```