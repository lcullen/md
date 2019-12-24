接口隔离原则:
    接口的纯洁性:
        1. 接口粒度要尽可能的小，可拓展
        2. 接口粒度不能小于单一职责 (only one thine can change status)
        3. 高内聚 
迪米特法则
    最少知识原则
    1. 朋友类：
        通过参数、成员变量出现在类中， 叫做朋友类
    2. talking to u immediate friend 只和最亲密的朋友交流
    3. know less about u friend and keep distance 不应该通过朋友类的返回值当做执行朋友下一个动作的条件
    
开闭原则
    对修改关闭对拓展开放
    1. 如果原有的实现方式不需要改变， 但是规则变化了 需要对最后的输出结果根据规则进行改变。继承原有的类 根据原有的类的结果进行包装后再输出
        如果以测试驱动开发，修改了原有的实现 需要修改原来的单元测试。
    2. 抽象约束
        + 依赖朋友类的interface 而不是具体的类，interface 中还可以由实类去实现 interface 中的约束
        + 元数据 metadata 控制模块行为
    3. 配置优先
    4. 协议先行
    5. 封装变化

1. option 模式 (皮裤套棉裤) 可以参考[refer GRPC](https://github.com/grpc/grpc-go)
   Go 是支持头等函数的语言 也就是 func 能够成为一个变量
   用func(*V) 模式自定义可改变的 func
   New("init", func changeModel(*V)) struct V
