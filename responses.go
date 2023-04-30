package main

import "time"

type Database struct {
	Object string `json:"object"`
	ID     string `json:"id"`
	Icon   struct {
		Type  string `json:"type"`
		Emoji string `json:"emoji"`
	} `json:"icon"`
	CreatedTime time.Time `json:"created_time"`
	CreatedBy   struct {
		Object string `json:"object"`
		ID     string `json:"id"`
	} `json:"created_by"`
	LastEditedBy struct {
		Object string `json:"object"`
		ID     string `json:"id"`
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
			ID     string `json:"id"`
			Name   string `json:"name"`
			Type   string `json:"type"`
			Status struct {
				Options []struct {
					ID    string `json:"id"`
					Name  string `json:"name"`
					Color string `json:"color"`
				} `json:"options"`
				Groups []struct {
					ID        string   `json:"id"`
					Name      string   `json:"name"`
					Color     string   `json:"color"`
					OptionIds []string `json:"option_ids"`
				} `json:"groups"`
			} `json:"status"`
		} `json:"Status"`
		Assign struct {
			ID     string `json:"id"`
			Name   string `json:"name"`
			Type   string `json:"type"`
			People struct {
			} `json:"people"`
		} `json:"Assign"`
		Name struct {
			ID    string `json:"id"`
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
	URL      string `json:"url"`
	Archived bool   `json:"archived"`
}
