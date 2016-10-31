package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"strconv"
	"time"
)

//用户
type User struct {
	Id       int64
	Username string
	Password string
	UserId   string
	Token    string
	Time     int64
}

type BeerVideo struct {
	Id    int64
	Title string
	Url   string
	Time  int64
}

/******admin******/
type AdminUser struct {
	Id       int64
	Username string
	Password string
	Time     int64
}

/******admin******/
func RegisterDB() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:@/niang?charset=utf8")
	beego.Debug("root:@/niang?charset=utf8")
	// register model
	orm.RegisterModel(new(User))
	orm.RegisterModel(new(BeerVideo))

	orm.RegisterModel(new(AdminUser))

	// create table
	orm.RunSyncdb("default", false, true)
}

//注册用户
func AddUser(username string, password string) (string, error) {
	o := orm.NewOrm()
	create_time := time.Now().Unix()
	userid, err := GetUserId()
	if err != nil {
		return "", err
	}
	token, err := GetToken(username, password)
	if err != nil {
		return "", err
	}
	obj := &User{Username: username, Password: password, Time: create_time, UserId: userid, Token: token}
	// 插入数据
	_, err = o.Insert(obj)
	if err != nil {
		return "", err
	}
	return token, nil
}

func GetUser(username string) (*User, error) {
	o := orm.NewOrm()
	var objs []User
	_, err := o.Raw("SELECT * FROM user WHERE username = ? ", username).QueryRows(&objs)
	obj := &User{}
	if len(objs) > 0 {
		obj = &objs[0]
	}
	return obj, err
}
func GetUserCount() (int32, error) {
	o := orm.NewOrm()
	count, err := o.QueryTable("user").Count()
	return int32(count), err
}

func GetUserId() (string, error) {
	count, err := GetUserCount()
	if err != nil {
		return "", err
	}
	count = count + 1
	idStr := strconv.Itoa(int(count))
	s := ""
	max := 8
	for i := 0; i < max-len(idStr); i++ {
		s = s + "0"
	}
	s = s + idStr
	return s, nil
}
func GetToken(username string, password string) (string, error) {
	token := ""
	token = username + password
	return token, nil
}

func GetUserInfo(token string) (*User, error) {
	o := orm.NewOrm()
	var objs []User
	_, err := o.Raw("SELECT * FROM user WHERE token = ? ", token).QueryRows(&objs)
	obj := &User{}
	if len(objs) > 0 {
		obj = &objs[0]
	}
	return obj, err
}

//添加啤酒视频
func AddBeerVideo(title string, url string) error {
	o := orm.NewOrm()
	create_time := time.Now().Unix()
	obj := &BeerVideo{Title: title, Url: url, Time: create_time}
	// 插入数据
	_, err := o.Insert(obj)
	if err != nil {
		return err
	}
	return nil
}

func GetAllBeerVideo() ([]BeerVideo, error) {
	o := orm.NewOrm()
	var objs []BeerVideo
	_, err := o.Raw("SELECT * FROM beer_video  ORDER BY id DESC").QueryRows(&objs)
	return objs, err
}
func DelBeerVideo(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &BeerVideo{Id: cid}
	_, err = o.Delete(cate)
	return err
}

/******admin******/
func AddAdminUser(username string, password string) error {
	o := orm.NewOrm()
	create_time := time.Now().Unix()
	obj := &AdminUser{Username: username, Password: password, Time: create_time}
	// 插入数据
	_, err := o.Insert(obj)
	if err != nil {
		return err
	}
	return nil
}
func GetAdminUser(username string) (*User, error) {
	o := orm.NewOrm()
	var objs []User
	_, err := o.Raw("SELECT * FROM admin_user WHERE username = ? ", username).QueryRows(&objs)
	obj := &User{}
	if len(objs) > 0 {
		obj = &objs[0]
	}
	return obj, err
}

/******admin******/
