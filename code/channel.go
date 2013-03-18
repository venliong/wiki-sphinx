package main  

import "fmt"

func Count(ch chan int) {
    fmt.Println("Counting")
    ch <- 1 
}

func main() {
    chs := make([]chan int, 10)
    for i := 0; i < 10; i++ {
        chs[i] = make(chan int)
        go Count(chs[i])
    }

    for _, ch := range(chs) {
       num :=  <-ch
       fmt.Println(num)
    } 
}

/*
在这个例子中，我们定义了一个包含10个channel的数组（名为chs），
并把数组中的每个channel分配给10个不同的goroutine。
在每个goroutine的Add()函数完成后，
我们通过ch <- 1语句向对应的channel中写入一个数据。
在这个channel被读取前，这个操作是阻塞的。
在所有的goroutine启动完成后，
我们通过<-ch语句从10个channel中依次读取数据。
在对应的channel写入数据前，这个操作也是阻塞的。
这样，我们就用channel实现了类似锁的功能，
进而保证了所有goroutine完成后主函数才返回。是不是比共享内存的方式更简单、优雅呢?
*/