# **Các khái niệm liên quan đến Slice** #
### Slice ###
![Cấu tạo của slice](/Image/Slice.png)
- Cấu trúc của slice thì tương tự như string tuy nhiên các giới hạn của string được lược bỏ. 
- Cấu trúc của Slice.
```go
type SliceHeader struct {
    Data intptr
    Len int
    Cap int
}
```
-> Slice được xem là **fat pointer**. Vì con trỏ bình thường chỉ chứa địa chỉ, nhưng Slice lại chứa thêm data về Len và Cap.
- Mỗi Slice trỏ đến 1 mảng cơ sở (Underlying Array).
---
### Underlying Array (Mảng cơ sở) ###
- Là mảng con trỏ mà slice trỏ tới. 
- Nhiều slice có thể dùng chung một Underlying Array. 

---
### Các cách định nghĩa Slice ###
```go 
a = []int // nil slice (Slice không trỏ đến bất kỳ Underlying Array nào).
b = []int{} //empty slice (Slice trỏ đến Underlying Array rỗng).
c = []int {1,2,3} //Khởi tạo slice có 3 phần tử, cả len và cap đều bằng 3.
d = c[:2] // Có 2 phần tử trong Slice, len bằng 2 và cap bằng 3.
e = make ([]int, 2, 3) // Tạo slice có 2 phần tử 
```
- **'d = c[:2]'** cú pháp tạo một slice từ một slice cho trước -> d tạo ra một slice mới trỏ đến mảng cơ sở Underlying Array của c -> 2 thằng này dùng chung 1 Underlying Array -> Khi thay đổi giá trị trong d thì c cũng thay đổi.
- 1 ví dụ rõ ràng hơn về việc 2 Slice dùng chung một mảng cơ sở.
```go
x := []int{2,3,5,7,11}
y := x[1:3]
```
![Hình ảnh mô phỏng cho ví dụ trên](/Image/SliceLayout.png)
- Khi thực hiện thay đổi ở Slice y, Slice x cũng thay đổi theo.
  - Chương trình cho thấy, Slice a thay đổi khi thực hiện thay đổi Slice b -> a và b dùng chung Underlying Array.
  ```go
  package main
  import "fmt"

  func main() {
	  a:= []int{2,3,5,7,11}
	  fmt.Println("Slice 'a' before:")
	  fmt.Println(a)
	  b:= a[:3]
	  b[0] = 1
	  fmt.Println("Slice 'a' after:")
	  fmt.Println(a)
	  fmt.Println(b)
  }   
  ```
  - Kết quả hiện thực chương trình.
  ![Kết quả](/Image/example.png)

---
### Các tác vụ cơ bản trong slice. ###
- Duyệt qua Slice.  
Sử dụng ``` for ... range ``` để duyệt qua slice.
```go
for i := range a { 
  fmt.Println("a[%d]: %d", i, a[i])
}
for i, v := range b{
  fmt.Println("b[%d]: %d", i, v)
}

for i:=0; i < len(c);i++ { 
  fmt.Println("c[%d]: %d",i, c[i])
}
```
![Kết quả duyệt Slice](/Image/DuyetSlice.png)

- Thêm phần tử vào Slice.  
Hàm ``` append ``` có thể thêm phần tử thứ N vào cuối Slice.
```go
var a []int
a = append(a,1) //Nối thêm 1 phần tử 
a = append(a,2,3,4) //Nối thêm 3 phần tử 2, 3, 4
a = append(a,[]int{5,6,7}) //Nối thêm 3 phần tử 5,6,7 bằng cách truyền tham số.
```
Trong trường hợp Slice ko đủ kích thước. Hàm ``` append ``` sẽ tiến hành cấp phát lại vùng nhớ có kích thước.
  - Nếu kích thước cũ (cap) < 1024: Cấp phát x2 vùng nhớ cũ.
  - Nếu kích thước cũ (cap) >= 1024: Cấp phát x1.5 vùng nhớ cũ.
Sau đó dữ liệu cũ sẽ được sao chép sang.
**Lưu ý:** ``` append ``` sẽ không cấp phát lại vùng nhớ khi cap chưa đạt giá trị tối đa

![Cấu trúc Slice](/Image/structure.png)
Ví dụ cho thấy giá trị cap tăng 2 lần khi thực thi hàm append vượt quá kích thước ban đầu.
```go
package main
import "fmt"

func myAppend(sl []int, val int) []int{
  sl = append (sl, val)
  printSlice(sl)
  return sl
}

func printSlice(sl []int){
  fmt.Printf("Slice: %v || ", sl)
  fmt.Printf("Len: %v, Cap %v", len(sl), cap(sl))
  fmt.Println
}

func main(){
  var mySlice = make([]int, 1)
  printSlice(mySlice)
  for i:=1; i<= 5; i++ {
    mySlice = myAppend(mySlice,i)
  }
}
```
Kết quả mô phỏng chương trình.
![Kết quả](/Image/MoPhong.png)

Bên cạnh thêm phần tử vào cuối Slice,cũng có thể thêm phần tử vào đầu Slice (Nối mảng chứ ko thể thêm từng phần tử như khi thêm vào cuối mảng).
```go
var a = []int{1,2,3}
// thêm phần tử 0 vào đầu slice a
a = append([]int{0}, a...)
// thêm các phần tử -3, -2, -1 vào đầu slice a
a = append([]int{-3,-2,-1}, a...)
```
Kết quả mô phỏng. 
![Kết quả](/Image/KetQua.png)

Do hàm ``` append ``` sẽ trả về một Slice mới, có thể kết hợp nhiều hàm ``` append ``` để chèn 1 vài phần tử vào giữa Slice.
```go
//Thêm 1 phần tử x vào vị trí thứ i trong Slice
sl = append(sl[:i],append([]int{x},sl[i:]...)...)
//Chèn 1 slice con vào vị trí thứ i
sl = append(sl[:i],append([]int{3,4,5},sl[i:]...)...)
```
Kết quả mô phỏng.
![Kết quả](/Image/middle.png)

- Xóa phần tử trong Slice.  
Có 3 trường hợp xóa phần tử.  
  - Ở đầu
  - Ở giữa
  - Ở cuối  
Trong đó xóa phần tử ở cuối là nhanh nhất.
```go
a := []int{1,2,3}
//Xóa 1 phần tử ở cuối.
a = a[:len(a)-1]
//Xóa N phần tử ở cuối.
a = a[:len(a)-N]
```
Kết quả mô phỏng.
![Kết quả](/Image/delfinal.png)

Xóa phần tử ở đầu thì thực chất là đưa con trỏ dữ liệu về sau.
```go
a := []int{1,2,3,4,5}
//Xóa 1 phần tử ở đầu.
a = a[1:]
//Xóa N phần tử ở đầu.
a = a[N:]
```
Kết quả mô phỏng.
![Kết quả](/Image/delfirst.png)

Xóa phần tử ở giữa thì cần dịch chuyển các phần tử từ phía sau lên trước.
```go
a := []int{1,2,3,4,5}
//Xóa phần tử ở vị trí thứ i
a = append(a[:i],a[i+1:]...)
//Xóa N phần tử ở vị trí thứ i
a = append(a[:i],a[i+N:]...)
```

Kết quả mô phỏng.
![Kết quả](/Image/delmid.png)

- Kỹ thuật quản lý vùng nhớ trong Slice.  
Hàm ``` TrimSpace ```được sử dụng để xóa đi các khoảng trắng. Hiện thực hàm trên với độ phức tạp O(n) để đạt được sự hiệu quả và đơn giản.
```go
func TrimSpace(s []byte) []byte {
    b := s[:0]
    // duyệt qua slice s để tìm phần tử thỏa điều kiện
    for _, x := range s {
        // kiểm tra điều kiện
        if x != ' ' {
        // tạo ra slice mới từ slice ban đầu thêm vào phần tử x
            b = append(b, x)
        }
    }
    return b
}
```
**Note:**  
Ký hiệu _ được gọi là "blank identifier" và nó được sử dụng để bỏ qua một giá trị mà không cần sử dụng. Khi sử dụng trong một vòng lặp for range, _ thường được sử dụng để bỏ qua chỉ số (index) của phần tử trong slice, map, hoặc kênh (channel)
---