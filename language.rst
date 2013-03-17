.. _language:

Programming language
======================


Golang Notes
------------------

Golang 语言学习笔记,参考 

* `golang book <http://www.golang-book.com>`_
* `learn go by example <https://gobyexample.com/>`_



指针
^^^^^^^^^^^^^^^^^^

对于下面的代码, 在main调用zero() 函数, main函数里的x输出是5,如果我们想要改变main函数里的x值,需要用到指针

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
