package thread

import (
	"testing"

	"fmt"
)

func TestConvertHtml(t *testing.T) {
	h, _ := ConvertHtml(` 宣传帖链接：<a href="https://www.mcbbs.net/forum.php?mod=viewthread&amp;tid=1145121&amp;page=1#pid20532128" target="_blank">https://www.mcbbs.net/forum.php? ... ;page=1#pid20532128</a><br />\r\n图片地址：<img id="aimg_oy8Z8" onclick="zoom(this, this.src, 0, 0, 0)" class="zoom" file="https://sc01.alicdn.com/kf/H43659cd647cf48d2accd6d7de244f802e.jpg" onmouseover="img_onmouseoverfunc(this)" lazyloadthumb="1" border="0" alt="" /><br />\r\n是否为本人帖子：否</td></tr></table>\n\n\n</div>`)
	fmt.Println(h)
}
