package controller

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/astaxie/beego/logs"
	"github.com/labstack/echo"
)

func ListProduct(c echo.Context) (err error) {

	retJ := map[string]interface{}{}

	ss := strings.TrimSpace(searchProductListModel)
	err = json.Unmarshal([]byte(ss), &retJ)
	if err != nil {
		logs.Error("umarshal err:%v", err)
	}

	outM := retJ["ProductListModel"].([]interface{})
	outJ := []map[string]interface{}{}
	for _, v := range outM {
		v1 := v.(map[string]interface{})

		v2 := map[string]interface{}{
			"Name":  v1["Name"],
			"Price": v1["Price"],
		}
		outJ = append(outJ, v2)
	}

	err = c.JSON(http.StatusOK, outJ)
	return
}
