墨云视窗 后端服务项目模板
---

### 使用:

保持最新：`cookiecutter https://github.com/inkneko-software/nekowindow-service-layout`

使用本地：`cookiecutter nekowindow-service-layout`

### 服务编写

1. 将api目录移动至大仓
2. 修改proto文件，定义服务，`make api`生成代码
3. 进入`internal/service`目录，拷贝生成的grpc代码中XXXService.UnimplementedServer中的所有函数到服务定义文件.go中
4. 编写biz层函数，编写biz层所需的repo接口
5. 在data层编写数据源，实现biz层所需的repo接口
6. 实现http.go中代码外部接口

