> Sponsored by [Warp](https://go.warp.dev/sing-box), built for coding with multiple AI agents

<a href="https://go.warp.dev/sing-box">
<img alt="Warp sponsorship" width="400" src="https://github.com/warpdotdev/brand-assets/raw/refs/heads/main/Github/Sponsor/Warp-Github-LG-02.png">
</a>

---

# sing-box

The universal proxy platform (forked edition).

[![Packaging status](https://repology.org/badge/vertical-allrepos/sing-box.svg)](https://repology.org/project/sing-box/versions)

## Upstream / Fork

- Upstream project: https://github.com/SagerNet/sing-box
- This fork: https://github.com/lovitus/singbox-withfallback

## What Is Different In This Fork

1. Added a new outbound group type: `fallback`.
2. `fallback` supports ordered failover and automatic fallback-to-primary recovery.
3. `fallback` supports configurable switch thresholds:
   - `failure_threshold` (default `3`)
   - `success_threshold` (default `3`)
4. Release binaries are built with upstream default full-feature tags:
   - `with_gvisor,with_quic,with_dhcp,with_wireguard,with_utls,with_acme,with_clash_api,with_tailscale,with_ccm,with_ocm,badlinkname,tfogo_checklinkname0`

## Fallback Parameters

`fallback` fields and behavior are documented in detail here:

- English: `docs/configuration/outbound/fallback.md`
- 中文: `docs/configuration/outbound/fallback.zh.md`

Quick guidance:

- Use default `3` for faster failover/recovery.
- Use `10` if your network is noisy and you want fewer switch oscillations.

## Verify Build Features

Run:

```bash
sing-box version
```

Check that `Tags:` includes at least `with_quic` and other required feature tags.

## Documentation

https://sing-box.sagernet.org

Fork notes:

- English: `docs/fork.md`
- 中文: `docs/fork.zh.md`

## Upstream Merge Guide

See:

- `UPSTREAM_MERGE.md`

## License

```
Copyright (C) 2022 by nekohasekai <contact-sagernet@sekai.icu>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.

In addition, no derivative work may use the name or imply association
with this application without prior consent.
```
