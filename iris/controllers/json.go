package controllers

import "github.com/kataras/iris/v12"

func ReturnJSON(c iris.Context) {
	data := new(interface{})

	if err := c.ReadJSON(data); err != nil {
		c.StatusCode(iris.StatusBadRequest)
		c.WriteString(err.Error())
		return
	}

	c.JSON(data)
}
