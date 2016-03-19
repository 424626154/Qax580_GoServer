package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"strconv"
	"time"
)

type RinseJson struct {
	Rtype   string `json:"rtype"`
	Data    string `json:"data"`
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	Phone   string `json:"phone"`
	Pwd     string `json:"pwd"`
}

/**
冲洗绑定
*/
// type RBinding struct {
// 	Id      int64
// 	BType   int16  //绑定类型 1 手机号
// 	Phone   string //手机号
// 	Pwd     string //密码
// 	Account string //唯一帐号
// }

/**
冲洗账户
*/
type RUser struct {
	Id      int64
	Phone   string //手机号
	Pwd     string //密码
	Account string //唯一帐号
	RId     int64  //显示id
}

func RegisterDB() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:@/photo?charset=utf8")
	beego.Debug("root:@/photo?charset=utf8")
	// register model
	// orm.RegisterModel(new(RBinding))        //冲洗绑定
	orm.RegisterModel(new(RUser)) //冲洗帐号
	// create table
	orm.RunSyncdb("default", false, true)
}

/***冲洗数据库***/

// func AddRUserPhone(phone string, pwd string) error {
// 	o := orm.NewOrm()
// 	cate := &RBinding{BType: int16(1), Phone: phone, Pwd: pwd}
// 	// 插入数据
// 	_, err := o.Insert(cate)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func GetRUserPhone(phone string, pwd string) (*RBinding, error) {
// 	o := orm.NewOrm()
// 	var objs []RBinding
// 	_, err := o.Raw("SELECT * FROM r_binding  WHERE phone = ? ORDER BY id DESC", phone).QueryRows(&objs)
// 	if err != nil {
// 		beego.Error(err)
// 		return nil, err
// 	}
// 	if len(objs) == 0 {
// 		return nil, nil
// 	}
// 	obj := &objs[0]
// 	return obj, nil
// }
/**
添加用户
*/
func AddRUser(phone string, pwd string, account string, rid int64) error {
	o := orm.NewOrm()
	obj := &RUser{Phone: phone, Pwd: pwd, Account: account, RId: rid}
	// 插入数据
	_, err := o.Insert(obj)
	if err != nil {
		return err
	}
	return nil
}

/**
获得用户
*/
func GetRUser(phone string) (*RUser, error) {
	o := orm.NewOrm()
	var objs []RUser
	_, err := o.Raw("SELECT * FROM r_user  WHERE phone = ? ORDER BY id DESC", phone).QueryRows(&objs)
	if err != nil {
		beego.Error(err)
		return nil, err
	}
	if len(objs) == 0 {
		return nil, nil
	}
	obj := &objs[0]
	return obj, nil
}

/**
获得用户account
*/
func GetRUserAccount(account string) (*RUser, error) {
	o := orm.NewOrm()
	var objs []RUser
	_, err := o.Raw("SELECT * FROM r_user  WHERE account = ? ORDER BY id DESC", account).QueryRows(&objs)
	if err != nil {
		beego.Error(err)
		return nil, err
	}
	if len(objs) == 0 {
		return nil, nil
	}
	obj := &objs[0]
	return obj, nil
}

/**
获得用户account
*/
func GetRUserPP(phone string) (*RUser, error) {
	o := orm.NewOrm()
	var objs []RUser
	_, err := o.Raw("SELECT * FROM r_user  WHERE phone = ? ORDER BY id DESC", phone).QueryRows(&objs)
	if err != nil {
		beego.Error(err)
		return nil, err
	}
	if len(objs) == 0 {
		return nil, nil
	}
	obj := &objs[0]
	return obj, nil
}

/**
活动用户数量
*/

func GetUserCount() (int64, error) {
	o := orm.NewOrm()
	count, err := o.QueryTable("r_user").Count()
	return count, err
}
