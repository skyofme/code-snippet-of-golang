package controllers

import (
	"crawl_movie/models"
	"github.com/astaxie/beego"
  "github.com/astaxie/beego/httplib"
  "time"
)

type CrawlMovieController struct {
	beego.Controller
}

/**
 目前这个爬虫只能爬取静态数据 对于像京东的部分动态数据 无法爬取
 对于动态数据 可以采用 一个组件 phantomjs
*/
func (c *CrawlMovieController) CrawlMovie() {
  var movieInfo models.MovieInfo
  //连接到redis
  models.ConnectRedis("127.0.0.1:6379")
    
  //爬虫入口url
  sUrl := "https://movie.douban.com/subject/25827935/"
  models.PutinQueue(sUrl)

  for{
        length := models.GetQueueLength()
        if length == 0{
            break //如果url队列为空 则退出当前循环
        }

        sUrl = models.PopfromQueue()
        //我们应当判断sUrl是否应该被访问过
        if models.IsVisit(sUrl){
           continue
        }

        rsp := httplib.Get(sUrl)
        sMovieHtml,err := rsp.String()
        if err != nil{
            panic(err)
        }

        movieInfo.Movie_name            = models.GetMovieName(sMovieHtml)
        //记录电影信息
        if movieInfo.Movie_name != ""{
            movieInfo.Movie_director        = models.GetMovieDirector(sMovieHtml)
            movieInfo.Movie_main_character  = models.GetMovieMainCharacters(sMovieHtml) 
            movieInfo.Movie_type            = models.GetMovieGenre(sMovieHtml)
            movieInfo.Movie_on_time         = models.GetMovieOnTime(sMovieHtml)
            movieInfo.Movie_grade           = models.GetMovieGrade(sMovieHtml)
            movieInfo.Movie_span            = models.GetMovieRunningTime(sMovieHtml)
            
            models.AddMovie(&movieInfo)
        }

        //提取该页面的所有连接
        urls := models.GetMovieUrls(sMovieHtml)

        for _,url := range urls{
            models.PutinQueue(url)
            c.Ctx.WriteString("<br>" + url + "</br>")
        }

        //sUrl 应当记录到 访问set中
        models.AddToSet(sUrl)

        time.Sleep(time.Second)
    }

    c.Ctx.WriteString("end of crawl!")
}
