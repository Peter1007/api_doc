package controllers

import (
	"code.oa.com/sean/api-doc/modules"
	//	"encoding/base64"
	"fmt"
	"html/template"
	//	"io/ioutil"
	"log"
	"net/http"
	//	"net/url"
	"strconv"
)

type QueryController struct {
	resp   http.ResponseWriter
	req    *http.Request
	module *modules.QueryModule
}

func (query *QueryController) SlowList(resp http.ResponseWriter, req *http.Request) {
	date := req.FormValue("date")
	page, _ := strconv.Atoi(req.FormValue("page"))
	count, _ := strconv.Atoi(req.FormValue("count"))

	list, total, _ := query.module.SlowQuery(date, page, count)
	t, err := template.ParseFiles("views/query/slow_list.tmpl")
	if err != nil {
		log.Panic(err)
	}

	t.ExecuteTemplate(resp, "slow_list", list)
	t.ExecuteTemplate(resp, "slow_list", total)
	fmt.Println(total)
}

//func (query *QueryController) makePage(url string, page int, count int) string {
//	str := ""
//        if(count <= 1)	return "<span class=\"page\">1</span>"
//        if($page > 1)	$str .= '<a href="' . $url . '&page=' .($page - 1) . '">﹤</a>';
//        if($pageCount <= 7)
//        {
//            for($i = 1; $i <= $pageCount; $i ++)
//            {
//                if($i == $page)
//                    $str .= '<span class="page">' . $i . '</span>';
//                else
//                    $str .= '<a href="' . $url . '&page=' . $i . '">' . $i . '</a>';
//            }
//        }
//        elseif($page <= 4)
//        {
//            for($i = 1; $i <= 6; $i ++)
//            {
//            if($i == $page)
//                $str .= '<span class="page">' . $i . '</span>';
//            else
//                $str .= '<a href="' . $url . '&page=' . $i . '">' . $i . '</a>';
//            }
//            if(( $pageCount - 6) > 2)
//                $str .= '<span>...</span><a href="' . $url . '&page=' .($pageCount - 1) . '">' .($pageCount - 1) . '</a><a href="' . $url . '&page=' . $pageCount . '">' . $pageCount . '</a>';
//        }
//        elseif($page >= $pageCount - 3)
//        {
//            $str .= '<a href="' . $url . '&page=1">1</a><a href="' . $url . '&page=2">2</a><span>...</span>';
//            for($i =($pageCount - 5); $i <= $pageCount; $i ++)
//            {
//                if($i == $page)
//                    $str .= '<span class="page">' . $i . '</span>';
//                else
//                    $str .= '<a href="' . $url . '&page=' . $i . '">' . $i . '</a>';
//            }
//        }
//        else
//        {
//            $str .= '<a href="' . $url . '&page=1">1</a><a href="' . $url . '&page=2">2</a><span>...</span>';
//            for($i =($page - 2); $i <=($page + 2); $i ++)
//            {
//                if($i == $page)
//                    $str .= '<span class="page">' . $i . '</span>';
//                else
//                    $str .= '<a href="' . $url . '&page=' . $i . '">' . $i . '</a>';
//            }
//            $str .= '<span>...</span><a href="' . $url . '&page=' .($pageCount - 1) . '">' .($pageCount - 1) . '</a><a href="' . $url . '&page=' . $pageCount . '">' . $pageCount . '</a>';
//        }
//        if($page < $pageCount)
//            $str .= '<a href="' . $url . '&page=' .($page + 1) . '">﹥</a>';

//        return $str;
//}
