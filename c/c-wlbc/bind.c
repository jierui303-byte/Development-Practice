/**
 * @file bind.c
 * @author your name (you@domain.com)
 * @brief 
 * @version 0.1
 * @date 2021-04-16
 * 
 * @copyright Copyright (c) 2021
 * 
 * udp编程步骤：
 * 1. 创建socket套接字
 * 2. 发送数据-sendto函数
 * 3. 绑定-bind函数
 * 4. 接收数据-recvfrom函数
 * 
 * 
 * 
 * 基于udp编程函数可以实现：tftp，广播，多播
 * 
 * 网络编程其实是指使用udp编程函数来实现不同的通信模式
 * 
 * 
 * udp编程(单播)【有对应的udp函数使用】
 * udp实现tftp上传/udp实现tftp下载
 * udp实现广播(局域网)【客户端设置是否允许广播方式发送消息，消息服务端无法拒绝，只能被动接收】
 * udp实现多播(广域网)【客户端谁都可以向多播组发送消息，但消息只有加入多播组的服务端才能接收消息】
 * 
 * 
 * TCP编程[有对应的tcp函数使用]
 * tcp客户端/tcp服务端
 * 
 * 
 * 
 * 
 */

int main(int argc, char *argv[])
{
    //第一步：创建套接字
    //第二步：将服务器的网络信息结构体绑定前进行填充
    //第三步：将网络信息结构体与套接字进行绑定
    //第四步：
}
