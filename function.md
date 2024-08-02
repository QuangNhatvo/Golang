# Funtion và các khái niệm #
- Hàm là thành phần cơ bản của chương trình, các hàm trong ngôn ngữ Go có thể có tên hoặc ẩn danh (anonymous function).
```go
//Hàm được đặt tên
func add(a int, b int) int {
    return a+b
}

//Hàm ẩn danh
var Sum = func Sum(x int, y int) int {
    return x+y
}
```
- Một hàm trong Go có thể gồm nhiều tham số và nhiều giá trị trả về. Cả tham số và giá trị trả về trao đổi dữ liệu với hàm theo cách truyền giá trị (pass by value)
- Vê mặt cú pháp hàm cũng  hỗ trợ *số lượng tham số thay đổi*, *Biến số lượng tham số phải là tham số cuối cùng được truyền vào hàm và biến này phải là kiểu [slice](slice.md).
```go
package main
import "fmt"
//Nhiều tham số và nhiều giá trị trả về.
func Swap(a,b int) (int, int) {
    return b,a
}
//Biến số lượng tham số 'more' Tương ứng với kiểu []int, là một slice
func Sum(a int, more ...int) int{
    for _, v := range more {
        a += v
    }
    return a
}
func main() {
    sl := []int{3,4,5,8,9,10,12}
    sl[3], sl[4] = Swap(sl[3], sl[4])
    fmt.Println (sl)
    s := 10
    s = Sum(s,1,2,3,4)
    fmt.Println (s)
}
```
Kết quả mô phỏng.
![Kết quả](/Image/function.png)

Cả tham số và các giá trị trả về đều có thể được đặt tên.
```go
func Find(m map[int]int, key int) (value int, ok bool){
    value, ok = m[key]
}
```

Khi đối sô là một kiểu interface null, việc người gọi có phân giải (unpack) đối số đó hay không sẽ dẫn đến những kết quả khác nhau.
```go
func main() {
    var a = []interface{}{123, "abc"}

    fmt.Println(a...) //Tương đương với lời gọi trực tiếp Print(123, "abc") -> kết quả là 123 abc
    fmt.Println(a) //Tương đương với lời gọi Print([]interface{}{123, "abc"}) -> Kết quả là [123 abc] (chưa được unpack)
}
```

Kết quả mô phỏng.
![Kết quả](/Image/pack_unpack.png)

**Defer trong Function**
Lệnh ``` defer ``` trì hoãn việc thực thi hàm cho tới khi hàm bao ngoài nó return. Các đối số trong lời gọi defer được đánh giá ngay lập tức nhưng lời gọi không được thực thi cho tới khi hàm bao ngoài nó return.
```go
func main(){
    defer fmt.Println("Word)
    fmt.Println ("Hello")
}
```

Kết quả thực thi
[Kết quả](/Image/defer.png)

Mỗi lời gọi ``` defer ``` được push vào stack và thực thi theo thứ tự ngược lại khi hàm bao ngoài nó kết thúc.  
Ta thường sử dụng ``` defer ``` cho việc đóng hoặc giải phóng tài nguyên.
  - Đóng file giống như ``` try finally ```
  ```go
  f,err := os.Create("file")
  
  ```