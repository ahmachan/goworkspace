
##假设:
> GO项目工作空间:/home/xmage/html/golang/goworkspace 
> GO安装在/home/xmage/developments/golang 

> 在当前使用的SHELL：(~/.bash_profile,~/.bashrc,~/.zshrc) 

### set golang env configuration 

```
export GOPATH=$HOME/html/golang/goworkspace    #设置为go安装的路径 
export GOROOT=/home/xmage/developments/golang         #go工作空间的路径 
export PATH=$PATH:$GOPATH 
export PATH=$PATH:$GOROOT/bin 
```
> source ~/.szhrc使设置立即生效
