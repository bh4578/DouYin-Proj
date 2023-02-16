# DouYin-Proj
极简抖音项目


教程参考：
1. https://juejin.cn/post/7195903568342482999
2. https://juejin.cn/post/7196468749179928633


修改为gin版本：
1.运行前修改publish.go 中ip地址为本机ipv4地址：
![image](https://user-images.githubusercontent.com/58996015/219223327-320045fd-0c77-44c8-b985-175b67f4a11c.png)

2.修改了数据库的数据类型，需要drop一下数据库再创建个

3.修改了一下项目结构

4.需要安装ffmpeg，并配置环境变量 用于截取视频第一帧作为封面
