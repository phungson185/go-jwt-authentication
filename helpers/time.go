package helpers

import (
	"jwt-authen/dtos"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GenerateTimeRequest(c *gin.Context) *dtos.Time {

	day := time.Now().Day()
	month := int(time.Now().Month())
	year := time.Now().Year()
	week := int(time.Now().Weekday())
	from := ""
	to := ""
	type_ := ""
	query := c.Request.URL.Query()

	for key, value := range query {
		queryValue := value[len(value)-1]

		switch key {
		case "day":
			day, _ = strconv.Atoi(queryValue)
			type_ = "particular"
			break
		case "month":
			month, _ = strconv.Atoi(queryValue)
			type_ = "particular"
			break
		case "year":
			year, _ = strconv.Atoi(queryValue)
			type_ = "particular"
			break
		case "week":
			week, _ = strconv.Atoi(queryValue)
			type_ = "week"
			break
		case "from":
			from = queryValue
			type_ = "range"
			break
		case "to":
			to = queryValue
			type_ = "range"
			break
		}
	}

	return &dtos.Time{Day: day, Month: month, Year: year, Week: week, From: from, To: to, Type_: type_}
}
