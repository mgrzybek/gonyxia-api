package core

type property struct {
	Type string `json:"type,omitempty"`

	Description string `json:"description,omitempty"`

	Title string `json:"title,omitempty"`

	Pattern string `json:"pattern,omitempty"`

	Media *media `json:"media,omitempty"`

	Minimum string `json:"minimum,omitempty"`

	Render string `json:"render,omitempty"`

	SliderMin int32 `json:"sliderMin,omitempty"`

	SliderMax int32 `json:"sliderMax,omitempty"`

	SliderStep int32 `json:"sliderStep,omitempty"`

	SliderUnit string `json:"sliderUnit,omitempty"`

	SliderExtremity string `json:"sliderExtremity,omitempty"`

	SliderExtremitySemantic string `json:"sliderExtremitySemantic,omitempty"`

	SliderRangeID string `json:"sliderRangeId,omitempty"`

	Hidden *hidden `json:"hidden,omitempty"`

	Default *interface{} `json:"default,omitempty"`

	Enum *interface{} `json:"enum,omitempty"`

	XForm *xForm `json:"x-form,omitempty"`

	XGenerated *xGenerated `json:"x-generated,omitempty"`
}
