---
description: Fork-specific notes for singbox-withfallback.
---

# Fork Notes

This repository is a fork of `SagerNet/sing-box`.

## Repositories

- Upstream: https://github.com/SagerNet/sing-box
- Fork: https://github.com/lovitus/singbox-withfallback

## Main Differences From Upstream

1. Added a new outbound group type: `fallback`.
2. `fallback` checks outbounds in configured order and always prefers higher-priority recovered nodes.
3. `fallback` has threshold-based switching:
   - `failure_threshold`: consecutive failures to mark current outbound unavailable (default `3`)
   - `success_threshold`: consecutive successes to mark unavailable outbound recovered (default `3`)
4. Fork release binaries are built with upstream default feature tags, including QUIC.

## Fallback Documentation

- English: [/configuration/outbound/fallback/](/configuration/outbound/fallback/)
- 中文: [/zh/configuration/outbound/fallback/](/zh/configuration/outbound/fallback/)

## Build Feature Tags Used In Fork Releases

```text
with_gvisor,with_quic,with_dhcp,with_wireguard,with_utls,with_acme,with_clash_api,with_tailscale,with_ccm,with_ocm,badlinkname,tfogo_checklinkname0
```

## How To Verify A Binary

Run:

```bash
sing-box version
```

Confirm the `Tags:` line contains the required feature tags (for example `with_quic`).

## Sync With Upstream

See repository file `UPSTREAM_MERGE.md` for merge/rebase workflow.
