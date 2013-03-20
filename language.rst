.. _language:

Programming language
======================


Golang Notes
------------------

Golang 语言学习笔记,参考 

* `golang book <http://www.golang-book.com>`_
* `learn go by example <https://gobyexample.com/>`_
* `go 网络编程 <http://jan.newmarch.name/go/zh/>`_


指针
^^^^^^^^^^^^^^^^^^

对于下面的代码, 在main调用zero() 函数, main函数里的x输出是5,如果我们想要改变main函数里的x值,
需要用到指针

.. code-block:: go

		func zero(x int) {
		    x = 0
		}
		func main() {
		    x := 5
		    zero(x)
		    fmt.Println(x) // x is still 5
		}

	
对于zero()函数参数类型为指针类型

.. code-block:: go

		func zero(xPtr *int) {
		    xPtr = 0
		}
		func main() {
		    x := 5
		    zero(&x)
		    fmt.Println(x) // x is 0
		}

*和&操作符	
""""""""""""""""

在zero函数里 xPtr 表示int型的指针, * xPtr = 0, 可以访问指针指向的存储单元的值, 即可以使用*访问指针指向的内存单元,  不能使用 xPtr = 0, 会得到编译器错误 , xPtr 是* int 类型,而不是int类型.

& 操作符用于取得变量的内存地址, &x 会返回一个int型的指针


new 
""""""""""""

不仅可以使用&操作符得到变量的内存地址, 也可以用内建的new函数

.. code-block:: go

		func one(xPtr *int) {
		    *xPtr = 1
		}
		func main() {
		    xPtr := new(int)
		    one(xPtr)
		    fmt.Println(*xPtr) // x is 1
		}

new 函数接受参数为某一类型, 分配内存, 返回指向它的指针



接口
^^^^^^^^^^^^^^^^^^^

简单的说，interface是一组method的组合，我们通过interface来定义对象的一组行为。
在golang中,一个类只需要实现了这个类所要求的所有函数,我们就说这个类实现了这个接口. 


接口例子:

.. literalinclude:: code/interfaces.go
	:language: go


网络编程 
^^^^^^^^^^^^^^^^^^^^^^^

一个简单的例子，展示个客户端连接到一个网页(HTTP)服务器。 

.. literalinclude:: code/go-net/tcpaddr.go
	:language: go


一个时间服务器的例子

有关调用::

	func ListenTCP(net string, laddr *TCPAddr) (l *TCPListener, err os.Error)
	func (l *TCPListener) Accept() (c Conn, err os.Error)

net参数可以设置为字符串"tcp", "tcp4"或者"tcp6"中的一个。如果你想监听所有网络接口，IP地址应设置为0，或如果你只是想监听一个简单网络接口，IP地址可以设置为该网络的地址。如果端口设置为0，O/S会为你选择一个端口。否则，你可以选择你自己的。需要注意的是，在Unix系统上，除非你是监控系统，否则不能监听低于1024的端口，小于128的端口是由IETF标准化。该示例程序选择端口1200没有特别的原因。TCP地址如下":1200" - 所有网络接口, 端口1200。




.. literalinclude:: code/go-net/time_server.go
	:language: go



Go 语言编程笔记 
^^^^^^^^^^^^^^^^^^

P79::

	go语言和c语言一样,类型都是基于值传递的,要想修改变量的值,要传递指针 

goroutine::

	go go语言中最重要的关键字,只要在函数前加上go关键字,这次调用会在一个新的goroutine中并发执行. 当被
	调用的函数返回时, 这个 goroutine 也结束了 ** 如果这个函数有返回值,那么这个返回值会被丢弃**


并发 通信::

		goroutine是Go语言中的轻量级线程实现，由Go运行时（runtime）管理

		并发单元间的通信是最大问题

		工程上最常用的两种并发通信模型, 共享数据和消息

		go 语言使用**消息机制**

		消息机制认为每个并发单元是自包含的、独立的个体，并且都有自己的变量，但在不同并发单元间这些变量不共享。每个并发单元的输入和输出只有一种，那就是消息。

		“不要通过共享内存来通信，而应该通过通信来共享内存。”


channel::
	
	channel是Go语言在语言级别提供的goroutine间的通信方式。我们可以使用channel在两个或多个goroutine之间传递消息。	

	channel是类型相关的。也就是说，一个channel只能传递一种类型的值

	一般channel的声明形式为： var chanName chan ElementType

	我们声明一个map，元素是bool型的channel: var m map[string] chan bool

channel example 

.. literalinclude:: code/channel.go
	:language: go

