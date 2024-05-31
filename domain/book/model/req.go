package model

type OrderReq struct {
	MemberCode string `form:"member_code"`
	BookCode   string `form:"book_code"`
	Status     string `form:"status"`
}
