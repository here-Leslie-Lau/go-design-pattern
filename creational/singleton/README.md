# singleton - 单例模式

<hr/>

单例是一种创建型设计模式， 让你能够保证一个类只有一个实例， 并提供一个访问该实例的全局节点。

>通常而言， 单例实例会在结构体首次初始化时创建。 为了实现这一操作， 我们在结构体中定义一个 get­Instance获取实例方法。 该方法将负责创建和返回单例实例。 创建后， 每次调用 get­Instance时都会返回相同的单例实例。