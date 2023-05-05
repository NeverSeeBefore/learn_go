#

    - go mod init 项目名
        初始化项目
    - go get 包
        下载包
    - go mod download
        根据go.mod 安装依赖
    - go mod tidy
        删除未使用的依赖，添加未引入的一来
    - go mod vendor
        将依赖添加到项目下的vendor文件夹，找不到则下载
    - go clean -modcache
        清除本地已缓存依赖包数据

# go.mod 文件

    require module/path v1.2.3 // direct|indirect
    - require：声明依赖的关键字
    - module/path：依赖包的引入路径
    - v1.2.3：依赖包的版本号。支持以下几种格式：
      - latest：最新版本
      - v1.0.0：详细版本号
    - commit hash：指定某次commit hash
    - 引入某些没有发布过tag版本标识的依赖包时，go.mod中记录的依赖版本信息就会出现类似v0.0.0-20210218074646-139b0bcd549d的格式，由版本号、commit时间和commit的hash值组成。
    - indirect 间接依赖

    replace liwenzhou.com/overtime  => ../overtime
    - 引入本地其他项目下的包

# init

    - 每个包 都可以有init函数
    - 每个包的init函数 会按照引入的顺序，以栈的顺序执行