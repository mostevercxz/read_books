## proto2 和 proto3区别,应该用哪个
We currently recommend trying proto3 only:

1. If you want to try using protocol buffers in one of our newly-supported languages（go, ruby, objective-c).
2. If you want to try our new open-source RPC implementation gRPC (currently also in alpha release) – we recommend using proto3 for all new gRPC servers and clients as it avoids compatibility issues.

Note that the two language version APIs are not completely compatible. To avoid inconvenience to existing users, we will continue to support the previous language version in new protocol buffers releases.

## 定义 .proto 文件
    package tutorial;
    
    message Person {
      required string name = 1;
      required int32 id = 2;
      optional string email = 3;
    
      enum PhoneType {
    	MOBILE = 0;
    	HOME = 1;
    	WORK = 2;
      }
    
      message PhoneNumber {
    	required string number = 1;
    	optional PhoneType type = 2 [default = HOME];
      }
    
      repeated PhoneNumber phone = 4;
    }
    
    message AddressBook {
      repeated Person person = 1;
    }

1. package tutorial 相当于 c++ 的 namespace tutorial ，防止不同工程间的命名冲突(c++中使用的时候用 tutorial::Person)；
2. 一条 message 就是一系列有类型的字段的集合。 protobuf中类型有下面几种：
	* 简单数据类型 : bool, int32, float, double, string
	* 其他的 message : 比如 message Person 中包含 PhoneNumber 类型
	* enum 类型 : 比如 PhoneType
3. 代码中的 =1,=2 是对应字段在protobuf最终的二进制编码中使用的唯一 id.(1-15 比16-∞ 所需要的字节少1个，所以尽量将1-15留给最常使用或者 repeated 类型的字段) 
4. 每个字段必须用下列修饰符来修饰：
	* **required**. 必须提供值给该字段,否则该 message 会被认为是未初始化的。 解析未初始化的message 总是会 return false. 除此之外， required 字段和 optional 字段的表现一样。
	* **optional**. 该字段可以有值或没值。如果没设置 optional 字段的值，将会使用默认值 (用 [default = HOME]声明的) 或者系统默认值(string="", int32,float,double=0,bool=false). 对于嵌套的message,默认值为任何一个字段都没被设置的 message. 调用 get 函数去取 值还未被设置的 optional(或required) 字段，将会得到该字段的默认值。
	* **repeated**. 该字段可以重复0次或多次。 该字段的顺序会在 protocol buffer 中被保留。


---
题外话：为什么是 1-15 而不是 1-255 呢？
Varints 是用1字节或者更多字节来序列化 integers 的一种方式。 在 varint 中的每个字节(除了最后一个byte)的最高位都会被置为1，用来标记接下来还有 1个byte 需要读取。 每个字节的低7位用来存储该 integer 的二补数(two's-complement)表示形式，每7位为1组，最小的组放在最前面(类似little endian,低位在前高位在后)。
 
对于 300 这个数,二补数形式是 `1 0010 1100`;
每7位分组的话，可以分2组: 10,0101100;
最小的组放在最前面,并且将除最后一个byte以外的其他 byte 的高位置为1,得到 varint 300的编码 : 10101100 00000010 

一条 protobuf message 是由一系列 key-value 对组成的集合。 protobuf message 的二进制形式使用字段的 number(required 字段名 = number) 作为key. 每个 key 都是值为 (field_number << 3) | wire_type 的 varint,也就是说key的最后3位存储 wire_type. 

key 最小为 1 byte, 该 byte 的第一位置0表示后面没有 byte, 最后3位留给 wire_type,只剩4位，所以能表示的id范围是 0-15.

为什么不能用0呢？从 protobuf 的设计层面上说，得看源代码。。
从应用层面上说，假设使用了0，在生成 .cc 和 .h文件的时候， protoc 会报错:

    test.proto:2:28: Field numbers must be positive integers.

参考链接 :
 
1. [google protobuf encoding](https://developers.google.com/protocol-buffers/docs/encoding)
2. [can protobuf fileds start with zero](http://stackoverflow.com/questions/26866911/can-proto-files-fields-start-at-zero)

---
## 由 .proto 生成 c++ 代码
    protoc -I=$SRC_DIR --cpp_out=$DST_DIR $SRC_DIR/addressbook.proto

    SRC_DIR 为 addressbook.proto 所在的目录
    DST_DIR 为 目标目录，即生成的 addressbook.pb.h 和 addressbook.pb.cc 存放的目录

## API 简介
    // name
    inline bool has_name() const;
    inline void clear_name();
    inline const ::std::string& name() const;
    inline void set_name(const ::std::string& value);
    inline void set_name(const char* value);
    inline ::std::string* mutable_name();
    
    // id
    inline bool has_id() const;
    inline void clear_id();
    inline int32_t id() const;
    inline void set_id(int32_t value);
    
    // email
    inline bool has_email() const;
    inline void clear_email();
    inline const ::std::string& email() const;
    inline void set_email(const ::std::string& value);
    inline void set_email(const char* value);
    inline ::std::string* mutable_email();
    
    // phone
    inline int phone_size() const;
    inline void clear_phone();
    inline const ::google::protobuf::RepeatedPtrField< ::tutorial::Person_PhoneNumber >& phone() const;
    inline ::google::protobuf::RepeatedPtrField< ::tutorial::Person_PhoneNumber >* mutable_phone();
    inline const ::tutorial::Person_PhoneNumber& phone(int index) const;
    inline ::tutorial::Person_PhoneNumber* mutable_phone(int index);
    inline ::tutorial::Person_PhoneNumber* add_phone();

### 常用方法
get 方法和 addressbook.proto 中定义的字段名转成全小写一样; set 方法 set_全小写字段名; has_ 方法检测 optional或required 字段是否被设置,如果被设置，return true; clear_ 方法将字段 unset,恢复到字段的 empty state

string 类型有 mutable_ 方法直接获取string 指针; 假设 email 字段没有被设置,此时调用 mutable_email(),email 将会被初始化为 empty string "".

对于 repeated 字段(比如phone),有些特殊的方法:
_size() 返回 repeated 字段的个数
get 方法可以用 index 获取 tutorial::Person_PhoneNumber
mutable_phone(index) 更新指定 index 的 tutorial::Person_PhoneNumber 数据
add_phone() 添加一个 tutorial::Person_PhoneNumber 到记录中去

### 枚举和嵌套类
枚举类型 tutorial::Person::PhoneType，枚举的值 tutorial::Person::MOBILE 等..

### 操作 message 的函数(实现了 Message 类中的接口)
    bool IsInitialized() const;// 检查是否所有 required 字段都已被设置
    string DebugString() const;// 返回一个人类可读的字符串
    void CopyFrom(const Person& from);// 拷贝构造函数
    void Clear();// 清除所有字段的值,所有字段设置为默认值
### 解析和序列化的函数
    bool SerializeToString(string* output) const;//serializes the message and stores the bytes in the given string. Note that the bytes are binary, not text; we only use the string class as a convenient container.
    bool ParseFromString(const string& data);// parses a message from the given string.
    bool SerializeToOstream(ostream* output) const;// writes the message to the given C++ ostream.
    bool ParseFromIstream(istream* input);// parses a message from the given C++ istream.

## 读写
    
## 修改 .proto 文件注意事项
1. 所有现有字段的 id 都不能更改
2. 不能增加或删除任何 required 字段
3. 可以删除 optional 或 repeated 字段
4. 可以增加任何 optional 或 repeated 字段,但必须使用新的id(那些被删除过的字段id也不能使用)
## 优化 tips
1. 复用声明过的 message objects.
2. 尝试使用 google tcmalloc, 对于从多个线程分配大量的小内存对象有优化。
## 小的 demo, 包含 stl, 基本类型,读写

GOOGLE_PROTOBUF_VERIFY_VERSION best practise.

// Optional:  Delete all global objects allocated by libprotobuf.
  google::protobuf::ShutdownProtobufLibrary();

## 待细查链接
https://developers.google.com/protocol-buffers/docs/reference/cpp-generated#invocation

https://developers.google.com/protocol-buffers/docs/encoding#embedded
