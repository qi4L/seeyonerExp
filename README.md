# seeyoner
致远OA漏洞利用工具

## Usage
```
PS C:\> seeyonerExp.exe -h
一个简单的致远OA安全测试工具，目的是为了协助漏洞自查、修复工作。

Usage:
  Seeyoner [command]

Available Commands:
  exploit     漏洞利用
  help        Help about any command
  list        列出所有漏洞信息
  scan        漏洞检测

Flags:
  -h, --help   help for Seeyoner

Use "Seeyoner [command] --help" for more information about a command.
```
### scan
全漏洞探测：
```
seeyonerExp.exe -u http://xxx.com -i 0
```

指定漏洞探测：
`-vn`指定漏洞编号，可通过`-show`参数查看：
```
D:\>seeyonerExp.exe list

漏洞列表：
1、seeyon<8.0_fastjson反序列化
2、thirdpartyController.do管理员session泄露
3、webmail.do任意文件下载（CNVD-2020-62422）
4、ajax.do未授权&任意文件上传
5、getSessionList泄露Session
6、htmlofficeservlet任意文件上传
7、initDataAssess.jsp信息泄露
8、DownExcelBeanServlet信息泄露
9、createMysql.jsp数据库信息泄露
10、test.jsp路径
11、setextno.jsp路径
12、status.jsp路径（状态监控页面）
```
探测seeyon<8.0_fastjson反序列化漏洞：
```
seeyonerExp.exe scan -u http://xxx.com -i 1
```

### run

以Session泄露+zip文件上传解压为例，指定编号为`2`：

```
seeyonerExp.exe exploit -u http://xxxx.com -i 2
```
  
