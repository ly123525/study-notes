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

## env
Rack用一个环境参数调用Rack应用程序，他是一个hash的实例
```ruby
#!/usr/bin/env ruby
require 'rack'
require 'webrick'
def pp(env)
 env.map {|k, v|  "#{k}=>#{v}"}.sort.join("\n")
end

Rack::Handler::WEBrick.run lambda {|env| [200, {}, [pp(env)]]}, :Port => 3000
```
浏览器输入`http://localhost:3000/hello`
```ruby
GATEWAY_INTERFACE=>CGI/1.1
HTTP_ACCEPT=>text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9
HTTP_ACCEPT_ENCODING=>gzip, deflate, br
HTTP_ACCEPT_LANGUAGE=>zh-CN,zh;q=0.9,en;q=0.8
HTTP_CONNECTION=>keep-alive
HTTP_COOKIE=>_sso_session=WVFqWWpFdmFJSndLczh6NXZBSkpHQU5ITHB1NWU4cXd3Ui9IcDkrSXRTNDhZK3VPL1cweVMwZ2tzbEN5cTZsN3V3SFZYdFdYb1FwdmF0amR6N3ZLSlcvdU83ZTFObDUxTzFIR2gwZ2M4ZlhuR3NrL2Q3c1gvdWI1MFpZdFp4djdCNnRtOGR3UmRSM21VSGcrWkYySSsvd1pERWpScGhtU2NTNzhuODFjY1hBekhJUnRrd0hJUGtYQ2FTdGkrQXhoRm42Z00rNUJFaXlOb09TbUZYR1dRQlFsNGdUV2NEVzZONDRLRGxUd0owdGhDeDY0NnRRS0k3Ymdpb1JtM0ZyN3RCRzJVbERaVWoweGFQcVFTM2R2aEM1endPVVBxSmJRUVd5YXRENWFmckw3WGJ2SllSa2ZVQjY1Z3dLRmt5OGItLVl4NktqdDlJRS9FSjZ3WjdMbGxEM2c9PQ%3D%3D--843c4395cb0c991dc437236fd3df06914fec75be; _session_id=ANskyocvRx2xkiKz3rTi8Q6a%2BgUpUZPUkupBmbTVf%2FuTC1fOmAoVd2XMyNePUYI4zeMfyUzl%2Fw7AAIz%2FQq%2BTP2LOM5OVKbJsi4nAJBHuIoG8wdjlTUTAlkR%2BKFsrh4Jw3YB4lf3ehr6h2dwv%2BGM0od%2BtW2q3kRf3eE%2B8OXf%2F1qS0NHepz3Qdd4iLuLcO%2BfXqghjVRFqUzAvRfIsqfIexwmVkcRRKaaKij8nCNlsTfACBFYdW4ulWO%2Fku6mZAWTaQLucpZ6EsC1yUf9pA2ziCmDDKNWtkDPX2uqJd5kuqDDgFLhQHrcH%2F%2FC7igqQFXpv%2BqPJWnRWD3kNTnCe5fM5ZOLXRb63a1vy0YZ7kGWSuylQgGnnscDqkOpyclALMNyUiXfOm8j3klXDlcJzo%2FP6n5JgkFBl6jbd4LTwLv6hiqwnsDPQ2Q%2F3k%2B8ITuZmMSrCnDohU%2FTMuyW14jQoqtPgCUsazLh9pj4cxuEMEUiwKM%2FKwIoEpgR3AaAwlzVkEZdHiFD5ZJNSkh19%2FH6fm%2FkrBiAEZatxnZYSopTQiU%2Fhlye6NzZG8oKPWMM6BxiR0x2kW7KRZsY7zqtFSvxOquU1%2FrknZA37wtVFDNt2rTHmzjkGQy2jPb%2BiucymA0g4wqbg%3D--0f1VYro22zioF8d7--DGmljcIKok2sTgU2RZyG%2FA%3D%3D
HTTP_HOST=>localhost:3000
HTTP_SEC_CH_UA=>" Not A;Brand";v="99", "Chromium";v="99", "Google Chrome";v="99"
HTTP_SEC_CH_UA_MOBILE=>?0
HTTP_SEC_CH_UA_PLATFORM=>"macOS"
HTTP_SEC_FETCH_DEST=>document
HTTP_SEC_FETCH_MODE=>navigate
HTTP_SEC_FETCH_SITE=>none
HTTP_SEC_FETCH_USER=>?1
HTTP_UPGRADE_INSECURE_REQUESTS=>1
HTTP_USER_AGENT=>Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36
HTTP_VERSION=>HTTP/1.1
PATH_INFO=>/hello
QUERY_STRING=>
REMOTE_ADDR=>::1
REMOTE_HOST=>::1
REQUEST_METHOD=>GET
REQUEST_PATH=>/hello
REQUEST_URI=>http://localhost:3000/hello
SCRIPT_NAME=>
SERVER_NAME=>localhost
SERVER_PORT=>3000
SERVER_PROTOCOL=>HTTP/1.1
SERVER_SOFTWARE=>WEBrick/1.7.0 (Ruby/3.0.0/2020-12-25)
rack.errors=>#<IO:0x00007fbfd30966a0>
rack.hijack=>#<Proc:0x00007fbfd48b5c90 /Users/hxadmin/.rvm/gems/ruby-3.0.0/gems/rack-2.2.3/lib/rack/handler/webrick.rb:83 (lambda)>
rack.hijack?=>true
rack.hijack_io=>
rack.input=>#<StringIO:0x00007fbfd48b5d08>
rack.multiprocess=>false
rack.multithread=>true
rack.run_once=>false
rack.url_scheme=>http
rack.version=>[1, 3]
```
我们可以看到env的key可以分为两类，一个是大写的类CGI的头，还有一个是rack特定的环境
## Rack相关变量
Rack要求环境中必须包含rack相关的一些变量，这些变量都是rack.xxx的形式。
```ruby
REQUEST_METHOD             http请求方法
PATH_INFO                  路径
QUERY_STRING               ?后的部分
```

## Request
Rack::Request为存取Rack环境提供了方便的接口
request = Rack::Request.new(env)

## Response
Response提供了对响应的状态,HTTP头和内容处理的方便接口
Response提供了两种方法来生产成响应体:
- 直接设置response.body,此时你必须自己设置响应头中Content-Length的值
- response.write增量写入内容，自动填充Content-Length的值
要注意的是你不应该混用这两种方法，浏览器需要用Content-Length头信息决定从服务器端读多少数据，因此这是必须的。不管用什么方法，最后使用response.finish完成。除了一些必要的检查工作外，finish将装配出符合Rack规范的一个数组-------这三个数组有三个成员: 状态码，响应头, 响应体，也就是我们原来手工返回的那个数组。
### response.body
```ruby
#!/usr/bin/env ruby
require 'rack'
require 'webrick'
rack_app = lambda {|env|  
  request = Rack::Request.new(env)
  response = Rack::Response.new
  body = "============header===============\n"
  if request.path_info == '/hello'
    body << 'say hello'
    client = request['client']
    body <<  "from #{client}" if client
  else
    body << "you must provide some info"
  end
  body << "\n============footer==============="
  response.body = [body]
  response.headers['Content-Length'] = body.bytesize
  response.finish
}

Rack::Handler::WEBrick.run rack_app, :Port => 3000
```

如果运行上述程序，你会得到入下错误:
```ruby
Internal Server Error
undefined method `split' for 77:Integer
```
原因在于`response.headers['Content-Length']`的值必须是`string`,改为
```ruby
response.headers['Content-Length'] = body.bytesize.to_s
```

### response.write
```ruby
#!/usr/bin/env ruby
require 'rack'
require 'webrick'
rack_app = lambda {|env|  
  request = Rack::Request.new(env)
  response = Rack::Response.new
  response.write("============header===============\n")
  if request.path_info == '/hello'
    response.write('say hello')
    client = request['client']
    response.write("from #{client}") if client
  else
    response.write("you must provide some info")
  end
  response.write("\n============footer===============")
  response.finish
}

Rack::Handler::WEBrick.run rack_app, :Port => 3000
```
### 状态码

我们可以直接存取Response对象来改变状态码。如果没有任何设置,那么状态码为200。
response.status = 200
Response提供一个重定向的方法#redirect
```ruby
#!/usr/bin/env ruby
require 'rack'
require 'webrick'
rack_app = lambda {|env|  
  request = Rack::Request.new(env)
  response = Rack::Response.new
  if request.path_info == '/redirect'
    response.redirect('http://google.com')
  else
    response.write("here")
  end
  response.finish
}

Rack::Handler::WEBrick.run rack_app, :Port => 3000
```
### 响应头
response.headers是一个hash，例如response.headers['Content-type'] = 'text/plain'
修改上面的代码，让它直接返回普通文本而不是html给浏览器

