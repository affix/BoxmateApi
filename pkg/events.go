package boxmate

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type EventResponse struct {
	Success bool    `json:"success"`
	Events  []Event `json:"results"`
}

type Event struct {
	ID                 int         `json:"id"`
	Object             string      `json:"object"`
	Name               string      `json:"name"`
	Description        string      `json:"description"`
	EventStatus        string      `json:"event_status"`
	AttendanceStatus   string      `json:"attendance_status"`
	StartTime          string      `json:"start_time"`
	EndTime            string      `json:"end_time"`
	UtcStart           string      `json:"utc_start"`
	RegistrationsOpen  string      `json:"registrations_open"`
	RegistrationsClose string      `json:"registrations_close"`
	CancellationCutoff string      `json:"cancellation_cutoff"`
	AttendingCount     int         `json:"attending_count"`
	WaitingCount       int         `json:"waiting_count"`
	AttendingMax       int         `json:"attending_max"`
	WaitlistMax        interface{} `json:"waitlist_max"`
	WaitlistPosition   interface{} `json:"waitlist_position"`
	OfferingType       struct {
		Name        string `json:"name"`
		ID          int    `json:"id"`
		Object      string `json:"object"`
		Description string `json:"description"`
		Type        string `json:"type"`
		Status      string `json:"status"`
	} `json:"offering_type"`
	Venue struct {
		ID          int         `json:"id"`
		Object      string      `json:"object"`
		Name        string      `json:"name"`
		Description interface{} `json:"description"`
		Address     string      `json:"address"`
		Timezone    string      `json:"timezone"`
		Lat         float64     `json:"lat"`
		Lng         float64     `json:"lng"`
	} `json:"venue"`
	Instructors []struct {
		ID          int         `json:"id"`
		Object      string      `json:"object"`
		Name        string      `json:"name"`
		PictureURL  interface{} `json:"picture_url"`
		Description string      `json:"description"`
	} `json:"instructors"`
	LateCancelMessage string `json:"late_cancel_message,omitempty"`
}

// GetTeamupEvents return teamup events
// params :
//   - apiKey string
//   - date string (2020-01-01)
func GetTeamupEventsForDate(apiKey string, date string) (EventResponse, error) {
	client := http.Client{}
	resp, err := client.PostForm("https://api.boxmateapp.co.uk/teamup/customer/events/"+date, url.Values{
		"member_token": {apiKey},
	})
	if err != nil {
		return EventResponse{}, err
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return EventResponse{}, err
	}
	var result EventResponse
	json.Unmarshal(data, &result)
	return result, nil
}
