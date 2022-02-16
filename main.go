package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Factro Task Replacer v1")

	//get console inputs
	fmt.Println("Mit welchem JWT sollen Aufgaben gelesen werden?")
	fmt.Print("-> ")
	getJWT, _ := reader.ReadString('\n')
	// clean console input from returns and newlines
	getJWT = strings.Replace(getJWT, "\r\n", "", -1)

	fmt.Println("Mit welchem JWT sollen Aufgaben geupdated werden?")
	fmt.Print("-> ")
	postJWT, _ := reader.ReadString('\n')
	// clean console input from returns and newlines
	postJWT = strings.Replace(postJWT, "\r\n", "", -1)

	fmt.Println("Welches Feld soll ersetzt werden?")
	fmt.Print("-> ")
	fieldToReplace, _ := reader.ReadString('\n')
	// clean console input from returns and newlines
	fieldToReplace = strings.Replace(fieldToReplace, "\r\n", "", -1)

	fmt.Println("Welcher Wert soll ersetzt werden?")
	fmt.Print("-> ")
	valueToFind, _ := reader.ReadString('\n')
	// clean console input from returns and newlines
	valueToFind = strings.Replace(valueToFind, "\r\n", "", -1)

	fmt.Println("Durch welchen Wert soll ersetzt werden?")
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
	fmt.Println(getJWT)
	req.Header.Add("Authorization", getJWT)
	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		time.Sleep(10 * time.Minute)
		return
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		time.Sleep(10 * time.Minute)
		return
	}

	var tasks []map[string]interface{}

	err = json.Unmarshal(responseBody, &tasks)
	if err != nil {
		fmt.Println(err)
		time.Sleep(10 * time.Minute)
		return
	}

	//find and replace
	for _, task := range tasks {
		fieldValue, ok := task[fieldToReplace].(string)
		if !ok {
			fmt.Println("Das angefragte Feld enthält keinen String")
			time.Sleep(10 * time.Minute)
			return
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
		time.Sleep(10 * time.Minute)
		return
	}

	fmt.Println(string(reqBody))

	req, err = http.NewRequest("PUT", "https://cloud.factro.com/api/core/tasks/tasks", bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println(err)
		time.Sleep(10 * time.Minute)
		return
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", postJWT)
	req.Header.Add("Content-Type", "application/json")

	response, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		time.Sleep(10 * time.Minute)
		return
	}

	responseBody, err = io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		time.Sleep(10 * time.Minute)
		return
	}

	fmt.Printf("Server responded: %v\n", string(responseBody))
	fmt.Println("Tasks wurden erfolgreich aktualisiert!")
	time.Sleep(10 * time.Minute)
}
