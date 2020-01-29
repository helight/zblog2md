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

package config

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/cihub/seelog"
	"github.com/go-sql-driver/mysql"
)

var db map[string]*sql.DB = make(map[string]*sql.DB)
var dbMutex sync.RWMutex

var (
	dbUsername = map[string]string{}
	dbPassword = map[string]string{}
	dbAddress  = map[string]string{}
	dbName     = map[string]string{}
)

func init() {
	seelog.Info("init")
}

// initializeDatabase 初始化数据库连接信息
func initializeDatabase(mode string) {
	switch mode {
	case "dev", "pre":
		//自动化测试库
		dbAddress["zblog"] = "localhost:3306"
		dbUsername["zblog"] = "zblog"
		dbPassword["zblog"] = "123123"
		dbName["zblog"] = "zblog"

	case "release":
		//自动化测试库
		dbAddress["zblog"] = "localhost:3306"
		dbUsername["zblog"] = "zblog"
		dbPassword["zblog"] = "123123"
		dbName["zblog"] = "zblog"

	default:
		dbAddress["zblog"] = "localhost:3306"
		dbUsername["zblog"] = "zblog"
		dbPassword["zblog"] = "123123"
		dbName["zblog"] = "zblog"
		fmt.Println("DB default!")
	}
}

func GetDBConnect(dbname string) func() *sql.DB {
	return func() *sql.DB {
		dbMutex.RLock()
		conn := db[dbname]
		if conn != nil {
			dbMutex.RUnlock()
			return conn
		}
		dbMutex.RUnlock()

		dbMutex.Lock()
		defer dbMutex.Unlock()
		conn = db[dbname]
		if conn != nil {
			return conn
		}

		connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&loc=Local&charset=utf8&timeout=5s&readTimeout=30s&writeTimeout=30s", dbUsername[dbname], dbPassword[dbname], dbAddress[dbname], dbName[dbname])
		fmt.Println("connStr: ", connStr)
		var err error
		conn, err = sql.Open("mysql", connStr)
		if err != nil {
			log.Fatal("Connect to database failed: ", connStr)
			log.Panicln(err)
		}
		conn.SetMaxIdleConns(10)
		db[dbname] = conn
		return conn
	}
}

func GetDBConnectByCfg(strDBUserName, strDBPassword, strDBHost string, iDBPort int, strDBName string) *sql.DB {
	strDBKey := (strDBUserName + strDBPassword + strDBHost + strDBName)
	dbMutex.RLock()
	conn := db[strDBKey]
	if conn != nil {
		dbMutex.RUnlock()
		return conn
	}
	dbMutex.RUnlock()

	dbMutex.Lock()
	defer dbMutex.Unlock()
	conn = db[strDBKey]
	if conn != nil {
		return conn
	}

	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local&charset=utf8&timeout=5s&readTimeout=30s&writeTimeout=30s", strDBUserName, strDBPassword, strDBHost, iDBPort, strDBName)
	var err error
	conn, err = sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal("Connect to database failed: ", connStr)
		log.Panicln(err)
	}
	conn.SetMaxIdleConns(10)
	db[strDBKey] = conn
	return conn
}

func init() {
	mysql.RegisterDial("failover", failoverDial)
}

func failoverDial(addr string) (conn net.Conn, err error) {
	addrs := strings.Split(addr, ",")
	for _, v := range addrs {
		nd := net.Dialer{Timeout: 3 * time.Second}
		conn, err = nd.Dial("tcp", v)
		if err == nil {
			if strings.HasSuffix(v, ":6090") {
				return &secConn{conn}, err
			}
			return conn, err
		}
		seelog.Error("dial tcp ", v, " failed:", err)
	}
	return
}

type secConn struct {
	net.Conn
}

func (s *secConn) Read(b []byte) (n int, err error) {
	n, err = s.Conn.Read(b)
	if n > 0 {
		for i := 0; i < n; i++ {
			b[i] = b[i] ^ 0x66
		}
	}
	return
}

func (s *secConn) Write(b []byte) (n int, err error) {
	cb := make([]byte, len(b))
	if len(b) > 0 {
		for i := 0; i < len(b); i++ {
			cb[i] = b[i] ^ 0x66
		}
	}
	n, err = s.Conn.Write(cb)
	return
}
