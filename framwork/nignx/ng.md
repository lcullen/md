1. TLS/SSL
    非对称加密:请求方利用自己公布的公钥进行加密，自己收到后利用私钥解密
    对称加密:加密方解密方能够用同一种加密方式进行解密
    
=== http 请求处理的11个阶段 [rf](https://time.geekbang.org/course/detail/138-71462)
postread 阶段 (realip module)
    如何获得真实用户的ip
    TCP 四元组 srcip srcport dstip dstport
    X-Forwarded-For: 是一个数组用来传递 反向代理的一路传递ip
    X-Real-Ip: 真实的用于ip
    基于变量传递 
 
rewrite module
    return code
    return code url 
    return url
        http 1.0 
            301 永久direct 可以被cache
            302 临时重定向 禁止cache 
        http 1.1
            303 temporary allowed modify method forbid cache
            307 temporary forbid modify method forbid cache
            308 forever forbid modify
    syntax rewrite regex replacement[flag]
        --last 用replacement 进行新的location 匹配
        --break 停止当前脚本指令的执行
        --redirect 返回302 重定向
        --permanent 返回301 重定向
    rewrite if syntax

find config module
    location 
        前缀字符串
            * 常规
            * = 精确匹配
            * ^~ 匹配上后则不再进行正则表达式匹配
        正则表达式
            * ~ 大小写敏感
            * ~* 大小写不敏感
            
    匹配规则，二叉匹配树
        1. 精确匹配 =  直接返回
        
        2. 正则匹配 ^~ 匹配后不再匹配 直接返回
        
        3. 记住最长匹配的前缀字符创
            按序 按照正则"最长"匹配 如果正则完全匹配则会优先于最长匹配
         

阶段 access
    ========
    pre_access:
        worker之间通过共享内存通信 
        limit_conn: 共享内存 remote_addr 当做key 限制用户的 并发连接 达到上限默认返回503
        对应 syntax limit_conn_zone key(ip) zone=name:size #default module
        ```nginx 
            limit_conn_zone $binary_remote_addr zone=addr:10m 
            server {
                server_name lxw.com
                root html/
                location /{
                    limit_conn_status 500; #返回code
                    limit_conn_log_level warn; #修改发生限流的时候的日志级别，减少日志io
                    limit_rate 50; #返回字节/每秒
                    limit_conn addr 1; #并发连接数
                }
            }
        ```
    如何限制每个客户端的每秒处理请求数(限流策略) leaky bucket 
        http_limit_req_module 
        syntax limit_req_zone key zone=name:size rate=rate;
    limit_request > limit_conn
    =======
    access:
        控制请求是否可以继续向下访问
        ip白名单
        allow / deny
        简单 auth_basic: 验证用户访问权
    precontent:
        try_files ngx_http_try_files_module 反向代理 先尝试本地文件 后转发到上游代理
            try_files file url
            try_files file =code
            依次试图访问多个url对应的文件,如果存在立即返回,否则返回最后对应 url or code
            location / {
                try_files $url $url2 = 404
            }
    content 
        static root alias; root 会把整个文件子路径加入到 root 指定的目录 alias 只是把匹配到
        location /root {
            root html # 如果匹配到则会访问文件路径 html/root
                       # 如果访问 domain/root/ 则匹配到 html/root/index.html
        }
        location /alias {
            alias html # domain/alias/ 匹配到文件路径 html/index.html
        }
        location ~/alias/(\w) {
            alias html/$1 #domain/alias/first/ 会
        }
        生成待访问文件的三个相关变量
            request_filename : 请求资源的完整路径
            document_root : 由uri 和 root/alias 构成的文件夹路径
            realpath_root : 真实非软链接 路径 
        访问 URL最后没有带/   nginx 直接返回301 重定向
            nginx 通过3个指令来处理301 重定向
            absolute_direct 
            servername_direct
            porr_direct  
        index && autoindex: 当访问url 以 / 结尾 index 模块先于 autoindex 模块生效，
            syntax index file //返回指定的file
        server {
             server_name ***
             listen 80
             location / {
                alias html/
                index index.html
                autoindex_format html /json /jsonp / xml
             }
        }
    ========
    http_log_module 
        syntax log_format 
        default combined
        syntax access_log path/off 其中path 可以添加变量
        日志cache  path 含有变量的时候
            open_log_file_cache max=缓存中的最大句柄数目LRU inactive 文件访问完的多少时间内 min_uses 被使用的多少次数才能进入缓存
        日志压缩
    ====== after content before log 
    http filter module 需要对返回的 header 和body 进行过滤 主要的 4个模块是
        http_copy_filter_module 用来将内核态的response copy 到用户态 ng 里面
        http_postpong_filter_module 处理子请求
        http_hader_filter 构造response header
        http_write_filter 调用内核send response to client
        ====
            ngx_http_sub_filter 对 response 中的字符串进行过滤
                syntax sub_filter string replacement
                sub_filter_once on/off
                sub_filter_types html/json
            ngx_http_addtion_filter_module 对响应之前添加 后者响应之后添加
                add_before_body url //访问上有服务器 在当前响应的内容之前添加内容
    =====ng 变量
        ng_http 框架提供的与用户有关的变量
    =====防盗链的解决方案
        
                
3. post access


4. 进程模型
    + 多进程结构:高可用性,多线程的问题是一个 错误导致整个进程退出
    + master 进程 用于监控管理worker
    + worker 进程 事件驱动:每个worker 一直占用 一个cpu 
    + 共享内存解决通信问题

5. 向master 进程发送信号

====
1. ng使用场景
    1. 静态资源服务器
        通过本地文件提供服务:
        
    2. 反向代理
        缓存
        负载均衡
    3. api服务
        OpenResty
2. ng 命令行
    -s: 表示发送信号:为什么要有-s？ 
    quit(Quit):
    stop(TREM):
    reload(HUP):
        向master发送reload HUP 信号
        master 检查配置文件语法-> 监听新的端口 -> 启动新worker 子进程
        master 发送TERM 给老的 worker
        老worker 关闭监听端口 处理
    reopen(USR1):
    
3. TLS通讯过程(提前准备网络7层协议) [参考](http://www.ruanyifeng.com/blog/2014/02/ssl_tls.html)
    大致流程:
    1. 验证身份
    2. 达成加密组件共识
    3. 传递秘钥 
    4. 加密通信

4. 请求处理流程(ng如何处理http请求)
    建立连接: 
        操作系统内核:
            三次握手 获得连接交给对应的监听端口
        事件模块:
            accept:
                分配连接内存池,并在 HTTP模块 设置回调方法，epoll_ctl 添加 连接 超时定时器
        HTTP模块:
            接受URL,
            接受Header,
    处理请求(11 个模块处理请求):

5. config
    值指令: 可以合并 向上覆盖 父模块 通常意义上讲的是 变量指令
        root:
        access_log:
        gzip:
    动作指令: 不可以合并 
        rewrite:
        proxy_pass:
        content:
     
    
======红黑树


http 模块处理请求 [refer](https://time.geekbang.org/course/detail/138-71458)
1. nginx 会根据hettp header 来确定 需要哪些模块
2. nginx连接建立：
    * 内核接受连接请求 (os core)
    * worker epoll_wait accept request 分配 connection pool(connection_pool_size) (event module)
    * accept 分配连接内存池 (event module)
    * ngx_http_init_connection 设置回调方法 并且注册过期时间 epoll_ctl (http module)
3. nginx 处理请求data (request_pool)
    * 接受uri  分配请求内存池(4k)
        * 状态机解析请求行
            * 方法名
            * url
            * 协议
        * 标识uri
     * 接受hader (cookie host)
        * 状态机解析header 分配内存
        * 标识header 确定server block
        * 移除定时器
        * 11个阶段的请求处理
    

======config 配置

reg
    分组与取值 匹配第一个分组()内容 用 $,只有用 $2 ...
   
server_name 域名
    主域名 server_name_in_redirect on  用于多域名跳转到主域名
  



