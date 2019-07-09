# boilerplate

## 初始化

首次启动时，可开启`kuu.json`中的`gorm:migrate`标记自动建表：

```json
{
  "gorm:migrate": true
}
```

> 注意：初始化完成后，请关闭`"gorm:migrate":false`，否则启动会很慢。

## internal 代码目录
- 模型 
   1. 新建模型：一个模型一个文件(主子表的可以放同一文件)，放到`[model-name]/model.go`下（也可再新建子文件夹进行二次分类）
   1. 注册模型：定义完成后，记得在`main.go`中通过指针进行注册
- 路由
   1. 新建路由：可以将一组前缀相同的路由定义在一个文件中，放到`[model-name]/route.go`下
   1. 注册路由：定义完成后，记得在`main.go`中进行注册
   1. 自定义路由需要在`docs/openapi.json`补充接口文档
   
## 中间件

1. 新建中间件：放到`middlewares`文件夹下

1. 注册中间件：定义完成后，记得在`main.go`中进行注册
