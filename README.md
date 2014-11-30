shadowsocks-configfile
======================
[![Build Status](https://travis-ci.org/billzhuang/shadowsocks-config-convert.svg)](https://travis-ci.org/billzhuang/shadowsocks-config-convert)

[shadowsocks-gui](https://github.com/shadowsocks/shadowsocks-gui) does not support 0.0.0.0
but [shadowsocks-local](https://github.com/shadowsocks/shadowsocks-go) does, so this "script" did the job to convert previous gui version's config json to shadowsocks-local one.

this tool also support [shadowsocks-csharp](https://github.com/clowwindy/shadowsocks-csharp) to [shadowsocks-local](https://github.com/shadowsocks/shadowsocks-go), but th csharp version already support share over LAN, i guess you no need this transfer in case.
