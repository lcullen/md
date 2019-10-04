1. error design
    错误分成三大类: 
        1. 资源错误：
        2. 程序错误：
        3. 用户错误：
    逻辑上区分 返回错误 vs 异常捕获
    1. 不期望错误发生使用 异常捕获
    2. 可能会发生 返回错误给客户端
    
    error in async
        [参考 Go promise 模式](https://github.com/fanliao/go-promise)
    
2. time management 
    + 将被动 转换 为主动
    + 有条件的说不：我不能说不，但是我要有条件地说是。而且，我要把你给我的压力再还给你，看似我给了需求方选择，实际上，我掌握了主动。
    + 形成长期习惯 + 需要正负反馈 （一个月一次的interview）
    + 反思 参与
    
    
3. 故障处理

4. 

5. promise 设计模式
    future 对象
    1. kafka producer 利用future 模式实现同步发送机制


11. gateway 与 facade 相似
    提供外部访问 系统的 入口
    设计要点:
       路由
       服务注册
       lb
       弹力设计
       安全
