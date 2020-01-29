// Copyright 2020 helight Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import (
    "time"
    "fmt"
)

/** JSGEN({type: "model", paged: true})
CREATE TABLE `zbp_category` (
  `cate_ID` int(11) NOT NULL AUTO_INCREMENT,
  `cate_Name` varchar(255) NOT NULL DEFAULT '',
  `cate_Order` int(11) NOT NULL DEFAULT '0',
  `cate_Count` int(11) NOT NULL DEFAULT '0',
  `cate_Alias` varchar(255) NOT NULL DEFAULT '',
  `cate_Intro` text NOT NULL,
  `cate_RootID` int(11) NOT NULL DEFAULT '0',
  `cate_ParentID` int(11) NOT NULL DEFAULT '0',
  `cate_Template` varchar(50) NOT NULL DEFAULT '',
  `cate_LogTemplate` varchar(50) NOT NULL DEFAULT '',
  `cate_Meta` longtext NOT NULL,
  PRIMARY KEY (`cate_ID`),
  KEY `zbp_cate_Order` (`cate_Order`)
) ENGINE=InnoDB AUTO_INCREMENT=102 DEFAULT CHARSET=utf8
JSGEN **/

var _ = time.Now

// ZbpCategoryConnection ZbpCategory连接类型
type ZbpCategoryConnection func() DBConnect

// DefaultZbpCategoryConnection DefaultZbpCategory默认连接
var DefaultZbpCategoryConnection ZbpCategoryConnection

// ZbpCategory ZbpCategory值类型
type ZbpCategory struct {
  CateID int `json:"cate_ID"`
  CateName string `json:"cate_Name"`
  CateOrder int `json:"cate_Order"`
  CateCount int `json:"cate_Count"`
  CateAlias string `json:"cate_Alias"`
  CateIntro string `json:"cate_Intro"`
  CateRootID int `json:"cate_RootID"`
  CateParentID int `json:"cate_ParentID"`
  CateTemplate string `json:"cate_Template"`
  CateLogTemplate string `json:"cate_LogTemplate"`
  CateMeta string `json:"cate_Meta"`
}

// Add 插入ZbpCategory
func (c ZbpCategoryConnection) Add(model *ZbpCategory) (int64, error) {
    sqlStr := "INSERT INTO `zbp_category` (`cate_Name`, `cate_Order`, `cate_Count`, `cate_Alias`, `cate_Intro`, `cate_RootID`, `cate_ParentID`, `cate_Template`, `cate_LogTemplate`, `cate_Meta`) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
    result, err := c().Exec(sqlStr, model.CateName, model.CateOrder, model.CateCount, model.CateAlias, model.CateIntro, model.CateRootID, model.CateParentID, model.CateTemplate, model.CateLogTemplate, model.CateMeta)
    if err != nil {
        return 0, err
    } 
    
    return result.LastInsertId()
}

// AddZbpCategory 插入ZbpCategory
func AddZbpCategory(model *ZbpCategory) (int64, error) {
    return DefaultZbpCategoryConnection.Add(model)
}

// Find 查询ZbpCategory
func (c ZbpCategoryConnection) Find(condition string, args ...interface{}) ([]*ZbpCategory, error) {
    sqlStr := "SELECT `cate_ID`, `cate_Name`, `cate_Order`, `cate_Count`, `cate_Alias`, `cate_Intro`, `cate_RootID`, `cate_ParentID`, `cate_Template`, `cate_LogTemplate`, `cate_Meta` FROM `zbp_category`"
    if len(condition) > 0 {
        sqlStr = sqlStr + " WHERE " + condition
    }
    results := make([]*ZbpCategory, 0)

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
            model := ZbpCategory{}
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
            }
            rows.Scan(values...)
            if *(values[0].(*interface{})) != nil {
                tmp := int((*(values[0].(*interface{}))).(int64))
                model.CateID = tmp
            }
            if *(values[1].(*interface{})) != nil {
                tmp := string((*(values[1].(*interface{}))).([]uint8))
                model.CateName = tmp
            }
            if *(values[2].(*interface{})) != nil {
                tmp := int((*(values[2].(*interface{}))).(int64))
                model.CateOrder = tmp
            }
            if *(values[3].(*interface{})) != nil {
                tmp := int((*(values[3].(*interface{}))).(int64))
                model.CateCount = tmp
            }
            if *(values[4].(*interface{})) != nil {
                tmp := string((*(values[4].(*interface{}))).([]uint8))
                model.CateAlias = tmp
            }
            if *(values[5].(*interface{})) != nil {
                tmp := string((*(values[5].(*interface{}))).([]uint8))
                model.CateIntro = tmp
            }
            if *(values[6].(*interface{})) != nil {
                tmp := int((*(values[6].(*interface{}))).(int64))
                model.CateRootID = tmp
            }
            if *(values[7].(*interface{})) != nil {
                tmp := int((*(values[7].(*interface{}))).(int64))
                model.CateParentID = tmp
            }
            if *(values[8].(*interface{})) != nil {
                tmp := string((*(values[8].(*interface{}))).([]uint8))
                model.CateTemplate = tmp
            }
            if *(values[9].(*interface{})) != nil {
                tmp := string((*(values[9].(*interface{}))).([]uint8))
                model.CateLogTemplate = tmp
            }
            if *(values[10].(*interface{})) != nil {
                tmp := string((*(values[10].(*interface{}))).([]uint8))
                model.CateMeta = tmp
            }
            results = append(results, &model)
        }
    return results, nil
}

// FindZbpCategory 查询ZbpCategory
func FindZbpCategory(condition string, args ...interface{}) ([]*ZbpCategory, error) {
    return DefaultZbpCategoryConnection.Find(condition, args...)
}
// PagedQuery 分页查询ZbpCategory
func (c ZbpCategoryConnection) PagedQuery(condition string, pageSize uint, page uint, args ...interface{}) (totalCount uint, rows []*ZbpCategory, err error) {
	sqlStr := "SELECT COUNT(1) as cnt FROM `zbp_category`"
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
		return totalCount, []*ZbpCategory{}, nil
	}

	if len(condition) == 0 {
		condition = fmt.Sprintf("1=1")
	}
	condition = condition + fmt.Sprintf(" LIMIT %d OFFSET %d", pageSize, offset)
	rows, err = c.Find(condition, args...)
	return
}

// ZbpCategoryPagedQuery 分页查询ZbpCategory
func ZbpCategoryPagedQuery(condition string, pageSize uint, page uint, args ...interface{}) (totalCount uint, rows []*ZbpCategory, err error) {
	return DefaultZbpCategoryConnection.PagedQuery(condition, pageSize, page, args...)
}

// Get 获取ZbpCategory
func (c ZbpCategoryConnection) Get(condition string, args ...interface{}) (*ZbpCategory, error) {
    results, err := c.Find(condition, args...)

    if err != nil {
        return nil, err
    } 
    
    if len(results) > 0 {
        return results[0], nil
    } 
        
    return nil, nil
}


// GetZbpCategory 获取ZbpCategory
func GetZbpCategory(condition string, args ...interface{}) (*ZbpCategory, error) {
    return DefaultZbpCategoryConnection.Get(condition, args...)
}

// Update 更新ZbpCategory
func (c ZbpCategoryConnection) Update(model *ZbpCategory) (int64, error) {
    sqlStr := "UPDATE `zbp_category` SET `cate_Name` = ?, `cate_Order` = ?, `cate_Count` = ?, `cate_Alias` = ?, `cate_Intro` = ?, `cate_RootID` = ?, `cate_ParentID` = ?, `cate_Template` = ?, `cate_LogTemplate` = ?, `cate_Meta` = ? WHERE `cate_ID` = ?"
    result, err := c().Exec(sqlStr, model.CateName, model.CateOrder, model.CateCount, model.CateAlias, model.CateIntro, model.CateRootID, model.CateParentID, model.CateTemplate, model.CateLogTemplate, model.CateMeta, model.CateID)
    if err != nil {
        return 0, err
    }
    return result.RowsAffected()
}

// UpdateZbpCategory 更新ZbpCategory
func UpdateZbpCategory(model *ZbpCategory) (int64, error) {
    return DefaultZbpCategoryConnection.Update(model)
}
