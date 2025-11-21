package controllers

import (
	"backend/g"
	"backend/modules/base/models"
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/nfnt/resize"
	"github.com/sirupsen/logrus"
)

type User struct {
	*g.X
}

func NewUser() *User {
	u := &User{X: g.NewX(&models.User{})}
	u.Tool = []*g.Tool{}
	u.Option = [][]any{}
	return u
}

func (u *User) ActionEdit(c *gin.Context) {
	u.RenderData(c, gin.H{
		"user":  u.GetUser(c).User,
		"rules": u.GetRules(),
	})
}

func (u *User) ActionSet(c *gin.Context) {
	key := c.Query("key")
	val := c.Query("val")
	if key != "fold" && key != "page_size" {
		u.JsonFail(c, fmt.Errorf("key错误"))
		return
	}
	sql := "UPDATE user SET " + key + " = ? WHERE username = ?"
	_, err := u.DB.Exec(sql, val, u.GetUsername(c))
	if err != nil {
		u.JsonFail(c, err)
		return
	}
	u.JsonSucc(c, "设置成功")
}

func (u *User) ActionSave(c *gin.Context) {
	username := u.GetUsername(c)
	user := &models.User{}
	if err := c.ShouldBindJSON(user); err != nil {
		u.JsonFail(c, err)
		return
	}
	if username != user.Username {
		u.JsonFail(c, fmt.Errorf("无权限修改其他用户信息"))
		return
	}
	old := new(models.User)
	if has, err := u.DB.Where("username = ?", username).Get(old); err != nil || !has {
		u.JsonFail(c, fmt.Errorf("用户不存在"))
		return
	}
	old.Email = user.Email
	old.CnName = user.CnName
	old.Fold = user.Fold
	old.PageSize = user.PageSize
	sess := u.BeginSess(u.DB, c)
	if err := old.Save(sess); err != nil {
		sess.Rollback()
		u.JsonFail(c, err)
		return
	}
	sess.Commit()
	u.JsonSucc(c, "保存成功")
}

func (u *User) ActionGetAvatar(c *gin.Context) {
	username := u.GetUsername(c)
	user := &models.User{}
	has, err := u.DB.Where("username = ?", username).Get(user)
	if err != nil || !has || len(user.Avatar) == 0 {
		u.JsonFail(c, fmt.Errorf("获取头像失败: %v", err))
		return
	}
	u.JsonSucc(c, base64.StdEncoding.EncodeToString(user.Avatar))
}

func (u *User) ActionUploadAvatar(c *gin.Context) {
	username := u.GetUsername(c)

	var req struct {
		Image string `json:"image"` // base64编码的图片数据
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		u.JsonFail(c, fmt.Errorf("参数错误: %v", err))
		return
	}

	// 解码base64图片
	imageData, err := base64.StdEncoding.DecodeString(req.Image)
	if err != nil {
		u.JsonFail(c, fmt.Errorf("图片格式错误: %v", err))
		return
	}

	// 解码图片
	img, _, err := image.Decode(bytes.NewReader(imageData))
	if err != nil {
		u.JsonFail(c, fmt.Errorf("图片解码失败: %v", err))
		return
	}

	// 转换为JPEG格式并压缩
	compressedData, err := compressImage(img, 64*1024) // 64KB
	if err != nil {
		u.JsonFail(c, fmt.Errorf("图片压缩失败: %v", err))
		return
	}

	// 更新用户头像
	old := new(models.User)
	if has, err := u.DB.Where("username = ?", username).Get(old); err != nil || !has {
		u.JsonFail(c, fmt.Errorf("用户不存在"))
		return
	}

	logrus.Infof("准备更新用户头像: %s, 压缩后大小: %d bytes", username, len(compressedData))
	old.Avatar = compressedData
	sess := u.BeginSess(u.DB, c)
	if err := old.Save(sess); err != nil {
		sess.Rollback()
		logrus.Errorf("更新头像失败: %v", err)
		u.JsonFail(c, fmt.Errorf("更新头像失败: %v", err))
		return
	}
	sess.Commit()

	// 验证保存是否成功
	verify := new(models.User)
	if has, err := u.DB.Where("username = ?", username).Get(verify); err == nil && has {
		logrus.Infof("头像保存验证: %s, 数据库中的大小: %d bytes", username, len(verify.Avatar))
	}

	u.JsonSucc(c, "头像上传成功")
}

func compressImage(img image.Image, maxSize int) ([]byte, error) {
	// 先调整大小到合理尺寸（最大800x800）
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	var resized image.Image
	if width > 800 || height > 800 {
		if width > height {
			resized = resize.Resize(800, 0, img, resize.Lanczos3)
		} else {
			resized = resize.Resize(0, 800, img, resize.Lanczos3)
		}
	} else {
		resized = img
	}

	// 压缩质量从90开始，逐步降低直到满足大小要求
	quality := 90
	var buf bytes.Buffer

	for quality >= 10 {
		buf.Reset()
		err := jpeg.Encode(&buf, resized, &jpeg.Options{Quality: quality})
		if err != nil {
			return nil, err
		}

		if buf.Len() <= maxSize {
			return buf.Bytes(), nil
		}

		quality -= 10
	}

	// 如果还是太大，进一步缩小尺寸
	if buf.Len() > maxSize {
		resized = resize.Resize(400, 0, resized, resize.Lanczos3)
		buf.Reset()
		err := jpeg.Encode(&buf, resized, &jpeg.Options{Quality: 70})
		if err != nil {
			return nil, err
		}
	}

	return buf.Bytes(), nil
}

func (u *User) ActionJoin(c *gin.Context) {
	u.RenderData(c, gin.H{"path": "/base/user/edit"})
}

type joinArg struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *User) loadJoinArg(c *gin.Context) (*joinArg, error) {
	var arg = &joinArg{}
	if err := c.ShouldBindJSON(arg); err != nil {
		return nil, err
	}
	if arg.Username == "" || arg.Password == "" {
		return nil, fmt.Errorf("用户名和密码不能为空")
	}
	return arg, nil
}

func (u *User) writeSession(c *gin.Context, username string) {
	session := sessions.Default(c)
	session.Set("username", username)
	if err := session.Save(); err != nil {
		logrus.Errorf("保存 session 失败: %v", err)
		u.JsonFail(c, fmt.Errorf("保存 session 失败: %v", err))
		return
	}
}

func (u *User) ActionLogIn(c *gin.Context) {
	arg, err := u.loadJoinArg(c)
	if err != nil {
		u.JsonFail(c, err)
		return
	}
	user := &models.User{}
	has, err := u.DB.Where("username = ?", arg.Username).Get(user)
	if err != nil || !has {
		u.JsonFail(c, fmt.Errorf("用户[%s]不存在", arg.Username))
		return
	}
	if user.Password != arg.Password {
		u.JsonFail(c, fmt.Errorf("密码错误"))
		return
	}

	u.writeSession(c, arg.Username)
	u.JsonSucc(c, "登录成功")
}

func (u *User) ActionSignUp(c *gin.Context) {
	arg, err := u.loadJoinArg(c)
	if err != nil {
		u.JsonFail(c, err)
		return
	}
	existingUser := &models.User{}
	has, _ := u.DB.Where("username = ?", arg.Username).Get(existingUser)
	if has {
		u.JsonFail(c, fmt.Errorf("用户名已存在"))
		return
	}
	newUser := &models.User{
		Username: arg.Username,
		Password: arg.Password,
		PageSize: 10,
	}
	sess := u.BeginSess(u.DB, c)
	if err := newUser.Save(sess); err != nil {
		sess.Rollback()
		u.JsonFail(c, fmt.Errorf("注册失败: %v", err))
		return
	}
	sess.Commit()

	u.writeSession(c, arg.Username)
	u.JsonSucc(c, "注册成功，已自动登录")
}
