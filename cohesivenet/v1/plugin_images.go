package cohesivenetv1

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func (c *Client) GetImage(imageId string) (ImageResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/container_system/images", c.HostURL), nil)
	if err != nil {
		log.Println(err)
	}
	body, err := c.doRequest(req)
	if err != nil {
		log.Println(err)
	}

	simpleJson := SimplifyImageJson(string(body), imageId)
	imageResponse := ImageResponse{}
	errUnmarshal := json.Unmarshal([]byte(simpleJson), &imageResponse)
	if err != nil {
		log.Println(errUnmarshal)
	}

	pluginImage := ImageResponse{}
	for _, i := range imageResponse.Images {
		if i.ID == imageId {
			pluginImage.Images = []Image{i}
		}
	}

	return pluginImage, nil
}

func (c *Client) CreateImage(image *PluginImage) (ImageResponse, error) {

	// synchronize creating a plugin image
	c.ReqLock.Lock()
	defer c.ReqLock.Unlock()

	rb, err := json.Marshal(image)
	if err != nil {
		log.Println(err)
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/container_system/images", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		log.Println(err)
	}

	body, err2 := c.doRequest(req)
	if err2 != nil {
		log.Println(err2)
	}
	newImage := CreateImageResponse{}
	errUnmarshal := json.Unmarshal([]byte(body), &newImage)
	if err != nil {
		log.Println(errUnmarshal)
	}

	// at this point, we have to poll for the image
	// until it is created and returns an id
	// we will poll by default for 300 seconds, every 5 seconds
	// default can be overridden with VNS3_IMAGE_POLLING_MAX env variable
	var response ImageResponse
	polling_max_in_seconds := ImagePollingMax()
	polling_interval := 5 //in seconds i.e. 5 seconds
	counter_limit := polling_max_in_seconds / polling_interval
	counter := 0

	timer := time.Tick(5 * time.Second)
	for _ = range timer {
		req, err := http.NewRequest("GET", fmt.Sprintf("%s/container_system/images", c.HostURL), nil)
		if err != nil {
			log.Println(err)
		}
		body, err := c.doRequest(req)
		if err != nil {
			log.Println(err)
		}

		simpleJson := SimplifyImageJson(string(body), newImage.NewImage.Import_uuid)
		imageResponse := ImageResponse{}
		errUnmarshal := json.Unmarshal([]byte(simpleJson), &imageResponse)
		if err != nil {
			log.Println(errUnmarshal)
		}

		for _, i := range imageResponse.Images {
			if i.ImportID == newImage.NewImage.Import_uuid && len(i.ID) != 0 {
				response.Images = []Image{i}
				break
			}
		}
		if len(response.Images) != 0 {
			break
		}
		counter++
		// can't poll forever, jump out after limit_counter
		if counter > counter_limit {
			return response, fmt.Errorf("Timeout error while polling for image after creation."+
				" Increase current polling seconds (%v) by setting VNS3_IMAGE_POLLING_MAX env variable", polling_max_in_seconds)
		}
	}
	return response, err
}

func ImagePollingMax() int {
	val := os.Getenv("VNS3_IMAGE_POLLING_MAX")
	if val == "" {
		val = "300"
	}
	polling, err := strconv.Atoi(val)
	if err != nil {
		return 300
	}
	return polling
}

func (c *Client) DeleteImage(imageId string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/container_system/images/%s", c.HostURL, imageId), nil)
	if err != nil {
		return err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return err
	}

	if string(body) == "Deleted route" {
		return err
	}

	return nil
}

func SimplifyImageJson(input string, import_uuid string) string {
	var result map[string]interface{}
	json.Unmarshal([]byte(input), &result)

	var response = result["response"].(map[string]interface{})
	var images = response["images"].([]interface{})
	var list string = "{\"response\" : [ "

	for _, image := range images {
		list += "{ "
		values := image.(map[string]interface{})
		list += "\"id\":\"" + unmarshalImageString(values, "id") + "\","
		list += "\"image_name\":\"" + unmarshalImageString(values, "image_name") + "\","
		list += "\"tag_name\":\"" + unmarshalImageString(values, "tag_name") + "\","
		list += "\"status\":\"" + unmarshalImageString(values, "status") + "\","
		list += "\"status_msg\":\"" + unmarshalImageString(values, "status_msg") + "\","
		list += "\"import_id\":\"" + unmarshalImageString(values, "import_id") + "\","
		list += "\"created\":\"" + unmarshalImageString(values, "created") + "\"," // date
		list += "\"description\":\"" + unmarshalImageString(values, "description") + "\","
		list += "\"comment\":\"" + unmarshalImageString(values, "comment") + "\","
		list += "\"import_uuid\":\"" + import_uuid + "\","
		list = strings.TrimSuffix(list, ",") //remove last comma
		list += " },"

	}
	list = strings.TrimSuffix(list, ",") //remove last comma
	list += "]}"
	return list
}

// handles missing properties and null values for strings
func unmarshalImageString(values map[string]interface{}, name string) string {
	if val, ok := values[name]; ok {
		if val != nil {
			return fmt.Sprintf("%v", val)
		}
	}
	return ""
}
