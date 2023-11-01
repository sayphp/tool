# test 测试命令

```bash
# http server test
curl "http://127.0.0.1:8080/dtp/test/test"

curl --unix-socket "/say/github/tool/dtp/uds-call-go.sock" "http://locahost/dtp/test/test"
```