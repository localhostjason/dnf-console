# DNF 后台管理系统

### 目录

1. console 
   - 后端源码
   - 技术栈：go+gorm+gin
2. web
   - 前端源码
   - 技术栈：vue3+ts+pinia+ep
3. dist
   - 最终编译程序

### 运行步骤

第一次运行：

```shell
main.exe -i  #初始化数据,只需要执行一次，后续只需要启停
```



后续：

```shell
main.exe -x #debug 程序跑，小白这样跑就行了

#### 下面 比较复制，非小白 只需要 -x 就跑起来了
main.exe -k install # windows 安装服务
main.exe -k start # 启动 deamon 服务
main.exe -k stop # 停止 deamon 服务
main.exe -k uninstall # windows 卸妆服务
```



### 注意点：

在 `dist/config` 中`server.json` 配置项:

更改 DNF 服务器 IP地址

1.  如果想将 系统 部署到 DNF 服务器上。将 `host` 改为 `127.0.0.1`
2. 如果部署到本地，将 `host` 改为 `dnf服务器IP`（允许数据库连接得上）



```shell
{
  "api": {
    "cors": true
  },
  "auth": {
    "realm": "test zone",
    "secret": "f450a7bdbde3416d22474b9fdc2a3636",
    "id_key": "username",
    "timeout": 43200,
    "max_refresh": 3600
  },
  "game_db": {
    "enable": true,
    "mysql": [
      {
        "key": "d_taiwan",
        "user": "game",
        "password": "uu5!^%jg",
        "host": "192.168.43.143",
        "port": 3306,
        "db": "d_taiwan",
        "charset": "utf8",
        "timeout": 5,
        "multi_statements": false,
        "debug": false
      },
      {
        "key": "taiwan_cain",
        "user": "game",
        "password": "uu5!^%jg",
        "host": "192.168.43.143",
        "port": 3306,
        "db": "taiwan_cain",
        "charset": "utf8",
        "timeout": 5,
        "multi_statements": false,
        "debug": false
      },
      {
        "key": "taiwan_cain_2nd",
        "user": "game",
        "password": "uu5!^%jg",
        "host": "192.168.43.143",
        "port": 3306,
        "db": "taiwan_cain_2nd",
        "charset": "utf8",
        "timeout": 5,
        "multi_statements": false,
        "debug": false
      },
      {
        "key": "taiwan_billing",
        "user": "game",
        "password": "uu5!^%jg",
        "host": "192.168.43.143",
        "port": 3306,
        "db": "taiwan_billing",
        "charset": "utf8",
        "timeout": 5,
        "multi_statements": false,
        "debug": false
      },
      {
        "key": "taiwan_login",
        "user": "game",
        "password": "uu5!^%jg",
        "host": "192.168.43.143",
        "port": 3306,
        "db": "taiwan_login",
        "charset": "utf8",
        "timeout": 5,
        "multi_statements": false,
        "debug": false
      }
    ]
  },
  "service": {
    "pid_file": "console.pid",
    "daemon_log": "daemon.log"
  }
}
```



### 最终

打开网页，访问地址： http://127.0.0.1:8088  （注：确保你的端口未被占用）

默认登录：

用户名：admin
密码：123



### 打赏

有木有大佬 请我喝杯咖啡

![](./wxp.png)
