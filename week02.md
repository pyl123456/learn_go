1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
- 不能，业务层不需要知道具体的sql错误，应该由数据层判断sql错误类型包装后返回其他错误信息。代码如下：
```
func insertUser(username, nickname, email, mobile, source string) {
	userid := int(time.Now().UnixNano()/1000000)
	d := rdb.GetDataBase()
	var name string
	err := d.QueryRow("select name from user where user_dir_id=1 and name=$1",username).Scan(&name)
	if err == sql.ErrNoRows {
		//新增数据
		stmt,err := d.Prepare("insert into user(user_id,name,nickname,email,phone_num,source) values($1,$2,$3,$4,$5,$6)")
		if err != nil {
			glog.Info("insert user prepare error ",err)
			return
		}
		_,err = stmt.Exec(userid,username,nickname,email,mobile)
		if err != nil {
			glog.Info("insert user error ",err)
			return
		}
	}else if err != nil {
		glog.Info("select user err ",err)
	}else{//不处理
		glog.Info("user exsits ")
	}
	 
}

```
