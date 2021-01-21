/*
	@Title
	api/verify_code.go

	@Description
	生成验证码给邮件

	@Func List（这个需打开函数检查）

	| func name            | develop  | unit test |

	|---------------------------------------------|

	| RandVerificationCode |    yes   |    yes	  |
*/
package verification

import (
	"math/rand"
	"strconv"
	"time"
)

/*
@Title
RandVerifyCode

@description
生成验证码并储存进redis数据库中

@param
无

@return
返回随机后的数字 string
*/
func RandVerificationCode() string {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(899999) + 100000
	time.Sleep(time.Second)
	return strconv.Itoa(randNum)
}
