select case 语句----->当多个case语句均满足时，随机执行一个便退出
switch case 语句----->每条case语句默认加上break，即执行完匹配的case语句后，便会退出switch。可以在每条case语句后面加上fallthrough关键字，强行执行后面的case语句。
defer  语句的嵌套入桟
panic 语句的执行，是一个内建函数,可以中断原有的控制流程,进入一个令人恐慌的流程中。当函数 F 调用
panic ,函数 F 的执行被中断,但是 F 中的延迟函数会正常执行,然后 F 返回到调用它的地方
map并发安全？ 1.8
闭包---?传值。
goto 语句跳转，只能在goto的函数内部进行跳转。
iota --->从0开始，每遇到const关键字便从0开始

struct 字段嵌套，当两个struct均有相同字段的时候，要遵循最外层的优先访问原则。
method 注意点：
原型语法 func (r ReceiverType) funcName(parameters) (results)
1.虽然 method 的名字一模一样,但是如果接收者不一样,那么 method 就不一样
2.method 里面可以访问接收者的字段
3.调用 method 通过 . 访问,就像 struct 里面访问字段一样
指针作为Receiver，变量作为Receiver，如果指针作为接收者，可以修该其值，而变量作为接收者，只是copy。
方法重写，及相应的调用逻辑

interface变量存储的类型，
1.可以直接判断是否是该类型的变量: value, ok = element.(T),
这里 value 就是变量的值,ok 是一个 bool 类型,element 是 interface 变量,T 是断言的类
型。如果 element 里面确实存储了 T 类型的数值,那么 ok 返回 true,否则返回 false。
2.这里有一点需要强调的是: valuetype:=element.(type) 语法不能在 switch 外的任何逻辑里面使用,如
果你要在 switch 外面判断一个类型就使用 comma-ok 。

接口赋值：
1.将对象实例赋值给接口（一个结构实现了接口的所有方法，可实例化一个对象，直接赋值给接口）
type A interface{} 方法多
type B interface{} 方法少
type C struct{} 实现了A，B的接口
c=new（C）
var a A 
a=c

2.将一个接口赋值给另一个接口（方法多的接口，可直接赋值给方法少的接口）
type A interface{}方法多
type B interface{}方法少
var a  A
var b  B
b=a

for range之坑：
 伪代码
    slice := []int{0, 1, 2, 3}
    myMap := make(map[int]*int)
	fmt.Println("原始数据",slice)
    for index, value := range slice {
		num:=value//修改处
	   // myMap[index] = &value //修改处
		myMap[index]=&num;
    }
note:
映射的值都相同且都是3。其实可以猜测映射的值都是同一个地址，遍历到切片的最后一个元素3时，将3写入了该地址，所以导致映射所有值都相同。其实真实原因也是如此，因为for range创建了每个元素的副本，而不是直接返回每个元素的引用，如果使用该值变量的地址作为指向每个元素的指针，就会导致错误，在迭代时，返回的变量是一个迭代过程中根据切片依次赋值的新变量，所以值的地址总是相同的，导致结果不如预期。


反射：


掌握垃圾回收机制
高并发工具，mysql，多表联及查询。
项目项目
grpc的应用。
