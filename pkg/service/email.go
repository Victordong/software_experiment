package service

//
//import (
//	"context"
//	"github.com/spf13/viper"
//	"gopkg.in/gomail.v2"
//	"software_experiment/pkg/comm/manager"
//	"software_experiment/pkg/comm/model"
//)
//
//func SendMail(ctx context.Context, address string, body string) bool {
//	m := gomail.NewMessage()
//	m.SetHeader("From", "1468767640@qq.com")
//	m.SetHeader("To", address)
//	m.SetHeader("Subject", "Hello!")
//	m.SetBody("text/html", body)
//
//	d := gomail.NewPlainDialer(
//		viper.GetString("change_password_email.smtp_host"),
//		viper.GetInt("change_password_email.smtp_port"),
//		viper.GetString("change_password_email.username"),
//		viper.GetString("change_password_email.password"))
//	if err := d.DialAndSend(m); err != nil {
//		return false
//	}
//	return true
//}
//
//type ShopMessageServiceClient struct {
//}
//
//func NewShopMessageClient() *ShopMessageServiceClient {
//	return &ShopMessageServiceClient{}
//}
//
//func (sms *ShopMessageServiceClient) SendMessage(ctx context.Context, shopMessage *model.ShopMessageModel) error {
//
//	shopMessage, err := manager.PostShopMessage(ctx, shopMessage)
//
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//type ShopMessageService struct {
//}
//
//func NewShopMessageServer() *ShopMessageService {
//	return &ShopMessageService{}
//}
//
//func (sms *ShopMessageService) Run() {
//	ctx := context.Background()
//
//	shopMessages, _, err := manager.QueryShopMessages(ctx, map[string][]string{
//		"status": []string{"unsent", "error"},
//	})
//	if err != nil {
//		return
//	}
//	var failMails []model.ShopMessageModel
//	for index, shopMessage := range shopMessages {
//
//		if shopMessage.Email == "all@all.com" {
//			// 全体消息 发邮件
//			shops, _, err := manager.QueryShops(ctx, map[string][]string{
//				"status":    []string{"close"},
//				"_per_page": []string{"100000"},
//			})
//			if err != nil {
//				continue
//			}
//			for _, shop := range shops {
//
//				ok := SendMail(ctx, shop.Email, shopMessage.Body)
//				if ok {
//					// 发送成功 修改状态
//					manager.PutShopMessage(ctx, shopMessage.ID, map[string]interface{}{
//						"status": "sent",
//					})
//				} else {
//					failMails = append(failMails, shopMessages[index])
//				}
//			}
//
//		} else {
//			ok := SendMail(ctx, shopMessage.Email, shopMessage.Body)
//			if ok {
//				// 发送成功 修改状态
//				manager.PutShopMessage(ctx, shopMessage.ID, map[string]interface{}{
//					"status": "unsent",
//				})
//			} else {
//				failMails = append(failMails, shopMessages[index])
//			}
//		}
//	}
//}
