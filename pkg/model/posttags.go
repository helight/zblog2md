package model

import (
    "time"
    "fmt"
)

/** JSGEN({type: "model", paged: true})
CREATE TABLE `zbp_tag` (
  `tag_ID` int(11) NOT NULL AUTO_INCREMENT,
  `tag_Name` varchar(255) NOT NULL DEFAULT '',
  `tag_Order` int(11) NOT NULL DEFAULT '0',
  `tag_Count` int(11) NOT NULL DEFAULT '0',
  `tag_Alias` varchar(255) NOT NULL DEFAULT '',
  `tag_Intro` text NOT NULL,
  `tag_Template` varchar(50) NOT NULL DEFAULT '',
  `tag_Meta` longtext NOT NULL,
  PRIMARY KEY (`tag_ID`)
) ENGINE=InnoDB AUTO_INCREMENT=116 DEFAULT CHARSET=utf8
JSGEN **/

var _ = time.Now

// ZbpTagConnection ZbpTag连接类型
type ZbpTagConnection func() DBConnect

// DefaultZbpTagConnection DefaultZbpTag默认连接
var DefaultZbpTagConnection ZbpTagConnection

// ZbpTag ZbpTag值类型
type ZbpTag struct {
  TagID int `json:"tag_ID"`
  TagName string `json:"tag_Name"`
  TagOrder int `json:"tag_Order"`
  TagCount int `json:"tag_Count"`
  TagAlias string `json:"tag_Alias"`
  TagIntro string `json:"tag_Intro"`
  TagTemplate string `json:"tag_Template"`
  TagMeta string `json:"tag_Meta"`
}

// Add 插入ZbpTag
func (c ZbpTagConnection) Add(model *ZbpTag) (int64, error) {
    sqlStr := "INSERT INTO `zbp_tag` (`tag_Name`, `tag_Order`, `tag_Count`, `tag_Alias`, `tag_Intro`, `tag_Template`, `tag_Meta`) VALUES(?, ?, ?, ?, ?, ?, ?)"
    result, err := c().Exec(sqlStr, model.TagName, model.TagOrder, model.TagCount, model.TagAlias, model.TagIntro, model.TagTemplate, model.TagMeta)
    if err != nil {
        return 0, err
    } 
    
    return result.LastInsertId()
}

// AddZbpTag 插入ZbpTag
func AddZbpTag(model *ZbpTag) (int64, error) {
    return DefaultZbpTagConnection.Add(model)
}

// Find 查询ZbpTag
func (c ZbpTagConnection) Find(condition string, args ...interface{}) ([]*ZbpTag, error) {
    sqlStr := "SELECT `tag_ID`, `tag_Name`, `tag_Order`, `tag_Count`, `tag_Alias`, `tag_Intro`, `tag_Template`, `tag_Meta` FROM `zbp_tag`"
    if len(condition) > 0 {
        sqlStr = sqlStr + " WHERE " + condition
    }
    results := make([]*ZbpTag, 0)

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
            model := ZbpTag{}
            values := []interface{}{
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
                model.TagID = tmp
            }
            if *(values[1].(*interface{})) != nil {
                tmp := string((*(values[1].(*interface{}))).([]uint8))
                model.TagName = tmp
            }
            if *(values[2].(*interface{})) != nil {
                tmp := int((*(values[2].(*interface{}))).(int64))
                model.TagOrder = tmp
            }
            if *(values[3].(*interface{})) != nil {
                tmp := int((*(values[3].(*interface{}))).(int64))
                model.TagCount = tmp
            }
            if *(values[4].(*interface{})) != nil {
                tmp := string((*(values[4].(*interface{}))).([]uint8))
                model.TagAlias = tmp
            }
            if *(values[5].(*interface{})) != nil {
                tmp := string((*(values[5].(*interface{}))).([]uint8))
                model.TagIntro = tmp
            }
            if *(values[6].(*interface{})) != nil {
                tmp := string((*(values[6].(*interface{}))).([]uint8))
                model.TagTemplate = tmp
            }
            if *(values[7].(*interface{})) != nil {
                tmp := string((*(values[7].(*interface{}))).([]uint8))
                model.TagMeta = tmp
            }
            results = append(results, &model)
        }
    return results, nil
}

// FindZbpTag 查询ZbpTag
func FindZbpTag(condition string, args ...interface{}) ([]*ZbpTag, error) {
    return DefaultZbpTagConnection.Find(condition, args...)
}
// PagedQuery 分页查询ZbpTag
func (c ZbpTagConnection) PagedQuery(condition string, pageSize uint, page uint, args ...interface{}) (totalCount uint, rows []*ZbpTag, err error) {
	sqlStr := "SELECT COUNT(1) as cnt FROM `zbp_tag`"
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
		return totalCount, []*ZbpTag{}, nil
	}

	if len(condition) == 0 {
		condition = fmt.Sprintf("1=1")
	}
	condition = condition + fmt.Sprintf(" LIMIT %d OFFSET %d", pageSize, offset)
	rows, err = c.Find(condition, args...)
	return
}

// ZbpTagPagedQuery 分页查询ZbpTag
func ZbpTagPagedQuery(condition string, pageSize uint, page uint, args ...interface{}) (totalCount uint, rows []*ZbpTag, err error) {
	return DefaultZbpTagConnection.PagedQuery(condition, pageSize, page, args...)
}

// Get 获取ZbpTag
func (c ZbpTagConnection) Get(condition string, args ...interface{}) (*ZbpTag, error) {
    results, err := c.Find(condition, args...)

    if err != nil {
        return nil, err
    } 
    
    if len(results) > 0 {
        return results[0], nil
    } 
        
    return nil, nil
}


// GetZbpTag 获取ZbpTag
func GetZbpTag(condition string, args ...interface{}) (*ZbpTag, error) {
    return DefaultZbpTagConnection.Get(condition, args...)
}

// Update 更新ZbpTag
func (c ZbpTagConnection) Update(model *ZbpTag) (int64, error) {
    sqlStr := "UPDATE `zbp_tag` SET `tag_Name` = ?, `tag_Order` = ?, `tag_Count` = ?, `tag_Alias` = ?, `tag_Intro` = ?, `tag_Template` = ?, `tag_Meta` = ? WHERE `tag_ID` = ?"
    result, err := c().Exec(sqlStr, model.TagName, model.TagOrder, model.TagCount, model.TagAlias, model.TagIntro, model.TagTemplate, model.TagMeta, model.TagID)
    if err != nil {
        return 0, err
    }
    return result.RowsAffected()
}

// UpdateZbpTag 更新ZbpTag
func UpdateZbpTag(model *ZbpTag) (int64, error) {
    return DefaultZbpTagConnection.Update(model)
}
