package api

import (
	"bytes"
	"os"
	"log"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"crypto/tls"
	"encoding/xml"
	"strings"
	"errors"
	"github.com/hashicorp/terraform/helper/schema"
	"os/exec"
)

type XMLElement struct {
	Key   string `xml:"name,attr"`
	Value string `xml:",chardata"`
}

type Linklist struct {
	XMLName struct{}    `xml:"Linklist"`
	XMLNS   string `xml:"xmlns,attr,omitempty"`
	Llist   []Link `xml:"link,omitempty"`
}

type Link struct {
	XMLName struct{}    `xml:"link"`
	Name        string      `xml:"name"`
	HREF        string      `xml:"href"`
	Id          string      `xml:"id"`
	Rel         string      `xml:"rel,omitempty"`
	Type        string      `xml:"type,omitempty"`
	Position    string      `xml:"position,omitempty"`
}

type Result struct {
	XMLName     struct{}    `xml:"link"`
	Name        string      `xml:"name"`
	Href        string      `xml:"href"`
	Id          string      `xml:"id"`
	Rel         string      `xml:"rel"`
	Type        string      `xml:"type"`
	Position    string      `xml:"position"`
}

type Config struct {
	AccessKey  string
	SecretKey  string
	MaxRetries string
	APIURL     string
	AWSSmall      string
	AWSMedium     string
	AWSLarge      string
	BCSmall      string
	BCMedium     string
	BCLarge      string
}

var configuration Config
var payload string
func init(){
	file, err1 := os.Open("./agility/api/conf.json")
	if err1 != nil {
		log.Println("file not found", err1)
	}
	decoder := json.NewDecoder(file)
	configuration = Config{}
	err := decoder.Decode(&configuration)
	if err != nil {
		log.Println("not able to decode", err)
	}

	/*err2 := file.Close()
	log.Printf("not able to close %v\n", err2)*/

}

func SimpleBlueprintDeploy(blueprintId string, environmentId string, username string, password string) []byte {
	//set up logging
	f, errf := os.OpenFile("agility.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if errf != nil {
		log.Println("error opening file: ", errf)
	}
	defer f.Close()

	log.SetOutput(f)

	var url bytes.Buffer

	// Create the URL for the call to the Agility API
	url.WriteString(configuration.APIURL)
	url.WriteString("current/blueprint/")
	url.WriteString(blueprintId)
	url.WriteString("/simpledeploy/")
	url.WriteString(environmentId)
	log.Println("URL:>", url.String())

	// Set the right HTTP Verb, and setup HTTP Basic Security
	req, err := http.NewRequest("POST", url.String(), nil)
	req.SetBasicAuth(username,password)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// make the HTTPS request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// log the response details for debugging
	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)

	//Stream the response body into a byte array and return it
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))
	return body
}

func DeploymentPlanBlueprintDeploy(blueprintId string, environmentId string, deploymentPlan string, username string, password string) []byte {
	var url bytes.Buffer
	// Create the URL for the call to the Agility API
	url.WriteString(configuration.APIURL)
	url.WriteString("current/blueprint/")
	url.WriteString(blueprintId)
	url.WriteString("/deploy/")
	url.WriteString(environmentId)
	log.Println("URL:>", url.String())

	// Set the right HTTP Verb, and setup HTTP Basic Security
	req, err := http.NewRequest("POST", url.String(), bytes.NewBuffer([]byte(deploymentPlan)))
	req.Header.Set("Content-Type", "application/xml; charset=utf-8")
	req.SetBasicAuth(username,password)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// make the HTTPS request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// log the response details for debugging
	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)

	//Stream the response body into a byte array and return it
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))
	return body
}

func GetBlueprintDetail(blueprintId string, username string, password string) []byte {
	var url bytes.Buffer
	// Create the URL for the call to the Agility API
	url.WriteString(configuration.APIURL)
	url.WriteString("current/blueprint/")
	url.WriteString(blueprintId)
	log.Println("URL:>", url.String())

	// Set the right HTTP Verb, and setup HTTP Basic Security
	req, err := http.NewRequest("GET", url.String(), nil)
	req.SetBasicAuth(username,password)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// make the HTTPS request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// log the response details for debugging
	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)

	//Stream the response body into a byte array and return it
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))
	return body
}

func StartTopology(topologyId string, username string, password string) []byte {
	var url bytes.Buffer
	// Create the URL for the call to the Agility API
	url.WriteString(configuration.APIURL)
	url.WriteString("current/topology/")
	url.WriteString(topologyId)
	url.WriteString("/start")
	log.Println("URL:>", url.String())

	// Set the right HTTP Verb, and setup HTTP Basic Security
	req, err := http.NewRequest("POST", url.String(), nil)
	req.SetBasicAuth(username,password)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// make the HTTPS request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// log the response details for debugging
	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)

	//Stream the response body into a byte array and return it
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))
	return body
}

func StopTopology(topologyId string, username string, password string) []byte {
	var url bytes.Buffer
	// Create the URL for the call to the Agility API
	url.WriteString(configuration.APIURL)
	url.WriteString("current/topology/")
	url.WriteString(topologyId)
	url.WriteString("/stop")
	log.Println("URL:>", url.String())

	// Set the right HTTP Verb, and setup HTTP Basic Security
	req, err := http.NewRequest("POST", url.String(), nil)
	req.SetBasicAuth(username,password)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// make the HTTPS request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// log the response details for debugging
	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)

	//Stream the response body into a byte array and return it
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))
	return body
}

func DestroyTopology(topologyId string, username string, password string) []byte {
	log.Println("topologyId is:", topologyId)
	var url bytes.Buffer
	// Create the URL for the call to the Agility API
	url.WriteString(configuration.APIURL)
	url.WriteString("current/topology/")
	url.WriteString(topologyId)
	log.Println("URL:>", url.String())

	// Set the right HTTP Verb, and setup HTTP Basic Security
	req, err := http.NewRequest("DELETE", url.String(), nil)
	req.SetBasicAuth(username,password)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// make the HTTPS request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// log the response details for debugging
	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)

	//Stream the response body into a byte array and return it
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))
	if resp.Status[:3] != "202" {
		return nil
	} else {
		return body
	}
}

func GetDeploymentPlans(blueprintId string, environmentId string, username string, password string) []byte {
	var url bytes.Buffer
	// Create the URL for the call to the Agility API
	url.WriteString(configuration.APIURL)
	url.WriteString("current/blueprint/")
	url.WriteString(blueprintId)
	url.WriteString("/deploymentplan/")
	url.WriteString(environmentId)
	log.Println("URL:>", url.String())

	// Set the right HTTP Verb, and setup HTTP Basic Security
	req, err := http.NewRequest("GET", url.String(), nil)
	req.SetBasicAuth(username, password)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// make the HTTPS request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// log the response details for debugging
	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)

	//Stream the response body into a byte array and return it
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))
	return body
}

func GetTopologyDetail(topologyId string, username string, password string) []byte {
	var url bytes.Buffer
	// Create the URL for the call to the Agility API
	url.WriteString(configuration.APIURL)
	url.WriteString("current/topology/")
	url.WriteString(topologyId)
	log.Println("URL:>", url.String())

	// Set the right HTTP Verb, and setup HTTP Basic Security
	req, err := http.NewRequest("GET", url.String(), nil)
	req.SetBasicAuth(username, password)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// make the HTTPS request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// log the response details for debugging
	log.Println("response Status:", resp.Status)

	//Stream the response body into a byte array and return it
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))
	return body
}

func UpdateTopology(topologyId string, toplogy string, username string, password string) []byte {
	var url bytes.Buffer
	// Create the URL for the call to the Agility API
	url.WriteString(configuration.APIURL)
	url.WriteString("current/topology/")
	url.WriteString(topologyId)
	log.Println("URL:>", url.String())

	// Set the right HTTP Verb, and setup HTTP Basic Security
	req, err := http.NewRequest("PUT", url.String(), bytes.NewBuffer([]byte(toplogy)))
	req.Header.Set("Content-Type", "application/xml; charset=utf-8")
	req.SetBasicAuth(username, password)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// make the HTTPS request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// log the response details for debugging
	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)

	//Stream the response body into a byte array and return it
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))
	return body
}

func GetTaskStatus(taskId string, username string, password string) []byte {
	var url bytes.Buffer
	// Create the URL for the call to the Agility API
	url.WriteString(configuration.APIURL)
	url.WriteString("current/task/")
	url.WriteString(taskId)
	log.Println("URL:>", url.String())

	// Set the right HTTP Verb, and setup HTTP Basic Security
	req, err := http.NewRequest("GET", url.String(), nil)
	req.SetBasicAuth(username, password)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// make the HTTPS request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// log the response details for debugging
	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)

	//Stream the response body into a byte array and return it
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))
	return body
}

func GetProjectId(projectName string, username string, password string) (string, error) {
	log.Println("projectName is: ", projectName)
	var url bytes.Buffer
	q := new(Result)

	// Create the URL for the call to the Agility API
	url.WriteString(configuration.APIURL)
	url.WriteString("current/project")
	log.Println("URL:>", url.String())

	// Set the right HTTP Verb, and setup HTTP Basic Security
	req, err := http.NewRequest("GET", url.String(), nil)
	req.SetBasicAuth(username, password)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// make the HTTPS request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// log the response details for debugging
	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)

	//Stream the response body into a byte array
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))

	//Parse the XML
	r := strings.NewReader(string(body))
	decoder := xml.NewDecoder(r)
	finish := false
	for {
		// Read tokens from the XML document in a stream.
		t, _ := decoder.Token()
		if t == nil {
			return "", errors.New("there are no Projects with this name")
		}
		if finish {
			break
		}
		// look for <link> element
		switch Element := t.(type) {
		case xml.StartElement:
			if Element.Name.Local == "link" {
				log.Println("Element name is : ", Element.Name.Local)

				// unmarshal the element into generic structure
				err := decoder.DecodeElement(&q, &Element)
				if err != nil {
					log.Println(err)
				}

				// if the project name matches the project defined to Terraform
				// then we are are the right place, so stop looking
				log.Println("Element value is :", string(q.Name))
				if string(q.Name) == projectName {
					log.Println("Found the Project : ", q.Name)
					finish = true
					break
				}
			}
			// if the element is the <Linklist> then go again
			if Element.Name.Local == "Linklist" {
				log.Println("Element name is : ", Element.Name.Local)
			} else {
				log.Println("Unknown Element name is : ", Element.Name.Local)
			}
		default:
		}

	}

	// return the ID for the project
	return string(q.Id), nil
}

func SearchTemplates(user string, username string, password string) []byte {
	var url bytes.Buffer
	// Create the URL for the call to the Agility API
	url.WriteString(configuration.APIURL)
	url.WriteString("current/template/search?qterm.field.creator.name=")
	url.WriteString(user)
	log.Println("URL:>", url.String())

	// Set the right HTTP Verb, and setup HTTP Basic Security
	req, err := http.NewRequest("GET", url.String(), nil)
	req.SetBasicAuth(username, password)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// make the HTTPS request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// log the response details for debugging
	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)

	//Stream the response body into a byte array and return it
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))
	return body
}

func GetInstanceDetail(instanceId string, username string, password string) []byte {
	var url bytes.Buffer
	// Create the URL for the call to the Agility API
	url.WriteString(configuration.APIURL)
	url.WriteString("current/compute/")
	url.WriteString(instanceId)
	log.Println("URL:>", url.String())

	// Set the right HTTP Verb, and setup HTTP Basic Security
	req, err := http.NewRequest("GET", url.String(), nil)
	req.SetBasicAuth(username, password)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// make the HTTPS request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// log the response details for debugging
	log.Println("response Status:", resp.Status)

	//Stream the response body into a byte array and return it
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))
	return body
}

func UpdateInstance(instanceId string, instance string, username string, password string) []byte {
	var url bytes.Buffer
	// Create the URL for the call to the Agility API
	url.WriteString(configuration.APIURL)
	url.WriteString("current/compute/")
	url.WriteString(instanceId)
	log.Println("URL:>", url.String())

	// Set the right HTTP Verb, and setup HTTP Basic Security
	req, err := http.NewRequest("PUT", url.String(), bytes.NewBuffer([]byte(instance)))
	req.Header.Set("Content-Type", "application/xml; charset=utf-8")
	req.SetBasicAuth(username, password)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// make the HTTPS request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// log the response details for debugging
	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)

	//Stream the response body into a byte array and return it
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))
	return body
}

func GetEnvironmentId(environmntName string, projectId string, username string, password string) (string, error) {
	var url bytes.Buffer
	q := new(Result)
	// Create the URL for the call to the Agility API
	url.WriteString(configuration.APIURL)
	url.WriteString("current/project/")
	url.WriteString(projectId)
	url.WriteString("/environment")
	log.Println("URL:>", url.String())

	// Set the right HTTP Verb, and setup HTTP Basic Security
	req, err := http.NewRequest("GET", url.String(), nil)
	req.SetBasicAuth(username, password)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// make the HTTPS request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// log the response details for debugging
	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)
	//Stream the response body into a byte array
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))

	//Parse the XML
	r := strings.NewReader(string(body))
	decoder := xml.NewDecoder(r)
	finish := false
	for {
		// Read tokens from the XML document in a stream.
		t, _ := decoder.Token()
		if t == nil {
			return "", errors.New("there are no Environments with this name")
		}
		if finish {
			break
		}
		// look for <link> element
		switch Element := t.(type) {
		case xml.StartElement:
			if Element.Name.Local == "link" {
				log.Println("Element name is : ", Element.Name.Local)

				// unmarshal the element into generic structure
				err := decoder.DecodeElement(&q, &Element)
				if err != nil {
					log.Println(err)
				}

				// if the environment name matches the environment defined to Terraform
				// then we are are the right place, so stop looking
				log.Println("Element value is :", string(q.Name))
				if string(q.Name) == environmntName {
					log.Println("Found the Environment : ", q.Name)
					finish = true
					break
				}
			}
			// if the element is the <Linklist> then go again
			if Element.Name.Local == "Linklist" {
				log.Println("Element name is : ", Element.Name.Local)
			} else {
				log.Println("Unknown Element name is : ", Element.Name.Local)
			}
		default:
		}

	}

	// return the ID for the environment
	return string(q.Id), nil

}

func GetBlueprintId(blueprintName string, projectId string, username string, password string) (string, error) {
	log.Println("The Blueprint name is: ", blueprintName)
	var url bytes.Buffer
	q := new(Result)

	// Create the URL for the call to the Agility API
	url.WriteString(configuration.APIURL)
	url.WriteString("current/project/")
	url.WriteString(projectId)
	url.WriteString("/blueprint")
	log.Println("URL:>", url.String())

	// Set the right HTTP Verb, and setup HTTP Basic Security
	req, err := http.NewRequest("GET", url.String(), nil)
	req.SetBasicAuth(username, password)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// make the HTTPS request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// log the response details for debugging
	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)

	//Stream the response body into a byte array
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))

	//Parse the XML
	r := strings.NewReader(string(body))
	decoder := xml.NewDecoder(r)
	finish := false
	for {
		// Read tokens from the XML document in a stream.
		t, _ := decoder.Token()
		if t == nil {
			return "", errors.New("there are no Blueprints with this name")
		}
		if finish {
			break
		}
		// look for <link> element
		switch Element := t.(type) {
		case xml.StartElement:
			if Element.Name.Local == "link" {
				log.Println("Element name is : ", Element.Name.Local)

				// unmarshal the element into generic structure
				err := decoder.DecodeElement(&q, &Element)
				if err != nil {
					log.Println(err)
				}

				// if the blueprint name matches the blueprint defined to Terraform
				// then we are are the right place, so stop looking
				log.Println("Element value is :", string(q.Name))
				if string(q.Name) == blueprintName {
					log.Println("Found the Blueprint : ", q.Name)
					finish = true
					break
				}
			}
		default:
		}

	}

	// return the ID for the blueprint
	return string(q.Id), nil
}

func GetBlueprintIdForVersion(blueprintName string, projectId string, version string, username string, password string) (string, error) {
	var url bytes.Buffer
	log.Println("The Blueprint name is: ", blueprintName)

	// call the internal function to get all the templates owned by the user and get the slot ID
	// for the storage of all the versions
	slotId, err := GetBlueprintVersionsSlot(blueprintName, projectId, version, username, password)
	if err != nil {
		return "", err
	}

	// stop if no slot ID as it means there are no versions
	if slotId == "" {
		return "", errors.New("there are no versions for this bluerint")
	}

	// Create the URL for the call to the Agility API
	// this gets all the versions for the blueprint
	url.WriteString(configuration.APIURL)
	url.WriteString("current/blueprint/")
	url.WriteString(slotId)
	url.WriteString("/version")
	log.Println("URL:>", url.String())

	// Set the right HTTP Verb, and setup HTTP Basic Security
	req, err := http.NewRequest("GET", url.String(), nil)
	req.SetBasicAuth(username, password)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// make the HTTPS request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// log the response details for debugging
	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))

	//unmarshall XML into a struct
	var list Linklist
	err = xml.Unmarshal(body, &list)
	if err != nil {
		log.Println(err)
		return "",err
	}
	log.Println("length of list.Llist is: ", len(list.Llist))
	var blueprintId string
	var element XMLElement
	finished := false

	// iterate through the struct looking for the right version, if there is a list
	if len(list.Llist) > 0 {
		for i := 0; i < len(list.Llist); i++ {
			if finished == true {
				break
			}
			log.Println("list.Llist[i].Name is: ", list.Llist[i].Name)
			log.Println("list.Llist[i].Id is: ", list.Llist[i].Id)

			// get the details of the blueprint for the current blueprint
			// in the list then parse the resulting XML
			statusResponse := GetBlueprintDetail(list.Llist[i].Id, username, password)
			sr := strings.NewReader(string(statusResponse))
			decoder := xml.NewDecoder(sr)
			for {
				if finished == true {
					break
				}
				// Read tokens from the XML document in a stream.
				st, _ := decoder.Token()
				if st == nil {
					break
				}

				// when we find the <version> element compare it with the one we are after
				switch Element := st.(type) {
				case xml.StartElement:
					if Element.Name.Local == "version" {
						log.Println("Element name is : ", Element.Name.Local)

						err := decoder.DecodeElement(&element, &Element)
						if err != nil {
							log.Println(err)
						}

						// if the values match we have found the right version
						log.Println("Element value is : ", element.Value)
						if element.Value == version {
							blueprintId = list.Llist[i].Id
							finished = true
							break
						}
					}
				default:
				}
			}
		}
	}

	// return the blueprint ID for the version we are after
	return blueprintId, nil

}

func GetBlueprintVersionsSlot(blueprintName string, projectId string, version string, username string, password string) (string, error) {
	// get the blueprint ID for the blueprint name within the project
	response, err := GetBlueprintId(blueprintName, projectId, username, password)
	if err != nil {
		return "", err
	}

	var url bytes.Buffer
	var slotId string
	var element XMLElement

	// Create the URL for the call to the Agility API
	// this gets detail the blueprint ID fetched above
	url.WriteString(configuration.APIURL)
	url.WriteString("current/blueprint/")
	url.WriteString(response)
	log.Println("URL:>", url.String())

	// Set the right HTTP Verb, and setup HTTP Basic Security
	req, err := http.NewRequest("GET", url.String(), nil)
	req.SetBasicAuth(username, password)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// make the HTTPS request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// log the response details for debugging
	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)
	// stream the result into a byte array
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))

	//Parse the XML
	r := strings.NewReader(string(body))
	decoder := xml.NewDecoder(r)
	finish := false

	for {
		// Read tokens from the XML document in a stream.
		t, _ := decoder.Token()
		if t == nil {
			break
		}
		if finish {
			break
		}
		// look for the <slotId> element
		switch Element := t.(type) {
		case xml.StartElement:
			if Element.Name.Local == "slotId" {
				log.Println("Element name is : ", Element.Name.Local)
				err := decoder.DecodeElement(&element, &Element)
				if err != nil {
					log.Println(err)
				}
				// finish once a slotId is found
				log.Println("Element value is : ", element.Value)
				slotId = element.Value
				finish = true
				break
			}
		default:
		}

	}

	return string(slotId), nil
}

func GetProject(projectId string, username string, password string) ([]byte, error) {
	var url bytes.Buffer
	// Create the URL for the call to the Agility API
	url.WriteString(configuration.APIURL)
	url.WriteString("current/project/")
	url.WriteString(projectId)
	log.Println("URL:>", url.String())

	// Set the right HTTP Verb, and setup HTTP Basic Security
	req, err := http.NewRequest("GET", url.String(), nil)
	req.SetBasicAuth(username, password)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// make the HTTPS request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// log the response details for debugging
	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)

	//Stream the response body into a byte array and return it
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))
	return body, nil
}

func GetContainerId(containerName string, username string, password string) (string, error){
	log.Println("containerName is: ", containerName)
	var url bytes.Buffer
	q := new(Result)

	// Create the URL for the call to the Agility API
	url.WriteString(configuration.APIURL)
	url.WriteString("current/container")
	log.Println("URL:>", url.String())

	// Set the right HTTP Verb, and setup HTTP Basic Security
	req, err := http.NewRequest("GET", url.String(), nil)
	req.SetBasicAuth(username, password)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// make the HTTPS request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// log the response details for debugging
	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)

	//Stream the response body into a byte array
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))

	//Parse the XML
	r := strings.NewReader(string(body))
	decoder := xml.NewDecoder(r)
	finish := false
	for {
		// Read tokens from the XML document in a stream.
		t, _ := decoder.Token()
		if t == nil {
			return "", errors.New("there are no Containers with this name")
		}
		if finish {
			break
		}
		// look for <link> element
		switch Element := t.(type) {
		case xml.StartElement:
			if Element.Name.Local == "link" {
				log.Println("Element name is : ", Element.Name.Local)

				// unmarshal the element into generic structure
				err := decoder.DecodeElement(&q, &Element)
				if err != nil {
					log.Println(err)
				}

				// if the container name matches the container defined to Terraform
				// then we are are the right place, so stop looking
				log.Println("Element value is :", string(q.Name))
				if string(q.Name) == containerName {
					log.Println("Found the Container : ", q.Name)
					finish = true
					break
				}
			}
			// if the element is the <Linklist> then go again
			if Element.Name.Local == "Linklist" {
				log.Println("Element name is : ", Element.Name.Local)
			} else {
				log.Println("Unknown Element name is : ", Element.Name.Local)
			}
		default:
		}

	}

	// return the ID for the container
	return string(q.Id), nil
}

func CreateSubContainer(ResourceData *schema.ResourceData, containerId string, username string, password string) ([]byte, error){
	f, errf := os.OpenFile("./agility/api/agility.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if errf != nil {
		log.Println("error opening file: ", errf)
	}
	defer f.Close()

	log.SetOutput(f)

	var url bytes.Buffer
	// Create the URL for the call to the Agility API
	url.WriteString(configuration.APIURL)
	url.WriteString("current/container/")
	url.WriteString(containerId)
	url.WriteString("/container")
	log.Println("URL:>", url.String())

	var payload bytes.Buffer
	//Create the payload for the request body
	s := ResourceData.Get("container").(string)
	t := ResourceData.Get("description").(string)
	payload.WriteString(`<?xml version="1.0" encoding="UTF-8" standalone="yes"?><ns1:Container xmlns:ns1="http://servicemesh.com/agility/api"><ns1:name>`)
	payload.WriteString(s)
	payload.WriteString(`</ns1:name><ns1:description>`)
	payload.WriteString(t)
	payload.WriteString(`</ns1:description></ns1:Container>`)
	payload1 := payload.String()
	log.Println("Payload=====>",payload1)
	payload2 := []byte(payload1)
	// Set the right HTTP Verb, and setup HTTP Basic Security
	req, err := http.NewRequest("POST", url.String(), bytes.NewBuffer([]byte(payload2)))
	req.Header.Set("Content-Type", "application/xml; charset=utf-8")
	req.SetBasicAuth(username,password)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// make the HTTPS request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// log the response details for debugging
	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)

	//Stream the response body into a byte array and return it
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))
	return body,nil
}

func CreateSubProject(ResourceData *schema.ResourceData, containerId string, username string, password string) ([]byte, error){
	f, errf := os.OpenFile("./agility/api/agility.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if errf != nil {
		log.Println("error opening file: ", errf)
	}
	defer f.Close()

	log.SetOutput(f)

	var url bytes.Buffer
	// Create the URL for the call to the Agility API
	url.WriteString(configuration.APIURL)
	url.WriteString("current/container/")
	url.WriteString(containerId)
	url.WriteString("/project")
	log.Println("URL:>", url.String())

	var payload bytes.Buffer
	//Create the payload for the request body
	s := ResourceData.Get("project").(string)
	t := ResourceData.Get("description").(string)
	payload.WriteString(`<?xml version="1.0" encoding="UTF-8" standalone="yes"?><ns1:Project xmlns:ns1="http://servicemesh.com/agility/api"><ns1:name>`)
	payload.WriteString(s)
	payload.WriteString(`</ns1:name><ns1:description>`)
	payload.WriteString(t)
	payload.WriteString(`</ns1:description></ns1:Project>`)
	payload1 := payload.String()
	log.Println("Payload=====>",payload1)
	payload2 := []byte(payload1)
	// Set the right HTTP Verb, and setup HTTP Basic Security
	req, err := http.NewRequest("POST", url.String(), bytes.NewBuffer([]byte(payload2)))
	req.Header.Set("Content-Type", "application/xml; charset=utf-8")
	req.SetBasicAuth(username,password)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// make the HTTPS request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// log the response details for debugging
	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)

	//Stream the response body into a byte array and return it
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))
	return body,nil
}

func CreateEnvironments(ResourceData *schema.ResourceData, projectId string, username string, password string) ([]byte, error){
	f, errf := os.OpenFile("./agility/api/agility.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if errf != nil {
		log.Println("error opening file: ", errf)
	}
	defer f.Close()

	log.SetOutput(f)

	var url bytes.Buffer
	// Create the URL for the call to the Agility API
	url.WriteString(configuration.APIURL)
	url.WriteString("current/project/")
	url.WriteString(projectId)
	url.WriteString("/environment")
	log.Println("URL:>", url.String())

	var payload bytes.Buffer
	//Create the payload for the request body
	s := ResourceData.Get("environment").(string)
	t := ResourceData.Get("description").(string)
	u := ResourceData.Get("environmenttype").(string)
	payload.WriteString(`<?xml version="1.0" encoding="UTF-8" standalone="yes"?><ns1:Environment xmlns:ns1="http://servicemesh.com/agility/api"><ns1:name>`)
	payload.WriteString(s)
	payload.WriteString(`</ns1:name><ns1:type><ns1:name>`)
	payload.WriteString(u)
	payload.WriteString(`</ns1:name></ns1:type><ns1:description>`)
	payload.WriteString(t)
	payload.WriteString(`</ns1:description></ns1:Environment>`)
	payload1 := payload.String()
	log.Println("Payload=====>",payload1)
	payload2 := []byte(payload1)
	// Set the right HTTP Verb, and setup HTTP Basic Security
	req, err := http.NewRequest("POST", url.String(), bytes.NewBuffer([]byte(payload2)))
	req.Header.Set("Content-Type", "application/xml; charset=utf-8")
	req.SetBasicAuth(username,password)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// make the HTTPS request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// log the response details for debugging
	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)

	//Stream the response body into a byte array and return it
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))
	return body,nil
}

func DestroySubConatiner(containerId string, username string, password string) []byte {
	log.Println("containerId is:", containerId)
	var url bytes.Buffer
	// Create the URL for the call to the Agility API
	url.WriteString(configuration.APIURL)
	url.WriteString("current/container/")
	url.WriteString(containerId)
	log.Println("URL:>", url.String())

	// Set the right HTTP Verb, and setup HTTP Basic Security
	req, err := http.NewRequest("DELETE", url.String(), nil)
	req.SetBasicAuth(username,password)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// make the HTTPS request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// log the response details for debugging
	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)

	//Stream the response body into a byte array and return it
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))
	if resp.Status[:3] != "202" {
		return nil
	} else {
		return body
	}
}

func DestroySubProject(ResourceData *schema.ResourceData, containerId string, username string, password string) ([]byte, error){
	f, errf := os.OpenFile("./agility/api/agility.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if errf != nil {
		log.Println("error opening file: ", errf)
	}
	defer f.Close()

	log.SetOutput(f)

	var url bytes.Buffer
	// Create the URL for the call to the Agility API
	url.WriteString(configuration.APIURL)
	url.WriteString("current/container/")
	url.WriteString(containerId)
	url.WriteString("/project")
	log.Println("URL:>", url.String())

	var payload bytes.Buffer
	//Create the payload for the request body
	s := ResourceData.Get("project").(string)
	t := ResourceData.Get("description").(string)
	payload.WriteString(`<?xml version="1.0" encoding="UTF-8" standalone="yes"?><ns1:Project xmlns:ns1="http://servicemesh.com/agility/api"><ns1:name>`)
	payload.WriteString(s)
	payload.WriteString(`</ns1:name><ns1:description>`)
	payload.WriteString(t)
	payload.WriteString(`</ns1:description></ns1:Project>`)
	payload1 := payload.String()
	log.Println("Payload=====>",payload1)
	payload2 := []byte(payload1)
	// Set the right HTTP Verb, and setup HTTP Basic Security
	req, err := http.NewRequest("POST", url.String(), bytes.NewBuffer([]byte(payload2)))
	req.Header.Set("Content-Type", "application/xml; charset=utf-8")
	req.SetBasicAuth(username,password)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// make the HTTPS request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// log the response details for debugging
	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)

	//Stream the response body into a byte array and return it
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))
	return body,nil
}

func DestroyEnvironments(ResourceData *schema.ResourceData, projectId string, username string, password string) ([]byte, error){
	f, errf := os.OpenFile("./agility/api/agility.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if errf != nil {
		log.Println("error opening file: ", errf)
	}
	defer f.Close()

	log.SetOutput(f)

	var url bytes.Buffer
	// Create the URL for the call to the Agility API
	url.WriteString(configuration.APIURL)
	url.WriteString("current/project/")
	url.WriteString(projectId)
	url.WriteString("/environment")
	log.Println("URL:>", url.String())

	var payload bytes.Buffer
	//Create the payload for the request body
	s := ResourceData.Get("environment").(string)
	t := ResourceData.Get("description").(string)
	u := ResourceData.Get("environmenttype").(string)
	payload.WriteString(`<?xml version="1.0" encoding="UTF-8" standalone="yes"?><ns1:Environment xmlns:ns1="http://servicemesh.com/agility/api"><ns1:name>`)
	payload.WriteString(s)
	payload.WriteString(`</ns1:name><ns1:type><ns1:name>`)
	payload.WriteString(u)
	payload.WriteString(`</ns1:name></ns1:type><ns1:description>`)
	payload.WriteString(t)
	payload.WriteString(`</ns1:description></ns1:Environment>`)
	payload1 := payload.String()
	log.Println("Payload=====>",payload1)
	payload2 := []byte(payload1)
	// Set the right HTTP Verb, and setup HTTP Basic Security
	req, err := http.NewRequest("POST", url.String(), bytes.NewBuffer([]byte(payload2)))
	req.Header.Set("Content-Type", "application/xml; charset=utf-8")
	req.SetBasicAuth(username,password)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// make the HTTPS request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// log the response details for debugging
	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)

	//Stream the response body into a byte array and return it
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))
	return body,nil
}

func GetCloudId(cloudName string, username string, password string) (string, error){
	log.Println("cloudName is: ", cloudName)
	var url bytes.Buffer
	q := new(Result)

	// Create the URL for the call to the Agility API
	url.WriteString(configuration.APIURL)
	url.WriteString("current/cloud")
	log.Println("URL:>", url.String())

	// Set the right HTTP Verb, and setup HTTP Basic Security
	req, err := http.NewRequest("GET", url.String(), nil)
	req.SetBasicAuth(username, password)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// make the HTTPS request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// log the response details for debugging
	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)

	//Stream the response body into a byte array
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))

	//Parse the XML
	r := strings.NewReader(string(body))
	decoder := xml.NewDecoder(r)
	finish := false
	for {
		// Read tokens from the XML document in a stream.
		t, _ := decoder.Token()
		if t == nil {
			return "", errors.New("there are no Clouds with this name")
		}
		if finish {
			break
		}
		// look for <link> element
		switch Element := t.(type) {
		case xml.StartElement:
			if Element.Name.Local == "link" {
				log.Println("Element name is : ", Element.Name.Local)

				// unmarshal the element into generic structure
				err := decoder.DecodeElement(&q, &Element)
				if err != nil {
					log.Println(err)
				}

				// if the cloud name matches the cloud defined to Terraform
				// then we are are the right place, so stop looking
				log.Println("Element value is :", string(q.Name))
				if string(q.Name) == cloudName {
					log.Println("Found the Cloud : ", q.Name)
					finish = true
					break
				}
			}
			// if the element is the <Linklist> then go again
			if Element.Name.Local == "Linklist" {
				log.Println("Element name is : ", Element.Name.Local)
			} else {
				log.Println("Unknown Element name is : ", Element.Name.Local)
			}
		default:
		}

	}

	// return the ID for the cloud
	return string(q.Id), nil
}

func AddCloudProvider(ResourceData *schema.ResourceData, username string, password string) ([]byte, error){
	f, errf := os.OpenFile("./agility/api/agility.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if errf != nil {
		log.Println("error opening file: ", errf)
	}
	defer f.Close()

	log.SetOutput(f)

	var url bytes.Buffer
	// Create the URL for the call to the Agility API
	url.WriteString(configuration.APIURL)
	url.WriteString("current/cloud/")
	log.Println("URL:>", url.String())

	var payload bytes.Buffer
	//Create the payload for the request body
	s := ResourceData.Get("cloudname").(string)
	t := ResourceData.Get("description").(string)
	u := ResourceData.Get("cloudtype").(string)
	v := ResourceData.Get("hostname").(string)
	w := ResourceData.Get("publickey").(string)
	x := ResourceData.Get("privatekey").(string)
	y := ResourceData.Get("awsaccountnumber").(string)
	payload.WriteString(`<?xml version="1.0" encoding="UTF-8" standalone="yes"?><ns1:Cloud xmlns:ns1="http://servicemesh.com/agility/api"><ns1:name>`)
	payload.WriteString(s)
	payload.WriteString(`</ns1:name><ns1:description>`)
	payload.WriteString(t)
	payload.WriteString(`</ns1:description><ns1:cloudType><ns1:name>`)
	payload.WriteString(u)
	payload.WriteString(`</ns1:name><ns1:type>application/com.servicemesh.agility.api.CloudType+xml</ns1:type></ns1:cloudType><ns1:hostname>`)
	payload.WriteString(v)
	payload.WriteString(`</ns1:hostname><ns1:enabled>true</ns1:enabled><ns1:cloudId>`)
	payload.WriteString(y)
	payload.WriteString(`</ns1:cloudId><ns1:cloudCredentials><ns1:assetType><ns1:name>credential</ns1:name><ns1:type>application/com.servicemesh.agility.api.AssetType+xml</ns1:type></ns1:assetType><ns1:credentialType>SSH</ns1:credentialType><ns1:publicKey>`)
	payload.WriteString(w)
	payload.WriteString(`</ns1:publicKey><ns1:privateKey>`)
	payload.WriteString(x)
	payload.WriteString(`</ns1:privateKey><ns1:encrypted>true</ns1:encrypted></ns1:cloudCredentials><ns1:priceEngine><ns1:top>false</ns1:top><ns1:removable>true</ns1:removable></ns1:priceEngine></ns1:Cloud>`)
	payload1 := payload.String()
	log.Println("Payload=====>",payload1)
	payload2 := []byte(payload1)
	// Set the right HTTP Verb, and setup HTTP Basic Security
	req, err := http.NewRequest("POST", url.String(), bytes.NewBuffer([]byte(payload2)))
	req.Header.Set("Content-Type", "application/xml; charset=utf-8")
	req.SetBasicAuth(username,password)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// make the HTTPS request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// log the response details for debugging
	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)

	//Stream the response body into a byte array and return it
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))
	return body,nil
}

func SyncCloudProvider(ResourceData *schema.ResourceData, cloudId string, username string, password string) ([]byte, error){
	f, errf := os.OpenFile("./agility/api/agility.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if errf != nil {
		log.Println("error opening file: ", errf)
	}
	defer f.Close()

	log.SetOutput(f)

	var url bytes.Buffer
	// Create the URL for the call to the Agility API
	url.WriteString(configuration.APIURL)
	url.WriteString("current/cloud/")
	url.WriteString(cloudId)
	url.WriteString("/resync")
	log.Println("URL:>", url.String())

	// Set the right HTTP Verb, and setup HTTP Basic Security
	req, err := http.NewRequest("POST", url.String(), nil)
	req.Header.Set("Content-Type", "application/xml; charset=utf-8")
	req.SetBasicAuth(username,password)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// make the HTTPS request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// log the response details for debugging
	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)

	//Stream the response body into a byte array and return it
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))
	return body,nil
}

func LicenseUpload (username string, password string) []byte{

	var url bytes.Buffer

	// url for licence upload
	url.WriteString(configuration.APIURL)
	url.WriteString("license")
	//reading payload from AgilityLicense file
	file, err1 := ioutil.ReadFile("./agility/api/AgilityLicense.xml")
	if err1 != nil {
		log.Println("error:", err1)
	}

	//Payload code ends

	log.Println("URL:>",url.String())
	req, err := http.NewRequest("POST", url.String(),bytes.NewBuffer([]byte(file)))
	req.Header.Set("Content-Type", "text/plain; charset=utf-8")
	req.SetBasicAuth(username, password)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// make the HTTPS request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// log the response details for debugging
	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)

	//Stream the response body into a byte array and return it
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", resp.Body)
	return body
}

func GetImageId(imageName string, username string, password string) (string, error) {
	log.Println("imageName is: ", imageName)
	var url bytes.Buffer
	q := new(Result)

	// Create the URL for the call to the Agility API
	url.WriteString(configuration.APIURL)
	url.WriteString("current/image")
	log.Println("URL:>", url.String())

	// Set the right HTTP Verb, and setup HTTP Basic Security
	req, err := http.NewRequest("GET", url.String(), nil)
	req.SetBasicAuth(username, password)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// make the HTTPS request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// log the response details for debugging
	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)

	//Stream the response body into a byte array
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))

	//Parse the XML
	r := strings.NewReader(string(body))
	decoder := xml.NewDecoder(r)
	finish := false
	for {
		// Read tokens from the XML document in a stream.
		t, _ := decoder.Token()
		if t == nil {
			return "", errors.New("there are no images with this name")
		}
		if finish {
			break
		}
		// look for <link> element
		switch Element := t.(type) {
		case xml.StartElement:
			if Element.Name.Local == "link" {
				log.Println("Element name is : ", Element.Name.Local)

				// unmarshal the element into generic structure
				err := decoder.DecodeElement(&q, &Element)
				if err != nil {
					log.Println(err)
				}

				// if the project name matches the project defined to Terraform
				// then we are are the right place, so stop looking
				log.Println("Element value is :", string(q.Name))
				if string(q.Name) == imageName {
					log.Println("Found the image : ", q.Name)
					finish = true
					break
				}
			}
			// if the element is the <Linklist> then go again
			if Element.Name.Local == "Linklist" {
				log.Println("Element name is : ", Element.Name.Local)
			} else {
				log.Println("Unknown Element name is : ", Element.Name.Local)
			}
		default:
		}

	}

	// return the ID for the image
	return string(q.Id), nil
}

/*func CreateStack(ResourceData *schema.ResourceData, projectId string, imageId string, username string, password string) ([]byte, error){
	f, errf := os.OpenFile("./agility/api/agility.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if errf != nil {
		log.Println("error opening file: ", errf)
	}
	defer f.Close()

	log.SetOutput(f)

	var url bytes.Buffer
	// Create the URL for the call to the Agility API
	url.WriteString(configuration.APIURL)
	url.WriteString("current/stack/")
	log.Println("URL:>", url.String())

	var payload bytes.Buffer
	//Create the payload for the request body
	s := ResourceData.Get("projectname").(string)
	pid, err := GetProjectId(s, username, password)
	t := ResourceData.Get("imagename").(string)
	iid, err := GetImageId(t, username, password)
	u := ResourceData.Get("operatingsystem").(string)
	v := ResourceData.Get("stackname").(string)
	w := ResourceData.Get("stackdescription").(string)
	payload.WriteString(`<?xml version="1.0" encoding="UTF-8" standalone="yes"?><ns1:Stack xmlns:ns1="http://servicemesh.com/agility/api"><ns1:name>`)
	payload.WriteString(v)
	payload.WriteString(`</ns1:name><ns1:description>`)
	payload.WriteString(w)
	payload.WriteString(`</ns1:description><ns1:parent><ns1:id>`)
	payload.WriteString(pid)
	payload.WriteString(`</ns1:id></ns1:parent><ns1:operatingSystem>`)
	payload.WriteString(u)
	payload.WriteString(`</ns1:operatingSystem><ns1:images><ns1:id>`)
	payload.WriteString(iid)
	payload.WriteString(`</ns1:id></ns1:images></ns1:Stack>`)
	payload1 := payload.String()
	log.Println("Payload=====>",payload1)
	payload2 := []byte(payload1)
	// Set the right HTTP Verb, and setup HTTP Basic Security
	req, err := http.NewRequest("POST", url.String(), bytes.NewBuffer([]byte(payload2)))
	req.Header.Set("Content-Type", "application/xml; charset=utf-8")
	req.SetBasicAuth(username,password)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// make the HTTPS request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// log the response details for debugging
	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)

	//Stream the response body into a byte array and return it
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))
	return body,nil
}
*/
func CreateScript(ResourceData *schema.ResourceData, projectId string,username string, password string)([]byte, error){
	f, errf := os.OpenFile("./agility/api/agility.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if errf != nil {
		log.Println("error opening file: ", errf)
	}
	defer f.Close()

	log.SetOutput(f)

	var url bytes.Buffer
	// Create the URL for the call to the Agility API
	url.WriteString(configuration.APIURL)
	url.WriteString("current/project/")
	url.WriteString(projectId)
	url.WriteString("/script")
	log.Println("URL:>", url.String())

	//payload
	var payload bytes.Buffer
	scriptname := ResourceData.Get("scriptname").(string)
	desc :=ResourceData.Get("desc").(string)
	operatingsystem :=ResourceData.Get("operatingsystem").(string)
	language := ResourceData.Get("language").(string)
	rebootrequired := ResourceData.Get("rebootrequired").(string)
	content:=ResourceData.Get("content").(string)
	var filelocation bytes.Buffer
	filelocation.WriteString( "./agility/scripts/")
	filelocation.WriteString(content)
	filelocation1 :=filelocation.String()
	log.Println("file location",filelocation1)
	payload.WriteString(`<ns1:Script xmlns:ns1="http://servicemesh.com/agility/api"><ns1:name>`)
	payload.WriteString(scriptname)
	payload.WriteString(`</ns1:name><ns1:description>`)
	payload.WriteString(desc)
	payload.WriteString(`</ns1:description><ns1:operatingSystem>`)
	payload.WriteString(operatingsystem)
	payload.WriteString(`</ns1:operatingSystem><ns1:enableExtensions>true</ns1:enableExtensions><ns1:body>`)
	// payload.WriteString(content)
	var file []byte
	file, err1 := ioutil.ReadFile(filelocation1)
	if err1 != nil {
		log.Println("error:", err1)
	}
	file1 := string(file)
	payload.WriteString(file1)

	payload.WriteString(`</ns1:body><ns1:type>Guest</ns1:type><ns1:language><ns1:name>`)
	payload.WriteString(language)
	payload.WriteString(`</ns1:name></ns1:language><ns1:runAsAdmin>true</ns1:runAsAdmin><ns1:timeout>3600</ns1:timeout><ns1:retries>0</ns1:retries><ns1:delay>60</ns1:delay><ns1:errorAction>Continue</ns1:errorAction><ns1:rebootRequired>`)
	payload.WriteString(rebootrequired)
	payload.WriteString(`</ns1:rebootRequired></ns1:Script>`)
	payload1 := payload.String()
	log.Println("Payload1 ===== " , payload1)
	req, err := http.NewRequest("POST",url.String(),bytes.NewBuffer([]byte(payload1)))
	req.Header.Set("Content-Type", "application/xml; charset=utf-8")
	req.SetBasicAuth(username,password)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// make the HTTPS request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// log the response details for debugging
	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)

	//Stream the response body into a byte array and return it
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))
	return body,nil
}

func GetScripttId(ResourceData *schema.ResourceData,scriptName string,projecId string, username string, password string) (string, error) {

	var url bytes.Buffer
	q := new(Result)
	log.Println("inside getscriptid ")
	// Create the URL for the call to the Agility API
	url.WriteString(configuration.APIURL)
	url.WriteString("current/project/")
	pid := projecId
	log.Println("pid === ", pid)
	url.WriteString(pid)
	url.WriteString("/script")

	log.Println("URL:>", url.String())

	// Set the right HTTP Verb, and setup HTTP Basic Security
	req, err := http.NewRequest("GET", url.String(), nil)
	req.SetBasicAuth(username, password)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// make the HTTPS request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// log the response details for debugging
	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)

	//Stream the response body into a byte array
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))

	//Parse the XML
	r := strings.NewReader(string(body))
	decoder := xml.NewDecoder(r)
	finish := false
	for {
		// Read tokens from the XML document in a stream.
		t, _ := decoder.Token()
		if t == nil {
			return "", errors.New("there are no scripts with this name")
		}
		if finish {
			break
		}
		// look for <link> element
		switch Element := t.(type) {
		case xml.StartElement:
			if Element.Name.Local == "link" {
				log.Println("Element name is : ", Element.Name.Local)

				// unmarshal the element into generic structure
				err := decoder.DecodeElement(&q, &Element)
				if err != nil {
					log.Println(err)
				}

				// if the script name matches the script defined to Terraform
				// then we are are the right place, so stop looking
				log.Println("Element value is :", string(q.Name))
				if string(q.Name) == scriptName {
					log.Println("Found the script : ", q.Name)
					finish = true
					break
				}
			}
			// if the element is the <Linklist> then go again
			if Element.Name.Local == "Linklist" {
				log.Println("Element name is : ", Element.Name.Local)
			} else {
				log.Println("Unknown Element name is : ", Element.Name.Local)
			}
		default:
		}

	}
	log.Println("script id ", q.Id)
	// return the ID for the script
	return string(q.Id), nil
}

func UploadAttachment(ResourceData *schema.ResourceData,scriptid string,username string, password string)([]byte){
	f, errf := os.OpenFile("./agility/api/agility.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if errf != nil {
		log.Println("error opening file: ", errf)
	}
	defer f.Close()

	log.SetOutput(f)
	var url bytes.Buffer
	// Create the URL for the call to the Agility API
	url.WriteString(configuration.APIURL)
	url.WriteString("current/script/")
	url.WriteString(scriptid)
	url.WriteString("/attachment")
	log.Println("URL:>", url.String())
	url1 := url.String()
	log.Println("URL=====>",url1)

	//String for attachment location
	attachmentname:=ResourceData.Get("attachmentname").(string)
	var attachmentlocation bytes.Buffer
	attachmentlocation.WriteString(`./agility/attachments/`)
	attachmentlocation.WriteString(attachmentname)
	attachmentlocation1 := attachmentlocation.String()
	log.Println("Attachment location=====>",attachmentlocation1)

	var payload bytes.Buffer
	//Create the payload for the request body
	payload.WriteString(`<?xml version="1.0" encoding="UTF-8" standalone="yes"?><Attachment xmlns="http://servicemesh.com/agility/api"><name>`)
	payload.WriteString(attachmentname)
	payload.WriteString(`</name></Attachment>`)
	payload1:=payload.String()
	log.Println("Payload=====>",payload1)

	//file, err1 := ioutil.WriteFile("./agility/api/attachment.xml",)
	log.Println("<=====Executing curl command=====>")
	methodCall:= "POST"
	HArgument:= "Content-Type:multipart/mixed"
	attachmentPayload:=`"file=@/usr/local/Cellar/go/1.9/libexec/src/github.com/csc/csc-agility-terraform-provider-plug-in/agility/attachments/attachment.xml;type=application/xml"`
	attachment:=`"file=@/usr/local/Cellar/go/1.9/libexec/src/github.com/csc/csc-agility-terraform-provider-plug-in/agility/attachments/a;type=application/x-compressed"`
	resp := exec.Command("curl", "-u",username+":"+password, "-X",methodCall,"-H",HArgument, "-F",attachmentPayload , "-F",attachment , "-k", url1)
	resp.Stdout=os.Stdout
	resp.Stderr=os.Stderr

	resp.Run()
	log.Println("resp stdout",resp.Stdout)
	log.Println("resp std err", resp.Stderr)
	log.Println("Response=====>",resp)
	//curl -upiyush:07June1990@ -X POST -H "Content-Type:multipart/mixed" -F "file=@/Users/pmohta/Documents/attachment.xml;type=application/xml" -F "file=@/Users/pmohta/Documents/agility-monitor-windows.zip;type=application/x-compressed" -k https://52.31.77.163:8443/agility/api/current/script/1861/attachment
	//log.Println("Payload1 ===== " , payload1)
	//req, err := http.NewRequest("POST",url.String(),bytes.NewBuffer([]byte(payload1)))
	//req.Header.Set("Content-Type", "multipart/mixed; charset=utf-8")
	//req.SetBasicAuth(username,password)
	//tr := &http.Transport{
	//	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	//}
	//client := &http.Client{Transport: tr}

	// make the HTTPS request
	//resp, err := client.Do(req)
	//if err != nil {
	//	panic(err)
	//}
	//defer resp.Body.Close()

	// log the response details for debugging
	//log.Println("response Status:", resp.Status)
	//log.Println("response Headers:", resp.Header)

	//Stream the response body into a byte array and return it
	//body, _ := ioutil.ReadAll(resp.Body)
	//log.Println("response Body:", string(body)
	//return body,nil
	return nil
}

func CreatePackage(ResourceData *schema.ResourceData, projectId string, username string, password string)([]byte, error){
	f, errf := os.OpenFile("./agility/api/agility.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if errf != nil {
		log.Println("error opening file: ", errf)
	}
	defer f.Close()

	log.SetOutput(f)

	var url bytes.Buffer
	// Create the URL for the call to the Agility API
	url.WriteString(configuration.APIURL)
	url.WriteString("current/project/")
	url.WriteString(projectId)
	url.WriteString("/package")
	log.Println("URL:>", url.String())

	//payload
	var payload bytes.Buffer
	packagename := ResourceData.Get("packagename").(string)
	packagedescription :=ResourceData.Get("packagedescription").(string)
	operatingsystem :=ResourceData.Get("operatingsystem").(string)
	installscriptname1 := ResourceData.Get("installscriptname1").(string)
	installscriptid1,err := GetScripttId(ResourceData,installscriptname1,projectId,username,password)
	installscriptname2 := ResourceData.Get("installscriptname2").(string)
	installscriptid2,err := GetScripttId(ResourceData,installscriptname2,projectId,username,password)
	startupscriptname1 := ResourceData.Get("startupscriptname1").(string)
	startupscriptid1,err := GetScripttId(ResourceData,startupscriptname1,projectId,username,password)
	startupscriptname2 := ResourceData.Get("startupscriptname2").(string)
	startupscriptid2,err := GetScripttId(ResourceData,startupscriptname2,projectId,username,password)
	operationalscriptname1 := ResourceData.Get("operationalscriptname1").(string)
	operationalscriptid1,err := GetScripttId(ResourceData,operationalscriptname1,projectId,username,password)
	operationalscriptname2 := ResourceData.Get("operationalscriptname2").(string)
	operationalscriptid2,err := GetScripttId(ResourceData,operationalscriptname2,projectId,username,password)
	operationalscriptname3 := ResourceData.Get("operationalscriptname3").(string)
	operationalscriptid3,err := GetScripttId(ResourceData,operationalscriptname3,projectId,username,password)
	operationalscriptname4 := ResourceData.Get("operationalscriptname4").(string)
	operationalscriptid4,err := GetScripttId(ResourceData,operationalscriptname4,projectId,username,password)
	log.Println("Install script name 1=====>",installscriptname1)
	log.Println("Install script id 1=====>",installscriptid1)
	log.Println("Install script name 2=====>",installscriptname2)
	log.Println("Install script id 2=====>",installscriptid2)
	log.Println("Startup script name 1=====>",startupscriptname1)
	log.Println("Startup script id 1=====>",startupscriptid1)
	log.Println("Startup script name 2=====>",startupscriptname2)
	log.Println("Startup script id 2=====>",startupscriptid2)
	log.Println("Operational script name 1=====>",operationalscriptname1)
	log.Println("Operational script id 1=====>",operationalscriptid1)
	log.Println("Operational script name 2=====>",operationalscriptname2)
	log.Println("Operational script id 2=====>",operationalscriptid2)
	log.Println("Operational script name 3=====>",operationalscriptname3)
	log.Println("Operational script id 3=====>",operationalscriptid3)
	log.Println("Operational script name 4=====>",operationalscriptname4)
	log.Println("Operational script id 4=====>",operationalscriptid4)

	//Create the payload

	payload.WriteString(`<?xml version="1.0" encoding="UTF-8" standalone="yes"?><ns1:Package xmlns:ns1="http://servicemesh.com/agility/api"><ns1:name>`)
	payload.WriteString(packagename)
	payload.WriteString(`</ns1:name><ns1:description>`)
	payload.WriteString(packagedescription)
	payload.WriteString(`</ns1:description><ns1:operatingSystem>`)
	payload.WriteString(operatingsystem)
	payload.WriteString(`</ns1:operatingSystem><ns1:install><ns1:id>`)
	payload.WriteString(installscriptid1)
	payload.WriteString(`</ns1:id></ns1:install><ns1:install><ns1:id>`)
	payload.WriteString(installscriptid2)
	payload.WriteString(`</ns1:id></ns1:install><ns1:startup><ns1:id>`)
	payload.WriteString(startupscriptid1)
	payload.WriteString(`</ns1:id></ns1:startup><ns1:startup><ns1:id>`)
	payload.WriteString(startupscriptid2)
	payload.WriteString(`</ns1:id></ns1:startup><ns1:operational><ns1:id>`)
	payload.WriteString(operationalscriptid1)
	payload.WriteString(`</ns1:id></ns1:operational><ns1:operational><ns1:id>`)
	payload.WriteString(operationalscriptid2)
	payload.WriteString(`</ns1:id></ns1:operational><ns1:operational><ns1:id>`)
	payload.WriteString(operationalscriptid3)
	payload.WriteString(`</ns1:id></ns1:operational><ns1:operational><ns1:id>`)
	payload.WriteString(operationalscriptid4)
	payload.WriteString(`</ns1:id></ns1:operational></ns1:Package>`)
	payload1 := payload.String()
	log.Println("Payload1 =====>" , payload1)
	req, err := http.NewRequest("POST",url.String(),bytes.NewBuffer([]byte(payload1)))
	req.Header.Set("Content-Type", "application/xml; charset=utf-8")
	req.SetBasicAuth(username,password)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// make the HTTPS request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// log the response details for debugging
	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)

	//Stream the response body into a byte array and return it
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))
	return body,nil
}

func CreateFirewallPolicy(ResourceData *schema.ResourceData, projectId string, username string, password string)([]byte, error){

	f, errf := os.OpenFile("./agility/api/agility.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if errf != nil {
		log.Println("error opening file: ", errf)
	}
	defer f.Close()

	log.SetOutput(f)

	var url bytes.Buffer
	// Create the URL for the call to the Agility API
	url.WriteString(configuration.APIURL)
	url.WriteString("current/policy")

	log.Println("URL:>", url.String())

	//payload
	var payload bytes.Buffer

	firewallname := ResourceData.Get("firewallname").(string)
	firewalldesc:= ResourceData.Get("firewalldesc").(string)
	direction:= ResourceData.Get("direction").(string)
	protocolprefix:= ResourceData.Get("protocolprefix").(string)
	protocolallowed:= ResourceData.Get("protocolallowed").(string)
	firewalltype:= ResourceData.Get("firewalltype").(string)

	protocolname1:= ResourceData.Get("protocolname1").(string)
	protocoldesc1:= ResourceData.Get("protocoldesc1").(string)
	protocolminport1:= ResourceData.Get("protocolminport1").(string)
	protocolmaxport1:= ResourceData.Get("protocolmaxport1").(string)
	protocol1:= ResourceData.Get("protocol1").(string)

	protocolname2:= ResourceData.Get("protocolname2").(string)
	protocoldesc2:= ResourceData.Get("protocoldesc2").(string)
	protocolminport2:= ResourceData.Get("protocolminport2").(string)
	protocolmaxport2:= ResourceData.Get("protocolmaxport2").(string)
	protocol2:= ResourceData.Get("protocol2").(string)

	protocolname3:= ResourceData.Get("protocolname3").(string)
	protocoldesc3:= ResourceData.Get("protocoldesc3").(string)
	protocolminport3:= ResourceData.Get("protocolminport3").(string)
	protocolmaxport3:= ResourceData.Get("protocolmaxport3").(string)
	protocol3:= ResourceData.Get("protocol3").(string)

	protocolname4:= ResourceData.Get("protocolname4").(string)
	protocoldesc4:= ResourceData.Get("protocoldesc4").(string)
	protocolminport4:= ResourceData.Get("protocolminport4").(string)
	protocolmaxport4:= ResourceData.Get("protocolmaxport4").(string)
	protocol4:= ResourceData.Get("protocol4").(string)

	payload.WriteString(`<?xml version="1.0" encoding="UTF-8" standalone="yes"?><ns1:Policy xmlns:ns1="http://servicemesh.com/agility/api"><ns1:name>`)
	payload.WriteString(firewallname)
	payload.WriteString(`</ns1:name><ns1:description>`)
	payload.WriteString(firewalldesc)
	payload.WriteString(`</ns1:description><ns1:parent><ns1:id>`)
	payload.WriteString(projectId)
	payload.WriteString(`</ns1:id></ns1:parent><ns1:definition>&lt;AccessList&gt;&lt;direction&gt;`)
	payload.WriteString(direction)
	payload.WriteString(`&lt;/direction&gt;&lt;protocols&gt;&lt;name&gt;`)
	payload.WriteString(protocolname1)
	payload.WriteString(`&lt;/name&gt;&lt;description&gt;`)
	payload.WriteString(protocoldesc1)
	payload.WriteString(`&lt;/description&gt;&lt;minPort&gt;`)
	payload.WriteString(protocolminport1)
	payload.WriteString(`&lt;/minPort&gt;&lt;maxPort&gt;`)
	payload.WriteString(protocolmaxport1)
	payload.WriteString(`&lt;/maxPort&gt;&lt;protocol&gt;`)
	payload.WriteString(protocol1)
	payload.WriteString(`&lt;/protocol&gt;&lt;prefixes&gt;`)
	payload.WriteString(protocolprefix)
	payload.WriteString(`&lt;/prefixes&gt;&lt;allowed&gt;`)
	payload.WriteString(protocolallowed)
	payload.WriteString(`&lt;/allowed&gt;&lt;/protocols&gt;&lt;protocols&gt;&lt;name&gt;`)
	payload.WriteString(protocolname2)
	payload.WriteString(`&lt;/name&gt;&lt;description&gt;`)
	payload.WriteString(protocoldesc2)
	payload.WriteString(`&lt;/description&gt;&lt;minPort&gt;`)
	payload.WriteString(protocolminport2)
	payload.WriteString(`&lt;/minPort&gt;&lt;maxPort&gt;`)
	payload.WriteString(protocolmaxport2)
	payload.WriteString(`&lt;/maxPort&gt;&lt;protocol&gt;`)
	payload.WriteString(protocol2)
	payload.WriteString(`&lt;/protocol&gt;&lt;prefixes&gt;`)
	payload.WriteString(protocolprefix)
	payload.WriteString(`&lt;/prefixes&gt;&lt;allowed&gt;`)
	payload.WriteString(protocolallowed)
	payload.WriteString(`&lt;/allowed&gt;&lt;/protocols&gt;&lt;protocols&gt;&lt;name&gt;`)
	payload.WriteString(protocolname3)
	payload.WriteString(`&lt;/name&gt;&lt;description&gt;`)
	payload.WriteString(protocoldesc3)
	payload.WriteString(`&lt;/description&gt;&lt;minPort&gt;`)
	payload.WriteString(protocolminport3)
	payload.WriteString(`&lt;/minPort&gt;&lt;maxPort&gt;`)
	payload.WriteString(protocolmaxport3)
	payload.WriteString(`&lt;/maxPort&gt;&lt;protocol&gt;`)
	payload.WriteString(protocol3)
	payload.WriteString(`&lt;/protocol&gt;&lt;prefixes&gt;`)
	payload.WriteString(protocolprefix)
	payload.WriteString(`&lt;/prefixes&gt;&lt;allowed&gt;`)
	payload.WriteString(protocolallowed)
	payload.WriteString(`&lt;/allowed&gt;&lt;/protocols&gt;&lt;protocols&gt;&lt;name&gt;`)
	payload.WriteString(protocolname4)
	payload.WriteString(`&lt;/name&gt;&lt;description&gt;`)
	payload.WriteString(protocoldesc4)
	payload.WriteString(`&lt;/description&gt;&lt;minPort&gt;`)
	payload.WriteString(protocolminport4)
	payload.WriteString(`&lt;/minPort&gt;&lt;maxPort&gt;`)
	payload.WriteString(protocolmaxport4)
	payload.WriteString(`&lt;/maxPort&gt;&lt;protocol&gt;`)
	payload.WriteString(protocol4)
	payload.WriteString(`&lt;/protocol&gt;&lt;prefixes&gt;`)
	payload.WriteString(protocolprefix)
	payload.WriteString(`&lt;/prefixes&gt;&lt;allowed&gt;`)
	payload.WriteString(protocolallowed)
	payload.WriteString(`&lt;/allowed&gt;&lt;/protocols&gt;&lt;/AccessList&gt;</ns1:definition><ns1:type><ns1:name>`)
	payload.WriteString(firewalltype)
	payload.WriteString(`</ns1:name></ns1:type></ns1:Policy>`)

	payload1 := payload.String()
	log.Println("payload====>",payload1)

	req, err := http.NewRequest("POST",url.String(),bytes.NewBuffer([]byte(payload1)))
	req.Header.Set("Content-Type", "application/xml; charset=utf-8")
	req.SetBasicAuth(username,password)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// make the HTTPS request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// log the response details for debugging
	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)

	//Stream the response body into a byte array and return it
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))
	return body,nil
}

func CheckIn(ResourceData *schema.ResourceData, projectId string, username string, password string)([]byte, error){

	f, errf := os.OpenFile("./agility/api/agility.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if errf != nil {
		log.Println("error opening file: ", errf)
	}
	defer f.Close()

	log.SetOutput(f)
	assetname := ResourceData.Get("assetname").(string)
	asset:=ResourceData.Get("asset").(string)
	var aid string
	if asset=="script"{
		assetid,err := GetScripttId(ResourceData,assetname,projectId,username,password)
		if err != nil {
			log.Println("no script found with the name: ", assetname)
		}
		aid=assetid
	} else if asset == "package" {
		assetid,err:=GetPackageId(ResourceData,assetname,projectId,username,password)
		if err != nil {
			log.Println("no package found with the name: ", assetname)
		}
		aid=assetid
	} else if asset == "stack"{
		assetid,err:=GetStackId(ResourceData,assetname,username,password)
		if err != nil {
			log.Println("no stack found with the name: ", assetname)
		}
		aid=assetid
	} else if asset == "blueprint"{
		assetid,err:=GetBlueprintId(assetname,projectId,username,password)
		if err != nil {
			log.Println("no blueprint found with the name: ", assetname)
		}
		aid=assetid
	} else if asset == "policy"{
		assetid,err:=GetPolicyId(ResourceData,assetname,projectId,username,password)
		if err != nil {
			log.Println("no policy found with the name: ", assetname)
		}
		aid=assetid
	} else{
		log.Println("please enter a valid asset")
	}
	containername:=ResourceData.Get("containername").(string)
	containerid,err1:=GetContainerId(containername,username,password)
	if err1 != nil {
		log.Println("no container found with the name: ", containername)
	}
	headversionallowed:=ResourceData.Get("headversionallowed").(string)

	var url bytes.Buffer
	// Create the URL for the call to the Agility API
	url.WriteString(configuration.APIURL)
	url.WriteString("current/")
	url.WriteString(asset)
	url.WriteString("/")
	url.WriteString(aid)
	url.WriteString("/checkin")

	log.Println("URL:>", url.String())

	//payload
	var payload bytes.Buffer

	payload.WriteString(`<ns1:Container xmlns:ns1="http://servicemesh.com/agility/api"><ns1:name>`)
	payload.WriteString(containername)
	payload.WriteString(`</ns1:name><ns1:id>`)
	payload.WriteString(containerid)
	payload.WriteString(`</ns1:id><ns1:headAllowed>`)
	payload.WriteString(headversionallowed)
	payload.WriteString(`</ns1:headAllowed></ns1:Container>`)

	payload1 := payload.String()
	log.Println("payload====>",payload1)

	req, err := http.NewRequest("POST",url.String(),bytes.NewBuffer([]byte(payload1)))
	req.Header.Set("Content-Type", "application/xml; charset=utf-8")
	req.SetBasicAuth(username,password)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// make the HTTPS request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// log the response details for debugging
	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)

	//Stream the response body into a byte array and return it
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))
	return body,nil

}

func Approve(ResourceData *schema.ResourceData, projectId string, username string, password string)([]byte, error){

	f, errf := os.OpenFile("./agility/api/agility.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if errf != nil {
		log.Println("error opening file: ", errf)
	}
	defer f.Close()

	log.SetOutput(f)

	asset:=ResourceData.Get("asset").(string)
	assetname:=ResourceData.Get("assetname").(string)
	var aid string
	if asset=="script"{
		assetid,err := GetScripttId(ResourceData,assetname,projectId,username,password)
		if err != nil {
			log.Println("no script found with the name: ", assetname)
		}
		aid=assetid
	} else if asset == "package" {
		assetid,err:=GetPackageId(ResourceData,assetname,projectId,username,password)
		if err != nil {
			log.Println("no package found with the name: ", assetname)
		}
		aid=assetid
	} else if asset == "stack"{
		assetid,err:=GetStackId(ResourceData,assetname,username,password)
		if err != nil {
			log.Println("no stack found with the name: ", assetname)
		}
		aid=assetid
	} else if asset == "blueprint"{
		assetid,err:=GetBlueprintId(assetname,projectId,username,password)
		if err != nil {
			log.Println("no blueprint found with the name: ", assetname)
		}
		aid=assetid
	} else if asset == "policy" {
		log.Println(" the asset is a policy , no need for approval")
	} else
	{
		log.Println("please enter a valid asset")
	}
	state:=ResourceData.Get("state").(string)
	comment:=ResourceData.Get("comment").(string)

	log.Println("asset id ", aid)
	var url bytes.Buffer
	// Create the URL for the call to the Agility API
	url.WriteString(configuration.APIURL)
	url.WriteString("current/")
	url.WriteString(asset)
	url.WriteString("/")
	url.WriteString(aid)
	url.WriteString("/")
	url.WriteString(state)
	log.Println("URL:>", url.String())

	//payload
	var payload bytes.Buffer
	payload.WriteString(`<WorkflowRequest xmlns="http://servicemesh.com/agility/api"><comment>`)
	payload.WriteString(comment)
	payload.WriteString(`</comment></WorkflowRequest>`)
	payload1 := payload.String()
	log.Println("payload====>",payload1)

	req, err := http.NewRequest("POST",url.String(),bytes.NewBuffer([]byte(payload1)))
	req.Header.Set("Content-Type", "application/xml; charset=utf-8")
	req.SetBasicAuth(username,password)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// make the HTTPS request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// log the response details for debugging
	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)

	//Stream the response body into a byte array and return it
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))
	return body,nil

}

func GetPackageId(ResourceData *schema.ResourceData,packageName string,projecId string, username string, password string) (string, error) {

	var url bytes.Buffer
	q := new(Result)

	// Create the URL for the call to the Agility API
	url.WriteString(configuration.APIURL)
	url.WriteString("current/project/")
	pid := projecId
	url.WriteString(pid)
	url.WriteString("/package")

	log.Println("URL:>", url.String())

	// Set the right HTTP Verb, and setup HTTP Basic Security
	req, err := http.NewRequest("GET", url.String(), nil)
	req.SetBasicAuth(username, password)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// make the HTTPS request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// log the response details for debugging
	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)

	//Stream the response body into a byte array
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))

	//Parse the XML
	r := strings.NewReader(string(body))
	decoder := xml.NewDecoder(r)
	finish := false
	for {
		// Read tokens from the XML document in a stream.
		t, _ := decoder.Token()
		if t == nil {
			return "", errors.New("there are no package with this name")
		}
		if finish {
			break
		}
		// look for <link> element
		switch Element := t.(type) {
		case xml.StartElement:
			if Element.Name.Local == "link" {
				log.Println("Element name is : ", Element.Name.Local)

				// unmarshal the element into generic structure
				err := decoder.DecodeElement(&q, &Element)
				if err != nil {
					log.Println(err)
				}

				// if the script name matches the script defined to Terraform
				// then we are are the right place, so stop looking
				log.Println("Element value is :", string(q.Name))
				if string(q.Name) == packageName {
					log.Println("Found the package : ", q.Name)
					finish = true
					break
				}
			}
			// if the element is the <Linklist> then go again
			if Element.Name.Local == "Linklist" {
				log.Println("Element name is : ", Element.Name.Local)
			} else {
				log.Println("Unknown Element name is : ", Element.Name.Local)
			}
		default:
		}

	}
	log.Println("package id ", q.Id)
	// return the ID for the script
	return string(q.Id), nil
}


func GetStackId(ResourceData *schema.ResourceData,stackName string, username string, password string) (string, error) {

	var url bytes.Buffer
	q := new(Result)

	// Create the URL for the call to the Agility API
	url.WriteString(configuration.APIURL)
	url.WriteString("current/stack")


	log.Println("URL:>", url.String())

	// Set the right HTTP Verb, and setup HTTP Basic Security
	req, err := http.NewRequest("GET", url.String(), nil)
	req.SetBasicAuth(username, password)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// make the HTTPS request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// log the response details for debugging
	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)

	//Stream the response body into a byte array
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))

	//Parse the XML
	r := strings.NewReader(string(body))
	decoder := xml.NewDecoder(r)
	finish := false
	for {
		// Read tokens from the XML document in a stream.
		t, _ := decoder.Token()
		if t == nil {
			return "", errors.New("there are no stack with this name")
		}
		if finish {
			break
		}
		// look for <link> element
		switch Element := t.(type) {
		case xml.StartElement:
			if Element.Name.Local == "link" {
				log.Println("Element name is : ", Element.Name.Local)

				// unmarshal the element into generic structure
				err := decoder.DecodeElement(&q, &Element)
				if err != nil {
					log.Println(err)
				}

				// if the script name matches the script defined to Terraform
				// then we are are the right place, so stop looking
				log.Println("Element value is :", string(q.Name))
				if string(q.Name) == stackName {
					log.Println("Found the stack : ", q.Name)
					finish = true
					break
				}
			}
			// if the element is the <Linklist> then go again
			if Element.Name.Local == "Linklist" {
				log.Println("Element name is : ", Element.Name.Local)
			} else {
				log.Println("Unknown Element name is : ", Element.Name.Local)
			}
		default:
		}

	}
	log.Println("stack id ", q.Id)
	// return the ID for the script
	return string(q.Id), nil
}

func GetPolicyId(ResourceData *schema.ResourceData,policyName string,projecId string, username string, password string) (string, error) {

	var url bytes.Buffer
	q := new(Result)

	// Create the URL for the call to the Agility API
	url.WriteString(configuration.APIURL)
	url.WriteString("current/project/")
	pid := projecId
	url.WriteString(pid)
	url.WriteString("/policy")

	log.Println("URL:>", url.String())

	// Set the right HTTP Verb, and setup HTTP Basic Security
	req, err := http.NewRequest("GET", url.String(), nil)
	req.SetBasicAuth(username, password)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// make the HTTPS request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// log the response details for debugging
	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)

	//Stream the response body into a byte array
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))

	//Parse the XML
	r := strings.NewReader(string(body))
	decoder := xml.NewDecoder(r)
	finish := false
	for {
		// Read tokens from the XML document in a stream.
		t, _ := decoder.Token()
		if t == nil {
			return "", errors.New("there are no policy with this name")
		}
		if finish {
			break
		}
		// look for <link> element
		switch Element := t.(type) {
		case xml.StartElement:
			if Element.Name.Local == "link" {
				log.Println("Element name is : ", Element.Name.Local)

				// unmarshal the element into generic structure
				err := decoder.DecodeElement(&q, &Element)
				if err != nil {
					log.Println(err)
				}

				// if the script name matches the script defined to Terraform
				// then we are are the right place, so stop looking
				log.Println("Element value is :", string(q.Name))
				if string(q.Name) == policyName {
					log.Println("Found the policy : ", q.Name)
					finish = true
					break
				}
			}
			// if the element is the <Linklist> then go again
			if Element.Name.Local == "Linklist" {
				log.Println("Element name is : ", Element.Name.Local)
			} else {
				log.Println("Unknown Element name is : ", Element.Name.Local)
			}
		default:
		}

	}
	log.Println("policy id ", q.Id)
	// return the ID for the script
	return string(q.Id), nil
}

func CreateBlueprint(ResourceData *schema.ResourceData,projectId string, username string, password string) ([]byte, error) {

	f, errf := os.OpenFile("./agility/api/agility.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if errf != nil {
		log.Println("error opening file: ", errf)
	}
	defer f.Close()

	log.SetOutput(f)

	blueprintname:=ResourceData.Get("blueprintname").(string)
	blueprintdesc:=ResourceData.Get("blueprintdesc").(string)
	stackname:=ResourceData.Get("stackname").(string)
	packagename:=ResourceData.Get("packagename").(string)
	policyname:=ResourceData.Get("policyname").(string)
	headversionallowed:=ResourceData.Get("headversionallowed").(string)
	workloadname:=ResourceData.Get("workloadname").(string)
	policyassignmentname:=ResourceData.Get("policyassignmentname").(string)
	log.Println("project id ", projectId)
	log.Println("policy assignment name", policyassignmentname)

	policyassignmentid,err3:=GetPolicyAssignmentId(ResourceData,policyassignmentname,projectId,username,password)
	if err3 != nil {
		log.Println("error opening file: ", err3)
	}
	stackid,err:=GetStackId(ResourceData,stackname,username,password)
	if err != nil {
		log.Println("error opening file: ", err)
	}
	packageid,err1:=GetPackageId(ResourceData,packagename,projectId,username,password)
	if err1 != nil {
		log.Println("error opening file: ", err1)
	}
	policyid,err2:=GetPolicyId(ResourceData,policyname,projectId,username,password)
	if err2 != nil {
		log.Println("error opening file: ", err2)
	}
	log.Println("policy assignment id ", policyassignmentid)
	var url bytes.Buffer
	// Create the URL for the call to the Agility API
	url.WriteString(configuration.APIURL)
	url.WriteString("current/project/")
	url.WriteString(projectId)
	url.WriteString("/blueprint")
	log.Println("URL:>", url.String())

	//payload
	var payload bytes.Buffer
	payload.WriteString(`<?xml version="1.0" encoding="UTF-8" standalone="yes"?><ns1:Blueprint xmlns:ns1="http://servicemesh.com/agility/api"><ns1:name>`)
	payload.WriteString(blueprintname)
	payload.WriteString(`</ns1:name><ns1:description>`)
	payload.WriteString(blueprintdesc)
	payload.WriteString(`</ns1:description><ns1:policyAssignment><ns1:name>`)
	payload.WriteString(policyassignmentname)
	payload.WriteString(`</ns1:name><ns1:id>`)
	payload.WriteString(policyassignmentid)
	payload.WriteString(`</ns1:id><ns1:policy xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:type="ns1:VersionedItemLink"><ns1:id>`)
	payload.WriteString(policyid)
	payload.WriteString(`</ns1:id><ns1:type>application/com.servicemesh.agility.api.Policy+xml</ns1:type></ns1:policy></ns1:policyAssignment><ns1:headAllowed>`)
	payload.WriteString(headversionallowed)
	payload.WriteString(`</ns1:headAllowed><ns1:anyOrderItem xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:type="ns1:Workload"><ns1:name>`)
	payload.WriteString(workloadname)
	payload.WriteString(`</ns1:name><ns1:packages xsi:type="ns1:VersionedItemLink"><ns1:id>`)
	payload.WriteString(packageid)
	payload.WriteString(`</ns1:id></ns1:packages><ns1:baseStack xsi:type="ns1:VersionedItemLink"><ns1:id>`)
	payload.WriteString(stackid)
	payload.WriteString(`</ns1:id></ns1:baseStack></ns1:anyOrderItem></ns1:Blueprint>`)

	payload1 := payload.String()
	log.Println("payload====>",payload1)

	req, err := http.NewRequest("POST",url.String(),bytes.NewBuffer([]byte(payload1)))
	req.Header.Set("Content-Type", "application/xml; charset=utf-8")
	req.SetBasicAuth(username,password)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// make the HTTPS request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// log the response details for debugging
	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)

	//Stream the response body into a byte array and return it
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))
	return body,nil
}
func AssignPolicy(ResourceData *schema.ResourceData,projectId string, username string, password string) ([]byte, error) {

	f, errf := os.OpenFile("./agility/api/agility.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if errf != nil {
		log.Println("error opening file: ", errf)
	}
	defer f.Close()

	log.SetOutput(f)

	policyname:=ResourceData.Get("policyname").(string)
	policyid,err:=GetAllPolicy(ResourceData,policyname,username,password)
	if err != nil {
		log.Println("there is no policy found with this name : ", policyname)
	}
	log.Println("policy id ",policyid)

	var url bytes.Buffer
	// Create the URL for the call to the Agility API
	url.WriteString(configuration.APIURL)
	url.WriteString("current/project/")
	url.WriteString(projectId)
	url.WriteString("/policy/")
	url.WriteString(policyid)
	log.Println("URL:>", url.String())

	req, err := http.NewRequest("PUT", url.String(), nil)
	req.Header.Set("Content-Type", "application/xml; charset=utf-8")
	req.SetBasicAuth(username, password)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// make the HTTPS request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// log the response details for debugging
	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)

	//Stream the response body into a byte array and return it
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))
	return body,nil

}

func UnassignPolicy(ResourceData *schema.ResourceData,projectId string, username string, password string) ([]byte, error) {

	f, errf := os.OpenFile("./agility/api/agility.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if errf != nil {
		log.Println("error opening file: ", errf)
	}
	defer f.Close()

	log.SetOutput(f)

	policyname:=ResourceData.Get("policyname").(string)
	policyid,err:=GetPolicyId(ResourceData,policyname,projectId,username,password)
	if err != nil {
		log.Println("there is no policy found with this name : ", policyname)
	}

	var url bytes.Buffer
	// Create the URL for the call to the Agility API
	url.WriteString(configuration.APIURL)
	url.WriteString("current/project/")
	url.WriteString(projectId)
	url.WriteString("/policy/")
	url.WriteString(policyid)
	log.Println("URL:>", url.String())

	req, err := http.NewRequest("DELETE", url.String(), nil)
	req.Header.Set("Content-Type", "application/xml; charset=utf-8")
	req.SetBasicAuth(username, password)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// make the HTTPS request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// log the response details for debugging
	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)

	//Stream the response body into a byte array and return it
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))
	return body,nil

}

func GetAllPolicy(ResourceData *schema.ResourceData,policyName string, username string, password string) (string, error) {

	var url bytes.Buffer
	q := new(Result)

	// Create the URL for the call to the Agility API
	url.WriteString(configuration.APIURL)
	url.WriteString("current/")
	url.WriteString("policy")

	log.Println("URL:>", url.String())

	// Set the right HTTP Verb, and setup HTTP Basic Security
	req, err := http.NewRequest("GET", url.String(), nil)
	req.SetBasicAuth(username, password)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// make the HTTPS request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// log the response details for debugging
	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)

	//Stream the response body into a byte array
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))

	//Parse the XML
	r := strings.NewReader(string(body))
	decoder := xml.NewDecoder(r)
	finish := false
	for {
		// Read tokens from the XML document in a stream.
		t, _ := decoder.Token()
		if t == nil {
			return "", errors.New("there are no policy with this name")
		}
		if finish {
			break
		}
		// look for <link> element
		switch Element := t.(type) {
		case xml.StartElement:
			if Element.Name.Local == "link" {
				log.Println("Element name is : ", Element.Name.Local)

				// unmarshal the element into generic structure
				err := decoder.DecodeElement(&q, &Element)
				if err != nil {
					log.Println(err)
				}

				// if the script name matches the script defined to Terraform
				// then we are are the right place, so stop looking
				log.Println("Element value is :", string(q.Name))
				if string(q.Name) == policyName {
					log.Println("Found the policy : ", q.Name)
					finish = true
					break
				}
			}
			// if the element is the <Linklist> then go again
			if Element.Name.Local == "Linklist" {
				log.Println("Element name is : ", Element.Name.Local)
			} else {
				log.Println("Unknown Element name is : ", Element.Name.Local)
			}
		default:
		}

	}
	log.Println("policy id ", q.Id)
	// return the ID for the script
	return string(q.Id), nil
}
func GetPolicyAssignmentId(ResourceData *schema.ResourceData,policyAssignmentName string,projectId string ,username string, password string) (string, error) {
	log.Println("inside policy assignment id")
	var url bytes.Buffer
	q := new(Result)

	// Create the URL for the call to the Agility API
	url.WriteString(configuration.APIURL)
	url.WriteString("current/project/")
	url.WriteString(projectId)
	url.WriteString("/policyassignment/search")

	log.Println("URL:>", url.String())

	// Set the right HTTP Verb, and setup HTTP Basic Security
	req, err := http.NewRequest("GET", url.String(), nil)
	req.SetBasicAuth(username, password)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// make the HTTPS request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// log the response details for debugging
	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)

	//Stream the response body into a byte array
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))

	//Parse the XML
	r := strings.NewReader(string(body))
	decoder := xml.NewDecoder(r)
	finish := false
	for {
		// Read tokens from the XML document in a stream.
		t, _ := decoder.Token()
		if t == nil {
			return "", errors.New("there are no policy assignments with this name")
		}
		if finish {
			break
		}
		// look for <link> element
		switch Element := t.(type) {
		case xml.StartElement:
			if Element.Name.Local == "link" {
				log.Println("Element name is : ", Element.Name.Local)

				// unmarshal the element into generic structure
				err := decoder.DecodeElement(&q, &Element)
				if err != nil {
					log.Println(err)
				}

				// if the script name matches the script defined to Terraform
				// then we are are the right place, so stop looking
				log.Println("Element value is :", string(q.Name))
				if string(q.Name) == policyAssignmentName {
					log.Println("Found the policy assignment: ", q.Name)
					finish = true
					break
				}
			}
			// if the element is the <Linklist> then go again
			if Element.Name.Local == "Linklist" {
				log.Println("Element name is : ", Element.Name.Local)
			} else {
				log.Println("Unknown Element name is : ", Element.Name.Local)
			}
		default:
		}

	}
	log.Println("policyassignment id ", q.Id)
	// return the ID for the script
	return string(q.Id), nil
}

func CreateStack (username string, password string) []byte{

	var url bytes.Buffer
	// Create the URL for the call to the Agility API
	url.WriteString(configuration.APIURL)
	url.WriteString("current/stack/")
	log.Println("URL:>", url.String())
	//reading payload from AgilityLicense file
	file, err1 := ioutil.ReadFile("./agility/api/stackwithpolicy.xml")
	if err1 != nil {
		log.Println("error:", err1)
	}

	//Payload code ends

	log.Println("URL:>",url.String())
	req, err := http.NewRequest("POST", url.String(),bytes.NewBuffer([]byte(file)))
	req.Header.Set("Content-Type", "application/xml; charset=utf-8")
	req.SetBasicAuth(username, password)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// make the HTTPS request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// log the response details for debugging
	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)

	//Stream the response body into a byte array and return it
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", resp.Body)
	return body
}

func Publish(ResourceData *schema.ResourceData,username string, password string)([]byte, error){

	f, errf := os.OpenFile("./agility/api/agility.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if errf != nil {
		log.Println("error opening file: ", errf)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println("inside publish func")
	var url bytes.Buffer
	// Create the URL for the call to the Agility API
	url.WriteString(configuration.APIURL)
	url.WriteString("current/storeproduct/")
	log.Println("URL:>", url.String())

	productname:=ResourceData.Get("productname").(string)
	productdesc:=ResourceData.Get("productdesc").(string)
	producttype:=ResourceData.Get("producttype").(string)
	itemtype:=ResourceData.Get("itemtype").(string)
	itemname:=ResourceData.Get("itemname").(string)
	category:=ResourceData.Get("category").(string)
	operatingsystem:=ResourceData.Get("operatingsystem").(string)
	assetid,err:=GetBlueprintIdAfterCheckin(productname,username,password)
	if err != nil {
		log.Println("cannot get blueprint id ", err)
	}

	//payload
	var payload bytes.Buffer
	payload.WriteString(`<ns1:StoreProduct xmlns:ns1="http://servicemesh.com/agility/api"><ns1:name>`)
	payload.WriteString(productname)
	payload.WriteString(`</ns1:name><ns1:description>`)
	payload.WriteString(productdesc)
	payload.WriteString(`</ns1:description><ns1:productType><ns1:name>`)
	payload.WriteString(producttype)
	payload.WriteString(`</ns1:name></ns1:productType><ns1:itemId>`)
	payload.WriteString(assetid)
	payload.WriteString(`</ns1:itemId><ns1:itemType>`)
	payload.WriteString(itemtype)
	payload.WriteString(`</ns1:itemType><ns1:itemName>`)
	payload.WriteString(itemname)
	payload.WriteString(`</ns1:itemName><ns1:category><ns1:name>`)
	payload.WriteString(category)
	payload.WriteString(`</ns1:name></ns1:category><ns1:category><ns1:name>`)
	payload.WriteString(operatingsystem)
	payload.WriteString(`</ns1:name></ns1:category></ns1:StoreProduct>`)
	payload1 := payload.String()
	log.Println("payload====>",payload1)

	req, err := http.NewRequest("POST",url.String(),bytes.NewBuffer([]byte(payload1)))
	req.Header.Set("Content-Type", "application/xml; charset=utf-8")
	req.SetBasicAuth(username,password)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// make the HTTPS request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// log the response details for debugging
	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)

	//Stream the response body into a byte array and return it
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))
	return body,nil

}

func GetBlueprintIdAfterCheckin(blueprintName string,username string, password string) (string, error) {
	log.Println("The Blueprint name is: ", blueprintName)
	var url bytes.Buffer
	q := new(Result)

	// Create the URL for the call to the Agility API
	url.WriteString(configuration.APIURL)
	url.WriteString("current/blueprint/")
	log.Println("URL:>", url.String())

	// Set the right HTTP Verb, and setup HTTP Basic Security
	req, err := http.NewRequest("GET", url.String(), nil)
	req.SetBasicAuth(username, password)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// make the HTTPS request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// log the response details for debugging
	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)

	//Stream the response body into a byte array
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))

	//Parse the XML
	r := strings.NewReader(string(body))
	decoder := xml.NewDecoder(r)
	finish := false
	for {
		// Read tokens from the XML document in a stream.
		t, _ := decoder.Token()
		if t == nil {
			return "", errors.New("there are no Blueprints with this name")
		}
		if finish {
			break
		}
		// look for <link> element
		switch Element := t.(type) {
		case xml.StartElement:
			if Element.Name.Local == "link" {
				log.Println("Element name is : ", Element.Name.Local)

				// unmarshal the element into generic structure
				err := decoder.DecodeElement(&q, &Element)
				if err != nil {
					log.Println(err)
				}

				// if the blueprint name matches the blueprint defined to Terraform
				// then we are are the right place, so stop looking
				log.Println("Element value is :", string(q.Name))
				if string(q.Name) == blueprintName {
					log.Println("Found the Blueprint : ", q.Name)
					finish = true
					break
				}
			}
		default:
		}

	}

	// return the ID for the blueprint
	return string(q.Id), nil
}
func WriteScriptIdToFile(ResourceData *schema.ResourceData,scriptname,projectId,username string, password string)(error){

	scriptId,err :=  GetScripttId(ResourceData,scriptname,projectId,username,password)
	if err != nil {
		log.Println("error in getting the script id: ", err)
	}
	f, errf := os.OpenFile("./agility/attachments/id.txt", os.O_RDWR | os.O_CREATE, 0666)
	if errf != nil {
		log.Println("error opening file: ", errf)
	}

	f.WriteString(scriptId)
	defer f.Close()
	return nil
}
