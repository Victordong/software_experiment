package plugin

import (
	"context"
	"github.com/jinzhu/gorm"
	"strconv"
	"strings"
)

type PageMap struct {
	PerPage int64
	Page    int64
}

type OrderMap struct {
	Sort  string
	Order string
}

// 获取修饰的表
// 传入key 如果有修饰表返回修饰表
func GetPrefixForm(key string) string {
	res := strings.Split(key, ".")
	if len(res) > 1 {
		return res[0]
	} else {
		return ""
	}
}

// 根据CompareMap 获取 查询的 运算符 和 字段名
func GetCompareMap(key string) (string, string) {
	compareChoiceMap := make(map[string]string)
	initCompareMap(&compareChoiceMap)
	op := "in"
	field := key
	for compareKey, compareOp := range compareChoiceMap {
		if strings.HasSuffix(key, compareKey) {
			op = compareOp
			field = key[:len(key)-len(compareKey)]
		}
	}
	return op, field
}

// 过滤查询
func FilterSQL(db *gorm.DB, filterMap map[string][]string, baseForm string) *gorm.DB {
	for key, value := range filterMap {
		form := GetPrefixForm(key)
		if form == "" {
			form = baseForm
		}

		if form != baseForm {
			// 修饰表不等于基础表时join修饰表
			db = db.Joins(strings.Join([]string{"join ", form, " on ", form, ".", "id", " = ", baseForm, ".", form[:len(form)-1], "_id"}, ""))
		} else {
			// 等于时附加表名
			key = baseForm + "." + key
		}

		// 获取filter的方式和字段
		op, field := GetCompareMap(key)
		if op == "like" {
			db = db.Where(strings.Join([]string{field, op, "?"}, " "), (value[0])+"%")
		} else if op == "in" {
			db = db.Where(strings.Join([]string{field, "in (?)"}, " "), value)
		} else {
			db = db.Where(strings.Join([]string{field, op, "?"}, " "), value[0])
		}
	}
	return db
}

// 排序查询
func SortSQL(db *gorm.DB, orderMap *OrderMap, baseForm string) *gorm.DB {
	db = db.Order(strings.Join([]string{orderMap.Sort, orderMap.Order}, " "))
	return db
}

// 分页查询
func PageSql(db *gorm.DB, pageMap *PageMap, baseForm string) *gorm.DB {
	db = db.Offset((pageMap.Page - 1) * pageMap.PerPage).Limit(pageMap.PerPage)
	return db
}

// 传入queryMap 和 基础表 返回 分页 排序 和 查询
func UrlCondition(queryMap map[string][]string, baseForm string) (*PageMap, *OrderMap, map[string][]string) {
	pageMap := PageMap{
		PerPage: 20,
		Page:    1,
	}

	orderMap := OrderMap{
		Sort:  "id",
		Order: "desc",
	}

	filterMap := make(map[string][]string)

	for key, value := range queryMap {
		key = strings.TrimSpace(key)
		if key == "_per_page" {
			perPage, _ := strconv.ParseInt(value[0], 10, 64)
			pageMap.PerPage = perPage
		} else if key == "_page" {
			page, _ := strconv.ParseInt(value[0], 10, 64)
			pageMap.Page = page
		} else if key == "_sort" {
			orderMap.Sort = value[0]
		} else if key == "_order" {
			orderMap.Order = value[0]
		} else {
			filterMap[key] = value
		}
	}

	return &pageMap, &orderMap, filterMap
}

// 初始化filter条件映射
func initCompareMap(compareMap *map[string]string) {
	(*compareMap)["_lt"] = "<"
	(*compareMap)["_gt"] = ">"
	(*compareMap)["_lte"] = "<="
	(*compareMap)["_gte"] = ">="
	(*compareMap)["_ne"] = "!="
	(*compareMap)["_like"] = "like"
}

// 进行全查询
func ProcessQuery(db *gorm.DB, queryMap map[string][]string, baseForm string) (*gorm.DB, int64, error) {
	pager, order, filter := UrlCondition(queryMap, baseForm)
	var err error
	db = FilterSQL(db, filter, baseForm)
	db = SortSQL(db, order, baseForm)
	var num int64
	db.Where(strings.Join([]string{baseForm, ".deleted_at is NULL"}, "")).Count(&num)
	db = PageSql(db, pager, baseForm)
	err = db.Error
	return db, num, err
}

// 将上下文的查询内容放入url参数解析的结果
func CtxQueryMap(ctx context.Context, queryMap map[string][]string) {
	filterMapValue := ctx.Value("filterMap")
	var filterMap map[string]interface{}
	switch filterMapValue.(type) {
	case map[string]interface{}:
		filterMap = filterMapValue.(map[string]interface{})
	default:
		filterMap = make(map[string]interface{})
	}
	for key, value := range filterMap {
		switch value.(type) {
		case string:
			queryMap[key] = []string{value.(string)}
		case int:
			queryMap[key] = []string{strconv.Itoa(value.(int))}
		case uint:
			queryMap[key] = []string{strconv.Itoa(int(value.(uint)))}
		}
	}
}
