package model

import (
    "time"
    "fmt"
)

/** JSGEN({type: "model", paged: true})
CREATE TABLE `zbp_post` (
  `log_ID` int(11) NOT NULL AUTO_INCREMENT,
  `log_CateID` int(11) NOT NULL DEFAULT '0',
  `log_AuthorID` int(11) NOT NULL DEFAULT '0',
  `log_Tag` varchar(255) NOT NULL DEFAULT '',
  `log_Status` tinyint(4) NOT NULL DEFAULT '0',
  `log_Type` tinyint(4) NOT NULL DEFAULT '0',
  `log_Alias` varchar(255) NOT NULL DEFAULT '',
  `log_IsTop` int(11) NOT NULL DEFAULT '0',
  `log_IsLock` tinyint(1) NOT NULL DEFAULT '0',
  `log_Title` varchar(255) NOT NULL DEFAULT '',
  `log_Intro` text NOT NULL,
  `log_Content` longtext NOT NULL,
  `log_PostTime` int(11) NOT NULL DEFAULT '0',
  `log_CommNums` int(11) NOT NULL DEFAULT '0',
  `log_ViewNums` int(11) NOT NULL DEFAULT '0',
  `log_Template` varchar(50) NOT NULL DEFAULT '',
  `log_Meta` longtext NOT NULL,
  PRIMARY KEY (`log_ID`),
  KEY `zbp_log_TPISC` (`log_Type`,`log_PostTime`,`log_IsTop`,`log_Status`,`log_CateID`),
  KEY `zbp_log_VTSC` (`log_ViewNums`,`log_Type`,`log_Status`,`log_CateID`)
) ENGINE=InnoDB AUTO_INCREMENT=1275 DEFAULT CHARSET=utf8
JSGEN **/

var _ = time.Now

// ZbpPostConnection ZbpPost连接类型
type ZbpPostConnection func() DBConnect

// DefaultZbpPostConnection DefaultZbpPost默认连接
var DefaultZbpPostConnection ZbpPostConnection

// ZbpPost ZbpPost值类型
type ZbpPost struct {
  LogID int `json:"log_ID"`
  LogCateID int `json:"log_CateID"`
  LogAuthorID int `json:"log_AuthorID"`
  LogTag string `json:"log_Tag"`
  LogStatus int `json:"log_Status"`
  LogType int `json:"log_Type"`
  LogAlias string `json:"log_Alias"`
  LogIsTop int `json:"log_IsTop"`
  LogIsLock int `json:"log_IsLock"`
  LogTitle string `json:"log_Title"`
  LogIntro string `json:"log_Intro"`
  LogContent string `json:"log_Content"`
  LogPostTime int `json:"log_PostTime"`
  LogCommNums int `json:"log_CommNums"`
  LogViewNums int `json:"log_ViewNums"`
  LogTemplate string `json:"log_Template"`
  LogMeta string `json:"log_Meta"`
}

// Add 插入ZbpPost
func (c ZbpPostConnection) Add(model *ZbpPost) (int64, error) {
    sqlStr := "INSERT INTO `zbp_post` (`log_CateID`, `log_AuthorID`, `log_Tag`, `log_Status`, `log_Type`, `log_Alias`, `log_IsTop`, `log_IsLock`, `log_Title`, `log_Intro`, `log_Content`, `log_PostTime`, `log_CommNums`, `log_ViewNums`, `log_Template`, `log_Meta`) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
    result, err := c().Exec(sqlStr, model.LogCateID, model.LogAuthorID, model.LogTag, model.LogStatus, model.LogType, model.LogAlias, model.LogIsTop, model.LogIsLock, model.LogTitle, model.LogIntro, model.LogContent, model.LogPostTime, model.LogCommNums, model.LogViewNums, model.LogTemplate, model.LogMeta)
    if err != nil {
        return 0, err
    } 
    
    return result.LastInsertId()
}

// AddZbpPost 插入ZbpPost
func AddZbpPost(model *ZbpPost) (int64, error) {
    return DefaultZbpPostConnection.Add(model)
}

// Find 查询ZbpPost
func (c ZbpPostConnection) Find(condition string, args ...interface{}) ([]*ZbpPost, error) {
    sqlStr := "SELECT `log_ID`, `log_CateID`, `log_AuthorID`, `log_Tag`, `log_Status`, `log_Type`, `log_Alias`, `log_IsTop`, `log_IsLock`, `log_Title`, `log_Intro`, `log_Content`, `log_PostTime`, `log_CommNums`, `log_ViewNums`, `log_Template`, `log_Meta` FROM `zbp_post`"
    if len(condition) > 0 {
        sqlStr = sqlStr + " WHERE " + condition
    }
    results := make([]*ZbpPost, 0)

	stmt, err := c().Prepare(sqlStr)
	if err != nil {
		return results, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(args...)
    if err != nil {
        return results, err
    } 
    
        defer rows.Close()
        for rows.Next() {
            model := ZbpPost{}
            values := []interface{}{
              new(interface{}),
              new(interface{}),
              new(interface{}),
              new(interface{}),
              new(interface{}),
              new(interface{}),
              new(interface{}),
              new(interface{}),
              new(interface{}),
              new(interface{}),
              new(interface{}),
              new(interface{}),
              new(interface{}),
              new(interface{}),
              new(interface{}),
              new(interface{}),
              new(interface{}),
            }
            rows.Scan(values...)
            if *(values[0].(*interface{})) != nil {
                tmp := int((*(values[0].(*interface{}))).(int64))
                model.LogID = tmp
            }
            if *(values[1].(*interface{})) != nil {
                tmp := int((*(values[1].(*interface{}))).(int64))
                model.LogCateID = tmp
            }
            if *(values[2].(*interface{})) != nil {
                tmp := int((*(values[2].(*interface{}))).(int64))
                model.LogAuthorID = tmp
            }
            if *(values[3].(*interface{})) != nil {
                tmp := string((*(values[3].(*interface{}))).([]uint8))
                model.LogTag = tmp
            }
            if *(values[4].(*interface{})) != nil {
                tmp := int((*(values[4].(*interface{}))).(int64))
                model.LogStatus = tmp
            }
            if *(values[5].(*interface{})) != nil {
                tmp := int((*(values[5].(*interface{}))).(int64))
                model.LogType = tmp
            }
            if *(values[6].(*interface{})) != nil {
                tmp := string((*(values[6].(*interface{}))).([]uint8))
                model.LogAlias = tmp
            }
            if *(values[7].(*interface{})) != nil {
                tmp := int((*(values[7].(*interface{}))).(int64))
                model.LogIsTop = tmp
            }
            if *(values[8].(*interface{})) != nil {
                tmp := int((*(values[8].(*interface{}))).(int64))
                model.LogIsLock = tmp
            }
            if *(values[9].(*interface{})) != nil {
                tmp := string((*(values[9].(*interface{}))).([]uint8))
                model.LogTitle = tmp
            }
            if *(values[10].(*interface{})) != nil {
                tmp := string((*(values[10].(*interface{}))).([]uint8))
                model.LogIntro = tmp
            }
            if *(values[11].(*interface{})) != nil {
                tmp := string((*(values[11].(*interface{}))).([]uint8))
                model.LogContent = tmp
            }
            if *(values[12].(*interface{})) != nil {
                tmp := int((*(values[12].(*interface{}))).(int64))
                model.LogPostTime = tmp
            }
            if *(values[13].(*interface{})) != nil {
                tmp := int((*(values[13].(*interface{}))).(int64))
                model.LogCommNums = tmp
            }
            if *(values[14].(*interface{})) != nil {
                tmp := int((*(values[14].(*interface{}))).(int64))
                model.LogViewNums = tmp
            }
            if *(values[15].(*interface{})) != nil {
                tmp := string((*(values[15].(*interface{}))).([]uint8))
                model.LogTemplate = tmp
            }
            if *(values[16].(*interface{})) != nil {
                tmp := string((*(values[16].(*interface{}))).([]uint8))
                model.LogMeta = tmp
            }
            results = append(results, &model)
        }
    return results, nil
}

// FindZbpPost 查询ZbpPost
func FindZbpPost(condition string, args ...interface{}) ([]*ZbpPost, error) {
    return DefaultZbpPostConnection.Find(condition, args...)
}
// PagedQuery 分页查询ZbpPost
func (c ZbpPostConnection) PagedQuery(condition string, pageSize uint, page uint, args ...interface{}) (totalCount uint, rows []*ZbpPost, err error) {
	sqlStr := "SELECT COUNT(1) as cnt FROM `zbp_post`"
	if len(condition) > 0 {
		sqlStr = sqlStr + " WHERE " + condition
	}

	cr := c().QueryRow(sqlStr, args...)

	err = cr.Scan(&totalCount)
	if err != nil {
		return 0, nil, err
	}
	if page > 0 {
		page = page - 1
	}
	offset := page * pageSize
	if totalCount <= offset {
		return totalCount, []*ZbpPost{}, nil
	}

	if len(condition) == 0 {
		condition = fmt.Sprintf("1=1")
	}
	condition = condition + fmt.Sprintf(" LIMIT %d OFFSET %d", pageSize, offset)
	rows, err = c.Find(condition, args...)
	return
}

// ZbpPostPagedQuery 分页查询ZbpPost
func ZbpPostPagedQuery(condition string, pageSize uint, page uint, args ...interface{}) (totalCount uint, rows []*ZbpPost, err error) {
	return DefaultZbpPostConnection.PagedQuery(condition, pageSize, page, args...)
}

// Get 获取ZbpPost
func (c ZbpPostConnection) Get(condition string, args ...interface{}) (*ZbpPost, error) {
    results, err := c.Find(condition, args...)

    if err != nil {
        return nil, err
    } 
    
    if len(results) > 0 {
        return results[0], nil
    } 
        
    return nil, nil
}


// GetZbpPost 获取ZbpPost
func GetZbpPost(condition string, args ...interface{}) (*ZbpPost, error) {
    return DefaultZbpPostConnection.Get(condition, args...)
}

// Update 更新ZbpPost
func (c ZbpPostConnection) Update(model *ZbpPost) (int64, error) {
    sqlStr := "UPDATE `zbp_post` SET `log_CateID` = ?, `log_AuthorID` = ?, `log_Tag` = ?, `log_Status` = ?, `log_Type` = ?, `log_Alias` = ?, `log_IsTop` = ?, `log_IsLock` = ?, `log_Title` = ?, `log_Intro` = ?, `log_Content` = ?, `log_PostTime` = ?, `log_CommNums` = ?, `log_ViewNums` = ?, `log_Template` = ?, `log_Meta` = ? WHERE `log_ID` = ?"
    result, err := c().Exec(sqlStr, model.LogCateID, model.LogAuthorID, model.LogTag, model.LogStatus, model.LogType, model.LogAlias, model.LogIsTop, model.LogIsLock, model.LogTitle, model.LogIntro, model.LogContent, model.LogPostTime, model.LogCommNums, model.LogViewNums, model.LogTemplate, model.LogMeta, model.LogID)
    if err != nil {
        return 0, err
    }
    return result.RowsAffected()
}

// UpdateZbpPost 更新ZbpPost
func UpdateZbpPost(model *ZbpPost) (int64, error) {
    return DefaultZbpPostConnection.Update(model)
}
