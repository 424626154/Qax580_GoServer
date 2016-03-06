package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	"path"
	"qax580go/models"
	"strings"
	"time"
)

type PollController struct {
	beego.Controller
}

/**
投票后台
*/
func (c *PollController) Adminpolls() {
	bool, username := chackAccount(c.Ctx)
	if bool {
		c.Data["isUser"] = bool
		c.Data["User"] = username
	} else {
		c.Redirect("/admin", 302)
		return
	}
	if c.Ctx.Input.IsGet() {
		beego.Debug("Adminpolls Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("Adminpolls Post")
	}

	op := c.Input().Get("op")
	id := c.Input().Get("id")
	beego.Debug("op :", op)
	switch op {
	case "state":
		err := models.UpPollsState(id, 1)
		if err != nil {
			beego.Error(err)
		}
		beego.Debug("state0 id :", id)
		c.Redirect("/poll/adminpolls", 302)
		return
		return
	case "state1":
		err := models.UpPollsState(id, 0)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/poll/adminpolls", 302)
		return
	case "del":
		err := models.DeletePolls(id)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/poll/adminpolls", 302)
		return
	}

	objs, err := models.GetAllPolls()
	if err != nil {
		beego.Error(err)
	}
	c.Data["Objs"] = objs
	c.TplNames = "adminpolls.html"
}

func (c *PollController) AdminUppollsInfo() {
	bool, username := chackAccount(c.Ctx)
	if bool {
		c.Data["isUser"] = bool
		c.Data["User"] = username
	} else {
		c.Redirect("/admin", 302)
		return
	}
	if c.Ctx.Input.IsGet() {
		beego.Debug("AdminUppollsInfo Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("AdminUppollsInfo Post")
		id := c.Input().Get("id")
		title := c.Input().Get("title")
		info := c.Input().Get("info")
		more := c.Input().Get("more")
		endtime := c.Input().Get("endtime")
		appid := c.Input().Get("appid")
		secret := c.Input().Get("secret")
		prize := c.Input().Get("prize")
		ext := c.Input().Get("ext")
		if len(title) != 0 && len(info) != 0 && len(more) != 0 && len(appid) != 0 && len(secret) != 0 {
			beego.Debug("endtime", endtime)
			//获取本地location
			toBeCharge := endtime                                           //待转化为时间戳的字符串 注意 这里的小时和分钟还要秒必须写 因为是跟着模板走的 修改模板的话也可以不写
			timeLayout := "2006-01-02 15:04"                                //转化所需模板
			loc, _ := time.LoadLocation("Local")                            //重要：获取时区
			theTime, _ := time.ParseInLocation(timeLayout, toBeCharge, loc) //使用模板在对应时区转化为time.time类型
			endtimelong := theTime.Unix()                                   //转化为时间戳 类型是int64
			beego.Debug("theTime", theTime)                                 //打印输出theTime 2015-01-01 15:15:00 +0800 CST
			beego.Debug("endtimelong ", endtimelong)                        //打印输出时间戳 1420041600
			t := time.Now().Unix()
			beego.Debug("local time", t)
			if endtimelong < t {
				beego.Error("select end time error")
			} else {
				err := models.UpPollsInfo(id, title, info, more, endtimelong, appid, secret, prize, ext)
				if err != nil {
					beego.Error(err)
				}
			}
			url := "/poll/adminpolls"
			c.Redirect(url, 302)
			return
		}
	}

	id := c.Input().Get("id")
	obj, err := models.GetOnePolls(id)
	if err != nil {
		beego.Error(err)
	}
	beego.Debug("id :", id)
	beego.Debug("obj :", obj)
	c.Data["Id"] = id
	c.Data["Obj"] = obj
	c.TplNames = "adminuppollsinfo.html"
}

func (c *PollController) AdminUppollsImg() {
	bool, username := chackAccount(c.Ctx)
	if bool {
		c.Data["isUser"] = bool
		c.Data["User"] = username
	} else {
		c.Redirect("/admin", 302)
		return
	}
	if c.Ctx.Input.IsGet() {
		beego.Debug("AdminUppollsImg Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("AdminUppollsImg Post")

		id := c.Input().Get("id")
		originalimg := c.Input().Get("originalimg")
		image_name := originalimg
		if len(id) != 0 {
			// 获取附件
			_, fh, err := c.GetFile("image")
			// beego.Debug("上传图片:", fh)
			if err != nil {
				beego.Error(err)
			}
			var attachment string
			if fh != nil {
				// 保存附件
				attachment = fh.Filename
				t := time.Now().Unix()
				str2 := fmt.Sprintf("%d", t)
				s := []string{attachment, str2}
				h := md5.New()
				h.Write([]byte(strings.Join(s, ""))) // 需要加密的字符串
				image_name = hex.EncodeToString(h.Sum(nil))
				beego.Info(image_name) // 输出加密结果
				err = c.SaveToFile("image", path.Join("imagehosting", image_name))
				if err != nil {
					beego.Error(err)
					image_name = originalimg
				}
			}

			beego.Debug("上传前图片", originalimg, "上传后图片", image_name)
			if len(image_name) != 0 {
				err := models.UpPollsImg(id, image_name)
				if err != nil {
					beego.Error(err)
				} else {
					c.Redirect("/poll/adminpolls", 302)
					return
				}
			}
		}
	}
	id := c.Input().Get("id")
	obj, err := models.GetOnePolls(id)
	if err != nil {
		beego.Error(err)
	}
	c.Data["Obj"] = obj
	c.TplNames = "adminuppollsimg.html"
}

func (c *PollController) Adminpollscon() {
	bool, username := chackAccount(c.Ctx)
	if bool {
		c.Data["isUser"] = bool
		c.Data["User"] = username
	} else {
		c.Redirect("/admin", 302)
		return
	}
	pollsid := c.Input().Get("pollsid")
	if c.Ctx.Input.IsGet() {
		beego.Debug("Adminpollscon Get")
		op := c.Input().Get("op")
		beego.Debug("op:", op)
		switch op {
		case "state":
			id := c.Input().Get("id")
			if len(id) == 0 {
				break
			}
			id = c.Input().Get("id")
			err := models.UpdatePollState(pollsid, id, 1)
			if err != nil {
				beego.Error(err)
			}
			url := fmt.Sprintf("/poll/adminpollscon?pollsid=%s", pollsid)
			c.Redirect(url, 302)
			return
		case "state1":
			id := c.Input().Get("id")
			if len(id) == 0 {
				break
			}
			id = c.Input().Get("id")
			err := models.UpdatePollState(pollsid, id, 0)
			if err != nil {
				beego.Error(err)
			}
			url := fmt.Sprintf("/poll/adminpollscon?pollsid=%s", pollsid)
			c.Redirect(url, 302)
			return
		}
	}

	if c.Ctx.Input.IsPost() {
		beego.Debug("Adminpollscon Post")
	}
	objs, err := models.GetAllPoll(pollsid)
	if err != nil {
		beego.Error(err)
	}
	beego.Debug("pollsid:", pollsid)
	beego.Debug("objs:", objs)
	c.Data["Objs"] = objs
	c.Data["PollsId"] = pollsid
	c.TplNames = "adminpollscon.html"
}

/**
投票后台添加新投票
*/
func (c *PollController) Adminaddpoll() {
	bool, username := chackAccount(c.Ctx)
	if bool {
		c.Data["isUser"] = bool
		c.Data["User"] = username
	} else {
		c.Redirect("/admin", 302)
		return
	}
	if c.Ctx.Input.IsGet() {
		beego.Debug("Adminaddpoll Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("Adminaddpoll Post")
		image_name := ""
		title := c.Input().Get("title")
		info := c.Input().Get("info")
		more := c.Input().Get("more")
		endtime := c.Input().Get("endtime")
		appid := c.Input().Get("appid")
		secret := c.Input().Get("secret")
		prize := c.Input().Get("prize")
		ext := c.Input().Get("ext")
		if len(title) != 0 && len(info) != 0 && len(more) != 0 && len(appid) != 0 && len(secret) != 0 {
			// 获取附件
			_, fh, err := c.GetFile("image")
			// beego.Debug("上传图片:", fh)
			if err != nil {
				beego.Error(err)
			}
			var attachment string
			if fh != nil {
				// 保存附件
				attachment = fh.Filename
				t := time.Now().Unix()
				str2 := fmt.Sprintf("%d", t)
				s := []string{attachment, str2}
				h := md5.New()
				h.Write([]byte(strings.Join(s, ""))) // 需要加密的字符串
				image_name = hex.EncodeToString(h.Sum(nil))
				beego.Info(image_name) // 输出加密结果
				err = c.SaveToFile("image", path.Join("imagehosting", image_name))
				if err != nil {
					beego.Error(err)
					image_name = ""
				}
			}
			beego.Debug("endtime", endtime)
			//获取本地location
			toBeCharge := endtime                                           //待转化为时间戳的字符串 注意 这里的小时和分钟还要秒必须写 因为是跟着模板走的 修改模板的话也可以不写
			timeLayout := "2006-01-02 15:04"                                //转化所需模板
			loc, _ := time.LoadLocation("Local")                            //重要：获取时区
			theTime, _ := time.ParseInLocation(timeLayout, toBeCharge, loc) //使用模板在对应时区转化为time.time类型
			endtimelong := theTime.Unix()                                   //转化为时间戳 类型是int64
			beego.Debug("theTime", theTime)                                 //打印输出theTime 2015-01-01 15:15:00 +0800 CST
			beego.Debug("endtimelong ", endtimelong)                        //打印输出时间戳 1420041600
			t := time.Now().Unix()
			beego.Debug("local time", t)
			if endtimelong < t {
				beego.Error("select end time error")
			} else {
				err = models.AddPolls(title, info, image_name, more, endtimelong, appid, secret, prize, ext)
				if err != nil {
					beego.Error(err)
				}
			}
			url := "/poll/adminpolls"
			c.Redirect(url, 302)
			return
		}
	}
	c.TplNames = "adminaddpoll.html"
}

/**
投票主页
*/
func (c *PollController) PollHome() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("PollHome Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("PollHome Post")
	}
	pollsid := c.Input().Get("pollsid")
	beego.Debug("pollsid :", pollsid)
	op := c.Input().Get("op")
	beego.Debug("op :", op)
	switch op {
	case "vote":
		id := c.Input().Get("id")
		err := models.AddVote(pollsid, id)
		if err != nil {
			beego.Debug(err)
		}
		url := fmt.Sprintf("/poll/pollhome?pollsid=%s", pollsid)
		c.Redirect(url, 302)
		return
	}

	if len(pollsid) != 0 {
		err := models.AddPollsPv(pollsid)
		if err != nil {
			beego.Error(err)
		}
		polls, err := models.GetOnePolls(pollsid)
		if err != nil {
			beego.Error(err)
		}
		beego.Debug("polls", polls)
		c.Data["Polls"] = polls
		pv, err := models.GetPollsPv(pollsid)
		c.Data["PV"] = pv
		pollnum, err := models.GetPollAllNum(pollsid)
		c.Data["PollNum"] = pollnum
		votenum, err := models.GetVoteAllNum(pollsid)
		c.Data["VoteNum"] = votenum

		endtime := polls.EndTimeLong
		curtime := time.Now().Unix()
		timestr := "活动已过期"
		if endtime-curtime > 0 {
			t := time.Unix(endtime, 0)

			_, mon, day := t.Date()
			_, cmon, cday := time.Now().Date()
			hour, min, _ := t.Clock()
			chour, cmin, _ := time.Now().Clock()
			beego.Debug("---------")
			timestr = fmt.Sprintf("%d月%d天%02d小时%02d分", mon-cmon, day-cday, hour-chour, min-cmin)
			// beego.Debug(timestr)
		}
		c.Data["TimeStr"] = timestr

		objs, err := models.GetAllPollState(pollsid, 1)
		if err != nil {
			beego.Debug(err)
		}
		for i := 0; i < len(objs); i++ {
			num, err := models.GetVoteNum(pollsid, objs[i].Id)
			if err != nil {
				beego.Error(err)
			}
			objs[i].VoteNum = num
		}
		beego.Debug("objs :", objs)
		c.Data["Objs"] = objs
	}
	c.Data["EndTimeLong"] = 0
	c.Data["PollsId"] = pollsid
	c.TplNames = "pollhome.html"
}

/**
投票详情
*/
func (c *PollController) PollHomeCon() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("PollHome Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("PollHome Post")
	}
	pollsid := c.Input().Get("pollsid")
	pollid := c.Input().Get("pollid")
	beego.Debug("pollsid:", pollsid)
	beego.Debug("pollid:", pollid)
	c.Data["PollsId"] = pollsid
	c.Data["PollId"] = pollid
	op := c.Input().Get("op")
	beego.Debug("op:", op)
	switch op {
	case "vote":
		err := models.AddVote(pollsid, pollid)
		if err != nil {
			beego.Debug(err)
		}
		url := fmt.Sprintf("/poll/pollhomecon?pollsid=%s&pollid=%s", pollsid, pollid)
		beego.Debug("url:", url)
		c.Redirect(url, 302)
		return
	}

	obj, err := models.GetOnePoll(pollsid, pollid)
	if err != nil {
		beego.Error(err)
	}
	num, err := models.GetVoteNum1(pollsid, pollid)
	if err != nil {
		beego.Error(err)
	}
	polls, err := models.GetOnePolls(pollsid)
	if err != nil {
		beego.Error(err)
	}
	beego.Debug("VoteNum", num)
	obj.VoteNum = num
	c.Data["Time"] = polls.EndTimeLong
	c.Data["Polls"] = polls
	c.Data["Obj"] = obj
	c.TplNames = "pollhomecon.html"
}

/**
投票搜索详情
*/
func (c *PollController) PollHomeSearch() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("PollHomeSearch Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("PollHomeSearch Post")
	}
	pollsid := c.Input().Get("pollsid")
	search := c.Input().Get("search")
	beego.Debug("pollsid:", pollsid)
	beego.Debug("search:", search)
	c.Data["PollsId"] = pollsid
	c.Data["Search"] = search
	op := c.Input().Get("op")
	beego.Debug("op:", op)
	switch op {
	case "vote":
		pollid := c.Input().Get("pollid")
		err := models.AddVote(pollsid, pollid)
		if err != nil {
			beego.Debug(err)
		}
		url := fmt.Sprintf("/poll/pollhomesearch?pollsid=%s&search=%s", pollsid, search)
		beego.Debug("url:", url)
		c.Redirect(url, 302)
		return
	}
	objs, err := models.GetAllPollOr(search)
	if err != nil {
		beego.Debug(err)
	}
	for i := 0; i < len(objs); i++ {
		num, err := models.GetVoteNum(pollsid, objs[i].Id)
		if err != nil {
			beego.Error(err)
		}
		objs[i].VoteNum = num
	}
	polls, err := models.GetOnePolls(pollsid)
	if err != nil {
		beego.Error(err)
	}
	c.Data["Time"] = polls.EndTimeLong
	c.Data["Polls"] = polls
	c.Data["Objs"] = objs
	c.TplNames = "pollhomesearch.html"
}

/**
查看排名
*/
func (c *PollController) PollHomeRanking() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("PollHomeRanking Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("PollHomeRanking Post")
	}
	pollsid := c.Input().Get("pollsid")
	beego.Debug("pollsid:", pollsid)
	c.Data["PollsId"] = pollsid
	if len(pollsid) != 0 {
		objs, err := models.GetAllPollState(pollsid, 1)
		if err != nil {
			beego.Debug(err)
		}
		for i := 0; i < len(objs); i++ {
			num, err := models.GetVoteNum(pollsid, objs[i].Id)
			if err != nil {
				beego.Error(err)
			}
			objs[i].VoteNum = num
		}
		for i := 0; i < len(objs); i++ {
			for j := 0; j < len(objs)-i-1; j++ {
				if objs[j].VoteNum < objs[j+1].VoteNum {
					objs[j], objs[j+1] = objs[j+1], objs[j]
				}
			}
		}
		for i := 0; i < len(objs); i++ {
			objs[i].Ranking = int32(i)
		}
		beego.Debug("objs :", objs)
		c.Data["Objs"] = objs
	}
	c.TplNames = "pollhomeranking.html"
}

/**
添加投票
*/
func (c *PollController) AddPoll() {
	openid := getPollCookie(c)
	pollsid := c.Input().Get("pollsid")
	openid = "o3AhEuB_wdTELvlErL4F1Em4Nck4"
	op := c.Input().Get("op")
	switch op {
	case "del":
		id := c.Input().Get("id")
		err := models.DelPoll(pollsid, id)
		if err != nil {
			beego.Error(err)
		}
		beego.Debug("add poll del :", id)
		url := fmt.Sprintf("/poll/addpoll?pollsid=%s", pollsid)
		c.Redirect(url, 302)
		return
	}
	poll, err := models.GetMyPoll(pollsid, openid)
	if err != nil {
		beego.Error(err)
	}
	if len(poll) > 0 {
		c.Data["IsAdd"] = true //是否已经添加过
		mypoll := poll[0]
		num, err := models.GetVoteNum(pollsid, mypoll.Id)
		if err != nil {
			beego.Error(err)
		}
		mypoll.VoteNum = num
		c.Data["Poll"] = mypoll
	} else {
		c.Data["IsAdd"] = false
	}
	if c.Ctx.Input.IsGet() {
		beego.Debug("PollHome Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("PollHome Post")
		image_name := ""
		title := c.Input().Get("title")
		info := c.Input().Get("info")
		contactway := c.Input().Get("contactway")
		if len(pollsid) != 0 && len(title) != 0 && len(info) != 0 && len(contactway) != 0 {
			// 获取附件
			_, fh, err := c.GetFile("image")
			// beego.Debug("上传图片:", fh)
			if err != nil {
				beego.Error(err)
			}
			var attachment string
			if fh != nil {
				// 保存附件
				attachment = fh.Filename
				t := time.Now().Unix()
				str2 := fmt.Sprintf("%d", t)
				s := []string{attachment, str2}
				h := md5.New()
				h.Write([]byte(strings.Join(s, ""))) // 需要加密的字符串
				image_name = hex.EncodeToString(h.Sum(nil))
				beego.Info(image_name) // 输出加密结果
				err = c.SaveToFile("image", path.Join("imagehosting", image_name))
				if err != nil {
					beego.Error(err)
					image_name = ""
				}
			}
			err = models.AddPoll(openid, pollsid, title, info, image_name, contactway)
			if err != nil {
				beego.Error(err)
			}
			url := fmt.Sprintf("/poll/pollhome?pollsid=%s", pollsid)
			c.Redirect(url, 302)
			return
		}
	}
	c.Data["PollsId"] = pollsid
	c.TplNames = "addpoll.html"
}

func getPollCookie(c *PollController) string {
	isUser := false
	openid := c.Ctx.GetCookie(COOKIE_WX_OPENID)
	// beego.Debug("------------openid--------")
	// beego.Debug(openid)
	if len(openid) != 0 {
		wxuser, err := models.GetOneWxUserInfo(openid)
		if err != nil {
			beego.Error(err)
		} else {
			isUser = true
			// beego.Debug("--------------wxuser----------")
			// beego.Debug(wxuser)
			c.Data["WxUser"] = wxuser
		}
	}
	c.Data["isUser"] = isUser
	return openid
}
