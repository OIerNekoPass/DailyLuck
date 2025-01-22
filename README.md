# 每日运势图生成

简陋的每日运势图片生成库

### 使用前的配置：

1. 配置运势图片底图：
   
   在本地创建一个文件夹该文件夹下存放运势背景图片 请保证底图分辨率至少有190x412 

   文件夹内的运势图片的编号应从1开始，并且是png格式。如 1.png 2.png 3.png ....

   （仓库的back文件夹中有自用底图）

2. 配置运势背景板图
   
   仓库的front文件夹以及其中的文件请下载到本地并保存，这些图片是用于在底图上绘制今日运势的背景板的。

   例如：你可以保存在本地的"/root/Luck_front"文件夹中
   
3. Set_Pic_Num(num int) 
   
   使用生成运势图的函数前请使用该函数用于设置有多少张运势图片底图

4. Set_event_list(list []string)
   
   该函数用于设置今日运势事件，如果不设置则使用默认的事件列表

### 生成运势图的函数：
```
Gen_Pic(uid int64, bk_dir string, ft_dir string, ttf_path string, out_path string)

uid用于生成随机种子，你可以使用用户的QQ号，或者用户名称的哈希值作为该值

bk_dir为底图的文件夹路径 如 "/root/bk"

ft_dir为运势背景板的文件夹路径 里面放的文件应该是本仓库的front文件夹内的文件 如 "/root/Luck_front"

ttf_path为字体文件的路径 如 "/root/ttf/cute.ttf" 您可以前往各种字体网站进行下载 也可以选择直接下载本项目中包含的ttf字体文件

out_path为输出的文件的路径 如 "/root/output/output.png"
```