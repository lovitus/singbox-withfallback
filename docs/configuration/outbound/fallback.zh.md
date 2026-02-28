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

按优先级顺序排列的出站标签列表。

Fallback 总是优先选择列表中最靠前且可用的出站。

#### url

健康检查使用的 URL。
为空时默认使用 `https://www.gstatic.com/generate_204`。

#### interval

健康检查间隔。
为空时默认 `3m`。

#### idle_timeout

定时健康检查的空闲超时。
策略组长时间无流量时会暂停定时检查，有新流量再恢复。
为空时默认 `30m`。

#### failure_threshold

将当前可用节点标记为不可用前，需要连续失败的次数。
为空或 `0` 时默认 `3`。

#### success_threshold

将不可用节点标记为恢复前，需要连续成功的次数。
为空或 `0` 时默认 `3`。

#### interrupt_exist_connections

当选定出站发生变化时，中断现有连接。

仅入站连接受此设置影响，内部连接将始终被中断。

### 行为说明

1. 策略组会对配置中的所有出站维护可用性状态。
2. 选择顺序为“按配置顺序取第一个可用出站”。
3. 当前出站达到 `failure_threshold` 后会被标记为不可用，并切换到下一个可用出站。
4. 更高优先级出站达到 `success_threshold` 后会被标记为恢复，并自动回切。
5. 不仅定时 URL 检查失败会计入失败次数，运行时的 Dial/Listen 失败也会计入失败次数。
6. 如果当前全部标记为不可用，仍会退化为选择第一个支持目标网络的出站进行兜底尝试。

### 调优建议

- `3`（默认）: 切换和回切更快。
- `10`: 更保守，适合链路波动较大、想减少抖动切换的场景。
