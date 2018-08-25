# net

1. Package net provides a **portable** interface for network I/O, including TCP/IP, UDP, domain name resolution, and Unix domain sockets.
   portable对IPV4和IPV6的支持
2. Although the package provides access to low-level networking primitives. 但是不建议使用
3. 重要函数：Dial, Listen, Accept. 重要接口: Conn, Listener
4. The crypto/tls package uses the same interfaces and similar Dial and Listen functions. 支持TLS
