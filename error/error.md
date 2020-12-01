###第二课 error
 
-   Error vs Exception
     
    errors.New() 返回为什么是一个指针？
    pointer.go
        新建的的值判定
        标准库的New返回的地址，比较地址是否一致。避免你自己写的errors.New和别人的写的一样了
    出现奇奇怪怪的bug。若是值，就可能判定相等。取地址是判定两个内存地址是否是一样的。
    
    各语言的演进历史：
        C : 单返回值，一般通过传递指针作为入参，返回值为int表示成功还是失败。
        C++: 引入excption，但是无法知道被调用方会抛出什么异常
        Java：引入了checked exception，方法的所有者必须申明，调用者必须处理，在启动时抛出大量
    的异常时司空见惯的事情，并在他们的调用堆栈中尽职地记录下来。Java 异常不再是异常，
    而是变得司空见惯，他们从良性到灾难性都有使用，异常的严重性由函数的调用者来区分。
   
   go 的处理异常逻辑不引入exception，支持多参数返回，所以在函数签名中带上实现error interface的
   对象，交由调用者判定。
   **如果一个函数返回了（value,error）。你不能对这个value做任何假设，必须先判定error。
   唯一可以忽略error的是，如果你连value也不关心。** 
        errorLogic.go
   
   Go中有panic 的机制，意味着fatal error(就是挂了)。不能假设调用者来解决panic，
   意味着代码不能运行
   **使用多个返回值和一个简单的约定，GO解决了让程序员知道什么时候出了问题，
   并为真正的异常情况保留了panic**
   
   野生goroutine 导致服务挂掉 panic 
   开异步的goroutine
   通常做保护：在底层公共包加上sync包，protect.go
   
   将任务发放到channel 里面，通过任务池来处理
   
   什么情况panic 不处理：
        main 函数里面的初始化强依赖的，比如配置文件内容
   
   强依赖：服务不能起来
   弱依赖：服务能起来
   
   数据启动失败，让服务是否可用，决定于使用场景，读多写少（弱依赖）
   main 中 grpc client，初始化，不成功有两种（blocking， non-blocking）
    1.blocking: 不可用，对业务有影响 
    2.non-blocking : 可用，所依赖的服务可能会降级处理。（通常选用）缺点：启动起来，
        流量一下进来，依赖的
    还没完全起来，可能会报一些错误，之后会好。
    3.non-blocking + 10s ：给你10s去连接，
    
   panic 三个场景：
        1.main函数初始化，看需求是强依赖，还是弱依赖
        2.配置文件的编程，panic 出来
        3.init函数里面一些资源的初始化，不成功也会panic
    
   error的几个demo
        demo.go
        
   **对于真正意外的情况，那些表示不可恢复的程序错误，例如索引越界，不可恢复的环境问题，栈溢出，
   我们才使用panic。对于其他的错误情况，我们应该是期望使用error来判断
   you only need to check the error value if you care about the result**
   
        简单
        考虑失败，而不是成功（plan for failure, not success）有错误立即处理
        没有隐藏的控制流
        完全交给你来控制error
        Error are values
        
###   Handling error
   解决error的几种套路
   
-   一、sentinel error


    预定义的特定错误。使用一个特定值来表示不可能进行进一步处理的做法。对Go，我们使用特定的值来表示
    错误。
    if err == ErrSomething {
    ...
    }
    类似 io.EOF ，
    使用 sentinel 值是最不灵活的错误处理策略，因为调用方必须使用 == 将结果与预先声明的值进行比较，
    当你想要提供更多的上下文是，这就出现了一个问题，因为返回一个不同的错误将破坏相等性检查。
    甚至是一些有意义的fmt.Errorf 携带一些上下文，也会破坏调用者的 == 等值判断，调用者将被迫查看
    error.Error()方法的输出，已查看它是否与特定的字符串匹配。
        handlingError.go  one
    不依赖检查 error.Error 输出
        不应该依赖检测 error.Error的输出，Error方法存在于error 接口主要用于方便程序员使用，
      但不是程序（编写测试可能会依赖这个返回）。这个输出的字符串用于记录日志，输出到stdout等。
    
    没有携带上下文信息
    弊端：
       1. sentinel errors 成为你API公共部分
       
       2. sentinerl errors 在两个包之间创建了依赖
       
       结论：尽可能避免sentinel errors。 
       避免编写的代码中使用sentinel  errors。在标准库有一些使用他们的情况，但这不是一个你应该
       模仿的模式。
       
       对于业务错误来说会这样吗？？
       
-   二、Error Types
    
    
        Error Type 是实现了 error 接口的自定义类型。
            handlingError.go   two    
        调用者可以使用类型断言转换成这个类型，来获取更多的上下文信息。
        与错误值相比，错误类型的一大改进是他们能够包装底层错误以提供更多上下文。比如地址什么操作，
        哪个路径出了什么问题。
        
    调用者要是有类型断言和类型 switch，就要让自定义的 error 变为 public。这种模型会导致和
    调用者产生强耦合，从而导致API 变得脆弱。
    
    结论是尽量避免使用 error types，虽然错误类型比 sentinel errors 更好，因为他们可以捕获
    关于出错的更多上下文，但是 error types 共享error values 许多相同的问题。
    建议是避免错误类型，或者至少避免将他们作为公共API的一部分。
    
-   三、Opaque errors


    不透明错误处理，它要求代码和调用者之间的耦合最少。因为虽知道发生了错误，但你没有能力看到错误
    的内部。作为调用者，关于操作的结果，你所知道的就是它起作用了，或者没有其作用（成功或失败）
    
    不透明错误处理的全部功能-只需返回错误而不假设其内容。
    
    **Assert errors for behaviour,not type**
        handlingError.go
    使用断言错误实现特定的行为，而不是断言错误是特定的类型或值。
    
### 几种套路：
-    一、Indented flow is for errors


        无错误的正常流程代码，将成为一条直线，而不是缩进的代码。
        ways.go  
        正常做法是demo1 ，无缩进
    
-    二、Eliminate error handling by eliminating errors

    //通过删除错误来消除错误处理
        ways.go  
    one 有什么问题，很繁杂，错误的类型和里面的类型一样，直接写成two
    文件建议用 two
    
-    三、ways.go


        建议推荐第三种