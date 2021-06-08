2. 进程:
    内存中的数据+寄存器中的值+堆栈中的指令+打开的文件io+其他设备的状态 的镜像集合
3. namespace:
    docker 项目帮助启动进程时 添加不同的 namespace 参数，使用namespace 达到沙盒的 目的
4. Cgroup: 是用作限制当前进程使用宿主机资源的配置，容器在一启动的时候可以添加参数

   rootfs
    
5. 虚拟机和docker的区别
    虚拟机通过hypervisor 隔离了 宿主 和 app
    docker 只是在进程启动的时候添加了一些参数使得 进程能够 隔离 
    但是 与宿主的 隔离性没有 虚拟机强，安全性也不是很高

6. 容器的镜像
    1. docker exec 原理
        docker exec 能够进入容器内部进行操作
        每个pid 都有对应着所关联的 所有ns namespace
        一旦一个进程通过setns(ns_path,进程)
        那么这个进程看到的视图是当前的ns视图  
    2. docker volume 
        docker -v /home:/test ...
        dockerinit 进程启动时,将宿主机的 home 目录真个挂载到 容器/test
        docker commit 不会提交写入到mnt中的修改，对于commit 来说 只有一个/test空白文件 
    3. 容器的层次
        从下到上:
            只读层(镜像层)：
            可写层(容器层)：从下往上找要修改的文件，有就修改完后替换下层文件
    4. volumn 挂载
        宿主机获取 容器进程创建的文件
        容器进程获取 到宿主机的文件
7. k8s
    master 控制节点 - node 计算节点 模式
    控制节点: 
        master 包括3部分: api server, schedule 调度, controller manager 编排
    计算节点:
        kublet 进行 net storage ,容器运行时交互 等
        node 节点的状态存储(etcd/consul)
        
    master 节点是核心
        pod: pod里的容器共享一个 network ns, 同一个数组卷 提高容器间的访问效率
        service: 为每个pod 绑定 service，service提供 pod ip等信息，k8s用于维护service的更新
        
8. mount namespace && change rootfs 
    是对容器进程挂载点的认知修改，只有在挂载动作发生后，才会改变容器的挂载视图，之前都是继承宿主机的视图
    作用: 改变容器的 运行根目录。对于宿主机是不可见的，也称为容器镜像
    
    过程: 
    启动namespace，指定cgroup 参数 ，change root fs 更改容器根目录
    
9. docker 镜像的层次
    rootfs 是整个镜像的基础
    后续的操作只对 rootfs 做增量的 更新 -> union fs
    1. 只读层(镜像层)
    2. 可读可写(容器层): 对文件的增删改都是通过增量文件的方式进行 添加
    3. init: 间与 可读可写 与 只读 之间，如修改hostname 文件
    
    评论区 copy on write 的方式作用在 docker 增量更新：1. 上面的读写层通常也称为容器层，下面的只读层称为镜像层，所有的增删查改操作都只会作用在容器层，相同的文件上层会覆盖掉下层。知道这一点，就不难理解镜像文件的修改，比如修改一个文件的时候，首先会从上到下查找有没有这个文件，找到，就复制到容器层中，修改，修改的结果就会作用到下层的文件，这种方式也被称为copy-on-write。
    
10. 挂载volume
    容器启动之前 所有的 aufs 都会在.../diff 文件下
    diff.* 在容器启动之后 会被挂载在 mnt 目录下(容器静态视图=容器镜像)
    所以在 chroot 之前把外部的 mnt 挂载到mnt 目录下
    
    mnt 的实现原理:
        dentry(目录) 修改指向 inode(文件) 节点的指针
        
11. kubelet 
    通过 CRI container runtime interface 与 docker 容器进行交互
    通过 CNI container net interface
    通过 CSI container storage interface



12. k8s 中 kubectl apply -f yaml 
    1. namespace
    2. deployment: deployment control spc.template
    
   ```yaml
   k8s api obj 描述文件
apiVersion: apps/v1
kind: Deployment
metadata: #对象的元数据 k8s 管理和标识 api obj
  name: nginx-deployment
spec: #对象的特有数据
  selector:
    matchLabels: # deployment label 筛选器
      app: nginx  
  replicas: 2
  template: #具体的pod 细节
    metadata:
      labels:
        app: nginx # 标签 用于deploment 管理识别 api obj
    spec:
      containers:
      - name: nginx
        image: nginx:1.7.9
        ports:
        - containerPort: 80
```

13. projected volume

14 kube-apiserver

//===
deployment.spec 是用来描述 你想要的 实例信息 
  
//==== controller 
statefulset:
    维护 pod 的拓扑状态，通过 headless service dns的方式 解析pod的网络地址
    维护 container 中的 依赖顺序
    
volume:
    k8s:提供什么样的,并将这些容器挂载到 container 中
    pv: 提供 存储资源
    pvc: 消耗 存储资源描述文件
    storageClass: 存储的驱动
    

服务发现和负载均衡
    无状态应用的 亲和性(session)
        使用service(网络四层模型) ready + Endpoints => 域名检查
        概念: 虚ip clusterIP 
            service
            port
            target port
            
            endpoint controller: watch service + port
            kubeproxy: load balance 
1. iptables
    简介:
        NetFilter: 4个表 5个chain
            pre      
        DNAT: 使用ip table 映射 起到流量转发的作用 ;并且可以 对每个 转发增加 权重 起到负载均衡的效果；并且可以增加 会话保持的功能 
    问题:
        规则线性匹配时延: chain 太长太多
        规则更新时延: 非增量
        可拓展性
        可用性
    解决方案: 
        ipvs
            路由转发步骤:
                绑定 VIP 网卡: dummy 创建dummy 命令:
2. ipvs
    DirectRouting dr
    Tunneling
    NAT
    
========
DaemonSet
1. 


-=-=-=-=-=-=-=-


pod:
1. 共享 cgroup, namespace, 的单位
 