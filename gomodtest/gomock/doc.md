## 执行命令
mockgen -destination=mocks/mock_doer.go -package=mocks github.com/io/gomockt/doer Doer

### 参数含义
-destination: 生成文件的目标位置
-package: 将生成的 mock 代码放到 mocks 包中
github.com/io/gomockt/doer: go mod 中的 module，注意这里必须要遵守类似的格式，如果单纯写成类似 gomockt 这种会无法执行成功
Doer: 为这个接口生成 mock 代码

# 用下面这个
mockgen -destination=./vkecluster_mock.go -source=../vkecluster/client.go -package=mock 