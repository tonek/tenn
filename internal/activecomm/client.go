package activecomm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const jsonCT = "application/json"

func NewClient() *Client {
	return &Client{
		baseURL: "https://anc.apm.activecommunities.com",
		cl: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

type Client struct {
	baseURL string
	cl      *http.Client
}

func (c *Client) GetReservations(
	req ReservationTimeGroupRequest,
) (*ReservationTimeGroupResponse, error) {
	url := fmt.Sprintf("%s/seattle/rest/reservation/resource/reservationtimegroup", c.baseURL)
	resp, err := c.post(url, req)
	if err != nil {
		return nil, err
	}
	res := &ReservationTimeGroupResponse{}
	if err := json.Unmarshal(resp, res); err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) FindResources(
	name string,
) (*ResourceResponse, error) {
	url := fmt.Sprintf("%s/seattle/rest/reservation/resource", c.baseURL)
	req := ResourceRequest{
		Name:          name,
		FacilityTypes: []uint32{39, 115},
		PageSize:      100,
	}
	resp, err := c.post(url, req)
	if err != nil {
		return nil, err
	}
	res := &ResourceResponse{}
	if err := json.Unmarshal(resp, res); err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) post(url string, req interface{}) ([]byte, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	resp, err := c.cl.Post(url, jsonCT, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("failed request to [%s] with status [%v]: %v",
			url, resp.StatusCode, string(body))
	}
	// fmt.Printf("Faull body: %s\n", string(body))
	return body, nil
}

func AsDate(t time.Time) string {
	return t.Format("2006-01-02 15:00:00")
}

type ReservationTimeGroupRequest struct {
	ResourceID      string           `json:"resource_id"`
	Periods         []DateTimePeriod `json:"datetime_periods"`
	ReservationUnit uint32           `json:"reservation_unit"`
}

type DateTimePeriod struct {
	From string `json:"from_date_time"`
	To   string `json:"to_date_time"`
}

type ReservationTimeGroupResponse struct {
	Body ReservationTimeGroupResponseBody `json:"body"`
}

type ReservationTimeGroupResponseBody struct {
	ReservationTimes []ReservationTime `json:"reservation_times"`
}

type ReservationTime struct {
	Availability string `json:"availability"`
	Start        string `json:"start_event_datetime"`
	End          string `json:"end_event_datetime"`
}

type ResourceRequest struct {
	Name          string   `json:"name"`
	Attendee      uint32   `json:"attendee"`
	FacilityTypes []uint32 `json:"facility_type_ids"` // 39, 115
	PageSize      uint32   `json:"page_size"`
}

type ResourceResponse struct {
	Body ResourceResponseBody
}

type ResourceResponseBody struct {
	Items []Resource `json:"items"`
}

type Resource struct {
	ID       uint32 `json:"id"`
	Name     string `json:"name"`
	TypeName string `json:"type_name"`
}
