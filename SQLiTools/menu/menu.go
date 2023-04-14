package menu

import (
	"fmt"
	"pstoolsmod/sql_attacks"
	"log"
	"os"
)

func Display_menu() string {
	content, err := os.ReadFile("menu/menu.txt")

	if err != nil {
		fmt.Println("Error Oops.. Could not open menu. Exiting...")
		log.Fatal(err)
	}

	return string(content)
}

func get_required_input_data(url *string, payload *string, with_payload bool) {
	fmt.Println("Enter URL:\n")
	fmt.Scanln(url)

	if with_payload {
		fmt.Println("Enter payload:\n")
		fmt.Scanln(url)
	}
}

func Choose_action() {
	var input int
	var url string
	var payload string


	for loopCount := 0; loopCount < loopCount + 1; loopCount++ {

		fmt.Println("Choose action: ")
		fmt.Scanln(&input)
		switch input {
		case 1:
			{
				get_required_input_data(&url, &payload, false)
				response, _ := sql_attacks.Get_html(url, "")
				fmt.Println(response)
			}
		case 2:
			{
				get_required_input_data(&url, &payload, true)
				response, _ := sql_attacks.Get_html(url, "")
				fmt.Println(response)
			}
		case 3:
			{
				var max_itr int
				get_required_input_data(&url, &payload, false)
				fmt.Println("Enter max number of columns to test:")
				fmt.Scanln(&max_itr)

				err := sql_attacks.Get_columns_sql_attack(url, max_itr)

				if err == 999 {
					fmt.Println("\n\nSQLi attack failed...")
				}
			}
		case 4:
			{
				var max_itr int
				get_required_input_data(&url, &payload, false)
				fmt.Println("Enter max number of columns to test:")
				fmt.Scanln(&max_itr)

				err := sql_attacks.Columns_containing_text(url, max_itr)

				if err == 999 {
					fmt.Println("\n\nSQLi attack failed...")
				}
			}
		case 5:
			{
				get_required_input_data(&url, nil, false)
				body, err := sql_attacks.Get_robots(url)

				if (err == nil) {
					fmt.Println(body)
				}
			}
		case 6:
			return
		}
		}
	

}
