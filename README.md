# updateGithubIP
因Github被墙，总是登不上，Git用不了。总是自己去查找Github可用IP挺麻烦的，所以写了个这样的程序自动更新Github服务器地址到hosts文件的配置中。系统默认为windows系统的hosts文件路径，如果为其它系统，可在conf文件中对hosts文件路径进行修改。

# 2023年05月19日更新：
采用了521xueweihan提供的Github可用IP地址，地址为：https://raw.hellogithub.com/hosts.json

感谢[@521xueweihan](https://github.com/521xueweihan) 在其项目 [521xueweihan/GitHub520](https://github.com/521xueweihan/GitHub520) 中提供的Github可用IP地址，地址为：https://raw.hellogithub.com/hosts.json
# 使用方法：
## 第一步
首先运行 go build 命令进行编译得到可执行文件

或者从release（发行版）中直接下载相应版本的可执行文件，注意操作系统的区别：macos请选择darwin版本，linux请选择linux版本，windows请选择windows版本

在最新的发行版中，相应操作系统应当下载的版本已在其中描述，可直接至发行版中找相应发行版进行下载。
## 第二步
将conf文件和可执行文件放在同一目录中
## 第三步（Windows系统可略过）
若为linux或macos系统，需要更改conf文件中的配置为相应系统的hosts文件路径
## 第四步
若为linux或macos系统，请以root用户运行该可执行文件

若为windows系统，请右键该文件，选择“以管理员身份运行”该文件
## 第五步
上述四步完毕后，即可访问github以及使用git。如果仍然出现无法访问的情况，请再次运行该可执行文件以获取新的IP地址，直至能访问成功为止。

## 第六步
若所有IP均已测试完毕，则手动清空bannedIP.txt文件中的内容，然后再次尝试即可。

# 注
一般情况下，运行一次即可访问github了。如果不行，那就多运行几次，换新IP就好了。
