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
	Examine    int16     //审核状态0未审核1 审核
	Label      int16     // 1个人 2 官方
	Image      string
	Type       int16  //0 默认 1房产 2 二手 3 出兑 4 招聘
	OpenId     string `orm:"size(500)"`
	NickeName  string `orm:"size(100)"`
	Sex        int32
	HeadImgurl string `orm:"size(500)"`
	Time       int64
	City       string //城市
	Bfrom      bool   //来源
	Fromshow   string //来源显示
	Fromurl    string //来源链接
	State1     int16  //是否违规 0 违规 1 违规
}

//意见反馈
type Feedback struct {
	Id         int64
	Info       string    `orm:"size(1000)"`
	CreateTime time.Time `orm:"index"`
	State      int16
	OpenId     string `orm:"size(500)"`
	NickeName  string `orm:"size(100)"`
	Sex        int32
	HeadImgurl string `orm:"size(500)"`
	Time       int64
}

//微信公众号
type Wxnum struct {
	Id       int64
	Title    string `orm:"size(100)"`
	Info     string `orm:"size(1000)"`
	Num      string `orm:"size(100)"`
	Evaluate string `orm:"size(1000)"` //评价
	Image    string
	Time     int64
	State    int8 //0未上线  1 已上线
}

//微信号
type WeixinNumber struct {
	Id       int64
	Name     string `orm:"size(100)"`
	Info     string `orm:"size(1000)"`
	Num      string `orm:"size(100)"`
	Evaluate string `orm:"size(1000)"` //评价
	Image    string
	Time     int64
	State    int8 //0未上线  1 已上线
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

type QRCodeJson struct {
	Ticket         string `json:"ticket"`
	ExpiresSeconds int64  `json:"expire_seconds"`
	Url            string `json:"url"`
	ErrCode        int64  `json:"errcode"`
	ErrMsg         string `json:"errmsg"`
}

type Wxuserinfo struct {
	Id            int64
	Subscribe     int32  `json:"subscribe"`      //用户是否订阅该公众号标识，值为0时，代表此用户没有关注该公众号，拉取不到其余信息。
	OpenId        string `json:"openid"`         //用户的标识，对当前公众号唯一
	NickeName     string `json:"nickname"`       //用户的昵称
	Sex           int32  `json:"sex"`            //用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
	City          string `json:"city"`           //用户所在城市
	Country       string `json:"country"`        //用户所在国家
	Province      string `json:"province"`       //	用户所在省份
	Language      string `json:"language"`       //用户的语言，简体中文为zh_CN
	HeadImgurl    string `json:"headimgurl"`     //用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），用户没有头像时该项为空。若用户更换头像，原有头像URL将失效。
	SubscribeTime int64  `json:"subscribe_time"` //用户关注时间，为时间戳。如果用户曾多次关注，则取最后关注时间
	Unionid       string `json:"unionid"`        //只有在用户将公众号绑定到微信开放平台帐号后，才会出现该字段。详见：获取用户个人信息（UnionID机制）
	Remark        string `json:"remark"`         //公众号运营者对粉丝的备注，公众号运营者可在微信公众平台用户管理界面对粉丝添加备注
	Groupid       int32  `json:"groupid"`        //用户所在的分组ID
	ErrCode       int32  `json:"errcode"`        //0 成功 1 失败 参数错误
	ErrMsg        string `json:"errmsg"`         //失败详情
	Experience    int64  //经验
	Money         int64  //金钱
}

/*
用户金钱纪录
*/
type UserMoneyRecord struct {
	Id        int64
	OpenId    string
	Money     int64
	Time      int64
	MoneyType int64 //事件类型 1注册 2 帖子被赞 3兑换商品
}
type JsApiTicketJson struct {
	Id        int64
	Ticket    string `json:"ticket"`
	ExpiresIn int64  `json:"expires_in"`
	ErrCode   int64  `json:"errcode"`
	ErrMsg    string `json:"errmsg"`
}

type WeatherJson struct {
	Resultcode string `json:"resultcode"`
	Reason     string `json:"reason"`
	Result     Result `json:"result"`
	ErrorCode  int64  `json:"error_code"`
}

type Result struct {
	Today Today `json:"today"`
}
type Today struct {
	City           string `json:"city"`
	DateY          string `json:"date_y"`
	Week           string `json:"week"`
	Temperature    string `json:"temperature"`
	Weather        string `json:"weather"`
	Wind           string `json:"wind"`
	Dressingindex  string `json:"dressing_index"`
	DressingAdvice string `json:"dressing_advice"`
	UvIndex        string `json:"uv_index"`
	WashIndex      string `json:"wash_index"`
	WravelIndex    string `json:"wravel_index"`
	WxerciseIndex  string `json:"wxercise_index"`
}

type NewsKey struct {
	Id         int64
	Info       string    `orm:"size(1000)"`
	CreateTime time.Time `orm:"index"`
	Op         int32
}

type QureyUser struct {
	Id      int64
	UserId  string `orm:"size(200)"`
	Time    int64
	Index   int32  //索引
	Keyword string `orm:"size(200)"` //搜索关键字
}

type Guanggao struct {
	Id         int64
	Image      string
	Title      string
	Content    string `orm:"size(4096)"`
	State      int32  //0未上线 1 上线
	Time       int64
	Blink      bool
	Link       string
	BImage     bool   //广告页图片是否显示主页
	ImageItem0 string //内容页图片0
	ImageItem1 string //内容页图片1
	ImageItem2 string //内容页图片2
}

//外卖
type Canting struct {
	Id          int64
	Name        string
	Address     string
	Image       string
	Time        int64
	State       int8
	Starthour   int8   //开始时间 小时
	Startminute int8   //开始时间 分钟
	Endhour     int8   //结束时间 小时
	Endminute   int8   //结束时间 分钟
	Phonenumber string //电话号码
}

//菜单
type Caidan struct {
	Id    int64
	Fid   int64
	Name  string
	Info  string
	Image string
	Time  int64
	State int8
	Price string //价格
	Mtype string //菜品类型
}

//帖子帮助
type Posthelp struct {
	Id     int64
	PostId int64
	OpenId string
	State  int32 //0 无帮助 1 有帮助
	Time   int64
}

//商品
type Commodity struct {
	Id         int64
	Name       string    //名称
	Info       string    //简介
	Image      string    //图片地址
	Money      int64     //价格
	State      int16     //0 未上架 1 上架
	CreateTime time.Time `orm:"index"` //创建时间
	Time       int64     //显示时间
}

//订单
type Uorder struct {
	Id           int64
	OpenId       string
	CommodityId  int64     //商品ID
	State        int16     //0 未兑换 1 已兑换
	CreateTime   time.Time `orm:"index"` //创建时间
	Time         int64     //显示时间
	ExchangeTime time.Time `orm:"index"` //兑换时间
	Time1        int64     //显示兑换时间
}

//显示订单
type ShowOrder struct {
	Id           int64
	OpenId       string
	CommodityId  int64     //商品ID
	State        int16     //0 未兑换 1 已兑换
	CreateTime   time.Time `orm:"index"` //创建时间
	Time         int64     //显示时间
	ExchangeTime time.Time `orm:"index"` //兑换时间
	Time1        int64     //显示兑换时间
	Commodity    *Commodity
	NickeName    string
	HeadImgurl   string
}

//商户
type ShangHu struct {
	Id         int64
	Name       string    //名称
	Info       string    //简介
	Image      string    //图片地址
	Type       int16     //类型
	Recommend  int16     //0 未推荐 1 推荐
	State      int16     //0 未上线 1 上线
	CreateTime time.Time `orm:"index"` //创建时间
	Time       int64     //显示时间
}

//搜索关键字
type Keywords struct {
	Id         int64
	KeyName    string    //关键字
	State      int16     //0 未上线 1 上线
	CreateTime time.Time `orm:"index"` //创建时间
	Time       int64     //显示时间
}

//关键字对象
type Keyobj struct {
	Id         int64
	KeyId      int64
	Image      string
	Title      string
	Info       string
	Url        string
	State      int16     //0 未上线 1 上线
	CreateTime time.Time `orm:"index"` //创建时间
	Time       int64     //显示时间
}

/**
投票组
*/
type Polls struct {
	Id            int64
	Title         string //投票组标题
	Info          string //投票组内容
	Image         string //投票组图片
	State         int16  //投票组状态 0未上线 1 上线
	StartTimeLong int64  //投票组开始时间
	EndTimeLong   int64  //投票组结束时间
	Pageview      int64  //访问量
	More          string //更多链接
	Appid         string //验证的投票Appid
	Secret        string //验证的投票Secret
	Prize         string //活动奖品
	Ext           string //扩展信息
}

/**
投票对象
*/
type Poll struct {
	Id         int64
	PollsId    int64  //投票列表id
	Title      string //投票标题
	Info       string //投票内容
	ContactWay string //联系方式
	Image      string //投票图片
	State      int16  //状态 0未上线 1 上线
	CreateTime int64  //创建时间
	OpenId     string //创建投票人
	VoteNum    int32  //投票数
	Ranking    int32  //排名
	Del        int16  //删除 0用户未删除 1 用户删除
}

/**
选票
*/
type Vote struct {
	Id         int64
	PollsId    int64 //投票列表id
	PollId     int64 //投票id
	State      int16 //状态 0未上线 1 上线
	OpneId     string
	CreateTime int64 //投票时间
}

/**
通知
*/
type Notice struct {
	Id         int64
	FromId     string //发送者
	ToId       string //接受者
	ToRead     int16  //接受者是否已读
	ToDel      int16  //接受者是否删除
	Msg        string //消息内容
	State      int16  //状态 0未上线 1 上线
	CreateTime int64  //投票时间
	Ext        string //消息扩展 1 帖子ID 2贴在ID
	NType      int16  //消息类型 1发布信息通过审核通知 2帖子存在违规通知
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
		beego.Debug("root:@/qax580?charset=utf8")
	} else {
		orm.RegisterDataBase("default", "mysql", "root:sbb890503@/qax580go?charset=utf8")
		beego.Debug("root:sbb890503@/qax580go?charset=utf8")
	}
	// register model
	orm.RegisterModel(new(Post))
	orm.RegisterModel(new(Feedback))
	orm.RegisterModel(new(Wxnum))
	orm.RegisterModel(new(WeixinNumber))
	orm.RegisterModel(new(Admin))
	orm.RegisterModel(new(Wxuserinfo)) //微信用户
	orm.RegisterModel(new(NewsKey))
	orm.RegisterModel(new(QureyUser))
	orm.RegisterModel(new(Guanggao))
	orm.RegisterModel(new(Canting))         //餐厅
	orm.RegisterModel(new(Caidan))          //菜单
	orm.RegisterModel(new(Posthelp))        //帖子帮助
	orm.RegisterModel(new(UserMoneyRecord)) //用户事件类型
	orm.RegisterModel(new(Commodity))       //商品
	orm.RegisterModel(new(Uorder))          //订单
	orm.RegisterModel(new(ShangHu))         //商户
	orm.RegisterModel(new(Keywords))        //关键字
	orm.RegisterModel(new(Keyobj))          //关键字对象
	orm.RegisterModel(new(Polls))           //投票组
	orm.RegisterModel(new(Poll))            //投票对象
	orm.RegisterModel(new(Vote))            //选票
	orm.RegisterModel(new(Notice))          //通知
	// create table
	orm.RunSyncdb("default", false, true)
}

//修改帖子内容
func UpdatePostInfo(id string, title string, info string, city string, bfrom bool, fromshow string, fromurl string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &Post{Id: cid}
	cate.Title = title
	cate.Info = info
	cate.City = city
	cate.Bfrom = bfrom
	cate.Fromshow = fromshow
	cate.Fromurl = fromurl
	_, err = o.Update(cate, "title", "info", "city", "bfrom", "fromshow", "fromurl")
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

//修改帖子违规状态
func UpdatePostState1(id string, exa int16) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &Post{Id: cid}
	cate.State1 = exa
	_, err = o.Update(cate, "state1")
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

/*
未审核帖子
*/
func GetAllState0Posts() ([]Post, error) {
	o := orm.NewOrm()
	var posts []Post
	_, err := o.Raw("SELECT * FROM post WHERE examine = ? ORDER BY id DESC ", 0).QueryRows(&posts)
	return posts, err
}

/*
未审核帖子数量
*/
func GetAllStateNum() (int, error) {
	o := orm.NewOrm()
	var posts []Post
	_, err := o.Raw("SELECT * FROM post WHERE examine = ? ORDER BY id DESC ", 0).QueryRows(&posts)
	return len(posts), err
}

//获得我发布的帖子
func GetAllPostsOpenid(openid string) ([]Post, error) {
	o := orm.NewOrm()
	var posts []Post
	_, err := o.Raw("SELECT * FROM post  WHERE open_id = ? ORDER BY id DESC", openid).QueryRows(&posts)
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
	create_time := time.Now()
	my_time := time.Now().Unix()
	cate := &Post{Title: title, Info: info, CreateTime: create_time, Time: my_time, Image: image}
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
func AddPostLabel(title string, info string, label int16, image string, city string, bfrom bool, fromshow string, fromurl string) (int64, error) {
	o := orm.NewOrm()

	create_time := time.Now()
	my_time := time.Now().Unix()
	cate := &Post{Title: title, Info: info, CreateTime: create_time, Time: my_time, Label: label, Image: image, City: city, Bfrom: bfrom, Fromshow: fromshow, Fromurl: fromurl}
	// 插入数据
	id, err := o.Insert(cate)
	return id, err
}
func AddPostLabelWx(title string, info string, label int16, image string, openid string, name string, sex int32, head string, city string) error {
	o := orm.NewOrm()
	create_time := time.Now()
	my_time := time.Now().Unix()
	beego.Debug("time :", my_time)
	cate := &Post{Title: title, Info: info, CreateTime: create_time, Time: my_time, Label: label, Image: image, OpenId: openid, NickeName: name, Sex: sex, HeadImgurl: head, City: city}
	// 插入数据
	_, err := o.Insert(cate)
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

//根据openid删除
func DeletePostOpenid(id string, openid string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &Post{Id: cid, OpenId: openid}
	_, err = o.Delete(cate)
	return err
}

func QueryLimitPost(nums int64) ([]Post, error) {
	o := orm.NewOrm()
	var posts []Post
	_, err := o.Raw("SELECT * FROM post ORDER BY id DESC LIMIT ? ", nums).QueryRows(&posts)
	return posts, err
}
func QueryPagePost(page int32, nums int32) ([]Post, error) {
	o := orm.NewOrm()
	var posts []Post
	_, err := o.Raw("SELECT * FROM post WHERE examine = 1 ORDER BY id DESC LIMIT ?,? ", page*nums, nums).QueryRows(&posts)
	return posts, err
}
func QueryCityPagePost(page int32, nums int32, city string) ([]Post, error) {
	o := orm.NewOrm()
	var posts []Post
	_, err := o.Raw("SELECT * FROM post WHERE examine = 1 AND city = ? ORDER BY id DESC LIMIT ?,? ", city, page*nums, nums).QueryRows(&posts)
	return posts, err
}
func QueryFuzzyLimitPost(fuzzy string, nums int64) ([]Post, error) {
	o := orm.NewOrm()
	var posts []Post
	_, err := o.Raw("SELECT * FROM post WHERE info LIKE ? ORDER BY id DESC LIMIT ? ", "%"+fuzzy+"%", nums).QueryRows(&posts)
	return posts, err
}

/*
返回帖子数量
*/
func GetPostCount() (int32, error) {
	o := orm.NewOrm()
	// sql := "select count(*) from post where examine = 1"
	// count := 0
	// err := o.Raw(sql).QueryRow(count)
	// if err != nil {

	// }
	count, err := o.QueryTable("post").Filter("examine", 1).Count()
	return int32(count), err
}

/*
根据城市返回帖子数量
*/
func GetCityPostCount(city string) (int32, error) {
	o := orm.NewOrm()
	// sql := "select count(*) from post where examine = 1"
	// count := 0
	// err := o.Raw(sql).QueryRow(count)
	// if err != nil {

	// }
	count, err := o.QueryTable("post").Filter("examine", 1).Filter("city", city).Count()
	return int32(count), err
}

/*******************意见反馈********************/

func AddFeedback(info string, openid string, name string, sex int32, head string) error {
	o := orm.NewOrm()
	create_time := time.Now()
	my_time := time.Now().Unix()
	cate := &Feedback{Info: info, CreateTime: create_time, Time: my_time, OpenId: openid, NickeName: name, Sex: sex, HeadImgurl: head}
	// 插入数据
	_, err := o.Insert(cate)
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

func AddPublicNumber(title string, info string, num string, evaluate string, image string) error {
	o := orm.NewOrm()
	my_time := time.Now().Unix()
	cate := &Wxnum{Title: title, Info: info, Num: num, Evaluate: evaluate, Image: image, Time: my_time, State: int8(0)}

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
func DeleteWxnum(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &Wxnum{Id: cid}
	_, err = o.Delete(cate)
	return err
}

func UpdateWxnum(id string, state int8) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	obj := &Wxnum{Id: cid}
	obj.State = state
	_, err = o.Update(obj, "state")
	return err
}

func GetAllWxnums() ([]Wxnum, error) {
	o := orm.NewOrm()
	var wxnums []Wxnum
	_, err := o.Raw("SELECT * FROM wxnum  ORDER BY id DESC").QueryRows(&wxnums)
	return wxnums, err
}
func GetAllWxnumsState1() ([]Wxnum, error) {
	o := orm.NewOrm()
	var objs []Wxnum
	_, err := o.Raw("SELECT * FROM wxnum  WHERE state = ? ORDER BY id DESC", 1).QueryRows(&objs)
	return objs, err
}

func GetOneWxnum(id string) (*Wxnum, error) {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	obj := &Wxnum{Id: cid}
	err = o.Read(obj)
	return obj, err
}

func UpdateWxnumInfo(id string, title string, info string, number string, evaluate string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &Wxnum{Id: cid}
	cate.Title = title
	cate.Info = info
	cate.Num = number
	cate.Evaluate = evaluate
	_, err = o.Update(cate, "title", "info", "num", "evaluate")
	return err
}

func UpdateWxnumImg(id string, img string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &Wxnum{Id: cid}
	cate.Image = img
	_, err = o.Update(cate, "image")
	return err
}

/***********************推荐公众号**********************/
//添加微信公众号
func AddWeixinNumber(name string, info string, num string, evaluate string, image string) error {
	o := orm.NewOrm()
	my_time := time.Now().Unix()
	cate := &WeixinNumber{Name: name, Info: info, Num: num, Evaluate: evaluate, Image: image, Time: my_time, State: int8(0)}

	// 查询数据
	qs := o.QueryTable("weixin_number")
	err := qs.Filter("num", num).One(cate)
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
func DeleteWeixinNumber(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &WeixinNumber{Id: cid}
	_, err = o.Delete(cate)
	return err
}

func UpdateWeixinNumber(id string, state int8) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	obj := &WeixinNumber{Id: cid}
	obj.State = state
	_, err = o.Update(obj, "state")
	return err
}

func GetAllWeixinNumbers() ([]WeixinNumber, error) {
	o := orm.NewOrm()
	var wxnums []WeixinNumber
	_, err := o.Raw("SELECT * FROM weixin_number  ORDER BY id DESC").QueryRows(&wxnums)
	return wxnums, err
}
func GetAllWeixinNumbersState1() ([]WeixinNumber, error) {
	o := orm.NewOrm()
	var objs []WeixinNumber
	_, err := o.Raw("SELECT * FROM weixin_number  WHERE state = ? ORDER BY id DESC", 1).QueryRows(&objs)
	return objs, err
}

func GetOneWeixinNumber(id string) (*WeixinNumber, error) {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	obj := &WeixinNumber{Id: cid}
	err = o.Read(obj)
	return obj, err
}

func UpdateWeixinNumberInfo(id string, name string, info string, number string, evaluate string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &WeixinNumber{Id: cid}
	cate.Name = name
	cate.Info = info
	cate.Num = number
	cate.Evaluate = evaluate
	_, err = o.Update(cate, "name", "info", "num", "evaluate")
	return err
}

func UpdateWeixinNumberImg(id string, img string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &WeixinNumber{Id: cid}
	cate.Image = img
	_, err = o.Update(cate, "image")
	return err
}

/***********************推荐公众号**********************/
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

//--------------------微信用户-------------
func AddWxUserInfo(wxUserInfo Wxuserinfo) error {
	beego.Debug("-----------AddWxUserInfo----------")
	beego.Debug(wxUserInfo)
	o := orm.NewOrm()
	cate := &Wxuserinfo{OpenId: wxUserInfo.OpenId, NickeName: wxUserInfo.NickeName, Sex: wxUserInfo.Sex,
		Province: wxUserInfo.Province, City: wxUserInfo.City, Country: wxUserInfo.Country,
		HeadImgurl: wxUserInfo.HeadImgurl, Unionid: wxUserInfo.Unionid,
		ErrCode: wxUserInfo.ErrCode, ErrMsg: wxUserInfo.ErrMsg}

	// 查询数据
	qs := o.QueryTable("wxuserinfo")
	err := qs.Filter("open_id", wxUserInfo.OpenId).One(cate)
	if err == nil { //存在则更新
		// beego.Debug("cate:", cate)
		cate.NickeName = wxUserInfo.NickeName
		cate.Sex = wxUserInfo.Sex
		cate.HeadImgurl = wxUserInfo.HeadImgurl
		_, err = o.Update(cate, "nicke_name", "sex", "head_imgurl")

		return err
	}

	// 插入数据
	_, err = o.Insert(cate)
	if err != nil {
		return err
	}

	return nil
}

func SunscribeWxUserInfo(wxUserInfo Wxuserinfo) (int, error) {
	beego.Debug("-----------AddWxUserInfo----------")
	beego.Debug(wxUserInfo)
	o := orm.NewOrm()
	cate := &Wxuserinfo{OpenId: wxUserInfo.OpenId, NickeName: wxUserInfo.NickeName, Sex: wxUserInfo.Sex,
		Province: wxUserInfo.Province, City: wxUserInfo.City, Country: wxUserInfo.Country,
		HeadImgurl: wxUserInfo.HeadImgurl, Unionid: wxUserInfo.Unionid,
		ErrCode: wxUserInfo.ErrCode, ErrMsg: wxUserInfo.ErrMsg}

	// 查询数据
	qs := o.QueryTable("wxuserinfo")
	err := qs.Filter("open_id", wxUserInfo.OpenId).One(cate)
	if err == nil {
		return 1, err
	}

	// 插入数据
	_, err = o.Insert(cate)
	if err != nil {
		return 0, err
	}

	return 0, nil
}

/*
添加用户金钱
*/
func AddWxUserMoney(openid string, money int64) error {
	beego.Debug("AddWxUserMoney openid:", openid)
	o := orm.NewOrm()
	cate := &Wxuserinfo{}
	// 查询数据
	qs := o.QueryTable("wxuserinfo")
	err := qs.Filter("open_id", openid).One(cate)
	if err != nil {
		return err
	}
	beego.Debug("AddWxUserMoney Id:", cate.Id)
	// cate = &Wxuserinfo{Id: cate.Id}
	cate.Money = cate.Money + money
	_, err = o.Update(cate, "money")
	return err
}

/*
消耗金钱
*/
func ConsumeWxUserMoney(openid string, money int64) error {
	beego.Debug("ConsumeWxUserMoney openid:", openid)
	o := orm.NewOrm()
	cate := &Wxuserinfo{}
	// 查询数据
	qs := o.QueryTable("wxuserinfo")
	err := qs.Filter("open_id", openid).One(cate)
	if err != nil {
		return err
	}
	beego.Debug("ConsumeWxUserMoney Id:", cate.Id)
	// cate = &Wxuserinfo{Id: cate.Id}
	cate.Money = cate.Money - money
	if cate.Money < 0 {
		cate.Money = 0
	}
	_, err = o.Update(cate, "money")
	return err
}

/*
修改用户关注状态
*/
func UpWxUserSubscribe(openid string, subscribe_s string) error {
	subscribe, err := strconv.ParseInt(subscribe_s, 10, 64)
	if err != nil {
		return err
	}
	beego.Debug("UpWxUserMoney openid:", openid, "subscribe_s:", subscribe_s)
	o := orm.NewOrm()
	cate := &Wxuserinfo{}
	// 查询数据
	qs := o.QueryTable("wxuserinfo")
	err = qs.Filter("open_id", openid).One(cate)
	if err != nil {
		return err
	}
	beego.Debug("UpWxUserMoney Id:", cate.Id)
	// cate = &Wxuserinfo{Id: cate.Id}
	cate.Subscribe = int32(subscribe)
	_, err = o.Update(cate, "subscribe")
	return err
}

/*
修改用户金钱
*/
func UpWxUserMoney(openid string, money_s string) error {
	money, err := strconv.ParseInt(money_s, 10, 64)
	if err != nil {
		return err
	}
	beego.Debug("UpWxUserMoney openid:", openid, "money_s:", money_s)
	o := orm.NewOrm()
	cate := &Wxuserinfo{}
	// 查询数据
	qs := o.QueryTable("wxuserinfo")
	err = qs.Filter("open_id", openid).One(cate)
	if err != nil {
		return err
	}
	beego.Debug("UpWxUserMoney Id:", cate.Id)
	// cate = &Wxuserinfo{Id: cate.Id}
	cate.Money = money
	_, err = o.Update(cate, "money")
	return err
}

func GetOneWxUserInfo(open_id string) (*Wxuserinfo, error) {
	o := orm.NewOrm()
	var wxusers []Wxuserinfo
	_, err := o.Raw("SELECT * FROM wxuserinfo WHERE open_id = ? ", open_id).QueryRows(&wxusers)
	wxuser := &Wxuserinfo{}
	if len(wxusers) > 0 {
		wxuser = &wxusers[0]
	}
	return wxuser, err
}
func GetOneWxUserInfoId(id string) (*Wxuserinfo, error) {
	o := orm.NewOrm()
	var wxusers []Wxuserinfo
	_, err := o.Raw("SELECT * FROM wxuserinfo WHERE id = ? ", id).QueryRows(&wxusers)
	wxuser := &Wxuserinfo{}
	if len(wxusers) > 0 {
		wxuser = &wxusers[0]
	}
	return wxuser, err
}

func GetAllWxUsers() ([]Wxuserinfo, error) {
	o := orm.NewOrm()
	var wxusers []Wxuserinfo
	_, err := o.Raw("SELECT * FROM wxuserinfo  ORDER BY id DESC").QueryRows(&wxusers)
	return wxusers, err
}

func DeleteWxUser(id int64) error {
	o := orm.NewOrm()
	obj := &Wxuserinfo{Id: id}
	_, err := o.Delete(obj)
	return err
}

func AddNewsKey(info string) error {
	o := orm.NewOrm()

	time := time.Now()
	cate := &NewsKey{Info: info, CreateTime: time, Op: 0}
	// 查询数据
	qs := o.QueryTable("news_key")
	err := qs.Filter("info", info).One(cate)
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

func GetAllNewsKey() ([]NewsKey, error) {
	o := orm.NewOrm()
	var newskey []NewsKey
	_, err := o.Raw("SELECT * FROM news_key  ORDER BY id DESC").QueryRows(&newskey)
	return newskey, err
}

func UpdateNewsKey(id string, op int32) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &NewsKey{Id: cid}
	cate.Op = op
	_, err = o.Update(cate, "op")
	return err
}

func DeleteNewsKey(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &NewsKey{Id: cid}
	_, err = o.Delete(cate)
	return err
}

func GetOpNewsKey(op int32) ([]NewsKey, error) {
	o := orm.NewOrm()
	var newskey []NewsKey
	_, err := o.Raw("SELECT * FROM news_key WHERE op = ? ORDER BY id DESC", 1).QueryRows(&newskey)
	return newskey, err
}

func GetQueryIndex(userid string) (int32, error) {
	o := orm.NewOrm()
	my_time := time.Now().Unix()
	beego.Debug("my_time :", my_time)
	cate := &QureyUser{UserId: userid, Time: my_time, Index: 0}
	// 查询数据
	qs := o.QueryTable("qurey_user")
	err := qs.Filter("user_id", userid).One(cate)
	if err != nil {
		beego.Debug("err :", err)
		// 插入数据
		id, err := o.Insert(cate)
		if err != nil {
			return 0, err
		}
		beego.Debug("id :", id)
	} else {
		index := cate.Index
		_, _, day := time.Unix(cate.Time, 0).Date()
		beego.Debug("day :", day)
		_, _, new_day := time.Unix(my_time, 0).Date()
		beego.Debug("new_day :", new_day)
		if day != new_day { //次日清零
			index = 0
		} else {
			index = index + 1
		}
		cate.Index = index
		cate.Time = my_time
		_, err = o.Update(cate, "time", "index")
		if err != nil {
			beego.Debug("err :", err)
		}
		beego.Debug("cate :", cate)
	}

	return cate.Index, nil
}
func AddGuanggao(title string, info string, image string, blink bool, link string, bimg bool, item0 string, item1 string, item2 string) (int64, error) {
	o := orm.NewOrm()
	my_time := time.Now().Unix()
	cate := &Guanggao{Title: title, Content: info, Time: my_time, State: 0, Image: image, Blink: blink, Link: link, BImage: bimg, ImageItem0: item0, ImageItem1: item1, ImageItem2: item2}
	// 插入数据
	id, err := o.Insert(cate)
	return id, err
}

func GetAllGuanggaos() ([]Guanggao, error) {
	o := orm.NewOrm()
	var guanggaos []Guanggao
	_, err := o.Raw("SELECT * FROM guanggao  ORDER BY id DESC").QueryRows(&guanggaos)
	return guanggaos, err
}
func GetAllGuanggaosState1() ([]Guanggao, error) {
	o := orm.NewOrm()
	var guanggaos []Guanggao
	_, err := o.Raw("SELECT * FROM guanggao  WHERE state = ? ORDER BY id DESC", 1).QueryRows(&guanggaos)
	return guanggaos, err
}
func UpdateGuanggao(id string, exa int32) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &Guanggao{Id: cid}
	cate.State = exa
	_, err = o.Update(cate, "state")
	return err
}
func DeleteGuanggao(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	guanggao := &Guanggao{Id: cid}
	_, err = o.Delete(guanggao)
	return err
}
func GetOneGuanggao(id string) (*Guanggao, error) {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	guanggao := &Guanggao{Id: cid}
	err = o.Read(guanggao)
	return guanggao, err
}
func UpdateGuanggaoImg(id string, img string, bimg bool, item0 string, item1 string, item2 string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &Guanggao{Id: cid}
	cate.Image = img
	cate.BImage = bimg
	cate.ImageItem0 = item0
	cate.ImageItem1 = item1
	cate.ImageItem2 = item2
	_, err = o.Update(cate, "image", "b_image", "image_item0", "image_item1", "image_item2")
	return err
}
func UpdateGuanggaoInfo(id string, title string, info string, blink bool, link string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &Guanggao{Id: cid}
	cate.Title = title
	cate.Content = info
	cate.Blink = blink
	cate.Link = link
	_, err = o.Update(cate, "title", "content", "blink", "link")
	return err
}

/***********餐厅**********/
//添加餐厅
func AddCanting(name string, address string, image string, state int8, phone string, starth string, startm string, endh string, endm string) (int64, error) {
	o := orm.NewOrm()
	my_time := time.Now().Unix()
	starth_i, _ := strconv.ParseInt(starth, 10, 64)
	startm_i, _ := strconv.ParseInt(startm, 10, 64)
	endh_i, _ := strconv.ParseInt(endh, 10, 64)
	endm_i, _ := strconv.ParseInt(endm, 10, 64)
	obj := &Canting{Name: name, Address: address, Time: my_time, State: state, Image: image, Phonenumber: phone, Starthour: int8(starth_i), Startminute: int8(startm_i), Endhour: int8(endh_i), Endminute: int8(endm_i)}
	// 插入数据
	id, err := o.Insert(obj)
	return id, err
}
func GetAllCanting() ([]Canting, error) {
	o := orm.NewOrm()
	var objs []Canting
	_, err := o.Raw("SELECT * FROM canting ORDER BY id DESC").QueryRows(&objs)
	return objs, err
}
func GetAllCantingState1() ([]Canting, error) {
	o := orm.NewOrm()
	var objs []Canting
	_, err := o.Raw("SELECT * FROM canting WHERE state = 1 ORDER BY id DESC").QueryRows(&objs)
	return objs, err
}
func DeleteCanting(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	obj := &Canting{Id: cid}
	_, err = o.Delete(obj)
	return err
}
func UpdateCanting(id string, state int8) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	obj := &Canting{Id: cid}
	obj.State = state
	_, err = o.Update(obj, "state")
	return err
}
func GetOneCanting(id string) (*Canting, error) {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	obj := &Canting{Id: cid}
	err = o.Read(obj)
	return obj, err
}

//菜单
func AddCaidan(fid string, name string, info string, image string, state int8, mtype string, price string) (int64, error) {
	fid_i, err := strconv.ParseInt(fid, 10, 64)
	if err != nil {

	}
	o := orm.NewOrm()
	my_time := time.Now().Unix()
	obj := &Caidan{Fid: fid_i, Name: name, Info: info, Time: my_time, State: state, Image: image, Mtype: mtype, Price: price}
	// 插入数据
	id, err := o.Insert(obj)
	return id, err
}
func GetAllCaidan(fid string) ([]Caidan, error) {
	o := orm.NewOrm()
	var objs []Caidan
	_, err := o.Raw("SELECT * FROM caidan WHERE fid = ? ORDER BY id DESC", fid).QueryRows(&objs)
	return objs, err
}
func DeleteCaidan(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	obj := &Caidan{Id: cid}
	_, err = o.Delete(obj)
	return err
}

//添加帮助
func AddPosthelp(postid int64, openid string, state int32) (int64, error) {
	// postid_i, err := strconv.ParseInt(postid, 10, 64)
	// if err != nil {

	// }
	o := orm.NewOrm()
	my_time := time.Now().Unix()
	obj := &Posthelp{PostId: postid, OpenId: openid, State: state, Time: my_time}
	// 插入数据
	id, err := o.Insert(obj)
	return id, err
}

/*
帖子的帮助数
*/
func GatPostHelpNum(postid string) (int, error) {
	o := orm.NewOrm()
	var objs []Posthelp
	_, err := o.Raw("SELECT * FROM posthelp WHERE post_id = ? ORDER BY id DESC", postid).QueryRows(&objs)
	return len(objs), err

}

func GatPaseHelpState(postid string, openid string) (int32, error) {
	o := orm.NewOrm()
	var objs []Posthelp
	_, err := o.Raw("SELECT * FROM posthelp WHERE post_id = ? AND open_id = ? ORDER BY id DESC", postid, openid).QueryRows(&objs)
	if len(objs) != 0 {
		return objs[0].State, err
	} else {
		return 0, err
	}
}

/*
增加金钱事件
*/
func AddUserMoneyRecord(openid string, money int64, money_type int64) (int64, error) {
	beego.Debug("AddUserMoneyRecord openid:", openid, "money:", money, "money_type:", money_type)
	o := orm.NewOrm()
	my_time := time.Now().Unix()
	obj := &UserMoneyRecord{OpenId: openid, MoneyType: money_type, Money: money, Time: my_time}
	// 插入数据
	id, err := o.Insert(obj)
	return id, err
}
func GetAllUserMoneyRecord(openid string) ([]UserMoneyRecord, error) {
	beego.Debug("GetAllUserMoneyRecord openid:", openid)
	o := orm.NewOrm()
	var objs []UserMoneyRecord
	_, err := o.Raw("SELECT * FROM user_money_record WHERE open_id = ? ORDER BY id DESC", openid).QueryRows(&objs)
	return objs, err
}
func DeleteUserMoneyRecord(openid string) error {
	beego.Debug("DeleteUserMoneyRecord openid:", openid)
	o := orm.NewOrm()
	_, err := o.QueryTable("user_money_record").Filter("open_id", openid).Delete()
	return err
}

/*******商品**********/
/*
添加商品
*/
func AddCommodity(name string, info string, image string, money string) error {
	imoney, err := strconv.ParseInt(money, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	create_time := time.Now()
	my_time := time.Now().Unix()
	cate := &Commodity{Name: name, Info: info, CreateTime: create_time, Time: my_time, Image: image, Money: imoney}
	// 查询数据
	qs := o.QueryTable("commodity")
	err = qs.Filter("name", name).One(cate)
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

/*
获得一个商品
*/
func GetOneCommodity(id string) (*Commodity, error) {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	obj := &Commodity{Id: cid}
	err = o.Read(obj)
	return obj, err
}
func GetOneCommodity1(id int64) (*Commodity, error) {
	o := orm.NewOrm()
	obj := &Commodity{Id: id}
	err := o.Read(obj)
	return obj, err
}

/*
后台商品
*/
func GetAllCommoditysAdmin() ([]Commodity, error) {
	o := orm.NewOrm()
	var commoditys []Commodity
	_, err := o.Raw("SELECT * FROM commodity  ORDER BY id DESC").QueryRows(&commoditys)
	beego.Debug("GetAllCommoditysAdmin", commoditys)
	return commoditys, err
}

/*
用户查看商品
*/
func GetAllCommoditys() ([]Commodity, error) {
	o := orm.NewOrm()
	var commoditys []Commodity
	_, err := o.Raw("SELECT * FROM commodity WHERE state = ? ORDER BY id DESC ", 1).QueryRows(&commoditys)
	return commoditys, err
}

/*
删除商品
*/
func DeleteCommodity(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &Commodity{Id: cid}
	_, err = o.Delete(cate)
	return err
}

//修改商品内容
func UpdateCommodityInfo(id string, name string, info string, money string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	imoney, err := strconv.ParseInt(money, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &Commodity{Id: cid}
	cate.Name = name
	cate.Info = info
	cate.Money = imoney
	_, err = o.Update(cate, "name", "info", "money")
	return err
}

//修改商品图片
func UpdateCommodityimg(id string, img string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &Commodity{Id: cid}
	cate.Image = img
	_, err = o.Update(cate, "image")
	return err
}

//修改商品状态
func UpdateCommodityState(id string, state int16) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &Commodity{Id: cid}
	cate.State = state
	_, err = o.Update(cate, "state")
	return err
}

/*******商品**********/

/*******订单**********/
/*
添加订单
*/
func AddUorder(openid string, commid string) error {
	cid, err := strconv.ParseInt(commid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	create_time := time.Now()
	my_time := time.Now().Unix()
	cate := &Uorder{OpenId: openid, CommodityId: cid, CreateTime: create_time, Time: my_time, ExchangeTime: create_time, Time1: my_time}
	// 插入数据
	_, err = o.Insert(cate)
	if err != nil {
		return err
	}
	return nil
}

/*
获得一个订单
*/
func GetOneUorder(id string) (*Uorder, error) {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	obj := &Uorder{Id: cid}
	err = o.Read(obj)
	return obj, err
}

/*
用户查看订单
*/
func GetAllUorders() ([]Uorder, error) {
	o := orm.NewOrm()
	var objs []Uorder
	_, err := o.Raw("SELECT * FROM uorder  ORDER BY id DESC").QueryRows(&objs)
	beego.Debug("GetAllUorders", objs)
	return objs, err
}

func GetAllUserUorders(openid string) ([]Uorder, error) {
	o := orm.NewOrm()
	var objs []Uorder
	_, err := o.Raw("SELECT * FROM uorder WHERE open_id = ? ORDER BY id DESC", openid).QueryRows(&objs)
	beego.Debug("GetAllUorders", objs)
	return objs, err
}
func UpdateUorderState(id string, state int16) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &Uorder{Id: cid}
	cate.State = state
	_, err = o.Update(cate, "state")
	return err
}

/*******订单**********/
/*******商户**********/
/*
添加商户
*/
func AddShangHu(name string, info string, image string, shanghu_type int16) error {
	o := orm.NewOrm()
	create_time := time.Now()
	my_time := time.Now().Unix()
	cate := &ShangHu{Name: name, Info: info, Image: image, CreateTime: create_time, Time: my_time, Type: shanghu_type}
	// 插入数据
	_, err := o.Insert(cate)
	if err != nil {
		return err
	}
	return nil
}

func GetOneShanghu(id string) (*ShangHu, error) {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	obj := &ShangHu{Id: cid}
	err = o.Read(obj)
	return obj, err
}

/*
用户查看订单
*/
func GetAdminAllShangHus() ([]ShangHu, error) {
	o := orm.NewOrm()
	var objs []ShangHu
	_, err := o.Raw("SELECT * FROM shang_hu  ORDER BY id DESC").QueryRows(&objs)
	beego.Debug("GetAllUorders", objs)
	return objs, err
}

func GetAllShangHus() ([]ShangHu, error) {
	o := orm.NewOrm()
	var objs []ShangHu
	_, err := o.Raw("SELECT * FROM shang_hu WHERE state = ? ORDER BY id DESC", 1).QueryRows(&objs)
	beego.Debug("GetAllUorders", objs)
	return objs, err
}
func GetAllTypeShangHus(mytype string) ([]ShangHu, error) {
	o := orm.NewOrm()
	var objs []ShangHu
	_, err := o.Raw("SELECT * FROM shang_hu WHERE state = ? And type = ? ORDER BY id DESC", 1, mytype).QueryRows(&objs)
	beego.Debug("GetAllTypeShangHus mytype ", mytype, objs)
	return objs, err
}
func UpdateShangHuState(id string, state int16) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &ShangHu{Id: cid}
	cate.State = state
	_, err = o.Update(cate, "state")
	return err
}
func UpdateShangHuRecommend(id string, recommend int16) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &ShangHu{Id: cid}
	cate.Recommend = recommend
	_, err = o.Update(cate, "recommend")
	beego.Debug("UpdateShangHuRecommend :", recommend)
	return err
}

func UpdateShangHuInfo(id string, name string, info string, sh_type int16) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &ShangHu{Id: cid}
	cate.Name = name
	cate.Info = info
	cate.Type = sh_type
	_, err = o.Update(cate, "name", "info", "type")
	return err
}
func UpdateShangHuImg(id string, img string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &ShangHu{Id: cid}
	cate.Image = img
	_, err = o.Update(cate, "image")
	return err
}

/*
删除商户
*/
func DeleteShangHu(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &ShangHu{Id: cid}
	_, err = o.Delete(cate)
	return err
}

/*******商户**********/
/*******关键字**********/
/*
添加关键字
*/
func AddKeywords(key string) error {
	o := orm.NewOrm()
	create_time := time.Now()
	my_time := time.Now().Unix()
	cate := &Keywords{KeyName: key, CreateTime: create_time, Time: my_time}
	// 插入数据
	_, err := o.Insert(cate)
	if err != nil {
		return err
	}
	return nil
}

/*
获得关键字列表
*/
func GetAllKeywords() ([]Keywords, error) {
	o := orm.NewOrm()
	var objs []Keywords
	_, err := o.Raw("SELECT * FROM keywords  ORDER BY id DESC").QueryRows(&objs)
	beego.Debug("GetAllUorders", objs)
	return objs, err
}

/*
修改关键字
*/
func UpdateKeywordsState(id string, state int16) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &Keywords{Id: cid}
	cate.State = state
	_, err = o.Update(cate, "state")
	return err
}

/*
删除商户
*/
func DeleteKeywords(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &Keywords{Id: cid}
	_, err = o.Delete(cate)
	return err
}

/*
获得关键字数量
*/

func GetKeywordsCount(key string) (int32, error) {
	o := orm.NewOrm()
	var objs []Keywords
	num, err := o.Raw("SELECT * FROM keywords WHERE key_name = ?", key).QueryRows(&objs)
	if err != nil {
		beego.Error(err)
		return int32(0), err
	}
	return int32(num), err
}

func GetOneKeywords(key string) (*Keywords, error) {
	o := orm.NewOrm()
	keywords := &Keywords{}
	err := o.QueryTable("keywords").Filter("key_name", key).One(keywords)
	if err != nil {
		beego.Error(err)
		return nil, err
	}
	return keywords, err
}

/*******关键字**********/

/********关键字对象*********/
/*
添加关键字对象
*/
func AddKeyobj(keyid string, title string, info string, img string, url string) error {
	ckeyid, err := strconv.ParseInt(keyid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	create_time := time.Now()
	my_time := time.Now().Unix()
	cate := &Keyobj{KeyId: ckeyid, Title: title, Info: info, Image: img, Url: url, CreateTime: create_time, Time: my_time}
	// 插入数据
	_, err = o.Insert(cate)
	if err != nil {
		return err
	}
	return nil
}

/*
获得关键字列表
*/
func GetAllKeyobj(keyid string) ([]Keyobj, error) {
	ckeyid, err := strconv.ParseInt(keyid, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	var objs []Keyobj
	_, err = o.Raw("SELECT * FROM keyobj WHERE key_id = ? ORDER BY id DESC", ckeyid).QueryRows(&objs)
	beego.Debug("GetAllUorders", objs)
	return objs, err
}

/*
修改关键字
*/
func UpdateKeyobjState(id string, state int16) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &Keyobj{Id: cid}
	cate.State = state
	_, err = o.Update(cate, "state")
	return err
}

/*
删除商户
*/
func DeleteKeyobj(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &Keyobj{Id: cid}
	_, err = o.Delete(cate)
	return err
}

/*
关键字查询
*/
func QueryFuzzyLimitKeyobj(keyid int64, nums int64) ([]Keyobj, error) {
	o := orm.NewOrm()
	var objs []Keyobj
	_, err := o.Raw("SELECT * FROM keyobj WHERE key_id = ? AND state = 1 ORDER BY id DESC LIMIT ? ", keyid, nums).QueryRows(&objs)
	return objs, err
}

/*******关键字对象**********/
/********投票组********/
/**
添加投票组
*/
func AddPolls(title string, info string, img string, more string, endtime int64, appid string, secret string, prize string, ext string) error {
	// 	StartTime     time.Time `orm:"index"` //投票组开始时间
	// StartTimeLong int64
	// EndTime       time.Time `orm:"index"` //投票组结束时间
	// EndTimeLong   int64
	o := orm.NewOrm()
	my_time := time.Now().Unix()
	cate := &Polls{Title: title, Info: info, Image: img, More: more, StartTimeLong: my_time, EndTimeLong: endtime, Appid: appid, Secret: secret, Prize: prize, Ext: ext}
	// 插入数据
	_, err := o.Insert(cate)
	if err != nil {
		return err
	}
	return nil
}

/**
修改投票组内容
*/
func UpPollsInfo(id string, title string, info string, more string, endtime int64, appid string, secret string, prize string, ext string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &Polls{Id: cid}
	cate.Title = title
	cate.Info = info
	cate.More = more
	cate.EndTimeLong = endtime
	cate.Appid = appid
	cate.Secret = secret
	cate.Prize = prize
	cate.Ext = ext
	_, err = o.Update(cate, "title", "info", "more", "end_time_long", "appid", "secret", "prize", "ext")
	return err
}

/**
修改投票组图片
*/
func UpPollsImg(id string, img string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &Polls{Id: cid}
	cate.Image = img
	_, err = o.Update(cate, "image")
	return err
}

/**
修改投票组状态
*/
func UpPollsState(id string, state int16) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &Polls{Id: cid}
	cate.State = state
	_, err = o.Update(cate, "state")
	if err != nil {
		beego.Error(err)
	}
	return err
}

/**
获得所有投票组
*/
func GetAllPolls() ([]Polls, error) {
	o := orm.NewOrm()
	var objs []Polls
	_, err := o.Raw("SELECT * FROM polls ORDER BY id DESC").QueryRows(&objs)
	beego.Debug("GetAllUorders", objs)
	return objs, err
}

/**
获得某个投票组
*/
func GetOnePolls(id string) (*Polls, error) {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	obj := &Polls{}
	err = o.QueryTable("polls").Filter("id", cid).One(obj)
	if err != nil {
		beego.Error(err)
		return nil, err
	}
	return obj, err
}

/**
获得访问量
*/
func GetPollsPv(id string) (int64, error) {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	obj := &Polls{}
	err = o.QueryTable("polls").Filter("id", cid).One(obj)
	if err != nil {
		beego.Error(err)
		return 0, err
	}
	return obj.Pageview, err
}

/**
增加访问量
*/
func AddPollsPv(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &Polls{}
	// 查询数据
	qs := o.QueryTable("polls")
	err = qs.Filter("id", cid).One(cate)
	if err != nil {
		return err
	}
	cate.Pageview = cate.Pageview + 1
	_, err = o.Update(cate, "pageview")
	return err
}

/**
删除投票组
*/
func DeletePolls(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &Polls{Id: cid}
	_, err = o.Delete(cate)
	return err
}

/********投票组********/
/********投票********/
/**
添加投票
*/
func AddPoll(openid string, pollsid string, title string, info string, img string, contactway string) error {
	cpollsid, err := strconv.ParseInt(pollsid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	my_time := time.Now().Unix()
	cate := &Poll{OpenId: openid, PollsId: cpollsid, Title: title, Info: info, Image: img, ContactWay: contactway, CreateTime: my_time}
	// 插入数据
	_, err = o.Insert(cate)
	if err != nil {
		return err
	}
	return nil
}

/**
参与用户
*/
func GetPollAllNum(pollsid string) (int32, error) {
	o := orm.NewOrm()
	var objs []Poll
	num, err := o.Raw("SELECT * FROM poll WHERE polls_id = ?", pollsid).QueryRows(&objs)
	if err != nil {
		beego.Error(err)
		return int32(0), err
	}
	return int32(num), err
}

/**
获得参与用户 one
*/
func GetOnePoll(pollsid string, id string) (*Poll, error) {
	cpollsid, err := strconv.ParseInt(pollsid, 10, 64)
	if err != nil {
		return nil, err
	}
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	obj := &Poll{}
	err = o.QueryTable("poll").Filter("id", cid).Filter("polls_id", cpollsid).One(obj)
	if err != nil {
		beego.Error(err)
		return nil, err
	}
	return obj, err
}

/**
获得用户发布的帖子
*/
func GetMyPoll(pollsid string, openid string) ([]Poll, error) {
	o := orm.NewOrm()
	var objs []Poll
	_, err := o.Raw("SELECT * FROM poll WHERE  polls_id = ? AND open_id = ? AND del = 0 ORDER BY id DESC", pollsid, openid).QueryRows(&objs)
	beego.Debug("GetAllPoll", objs)
	return objs, err
}

/**
活动参与用户
*/
func GetAllPoll(pollsid string) ([]Poll, error) {
	o := orm.NewOrm()
	var objs []Poll
	_, err := o.Raw("SELECT * FROM poll WHERE  polls_id = ? ORDER BY id DESC", pollsid).QueryRows(&objs)
	beego.Debug("GetAllPoll", objs)
	return objs, err
}

/**
根据条件搜索
*/
func GetAllPollOr(search string) ([]Poll, error) {
	o := orm.NewOrm()
	var objs []Poll
	_, err := o.Raw("SELECT * FROM poll WHERE  id = ? OR title = ? ORDER BY id DESC", search, search).QueryRows(&objs)
	beego.Debug("GetAllPollOr", objs)
	return objs, err
}

/**
刷新参与用户状态
*/
func UpdatePollState(pollid string, id string, state int16) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	cpollid, err := strconv.ParseInt(pollid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &Poll{Id: cid, PollsId: cpollid}
	cate.State = state
	_, err = o.Update(cate, "state")
	if err != nil {
		beego.Error(err)
	}
	return err
}

/**
活动参与用户已审核过
*/
func GetAllPollState(pollsid string, state int16) ([]Poll, error) {
	o := orm.NewOrm()
	var objs []Poll
	_, err := o.Raw("SELECT * FROM poll WHERE  polls_id = ? AND state = ? ORDER BY id DESC", pollsid, state).QueryRows(&objs)
	beego.Debug("GetAllPoll", objs)
	return objs, err
}

/**
刷新参与用户状态
*/
func DelPoll(pollid string, id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	cpollid, err := strconv.ParseInt(pollid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &Poll{Id: cid, PollsId: cpollid}
	cate.Del = 1
	_, err = o.Update(cate, "del")
	if err != nil {
		beego.Error(err)
	}
	return err
}

/********投票********/
/*********选票********/
/**
投票
*/

func AddVote(pollsid string, pollid string) error {
	cpollsid, err := strconv.ParseInt(pollsid, 10, 64)
	if err != nil {
		return err
	}
	cpollid, err := strconv.ParseInt(pollid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	my_time := time.Now().Unix()
	cate := &Vote{PollsId: cpollsid, PollId: cpollid, CreateTime: my_time}
	// 插入数据
	_, err = o.Insert(cate)
	if err != nil {
		return err
	}
	return nil
}

/**
获得投票
*/

func GetVoteNum(pollsid string, pollid int64) (int32, error) {
	o := orm.NewOrm()
	var objs []Vote
	num, err := o.Raw("SELECT * FROM vote WHERE polls_id = ? AND poll_id = ?", pollsid, pollid).QueryRows(&objs)
	if err != nil {
		beego.Error(err)
		return int32(0), err
	}
	return int32(num), err
}
func GetVoteNum1(pollsid string, pollid string) (int32, error) {
	o := orm.NewOrm()
	var objs []Vote
	_, err := o.Raw("SELECT * FROM vote WHERE polls_id = ? AND poll_id = ?", pollsid, pollid).QueryRows(&objs)
	if err != nil {
		beego.Error(err)
		return int32(0), err
	}
	return int32(len(objs)), err
}

/**
获得累计投票
*/
func GetVoteAllNum(pollsid string) (int32, error) {
	o := orm.NewOrm()
	var objs []Vote
	num, err := o.Raw("SELECT * FROM vote WHERE polls_id = ?", pollsid).QueryRows(&objs)
	if err != nil {
		beego.Error(err)
		return int32(0), err
	}
	return int32(num), err
}

/***********通知*************/
func AddNotice(from string, to string, bsend bool, msg string, ext string, ntype int16) error {
	// 	Id         int64
	// From       string //发送者
	// To         string //接受者
	// ToRead     int16  //接受者是否已读
	// ToDel      int16  //接受者是否删除
	// State      int16  //状态 0未上线 1 上线
	// CreateTime int64  //投票时间
	// Ext        string //消息扩展 1 帖子ID 2贴在ID
	// Type       int16  //消息类型 1发布信息通过审核通知 2帖子存在违规通知
	send := int16(0)
	if bsend {
		send = int16(1)
	}
	o := orm.NewOrm()
	my_time := time.Now().Unix()
	cate := &Notice{FromId: from, ToId: to, Msg: msg, Ext: ext, NType: ntype, State: send, CreateTime: my_time}
	// 插入数据
	_, err := o.Insert(cate)
	if err != nil {
		return err
	}
	return nil
}

/**
修改通知读取状态
*/
func UpUeserNoticeRead(id int64, read int16) error {
	o := orm.NewOrm()
	cate := &Notice{Id: id}
	cate.ToRead = read
	_, err := o.Update(cate, "to_read")
	if err != nil {
		beego.Error(err)
	}
	return err
}

/**
获得所有通知
*/
func GetAllNotice() ([]Notice, error) {
	o := orm.NewOrm()
	var objs []Notice
	_, err := o.Raw("SELECT * FROM notice  ORDER BY id DESC").QueryRows(&objs)
	return objs, err
}

/**
获得用户所有消息
*/
func GetUeserAllNotice(openid string) ([]Notice, error) {
	o := orm.NewOrm()
	var objs []Notice
	_, err := o.Raw("SELECT * FROM notice WHERE to_id = ?  AND to_del = 0 ORDER BY id DESC", openid).QueryRows(&objs)
	return objs, err
}

/**
获得用户未读消息数量
*/
func GetUserNoticeNum(openid string) (int32, error) {
	beego.Debug("GetUserNoticeNum openid :", openid)
	o := orm.NewOrm()
	var objs []Notice
	num, err := o.Raw("SELECT * FROM notice WHERE to_id = ? AND to_read = 0 ORDER BY id DESC", openid).QueryRows(&objs)
	if err != nil {
		beego.Error(err)
		return int32(0), err
	}
	return int32(num), err
}

/**
用户删除通知
*/
func DeleteUserNotice(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &Notice{Id: cid}
	cate.ToDel = 1
	_, err = o.Update(cate, "to_del")
	if err != nil {
		beego.Error(err)
	}
	return err
}

/**
管理员删除
*/
func DeleteAdminUserNotice(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &Notice{Id: cid}
	_, err = o.Delete(cate)
	return err
}
