package main

import "github.com/gin-gonic/gin"

func Register(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "注册",
	})
}

func Login(c *gin.Context) {
  c.JSON(200,gin.H{
    "msg" :"登录"
  })
}

func Captcha(c *gin.Context) {
  c.JSON(200,gin.H{
    "msg" :"验证码"
  })
}

func SearchSong(c *gin.Context) {
  c.JSON(200,gin.H{
    "msg" :"搜索歌曲"
  })
}
func main() {
    //启动配置
	router := gin.Default()
	//接口
	
	//注册
	router.POST("/V1/Register", Register)
	//登录
	router.POST("/V1/Login", Login)
	//验证码
	router.GET("/V1/Captcha", Captcha)
	//搜索歌曲
	router.GET("/V1/SearchSong", SearchSong)

	//收藏列表
	//添加收藏
	//删除收藏

	//歌曲下载
	//歌曲上传
	//歌曲举报
	//歌曲删除
	//歌曲评论
	//删除评论
	
	//帖子列表
	//发布帖子
	//修改帖子
	//删除帖子
	//帖子浏览
	//帖子点赞
	//帖子评论
	//删除评论
	//搜索帖子
	
	//房间列表
	//房间连接
	//聊天记录
	//搜索房间
	
	//热门歌曲
	//热门帖子
	//热门房间
	
	
	//router.POST("/somePost", posting)
	//router.PUT("/somePut", putting)
	//router.DELETE("/someDelete", deleting)
	//router.PATCH("/somePatch", patching)
	//router.HEAD("/someHead", head)
	//router.OPTIONS("/someOptions", options)
	router.Run()
}
