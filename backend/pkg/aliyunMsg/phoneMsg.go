package aliyunMsg

import (
	"Skywing/settings"
	"Skywing/store/redis"
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"go.uber.org/zap"
)

/*CreateClient
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func createClient(accessKeyId *string, accessKeySecret *string) (_result *dysmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		// 您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}

func SendVerificationCode(phone, code string) error {
	client, err := createClient(tea.String(settings.Conf.AliyunMsgConf.AccessKeyId), tea.String(settings.Conf.AliyunMsgConf.AccessKeySecret))
	if err != nil {
		return err
	}
	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		SignName:      tea.String("SKYLAB"),
		TemplateCode:  tea.String("SMS_154950909"),
		PhoneNumbers:  tea.String(phone),
		TemplateParam: tea.String(fmt.Sprintf(`{"code":"%s"}`, code)),
	}
	runtime := &util.RuntimeOptions{}
	tryErr := func() (e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				e = r
			}
		}()
		_, err = client.SendSmsWithOptions(sendSmsRequest, runtime)
		if err != nil {
			return err
		}
		return nil
	}()

	if tryErr != nil {
		var e = &tea.SDKError{}
		if t, ok := tryErr.(*tea.SDKError); ok {
			e = t
		} else {
			e.Message = tea.String(tryErr.Error())
		}
		_, err = util.AssertAsString(e.Message)
		if err != nil {
			zap.L().Error("发送短信失败！", zap.Error(err))
			return err
		}
	}
	return err
}
func GetVerificationCode(phone string) (string, error) {
	randCode := redis.RdbClient.Client.Get(phone).String()
	if err := redis.RdbClient.Client.Del(phone).Err(); err != nil {
		return "", err
	}
	return randCode, nil
}
