package buildsql

import (
	"ant/reflectutil"
	"ant/stringutil"
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

type Closure = func(this *BuildSql) //用于构建where条件的闭包函数

//构建sql对象
//默认情况下零值不会参与sql的构建，若需要则使用Emphasize函数对其进行强调
type BuildSql struct {
	table        string      //表名
	tableAlias   string      //主表别名
	whereStr     string      //where条件字串
	closureRun   bool        //标识闭包函数正在被调用
	limit        string      //数据条数限制
	joinStr      string      //连表字串
	fieldStr     string      //查询字段
	groupBy      string      //分组字串
	orderBy      string      //排序字串
	primaryModel interface{} //主表模型

	//以下属性用于解析响应数据为对应的结构体
	modelFieldMap   map[string][2]string   //结构体的字段与模型名称和数据库字段对应的映射关系,用于http查询反解析,[2]string中0=model的名称,1=model的字段
	fieldMap        map[string]interface{} //数据库字段和值的映射关系
	fieldMapSortKey []string               //存放排过序的fieldMap的key
	emphasizeKey    []string               //强调字段,如果model的字段在这里面,那么构建sql的时候零值将不会被丢弃
	selectField     []string               //Select 查询，指定需要查询的字段，默认全部查询

	errors []error //构建sql中产生的错误
	Debug  bool    //是否开启调试，开启后会打印出构建的sql语句
}

//分页对象
type Pagination struct {
	Total    int    //总页数k
	Page     int    //当前页 url中为page参数
	PageSize int    //分页数量 url 中为per 参数
	Link     string //生成的a标签
}

//构造函数
func NewModel(model interface{}, params ...string) *BuildSql {
	buildSql := &BuildSql{}

	if len(params) == 1 {
		buildSql.tableAlias = params[0]
	}else{
		buildSql.tableAlias,_=GetTableNameFromModel(model)
	}

	buildSql.fieldMap = make(map[string]interface{})
	buildSql.modelFieldMap = make(map[string][2]string)

	buildSql.primaryModel = model

	return buildSql
}


//基本插入
func (this *BuildSql)baseInsert(insert string) (string, error) {
	var err error

	//检查错误
	if err = this.checkError(); err != nil {
		return "", err
	}

	//获取主表名称
	this.table, err =GetTableNameFromModel(this.primaryModel)
	if err != nil {
		return "", err
	}

	//解析字段
	if err := this.setFieldMap(this.primaryModel, true, ""); err != nil {
		return "", err
	}

	//字段解析
	if len(this.fieldMap)== 0 {
		return "",errors.New("字段缺失")
	}

	column := ""
	value := ""
	for k, v := range this.fieldMap {
		column += k + ","
		if str, ok := v.(string); ok {
			value += fmt.Sprintf("'%s',", str)
		} else {
			value += fmt.Sprintf("%v,", v)
		}
	}

	column = strings.Trim(column, ",")
	value = strings.Trim(value, ",")


	return fmt.Sprintf("%s INTO %s (%s) VALUES (%s);",insert, this.table, column, value), nil

}

//插入数据
func (this *BuildSql) Insert() (string, error) {
	return this.baseInsert("INSERT")

}

//更新数据
func (this *BuildSql) Update() (string, error) {
	var err error

	//检查错误
	if err = this.checkError(); err != nil {
		return "", err
	}

	//获取主表名称
	this.table, err = GetTableNameFromModel(this.primaryModel)
	if err != nil {
		return "", err
	}

	if err := this.setFieldMap(this.primaryModel, true, ""); err != nil {
		return "", err
	}
	//字段解析
	if len(this.fieldMap)== 0 {
		return "",errors.New("字段缺失")
	}


	setStr := ""
	for k, v := range this.fieldMap {
		tmp := k + "="
		if str, ok := v.(string); ok {
			tmp += fmt.Sprintf("'%s',", str)
		} else {
			tmp += fmt.Sprintf("%d,", v)
		}

		setStr += tmp
	}

	setStr = strings.Trim(setStr, ",")

	if this.whereStr== "" {
		return "",errors.New("where 条件缺失")
	}

	return fmt.Sprintf("UPDATE %s SET %s%s;", this.table, setStr, this.whereStr), nil

}

//删除数据
func (this *BuildSql) Delete() (string, error) {
	var err error

	//检查错误
	if err = this.checkError(); err != nil {
		return "", err
	}

	//获取主表名称
	if str, ok := this.primaryModel.(string); ok {
		this.table = str
	} else {
		this.table, err = GetTableNameFromModel(this.primaryModel)
		if err != nil {
			return "", err
		}
	}

	if this.whereStr == "" {
		return "", errors.New("where 条件缺失")
	}
	return fmt.Sprintf("DELETE FROM %s %s;", this.table, this.whereStr), nil
}

//where
func (this *BuildSql) Where(where ...interface{}) *BuildSql {
	return this.baseWhere("AND", where)
}

//or where
func (this *BuildSql) OrWhere(where ...interface{}) *BuildSql {
	return this.baseWhere("OR", where)
}

//基本where,使用闭包来打()号
func (this *BuildSql) baseWhere(flag string, where []interface{}) *BuildSql {

	if len(where) == 1 {
		//一个参数,闭包
		switch where0 := where[0].(type) {
		case string:
			this.appendWhereStr(flag, where0)
		case Closure:
			this.whereStr += fmt.Sprintf(" %s (", flag)
			this.closureRun = true
			where0(this)

			this.whereStr += ") "
		}

	} else if len(where) == 2 {

		//第一个参数必须是字符串
		if _, ok := where[0].(string); !ok {
			return this
		}

		//两个参数,等于
		switch where1 := where[1].(type) {
		case string:
			this.appendWhereStr(flag, fmt.Sprintf("%s='%s'", where[0], where1))
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
			this.appendWhereStr(flag, fmt.Sprintf("%s=%v", where[0], where1))

		}

	} else if len(where) == 3 {
		//三个参数,操作符
		//第一个参数必须是字符串
		if _, ok := where[0].(string); !ok {
			return this
		}

		//第二个参数必须是字符串
		if _, ok := where[1].(string); !ok {
			return this
		}

		switch where2 := where[2].(type) {
		case string:
			this.appendWhereStr(flag, fmt.Sprintf("%s %s '%s'", where[0], where[1], where2))
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
			this.appendWhereStr(flag, fmt.Sprintf("%s %s %v", where[0], where[1], where2))
		case []int:
			tmpStr := ""
			for _, v := range where2 {
				tmpStr += fmt.Sprintf("%d,", v)
			}
			tmpStr = strings.Trim(tmpStr, ",")
			this.appendWhereStr(flag, fmt.Sprintf("%s %s (%s)", where[0], where[1], tmpStr))
		case []uint32:
			tmpStr := ""
			for _, v := range where2 {
				tmpStr += fmt.Sprintf("%d,", v)
			}
			tmpStr = strings.Trim(tmpStr, ",")
			this.appendWhereStr(flag, fmt.Sprintf("%s %s (%s)", where[0], where[1], tmpStr))
		case []string:
			tmpStr := ""
			for _, v := range where2 {
				tmpStr += fmt.Sprintf("'%s',", v)
			}
			tmpStr = strings.Trim(tmpStr, ",")
			this.appendWhereStr(flag, fmt.Sprintf("%s %s (%s)", where[0], where[1], tmpStr))
		}
	}

	return this
}

func (this *BuildSql) appendWhereStr(flag string, str string) {
	if this.whereStr == "" {
		this.whereStr += " WHERE " + str
	} else {

		if this.closureRun {
			this.whereStr += str
			this.closureRun = false
		} else {
			this.whereStr += fmt.Sprintf(" %s ", flag) + str
		}

	}
}

/**
 * 强调字段
 * 默认情况：结构体整型为0，字符串为空的字段不会构建在SQL中
 * 如果希望构建在SQL中，则需要使用强调字段。
 * Emphasize("AAA", "bbb")
 * 其中：AAA,bbb为Golang当中结构体定义的字段值，而非数据库中定义的字段值
 */
func (this *BuildSql) Emphasize(params ...string) *BuildSql {
	this.emphasizeKey = params
	return this
}

//设置查询字段
func (this *BuildSql) Field(fields ...string) *BuildSql {

	this.selectField = fields

	return this
}

//分组
func (this *BuildSql) GroupBy(fields string) *BuildSql {
	this.groupBy = " GROUP BY " + fields

	return this
}

//排序
func (this *BuildSql) OrderBy(fields string) *BuildSql {
	this.orderBy = " Order BY " + fields

	return this
}

//从model 中获取表名称
func GetTableNameFromModel(model interface{}) (string, error) {

	tableName:=""

	objT,_,err := reflectutil.GetStructTV(model)

	if err != nil {
		return "",err
	}

	for i := 0; i < objT.NumField(); i++ {
		tableName =reflectutil.GetStructTagFuncName(objT.Field(i).Tag,"orm","table")
		if tableName!= "" {
			break
		}
	}

	return tableName, nil
}

//设置选中的字段,并生成查询字段
func (this *BuildSql) setFieldStr() *BuildSql {

	//填充过了,不再填充
	if this.fieldStr != "" {
		return this
	}

	//默认全部选中
	if len(this.selectField) != 0 {
		for k, _ := range this.fieldMap {
			if !stringutil.InSliceString(k, this.selectField) {
				//删除没有使用的字段
				delete(this.fieldMap, k)
			}
		}
	}

	for k, _ := range this.fieldMap {
		this.fieldMapSortKey = append(this.fieldMapSortKey, k)
	}
	sort.Strings(this.fieldMapSortKey)

	for _, key := range this.fieldMapSortKey {
		//剥离别名
		keyArr := strings.Split(key, ".")

		if len(keyArr) == 2 {
			//聚合字段
			if this.isPolymerization(keyArr[1]) {

				this.fieldStr += fmt.Sprintf("%s AS `%s`,", keyArr[1], key)
			} else {
				this.fieldStr += fmt.Sprintf("%v,", key)
			}

		} else {
			this.fieldStr += fmt.Sprintf("%v,", key)
		}

	}

	this.fieldStr = strings.Trim(this.fieldStr, ",")

	return this
}

//leftJoin,table 需要包含表名和别名
func (this *BuildSql) LeftJoin(model interface{}, alias string, on string) *BuildSql {

	this.join("LEFT", alias, on, model)

	return this
}

//Join
func (this *BuildSql) join(option string, alias string, on string, model interface{}) *BuildSql {

	table, err := GetTableNameFromModel(model)
	if err != nil {
		this.appendError(err)
		return this
	}

	//连接字符串
	this.joinStr += fmt.Sprintf(" %s JOIN %s %s ON %s ", option, table, alias, on)
	if err := this.setFieldMap(model, false, alias); err != nil {
		this.errors = append(this.errors, err)
	}

	return this
}

//查询所有
func (this *BuildSql) Select() (string, error) {
	var err error

	//检查错误
	if err = this.checkError(); err != nil {
		return "", err
	}

	//获取主表名称
	this.table, err = GetTableNameFromModel(this.primaryModel)
	if err != nil {
		return "", err
	}

	//没有进行连表操作
	if this.joinStr== "" {
		this.tableAlias=""
	}

	//解析模型字段
	err = this.setFieldMap(this.primaryModel, false, this.tableAlias)
	if err != nil {
		return "", err
	}

	//生成查询字段
	this.setFieldStr()

	if this.fieldStr== "" {
		return "",errors.New("未解析出查询字段")
	}

	sql := fmt.Sprintf("SELECT %s FROM %s %s%s%s%s%s%s;",
		this.fieldStr,
		this.table,
		this.tableAlias,
		this.joinStr,
		this.whereStr,
		this.groupBy,
		this.orderBy,
		this.limit)

	return sql, nil

}

//添加错误
func (this *BuildSql) appendError(err error) {
	if err != nil {
		this.errors = append(this.errors, err)
	}
}

//检查错误
func (this *BuildSql) checkError() error {
	if len(this.errors) > 1 {
		return this.errors[0]
	}

	return nil
}

//limit
func (this *BuildSql) Limit(params ...int) *BuildSql {

	if this.limit != "" {
		return this
	}

	if len(params) == 1 {
		this.limit = fmt.Sprintf(" LIMIT %d ", params[0])
	} else if len(params) == 2 {
		this.limit = fmt.Sprintf(" LIMIT %d OFFSET %d", params[1], params[0])

	}

	return this
}

//解析结构体的值成map
func (this *BuildSql) setFieldMap(model interface{}, dropEmpty bool, alias string) error {

	objT,objV,err := reflectutil.GetStructTV(model)
	if err != nil {
		return err
	}


	//遍历解析出NewFieldsMap
	for i := 0; i < objT.NumField(); i++ {
		value := objV.Field(i).Interface()

		if !stringutil.InSliceString(objT.Field(i).Name, this.emphasizeKey) && dropEmpty && reflectutil.IsEmpty(value) {
			//没有在强调map里,dropEmpty=true,值为零,会被丢弃
			continue
		}


		ormTag := objT.Field(i).Tag.Get("orm")
		field := GetColumnName(ormTag,"column")
		if field == "" {
			continue
		}

		if alias != "" {
			field = alias + "." + field
		}

		//聚合字段
		if this.isPolymerization(field) && !stringutil.InSliceString(field, this.emphasizeKey) {
			//是聚合字段,但是不是强调字段,丢弃
			continue
		}

		this.fieldMap[field] = value
		this.modelFieldMap[field] = [2]string{objT.Name(), objT.Field(i).Name}

	}

	return nil
}

//获取字段的名称
func GetColumnName(str string,funcName string) string {

	re := regexp.MustCompile(fmt.Sprintf(`%s\(([^(]*)\)`,funcName))

	res := re.FindStringSubmatch(str)

	if len(res) > 0 {
		return res[1]
	}

	return ""
}


//判定是否为聚合字段
func (this *BuildSql) isPolymerization(field string) bool {

	re := regexp.MustCompile(`(COUNT|SUM|MAX|MIN|AVG)\((.*)\)$`)

	return re.Match([]byte(field))

}

