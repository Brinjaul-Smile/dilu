package dto

import (
	"time"

	"github.com/baowk/dilu-core/core/base"
)

type SysMemberGetPageReq struct {
	base.ReqPage `search:"-"`
	Status       int    `json:"status" form:"status"` //状态 1正常
	TeamId       int    `json:"teamId"`               //团队id
	DeptPath     string `json:"deptPath"`             //部门路径
	Name         string `json:"name"`                 //名字
	Phone        string `json:"phone"`                //手机号
}

// 会员
type SysMemberDto struct {
	Id       int    `json:"id"`       //主键
	TeamId   int    `json:"teamId"`   //团队id
	UserId   int    `json:"userId"`   //用户id
	Nickname string `json:"nickname"` //昵称
	Name     string `json:"name"`     //姓名
	Phone    string `json:"phone"`    //电话
	DeptPath string `json:"deptPath"` //部门路径
	DeptId   int    `json:"deptId"`   //部门id
	PostTag  int    `json:"postTag"`  //职位标签 1主管 2副主管 3员工
	Status   int    `json:"status"`   //状态 1正常
}

type TeamMemberResp struct {
	Owner     int       `json:"owner"`     //团队拥有者id
	TeamId    int       `json:"teamId"`    //团队id
	TeamName  string    `json:"teamName"`  //团队名
	UserId    int       `json:"userId"`    //用户id
	Nickname  string    `json:"nickname"`  //昵称
	Name      string    `json:"name"`      //姓名
	Phone     string    `json:"phone"`     //电话
	DeptPath  string    `json:"deptPath"`  //部门路径
	DeptId    int       `json:"deptId"`    //部门id
	PostId    int       `json:"postId"`    //职位标签 1主管 2副主管 3员工
	RoleId    int       `json:"role_id"`   //角色id
	Gender    int       `json:"gender"`    //性别
	EntryTime time.Time `json:"entryTime"` //入职时间
	Birthday  time.Time `json:"birthday"`  //入职时间
}
