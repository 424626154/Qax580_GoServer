package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"strconv"
	"time"
)

//帖子
type Post struct {
	Id         int64
	Title      string    `orm:"size(100)"`
	Info       string    `orm:"size(1000)"`
	CreateTime time.Time `orm:"index"`
	Examine    int16
	Label      int16 // 1个人 2 官方
	Image      string
	Type       int16 //0 默认 1房产 2 二手 3 出兑 4 招聘
}

//意见反馈
type Feedback struct {
	Id         int64
	Info       string    `orm:"size(1000)"`
	CreateTime time.Time `orm:"index"`
	State      int16
}

//微信公众号
type Wxnum struct {
	Id         int64
	Title      string    `orm:"size(100)"`
	Info       string    `orm:"size(1000)"`
	Num        string    `orm:"size(100)"`
	CreateTime time.Time `orm:"index"`
	Image      string
}

type Admin struct {
	Id       int64
	Username string `orm:"size(100)"`
	Password string `orm:"size(1000)"`
}

type AccessTokenJson struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenID       string `json:"openid"`
	Scope        string `json:"scope"`
	Unionid      string `json:"unionid"`
	ErrCode      int64  `json:"errcode"`
	ErrMsg       string `json:"errmsg"`
}

type UserInfoJson struct {
	OpenId     string   `json:"openid"`
	NickeName  string   `json:"nickname"`
	Sex        int32    `json:"sex"`
	Province   string   `json:"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	HeadImgurl string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
	Unionid    string   `json:"unionid"`
	ErrCode    int64    `json:"errcode"`
	ErrMsg     string   `json:"errmsg"`
}

func RegisterDB() {
	// set default database
	isdebug := "true"
	iniconf, err := config.NewConfig("json", "conf/myconfig.json")
	if err != nil {
		beego.Debug(err)
	} else {
		isdebug = iniconf.String("qax580::isdebug")
	}
	if isdebug == "true" {
		orm.RegisterDataBase("default", "mysql", "root:@/qax580?charset=utf8")
	} else {
		orm.RegisterDataBase("default", "mysql", "root:sbb890503@/qax580go?charset=utf8")
	}
	// register model
	orm.RegisterModel(new(Post))
	orm.RegisterModel(new(Feedback))
	orm.RegisterModel(new(Wxnum))
	orm.RegisterModel(new(Admin))
	// create table
	orm.RunSyncdb("default", false, true)
}

//修改帖子内容
func UpdatePostInfo(id string, title string, info string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &Post{Id: cid}
	cate.Title = title
	cate.Info = info
	_, err = o.Update(cate, "title", "info")
	return err
}

//修改帖子审核状态
func UpdatePost(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &Post{Id: cid}
	cate.Examine = 1
	_, err = o.Update(cate, "examine")
	return err
}

//修改帖子审核状态1
func UpdatePostExamine(id string, exa int16) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &Post{Id: cid}
	cate.Examine = exa
	_, err = o.Update(cate, "examine")
	return err
}

func GetAllPostsAdmin() ([]Post, error) {
	o := orm.NewOrm()
	var posts []Post
	_, err := o.Raw("SELECT * FROM post  ORDER BY id DESC").QueryRows(&posts)
	return posts, err
}

func GetAllPosts() ([]Post, error) {
	o := orm.NewOrm()
	var posts []Post
	_, err := o.Raw("SELECT * FROM post WHERE examine = ? ORDER BY id DESC ", 1).QueryRows(&posts)
	return posts, err
}

func GetOnePost(id string) (*Post, error) {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	post := &Post{Id: cid}
	err = o.Read(post)
	return post, err
}
func AddPost(title string, info string, image string) error {
	o := orm.NewOrm()

	cate := &Post{Title: title, Info: info, CreateTime: time.Now(), Image: image}

	// 查询数据
	qs := o.QueryTable("post")
	err := qs.Filter("title", title).One(cate)
	if err == nil {
		return err
	}

	// 插入数据
	_, err = o.Insert(cate)
	if err != nil {
		return err
	}

	return nil
}
func AddPostLabel(title string, info string, label int16, image string) error {
	o := orm.NewOrm()

	cate := &Post{Title: title, Info: info, CreateTime: time.Now(), Label: label, Image: image}

	// 查询数据
	qs := o.QueryTable("post")
	err := qs.Filter("title", title).One(cate)
	if err == nil {
		return err
	}

	// 插入数据
	_, err = o.Insert(cate)
	if err != nil {
		return err
	}

	return nil
}

func DeletePost(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &Post{Id: cid}
	_, err = o.Delete(cate)
	return err
}

func QueryLimitPost(nums int64) ([]Post, error) {
	o := orm.NewOrm()
	var posts []Post
	_, err := o.Raw("SELECT * FROM post ORDER BY id DESC LIMIT ? ", nums).QueryRows(&posts)
	return posts, err
}
func QueryFuzzyLimitPost(fuzzy string, nums int64) ([]Post, error) {
	o := orm.NewOrm()
	var posts []Post
	_, err := o.Raw("SELECT * FROM post WHERE info LIKE ? ORDER BY id DESC LIMIT ? ", "%"+fuzzy+"%", nums).QueryRows(&posts)
	return posts, err
}

/*******************意见反馈********************/

func AddFeedback(info string) error {
	o := orm.NewOrm()
	time := time.Now()
	cate := &Feedback{Info: info, CreateTime: time}

	// 查询数据
	qs := o.QueryTable("feedback")
	err := qs.Filter("createTime", time).One(cate)
	if err == nil {
		return err
	}

	// 插入数据
	_, err = o.Insert(cate)
	if err != nil {
		return err
	}

	return nil
}

//意见反馈列表
func GetAllFeedbacks() ([]Feedback, error) {
	o := orm.NewOrm()
	var feedbacks []Feedback
	_, err := o.Raw("SELECT * FROM feedback  ORDER BY id DESC").QueryRows(&feedbacks)
	return feedbacks, err
}

func GetOneFeedback(id string) (*Feedback, error) {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	feedback := &Feedback{Id: cid}
	err = o.Read(feedback)
	return feedback, err
}

//添加微信公众号

func AddPublicNumber(title string, info string, num string, image string) error {
	o := orm.NewOrm()

	cate := &Wxnum{Title: title, Info: info, Num: num, CreateTime: time.Now(), Image: image}

	// 查询数据
	qs := o.QueryTable("wxnum")
	err := qs.Filter("title", title).One(cate)
	if err == nil {
		return err
	}

	// 插入数据
	_, err = o.Insert(cate)
	if err != nil {
		return err
	}

	return nil
}

func GetAllWxnums() ([]Wxnum, error) {
	o := orm.NewOrm()
	var wxnums []Wxnum
	_, err := o.Raw("SELECT * FROM wxnum  ORDER BY id DESC").QueryRows(&wxnums)
	return wxnums, err
}

/*
添加后台用户
*/
func AddAdmin(username string, password string) error {
	o := orm.NewOrm()

	admin := &Admin{Username: username, Password: password}
	// 查询数据
	qs := o.QueryTable("admin")
	err := qs.Filter("username", username).One(admin)
	if err == nil {
		return err
	}
	// 插入数据
	_, err = o.Insert(admin)
	if err != nil {
		return err
	}

	return nil
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

func GetAllAdmins() ([]Admin, error) {
	o := orm.NewOrm()
	var admins []Admin
	_, err := o.Raw("SELECT * FROM admin  ORDER BY id DESC").QueryRows(&admins)
	return admins, err
}

func DeleteAdmin(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	admin := &Admin{Id: cid}
	_, err = o.Delete(admin)
	return err
}
