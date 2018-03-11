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
    url.WriteString("blueprint/") 
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
    url.WriteString("blueprint/") 
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
    url.WriteString("blueprint/") 
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
    url.WriteString("topology/") 
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
    url.WriteString("topology/") 
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
    url.WriteString("topology/") 
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
    url.WriteString("blueprint/") 
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
    url.WriteString("topology/") 
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
    url.WriteString("topology/") 
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
    url.WriteString("task/") 
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
    url.WriteString("project") 
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
    url.WriteString("template/search?qterm.field.creator.name=") 
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
    url.WriteString("compute/") 
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
    url.WriteString("compute/") 
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
    url.WriteString("project/") 
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
    url.WriteString("project/") 
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
    url.WriteString("blueprint/") 
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
    url.WriteString("blueprint/") 
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
    url.WriteString("project/") 
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
	url.WriteString("container")
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
	url.WriteString("container/")
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
	url.WriteString("container/")
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
	url.WriteString("project/")
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
	url.WriteString("container/")
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
	url.WriteString("container/")
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
	url.WriteString("project/")
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
	url.WriteString("cloud")
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
	url.WriteString("cloud/")
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
    url.WriteString("cloud/")
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
    //req.Header.Set("Content-Type", "text; charset=utf-8")
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
