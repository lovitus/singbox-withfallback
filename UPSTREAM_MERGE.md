# Upstream Merge Guide

This repository is a fork of `SagerNet/sing-box` with custom fallback outbound support.

## 1. Ensure remotes

```bash
git remote -v
```

Expected:

- `origin` -> your fork (`lovitus/singbox-withfallback`)
- `upstream` -> original project (`SagerNet/sing-box`)

If `upstream` is missing:

```bash
git remote add upstream https://github.com/SagerNet/sing-box.git
```

## 2. Fetch latest upstream

```bash
git fetch upstream --tags
```

## 3. Rebase your main branch onto upstream

```bash
git checkout main
git pull origin main
git rebase upstream/dev-next
```

Why `dev-next`: this fork was based on upstream `dev-next`.

## 4. Resolve conflicts (if any)

Likely conflict files for fallback feature:

- `option/group.go`
- `constant/proxy.go`
- `include/registry.go`
- `protocol/group/fallback.go`
- `protocol/group/fallback_test.go`
- docs files under `docs/configuration/outbound/`

After resolving:

```bash
git add <resolved files>
git rebase --continue
```

Repeat until rebase finishes.

## 5. Validate build and tests

```bash
go test ./protocol/group ./include ./option ./constant
```

For a quick binary check with upstream default feature tags:

```bash
go build -trimpath -tags "with_gvisor,with_quic,with_dhcp,with_wireguard,with_utls,with_acme,with_clash_api,with_tailscale,with_ccm,with_ocm,badlinkname,tfogo_checklinkname0" -ldflags "-X github.com/sagernet/sing-box/constant.Version=dev -X internal/godebug.defaultGODEBUG=multipathtcp=0 -s -w -buildid= -checklinkname=0" ./cmd/sing-box
```

## 6. Push updated branch

```bash
git push origin main --force-with-lease
```

Use `--force-with-lease` because rebase rewrites history.

## 7. Re-run release workflow

In GitHub Actions, run `Release All Platforms` and pass a new version string, for example:

- `1.13.0-fallback.4`

## Optional alternative: merge strategy (no history rewrite)

If you prefer not to rebase:

```bash
git checkout main
git pull origin main
git merge upstream/dev-next
git push origin main
```

This keeps merge commits and avoids forced push.
