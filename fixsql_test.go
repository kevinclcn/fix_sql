package main

import (
	"fmt"
	"testing"

	"github.com/xwb1989/sqlparser"
)

func TestFixInsert(t *testing.T) {
	str := "INSERT INTO `customer_order_detail` (`tenant_code`, `customer_order_id`, `product_id`, `order_create_time`, `customer_id`, `order_no`, `id`, `external_product_id`, `sku_no`, `product_no`, `sku_bar_code`, `product_name`, `pic_path`, `brand_name`, `category`, `qty`, `price_unit`, `price_sub_total`, `price_sub_paid`, `create_time`, `update_time`, `post_fee`) VALUES ('380218', 1741605172130787331, 1676091015108366339, '2024-01-01 07:36:45', 1675844337542402311, '6925366924928423209', 1741605187775541252, NULL, '1764943235817487', '3545005398065072613', NULL, '【含赠品到手13件】养元青防脱育发洗发乳育发液套装侧柏叶控油强韧', NULL, '', '个人护理-洗发护发-洗护套装', 1, 149.00, 149.00, 139.00, '2024-01-01 07:40:08', '', NULL);"

	stmt, err := sqlparser.Parse(str)
	if err != nil {
		panic(err)
	}

	columns := stmt.(*sqlparser.Insert).Columns

	to := sqlparser.NewValArg([]byte("'2024-01-01 07:040:08'"))

	froms := []sqlparser.Expr{}
	for _, values := range stmt.(*sqlparser.Insert).Rows.(sqlparser.Values) {
		for i, val := range values {
			// fmt.Println(columns[i].String(), sqlparser.String(val))

			switch v := val.(type) {
			case *sqlparser.SQLVal:

				if columns[i].String() == "update_time" && v.Type == sqlparser.StrVal && string(v.Val) == "" {
					froms = append(froms, val)
				}
			}
		}
	}
	var expr sqlparser.Expr = stmt.(*sqlparser.Insert).Rows.(sqlparser.Values)[0]
	for _, from := range froms {
		expr = sqlparser.ReplaceExpr(expr, from, to)
	}
	fmt.Println(sqlparser.String(expr))
}
