package cohesivenetv1

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
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
		if i.ImportID == imageId {
			pluginImage.Images = []Image{i}
		}
	}

	return pluginImage, nil
}

func (c *Client) CreateImage(image *PluginImage) (CreateImageResponse, error) {

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
	return newImage, err
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
