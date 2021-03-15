package web

import (
	"OnlineJudge-RearEnd/api/verification"
	"OnlineJudge-RearEnd/configs"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

/*
	bug list
	没有掉线机制，容易被ddos
*/

func Init() {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	// r.Any("/", Websocket)
	router.POST("/snowflakeID", authSnowflakeID)
	router.POST("/authLogin", authLogin)

	//需要添加函数
	account := router.Group("/account")
	{
		login := account.Group("/login")
		{
			login.POST("/user", loginByUser)
			login.POST("/email", loginByEmail)
		}
		account.POST("/logout", logout)
		regist := account.Group("/regist")
		{
			regist.POST("/email", registByEmail)
		}
		verifyCode := account.Group("/verifyCode")
		{
			verifyCode.POST("/email", getEmailVerifyCode)
		}
	}

	//problem
	router.GET("/problem", getProblemList)
	router.GET("/problem/:id", getProblemDetail)

	//未完成
	router.GET("/contest", getContestList)
	router.POST("/contest/:id", getContestDetail)

	//未完成
	// router.GET("/userInfo/:id")
	// router.PUT("/userInfo/:id")

	//未完成
	router.POST("/submit/list", getSubmit)
	router.POST("/submit", submitAnswer)

	/*
		管理函数
		需要添加验证函数
	*/
	//未完成
	admin := router.Group("/admin", authAdmin)
	{
		fmt.Println("route")

		admin.POST("/problem/edit/:id", getProblemEdit)
		// admin.POST("/problem/edit/:id", editProblem)
		// admin.POST("/problem/add", addProblem)
		// admin.DELETE("/problem/delete/:id", deleteProblem)
		admin.POST("/user/list")
	}

	router.Run(configs.REAREND_SERVER_IP + ":" + configs.REAREND_SERVER_PORT)
}

func getSubmit(c *gin.Context) {
	type RD struct {
		Submit Submit `json:"submit"`
		Page   Page   `json:"page"`
	}

	var rd RD

	err := c.BindJSON(&rd)
	if err != nil {
		fmt.Println(err)
		return
	}

	var submit Submit
	submits, httpStatus, total := submit.List(rd.Page.PageIndex, rd.Page.PageSize)

	type Tmp struct {
		Submits    []Submit   `json:"submit"`
		HTTPStatus HTTPStatus `json:"httpStatus"`
		Page       Page       `json:"page"`
	}
	var tmp Tmp
	tmp.Submits = submits
	tmp.HTTPStatus = httpStatus
	tmp.Page.Total64 = total
	tmp.Page.PageIndex = rd.Page.PageIndex
	tmp.Page.PageSize = rd.Page.PageSize

	c.JSONP(http.StatusOK, tmp)
}

func submitAnswer(c *gin.Context) {
	type Tmp struct {
		LoginInfo LoginInfo `json:"loginInfo"`
		Submit    Submit    `json:"submit"`
	}
	var tmp Tmp

	var httpStatus HTTPStatus
	err := c.BindJSON(&tmp)
	if err != nil {
		fmt.Println(err)
		c.JSONP(http.StatusOK, HTTPStatus{
			Message: "服务器发生错误，请稍后尝试",
			IsError: true,
		})
	}
	// fmt.Println(tmp.Submit)

	httpStatus = tmp.LoginInfo.AuthLogin()
	if httpStatus.IsError == true {
		c.JSONP(http.StatusOK, httpStatus)
	}
	httpStatus = tmp.Submit.SubmitAnswer()
	c.JSONP(http.StatusOK, httpStatus)
}

func authSnowflakeID(c *gin.Context) {
	// snowflakeID, err := c.Cookie("snowflakeID")
	type SnowflakeID struct {
		SnowflakeID string `json:"snowflakeID"`
	}
	var snowflakeID SnowflakeID
	err := c.BindJSON(&snowflakeID)
	if err != nil {
		c.JSONP(http.StatusNotFound, nil)
	}

	// fmt.Println(snowflakeID)

	if snowflakeID.SnowflakeID == "" {
		snowflakeID.SnowflakeID = verification.Snowflake()
	}
	c.JSONP(http.StatusOK, snowflakeID)
}

func authLogin(c *gin.Context) {
	var loginInfo LoginInfo
	var httpStatus HTTPStatus
	err := c.BindJSON(&loginInfo)
	if err != nil {
		httpStatus.IsError = true
		httpStatus.Message = "服务器发生错误"
		c.JSONP(http.StatusUnauthorized, httpStatus)
	}

	// if loginInfo.UserID == 0 {
	// 	httpStatus.IsError = false
	// 	httpStatus.Message = ""
	// 	c.JSONP(http.StatusOK, httpStatus)
	// }

	httpStatus = loginInfo.AuthLogin()
	if err != nil {
		httpStatus.IsError = true
		httpStatus.Message = "服务器发生错误"
		c.JSONP(http.StatusUnauthorized, httpStatus)
	}

	// fmt.Println(loginInfo)
	// fmt.Println(httpStatus)

	c.JSONP(http.StatusOK, httpStatus)
}

func authAdmin(c *gin.Context) {
	fmt.Println("auth admin")
	fmt.Println(c.Request.Body)

	var loginInfo LoginInfo
	err := c.BindJSON(&loginInfo)
	if err != nil{
		// fmt.Println(err)
		c.JSONP(http.StatusNotFound, nil)
	}
	httpStatus := loginInfo.AuthAdmin()
	type Tmp struct {
		HTTPStatus HTTPStatus `json:"httpStatus"`
	}
	var tmp Tmp
	tmp.HTTPStatus = httpStatus

	if tmp.HTTPStatus.IsError {
		c.JSONP(http.StatusNotFound, tmp)
	}
}

func loginByUser(c *gin.Context) {

	// var user User

	// User.Login()
}

func loginByEmail(c *gin.Context) {
	// var frontEndData FrontEndData
	var loginInfo LoginInfo
	err := c.BindJSON(&loginInfo)
	if err != nil {
		fmt.Println(err)
	}
	var email Email
	email.Email = loginInfo.Account
	email.User.Password = loginInfo.Password

	userID, authority, userName, httpStatus := email.Login(loginInfo.SnowflakeID)

	loginInfo.UserID = userID
	loginInfo.Password = email.User.Password
	loginInfo.Authority = authority
	loginInfo.UserName = userName

	type TmpStruct struct {
		HTTPStatus HTTPStatus `json:"httpStatus"`
		LoginInfo  LoginInfo  `json:"loginInfo"`
	}

	var tmp TmpStruct
	tmp.LoginInfo = loginInfo
	tmp.HTTPStatus = httpStatus

	c.JSONP(http.StatusOK, tmp)
}

func logout(c *gin.Context) {
	var loginInfo LoginInfo
	err := c.BindJSON(&loginInfo)
	if err != nil {
		fmt.Println(err)
	}
	httpStatus := loginInfo.AuthLogin()

	if httpStatus.IsError == false {
		c.JSONP(http.StatusOK, httpStatus)
	} else {
		c.JSONP(http.StatusBadRequest, httpStatus)
	}
}

func registByEmail(c *gin.Context) {
	var loginInfo LoginInfo
	err := c.BindJSON(&loginInfo)
	if err != nil {
		fmt.Println(err)
	}
	var email Email
	email.Email = loginInfo.Account
	email.User.Password = loginInfo.Password

	type Tmp struct {
		User       User       `json:"user"`
		HTTPStatus HTTPStatus `json:"httpStatus"`
	}
	var tmp Tmp
	tmp.User, tmp.HTTPStatus = email.Regist(loginInfo.SnowflakeID, loginInfo.VerifyCode)
	if tmp.HTTPStatus.IsError == false {
		c.JSONP(http.StatusBadRequest, tmp)
	} else {
		c.JSONP(http.StatusOK, tmp)
	}
}

func getEmailVerifyCode(c *gin.Context) {
	var loginInfo LoginInfo
	err := c.BindJSON(&loginInfo)
	if err != nil {
		fmt.Println(err)
	}
	var email Email
	email.Email = loginInfo.Account
	httpStatus := email.SendVerifyCode()
	if httpStatus.IsError == false {
		c.JSONP(http.StatusOK, httpStatus)
	} else {
		c.JSONP(http.StatusBadRequest, httpStatus)
	}
}

func getProblemList(c *gin.Context) {
	var page Page
	var err error
	var httpStatus HTTPStatus
	page.PageIndex, err = strconv.Atoi(c.DefaultQuery("pageIndex", "1"))
	if err != nil {
		httpStatus.Message = "服务器内部转int错误"
		httpStatus.IsError = true
	}
	page.PageSize, err = strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	if err != nil {
		httpStatus.Message = "服务器内部转int错误"
		httpStatus.IsError = true
	}

	var problem []Problem
	var tempProblem Problem
	var total int64
	problem = make([]Problem, page.PageSize)
	problem, httpStatus, total = tempProblem.List(page.PageIndex, page.PageSize)

	type Tmp struct {
		HTTPStatus HTTPStatus `json:"httpStatus"`
		Problem    []Problem  `json:"problem"`
		Total      int64      `json:"total"`
	}
	var tmp Tmp
	tmp.Problem = problem
	tmp.HTTPStatus = httpStatus
	tmp.Total = total

	fmt.Println(total)

	c.JSONP(http.StatusOK, tmp)
}

func getProblemDetail(c *gin.Context) {
	id := c.Param("id")
	var problem Problem
	var err error
	var httpStatus HTTPStatus
	problem.ID, err = strconv.Atoi(id)
	type Tmp struct {
		Problem    Problem    `json:"problem"`
		HTTPStatus HTTPStatus `json:"httpStatus"`
		Languages  []Language `json:"languages"`
	}
	if err != nil {
		httpStatus.IsError = true
		httpStatus.Message = "服务器出现错误"
		httpStatus.SubMessage = "string转int出现错误，server.getProblemDetail"
		c.JSONP(http.StatusOK, &Tmp{
			HTTPStatus: httpStatus,
		})
	}
	problem, httpStatus = problem.Detail()
	var language Language
	var languages []Language
	languages, httpStatus = language.List()
	httpStatus.IsError = false
	httpStatus.Message = ""

	c.JSONP(http.StatusOK, &Tmp{
		Problem:    problem,
		HTTPStatus: httpStatus,
		Languages:  languages,
	})
}

func getProblemEdit(c *gin.Context) {
	id := c.Param("id")
	var problem Problem
	var err error
	var httpStatus HTTPStatus
	problem.ID, err = strconv.Atoi(id)
	type Tmp struct {
		Problem    Problem    `json:"problem"`
		HTTPStatus HTTPStatus `json:"httpStatus"`
		Languages  []Language `json:"languages"`
	}
	if err != nil {
		httpStatus.IsError = true
		httpStatus.Message = "服务器出现错误"
		httpStatus.SubMessage = "string转int出现错误，server.getProblemDetail"
		c.JSONP(http.StatusOK, &Tmp{
			HTTPStatus: httpStatus,
		})
	}
	problem, httpStatus = problem.Detail()
	httpStatus.IsError = false
	httpStatus.Message = ""

	c.JSONP(http.StatusOK, &Tmp{
		Problem:    problem,
		HTTPStatus: httpStatus,
	})
}

func getContestList(c *gin.Context) {
	var page Page
	var err error
	type Tmp struct {
		Page       Page       `json:"page"`
		Contest    []Contest  `json:"contest"`
		HTTPStatus HTTPStatus `json:"httpStatus"`
	}
	var tmp Tmp

	page.PageIndex, err = strconv.Atoi(c.DefaultQuery("pageIndex", "1"))
	if err != nil {
		tmp.HTTPStatus.Message = "服务器内部转int错误"
		tmp.HTTPStatus.IsError = true
		tmp.HTTPStatus.RequestPath = "get contest list"
		c.JSONP(http.StatusOK, tmp)
	}
	page.PageSize, err = strconv.Atoi(c.DefaultQuery("pageSize", "5"))
	if err != nil {
		tmp.HTTPStatus.Message = "服务器内部转int错误"
		tmp.HTTPStatus.IsError = true
		tmp.HTTPStatus.RequestPath = "get contest list"
		c.JSONP(http.StatusOK, tmp)
	}

	var tempContest Contest
	var total int64
	tmp.Contest, tmp.HTTPStatus, tmp.Page.Total = tempContest.List(page.PageIndex, page.PageSize)

	fmt.Println(total)

	c.JSONP(http.StatusOK, tmp)
}

func getContestDetail(c *gin.Context) {
	id := c.Param("id")

	type Tmp struct {
		LoginInfo  LoginInfo  `json:"loginInfo"`
		HTTPStatus HTTPStatus `json:"httpStatus"`
		Contest    Contest    `json:"contest"`
		Problem    []Problem  `json:"problems"`
		Language   []Language `json:"languages"`
	}
	var tmp Tmp
	var err error
	c.BindJSON(&tmp)
	tmp.Contest.ID, err = strconv.Atoi(id)
	if err != nil {
		tmp.HTTPStatus.IsError = true
		tmp.HTTPStatus.Message = "服务器发生错误，请稍后尝试"
		tmp.HTTPStatus.SubMessage = "string to int error"
		tmp.HTTPStatus.RequestPath = "get contest detail"
		c.JSONP(http.StatusOK, tmp)
	}

	tmp.HTTPStatus = tmp.LoginInfo.AuthLogin()
	if tmp.HTTPStatus.IsError == true {
		c.JSONP(http.StatusOK, tmp)
	}

	//查询有无竞赛资格
	var contest Contest
	contest, tmp.Problem, tmp.Language, tmp.HTTPStatus = tmp.Contest.Detail(tmp.LoginInfo.UserID)
	tmp.Contest = contest

	tmp.LoginInfo.UserID = 0
	tmp.LoginInfo.Password = ""
	tmp.LoginInfo.SnowflakeID = ""
	tmp.LoginInfo.Authority = ""

	c.JSONP(http.StatusOK, tmp)
}

// /*
// 	获取前端接收数据，并返回数据
// */
// func Websocket(c *gin.Context) {
// 	var upGrader = websocket.Upgrader{
// 		CheckOrigin: func(r *http.Request) bool {
// 			return true
// 		},
// 	}

// 	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
// 	if err != nil {
// 		return
// 	}
// 	defer ws.Close()

// 	for {
// 		var receiveData FrontEndData

// 		err = ws.ReadJSON(&receiveData)
// 		if err != nil {
// 			fmt.Println("error read json data")
// 			fmt.Println(err)
// 			break
// 		} else {
// 			sendData := Router(receiveData)

// 			fmt.Println(sendData)

// 			err = ws.WriteJSON(sendData)
// 			if err != nil {
// 				fmt.Println("error send json data")
// 				fmt.Println(err)
// 				break
// 			}
// 		}
// 	}
// }

// /*
// 	msg格式：login/xxx/xxx

// 	route list

// 	--user
// 		--login
// 			--email
// 			--auto
// 		--regist
// 			--email
// 		--userInfo
// 		--verifyCode
// 			--email

// 	--problems
// 		--list
// 		--detail
// 		--submit
// */
// func Router(receiveData FrontEndData) FrontEndData {
// 	//检测是否为404，解析请求路径
// 	var sendData FrontEndData
// 	var isNot404 bool = false
// 	var requestPath []string = strings.Split(receiveData.HTTPStatus.RequestPath, "/")

// 	//验证登录
// 	// sendData.HTTPStatus = receiveData.Data.LoginInfo.AuthLogin()
// 	// if sendData.HTTPStatus.IsError == true {
// 	// 	return sendData
// 	// }

// 	//test
// 	fmt.Println("Router output test\n", receiveData, "\n", requestPath)

// 	//account
// 	if requestPath[0] == "account" {
// 		if requestPath[1] == "login" {
// 			if requestPath[2] == "email" {
// 				isNot404 = true
// 				sendData.Data.Email = make([]Email, 1)
// 				sendData.Data.Email[0].User.ID, sendData.HTTPStatus = receiveData.Data.Email[0].Login(receiveData.WebsocketID)
// 			}
// 			if requestPath[2] == "user" {
// 				isNot404 = true
// 				_, sendData.HTTPStatus = receiveData.Data.User[0].Login(receiveData.WebsocketID)
// 			}
// 		}
// 		if requestPath[1] == "regist" {
// 			if requestPath[2] == "email" {
// 				isNot404 = true
// 				sendData.Data.Email = make([]Email, 1)
// 				sendData.Data.Email[0].User, sendData.HTTPStatus = receiveData.Data.Email[0].Regist(receiveData.WebsocketID, receiveData.Data.LoginInfo.VerifyCode)
// 			}
// 		}
// 		if requestPath[1] == "userInfo" {
// 		}
// 		if requestPath[1] == "verifyCode" {
// 			if requestPath[2] == "email" {
// 				isNot404 = true
// 				sendData.HTTPStatus = receiveData.Data.Email[0].SendVerifyCode()
// 			}
// 		}
// 	}

// 	//problems
// 	if requestPath[0] == "problem" {
// 		if requestPath[1] == "list" {
// 			isNot404 = true
// 			sendData.Data.Problem = make([]Problem, receiveData.Data.Page.PageSize)
// 			sendData.Data.Problem, sendData.HTTPStatus = sendData.Data.Problem[0].List(receiveData.Data.Page.PageIndex, receiveData.Data.Page.PageSize)
// 		}
// 		if requestPath[1] == "detail" {
// 			//需要判断题目是否存在，如果不存在返回404
// 			isNot404 = true
// 			sendData.Data.Problem = make([]Problem, 1)
// 			sendData.Data.Problem[0], sendData.HTTPStatus = receiveData.Data.Problem[0].Detail()
// 		}
// 	}

// 	//submit
// 	if requestPath[0] == "submit" && receiveData.Data.LoginInfo.UserID != 0 && receiveData.Data.LoginInfo.Password != "" && receiveData.Data.LoginInfo.WebsocketID != "" {
// 		receiveData.Data.Submit[0].SubmitAnswer()
// 	}

// 	//404 not found
// 	if isNot404 == false {
// 		sendData.HTTPStatus = HTTPStatus{
// 			Message:     "页面走丢了ToT",
// 			IsError:     true,
// 			ErrorCode:   404,
// 			SubMessage:  "404",
// 			RequestPath: receiveData.HTTPStatus.RequestPath,
// 		}
// 	}

// 	return sendData
// }
