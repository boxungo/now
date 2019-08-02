# 当前时间

> 系统配置的时区的当前时间

```golang
import "github.com/boxungo/now"

log.Debug(fmt.Sprintf("Timestamp: %v", now.Timestamp()))
log.Info(fmt.Sprintf("Microsecond: %v", now.Microsecond()))
log.Warn(fmt.Sprintf("Date: %v", now.Date()))
log.Error(fmt.Sprintf("Datetime: %v", now.Datetime()))
log.Fatal(fmt.Sprintf("String: %v", now.String()))
log.Fatal(fmt.Sprintf("Time: %v", now.Time()))
```

```text
2019-08-02 16:10:45.786418 [DEBUG] Timestamp: 1564733445
2019-08-02 16:10:45.786556 [INFO] Microsecond: 1564733445786553
2019-08-02 16:10:45.786562 [WARN] Date: 2019-08-02
2019-08-02 16:10:45.786566 [ERROR] Datetime: 2019-08-02 16:10:45
2019-08-02 16:10:45.786579 [FATAL] String: 2019-08-02 16:10:45.786576
2019-08-02 16:10:45.786598 [FATAL] Time: 2019-08-02 16:10:45.786581 +0800 CST
```