package handlers

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func CreateDynamicChargeJD(c *gin.Context) {
	response := map[string]interface{}{
        "data": map[string]interface{}{
            "idDocumento":           "52c43361-caa1-4ddb-9152-708426a25db3",
            "imagemQrCodeInBase64": "SkQuUElYLlFSQ29kZS5EaW5hbWljbw==......",
            "payloadBase64" : `MDAwMjAxMjYzNjAwMTRCUi5HT1YuQkNCLlBJWTAxMTQrNTUxMTk5OTk5OTk5NTIwNDAwMDA1MzAzOTg2NTQwNjEwMC4wMDU4MDJCUjU5MTNGdWxhbm8gZGUgVGFsNjAwOVNhbyBQYXVsbzYyMTAwNTA2QUJDMTIzNjMwNEMxM0Q`,
            "payloadJws": `eyJhbGciOiJQUzUxMiIsImtpZCI6IjUwQTM5Qzc0MUE0RTFDQjQxN0Y2OEM2Q0MwMkY2M0JFODg2RDg1MzIiLCJ0eXAiOiJKV1QiLCJ4NXQiOiJVS09jZEJwT0hMUVg5b3hzd0M5anZvaHRoVEkiLCJqa3UiOiJodHRwczovL2V4ZW1wbG8uY29tL3BpeC9qZHBpL3FyYy9qd2tzIn0.eyJ2ZXJzYW8iOiIxLjAuMCIsImRvY3VtZW50byI6eyJpZCI6IjNmOTk0YTY2LTgyMzAtNGM5YS1hZDJhLTBmZmU0YjY5YWQ2OCIsInJldmlzYW8iOjB9LCJjYWxlbmRhcmlvIjp7ImNyaWFjYW8iOiIxMy8xMC8yMDIwIDE4OjQ0OjExIiwiYXByZXNlbnRhY2FvIjoiMTMvMTAvMjAyMCAxODo0NDoxMSIsImV4cGlyYWNhbyI6IjIzLzEyLzIwMjAgMjM6NTk6MDUiLCJ2ZW5jaW1lbnRvIjoiMjMvMDEvMjAyMCAwMDowMDowMCIsInJlY2ViaXZlbEFwb3NWZW5jaW1lbnRvIjp0cnVlfSwicGFnYWRvciI6eyJjbnBqIjoiNDM1Odc5ODAwMDEwNyIsIm5vbWUiOiJDaWNsYW5vIGRlIFRhbCJ9LCJ2YwxvciI6eyJvcmlnaW5hbCI6MTAwMC4wMSwiZmluYWwiOjExMDAuMDEsImp1cm9zIjoxMDAuMCwibXVsdGEiOjEwLjAsImRlc2NvbnRvIjoxMC4wLCJwZXJtaXRlQWx0ZXJhY2FvIjpmYWxzZX0sImNoYXZlIjoiZnVsYW5vLnRhbEBwcm92ZWRvci5jb20uYnIiLCJ0eGlkIjoiSkRQSTIwMjAwMTAzMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDEiLCJzb2xpY2l0YWNhb1BhZ2Fkb3IiOiJQZXNxdWlzYSBkZSBzYXRpc2fDp8OjbyIsImluZm9BZGljaW9uYWlzIjpbeyJub21lIjoiRGV0YWxoZXMgZG8gUGFnYW1lbnRvIiwidmFsb3IiOiJJbmZvcm1hw6fDoW8gQWRpY2lvbmFsIGRvIFBTUCBkbyBSZWNlYmVkb3IifSx7Im5vbWUiOiJEZXRhbGhlcyBkbyBQYWdhbWVudG8gMiIsInZhbG9yIjoiSW5mb3JtYcOnw6NvIEFkaWNpb25hbCBkbyBQU1AgZG8gUmVjZWJlZG9yIn1dLCJleHAiOjE2MDg3Nzg3NDUsIm5iZiI6MTYwMjYxNDY1NiwiaWF0IjoxNjAyNjE0NjU2fQ.Ml-RauOQlIqPxRf4sjnoMuJRzbFLcEaF4KQmP5Hm9ilCht89kXQCiGdNHQZ3cU0_Civf2zQsNwe3w98nDvjasw0XBmBkCCAAY46H4CbXk26qROxeE9wiOavBq-O47C5s-hlmg_fhbdJX7xdmTdtuHN6RWqCII0JzIGcTyiEwJxsKt5Qb1dYWFINnh9xhb3J9KsWJRwESK4jB-WyIIGLr7zVKot5lFUt7hTgU8c6QeQijwjwqmxMUF8z2h-y7dV0prCt7o-JpbdcdTXAL7CuTte8WfH7nGiljeiDRdEkY-neBvzBRpd84RMR1392bQ4gppaxHU7S-ZkakJ5hSHoQKzg`,
        },
    }
	c.JSON(http.StatusOK, response)
}

func UpdateDynamicChargeJD(c *gin.Context) {
	response := map[string]interface{}{
        "data": map[string]interface{}{
            "idDocumento":"52c43361-caa1-4ddb-9152-708426a25db3",
            "revisao":2,
            "imagemQrCodeInBase64":"SkQuUElYLlFSQ29kZS5EaW5hbWljbw==......",
            "payloadBase64" : `MDAwMjAxMjYzNjAwMTRCUi5HT1YuQkNCLlBJWTAxMTQrNTUxMTk5OTk5OTk5NTIwNDAwMDA1MzAzOTg2NTQwNjEwMC4wMDU4MDJCUjU5MTNGdWxhbm8gZGUgVGFsNjAwOVNhbyBQYXVsbzYyMTAwNTA2QUJDMTIzNjMwNEMxM0Q`,
            "payloadJws": `eyJhbGciOiJQUzUxMiIsImtpZCI6IjUwQTM5Qzc0MUE0RTFDQjQxN0Y2OEM2Q0MwMkY2M0JFODg2RDg1MzIi
LCJ0eXAiOiJKV1QiLCJ4NXQiOiJVS09jZEJwT0hMUVg5b3hzd0M5anZvaHRoVEkiLCJqa3UiOiJodHRwczovL2V4ZW1wbG8uY29t
L3BpeC9qZHBpL3FyYy9qd2tzIn0.eyJ2ZXJzYW8iOiIxLjAuMCIsImRvY3VtZW50byI6eyJpZCI6IjNmOTk0YTY2LTgyMzAtNGM5
YS1hZDJhLTBmZmU0YjY5YWQ2OCIsInJldmlzYW8iOjB9LCJjYWxlbmRhcmlvIjp7ImNyaWFjYW8iOiIxMy8xMC8yMDIwIDE4OjQ0
OjExIiwiYXByZXNlbnRhY2FvIjoiMTMvMTAvMjAyMCAxODo0NDoxMSIsImV4cGlyYWNhbyI6IjIzLzEyLzIwMjAgMjM6NTk6MDUi
LCJ2ZW5jaW1lbnRvIjoiMjMvMDEvMjAyMCAwMDowMDowMCIsInJlY2ViaXZlbEFwb3NWZW5jaW1lbnRvIjp0cnVlfSwicGFnYWRv
ciI6eyJjbnBqIjoiNDM1Odc5ODAwMDEwNyIsIm5vbWUiOiJDaWNsYW5vIGRlIFRhbCJ9LCJ2YwxvciI6eyJvcmlnaW5hbCI6MTAw
MC4wMSwiZmluYWwiOjExMDAuMDEsImp1cm9zIjoxMDAuMCwibXVsdGEiOjEwLjAsImRlc2NvbnRvIjoxMC4wLCJwZXJtaXRlQWx0
ZXJhY2FvIjpmYWxzZX0sImNoYXZlIjoiZnVsYW5vLnRhbEBwcm92ZWRvci5jb20uYnIiLCJ0eGlkIjoiSkRQSTIwMjAwMTAzMDAw
MDAwMDAwMDAwMDAwMDAwMDAwMDEiLCJzb2xpY2l0YWNhb1BhZ2Fkb3IiOiJQZXNxdWlzYSBkZSBzYXRpc2Zhw6fDo28iLCJpbmZv
QWRpY2lvbmFpcyI6W3sibm9tZSI6IkRldGFsaGVzIGRvIFBhZ2FtZW50byIsInZhbG9yIjoiSW5mb3JtYcOnw6NvIEFkaWNpb25h
bCBkbyBQU1AgZG8gUmVjZWJlZG9yIn0seyJub21lIjoiRGV0YWxoZXMgZG8gUGFnYW1lbnRvIDIiLCJ2YwxvciI6IkluZm9ybWHD
p8OjbyBBZGljaW9uYWwgZG8gUFNQIGRvIFJlY2ViZWRvciJ9XswiZXhwIjoxNjA4Nzc4NzQ1LCJuYmYiOjE2MDI2MTQ2NTYsImlh
dCI6MTYwMjYxNDY1Nn0.Ml-RauOQlIqPxRf4sjnoMuJRzbFLcEaF4KQmP5Hm9ilCht89kXQCiGdNHQZ3cU0_Civf2zQsNwe3w98nDvjasw0XBmBkCCAAY46H4Cb
Xk26qROxeE9wiOavBq-O47C5s-
hlmg_fhbdJX7xdmTdtuHN6RWqCII0JzIGcTyiEwJxsKt5Qb1dYWFINnh9xhb3J9KsWJRwESK4jB-
WyIIGLr7zVKot5lFUt7hTgU8c6QeQijwjwqmxMUF8z2h-y7dV0prCt7o-JpbdcdTXAL7CuTte8WfH7nGiljeiDRdEkY-
neBvzBRpd84RMR1392bQ4gppaxHU7S-ZkakJ5hSHoQKzg`,
        },
    }
	c.JSON(http.StatusOK, response)
}

func CreateDynamicDueChargeJD(c *gin.Context) {
	response := map[string]interface{}{
        "data": map[string]interface{}{
            "idDocumento":"52c43361-caa1-4ddb-9152-708426a25db3",
            "revisao":2,
            "payloadBase64" : `MDAwMjAxMjYzNjAwMTRCUi5HT1YuQkNCLlBJWTAxMTQrNTUxMTk5OTk5OTk5NTIwNDAwMDA1MzAzOTg2NTQwNjEwMC4wMDU4MDJCUjU5MTNGdWxhbm8gZGUgVGFsNjAwOVNhbyBQYXVsbzYyMTAwNTA2QUJDMTIzNjMwNEMxM0Q`,
            "imagemQrCodeInBase64":"SkQuUElYLlFSQ29kZS5EaW5hbWljbw==......",
        },
    }
	c.JSON(http.StatusOK, response)
}

func UpdateDynamicDueChargeJD(c *gin.Context) {
	response := map[string]interface{}{
        "data": map[string]interface{}{
            "idDocumento":"52c43361-caa1-4ddb-9152-708426a25db3",
            "revisao":2,
            "payloadBase64" : `MDAwMjAxMjYzNjAwMTRCUi5HT1YuQkNCLlBJWTAxMTQrNTUxMTk5OTk5OTk5NTIwNDAwMDA1MzAzOTg2NTQwNjEwMC4wMDU4MDJCUjU5MTNGdWxhbm8gZGUgVGFsNjAwOVNhbyBQYXVsbzYyMTAwNTA2QUJDMTIzNjMwNEMxM0Q`,
            "imagemQrCodeInBase64":"SkQuUElYLlFSQ29kZS5EaW5hbWljbw==......",
        },
    }
	c.JSON(http.StatusOK, response)
}
