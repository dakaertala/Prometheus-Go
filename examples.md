# Examples

To see the percentage of total requests resulting in an exception:

```PromQL
rate(hello_world_exceptions_total[1m])/rate(hello_world_total[1m]) * 100
```

To see an average rate of change over 1 minute for a counter:

```PromQL
rate(hello_world_total[1m])
```

This can be used to determine when the last Hello World was served. The main use case for such a metric is detecting if it has been too long since a request was handled:

```PromQL
hello_world_last_time_seconds
```

This will tell you how many seconds it is since the last request:

```PromQL
time() - hello_world_last_time_seconds
```


