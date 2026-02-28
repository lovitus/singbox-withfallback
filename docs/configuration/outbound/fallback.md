### Structure

```json
{
  "type": "fallback",
  "tag": "auto",

  "outbounds": [
    "proxy-a",
    "proxy-b",
    "proxy-c"
  ],
  "url": "",
  "interval": "",
  "idle_timeout": "",
  "failure_threshold": 3,
  "success_threshold": 3,
  "interrupt_exist_connections": false
}
```

### Fields

#### outbounds

==Required==

List of outbound tags, ordered by priority.

Fallback always tries to use the first available outbound in this list.

#### url

The URL used for health checks.
`https://www.gstatic.com/generate_204` will be used if empty.

#### interval

Health check interval.
`3m` will be used if empty.

#### idle_timeout

Idle timeout for scheduled health checks.
When the group is idle, periodic checks are paused and resumed on new traffic.
`30m` will be used if empty.

#### failure_threshold

Consecutive failures required before marking a currently available outbound as unavailable.
`3` will be used if empty or `0`.

#### success_threshold

Consecutive successes required before marking an unavailable outbound as recovered.
`3` will be used if empty or `0`.

#### interrupt_exist_connections

Interrupt existing connections when selected outbound has changed.

Only inbound connections are affected by this setting, internal connections will always be interrupted.

### Behavior

1. The group checks all configured outbounds and tracks availability state.
2. It selects the first available outbound by configured order.
3. When the current outbound reaches `failure_threshold`, it is marked unavailable and switches to the next available outbound.
4. When a higher-priority outbound reaches `success_threshold`, it is marked recovered and selected again automatically.
5. Runtime dial/listen failures of the selected outbound also count as failures (not only scheduled URL checks).
6. If all outbounds are currently unavailable, it still picks the first outbound that supports the requested network as a best-effort fallback.

### Tuning Tips

- `3` (default): faster failover and faster recovery.
- `10`: more conservative switching, useful on unstable links to reduce flapping.
