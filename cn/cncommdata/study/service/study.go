package service

/**
 * 1。GO语言在命名的时候，一般会在单词后面添加er，如有写操作的接口叫writer，有字符串功能叫stringer，又关闭功能的 接口叫closer
 * 2.方法名：当方法名首字母大写时，且这个接口类型名首字母也是大写是，这个方法可以被接口所在的包之外的代码访问
 * 3. 参数列表、返回值列表：参数列表和返回值列表中的参数变量名可以被忽略
 *
 */
type writer interface {
	Writer([]byte) error
}

/**
 * 开发中常见的接口及写法
 * 这个接口可以调用Writer()方法写入一个字节数组，返回值告诉写入字节数，和可能发生的错误
 */
type Writer interface {
	Write(p []byte) (n int, err error)
}

/**
 * stringer 接口在Go语言中的使用频率非常高，功能类似于Java 语言里的ToString方法
 * Go语言的每个接口中的方法数不会很多。Go语言希望通过一个接口精准描述它自己的功能，而通过多个接口的嵌入和组合的方式将简单的接口扩展为复杂的接口。本章后面的小结中会介绍如何使用组合来扩充接口
 */
type Stringer interface {
	String() string
}
