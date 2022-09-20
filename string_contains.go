package main

import (
	"fmt"
	"strings"
)

func main() {
	s := `<?xml version="1.0" encoding="utf-8"?>
<root><![CDATA[抱歉，您所在的用户组每小时限制发回帖 255 个，请稍候再发表<script type="text/javascript" reload="1">if(typeof errorhandle_reply=='function') {errorhandle_reply('抱歉，您所在的用户组每小时限制发回帖 255 个，请稍候再发表', {'posts_per_hour':'255'});}</script>]]></root>`

	b := strings.Contains(s, "抱歉")
	fmt.Println(b)
}
