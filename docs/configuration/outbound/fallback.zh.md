### 结构

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

### 字段

#### outbounds

==必填==

按优先级顺序检查的出站标签列表。

Fallback 总是优先选择列表里最靠前且可用的出站。
当前出站不可用时，会切换到下一个可用出站。
高优先级出站恢复后，会自动切回。

#### url

用于测试的链接。默认使用 `https://www.gstatic.com/generate_204`。

#### interval

检查间隔。默认使用 `3m`。

#### idle_timeout

空闲超时。默认使用 `30m`。

#### failure_threshold

将当前出站标记为不可用前所需的连续失败次数。
默认使用 `3`。

#### success_threshold

将不可用出站标记为恢复前所需的连续成功次数。
默认使用 `3`。

#### interrupt_exist_connections

当选定的出站发生更改时，中断现有连接。

仅入站连接受此设置影响，内部连接将始终被中断。
