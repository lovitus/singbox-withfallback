---
description: singbox-withfallback 的 fork 说明。
---

# Fork 说明

此仓库是 `SagerNet/sing-box` 的 fork。

## 仓库地址

- 上游仓库: https://github.com/SagerNet/sing-box
- 本 fork: https://github.com/lovitus/singbox-withfallback

## 与上游的主要区别

1. 新增了 `fallback` 出站策略组类型。
2. `fallback` 按配置顺序检查节点，并在高优先级节点恢复后自动回切。
3. `fallback` 使用阈值控制切换:
   - `failure_threshold`: 连续失败多少次后判定当前节点不可用（默认 `3`）
   - `success_threshold`: 连续成功多少次后判定节点恢复（默认 `3`）
4. 本 fork 的 release 二进制使用上游默认特性标签构建（包含 QUIC）。

## Fallback 文档

- English: [/configuration/outbound/fallback/](/configuration/outbound/fallback/)
- 中文: [/zh/configuration/outbound/fallback/](/zh/configuration/outbound/fallback/)

## 本 Fork Release 使用的构建标签

```text
with_gvisor,with_quic,with_dhcp,with_wireguard,with_utls,with_acme,with_clash_api,with_tailscale,with_ccm,with_ocm,badlinkname,tfogo_checklinkname0
```

## 如何验证二进制特性

运行:

```bash
sing-box version
```

确认 `Tags:` 行包含所需标签（例如 `with_quic`）。

## 如何同步上游

参阅仓库根目录文件 `UPSTREAM_MERGE.md`。
