package jvmgo.book.ch06;

public class MyObject{
    public static int staticVar;
    public int instanceVar;
    public static void main(String[] args){
        int x = 32768;
        MyObject myObj = new MyObject();
        MyObject.staticVar = x;
        x = MyObject.staticVar;
        myObj.instanceVar = x;
        x = myObj.instanceVar;
        Object obj = myObj;
        if(obj instanceOf MyObject){
            myObj = (MyObject) obj;
        }
    }

}