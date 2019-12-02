1. "gopkg.in/mgo.v2"
    statement
    prepare
2. 获取db连接池中的一个session
    + 连接池中又分为好几种模式 从库和主库
    + 在query 之前加入 拦截器 用于日志
