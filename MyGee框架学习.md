## Gee框架

### 什么是web框架

#### 为什么要用框架
- 各种语言都会自带一些基础的web功能，例如端口监听，映射静态路由，解析http报文等
- 但是相对于实际web开发中，基础web功能缺少了：
    - 动态路由：例如hello/:name，hello/*这类的规则。
    - 鉴权：没有分组/统一鉴权的能力，需要在每个路由映射的handler中实现。
    - 模板：没有统一简化的HTML机制。
    - ...
- 总而言之，就是语言自带的web功能不够灵活，使用起来不够简洁，或者不能满足一些实际的业务需求

#### 框架的功能
- 路由(Routing)：将请求映射到函数，支持动态路由。例如'/hello/:name。
- 模板(Templates)：使用内置模板引擎提供模板渲染机制。
- 工具集(Utilites)：提供对 cookies，headers 等处理机制。
- 插件(Plugin)：本身功能有限，但提供了插件机制。可以选择安装到全局，也可以只针对某几个路由生效。


### Day1 搭建Gee基本框架
#### 实现http.Handler接口
- Go启动web监听的方法是这样的
```golang
http.ListenAndServe(address string,h handler)
```
- 其中，`address`是监听的地址，而`handler`是访问该地址时运行的实例
- 而`handler`是一个接口，需要实现方法`ServeHTTP(w ResponseWriter,r *Request)`
- 所以只要传入任何实现了`ServeHTTP`接口的实例，所有的http请求就都交给了该实例处理


### Day2 实现上下文Context
#### 为什要使用上下文Context
- 对Web服务来说，无非是根据请求*http.Request，构造响应http.ResponseWriter。但是这两个对象提供的接口粒度太细，比如我们要构造一个完整的响应，需要考虑消息头(Header)和消息体(Body)，而 Header 包含了状态码(StatusCode)，消息类型(ContentType)等几乎每次请求都需要设置的信息。因此，如果不进行有效的封装，那么框架的用户将需要写大量重复，繁杂的代码，而且容易出错。针对常用场景，能够高效地构造出 HTTP 响应是一个好的框架必须考虑的点
- 针对使用场景，封装*http.Request和http.ResponseWriter的方法，简化相关接口的调用，只是设计 Context 的原因之一。对于框架来说，还需要支撑额外的功能。例如，将来解析动态路由/hello/:name，参数:name的值放在哪呢？再比如，框架需要支持中间件，那中间件产生的信息放在哪呢？Context 随着每一个请求的出现而产生，请求的结束而销毁，和当前请求强相关的信息都应由 Context 承载。因此，设计 Context 结构，扩展性和复杂性留在了内部，而对外简化了接口。路由的处理函数，以及将要实现的中间件，参数都统一使用 Context 实例， Context 就像一次会话的百宝箱，可以找到任何东西

### day3 实现前缀路由树Router
#### 为什么要是用前缀树存储路由
- 使用map存储路由虽然方便，但是不能存储动态路由，只能存储指定的静态路由
- 使用前缀树存储与解析路由可以实现：
    - 参数匹配`:`。例如 /p/:lang/doc，可以匹配 /p/c/doc 和 /p/go/doc。
    - 通配`*`。例如 /static/*filepath，可以匹配/static/fav.ico，也可以匹配/static/js/jQuery.js，这种模式常用于静态服务器，能够递归地匹配子路径。


### day4 实现分组控制Group
- 中间件(middlewares)，简单说，就是非业务的技术类组件。Web 框架本身不可能去理解所有的业务，因而不可能实现所有的功能。因此，框架需要有一个插口，允许用户自己定义功能，嵌入到框架中，仿佛这个功能是框架原生支持的一样。


