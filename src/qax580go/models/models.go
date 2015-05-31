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
}

func RegisterDB() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:@/qax580go?charset=utf8")
	// register model
	orm.RegisterModel(new(Post))

	// create table
	orm.RunSyncdb("default", false, true)
}

func GetAllPosts() ([]*Post, error) {
	o := orm.NewOrm()

	posts := make([]*Post, 0)

	qs := o.QueryTable("post")
	_, err := qs.All(&posts)
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
func AddPost(title string, info string) error {
	o := orm.NewOrm()

	cate := &Post{Title: title, Info: info, CreateTime: time.Now()}

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
