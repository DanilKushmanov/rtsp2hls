package main

import (
	"github.com/AlexxIT/go2rtc/internal/api"
	"github.com/AlexxIT/go2rtc/internal/api/ws"
	"github.com/AlexxIT/go2rtc/internal/app"
	"github.com/AlexxIT/go2rtc/internal/debug"
	"github.com/AlexxIT/go2rtc/internal/hls"
	"github.com/AlexxIT/go2rtc/internal/mjpeg"
	"github.com/AlexxIT/go2rtc/internal/mp4"
	"github.com/AlexxIT/go2rtc/internal/ngrok"
	"github.com/AlexxIT/go2rtc/internal/rtsp"
	"github.com/AlexxIT/go2rtc/internal/srtp"
	"github.com/AlexxIT/go2rtc/internal/streams"
	"github.com/AlexxIT/go2rtc/internal/webrtc"
	"github.com/AlexxIT/go2rtc/pkg/shell"
)

func main() {
	app.Version = "1.9.8"

	// 1. Core modules: app, api/ws, streams

	app.Init() // init config and logs

	api.Init() // init API before all others
	ws.Init()  // init WS API endpoint

	streams.Init() // streams module

	// 2. Main sources and servers

	rtsp.Init()   // rtsp source, RTSP server
	webrtc.Init() // webrtc source, WebRTC server

	// 3. Main API

	mp4.Init()   // MP4 API
	hls.Init()   // HLS API
	mjpeg.Init() // MJPEG API

	// 4. Other sources and servers

	// 5. Other sources

	// 6. Helper modules

	ngrok.Init() // ngrok module
	srtp.Init()  // SRTP server
	debug.Init() // debug API

	// 7. Go

	shell.RunUntilSignal()
}
