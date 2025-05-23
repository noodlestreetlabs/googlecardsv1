// We "forked" the v1 chat API so that we could address a couple of small bugs
// (namely that the GridItem didn't have a textAlignment field).
//
// 1. cp ~/go/pkg/mod/google.golang.org/api@v0.158.0/chat/v1/chat-gen.go chat-gen.go
// 2. delete everything not referenced by one of the GoogleAppsCardV1 structs.
// 3. replace GoogleAppsCardV1 -> ""

package googlecardsv1

import (
	"encoding/json"
)

// ChatClientDataSourceMarkup: For a `SelectionInput` widget that uses a
// multiselect menu, a data source from Google Chat. The data source
// populates selection items for the multiselect menu. For example, a
// user can select Google Chat spaces that they're a member of. Google
// Chat apps (https://developers.google.com/chat):
type ChatClientDataSourceMarkup struct {
	// SpaceDataSource: Google Chat spaces that the user is a member of.
	// NOTE: omitted
	// SpaceDataSource *SpaceDataSource `json:"spaceDataSource,omitempty"`

	// ForceSendFields is a list of field names (e.g. "SpaceDataSource") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "SpaceDataSource") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ChatClientDataSourceMarkup) MarshalJSON() ([]byte, error) {
	type NoMethod ChatClientDataSourceMarkup
	raw := NoMethod(*s)
	return marshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type Color struct {
	// Alpha: The fraction of this color that should be applied to the
	// pixel. That is, the final pixel color is defined by the equation:
	// `pixel color = alpha * (this color) + (1.0 - alpha) * (background
	// color)` This means that a value of 1.0 corresponds to a solid color,
	// whereas a value of 0.0 corresponds to a completely transparent color.
	// This uses a wrapper message rather than a simple float scalar so that
	// it is possible to distinguish between a default value and the value
	// being unset. If omitted, this color object is rendered as a solid
	// color (as if the alpha value had been explicitly given a value of
	// 1.0).
	Alpha float64 `json:"alpha,omitempty"`

	// Blue: The amount of blue in the color as a value in the interval [0,
	// 1].
	Blue float64 `json:"blue,omitempty"`

	// Green: The amount of green in the color as a value in the interval
	// [0, 1].
	Green float64 `json:"green,omitempty"`

	// Red: The amount of red in the color as a value in the interval [0,
	// 1].
	Red float64 `json:"red,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Alpha") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Alpha") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Color) MarshalJSON() ([]byte, error) {
	type NoMethod Color
	raw := NoMethod(*s)
	return marshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CommonEventObject: Represents information about the user's client,
// such as locale, host app, and platform. For Chat apps,
// `CommonEventObject` includes data submitted by users interacting with
// cards, like data entered in dialogs
// (https://developers.google.com/chat/how-tos/dialogs).
type CommonEventObject struct {
	// FormInputs: A map containing the current values of the widgets in a
	// card. The map keys are the string IDs assigned to each widget, and
	// the values represent inputs to the widget. Depending on the input
	// data type, a different object represents each input: For single-value
	// widgets, `StringInput`. For multi-value widgets, an array of
	// `StringInput` objects. For a date-time picker, a `DateTimeInput`. For
	// a date-only picker, a `DateInput`. For a time-only picker, a
	// `TimeInput`. Corresponds with the data entered by a user on a card in
	// a dialog (https://developers.google.com/chat/how-tos/dialogs).
	FormInputs map[string]Inputs `json:"formInputs,omitempty"`

	// HostApp: The hostApp enum which indicates the app the add-on is
	// invoked from. Always `CHAT` for Chat apps.
	//
	// Possible values:
	//   "UNSPECIFIED_HOST_APP" - Google can't identify a host app.
	//   "GMAIL" - The add-on launches from Gmail.
	//   "CALENDAR" - The add-on launches from Google Calendar.
	//   "DRIVE" - The add-on launches from Google Drive.
	//   "DEMO" - Not used.
	//   "DOCS" - The add-on launches from Google Docs.
	//   "MEET" - The add-on launches from Google Meet.
	//   "SHEETS" - The add-on launches from Google Sheets.
	//   "SLIDES" - The add-on launches from Google Slides.
	//   "DRAWINGS" - The add-on launches from Google Drawings.
	//   "CHAT" - A Google Chat app. Not used for Google Workspace Add-ons.
	HostApp string `json:"hostApp,omitempty"`

	// InvokedFunction: Name of the invoked function associated with the
	// widget. Only set for Chat apps.
	InvokedFunction string `json:"invokedFunction,omitempty"`

	// Parameters: Custom parameters
	// (/chat/api/reference/rest/v1/cards#ActionParameter) passed to the
	// invoked function. Both keys and values must be strings.
	Parameters map[string]string `json:"parameters,omitempty"`

	// Platform: The platform enum which indicates the platform where the
	// event originates (`WEB`, `IOS`, or `ANDROID`). Not supported by Chat
	// apps.
	//
	// Possible values:
	//   "UNKNOWN_PLATFORM"
	//   "WEB"
	//   "IOS"
	//   "ANDROID"
	Platform string `json:"platform,omitempty"`

	// TimeZone: The timezone ID and offset from Coordinated Universal Time
	// (UTC). Only supported for the event types `CARD_CLICKED`
	// (https://developers.google.com/chat/api/reference/rest/v1/EventType#ENUM_VALUES.CARD_CLICKED)
	// and `SUBMIT_DIALOG`
	// (https://developers.google.com/chat/api/reference/rest/v1/DialogEventType#ENUM_VALUES.SUBMIT_DIALOG).
	TimeZone *TimeZone `json:"timeZone,omitempty"`

	// UserLocale: The full `locale.displayName` in the format of [ISO 639
	// language code]-[ISO 3166 country/region code] such as "en-US".
	UserLocale string `json:"userLocale,omitempty"`

	// ForceSendFields is a list of field names (e.g. "FormInputs") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "FormInputs") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CommonEventObject) MarshalJSON() ([]byte, error) {
	type NoMethod CommonEventObject
	raw := NoMethod(*s)
	return marshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Action: An action that describes the behavior when
// the form is submitted. For example, you can invoke an Apps Script
// script to handle the form. If the action is triggered, the form
// values are sent to the server. Google Workspace Add-ons and Chat apps
// (https://developers.google.com/workspace/extend):
type Action struct {
	// Function: A custom function to invoke when the containing element is
	// clicked or othrwise activated. For example usage, see Create
	// interactive cards
	// (https://developers.google.com/chat/how-tos/cards-onclick).
	Function string `json:"function,omitempty"`

	// Interaction: Optional. Required when opening a dialog
	// (https://developers.google.com/chat/how-tos/dialogs). What to do in
	// response to an interaction with a user, such as a user clicking a
	// button in a card message. If unspecified, the app responds by
	// executing an `action`—like opening a link or running a
	// function—as normal. By specifying an `interaction`, the app can
	// respond in special interactive ways. For example, by setting
	// `interaction` to `OPEN_DIALOG`, the app can open a dialog
	// (https://developers.google.com/chat/how-tos/dialogs). When specified,
	// a loading indicator isn't shown. If specified for an add-on, the
	// entire card is stripped and nothing is shown in the client. Google
	// Chat apps (https://developers.google.com/chat):
	//
	// Possible values:
	//   "INTERACTION_UNSPECIFIED" - Default value. The `action` executes as
	// normal.
	//   "OPEN_DIALOG" - Opens a
	// [dialog](https://developers.google.com/chat/how-tos/dialogs), a
	// windowed, card-based interface that Chat apps use to interact with
	// users. Only supported by Chat apps in response to button-clicks on
	// card messages. If specified for an add-on, the entire card is
	// stripped and nothing is shown in the client. [Google Chat
	// apps](https://developers.google.com/chat):
	Interaction string `json:"interaction,omitempty"`

	// LoadIndicator: Specifies the loading indicator that the action
	// displays while making the call to the action.
	//
	// Possible values:
	//   "SPINNER" - Displays a spinner to indicate that content is loading.
	//   "NONE" - Nothing is displayed.
	LoadIndicator string `json:"loadIndicator,omitempty"`

	// Parameters: List of action parameters.
	Parameters []*ActionParameter `json:"parameters,omitempty"`

	// PersistValues: Indicates whether form values persist after the
	// action. The default value is `false`. If `true`, form values remain
	// after the action is triggered. To let the user make changes while the
	// action is being processed, set `LoadIndicator`
	// (https://developers.google.com/workspace/add-ons/reference/rpc/google.apps.card.v1#loadindicator)
	// to `NONE`. For card messages
	// (https://developers.google.com/chat/api/guides/v1/messages/create#create)
	// in Chat apps, you must also set the action's `ResponseType`
	// (https://developers.google.com/chat/api/reference/rest/v1/spaces.messages#responsetype)
	// to `UPDATE_MESSAGE` and use the same `card_id`
	// (https://developers.google.com/chat/api/reference/rest/v1/spaces.messages#CardWithId)
	// from the card that contained the action. If `false`, the form values
	// are cleared when the action is triggered. To prevent the user from
	// making changes while the action is being processed, set
	// `LoadIndicator`
	// (https://developers.google.com/workspace/add-ons/reference/rpc/google.apps.card.v1#loadindicator)
	// to `SPINNER`.
	PersistValues bool `json:"persistValues,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Function") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Function") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Action) MarshalJSON() ([]byte, error) {
	type NoMethod Action
	raw := NoMethod(*s)
	return marshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ActionParameter: List of string parameters to supply
// when the action method is invoked. For example, consider three snooze
// buttons: snooze now, snooze one day, or snooze next week. You might
// use `action method = snooze()`, passing the snooze type and snooze
// time in the list of string parameters. To learn more, see
// `CommonEventObject`
// (https://developers.google.com/chat/api/reference/rest/v1/Event#commoneventobject).
// Google Workspace Add-ons and Chat apps
// (https://developers.google.com/workspace/extend):
type ActionParameter struct {
	// Key: The name of the parameter for the action script.
	Key string `json:"key,omitempty"`

	// Value: The value of the parameter.
	Value string `json:"value,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Key") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Key") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ActionParameter) MarshalJSON() ([]byte, error) {
	type NoMethod ActionParameter
	raw := NoMethod(*s)
	return marshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// BorderStyle: The style options for the border of a
// card or widget, including the border type and color. Google Workspace
// Add-ons and Chat apps
// (https://developers.google.com/workspace/extend):
type BorderStyle struct {
	// CornerRadius: The corner radius for the border.
	CornerRadius int64 `json:"cornerRadius,omitempty"`

	// StrokeColor: The colors to use when the type is `BORDER_TYPE_STROKE`.
	StrokeColor *Color `json:"strokeColor,omitempty"`

	// Type: The border type.
	//
	// Possible values:
	//   "BORDER_TYPE_UNSPECIFIED" - Don't use. Unspecified.
	//   "NO_BORDER" - Default value. No border.
	//   "STROKE" - Outline.
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CornerRadius") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CornerRadius") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *BorderStyle) MarshalJSON() ([]byte, error) {
	type NoMethod BorderStyle
	raw := NoMethod(*s)
	return marshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Button: A text, icon, or text and icon button that
// users can click. For an example in Google Chat apps, see Button list
// (https://developers.google.com/chat/ui/widgets/button-list). To make
// an image a clickable button, specify an `Image` (not an
// `ImageComponent`) and set an `onClick` action. Google Workspace
// Add-ons and Chat apps
// (https://developers.google.com/workspace/extend):
type Button struct {
	// AltText: The alternative text that's used for accessibility. Set
	// descriptive text that lets users know what the button does. For
	// example, if a button opens a hyperlink, you might write: "Opens a new
	// browser tab and navigates to the Google Chat developer documentation
	// at https://developers.google.com/chat".
	AltText string `json:"altText,omitempty"`

	// Color: If set, the button is filled with a solid background color and
	// the font color changes to maintain contrast with the background
	// color. For example, setting a blue background likely results in white
	// text. If unset, the image background is white and the font color is
	// blue. For red, green, and blue, the value of each field is a `float`
	// number that you can express in either of two ways: as a number
	// between 0 and 255 divided by 255 (153/255), or as a value between 0
	// and 1 (0.6). 0 represents the absence of a color and 1 or 255/255
	// represent the full presence of that color on the RGB scale.
	// Optionally set `alpha`, which sets a level of transparency using this
	// equation: ``` pixel color = alpha * (this color) + (1.0 - alpha) *
	// (background color) ``` For `alpha`, a value of `1` corresponds with a
	// solid color, and a value of `0` corresponds with a completely
	// transparent color. For example, the following color represents a half
	// transparent red: ``` "color": { "red": 1, "green": 0, "blue": 0,
	// "alpha": 0.5 } ```
	Color *Color `json:"color,omitempty"`

	// Disabled: If `true`, the button is displayed in an inactive state and
	// doesn't respond to user actions.
	Disabled bool `json:"disabled,omitempty"`

	// Icon: The icon image. If both `icon` and `text` are set, then the
	// icon appears before the text.
	Icon *Icon `json:"icon,omitempty"`

	// OnClick: Required. The action to perform when a user clicks the
	// button, such as opening a hyperlink or running a custom function.
	OnClick *OnClick `json:"onClick,omitempty"`

	// Text: The text displayed inside the button.
	Text string `json:"text,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AltText") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AltText") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Button) MarshalJSON() ([]byte, error) {
	type NoMethod Button
	raw := NoMethod(*s)
	return marshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ButtonList: A list of buttons layed out horizontally.
// For an example in Google Chat apps, see Button list
// (https://developers.google.com/chat/ui/widgets/button-list). Google
// Workspace Add-ons and Chat apps
// (https://developers.google.com/workspace/extend):
type ButtonList struct {
     	// Unclear if this works
     	HorizontalAlignment string `json:"horizontalAlignment,omitempty"`


	// Buttons: An array of buttons.
	Buttons []*Button `json:"buttons,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Buttons") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Buttons") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ButtonList) MarshalJSON() ([]byte, error) {
	type NoMethod ButtonList
	raw := NoMethod(*s)
	return marshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Card: A card interface displayed in a Google Chat
// message or Google Workspace Add-on. Cards support a defined layout,
// interactive UI elements like buttons, and rich media like images. Use
// cards to present detailed information, gather information from users,
// and guide users to take a next step. Card builder
// (https://addons.gsuite.google.com/uikit/builder) To learn how to
// build cards, see the following documentation: * For Google Chat apps,
// see Design dynamic, interactive, and consistent UIs with cards
// (https://developers.google.com/chat/ui). * For Google Workspace
// Add-ons, see Card-based interfaces
// (https://developers.google.com/apps-script/add-ons/concepts/cards).
// **Example: Card message for a Google Chat app** !Example contact card
// (https://developers.google.com/chat/images/card_api_reference.png) To
// create the sample card message in Google Chat, use the following
// JSON: ``` { "cardsV2": [ { "cardId": "unique-card-id", "card": {
// "header": { "title": "Sasha", "subtitle": "Software Engineer",
// "imageUrl":
// "https://developers.google.com/chat/images/quickstart-app-avatar.png",
//
//	"imageType": "CIRCLE", "imageAltText": "Avatar for Sasha", },
//
// "sections": [ { "header": "Contact Info", "collapsible": true,
// "uncollapsibleWidgetsCount": 1, "widgets": [ { "decoratedText": {
// "startIcon": { "knownIcon": "EMAIL", }, "text": "sasha@example.com",
// } }, { "decoratedText": { "startIcon": { "knownIcon": "PERSON", },
// "text": "Online", }, }, { "decoratedText": { "startIcon": {
// "knownIcon": "PHONE", }, "text": "+1 (555) 555-1234", } }, {
// "buttonList": { "buttons": [ { "text": "Share", "onClick": {
// "openLink": { "url": "https://example.com/share", } } }, { "text":
// "Edit", "onClick": { "action": { "function": "goToView",
// "parameters": [ { "key": "viewType", "value": "EDIT", } ], } } }, ],
// } }, ], }, ], }, } ], } ```
type Card struct {
	// CardActions: The card's actions. Actions are added to the card's
	// toolbar menu. Google Workspace Add-ons
	// (https://developers.google.com/workspace/add-ons): For example, the
	// following JSON constructs a card action menu with `Settings` and
	// `Send Feedback` options: ``` "card_actions": [ { "actionLabel":
	// "Settings", "onClick": { "action": { "functionName": "goToView",
	// "parameters": [ { "key": "viewType", "value": "SETTING" } ],
	// "loadIndicator": "LoadIndicator.SPINNER" } } }, { "actionLabel":
	// "Send Feedback", "onClick": { "openLink": { "url":
	// "https://example.com/feedback" } } } ] ```
	CardActions []*CardAction `json:"cardActions,omitempty"`

	// DisplayStyle: In Google Workspace Add-ons, sets the display
	// properties of the `peekCardHeader`. Google Workspace Add-ons
	// (https://developers.google.com/workspace/add-ons):
	//
	// Possible values:
	//   "DISPLAY_STYLE_UNSPECIFIED" - Don't use. Unspecified.
	//   "PEEK" - The header of the card appears at the bottom of the
	// sidebar, partially covering the current top card of the stack.
	// Clicking the header pops the card into the card stack. If the card
	// has no header, a generated header is used instead.
	//   "REPLACE" - Default value. The card is shown by replacing the view
	// of the top card in the card stack.
	DisplayStyle string `json:"displayStyle,omitempty"`

	// FixedFooter: The fixed footer shown at the bottom of this card.
	// Setting `fixedFooter` without specifying a `primaryButton` or a
	// `secondaryButton` causes an error. For Chat apps, you can use fixed
	// footers in dialogs
	// (https://developers.google.com/chat/how-tos/dialogs), but not card
	// messages
	// (https://developers.google.com/chat/api/guides/v1/messages/create#create).
	// Google Workspace Add-ons and Chat apps
	// (https://developers.google.com/workspace/extend):
	FixedFooter *CardFixedFooter `json:"fixedFooter,omitempty"`

	// Header: The header of the card. A header usually contains a leading
	// image and a title. Headers always appear at the top of a card.
	Header *CardHeader `json:"header,omitempty"`

	// Name: Name of the card. Used as a card identifier in card navigation.
	// Google Workspace Add-ons
	// (https://developers.google.com/workspace/add-ons):
	Name string `json:"name,omitempty"`

	// PeekCardHeader: When displaying contextual content, the peek card
	// header acts as a placeholder so that the user can navigate forward
	// between the homepage cards and the contextual cards. Google Workspace
	// Add-ons (https://developers.google.com/workspace/add-ons):
	PeekCardHeader *CardHeader `json:"peekCardHeader,omitempty"`

	// SectionDividerStyle: The divider style between sections.
	//
	// Possible values:
	//   "DIVIDER_STYLE_UNSPECIFIED" - Don't use. Unspecified.
	//   "SOLID_DIVIDER" - Default option. Render a solid divider between
	// sections.
	//   "NO_DIVIDER" - If set, no divider is rendered between sections.
	SectionDividerStyle string `json:"sectionDividerStyle,omitempty"`

	// Sections: Contains a collection of widgets. Each section has its own,
	// optional header. Sections are visually separated by a line divider.
	// For an example in Google Chat apps, see Card section
	// (https://developers.google.com/chat/ui/widgets/card-section).
	Sections []*Section `json:"sections,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CardActions") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CardActions") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Card) MarshalJSON() ([]byte, error) {
	type NoMethod Card
	raw := NoMethod(*s)
	return marshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CardAction: A card action is the action associated
// with the card. For example, an invoice card might include actions
// such as delete invoice, email invoice, or open the invoice in a
// browser. Google Workspace Add-ons
// (https://developers.google.com/workspace/add-ons):
type CardAction struct {
	// ActionLabel: The label that displays as the action menu item.
	ActionLabel string `json:"actionLabel,omitempty"`

	// OnClick: The `onClick` action for this action item.
	OnClick *OnClick `json:"onClick,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ActionLabel") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ActionLabel") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CardAction) MarshalJSON() ([]byte, error) {
	type NoMethod CardAction
	raw := NoMethod(*s)
	return marshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CardFixedFooter: A persistent (sticky) footer that
// that appears at the bottom of the card. Setting `fixedFooter` without
// specifying a `primaryButton` or a `secondaryButton` causes an error.
// For Chat apps, you can use fixed footers in dialogs
// (https://developers.google.com/chat/how-tos/dialogs), but not card
// messages
// (https://developers.google.com/chat/api/guides/v1/messages/create#create).
// For an example in Google Chat apps, see Card footer
// (https://developers.google.com/chat/ui/widgets/card-fixed-footer).
// Google Workspace Add-ons and Chat apps
// (https://developers.google.com/workspace/extend):
type CardFixedFooter struct {
	// PrimaryButton: The primary button of the fixed footer. The button
	// must be a text button with text and color set.
	PrimaryButton *Button `json:"primaryButton,omitempty"`

	// SecondaryButton: The secondary button of the fixed footer. The button
	// must be a text button with text and color set. If `secondaryButton`
	// is set, you must also set `primaryButton`.
	SecondaryButton *Button `json:"secondaryButton,omitempty"`

	// ForceSendFields is a list of field names (e.g. "PrimaryButton") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "PrimaryButton") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CardFixedFooter) MarshalJSON() ([]byte, error) {
	type NoMethod CardFixedFooter
	raw := NoMethod(*s)
	return marshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CardHeader: Represents a card header. For an example
// in Google Chat apps, see Card header
// (https://developers.google.com/chat/ui/widgets/card-header). Google
// Workspace Add-ons and Chat apps
// (https://developers.google.com/workspace/extend):
type CardHeader struct {
	// ImageAltText: The alternative text of this image that's used for
	// accessibility.
	ImageAltText string `json:"imageAltText,omitempty"`

	// ImageType: The shape used to crop the image. Google Workspace Add-ons
	// and Chat apps (https://developers.google.com/workspace/extend):
	//
	// Possible values:
	//   "SQUARE" - Default value. Applies a square mask to the image. For
	// example, a 4x3 image becomes 3x3.
	//   "CIRCLE" - Applies a circular mask to the image. For example, a 4x3
	// image becomes a circle with a diameter of 3.
	ImageType string `json:"imageType,omitempty"`

	// ImageUrl: The HTTPS URL of the image in the card header.
	ImageUrl string `json:"imageUrl,omitempty"`

	// Subtitle: The subtitle of the card header. If specified, appears on
	// its own line below the `title`.
	Subtitle string `json:"subtitle,omitempty"`

	// Title: Required. The title of the card header. The header has a fixed
	// height: if both a title and subtitle are specified, each takes up one
	// line. If only the title is specified, it takes up both lines.
	Title string `json:"title,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ImageAltText") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ImageAltText") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CardHeader) MarshalJSON() ([]byte, error) {
	type NoMethod CardHeader
	raw := NoMethod(*s)
	return marshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Column: A column. Google Chat apps
// (https://developers.google.com/chat):
type Column struct {
	// HorizontalAlignment: Specifies whether widgets align to the left,
	// right, or center of a column.
	//
	// Possible values:
	//   "HORIZONTAL_ALIGNMENT_UNSPECIFIED" - Don't use. Unspecified.
	//   "START" - Default value. Aligns widgets to the start position of
	// the column. For left-to-right layouts, aligns to the left. For
	// right-to-left layouts, aligns to the right.
	//   "CENTER" - Aligns widgets to the center of the column.
	//   "END" - Aligns widgets to the end position of the column. For
	// left-to-right layouts, aligns widgets to the right. For right-to-left
	// layouts, aligns widgets to the left.
	HorizontalAlignment string `json:"horizontalAlignment,omitempty"`

	// HorizontalSizeStyle: Specifies how a column fills the width of the
	// card. Google Chat apps (https://developers.google.com/chat):
	//
	// Possible values:
	//   "HORIZONTAL_SIZE_STYLE_UNSPECIFIED" - Don't use. Unspecified.
	//   "FILL_AVAILABLE_SPACE" - Default value. Column fills the available
	// space, up to 70% of the card's width. If both columns are set to
	// `FILL_AVAILABLE_SPACE`, each column fills 50% of the space.
	//   "FILL_MINIMUM_SPACE" - Column fills the least amount of space
	// possible and no more than 30% of the card's width.
	HorizontalSizeStyle string `json:"horizontalSizeStyle,omitempty"`

	// VerticalAlignment: Specifies whether widgets align to the top,
	// bottom, or center of a column. Google Chat apps
	// (https://developers.google.com/chat):
	//
	// Possible values:
	//   "VERTICAL_ALIGNMENT_UNSPECIFIED" - Don't use. Unspecified.
	//   "CENTER" - Default value. Aligns widgets to the center of a column.
	//   "TOP" - Aligns widgets to the top of a column.
	//   "BOTTOM" - Aligns widgets to the bottom of a column.
	VerticalAlignment string `json:"verticalAlignment,omitempty"`

	// Widgets: An array of widgets included in a column. Widgets appear in
	// the order that they are specified.
	Widgets []*Widgets `json:"widgets,omitempty"`

	// ForceSendFields is a list of field names (e.g. "HorizontalAlignment")
	// to unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "HorizontalAlignment") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *Column) MarshalJSON() ([]byte, error) {
	type NoMethod Column
	raw := NoMethod(*s)
	return marshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Columns: The `Columns` widget displays up to 2
// columns in a card message or dialog. You can add widgets to each
// column; the widgets appear in the order that they are specified. For
// an example in Google Chat apps, see Columns
// (https://developers.google.com/chat/ui/widgets/columns). The height
// of each column is determined by the taller column. For example, if
// the first column is taller than the second column, both columns have
// the height of the first column. Because each column can contain a
// different number of widgets, you can't define rows or align widgets
// between the columns. Columns are displayed side-by-side. You can
// customize the width of each column using the `HorizontalSizeStyle`
// field. If the user's screen width is too narrow, the second column
// wraps below the first: * On web, the second column wraps if the
// screen width is less than or equal to 480 pixels. * On iOS devices,
// the second column wraps if the screen width is less than or equal to
// 300 pt. * On Android devices, the second column wraps if the screen
// width is less than or equal to 320 dp. To include more than 2
// columns, or to use rows, use the `Grid` widget. Google Chat apps
// (https://developers.google.com/chat):
type Columns struct {
	// ColumnItems: An array of columns. You can include up to 2 columns in
	// a card or dialog.
	ColumnItems []*Column `json:"columnItems,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ColumnItems") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ColumnItems") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Columns) MarshalJSON() ([]byte, error) {
	type NoMethod Columns
	raw := NoMethod(*s)
	return marshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DateTimePicker: Lets users input a date, a time, or
// both a date and a time. For an example in Google Chat apps, see Date
// time picker
// (https://developers.google.com/chat/ui/widgets/date-time-picker).
// Users can input text or use the picker to select dates and times. If
// users input an invalid date or time, the picker shows an error that
// prompts users to input the information correctly. Google Workspace
// Add-ons and Chat apps
// (https://developers.google.com/workspace/extend):
type DateTimePicker struct {
	// Label: The text that prompts users to input a date, a time, or a date
	// and time. For example, if users are scheduling an appointment, use a
	// label such as `Appointment date` or `Appointment date and time`.
	Label string `json:"label,omitempty"`

	// Name: The name by which the `DateTimePicker` is identified in a form
	// input event. For details about working with form inputs, see Receive
	// form data (https://developers.google.com/chat/ui/read-form-data).
	Name string `json:"name,omitempty"`

	// OnChangeAction: Triggered when the user clicks **Save** or **Clear**
	// from the `DateTimePicker` interface.
	OnChangeAction *Action `json:"onChangeAction,omitempty"`

	// TimezoneOffsetDate: The number representing the time zone offset from
	// UTC, in minutes. If set, the `value_ms_epoch` is displayed in the
	// specified time zone. If unset, the value defaults to the user's time
	// zone setting.
	TimezoneOffsetDate int64 `json:"timezoneOffsetDate,omitempty"`

	// Type: Whether the widget supports inputting a date, a time, or the
	// date and time.
	//
	// Possible values:
	//   "DATE_AND_TIME" - Users input a date and time.
	//   "DATE_ONLY" - Users input a date.
	//   "TIME_ONLY" - Users input a time.
	Type string `json:"type,omitempty"`

	// ValueMsEpoch: The default value displayed in the widget, in
	// milliseconds since Unix epoch time
	// (https://en.wikipedia.org/wiki/Unix_time). Specify the value based on
	// the type of picker (`DateTimePickerType`): * `DATE_AND_TIME`: a
	// calendar date and time in UTC. For example, to represent January 1,
	// 2023 at 12:00 PM UTC, use `1672574400000`. * `DATE_ONLY`: a calendar
	// date at 00:00:00 UTC. For example, to represent January 1, 2023, use
	// `1672531200000`. * `TIME_ONLY`: a time in UTC. For example, to
	// represent 12:00 PM, use `43200000` (or `12 * 60 * 60 * 1000`).
	ValueMsEpoch int64 `json:"valueMsEpoch,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "Label") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Label") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *DateTimePicker) MarshalJSON() ([]byte, error) {
	type NoMethod DateTimePicker
	raw := NoMethod(*s)
	return marshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DecoratedText: A widget that displays text with
// optional decorations such as a label above or below the text, an icon
// in front of the text, a selection widget, or a button after the text.
// For an example in Google Chat apps, see Decorated text
// (https://developers.google.com/chat/ui/widgets/decorated-text).
// Google Workspace Add-ons and Chat apps
// (https://developers.google.com/workspace/extend):
type DecoratedText struct {
	// BottomLabel: The text that appears below `text`. Always wraps.
	BottomLabel string `json:"bottomLabel,omitempty"`

	// Button: A button that a user can click to trigger an action.
	Button *Button `json:"button,omitempty"`

	// EndIcon: An icon displayed after the text. Supports built-in
	// (https://developers.google.com/chat/format-messages#builtinicons) and
	// custom
	// (https://developers.google.com/chat/format-messages#customicons)
	// icons.
	EndIcon *Icon `json:"endIcon,omitempty"`

	// Icon: Deprecated in favor of `startIcon`.
	Icon *Icon `json:"icon,omitempty"`

	// OnClick: This action is triggered when users click `topLabel` or
	// `bottomLabel`.
	OnClick *OnClick `json:"onClick,omitempty"`

	// StartIcon: The icon displayed in front of the text.
	StartIcon *Icon `json:"startIcon,omitempty"`

	// SwitchControl: A switch widget that a user can click to change its
	// state and trigger an action.
	SwitchControl *SwitchControl `json:"switchControl,omitempty"`

	// Text: Required. The primary text. Supports simple formatting. For
	// more information about formatting text, see Formatting text in Google
	// Chat apps
	// (https://developers.google.com/chat/format-messages#card-formatting)
	// and Formatting text in Google Workspace Add-ons
	// (https://developers.google.com/apps-script/add-ons/concepts/widgets#text_formatting).
	Text string `json:"text,omitempty"`

	// TopLabel: The text that appears above `text`. Always truncates.
	TopLabel string `json:"topLabel,omitempty"`

	// WrapText: The wrap text setting. If `true`, the text wraps and
	// displays on multiple lines. Otherwise, the text is truncated. Only
	// applies to `text`, not `topLabel` and `bottomLabel`.
	WrapText bool `json:"wrapText,omitempty"`

	// ForceSendFields is a list of field names (e.g. "BottomLabel") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "BottomLabel") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *DecoratedText) MarshalJSON() ([]byte, error) {
	type NoMethod DecoratedText
	raw := NoMethod(*s)
	return marshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Divider: Displays a divider between widgets as a
// horizontal line. For an example in Google Chat apps, see Divider
// (https://developers.google.com/chat/ui/widgets/divider). Google
// Workspace Add-ons and Chat apps
// (https://developers.google.com/workspace/extend): For example, the
// following JSON creates a divider: ``` "divider": {} ```
type Divider struct {
}

// Grid: Displays a grid with a collection of items.
// Items can only include text or images. For responsive columns, or to
// include more than text or images, use `Columns`. For an example in
// Google Chat apps, see Grid
// (https://developers.google.com/chat/ui/widgets/grid). A grid supports
// any number of columns and items. The number of rows is determined by
// items divided by columns. A grid with 10 items and 2 columns has 5
// rows. A grid with 11 items and 2 columns has 6 rows. Google Workspace
// Add-ons and Chat apps
// (https://developers.google.com/workspace/extend): For example, the
// following JSON creates a 2 column grid with a single item: ```
// "grid": { "title": "A fine collection of items", "columnCount": 2,
// "borderStyle": { "type": "STROKE", "cornerRadius": 4 }, "items": [ {
// "image": { "imageUri": "https://www.example.com/image.png",
// "cropStyle": { "type": "SQUARE" }, "borderStyle": { "type": "STROKE"
// } }, "title": "An item", "textAlignment": "CENTER" } ], "onClick": {
// "openLink": { "url": "https://www.example.com" } } } ```
type Grid struct {
	// BorderStyle: The border style to apply to each grid item.
	BorderStyle *BorderStyle `json:"borderStyle,omitempty"`

	// ColumnCount: The number of columns to display in the grid. A default
	// value is used if this field isn't specified, and that default value
	// is different depending on where the grid is shown (dialog versus
	// companion).
	ColumnCount int64 `json:"columnCount,omitempty"`

	// Items: The items to display in the grid.
	Items []*GridItem `json:"items,omitempty"`

	// OnClick: This callback is reused by each individual grid item, but
	// with the item's identifier and index in the items list added to the
	// callback's parameters.
	OnClick *OnClick `json:"onClick,omitempty"`

	// Title: The text that displays in the grid header.
	Title string `json:"title,omitempty"`

	// ForceSendFields is a list of field names (e.g. "BorderStyle") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "BorderStyle") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Grid) MarshalJSON() ([]byte, error) {
	type NoMethod Grid
	raw := NoMethod(*s)
	return marshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GridItem: Represents an item in a grid layout. Items
// can contain text, an image, or both text and an image. Google
// Workspace Add-ons and Chat apps
// (https://developers.google.com/workspace/extend):
type GridItem struct {
	// Id: A user-specified identifier for this grid item. This identifier
	// is returned in the parent grid's `onClick` callback parameters.
	Id string `json:"id,omitempty"`

	// Image: The image that displays in the grid item.
	Image *ImageComponent `json:"image,omitempty"`

	// Layout: The layout to use for the grid item.
	//
	// Possible values:
	//   "GRID_ITEM_LAYOUT_UNSPECIFIED" - Don't use. Unspecified.
	//   "TEXT_BELOW" - The title and subtitle are shown below the grid
	// item's image.
	//   "TEXT_ABOVE" - The title and subtitle are shown above the grid
	// item's image.
	Layout string `json:"layout,omitempty"`

	// Subtitle: The grid item's subtitle.
	Subtitle string `json:"subtitle,omitempty"`

	// Layout: Specifies how the text is horizontally aligned.
	//
	// Possible values:
	//   "START" - Default value. Aligns widgets to the start position of
	// the column. For left-to-right layouts, aligns to the left. For
	// right-to-left layouts, aligns to the right.
	//   "CENTER" - Aligns widgets to the center of the column.
	//   "END" - Aligns widgets to the end position of the column. For
	// left-to-right layouts, aligns widgets to the right. For right-to-left
	// layouts, aligns widgets to the left.
	TextAlignment string `json:"textAlignment,omitempty"`

	// Title: The grid item's title.
	Title string `json:"title,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Id") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Id") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func GridItemNew(id string, image *ImageComponent, layout string, subtitle string, textAlignment string, title string, forceSendFields []string, nullFields []string) *GridItem {
	return &GridItem{
		Id:              id,
		Image:           image,
		Layout:          layout,
		Subtitle:        subtitle,
		TextAlignment:   textAlignment,
		Title:           title,
		ForceSendFields: forceSendFields,
		NullFields:      nullFields,
	}
}

func (s *GridItem) MarshalJSON() ([]byte, error) {
	type NoMethod GridItem
	raw := NoMethod(*s)
	return marshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Icon: An icon displayed in a widget on a card. For an
// example in Google Chat apps, see Icon
// (https://developers.google.com/chat/ui/widgets/icon). Supports
// built-in
// (https://developers.google.com/chat/format-messages#builtinicons) and
// custom
// (https://developers.google.com/chat/format-messages#customicons)
// icons. Google Workspace Add-ons and Chat apps
// (https://developers.google.com/workspace/extend):
type Icon struct {
	// AltText: Optional. A description of the icon used for accessibility.
	// If unspecified, the default value `Button` is provided. As a best
	// practice, you should set a helpful description for what the icon
	// displays, and if applicable, what it does. For example, `A user's
	// account portrait`, or `Opens a new browser tab and navigates to the
	// Google Chat developer documentation at
	// https://developers.google.com/chat`. If the icon is set in a
	// `Button`, the `altText` appears as helper text when the user hovers
	// over the button. However, if the button also sets `text`, the icon's
	// `altText` is ignored.
	AltText string `json:"altText,omitempty"`

	// IconUrl: Display a custom icon hosted at an HTTPS URL. For example:
	// ``` "iconUrl":
	// "https://developers.google.com/chat/images/quickstart-app-avatar.png"
	// ``` Supported file types include `.png` and `.jpg`.
	IconUrl string `json:"iconUrl,omitempty"`

	// ImageType: The crop style applied to the image. In some cases,
	// applying a `CIRCLE` crop causes the image to be drawn larger than a
	// built-in icon.
	//
	// Possible values:
	//   "SQUARE" - Default value. Applies a square mask to the image. For
	// example, a 4x3 image becomes 3x3.
	//   "CIRCLE" - Applies a circular mask to the image. For example, a 4x3
	// image becomes a circle with a diameter of 3.
	ImageType string `json:"imageType,omitempty"`

	// KnownIcon: Display one of the built-in icons provided by Google
	// Workspace. For example, to display an airplane icon, specify
	// `AIRPLANE`. For a bus, specify `BUS`. For a full list of supported
	// icons, see built-in icons
	// (https://developers.google.com/chat/format-messages#builtinicons).
	KnownIcon string `json:"knownIcon,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AltText") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AltText") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Icon) MarshalJSON() ([]byte, error) {
	type NoMethod Icon
	raw := NoMethod(*s)
	return marshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Image: An image that is specified by a URL and can
// have an `onClick` action. For an example, see Image
// (https://developers.google.com/chat/ui/widgets/image). Google
// Workspace Add-ons and Chat apps
// (https://developers.google.com/workspace/extend):
type Image struct {
	// AltText: The alternative text of this image that's used for
	// accessibility.
	AltText string `json:"altText,omitempty"`

	// ImageUrl: The HTTPS URL that hosts the image. For example: ```
	// https://developers.google.com/chat/images/quickstart-app-avatar.png
	// ```
	ImageUrl string `json:"imageUrl,omitempty"`

	// OnClick: When a user clicks the image, the click triggers this
	// action.
	OnClick *OnClick `json:"onClick,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AltText") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AltText") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func ImageNew(altText string, imageURL string, onClick *OnClick, forceSendFields []string, nullFields []string) *Image {
	return &Image{
		AltText:         altText,
		ImageUrl:        imageURL,
		OnClick:         onClick,
		ForceSendFields: forceSendFields,
		NullFields:      nullFields,
	}
}

func (s *Image) MarshalJSON() ([]byte, error) {
	type NoMethod Image
	raw := NoMethod(*s)
	return marshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ImageComponent: Represents an image. Google Workspace
// Add-ons and Chat apps
// (https://developers.google.com/workspace/extend):
type ImageComponent struct {
	// AltText: The accessibility label for the image.
	AltText string `json:"altText,omitempty"`

	// BorderStyle: The border style to apply to the image.
	BorderStyle *BorderStyle `json:"borderStyle,omitempty"`

	// CropStyle: The crop style to apply to the image.
	CropStyle *ImageCropStyle `json:"cropStyle,omitempty"`

	// ImageUri: The image URL.
	ImageUri string `json:"imageUri,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AltText") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AltText") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ImageComponent) MarshalJSON() ([]byte, error) {
	type NoMethod ImageComponent
	raw := NoMethod(*s)
	return marshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ImageCropStyle: Represents the crop style applied to
// an image. Google Workspace Add-ons and Chat apps
// (https://developers.google.com/workspace/extend): For example, here's
// how to apply a 16:9 aspect ratio: ``` cropStyle { "type":
// "RECTANGLE_CUSTOM", "aspectRatio": 16/9 } ```
type ImageCropStyle struct {
	// AspectRatio: The aspect ratio to use if the crop type is
	// `RECTANGLE_CUSTOM`. For example, here's how to apply a 16:9 aspect
	// ratio: ``` cropStyle { "type": "RECTANGLE_CUSTOM", "aspectRatio":
	// 16/9 } ```
	AspectRatio float64 `json:"aspectRatio,omitempty"`

	// Type: The crop type.
	//
	// Possible values:
	//   "IMAGE_CROP_TYPE_UNSPECIFIED" - Don't use. Unspecified.
	//   "SQUARE" - Default value. Applies a square crop.
	//   "CIRCLE" - Applies a circular crop.
	//   "RECTANGLE_CUSTOM" - Applies a rectangular crop with a custom
	// aspect ratio. Set the custom aspect ratio with `aspectRatio`.
	//   "RECTANGLE_4_3" - Applies a rectangular crop with a 4:3 aspect
	// ratio.
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AspectRatio") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AspectRatio") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ImageCropStyle) MarshalJSON() ([]byte, error) {
	type NoMethod ImageCropStyle
	raw := NoMethod(*s)
	return marshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *ImageCropStyle) UnmarshalJSON(data []byte) error {
	type NoMethod ImageCropStyle
	var s1 struct {
		AspectRatio jsonFloat64 `json:"aspectRatio"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.AspectRatio = float64(s1.AspectRatio)
	return nil
}

// OnClick: Represents how to respond when users click
// an interactive element on a card, such as a button. Google Workspace
// Add-ons and Chat apps
// (https://developers.google.com/workspace/extend):
type OnClick struct {
	// Action: If specified, an action is triggered by this `onClick`.
	Action *Action `json:"action,omitempty"`

	// Card: A new card is pushed to the card stack after clicking if
	// specified. Google Workspace Add-ons
	// (https://developers.google.com/workspace/add-ons):
	Card *Card `json:"card,omitempty"`

	// OpenDynamicLinkAction: An add-on triggers this action when the action
	// needs to open a link. This differs from the `open_link` above in that
	// this needs to talk to server to get the link. Thus some preparation
	// work is required for web client to do before the open link action
	// response comes back. Google Workspace Add-ons
	// (https://developers.google.com/workspace/add-ons):
	OpenDynamicLinkAction *Action `json:"openDynamicLinkAction,omitempty"`

	// OpenLink: If specified, this `onClick` triggers an open link action.
	OpenLink *OpenLink `json:"openLink,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Action") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Action") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *OnClick) MarshalJSON() ([]byte, error) {
	type NoMethod OnClick
	raw := NoMethod(*s)
	return marshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// OpenLink: Represents an `onClick` event that opens a
// hyperlink. Google Workspace Add-ons and Chat apps
// (https://developers.google.com/workspace/extend):
type OpenLink struct {
	// OnClose: Whether the client forgets about a link after opening it, or
	// observes it until the window closes. Google Workspace Add-ons
	// (https://developers.google.com/workspace/add-ons):
	//
	// Possible values:
	//   "NOTHING" - Default value. The card doesn't reload; nothing
	// happens.
	//   "RELOAD" - Reloads the card after the child window closes. If used
	// in conjunction with
	// [`OpenAs.OVERLAY`](https://developers.google.com/workspace/add-ons/ref
	// erence/rpc/google.apps.card.v1#openas), the child window acts as a
	// modal dialog and the parent card is blocked until the child window
	// closes.
	OnClose string `json:"onClose,omitempty"`

	// OpenAs: How to open a link. Google Workspace Add-ons
	// (https://developers.google.com/workspace/add-ons):
	//
	// Possible values:
	//   "FULL_SIZE" - The link opens as a full-size window (if that's the
	// frame used by the client).
	//   "OVERLAY" - The link opens as an overlay, such as a pop-up.
	OpenAs string `json:"openAs,omitempty"`

	// Url: The URL to open.
	Url string `json:"url,omitempty"`

	// ForceSendFields is a list of field names (e.g. "OnClose") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "OnClose") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *OpenLink) MarshalJSON() ([]byte, error) {
	type NoMethod OpenLink
	raw := NoMethod(*s)
	return marshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// PlatformDataSource: For a `SelectionInput` widget
// that uses a multiselect menu, a data source from Google Workspace.
// Used to populate items in a multiselect menu. Google Chat apps
// (https://developers.google.com/chat):
type PlatformDataSource struct {
	// CommonDataSource: A data source shared by all Google Workspace
	// applications, such as users in a Google Workspace organization.
	//
	// Possible values:
	//   "UNKNOWN" - Default value. Don't use.
	//   "USER" - Google Workspace users. The user can only view and select
	// users from their Google Workspace organization.
	CommonDataSource string `json:"commonDataSource,omitempty"`

	// HostAppDataSource: A data source that's unique to a Google Workspace
	// host application, such spaces in Google Chat.
	HostAppDataSource *HostAppDataSourceMarkup `json:"hostAppDataSource,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CommonDataSource") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CommonDataSource") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *PlatformDataSource) MarshalJSON() ([]byte, error) {
	type NoMethod PlatformDataSource
	raw := NoMethod(*s)
	return marshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Section: A section contains a collection of widgets
// that are rendered vertically in the order that they're specified.
// Google Workspace Add-ons and Chat apps
// (https://developers.google.com/workspace/extend):
type Section struct {
	// Collapsible: Indicates whether this section is collapsible.
	// Collapsible sections hide some or all widgets, but users can expand
	// the section to reveal the hidden widgets by clicking **Show more**.
	// Users can hide the widgets again by clicking **Show less**. To
	// determine which widgets are hidden, specify
	// `uncollapsibleWidgetsCount`.
	Collapsible bool `json:"collapsible,omitempty"`

	// Header: Text that appears at the top of a section. Supports simple
	// HTML formatted text. For more information about formatting text, see
	// Formatting text in Google Chat apps
	// (https://developers.google.com/chat/format-messages#card-formatting)
	// and Formatting text in Google Workspace Add-ons
	// (https://developers.google.com/apps-script/add-ons/concepts/widgets#text_formatting).
	Header string `json:"header,omitempty"`

	// UncollapsibleWidgetsCount: The number of uncollapsible widgets which
	// remain visible even when a section is collapsed. For example, when a
	// section contains five widgets and the `uncollapsibleWidgetsCount` is
	// set to `2`, the first two widgets are always shown and the last three
	// are collapsed by default. The `uncollapsibleWidgetsCount` is taken
	// into account only when `collapsible` is `true`.
	UncollapsibleWidgetsCount int64 `json:"uncollapsibleWidgetsCount,omitempty"`

	// Widgets: All the widgets in the section. Must contain at least one
	// widget.
	Widgets []*Widget `json:"widgets,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Collapsible") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Collapsible") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Section) MarshalJSON() ([]byte, error) {
	type NoMethod Section
	raw := NoMethod(*s)
	return marshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// SelectionInput: A widget that creates one or more UI
// items that users can select. For example, a dropdown menu or
// checkboxes. You can use this widget to collect data that can be
// predicted or enumerated. For an example in Google Chat apps, see
// Selection input
// (https://developers.google.com/chat/ui/widgets/selection-input). Chat
// apps can process the value of items that users select or input. For
// details about working with form inputs, see Receive form data
// (https://developers.google.com/chat/ui/read-form-data). To collect
// undefined or abstract data from users, use the TextInput widget.
// Google Workspace Add-ons and Chat apps
// (https://developers.google.com/workspace/extend):
type SelectionInput struct {
	// ExternalDataSource: An external data source, such as a relational
	// data base.
	ExternalDataSource *Action `json:"externalDataSource,omitempty"`

	// Items: An array of selectable items. For example, an array of radio
	// buttons or checkboxes. Supports up to 100 items.
	Items []*SelectionItem `json:"items,omitempty"`

	// Label: The text that appears above the selection input field in the
	// user interface. Specify text that helps the user enter the
	// information your app needs. For example, if users are selecting the
	// urgency of a work ticket from a drop-down menu, the label might be
	// "Urgency" or "Select urgency".
	Label string `json:"label,omitempty"`

	// MultiSelectMaxSelectedItems: For multiselect menus, the maximum
	// number of items that a user can select. Minimum value is 1 item. If
	// unspecified, defaults to 3 items.
	MultiSelectMaxSelectedItems int64 `json:"multiSelectMaxSelectedItems,omitempty"`

	// MultiSelectMinQueryLength: For multiselect menus, the number of text
	// characters that a user inputs before the Chat app queries
	// autocomplete and displays suggested items in the menu. If
	// unspecified, defaults to 0 characters for static data sources and 3
	// characters for external data sources.
	MultiSelectMinQueryLength int64 `json:"multiSelectMinQueryLength,omitempty"`

	// Name: The name that identifies the selection input in a form input
	// event. For details about working with form inputs, see Receive form
	// data (https://developers.google.com/chat/ui/read-form-data).
	Name string `json:"name,omitempty"`

	// OnChangeAction: If specified, the form is submitted when the
	// selection changes. If not specified, you must specify a separate
	// button that submits the form. For details about working with form
	// inputs, see Receive form data
	// (https://developers.google.com/chat/ui/read-form-data).
	OnChangeAction *Action `json:"onChangeAction,omitempty"`

	// PlatformDataSource: A data source from Google Workspace.
	PlatformDataSource *PlatformDataSource `json:"platformDataSource,omitempty"`

	// Type: The type of items that are displayed to users in a
	// `SelectionInput` widget. Selection types support different types of
	// interactions. For example, users can select one or more checkboxes,
	// but they can only select one value from a dropdown menu.
	//
	// Possible values:
	//   "CHECK_BOX" - A set of checkboxes. Users can select one or more
	// checkboxes.
	//   "RADIO_BUTTON" - A set of radio buttons. Users can select one radio
	// button.
	//   "SWITCH" - A set of switches. Users can turn on one or more
	// switches.
	//   "DROPDOWN" - A dropdown menu. Users can select one item from the
	// menu.
	//   "MULTI_SELECT" - A multiselect menu for static or dynamic data.
	// From the menu bar, users select one or more items. Users can also
	// input values to populate dynamic data. For example, users can start
	// typing the name of a Google Chat space and the widget autosuggests
	// the space. To populate items for a multiselect menu, you can use one
	// of the following types of data sources: * Static data: Items are
	// specified as `SelectionItem` objects in the widget. Up to 100 items.
	// * Google Workspace data: Items are populated using data from Google
	// Workspace, such as Google Workspace users or Google Chat spaces. *
	// External data: Items are populated from an external data source
	// outside of Google Workspace. For examples of how to implement
	// multiselect menus, see the [`SelectionInput` widget
	// page](https://developers.google.com/chat/ui/widgets/selection-input#mu
	// ltiselect-menu). [Google Chat
	// apps](https://developers.google.com/chat):
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ExternalDataSource")
	// to unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ExternalDataSource") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *SelectionInput) MarshalJSON() ([]byte, error) {
	type NoMethod SelectionInput
	raw := NoMethod(*s)
	return marshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// SelectionItem: An item that users can select in a
// selection input, such as a checkbox or switch. Google Workspace
// Add-ons and Chat apps
// (https://developers.google.com/workspace/extend):
type SelectionItem struct {
	// BottomText: For multiselect menus, a text description or label that's
	// displayed below the item's `text` field.
	BottomText string `json:"bottomText,omitempty"`

	// Selected: Whether the item is selected by default. If the selection
	// input only accepts one value (such as for radio buttons or a dropdown
	// menu), only set this field for one item.
	Selected bool `json:"selected,omitempty"`

	// StartIconUri: For multiselect menus, the URL for the icon displayed
	// next to the item's `text` field. Supports PNG and JPEG files. Must be
	// an `HTTPS` URL. For example,
	// `https://developers.google.com/chat/images/quickstart-app-avatar.png`.
	StartIconUri string `json:"startIconUri,omitempty"`

	// Text: The text that identifies or describes the item to users.
	Text string `json:"text,omitempty"`

	// Value: The value associated with this item. The client should use
	// this as a form input value. For details about working with form
	// inputs, see Receive form data
	// (https://developers.google.com/chat/ui/read-form-data).
	Value string `json:"value,omitempty"`

	// ForceSendFields is a list of field names (e.g. "BottomText") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "BottomText") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *SelectionItem) MarshalJSON() ([]byte, error) {
	type NoMethod SelectionItem
	raw := NoMethod(*s)
	return marshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// SuggestionItem: One suggested value that users can
// enter in a text input field. Google Workspace Add-ons and Chat apps
// (https://developers.google.com/workspace/extend):
type SuggestionItem struct {
	// Text: The value of a suggested input to a text input field. This is
	// equivalent to what users enter themselves.
	Text string `json:"text,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Text") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Text") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *SuggestionItem) MarshalJSON() ([]byte, error) {
	type NoMethod SuggestionItem
	raw := NoMethod(*s)
	return marshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Suggestions: Suggested values that users can enter.
// These values appear when users click inside the text input field. As
// users type, the suggested values dynamically filter to match what the
// users have typed. For example, a text input field for programming
// language might suggest Java, JavaScript, Python, and C++. When users
// start typing `Jav`, the list of suggestions filters to show `Java`
// and `JavaScript`. Suggested values help guide users to enter values
// that your app can make sense of. When referring to JavaScript, some
// users might enter `javascript` and others `java script`. Suggesting
// `JavaScript` can standardize how users interact with your app. When
// specified, `TextInput.type` is always `SINGLE_LINE`, even if it's set
// to `MULTIPLE_LINE`. Google Workspace Add-ons and Chat apps
// (https://developers.google.com/workspace/extend):
type Suggestions struct {
	// Items: A list of suggestions used for autocomplete recommendations in
	// text input fields.
	Items []*SuggestionItem `json:"items,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Items") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Items") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Suggestions) MarshalJSON() ([]byte, error) {
	type NoMethod Suggestions
	raw := NoMethod(*s)
	return marshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// SwitchControl: Either a toggle-style switch or a
// checkbox inside a `decoratedText` widget. Google Workspace Add-ons
// and Chat apps (https://developers.google.com/workspace/extend): Only
// supported in the `decoratedText` widget.
type SwitchControl struct {
	// ControlType: How the switch appears in the user interface. Google
	// Workspace Add-ons and Chat apps
	// (https://developers.google.com/workspace/extend):
	//
	// Possible values:
	//   "SWITCH" - A toggle-style switch.
	//   "CHECKBOX" - Deprecated in favor of `CHECK_BOX`.
	//   "CHECK_BOX" - A checkbox.
	ControlType string `json:"controlType,omitempty"`

	// Name: The name by which the switch widget is identified in a form
	// input event. For details about working with form inputs, see Receive
	// form data (https://developers.google.com/chat/ui/read-form-data).
	Name string `json:"name,omitempty"`

	// OnChangeAction: The action to perform when the switch state is
	// changed, such as what function to run.
	OnChangeAction *Action `json:"onChangeAction,omitempty"`

	// Selected: When `true`, the switch is selected.
	Selected bool `json:"selected,omitempty"`

	// Value: The value entered by a user, returned as part of a form input
	// event. For details about working with form inputs, see Receive form
	// data (https://developers.google.com/chat/ui/read-form-data).
	Value string `json:"value,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ControlType") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ControlType") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *SwitchControl) MarshalJSON() ([]byte, error) {
	type NoMethod SwitchControl
	raw := NoMethod(*s)
	return marshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TextInput: A field in which users can enter text.
// Supports suggestions and on-change actions. For an example in Google
// Chat apps, see Text input
// (https://developers.google.com/chat/ui/widgets/text-input). Chat apps
// receive and can process the value of entered text during form input
// events. For details about working with form inputs, see Receive form
// data (https://developers.google.com/chat/ui/read-form-data). When you
// need to collect undefined or abstract data from users, use a text
// input. To collect defined or enumerated data from users, use the
// SelectionInput widget. Google Workspace Add-ons and Chat apps
// (https://developers.google.com/workspace/extend):
type TextInput struct {
	// AutoCompleteAction: Optional. Specify what action to take when the
	// text input field provides suggestions to users who interact with it.
	// If unspecified, the suggestions are set by `initialSuggestions` and
	// are processed by the client. If specified, the app takes the action
	// specified here, such as running a custom function. Google Workspace
	// Add-ons (https://developers.google.com/workspace/add-ons):
	AutoCompleteAction *Action `json:"autoCompleteAction,omitempty"`

	// HintText: Text that appears below the text input field meant to
	// assist users by prompting them to enter a certain value. This text is
	// always visible. Required if `label` is unspecified. Otherwise,
	// optional.
	HintText string `json:"hintText,omitempty"`

	// InitialSuggestions: Suggested values that users can enter. These
	// values appear when users click inside the text input field. As users
	// type, the suggested values dynamically filter to match what the users
	// have typed. For example, a text input field for programming language
	// might suggest Java, JavaScript, Python, and C++. When users start
	// typing `Jav`, the list of suggestions filters to show just `Java` and
	// `JavaScript`. Suggested values help guide users to enter values that
	// your app can make sense of. When referring to JavaScript, some users
	// might enter `javascript` and others `java script`. Suggesting
	// `JavaScript` can standardize how users interact with your app. When
	// specified, `TextInput.type` is always `SINGLE_LINE`, even if it's set
	// to `MULTIPLE_LINE`. Google Workspace Add-ons and Chat apps
	// (https://developers.google.com/workspace/extend):
	InitialSuggestions *Suggestions `json:"initialSuggestions,omitempty"`

	// Label: The text that appears above the text input field in the user
	// interface. Specify text that helps the user enter the information
	// your app needs. For example, if you are asking someone's name, but
	// specifically need their surname, write `surname` instead of `name`.
	// Required if `hintText` is unspecified. Otherwise, optional.
	Label string `json:"label,omitempty"`

	// Name: The name by which the text input is identified in a form input
	// event. For details about working with form inputs, see Receive form
	// data (https://developers.google.com/chat/ui/read-form-data).
	Name string `json:"name,omitempty"`

	// OnChangeAction: What to do when a change occurs in the text input
	// field. For example, a user adding to the field or deleting text.
	// Examples of actions to take include running a custom function or
	// opening a dialog (https://developers.google.com/chat/how-tos/dialogs)
	// in Google Chat.
	OnChangeAction *Action `json:"onChangeAction,omitempty"`

	// PlaceholderText: Text that appears in the text input field when the
	// field is empty. Use this text to prompt users to enter a value. For
	// example, `Enter a number from 0 to 100`. Google Chat apps
	// (https://developers.google.com/chat):
	PlaceholderText string `json:"placeholderText,omitempty"`

	// Type: How a text input field appears in the user interface. For
	// example, whether the field is single or multi-line.
	//
	// Possible values:
	//   "SINGLE_LINE" - The text input field has a fixed height of one
	// line.
	//   "MULTIPLE_LINE" - The text input field has a fixed height of
	// multiple lines.
	Type string `json:"type,omitempty"`

	// Value: The value entered by a user, returned as part of a form input
	// event. For details about working with form inputs, see Receive form
	// data (https://developers.google.com/chat/ui/read-form-data).
	Value string `json:"value,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AutoCompleteAction")
	// to unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AutoCompleteAction") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *TextInput) MarshalJSON() ([]byte, error) {
	type NoMethod TextInput
	raw := NoMethod(*s)
	return marshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TextParagraph: A paragraph of text that supports
// formatting. For an example in Google Chat apps, see Text paragraph
// (https://developers.google.com/chat/ui/widgets/text-paragraph). For
// more information about formatting text, see Formatting text in Google
// Chat apps
// (https://developers.google.com/chat/format-messages#card-formatting)
// and Formatting text in Google Workspace Add-ons
// (https://developers.google.com/apps-script/add-ons/concepts/widgets#text_formatting).
// Google Workspace Add-ons and Chat apps
// (https://developers.google.com/workspace/extend):
type TextParagraph struct {
	// Text: The text that's shown in the widget.
	Text string `json:"text,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Text") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Text") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *TextParagraph) MarshalJSON() ([]byte, error) {
	type NoMethod TextParagraph
	raw := NoMethod(*s)
	return marshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Widget: Each card is made up of widgets. A widget is
// a composite object that can represent one of text, images, buttons,
// and other object types.
type Widget struct {
	// ButtonList: A list of buttons. For example, the following JSON
	// creates two buttons. The first is a blue text button and the second
	// is an image button that opens a link: ``` "buttonList": { "buttons":
	// [ { "text": "Edit", "color": { "red": 0, "green": 0, "blue": 1,
	// "alpha": 1 }, "disabled": true, }, { "icon": { "knownIcon": "INVITE",
	// "altText": "check calendar" }, "onClick": { "openLink": { "url":
	// "https://example.com/calendar" } } } ] } ```
	ButtonList *ButtonList `json:"buttonList,omitempty"`

	// Columns: Displays up to 2 columns. To include more than 2 columns, or
	// to use rows, use the `Grid` widget. For example, the following JSON
	// creates 2 columns that each contain text paragraphs: ``` "columns": {
	// "columnItems": [ { "horizontalSizeStyle": "FILL_AVAILABLE_SPACE",
	// "horizontalAlignment": "CENTER", "verticalAlignment": "CENTER",
	// "widgets": [ { "textParagraph": { "text": "First column text
	// paragraph" } } ] }, { "horizontalSizeStyle": "FILL_AVAILABLE_SPACE",
	// "horizontalAlignment": "CENTER", "verticalAlignment": "CENTER",
	// "widgets": [ { "textParagraph": { "text": "Second column text
	// paragraph" } } ] } ] } ```
	Columns *Columns `json:"columns,omitempty"`

	// DateTimePicker: Displays a widget that lets users input a date, time,
	// or date and time. For example, the following JSON creates a date time
	// picker to schedule an appointment: ``` "dateTimePicker": { "name":
	// "appointment_time", "label": "Book your appointment at:", "type":
	// "DATE_AND_TIME", "valueMsEpoch": "796435200000" } ```
	DateTimePicker *DateTimePicker `json:"dateTimePicker,omitempty"`

	// DecoratedText: Displays a decorated text item. For example, the
	// following JSON creates a decorated text widget showing email address:
	// ``` "decoratedText": { "icon": { "knownIcon": "EMAIL" }, "topLabel":
	// "Email Address", "text": "sasha@example.com", "bottomLabel": "This is
	// a new Email address!", "switchControl": { "name":
	// "has_send_welcome_email_to_sasha", "selected": false, "controlType":
	// "CHECKBOX" } } ```
	DecoratedText *DecoratedText `json:"decoratedText,omitempty"`

	// Divider: Displays a horizontal line divider between widgets. For
	// example, the following JSON creates a divider: ``` "divider": { } ```
	Divider *Divider `json:"divider,omitempty"`

	// Grid: Displays a grid with a collection of items. A grid supports any
	// number of columns and items. The number of rows is determined by the
	// upper bounds of the number items divided by the number of columns. A
	// grid with 10 items and 2 columns has 5 rows. A grid with 11 items and
	// 2 columns has 6 rows. Google Workspace Add-ons and Chat apps
	// (https://developers.google.com/workspace/extend): For example, the
	// following JSON creates a 2 column grid with a single item: ```
	// "grid": { "title": "A fine collection of items", "columnCount": 2,
	// "borderStyle": { "type": "STROKE", "cornerRadius": 4 }, "items": [ {
	// "image": { "imageUri": "https://www.example.com/image.png",
	// "cropStyle": { "type": "SQUARE" }, "borderStyle": { "type": "STROKE"
	// } }, "title": "An item", "textAlignment": "CENTER" } ], "onClick": {
	// "openLink": { "url": "https://www.example.com" } } } ```
	Grid *Grid `json:"grid,omitempty"`

	// HorizontalAlignment: Specifies whether widgets align to the left,
	// right, or center of a column.
	//
	// Possible values:
	//   "HORIZONTAL_ALIGNMENT_UNSPECIFIED" - Don't use. Unspecified.
	//   "START" - Default value. Aligns widgets to the start position of
	// the column. For left-to-right layouts, aligns to the left. For
	// right-to-left layouts, aligns to the right.
	//   "CENTER" - Aligns widgets to the center of the column.
	//   "END" - Aligns widgets to the end position of the column. For
	// left-to-right layouts, aligns widgets to the right. For right-to-left
	// layouts, aligns widgets to the left.
	HorizontalAlignment string `json:"horizontalAlignment,omitempty"`

	// Image: Displays an image. For example, the following JSON creates an
	// image with alternative text: ``` "image": { "imageUrl":
	// "https://developers.google.com/chat/images/quickstart-app-avatar.png",
	//  "altText": "Chat app avatar" } ```
	Image *Image `json:"image,omitempty"`

	// SelectionInput: Displays a selection control that lets users select
	// items. Selection controls can be checkboxes, radio buttons, switches,
	// or dropdown menus. For example, the following JSON creates a dropdown
	// menu that lets users choose a size: ``` "selectionInput": { "name":
	// "size", "label": "Size" "type": "DROPDOWN", "items": [ { "text": "S",
	// "value": "small", "selected": false }, { "text": "M", "value":
	// "medium", "selected": true }, { "text": "L", "value": "large",
	// "selected": false }, { "text": "XL", "value": "extra_large",
	// "selected": false } ] } ```
	SelectionInput *SelectionInput `json:"selectionInput,omitempty"`

	// TextInput: Displays a text box that users can type into. For example,
	// the following JSON creates a text input for an email address: ```
	// "textInput": { "name": "mailing_address", "label": "Mailing Address"
	// } ``` As another example, the following JSON creates a text input for
	// a programming language with static suggestions: ``` "textInput": {
	// "name": "preferred_programing_language", "label": "Preferred
	// Language", "initialSuggestions": { "items": [ { "text": "C++" }, {
	// "text": "Java" }, { "text": "JavaScript" }, { "text": "Python" } ] }
	// } ```
	TextInput *TextInput `json:"textInput,omitempty"`

	// TextParagraph: Displays a text paragraph. Supports simple HTML
	// formatted text. For more information about formatting text, see
	// Formatting text in Google Chat apps
	// (https://developers.google.com/chat/format-messages#card-formatting)
	// and Formatting text in Google Workspace Add-ons
	// (https://developers.google.com/apps-script/add-ons/concepts/widgets#text_formatting).
	// For example, the following JSON creates a bolded text: ```
	// "textParagraph": { "text": " *bold text*" } ```
	TextParagraph *TextParagraph `json:"textParagraph,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ButtonList") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ButtonList") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Widget) MarshalJSON() ([]byte, error) {
	type NoMethod Widget
	raw := NoMethod(*s)
	return marshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Widgets: The supported widgets that you can include
// in a column. Google Chat apps (https://developers.google.com/chat):
type Widgets struct {
	// ButtonList: ButtonList widget.
	ButtonList *ButtonList `json:"buttonList,omitempty"`

	// DateTimePicker: DateTimePicker widget.
	DateTimePicker *DateTimePicker `json:"dateTimePicker,omitempty"`

	// DecoratedText: DecoratedText widget.
	DecoratedText *DecoratedText `json:"decoratedText,omitempty"`

	// Image: Image widget.
	Image *Image `json:"image,omitempty"`

	// SelectionInput: SelectionInput widget.
	SelectionInput *SelectionInput `json:"selectionInput,omitempty"`

	// TextInput: TextInput widget.
	TextInput *TextInput `json:"textInput,omitempty"`

	// TextParagraph: TextParagraph widget.
	TextParagraph *TextParagraph `json:"textParagraph,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ButtonList") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ButtonList") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Widgets) MarshalJSON() ([]byte, error) {
	type NoMethod Widgets
	raw := NoMethod(*s)
	return marshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Group: A Google Group in Google Chat.
type Group struct {
	// Name: Resource name for a Google Group. Represents a group
	// (https://cloud.google.com/identity/docs/reference/rest/v1/groups) in
	// Cloud Identity Groups API. Format: groups/{group}
	Name string `json:"name,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Name") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Name") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Group) MarshalJSON() ([]byte, error) {
	type NoMethod Group
	raw := NoMethod(*s)
	return marshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// HostAppDataSourceMarkup: For a `SelectionInput` widget that uses a
// multiselect menu, a data source from a Google Workspace application.
// The data source populates selection items for the multiselect menu.
// Google Chat apps (https://developers.google.com/chat):
type HostAppDataSourceMarkup struct {
	// ChatDataSource: A data source from Google Chat.
	ChatDataSource *ChatClientDataSourceMarkup `json:"chatDataSource,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ChatDataSource") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ChatDataSource") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *HostAppDataSourceMarkup) MarshalJSON() ([]byte, error) {
	type NoMethod HostAppDataSourceMarkup
	raw := NoMethod(*s)
	return marshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type Inputs struct {
	// DateInput: Date input values.
	DateInput *DateInput `json:"dateInput,omitempty"`

	// DateTimeInput: Date and time input values.
	DateTimeInput *DateTimeInput `json:"dateTimeInput,omitempty"`

	// StringInputs: Input parameter for regular widgets. For single-valued
	// widgets, it is a single value list. For multi-valued widgets, such as
	// checkbox, all the values are presented.
	StringInputs *StringInputs `json:"stringInputs,omitempty"`

	// TimeInput: Time input values.
	TimeInput *TimeInput `json:"timeInput,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DateInput") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DateInput") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Inputs) MarshalJSON() ([]byte, error) {
	type NoMethod Inputs
	raw := NoMethod(*s)
	return marshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type DateInput struct {
	// MsSinceEpoch: Time since epoch time, in milliseconds.
	MsSinceEpoch int64 `json:"msSinceEpoch,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "MsSinceEpoch") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "MsSinceEpoch") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *DateInput) MarshalJSON() ([]byte, error) {
	type NoMethod DateInput
	raw := NoMethod(*s)
	return marshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DateTimeInput: Date and time input values.
type DateTimeInput struct {
	// HasDate: Whether the `datetime` input includes a calendar date.
	HasDate bool `json:"hasDate,omitempty"`

	// HasTime: Whether the `datetime` input includes a timestamp.
	HasTime bool `json:"hasTime,omitempty"`

	// MsSinceEpoch: Time since epoch time, in milliseconds.
	MsSinceEpoch int64 `json:"msSinceEpoch,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "HasDate") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "HasDate") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *DateTimeInput) MarshalJSON() ([]byte, error) {
	type NoMethod DateTimeInput
	raw := NoMethod(*s)
	return marshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// StringInputs: Input parameter for regular widgets. For single-valued
// widgets, it is a single value list. For multi-valued widgets, such as
// checkbox, all the values are presented.
type StringInputs struct {
	// Value: An array of strings entered by the user.
	Value []string `json:"value,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Value") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Value") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *StringInputs) MarshalJSON() ([]byte, error) {
	type NoMethod StringInputs
	raw := NoMethod(*s)
	return marshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TimeInput: Time input values.
type TimeInput struct {
	// Hours: The hour on a 24-hour clock.
	Hours int64 `json:"hours,omitempty"`

	// Minutes: The number of minutes past the hour. Valid values are 0 to
	// 59.
	Minutes int64 `json:"minutes,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Hours") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Hours") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *TimeInput) MarshalJSON() ([]byte, error) {
	type NoMethod TimeInput
	raw := NoMethod(*s)
	return marshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type TimeZone struct {
	// Id: The IANA TZ (https://www.iana.org/time-zones) time zone database
	// code, such as "America/Toronto".
	Id string `json:"id,omitempty"`

	// Offset: The user timezone offset, in milliseconds, from Coordinated
	// Universal Time (UTC).
	Offset int64 `json:"offset,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Id") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Id") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *TimeZone) MarshalJSON() ([]byte, error) {
	type NoMethod TimeZone
	raw := NoMethod(*s)
	return marshalJSON(raw, s.ForceSendFields, s.NullFields)
}
