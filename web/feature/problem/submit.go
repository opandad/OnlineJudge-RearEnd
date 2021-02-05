package problem

/*
	@Title
	problem.Submit

	@Description
	显示题目功能

	@Func List
	| func name           | develop  | unit test |
	|--------------------------------------------|
	| Submit              |    no    |    no	 |
*/

import "OnlineJudge-RearEnd/web/model"

/*
	正在开发

	@Title
	problems.Submit

	@description
	提交题目

	@param
	提交的题目id，提交题目语言，提交题目的代码

	@return
	提交结果
*/
func Submit(websocketInputData *model.WebsocketInputData, websocketOutputData *model.WebsocketOutputData) error {
	//查询是否有这个题目

	//是否具备提交资质

	//提交进消息队列内

	//提交成功返回nil
	return nil
}
