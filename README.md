## 动手写jvm

### ch01
    主要解析命令行指令
### ch02
    类加载文件实现，根据文件名读取二进制字节流，双亲委派机制的实现

### ch03
    字节码二进制解析
    -Xjre "/Users/mr.yang/Library/Java/JavaVirtualMachines/adopt-openjdk-1.8.0_292-1/Contents/Home/jre" java.lang.String

### ch04
    运行时数据

### ch05
    解释器
    go build
    ./ch05 -Xjre "/Users/mr.yang/Library/Java/JavaVirtualMachines/adopt-openjdk-1.8.0_292-1/Contents/Home/jre" jvmgo.book.ch05.GaussTest

### ch06

    ./ch06 -Xjre "/Users/mr.yang/Library/Java/JavaVirtualMachines/adopt-openjdk-1.8.0_292-1/Contents/Home/jre" jvmgo.book.ch06.MyObject

### ch07

windows java path
C:\Program Files\AdoptOpenJDK\jdk-8.0.292.10-hotspot\jre

    .\ch07.exe -Xjre "C:\\Program Files\\AdoptOpenJDK\\jdk-8.0.292.10-hotspot\\jre" jvmgo.book.ch07.InvokeDemo

### ch08 
    .\ch08.exe -Xjre "C:\\Program Files\\AdoptOpenJDK\\jdk-8.0.292.10-hotspot\\jre" jvmgo.book.ch08.HelloWorld


### ch09 
    .\ch09.exe -Xjre "C:\\Program Files\\AdoptOpenJDK\\jdk-8.0.292.10-hotspot\\jre" jvmgo.book.ch09.StringTest
### ch10 异常处理
    .\ch10.exe -Xjre "C:\\Program Files\\AdoptOpenJDK\\jdk-8.0.292.10-hotspot\\jre" jvmgo.book.ch10.ParseIntTest 123