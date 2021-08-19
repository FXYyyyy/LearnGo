package main

import (
	"bytes"
	"fmt"
	"io"
	"sync"
	"time"
)

/**
*  DataBucket
*  @Description: 数据bucket
*/
type DataBucket struct {
	buffer *bytes.Buffer	//缓冲区
	mutex *sync.RWMutex		//互斥锁
	cond  *sync.Cond		//条件变量
}
func NewDataBucket() *DataBucket {
	buf := make([]byte, 0)
	db := &DataBucket{
		buffer: bytes.NewBuffer(buf),
		mutex: new(sync.RWMutex),
	}
	db.cond = sync.NewCond(db.mutex.RLocker())	//通知的变量为一个等待的读锁

	return db
}

/**
* read
* @Description: 读取器
* @receiver db
* @param i
*/
func (db *DataBucket) read(i int)  {
	db.mutex.RLock()	//打开读锁
	defer db.mutex.RUnlock()	//结束后释放读锁

	var data []byte
	var d byte
	var err error
	for{
		//每次读取一个字符
		if d,err = db.buffer.ReadByte(); err != nil{
			if err == io.EOF {	//缓冲区数据为空时
				if string(data) != "" {
					fmt.Printf("read-%d: %s\n", i, data)
				}
				db.cond.Wait()	//缓冲区为空，通过wait方法等待通知，进入阻塞状态
				//data = data[:0]	//将data数据清空
				continue
			}
		}
		data = append(data, d)	//将读到的字符追加到data
	}
}

/**
* Put
* @Description: 写入器
* @receiver db
* @param d
* @return int
* @return error
*/
func (db *DataBucket) Put(d []byte) (int, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	//写入一个数据块
	n,err := db.buffer.Write(d)
	//db.cond.Signal()	//写入数据后通过 Signal 通知处于阻塞状态的读取器(RLock)
	db.cond.Broadcast()	//通知多个阻塞等待的进程
	return n, err
}

func main() {
	//testWithSingle()
	testWithMulti()
}

/**
* testWithSingle
* @Description: 仅开启一个读、一个写进行测试
*/
func testWithSingle()  {
	db := NewDataBucket()

	go db.read(1)	// 开启读取器协程
	go func(i int) {
		d := fmt.Sprintf("data-%d: zhizhi", i)
		db.Put([]byte(d))
	}(1)	// 开启写入器协程

	time.Sleep(100*time.Millisecond)
}

func testWithMulti()  {
	db := NewDataBucket()

	for i:=0; i<3; i++{
		go db.read(i)
	}

	for i:=0; i<5; i++ {
		go func(i int) {
			data := fmt.Sprintf("data-%d", i)
			go db.Put([]byte(data))
		}(i)

		time.Sleep(500*time.Millisecond)
	}
}
