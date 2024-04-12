package test

import (
	"fmt"
	"getcharzp.cn/helper"
	"getcharzp.cn/models"
	"testing"
)

func TestDBConn(t *testing.T) {
	data := new(models.UserBasic)
	err := models.DB.Where("name = ? AND password = ? ", "hhf", helper.GetMd5("123456")).First(&data).Error
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data.Name)
}
