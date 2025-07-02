package flags

import (
	"FastFiber_v2/global"
	"FastFiber_v2/models"
	"FastFiber_v2/utils/pwd"
	"fmt"
	"go.uber.org/zap"
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

type User struct {
}

func (User) Create() {
	var user models.UserModel
	fmt.Println("请选择角色 1 管理员 2 普通用户:")
	_, err := fmt.Scanln(&user.RoleID)
	if err != nil {
		fmt.Println("输入错误", err)
		return
	}
	if user.RoleID != 1 && user.RoleID != 2 {
		fmt.Println("用户角色输入错误", err)
		return
	}
	fmt.Println("请输入用户名:")
	fmt.Scanln(&user.Username)

	var u models.UserModel
	err = global.DB.Take(&u, "username = ?", user.Username).Error
	if err == nil {
		fmt.Println("用户名已存在")
		return
	}

	fmt.Println("请输入密码:")
	password, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("读取密码错误", err)
		return
	}

	fmt.Println("请再次输入密码:")
	rePassword, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("读取密码错误", err)
		return
	}
	if string(password) != string(rePassword) {
		fmt.Println("两次密码输入不一致")
		return
	}

	hashPwd, _ := pwd.GenerateFromPassword(string(password))
	err = global.DB.Create(&models.UserModel{
		Username: user.Username,
		Password: hashPwd,
		RoleID:   user.RoleID,
	}).Error
	if err != nil {
		zap.L().Error("用户创建失败", zap.Error(err))
		return
	}
	zap.L().Info("用户创建成功")
}

func (User) List() {
	var userList []models.UserModel
	global.DB.Order("created_at desc").Limit(10).Find(&userList)
	for _, model := range userList {
		fmt.Printf("用户id: %d 用户名: %s 用户角色: %d \n",
			model.ID,
			model.Username,
			model.RoleID,
		)
	}
}
