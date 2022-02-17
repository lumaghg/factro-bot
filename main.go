package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Factro Task Replacer v1")
	firstIteration := true
	for {
		if !firstIteration {
			fmt.Println("Programm beenden? (J/N)")
			fmt.Print("-> ")
			userInput, _ := reader.ReadString('\n')
			// clean console input from returns and newlines
			userInput = strings.Replace(userInput, "\r\n", "", -1)
			if userInput == "J" {
				break
			}
		}
		firstIteration = false

		//get testuser api token from file
		filePath := path.Join("config", "api_user_token.txt")
		fileContent, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Println(err)
			time.Sleep(10 * time.Minute)
			return
		}
		getJWT := string(fileContent)
		getJWT = strings.Replace(getJWT, "\n", "", -1)
		fmt.Println("API Token used: ", getJWT)

		fmt.Println("Mit welchem JWT (Token API Admin) sollen Aufgaben geupdated werden?")
		fmt.Print("-> ")
		postJWT, _ := reader.ReadString('\n')
		// clean console input from returns and newlines
		postJWT = strings.Replace(postJWT, "\r\n", "", -1)

		fmt.Println("Welches Feld soll ersetzt werden? (Um den Titel zu ändern, hier title eingeben)")
		fmt.Print("-> ")
		fieldToReplace, _ := reader.ReadString('\n')
		// clean console input from returns and newlines
		fieldToReplace = strings.Replace(fieldToReplace, "\r\n", "", -1)

		fmt.Println("Welcher Wert soll ersetzt werden? (Beispiel: Kursnummer = MM.YYYY)")
		fmt.Print("-> ")
		valueToFind, _ := reader.ReadString('\n')
		// clean console input from returns and newlines
		valueToFind = strings.Replace(valueToFind, "\r\n", "", -1)

		fmt.Println("Durch welchen Wert soll ersetzt werden? (Beispiel: Kursnummer = MM.YYYY)")
		fmt.Print("-> ")
		valueToInsert, _ := reader.ReadString('\n')
		// clean console input from returns and newlines
		valueToInsert = strings.Replace(valueToInsert, "\r\n", "", -1)

		client := &http.Client{}
		//request tasks from factro
		req, err := http.NewRequest("GET", "https://cloud.factro.com/api/core/tasks", nil)
		if err != nil {
			fmt.Println(err)
			time.Sleep(10 * time.Minute)
			return
		}
		req.Header.Add("accept", "application/json")
		req.Header.Add("Authorization", getJWT)
		response, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			continue
		}

		defer response.Body.Close()
		responseBody, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
			continue
		}

		var tasks []map[string]interface{}

		err = json.Unmarshal(responseBody, &tasks)
		if err != nil {
			fmt.Println(err)
			continue
		}

		//find and replace
		for _, task := range tasks {
			fieldValue, ok := task[fieldToReplace].(string)
			if !ok {
				fmt.Println("Das angefragte Feld enthält keinen String")
				continue
			}

			if strings.Contains(fieldValue, valueToFind) {
				newFieldValue := strings.ReplaceAll(fieldValue, valueToFind, valueToInsert)
				task[fieldToReplace] = newFieldValue
			}
		}

		fmt.Printf("new tasks: %v\n", tasks)

		//update tasks in factro
		reqBody, err := json.Marshal(tasks)
		if err != nil {
			fmt.Println(err)
			continue
		}

		req, err = http.NewRequest("PUT", "https://cloud.factro.com/api/core/tasks/tasks", bytes.NewBuffer(reqBody))
		if err != nil {
			fmt.Println(err)
			continue
		}
		req.Header.Add("accept", "application/json")
		req.Header.Add("Authorization", postJWT)
		req.Header.Add("Content-Type", "application/json")

		_, err = client.Do(req)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("Tasks wurden erfolgreich aktualisiert! Ansicht in Factro aktualisieren und überprüfen.")
		time.Sleep(10 * time.Minute)

	}
}
