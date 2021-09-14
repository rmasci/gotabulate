package main

import (
	"fmt"

	"github.com/rmasci/gotabulate"
)

func main() {

	csv2Print := gotabulate.NewCSVOut()
	csv2Print.Text = `date,mysql,percent,iblogs,percent,backup,percent,binlogs,percent,data,percent
2021-03-02 22:53,5.1G,12%,6.1G,71%,663G,55%,6.1G,2%,561G,40%
2021-03-02 22:54,5.1G,12%,6.1G,71%,664G,55%,6.1G,2%,561G,40%
2021-03-03 00:00,5.0G,12%,6.1G,71%,688G,57%,6.6G,2%,506G,37%
2021-03-03 04:00,5.0G,12%,6.1G,71%,795G,66%,6.0G,2%,508G,37%`
	csv2Print.Render="mysql"
	fmt.Println(csv2Print.Table())
}
