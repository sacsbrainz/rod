// This file is generated by "./lib/proto/generate"

package proto

/*

HeadlessExperimental

This domain provides experimental commands only supported in headless mode.

*/

// HeadlessExperimentalScreenshotParamsFormat enum
type HeadlessExperimentalScreenshotParamsFormat string

const (
	// HeadlessExperimentalScreenshotParamsFormatJpeg enum const
	HeadlessExperimentalScreenshotParamsFormatJpeg HeadlessExperimentalScreenshotParamsFormat = "jpeg"

	// HeadlessExperimentalScreenshotParamsFormatPng enum const
	HeadlessExperimentalScreenshotParamsFormatPng HeadlessExperimentalScreenshotParamsFormat = "png"

	// HeadlessExperimentalScreenshotParamsFormatWebp enum const
	HeadlessExperimentalScreenshotParamsFormatWebp HeadlessExperimentalScreenshotParamsFormat = "webp"
)

// HeadlessExperimentalScreenshotParams Encoding options for a screenshot.
type HeadlessExperimentalScreenshotParams struct {
	// Format (optional) Image compression format (defaults to png).
	Format HeadlessExperimentalScreenshotParamsFormat `json:"format,omitempty"`

	// Quality (optional) Compression quality from range [0..100] (jpeg only).
	Quality *int `json:"quality,omitempty"`

	// OptimizeForSpeed (optional) Optimize image encoding for speed, not for resulting size (defaults to false)
	OptimizeForSpeed bool `json:"optimizeForSpeed,omitempty"`
}

// HeadlessExperimentalBeginFrame Sends a BeginFrame to the target and returns when the frame was completed. Optionally captures a
// screenshot from the resulting frame. Requires that the target was created with enabled
// BeginFrameControl. Designed for use with --run-all-compositor-stages-before-draw, see also
// https://goo.gle/chrome-headless-rendering for more background.
type HeadlessExperimentalBeginFrame struct {
	// FrameTimeTicks (optional) Timestamp of this BeginFrame in Renderer TimeTicks (milliseconds of uptime). If not set,
	// the current time will be used.
	FrameTimeTicks *float64 `json:"frameTimeTicks,omitempty"`

	// Interval (optional) The interval between BeginFrames that is reported to the compositor, in milliseconds.
	// Defaults to a 60 frames/second interval, i.e. about 16.666 milliseconds.
	Interval *float64 `json:"interval,omitempty"`

	// NoDisplayUpdates (optional) Whether updates should not be committed and drawn onto the display. False by default. If
	// true, only side effects of the BeginFrame will be run, such as layout and animations, but
	// any visual updates may not be visible on the display or in screenshots.
	NoDisplayUpdates bool `json:"noDisplayUpdates,omitempty"`

	// Screenshot (optional) If set, a screenshot of the frame will be captured and returned in the response. Otherwise,
	// no screenshot will be captured. Note that capturing a screenshot can fail, for example,
	// during renderer initialization. In such a case, no screenshot data will be returned.
	Screenshot *HeadlessExperimentalScreenshotParams `json:"screenshot,omitempty"`
}

// ProtoReq name
func (m HeadlessExperimentalBeginFrame) ProtoReq() string { return "HeadlessExperimental.beginFrame" }

// Call the request
func (m HeadlessExperimentalBeginFrame) Call(c Client) (*HeadlessExperimentalBeginFrameResult, error) {
	var res HeadlessExperimentalBeginFrameResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// HeadlessExperimentalBeginFrameResult ...
type HeadlessExperimentalBeginFrameResult struct {
	// HasDamage Whether the BeginFrame resulted in damage and, thus, a new frame was committed to the
	// display. Reported for diagnostic uses, may be removed in the future.
	HasDamage bool `json:"hasDamage"`

	// ScreenshotData (optional) Base64-encoded image data of the screenshot, if one was requested and successfully taken.
	ScreenshotData []byte `json:"screenshotData,omitempty"`
}

// HeadlessExperimentalDisable (deprecated) Disables headless events for the target.
type HeadlessExperimentalDisable struct{}

// ProtoReq name
func (m HeadlessExperimentalDisable) ProtoReq() string { return "HeadlessExperimental.disable" }

// Call sends the request
func (m HeadlessExperimentalDisable) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// HeadlessExperimentalEnable (deprecated) Enables headless events for the target.
type HeadlessExperimentalEnable struct{}

// ProtoReq name
func (m HeadlessExperimentalEnable) ProtoReq() string { return "HeadlessExperimental.enable" }

// Call sends the request
func (m HeadlessExperimentalEnable) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}
