// Code generated by sdkcodegen; DO NOT EDIT.

package workwx

// execGetAccessToken 获取access_token
func (c *WorkwxApp) execGetAccessToken(req reqAccessToken) (respAccessToken, error) {
	var resp respAccessToken
	err := c.executeQyapiGet("/cgi-bin/gettoken", req, &resp, false)
	if err != nil {
		return respAccessToken{}, err
	}

	return resp, nil
}

// execUserGet 读取成员
func (c *WorkwxApp) execUserGet(req reqUserGet) (respUserGet, error) {
	var resp respUserGet
	err := c.executeQyapiGet("/cgi-bin/user/get", req, &resp, true)
	if err != nil {
		return respUserGet{}, err
	}

	return resp, nil
}

// execDeptList 获取部门列表
func (c *WorkwxApp) execDeptList(req reqDeptList) (respDeptList, error) {
	var resp respDeptList
	err := c.executeQyapiGet("/cgi-bin/department/list", req, &resp, true)
	if err != nil {
		return respDeptList{}, err
	}

	return resp, nil
}

// execAppchatCreate 创建群聊会话
func (c *WorkwxApp) execAppchatCreate(req reqAppchatCreate) (respAppchatCreate, error) {
	var resp respAppchatCreate
	err := c.executeQyapiJSONPost("/cgi-bin/appchat/create", req, &resp, true)
	if err != nil {
		return respAppchatCreate{}, err
	}

	return resp, nil
}

// execAppchatGet 获取群聊会话
func (c *WorkwxApp) execAppchatGet(req reqAppchatGet) (respAppchatGet, error) {
	var resp respAppchatGet
	err := c.executeQyapiGet("/cgi-bin/appchat/get", req, &resp, true)
	if err != nil {
		return respAppchatGet{}, err
	}

	return resp, nil
}

// execMessageSend 发送应用消息
func (c *WorkwxApp) execMessageSend(req reqMessage) (respMessageSend, error) {
	var resp respMessageSend
	err := c.executeQyapiJSONPost("/cgi-bin/message/send", req, &resp, true)
	if err != nil {
		return respMessageSend{}, err
	}

	return resp, nil
}

// execAppchatSend 应用推送消息
func (c *WorkwxApp) execAppchatSend(req reqMessage) (respMessageSend, error) {
	var resp respMessageSend
	err := c.executeQyapiJSONPost("/cgi-bin/appchat/send", req, &resp, true)
	if err != nil {
		return respMessageSend{}, err
	}

	return resp, nil
}

// execMediaUpload 上传临时素材
func (c *WorkwxApp) execMediaUpload(req reqMediaUpload) (respMediaUpload, error) {
	var resp respMediaUpload
	err := c.executeQyapiMediaUpload("/cgi-bin/media/upload", req, &resp, true)
	if err != nil {
		return respMediaUpload{}, err
	}

	return resp, nil
}

// execMediaUploadImg 上传永久图片
func (c *WorkwxApp) execMediaUploadImg(req reqMediaUploadImg) (respMediaUploadImg, error) {
	var resp respMediaUploadImg
	err := c.executeQyapiMediaUpload("/cgi-bin/media/uploadimg", req, &resp, true)
	if err != nil {
		return respMediaUploadImg{}, err
	}

	return resp, nil
}
