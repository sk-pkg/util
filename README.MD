# 常用辅助方法

## Encryptor 加密解密
```go
// AES解密
// key的长度必须是16、24或者32字节，分别用于选择AES-128, AES-192, or AES-256
func AesDecrypt(encryptedData, key, iv []byte) ([]byte, error)

// AES加密
// key的长度必须是16、24或者32字节，分别用于选择AES-128, AES-192, or AES-256
func AesEncrypt(origData, key, iv []byte) ([]byte, error)

// 计算Hmac散列值
func Hmac(key, data string) string

// 计算指定字符串32位md5哈希
// 如果不指定字符串，则返回当前时间的md5哈希
func MD5(str string) string

// PKCS7填充加密
func PKCS7Padding(ciphertext []byte, blockSize int) []byte

// PKCS7填充解密
func PKCS7UnPadding(origData []byte) []byte

// 计算指定字符串sha1散列值
func Sha1(data string) string
```
## Date 时间日期相关
```go
// 日期转化为时间戳
func DealDateUnix(date string) int64

// 格式化时间(string to string)
func DealTimeFormat(date, format string) string

// 格式化时间(time to string)
func FormatTime(t time.Time, f ...string) string

// 获取指定时间段的起止时间
// 可选：
// today 今天
// yesterday 昨天
// lately7 最近7天
// lately30 最近30天
// month 本月
// year 本年
func GetDateStartAndEnd(dateString string) [2]string

// 获取传入的时间所在月份的第一天0点时间。
func GetFirstDateOfMonth(d time.Time) time.Time

// 获取传入的时间所在年份的第一天0点时间。
func GetFirstDateOfYear(d time.Time) time.Time

// 获取传入的时间所在月份的最后一天最晚点时间。
func GetLastDateOfMonth(d time.Time) time.Time

// 获取传入的时间所在年份的最后一天最晚点时间。
func GetLastDateOfYear(d time.Time) time.Time

// 获取某一天的最晚点时间。
func GetLastTime(d time.Time) time.Time

// 获取当前时间到当天23:59:59的剩余时间（秒）
func GetNowToDayEndSecond() int

// 获取某一天的0点时间
func GetZeroTime(d time.Time) time.Time
```
## String 字符处理相关
```go
// 生成随机base32编码字符串
func RandBase32Str(length int) string

// 随机字符串，包含 0~9 和 a~z - [g,i,j,l,o,p,q,y]
func RandLowStr(length int) string

// 随机字符串，包含 英文字母和数字
func RandUpStr(length int) string

// 拼接字符串
func SpliceStr(p ...string) string

// 判断字符串切片中是否存在某个字符串
func IsContainStr(items []string, item string) bool
```
## Network 网络相关
```go
// 获取客户端真实IP
func GetRealIP(c *gin.Context) string
```
