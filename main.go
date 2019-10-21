package main

import (
	"fmt"
	"time"

	"github.com/g3n/engine/app"
	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/renderer"
	"github.com/g3n/engine/window"
)

func main() {
	winWidth := 800
	winHeight := 600
	// Create application and scene
	a := app.App()
	// Initialize window

	scene := core.NewNode()
	// Set the scene to be managed by the gui manager
	gui.Manager().Set(scene)
	// Create perspective camera
	cam := camera.New(1)
	cam.SetPosition(0, 0, 3)
	scene.Add(cam)
	// Set up orbit control for the camera
	camera.NewOrbitControl(cam)

	// Set up callback to update viewport and camera aspect ratio when the window is resized
	onResize := func(evname string, ev interface{}) {
		// Get framebuffer size and update viewport accordingly
		width, height := a.GetSize()
		a.Gls().Viewport(0, 0, int32(width), int32(height))
		// Update the camera's aspect ratio
		cam.SetAspect(float32(width) / float32(height))
	}
	a.Subscribe(window.OnWindowSize, onResize)
	onResize("", nil)

	inputLengthLable := gui.NewLabel("0/50")
	inputLengthLable.SetPosition(700, float32(winHeight-40))
	scene.Add(inputLengthLable)

	inputTextEdit := gui.NewEdit(680, "edit 2")
	inputTextEdit.SetPosition(10, float32(winHeight-40))
	inputTextEdit.SetFontSize(18)
	inputTextEdit.MaxLength = 50
	inputTextEdit.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		linelength := len(inputTextEdit.Text())
		// if linelength > 50{

		// }
		inputLengthLable.SetText(fmt.Sprintf("%d/50", linelength))
	})
	scene.Add(inputTextEdit)

	converseScroller := gui.NewVScroller(500, 500) //gui.NewScroller(500, 500, gui.ScrollVertical, img)
	converseScroller.SetPosition(10, 10)
	scene.Add(converseScroller)

	sendButton := gui.NewButton("Send")
	sendButton.SetPosition(float32(winWidth-50), float32(winHeight-50))
	sendButton.SetSize(45, 45)
	sendButton.Subscribe(gui.OnClick, func(name string, ev interface{}) {
		converseScroller.Add(gui.NewLabel(inputTextEdit.Text()))
		inputTextEdit.SetText("")
	})
	scene.Add(sendButton)

	scrollerImg, _ := gui.NewImage("images/maxresdefault.jpg")
	imgOriginalSize := float32(512)
	scrollerImg.SetSize(imgOriginalSize, imgOriginalSize)
	playersScroller := gui.NewScroller(240, 500, gui.ScrollVertical, scrollerImg)
	playersScroller.SetPosition(550, 10)
	scene.Add(playersScroller)

	a.Gls().ClearColor(0.5, 0.5, 0.5, 1.0)

	// Run the application
	a.Run(func(renderer *renderer.Renderer, deltaTime time.Duration) {
		a.Gls().Clear(gls.DEPTH_BUFFER_BIT | gls.STENCIL_BUFFER_BIT | gls.COLOR_BUFFER_BIT)
		renderer.Render(scene, cam)
	})
}
