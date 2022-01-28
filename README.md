# Maimu Streamer

Multithreaded and multiplatform live streaming recorder.

多线程和多平台录播姬。



### 特点

使用 sqlite 数据库进行数据交换，添加或删除直播间无需重启程序，可同时监控和录制多个直播间。



### 支持平台

- [x] BiliBIli
- [ ] Youtube（待开发）



### 安装方法

1. 自行编译

   需要安装：[go 1.17.x](https://go.dev/dl/)
   
   ```bash
   git clone https://github.com/ciisaichan/MaimuStreamer
   cd MaimuStreamer
   go mod tidy
   go build -ldflags="-s -w" .
   ```
   
2. 下载已编译版本

   请看：[releases](https://github.com/ciisaichan/MaimuStreamer/releases/)。



### 使用方法

#### 启动任务池

```bash
./MaimuStreamer run
```

启动任务池是关键步骤，监控以及录播都由任务池处理。

可选参数：

`-c`： 手动指定配置文件路径，默认读取目录下的 config.json 。



#### 添加需要录制的直播间

```bash
./MaimuStreamer addroom -p 平台 -r 直播间ID -n 名称
```

必须参数：

`-p`：直播平台的英文名称，程序内部固定（例如：bilibili），拼写错误无效。

`-r`：直播间ID，不同平台格式不一

`-n`：自定义名称，可以自己定义，用于生成录播文件名，尽可能避免使用特殊字符（`/\<>?:"|*` 等）

可选参数：

`-d`：手动指定数据库文件路径，默认读取目录下的 data.db



#### 查看已添加的直播间

```bash
./MaimuStreamer listrooms
```

可选参数：

`-d`：手动指定数据库文件路径，默认读取目录下的 data.db



#### 删除直播间

```bash
./MaimuStreamer delroom -i 数据表ID
```

必须参数：

`-i`：数据表ID，可以通过 **查看已添加的直播间** 方法查看第一列对应的ID。

可选参数：

`-d`：手动指定数据库文件路径，默认读取目录下的 data.db



#### 编辑直播间

```bash
./MaimuStreamer editroom -i 数据表ID -p 平台 -r 直播间ID -n 名称
```

必须参数：

`-i`：数据表ID，可以通过 **查看已添加的直播间** 方法查看第一列对应的ID。

`-p`：直播平台的英文名称，程序内部固定（例如：bilibili），拼写错误无效。

`-r`：直播间ID，不同平台格式不一

`-n`：自定义名称，可以自己定义，用于生成录播文件名，尽可能避免使用特殊字符（`/\<>?:"|*` 等）

可选参数：

`-d`：手动指定数据库文件路径，默认读取目录下的 data.db



#### 列出正在进行的任务

```bash
./MaimuStreamer listasks
```

可选参数：

`-d`：手动指定数据库文件路径，默认读取目录下的 data.db



### 关于开发

如果需要接入其他自动化程序，可以直接操作 sqlite 数据库，无需通过命令调用（上述增删改查操作也是通过操作数据库方式实现）。
