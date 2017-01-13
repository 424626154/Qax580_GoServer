package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"github.com/satori/go.uuid"
	"math/rand"
	"strconv"
	"time"
)

func RegisterDB() {
	// set default database
	isdebug := "true"
	iniconf, err := config.NewConfig("json", "conf/config.json")
	if err != nil {
		beego.Debug(err)
	} else {
		isdebug = iniconf.String("beernotes::isdebug")
	}
	beego.Debug("register db isdebug:", isdebug)
	if isdebug == "true" {
		orm.RegisterDataBase("default", "mysql", "root:@/bndb?charset=utf8")
		beego.Debug("root:@/bndb?charset=utf8")
	} else {
		orm.RegisterDataBase("default", "mysql", "root:sbb890503@/bndb?charset=utf8")
		beego.Debug("root:sbb890503@/bndb?charset=utf8")
	}
	// register model
	orm.RegisterModel(new(Admin))
	orm.RegisterModel(new(Related))
	orm.RegisterModel(new(Token))
	orm.RegisterModel(new(Push))
	orm.RegisterModel(new(User))
	// create table
	orm.RunSyncdb("default", false, true)
}

type ResponseJson struct {
	Errcode int32  `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Data    string `json:"data"`
}

type Admin struct {
	Id       int64
	Username string `orm:"size(100)"`
	Password string `orm:"size(1000)"`
}

type User struct {
	Id       int64
	Username string `orm:"size(500)"`
	Password string `orm:"size(500)"`
	Email    string `orm:"size(500)"`
	Token    string `orm:"size(500)"`
	Secret   string `orm:"size(100)"`
	Uid      string `orm:"size(100)"`
}

type Related struct {
	Id    int64
	Name  string `orm:"size(100)"`
	Brief string `orm:"size(1000)"`
	Link  string `orm:"size(500)"`
	State int16
	Time  int64
}

type Token struct {
	Id     int64
	Token  string `orm:"size(500)"`
	Ostype string
	Time   int64
}

type Push struct {
	Id      int64
	Title   string `orm:"size(500)"`
	Content string `orm:"size(2000)"`
	Tokens  string `orm:"size(2000)"`
	Ostype  string
	Time    int64
}

func GetOneAdmin(username string) (*Admin, error) {
	o := orm.NewOrm()
	var admins []Admin
	_, err := o.Raw("SELECT * FROM admin WHERE username = ? ", username).QueryRows(&admins)
	admin := &Admin{}
	if len(admins) > 0 {
		admin = &admins[0]
	}
	return admin, err
}

func GetAllRelated() ([]Related, error) {
	o := orm.NewOrm()
	var objs []Related
	_, err := o.Raw("SELECT * FROM related  ORDER BY id DESC").QueryRows(&objs)
	return objs, err
}

func GetAllRelatedState(state int16) ([]Related, error) {
	o := orm.NewOrm()
	var objs []Related
	_, err := o.Raw("SELECT * FROM related WHERE state = ? ORDER BY id DESC", state).QueryRows(&objs)
	return objs, err
}

func AddRelated(name string, brief string, link string) error {
	o := orm.NewOrm()
	my_time := time.Now().Unix()
	cate := &Related{Name: name, Brief: brief, Link: link, Time: my_time}
	// 插入数据
	_, err := o.Insert(cate)
	if err != nil {
		return err
	}
	return nil
}

func UpdateRelatedState(id string, state int16) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &Related{Id: cid}
	cate.State = state
	_, err = o.Update(cate, "state")
	return err
}

func DeleteRelated(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &Related{Id: cid}
	_, err = o.Delete(cate)
	return err
}

func GetOneRelated(id string) (*Related, error) {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	var objes []Related
	_, err = o.Raw("SELECT * FROM related WHERE id = ? ", cid).QueryRows(&objes)
	if err != nil {
		return nil, err
	}
	obj := &Related{}
	if len(objes) > 0 {
		obj = &objes[0]
	}
	return obj, err
}

func UpdateRelated(id string, name string, brief string, link string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	my_time := time.Now().Unix()
	o := orm.NewOrm()
	cate := &Related{Id: cid}
	cate.Name = name
	cate.Brief = brief
	cate.Link = link
	cate.Time = my_time
	_, err = o.Update(cate, "name", "brief", "link", "time")
	return err
}

func SaveToken(token string, ostype string) error {
	o := orm.NewOrm()
	my_time := time.Now().Unix()
	var objs []Token
	_, err := o.Raw("SELECT * FROM token WHERE token = ? ", token).QueryRows(&objs)
	if len(objs) > 0 {
		cid := objs[0].Id
		obj := &Token{Id: cid}
		obj.Ostype = ostype
		obj.Time = my_time
		_, err = o.Update(obj, "ostype", "time")
	} else {
		cate := &Token{Token: token, Ostype: ostype, Time: my_time}
		// 插入数据
		_, err = o.Insert(cate)
	}
	return err
}

func GetOstypeAllTokesn(ostype string) ([]Token, error) {
	o := orm.NewOrm()
	var objs []Token
	_, err := o.Raw("SELECT * FROM token WHERE ostype = ? ORDER BY id DESC", ostype).QueryRows(&objs)
	return objs, err
}

func AddPush(title string, content string, tokens string, ostype string) error {
	o := orm.NewOrm()
	my_time := time.Now().Unix()
	cate := &Push{Title: title, Content: content, Tokens: tokens, Ostype: ostype, Time: my_time}
	// 插入数据
	_, err := o.Insert(cate)
	if err != nil {
		return err
	}
	return nil
}

func GetAllPush() ([]Push, error) {
	o := orm.NewOrm()
	var objs []Push
	_, err := o.Raw("SELECT * FROM push ORDER BY id DESC").QueryRows(&objs)
	return objs, err
}

func RegisterUser(username string, password string, email string) (int64, *User, error) {
	o := orm.NewOrm()
	var objs []User
	_, err := o.Raw("SELECT * FROM user WHERE username = ? ORDER BY id DESC", username).QueryRows(&objs)
	if err != nil {
		return 1006, nil, err
	}
	if len(objs) > 0 {
		return 1004, nil, err
	}
	_, err = o.Raw("SELECT * FROM user WHERE email = ? ORDER BY id DESC", email).QueryRows(&objs)
	if err != nil {
		return 1006, nil, err
	}
	if len(objs) > 0 {
		return 1005, nil, err
	}
	uid := uuid.NewV4().String()
	token := getUserToken(uid)
	secret := getUserSecret(uid)
	obj := &User{Username: username, Password: password, Email: email, Uid: uid, Token: token, Secret: secret}
	// 插入数据
	_, err = o.Insert(obj)
	if err != nil {
		return 1006, nil, err
	}
	return 0, obj, nil
}

func GetOneUser(username string) (*User, error) {
	o := orm.NewOrm()
	var objs []User
	_, err := o.Raw("SELECT * FROM user WHERE username = ? ORDER BY id DESC", username).QueryRows(&objs)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return &objs[0], err
	}
	return nil, err
}

func ModifyPassword(username string, password string, newpass string) (int64, *User, error) {
	o := orm.NewOrm()
	var objs []User
	_, err := o.Raw("SELECT * FROM user WHERE username = ? ORDER BY id DESC", username).QueryRows(&objs)
	if err != nil {
		return 1003, nil, err
	}
	if len(objs) > 0 {
		if objs[0].Password == password {
			cid := objs[0].Id
			obj := &User{Id: cid}
			obj.Password = newpass
			_, err = o.Update(obj, "password")
			if err != nil {
				return 1003, nil, err
			}
			objs[0].Password = newpass
			return 0, &objs[0], err
		} else {
			return 1008, nil, err
		}
	} else {
		return 1007, nil, err
	}
}

func EmailUser(email string) (int64, *User, error) {
	o := orm.NewOrm()
	var objs []User
	_, err := o.Raw("SELECT * FROM user WHERE email = ? ORDER BY id DESC", email).QueryRows(&objs)
	if err != nil {
		return 1003, nil, err
	}
	if len(objs) > 0 {
		newpass := "123456"
		cid := objs[0].Id
		obj := &User{Id: cid}
		obj.Password = newpass
		_, err = o.Update(obj, "password")
		if err != nil {
			return 1003, nil, err
		}
		objs[0].Password = newpass
		return 0, &objs[0], err
	} else {
		return 1009, nil, err
	}
}

func getUserToken(uid string) string {
	my_time := time.Now().Unix()
	time_str := strconv.FormatInt(my_time, 10)
	rand.Seed(time.Now().UnixNano())
	ri := rand.Intn(3)
	ri_str := strconv.Itoa(ri)
	token := uid + "_" + time_str + "_" + ri_str
	return token
}

func getUserSecret(uid string) string {
	secret := "hellobeer"
	return secret
}
