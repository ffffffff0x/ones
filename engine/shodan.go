package engine

import (
	"fmt"
	ones "ones/mod"
)

func TodoShodan() {

	ShodanKeyValue := string(ones.Confs["shodan_key"])
	fmt.Println(ShodanKeyValue[1 : len(ShodanKeyValue)-1])

}
