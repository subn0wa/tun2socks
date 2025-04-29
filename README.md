Forked from https://github.com/xjasonlyu/tun2socks

Removed yaml config file
Added mutex command line parameter for control process lifetime in windows. Since windows does not support signals, and this allow to stop without "magic" (GenerateConsoleCtrlEvent is a bit a pain in the ass), just with releasing or abandoning of mutex
## Features

- **Universal Proxying**: Transparently routes all network traffic from any application through a proxy.
- **Multi-Protocol**: Supports HTTP/SOCKS4/SOCKS5/Shadowsocks proxies with optional authentication.
- **Cross-Platform**: Runs on Linux/macOS/Windows/FreeBSD/OpenBSD with platform-specific optimizations.
- **Gateway Mode**: Acts as a Layer 3 gateway to route traffic from other devices on the same network.
- **Full IPv6 Compatibility**: Natively supports IPv6; seamlessly tunnels IPv4 over IPv6 and vice versa.
- **User-Space Networking**: Leverages the **[gVisor](https://github.com/google/gvisor)** network stack for enhanced
  performance and flexibility.

## Documentation

- [Install from Source](https://github.com/xjasonlyu/tun2socks/wiki/Install-from-Source)
- [Quickstart Examples](https://github.com/xjasonlyu/tun2socks/wiki/Examples)
- [Memory Optimization](https://github.com/xjasonlyu/tun2socks/wiki/Memory-Optimization)

Full documentation and technical guides can be found at [Wiki](https://github.com/xjasonlyu/tun2socks/wiki).

