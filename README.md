## 接口

用来自https://domainhouse.buzz/ui/

子域名查询接口写了个小工具

```
https://domainhouse.buzz/query.php?token=9378409013576a0cb7c6fac863d5dfe03f0d288a&domain=
```

![image](https://user-images.githubusercontent.com/38073810/159868278-2c78cb3a-fe09-43e7-b270-b61b35f0edf3.png)



## 使用

~~~
[shym]% ./main -h
Usage of ./main:
  -d string
    	exalple [domainhouse.exe -d baidu.com]
  -f string
    	exalple [domainhouse.exe -f target.txt]
~~~

### 单条域名

~~~
domainhouse.exe -d baidu.com
~~~

### 批量查询

~~~
domainhouse.exe -f target.txt
~~~

### 导出结果

~~~
domainhouse.exe -d baidu.com > result.txt
domainhouse.exe -f target.txt > result.txt
~~~

## 下载

win

~~~
https://github.com/login546/domainhouse/releases/download/v0.1/default.zip
~~~

mac

~~~
https://github.com/login546/domainhouse/releases/download/v0.1/default.zip
~~~

linux

~~~
https://github.com/login546/domainhouse/releases/download/v0.1/default.zip
~~~

或自行编译