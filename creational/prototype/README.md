# prototype - 原型模式

<hr/>

原型是一种创建型设计模式， 使你能够复制对象， 甚至是复杂对象， 而又无需使代码依赖它们所属的类。

我们可以在原型模式的基础上，增加一个原型接口管理者(prototypeManager)这样一个角色，该角色可以更好地管理原型接口，使客户端无需知道其具体实现类。