package engine

import (
	"fmt"
	ones "ones/mod"
)

func TodoZoomeye() {

	ZoomKeyValue := string(ones.Confs["zoom_key"])
	fmt.Println(ZoomKeyValue[1 : len(ZoomKeyValue)-1])

}
