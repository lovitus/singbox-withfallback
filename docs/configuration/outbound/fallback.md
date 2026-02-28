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

List of outbound tags to check in order of priority.

Fallback always prefers the earliest available outbound in this list.
When the current outbound becomes unavailable, it switches to the next available one.
When a higher-priority outbound recovers, it switches back automatically.

#### url

The URL to test. `https://www.gstatic.com/generate_204` will be used if empty.

#### interval

The check interval. `3m` will be used if empty.

#### idle_timeout

The idle timeout. `30m` will be used if empty.

#### failure_threshold

Required consecutive failed checks before marking the current outbound unavailable.
`3` will be used if empty.

#### success_threshold

Required consecutive successful checks before marking an unavailable outbound recovered.
`3` will be used if empty.

#### interrupt_exist_connections

Interrupt existing connections when the selected outbound has changed.

Only inbound connections are affected by this setting, internal connections will always be interrupted.
