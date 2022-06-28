### 创建一个索引及文档

我们接下来创建一个叫做 **twitter** 的索引（index），并插入一个文档（document)。我们知道在 RDMS 中，我们通常需要有专用的语句来生产相应的数据库，表格，让后才可以让我们输入相应的记录，但是针对 Elasticsearch 来说，这个是不必须的。我们在左边的窗口中输入：

```ruby
POST twitter/_doc/1
{
  "user": "LY",
  "uid": 1,
  "city": "Beijing",
  "province": "Beijing",
  "country": "China"
}
```
在上面创建文档的过程中，Elasticsearch 可以自动动态地为你创建索 mapping。当我们建立一个索引的第一个文档时，如果你没有创建它的  schema，那么 Elasticsearch 会根据所输入字段的数据进行猜测它的数据类型，比如上面的 user 被被认为是 text 类型，而 uid 将被猜测为整数类型。这种方式我们称之为 schema on write,也即当我们写入第一个文档时，Elasticsearch 会自动帮我们创建相应的 schema。在 Elasticsearch 的术语中，mapping 被称作为 Elasticsearch 的数据 schema。


```ruby
GET twitter/_mapping
```

```ruby
{
  "twitter" : {
    "mappings" : {
      "properties" : {
        "city" : {
          "type" : "text",
          "fields" : {
            "keyword" : {
              "type" : "keyword",
              "ignore_above" : 256
            }
          }
        },
        "country" : {
          "type" : "text",
          "fields" : {
            "keyword" : {
              "type" : "keyword",
              "ignore_above" : 256
            }
          }
        },
        "message" : {
          "type" : "text",
          "fields" : {
            "keyword" : {
              "type" : "keyword",
              "ignore_above" : 256
            }
          }
        },
        "post_date" : {
          "type" : "date"
        },
        "province" : {
          "type" : "text",
          "fields" : {
            "keyword" : {
              "type" : "keyword",
              "ignore_above" : 256
            }
          }
        },
        "uid" : {
          "type" : "long"
        },
        "user" : {
          "type" : "text",
          "fields" : {
            "keyword" : {
              "type" : "keyword",
              "ignore_above" : 256
            }
          }
        }
      }
    }
  }
}
```
Elasticsearch 的数据类型：
- text：全文搜索字符串
- keyword：用于精确字符串匹配和聚合
- date 及 date_nanos：格式化为日期或数字日期的字符串
- byte, short, integer, long：整数类型
- boolean：布尔类型
- float，double，half_float：浮点数类型
- 分级的类型：object 及 nested


在默认的情况下，Elasticsearch 可以理解你正在索引的文档的结构并自动创建映射（mapping）定义。 这称为显式映射（Explicit mapping）创建。在绝大多数的情况下，它工作的非常好。使用显式映射可以开始使用无模式（schemaless）方法快速摄取数据，而无需担心字段类型。 因此，为了在索引中获得更好的结果和性能，我们有时需要需要手动定义映射。 微调映射带来了一些优势，例如：

- 减少磁盘上的索引大小（禁用自定义字段的功能）
- 仅索引感兴趣的字段（一般加速）
- 用于快速搜索或实时分析（例如聚合）
- 正确定义字段是否必须分词为多个 token 或单个 token
- 定义映射类型，例如地理点、suggester、向量等


假如，我们想创建一个索引 test，并且含有 id 及 message 字段。id 字段为 keyword 类型，而 message 字段为 text 类型，那么我们可以使用如下的方法来创建：
```ruby
PUT test
{
  "mappings": {
    "properties": {
      "id": {
        "type": "keyword"
      },
      "message": {
        "type": "text"
      }
    }
  }
```

如果不喜欢自动创建一个index，可以修改如下的一个设置:
```ruby
PUT _cluster/settings
{
    "persistent": {
        "action.auto_create_index": "false" 
    }
}
```

通常对一个通过上面方法写入到 Elasticsearch 的文档，在默认的情况下并不马上可以进行搜索。这是因为在 Elasticsearch 的设计中，有一个叫做 refresh 的操作。它可以使更改可见以进行搜索的操作。通常会有一个 refresh timer 来定时完成这个操作。这个周期为1秒。这也是我们通常所说的 Elasticsearch 可以实现秒级的搜索。当然这个 timer 的周期也可以在索引的[设置](https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-update-settings.html#reset-index-setting) 中进行配置。如果我们想让我们的结果马上可以对搜索可见，我们可以用如下的方法

```ruby
PUT twitter/_doc/1?refresh=true
{
  "user": "LY",
  "uid": 1,
  "city": "Beijing",
  "province": "Beijing",
  "country": "China"
}
```

上面的方式可以强制使 Elasticsearch 进行 refresh 的操作，当然这个是有代价的。频繁的进行这种操作，可以使我们的 Elasticsearch 变得非常慢。另外一种方式是通过设置 refresh=wait_for。这样相当于一个同步的操作，它等待下一个 refresh 周期发生完后，才返回。这样可以确保我们在调用上面的接口后，马上可以搜索到我们刚才录入的文档：

```ruby
PUT twitter/_doc/1?refresh=wait_for
{
  "user": "LY",
  "uid": 1,
  "city": "Beijing",
  "province": "Beijing",
  "country": "China"
}
```
![image](https://user-images.githubusercontent.com/26034217/176152454-362ffba3-1e61-4fa9-afba-0481923e4dce.png)

### 修改文档
使用PUT或使用POST(指定ID), 已经存在文档的时候是修改，不存在的时候是更新


 设置异步复制允许我们在主分片上同步执行索引操作，在副本分片上异步执行。这样，API 调用会更快地返回响应操作。我们可以这样来进行调用：
 ```ruby
POST my_index/_doc?replication=async
{
  "content": "this is really cool"
}
 ```

 在关系数据库中，我们通常是对数据库进行搜索，然后才进行修改。在这种情况下，我们事先通常并不知道文档的 id。我们需要通过查询的方式来进行查询，让后进行修改。ES 也提供了相应的 REST 接口。
 ```ruby
 
POST twitter/_update_by_query
{
  "query": {
    "match": {
      "user": "LY"
    }
  },
  "script": {
    "source": "ctx._source.city = params.city;ctx._source.province = params.province;ctx._source.country = params.country",
    "lang": "painless",
    "params": {
      "city": "Beijing",
      "province": "Beijing",
      "country": "China"
    }
  }
 ```

 检查一个文档是否存在
 ```ruby
 HEAD twitter/_doc/1
 ```
删除
```ruby
DELETE twitter/_doc/1
```
批处理
```ruby
POST _bulk
{ "index" : { "_index" : "twitter", "_id": 1} }
{"user":"双榆树-张三","message":"今儿天气不错啊，出去转转去","uid":2,"age":20,"city":"北京","province":"北京","country":"中国","address":"中国北京市海淀区","location":{"lat":"39.970718","lon":"116.325747"}}
{ "index" : { "_index" : "twitter", "_id": 2 }}
{"user":"东城区-老刘","message":"出发，下一站云南！","uid":3,"age":30,"city":"北京","province":"北京","country":"中国","address":"中国北京市东城区台基厂三条3号","location":{"lat":"39.904313","lon":"116.412754"}}
{ "index" : { "_index" : "twitter", "_id": 3} }
{"user":"东城区-李四","message":"happy birthday!","uid":4,"age":30,"city":"北京","province":"北京","country":"中国","address":"中国北京市东城区","location":{"lat":"39.893801","lon":"116.408986"}}
{ "index" : { "_index" : "twitter", "_id": 4} }
{"user":"朝阳区-老贾","message":"123,gogogo","uid":5,"age":35,"city":"北京","province":"北京","country":"中国","address":"中国北京市朝阳区建国门","location":{"lat":"39.718256","lon":"116.367910"}}
{ "index" : { "_index" : "twitter", "_id": 5} }
{"user":"朝阳区-老王","message":"Happy BirthDay My Friend!","uid":6,"age":50,"city":"北京","province":"北京","country":"中国","address":"中国北京市朝阳区国贸","location":{"lat":"39.918256","lon":"116.467910"}}
{ "index" : { "_index" : "twitter", "_id": 6} }
{"user":"虹桥-老吴","message":"好友来了都今天我生日，好友来了,什么 birthday happy 就成!","uid":7,"age":90,"city":"上海","province":"上海","country":"中国","address":"中国上海市闵行区","location":{"lat":"31.175927","lon":"121.383328"}}
```
数量
```ruby
GET twitter/_count
```
