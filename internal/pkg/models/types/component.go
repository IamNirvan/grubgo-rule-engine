package types

/*
	{
		"type": "component", // Have other types like mail, sms, etc...
		"payload": {
			// The content in here will vary depending on the type...
			// This is for a component...
			"type": "TAG", // Have other types like pop-up, etc.
			"status": "NEGATIVE" // Have other statuses like NEURAL and POSITIVE. Can use this to adjust colors, etc..
			"mainText": "This dish happens to have an ingredient you are allergic to"
			"secondaryText": null
		}
	}
*/

// type Component struct {
// 	Type string `json:"type"`
// 	Mood string `json:"mood"`
// 	Text string `json:"text"`
// }

type Component struct {
	Type   string `json:"type"`
	Status string `json:"status"`
	Text   string `json:"text"`
}
