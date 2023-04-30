package main

import "time"

type Database struct {
	Object string `json:"object"`
	Id     string `json:"Id"`
	Icon   struct {
		Type  string `json:"type"`
		Emoji string `json:"emoji"`
	} `json:"icon"`
	CreatedTime time.Time `json:"created_time"`
	CreatedBy   struct {
		Object string `json:"object"`
		Id     string `json:"Id"`
	} `json:"created_by"`
	LastEditedBy struct {
		Object string `json:"object"`
		Id     string `json:"Id"`
	} `json:"last_edited_by"`
	LastEditedTime time.Time `json:"last_edited_time"`
	Title          []struct {
		Type string `json:"type"`
		Text struct {
			Content string `json:"content"`
			Link    string `json:"link"`
		} `json:"text"`
		Annotations struct {
			Bold          bool   `json:"bold"`
			Italic        bool   `json:"italic"`
			Strikethrough bool   `json:"strikethrough"`
			Underline     bool   `json:"underline"`
			Code          bool   `json:"code"`
			Color         string `json:"color"`
		} `json:"annotations"`
		PlainText string `json:"plain_text"`
		Href      string `json:"href"`
	} `json:"title"`
	Description []string `json:"description"`
	IsInline    bool     `json:"is_inline"`
	Properties  struct {
		Status struct {
			Id     string `json:"Id"`
			Name   string `json:"name"`
			Type   string `json:"type"`
			Status struct {
				Options []struct {
					Id    string `json:"Id"`
					Name  string `json:"name"`
					Color string `json:"color"`
				} `json:"options"`
				Groups []struct {
					Id        string   `json:"Id"`
					Name      string   `json:"name"`
					Color     string   `json:"color"`
					OptionIds []string `json:"option_Ids"`
				} `json:"groups"`
			} `json:"status"`
		} `json:"Status"`
		Assign struct {
			Id     string `json:"Id"`
			Name   string `json:"name"`
			Type   string `json:"type"`
			People struct {
			} `json:"people"`
		} `json:"Assign"`
		Name struct {
			Id    string `json:"Id"`
			Name  string `json:"name"`
			Type  string `json:"type"`
			Title struct {
			} `json:"title"`
		} `json:"Name"`
	} `json:"properties"`
	Parent struct {
		Type      string `json:"type"`
		Workspace bool   `json:"workspace"`
	} `json:"parent"`
	Url      string `json:"url"`
	Archived bool   `json:"archived"`
}

type QueryResult struct {
	Object  string `json:"object"`
	Results []struct {
		Object         string    `json:"object"`
		Id             string    `json:"Id"`
		CreatedTime    time.Time `json:"created_time"`
		LastEditedTime time.Time `json:"last_edited_time"`
		CreatedBy      struct {
			Object string `json:"object"`
			Id     string `json:"Id"`
		} `json:"created_by"`
		LastEditedBy struct {
			Object string `json:"object"`
			Id     string `json:"Id"`
		} `json:"last_edited_by"`
		Parent struct {
			Type       string `json:"type"`
			DatabaseId string `json:"database_Id"`
		} `json:"parent"`
		Archived   bool `json:"archived"`
		Properties struct {
			State struct {
				Id     string `json:"Id"`
				Type   string `json:"type"`
				Status struct {
					Id    string `json:"Id"`
					Name  string `json:"name"`
					Color string `json:"color"`
				} `json:"status"`
			} `json:"State"`
			Assign struct {
				Id     string   `json:"Id"`
				Type   string   `json:"type"`
				People []string `json:"people"`
			} `json:"Assign"`
			Name struct {
				Id    string `json:"Id"`
				Type  string `json:"type"`
				Title []struct {
					Type string `json:"type"`
					Text struct {
						Content string `json:"content"`
						Link    string `json:"link"`
					} `json:"text"`
					Annotations struct {
						Bold          bool   `json:"bold"`
						Italic        bool   `json:"italic"`
						Strikethrough bool   `json:"strikethrough"`
						Underline     bool   `json:"underline"`
						Code          bool   `json:"code"`
						Color         string `json:"color"`
					} `json:"annotations"`
					PlainText string `json:"plain_text"`
					Href      string `json:"href"`
				} `json:"title"`
			} `json:"Name"`
		} `json:"properties"`
		Url string `json:"url"`
	} `json:"results"`
	NextCursor string `json:"next_cursor"`
	HasMore    bool   `json:"has_more"`
	Type       string `json:"type"`
	Page       struct {
	} `json:"page"`
}

// type Property string
type Status struct {
	Equals string `json:"equals"`
}

type Filter struct {
	Property string `json:"property"`
	Status   Status `json:"status"`
}

type FilterPages struct {
	Filter Filter `json:"filter"`
}
