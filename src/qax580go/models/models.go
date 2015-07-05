package models

import (
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
	Type       int16 //1房产 2 二手 3 出兑 4 招聘
}

//意见反馈
type Feedback struct {
	Id         int64
	Info       string    `orm:"size(1000)"`
	CreateTime time.Time `orm:"index"`
	State      int16
}

func RegisterDB() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:@/qax580?charset=utf8")
	// orm.RegisterDataBase("default", "mysql", "root:sbb890503@/qax580go?charset=utf8")
	// register model
	orm.RegisterModel(new(Post))
	orm.RegisterModel(new(Feedback))
	// create table
	orm.RunSyncdb("default", false, true)
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
