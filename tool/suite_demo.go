package tool

import (
	"fmt"
	"strings"
	"xorm.io/xorm"
)

type QueryMakerReq struct {
	Title string
}

func QueryMaker(session *xorm.Session, req QueryMakerReq) (*xorm.Session, string) {
	whereSQL := []string{}

	if req.Title != "" {
		session = session.Where("title ilike ?", fmt.Sprintf("%%%s%%", req.Title))
		whereSQL = append(whereSQL, fmt.Sprintf("title ilike '%%%s%%'", req.Title))
	}
	return session, strings.Join(whereSQL, " and ")
}
