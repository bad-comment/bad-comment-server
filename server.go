package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	router(e)
	e.Logger.Fatal(e.Start(":1323"))
}

/**



user
subject userId
1      1
2      1
3      2

userId subjectId yesOrNo
3      1
4      1
5      1
2      1
1      3

rId tId score deep
3   1         2
2   1         2
1   3         1


*/
