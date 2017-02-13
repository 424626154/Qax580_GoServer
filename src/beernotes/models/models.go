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
	orm.RegisterModel(new(Formula))
	orm.RegisterModel(new(FormulaComment))
	orm.RegisterModel(new(FormulaLove))
	orm.RegisterModel(new(Message))
	orm.RegisterModel(new(WNumber))
	orm.RegisterModel(new(KnowBrew))
	orm.RegisterModel(new(Business))
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

//配方
type Formula struct {
	Id           int64
	Fid          string `orm:"size(100)"`
	Uid          string `orm:"size(100)"`
	Username     string `orm:"size(500)"`
	Fname        string `orm:"size(500)"`  //配方名称
	Malts        string `orm:"size(2000)"` //麦芽
	Hopss        string `orm:"size(2000)"` //啤酒花
	Yeasts       string `orm:"size(2000)"` //酵母
	Water        int64  //水
	Accessoriess string `orm:"size(2000)"` //辅料
	Time         int64
	Lovenum      int64 //点赞数量
	Islove       bool  //是否点赞
	Isdele       bool  //是否可以删除
}

//配方评论
type FormulaComment struct {
	Id         int64
	Fid        string `orm:"size(100)"`
	Uid        string `orm:"size(100)"`
	Username   string `orm:"size(500)"`
	Comment    string `orm:"size(2000)"`
	Time       int64
	Fcid       int64  //被评论的评论
	Fcuid      string `orm:"size(100)"`
	Fcusername string `orm:"size(500)"`
	Lovenum    int64  //点赞数量
	Islove     bool   //是否点赞
	Isdele     bool   //是否可以删除
}

type FormulaLove struct {
	Id     int64
	Fid    string `orm:"size(100)"`
	Uid    string `orm:"size(100)"`
	Fcid   int64  //被评论的评论
	Islove bool   //是否点赞
	Time   int64
}

type Message struct {
	Id      int64
	Uid     string `orm:"size(100)"`
	Title   string `orm:"size(500)"`
	Content string `orm:"size(2000)"`
	Read    bool
	Time    int64
}

type WNumber struct {
	Id     int64
	Title  string `orm:"size(500)"`
	Info   string `orm:"size(2000)"`
	Number string `orm:"size(100)"`
	Image  string
	Time   int64
	State  int8 //0未上线  1 已上线
}


type KnowBrew struct {
	Id       int64
	Title    string `orm:"size(500)"`
	Brief    string `orm:"size(2000)"`
	Link     string `orm:"size(500)"`
	Uid      string `orm:"size(100)"`
	Username string `orm:"size(500)"`
	From     int8   // 0用户 1官方
	Time     int64
	State    int8 //0 未上线 1上线
	Examine  int8 //0 未审核 1合格 2 驳回
	Type     int8 // 1文章 2视频
}

type Business struct {
	Id    int64
	Title string `orm:"size(100)"`
	Brief string `orm:"size(1000)"`
	Link  string `orm:"size(500)"`
	State int16
	Time  int64
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
		return 1003, nil, err
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
func GetAllUser() ([]User, error) {
	o := orm.NewOrm()
	var objs []User
	_, err := o.Raw("SELECT * FROM user ORDER BY id DESC").QueryRows(&objs)
	if err != nil {
		return nil, err
	}
	return objs, err
}

func GetOneUserToken(token string) (*User, error) {
	o := orm.NewOrm()
	var objs []User
	_, err := o.Raw("SELECT * FROM user WHERE token = ? ORDER BY id DESC", token).QueryRows(&objs)
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

func AddFormula(uid string, username string, fid string, fname string, malts string, hopss string, yeasts string, water string, accessoriess string) (int64, *Formula, error) {
	iwater, err := strconv.ParseInt(water, 10, 64)
	if err != nil {
		return 1003, nil, err
	}
	o := orm.NewOrm()
	var objs []Formula
	_, err = o.Raw("SELECT * FROM formula WHERE fid = ? ORDER BY id DESC", fid).QueryRows(&objs)
	if err != nil {
		return 1003, nil, err
	}
	if len(objs) > 0 {
		return 1010, nil, err
	}
	my_time := time.Now().Unix()
	obj := &Formula{Uid: uid, Username: username, Fid: fid, Fname: fname, Malts: malts, Hopss: hopss, Yeasts: yeasts, Water: iwater, Accessoriess: accessoriess, Time: my_time}
	// 插入数据
	_, err = o.Insert(obj)
	if err != nil {
		return 1003, nil, err
	}
	return 0, obj, nil
}

func DeleteFormula(fid string, uid string) error {
	o := orm.NewOrm()
	var objs []Formula
	_, err := o.Raw("SELECT * FROM formula WHERE fid = ? AND uid = ? ORDER BY id DESC", fid, uid).QueryRows(&objs)
	if err != nil {
		return err
	}
	if len(objs) > 0 {
		o := orm.NewOrm()
		cate := &Formula{Id: objs[0].Id}
		_, err = o.Delete(cate)
		if err != nil {
			return err
		}
	}
	return err
}

func GetFormula() ([]Formula, error) {
	o := orm.NewOrm()
	var objs []Formula
	_, err := o.Raw("SELECT * FROM formula ORDER BY id DESC").QueryRows(&objs)
	if err != nil {
		return nil, err
	}
	return objs, err
}

func AddFormulaComment(fid string, uid string, username string, comment string, fcid string) (*FormulaComment, error) {
	o := orm.NewOrm()
	var objs []FormulaComment
	_, err := o.Raw("SELECT * FROM formula_comment WHERE id = ? ORDER BY id DESC", fcid).QueryRows(&objs)
	if err != nil {
		return nil, err
	}
	my_time := time.Now().Unix()
	beego.Debug("objs:", objs, "fcid:", fcid)
	if len(objs) > 0 {
		obj := &FormulaComment{Fid: fid, Uid: uid, Username: username, Comment: comment, Fcid: objs[0].Id, Fcuid: objs[0].Uid, Fcusername: objs[0].Username, Time: my_time}
		// 插入数据
		_, err = o.Insert(obj)
		if err != nil {
			return obj, err
		}
	} else {
		obj := &FormulaComment{Fid: fid, Uid: uid, Username: username, Comment: comment, Time: my_time}
		// 插入数据
		_, err = o.Insert(obj)
		if err != nil {
			return obj, err
		}
	}
	return nil, nil
}

func DeleteFormulaComment(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &FormulaComment{Id: cid}
	_, err = o.Delete(cate)
	if err != nil {
		return err
	}
	return err
}

func GetFormulaComment() ([]FormulaComment, error) {
	o := orm.NewOrm()
	var objs []FormulaComment
	_, err := o.Raw("SELECT * FROM formula_comment ORDER BY id DESC").QueryRows(&objs)
	if err != nil {
		return nil, err
	}
	return objs, err
}

func GetFormulaLoveNum(fid string, fcid int64) (int, error) {
	o := orm.NewOrm()
	var objs []FormulaLove
	_, err := o.Raw("SELECT * FROM formula_love WHERE fid = ? AND fcid = ? AND islove = true ORDER BY id DESC", fid, fcid).QueryRows(&objs)
	if err != nil {
		return 0, err
	}
	return len(objs), err
}

func GetFormulaLove(uid string, fid string, fcid int64) (bool, error) {
	o := orm.NewOrm()
	var objs []FormulaLove
	_, err := o.Raw("SELECT * FROM formula_love WHERE uid = ? AND fid = ? AND fcid = ? ORDER BY id DESC", uid, fid, fcid).QueryRows(&objs)
	if err != nil {
		return false, err
	}
	if len(objs) > 0 {
		return objs[0].Islove, err
	}
	return false, err
}

func AddFormulaLove(uid string, fid string, fcid string, islove bool) (*FormulaLove, error) {
	o := orm.NewOrm()
	var objs []FormulaLove
	_, err := o.Raw("SELECT * FROM formula_love WHERE uid = ? AND fid = ? AND fcid = ? ORDER BY id DESC", uid, fid, fcid).QueryRows(&objs)
	if err != nil {
		return nil, err
	}
	my_time := time.Now().Unix()
	if len(objs) > 0 {
		obj := &FormulaLove{Id: objs[0].Id}
		obj.Islove = islove
		obj.Time = my_time
		objs[0].Islove = islove
		_, err = o.Update(obj, "islove", "time")
		if err != nil {
			return nil, err
		}
		return &objs[0], nil
	} else {
		ifcid, err := strconv.ParseInt(fcid, 10, 64)
		if err != nil {
			return nil, err
		}
		obj := &FormulaLove{Uid: uid, Fid: fid, Fcid: ifcid, Islove: islove, Time: my_time}
		// 插入数据
		_, err = o.Insert(obj)
		if err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

func AddMessage(uid string, title string, content string) (*Message, error) {
	o := orm.NewOrm()
	my_time := time.Now().Unix()
	obj := &Message{Uid: uid, Title: title, Content: content, Time: my_time}
	// 插入数据
	_, err := o.Insert(obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
func GetReadMessage(uid string) ([]Message, error) {
	o := orm.NewOrm()
	var objs []Message
	_, err := o.Raw("SELECT * FROM message WHERE uid = ? ORDER BY id DESC", uid).QueryRows(&objs)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		var temp_objs []Message
		for i := 0; i < len(objs); i++ {
			beego.Debug(objs[i].Read)
			if objs[i].Read == false {
				temp_objs = append(temp_objs, objs[i])
			}
		}
		return temp_objs, err
	}
	return objs, err
}

func UpMessage(id int64, read bool) (*Message, error) {
	o := orm.NewOrm()
	my_time := time.Now().Unix()
	var objs []Message
	_, err := o.Raw("SELECT * FROM message WHERE id = ? ORDER BY id DESC", id).QueryRows(&objs)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		obj := &Message{Id: objs[0].Id}
		obj.Read = read
		obj.Time = my_time
		objs[0].Read = read
		_, err = o.Update(obj, "read", "time")
		if err != nil {
			return nil, err
		}
		return &objs[0], nil
	}
	return nil, err
}

func AddWNumber(title string, info string, number string, image string) (*WNumber, error) {
	o := orm.NewOrm()
	my_time := time.Now().Unix()
	obj := &WNumber{Title: title, Info: info, Number: number, Image: image, Time: my_time}
	// 插入数据
	_, err := o.Insert(obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
func GetWNumber(id string) (*WNumber, error) {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	var objs []WNumber
	_, err = o.Raw("SELECT * FROM w_number WHERE id = ? ORDER BY id DESC", cid).QueryRows(&objs)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return &objs[0], err
	}
	return nil, err
}

func GetAllWNumber() ([]WNumber, error) {
	o := orm.NewOrm()
	var objs []WNumber
	_, err := o.Raw("SELECT * FROM w_number ORDER BY id DESC").QueryRows(&objs)
	if err != nil {
		return nil, err
	}
	return objs, err
}

func GetWNumberState(state int8) ([]WNumber, error) {
	o := orm.NewOrm()
	var objs []WNumber
	_, err := o.Raw("SELECT * FROM w_number WHERE state = ? ORDER BY id DESC", state).QueryRows(&objs)
	if err != nil {
		return nil, err
	}
	return objs, err
}

func UpdateWNumberState(id string, state int8) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &WNumber{Id: cid}
	cate.State = state
	_, err = o.Update(cate, "state")
	return err
}

func DeleteWNumber(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &WNumber{Id: cid}
	_, err = o.Delete(cate)
	return err
}
func UpdateWNumberInfo(id string, title string, info string, number string) (*WNumber, error) {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	my_time := time.Now().Unix()
	cate := &WNumber{Id: cid}
	cate.Title = title
	cate.Info = info
	cate.Number = number
	cate.Time = my_time
	_, err = o.Update(cate, "title", "info", "number", "time")
	return cate, err
}

func UpdateWNumberImg(id string, image string) (*WNumber, error) {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	my_time := time.Now().Unix()
	cate := &WNumber{Id: cid}
	cate.Image = image
	cate.Time = my_time
	_, err = o.Update(cate, "image", "time")
	return cate, err
}

func AddKnowBrew(title string, brief string, link string, uid string, username string, kbtype int8) (*KnowBrew, error) {
	o := orm.NewOrm()
	my_time := time.Now().Unix()
	obj := &KnowBrew{Title: title, Brief: brief, Link: link, Uid: uid, Username: username, Type: kbtype, Time: my_time}
	// 插入数据
	_, err := o.Insert(obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
func AddAdminKnowBrew(title string, brief string, link string, kbtype int8) (*KnowBrew, error) {
	o := orm.NewOrm()
	my_time := time.Now().Unix()
	obj := &KnowBrew{Title: title, Brief: brief, Link: link, From: 1, Examine: 1, Type: kbtype, Time: my_time}
	// 插入数据
	_, err := o.Insert(obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func GetAllKnowBrew() ([]KnowBrew, error) {
	o := orm.NewOrm()
	var objs []KnowBrew
	_, err := o.Raw("SELECT * FROM know_brew ORDER BY id DESC").QueryRows(&objs)
	if err != nil {
		return nil, err
	}
	return objs, err
}

func UpdateKnowBrewState(id string, state int8) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &KnowBrew{Id: cid}
	cate.State = state
	_, err = o.Update(cate, "state")
	return err
}

func UpdateKnowBrewExamine(id string, examine int8) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &KnowBrew{Id: cid}
	cate.Examine = examine
	_, err = o.Update(cate, "examine")
	return err
}

func GetKnowBrewSAE(state int8, examine int8) ([]KnowBrew, error) {
	o := orm.NewOrm()
	var objs []KnowBrew
	_, err := o.Raw("SELECT * FROM know_brew WHERE state = ? AND examine = ? ORDER BY id DESC", state, examine).QueryRows(&objs)
	if err != nil {
		return nil, err
	}
	return objs, err
}

func GetKnowBrew(id string) (*KnowBrew, error) {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	var objs []KnowBrew
	_, err = o.Raw("SELECT * FROM know_brew WHERE id = ? ORDER BY id DESC", cid).QueryRows(&objs)
	if err != nil {
		return nil, err
	}
	if len(objs) > 0 {
		return &objs[0], err
	}
	return nil, err
}

func UpdateKnowBrew(id string, title string, brief string, link string, kbtype int8) (*KnowBrew, error) {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	cate := &KnowBrew{Id: cid}
	cate.Title = title
	cate.Brief = brief
	cate.Link = link
	cate.Type = kbtype
	_, err = o.Update(cate, "title", "brief", "link", "type")
	return cate, err
}

func DeleteKnowBrew(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &KnowBrew{Id: cid}
	_, err = o.Delete(cate)
	return err
}

func AddBusiness(title string, brief string, link string) (*Business, error) {
	o := orm.NewOrm()
	my_time := time.Now().Unix()
	obj := &Business{Title: title, Brief: brief, Link: link, Time: my_time}
	// 插入数据
	_, err := o.Insert(obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
func GetAllBusiness() ([]Business, error) {
	o := orm.NewOrm()
	var objs []Business
	_, err := o.Raw("SELECT * FROM business ORDER BY id DESC").QueryRows(&objs)
	if err != nil {
		return nil, err
	}
	return objs, err
}
func UpdateBusinessState(id string, state int16) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &Business{Id: cid}
	cate.State = state
	_, err = o.Update(cate, "state")
	return err
}

func DeleteBusiness(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &Business{Id: cid}
	_, err = o.Delete(cate)
	return err
}

func UpdateBusiness(id string, title string, brief string, link string) (*Business, error) {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	my_time := time.Now().Unix()
	cate := &Business{Id: cid}
	cate.Title = title
	cate.Brief = brief
	cate.Link = link
	cate.Time = my_time
	_, err = o.Update(cate, "title", "brief", "link", "time")
	return cate, err
}

func GetOneBusiness(id string) (*Business, error) {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	var objes []Business
	_, err = o.Raw("SELECT * FROM business WHERE id = ? ", cid).QueryRows(&objes)
	if err != nil {
		return nil, err
	}
	obj := &Business{}
	if len(objes) > 0 {
		obj = &objes[0]
	}
	return obj, err
}

func GetBusinessState(state int8) ([]Business, error) {
	o := orm.NewOrm()
	var objs []Business
	_, err := o.Raw("SELECT * FROM business WHERE state = ? ORDER BY id DESC", state).QueryRows(&objs)
	if err != nil {
		return nil, err
	}
	return objs, err
}
