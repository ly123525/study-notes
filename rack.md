### 什么是Rack
Rack是Ruby应用服务器和Ruby程序之间的一个接口。

<img width="703" alt="image" src="https://user-images.githubusercontent.com/26034217/161721373-40ce213f-dde8-4789-b80f-f3ff179ada8d.png">

### 为什么Rack
首先Rack提供了一种标准的接口，便于应用程序与应用服务器之间的交互。一个Rack应用程序可以被任何和Rack兼容的应用服务器调用。
目前所有的主流Ruby应用服务器都支持Rack接口，Rack通过一种叫做句柄(handler)的机制实现对应用服务器的支持。目前Rack本身带有的句柄包括：
- FastCGI
- CGI
- WEBrick
- LSWS
- SCGI
- Thin

下面的应用服务器也在他们的代码中包含了Rack Handler
- Ebb
- Fuzed
- Classfish v3
- Phusion Passenger
- Rainbows
- Unicorn
- Zbatery
这意味着所有上述的服务器都以Rack接口的形式调用Rack应用程序。
这些句柄都都位于Rack::Handler命名空间之下，Rack文档中我们可以看到这些类型:
- Rack::Handler::FastCGI
- Rack::Handler::CGI
- Rack::Handler::WEBrick
- Rack::Handler::LSWS
- Rack::Handler::SCGI
- Rack::Handler::Thin

查看Rack内嵌的handler
```ruby
3.0.0 :006 > require 'rack'
 => true
3.0.0 :007 > Rack::Handler.constants
 => [:FastCGI, :CGI, :WEBrick, :LSWS, :SCGI, :Thin]
3.0.0 :008 >
```
所有的handler都有一个run方法, 你可以用：
```ruby
Rack::Handler::WEBrick.run
```
## 一个可被call的对象
那么一个Rack应用程序需要符合什么条件
一个Rack应用程序是一个Ruby对象，只要这个对象能响应call, 这个call方法需要接受一个参数，即环境(env)对象，需要返回一个数组，这个数组有三个对象
- 一个状态(status),即http协议定义的状态码
- 一个头(headers), 他可能是一个hash，其中包含所有的http头
- 一个body，字符串数组

一个简单的Rack应用服务器
```ruby
:001 > require 'webrick'
=> false
:002 > require 'rack'
=> false
:003 > require 'webrick'
=> false
:004 > rack_app = lambda {|env| [200, {}, ['hello']]}
=> #<Proc:0x00007f901da13d08 (irb):21 (lambda)>
:005 > Rack::Handler::WEBrick.run rack_app, :Port => 3000
[2022-04-05 18:16:23] INFO  WEBrick 1.7.0
[2022-04-05 18:16:23] INFO  ruby 3.0.0 (2020-12-25) [x86_64-darwin18]
[2022-04-05 18:16:23] INFO  WEBrick::HTTPServer#start: pid=45141 port=3000
```

当然一个合法的Rack应用也可以是任何对象，只要他的类定义了call方法
```ruby
class AnyClass
  def call(env)
    [200, {}, ['hello from AnyClass instance with call defined']]
  end
end

rack_app = AnyClass.new
Rack::Handler::WEBrick.run rack_app, :Port => 3000
```
