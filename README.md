# Oss Picture

[![Build Status](https://travis-ci.org/goecology/oss-pic.svg?branch=master)](https://travis-ci.org/godefault/oss-pic)

一个类似于oss的图片系统

* 根据url，自动裁剪文件
    * 支持多个空间
    * 支持裁剪
    * 支持裁剪后保存图（todo）

* 上传图片
    * 支持本地硬盘(todo)
    * 支持ceph(todo)
    * 支持minio(todo)
    * 支持阿里云oss(todo)

* 授权使用(todo)

* 后台管理界面(todo)

## 使用方法
### 设置配置
在conf/conf.toml里填写里图片的空间名shopadmin，然后指定它的路径``/home/www/server/image/shopadmin``

```
[image]
     [image.egoshop]
      path = "/home/www/image/egoshop"
```

### 使用说明
* 在空间shopadmin的路径``/home/www/image/egoshop``上传一个1.jpg的文件
* 访问http://127.0.0.1:9003/命名空间/文件相对路径/宽度_高度，
例如curl http://127.0.0.1:9003/egoshop/1.jpg/80_80
