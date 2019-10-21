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
	// winWidth := 800
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

	// width, height := a.GetSize()
	// Create and add a button to the scene

	l1 := gui.NewLabel("0/50")
	l1.SetPosition(700, float32(winHeight-40))
	// l1.SetFontSize(30)
	scene.Add(l1)

	scroller3 := gui.NewVScroller(680, 80) //gui.NewScroller(500, 500, gui.ScrollVertical, img)
	scroller3.SetPosition(10, float32(winHeight-100))

	// es := gui.EditStyle{}
	// es.Border = gui.RectBounds{0, 0, 0, 0}
	// es.Paddings = gui.RectBounds{0, 0, 0, 0}
	// es.BorderColor = math32.Color4{0, 0, 0, 1}
	// es.BgColor = math32.Color4{0, 0, 0, 0}
	// es.FgColor = math32.Color4{0.85, 0.85, 0.85, 1}
	// ess := gui.EditStyles{
	// 	Normal:   es,
	// 	Over:     es,
	// 	Focus:    es,
	// 	Disabled: es,
	// }

	// Edit 2
	ed2 := gui.NewEdit(680, "edit 2")
	ed2.SetPosition(10, 10)
	// ed2.SetSize(100, 100)
	ed2.SetFontSize(18)
	ed2.MaxLength = 51
	ed2.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		linelength := len(ed2.Text())
		// if linelength > 50{

		// }
		l1.SetText(fmt.Sprintf("%d/50", linelength))

		// fmt.Println("Edit 2 OnChange:%s", ed2.Text())
	})
	// ed2.SetStyles(&ess)
	// scene.Add(ed2)

	ed1 := gui.NewEdit(680, "edit 1")
	// ed1.SetPosition(0, 0)
	ed1.SetFontSize(18)
	ed1.MaxLength = 50
	ed1.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		linelength := len(ed1.Text())
		if linelength > 30 {
			scroller3.Add(ed2)
			fmt.Println("Size should grow")
		}
		l1.SetText(fmt.Sprintf("%d/50", linelength))

		// fmt.Println("Edit 2 OnChange:%s", ed2.Text())
	})
	scroller3.Add(ed1)

	scene.Add(scroller3)

	// img, _ := gui.NewImage("images/maxresdefault.jpg")
	// imgOriginalSize := float32(512)
	// img.SetSize(imgOriginalSize, imgOriginalSize)
	// scroller := gui.NewVScroller(500, 500) //gui.NewScroller(500, 500, gui.ScrollVertical, img)
	// scroller.SetPosition(10, 10)

	// scene.Add(scroller)

	// btn := gui.NewButton("Send")
	// btn.SetPosition(float32(winWidth-50), float32(winHeight-50))
	// btn.SetSize(45, 45)
	// btn.Subscribe(gui.OnClick, func(name string, ev interface{}) {
	// 	scroller.Add(gui.NewLabel(ed2.Text()))

	// })
	// scene.Add(btn)

	// img, _ = gui.NewImage("images/maxresdefault.jpg")
	// imgOriginalSize = float32(512)
	// img.SetSize(imgOriginalSize, imgOriginalSize)
	// scroller1 := gui.NewScroller(240, 500, gui.ScrollVertical, img)
	// scroller1.SetPosition(550, 10)

	// scene.Add(scroller)
	// scene.Add(scroller1)

	a.Gls().ClearColor(0.5, 0.5, 0.5, 1.0)

	// Run the application
	a.Run(func(renderer *renderer.Renderer, deltaTime time.Duration) {
		a.Gls().Clear(gls.DEPTH_BUFFER_BIT | gls.STENCIL_BUFFER_BIT | gls.COLOR_BUFFER_BIT)
		renderer.Render(scene, cam)
	})
}
