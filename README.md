# DouYin-Proj
极简抖音项目


教程参考：
1. https://juejin.cn/post/7195903568342482999
2. https://juejin.cn/post/7196468749179928633


修改为gin版本：
1.上传视频功能暂未实现，需要在feed.go文件中修改url来播放视频
![image](https://user-images.githubusercontent.com/58996015/216958350-45bbf600-2041-4fbc-b7c1-cf8372b24b93.png)

2.修改了数据库的数据类型，需要drop一下数据库再创建个

3.修改了一下项目结构

4.需要安装ffmpeg，并配置环境变量 用于截取视频第一帧作为封面
