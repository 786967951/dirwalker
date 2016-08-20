# dirwalker
1.对整个目录下的所有文件进行遍历，获取所有文件的大小和计算文件的sha1哈希值，记录在一个文件里面                                           
2.可以指定忽略哪些目录、文件，支持通配符                                                                                                 
3.执行示例：go run dirwalker.go d:/test(遍历目标) d:/test/test.txt(结果存储文件) pathpass:d:/test/*1(pathpass1、pathpass2...)  filepass:d:/test/test2/????2.txt(filepass1、filepass2...)
