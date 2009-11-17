<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/exp/nacl/av/av.go</title>

  <link rel="stylesheet" type="text/css" href="../../../../../doc/style.css">
  <script type="text/javascript" src="../../../../../doc/godocs.js"></script>

</head>

<body>

  <script>
    // Catch 'enter' key down events and trigger the search form submission.
    function codesearchKeyDown(event) {
      if (event.which == 13) {
        var form = document.getElementById('codesearch');
        var query = document.getElementById('codesearchQuery');
        form.q.value = "lang:go package:go.googlecode.com " + query.value;
        document.getElementById('codesearch').submit();
}      return true;
}
    // Capture the submission event and construct the query parameter.
    function codeSearchSubmit() {
      var query = document.getElementById('codesearchQuery');
      var form = document.getElementById('codesearch');
      form.q.value = "lang:go package:go.googlecode.com " + query.value;
      return true;
}  </script>

<div id="topnav">
  <table summary="">
    <tr>
      <td id="headerImage">
        <a href="../../../../../index.html"><img src="../../../../../doc/logo-153x55.png" height="55" width="153" alt="Go Home Page" style="border:0" /></a>
      </td>
      <td>
        <div id="headerDocSetTitle">The Go Programming Language</div>
      </td>
      <td>
        <!-- <table>
          <tr>
            <td>
              <! The input box is outside of the form because we want to add
              a couple of restricts to the query before submitting. If we just
              add the restricts to the text box before submitting, then they
              appear in the box when the user presses 'back'. Thus we use a
              hidden field in the form. However, there's no way to stop the
              non-hidden text box from also submitting a value unless we move
              it outside of the form
              <input type="search" id="codesearchQuery" value="" size="30" onkeydown="return codesearchKeyDown(event);"/>
              <form method="GET" action="http://www.google.com/codesearch" id="codesearch" class="search" onsubmit="return codeSearchSubmit();" style="display:inline;">
                <input type="hidden" name="q" value=""/>
                <input type="submit" value="Code search" />
                <span style="color: red">(TODO: remove for now?)</span>
              </form>
            </td>
          </tr>
          <tr>
            <td>
              <span style="color: gray;">(e.g. &ldquo;pem&rdquo; or &ldquo;xml&rdquo;)</span>
            </td>
          </tr>
        </table> -->
      </td>
    </tr>
  </table>
</div>

<div id="linkList">
  <ul>
    <li class="navhead"><a href="../../../../../index.html">Home</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Documents</li>
    <li><a href="../../../../../doc/go_tutorial.html">Tutorial</a></li>
    <li><a href="../../../../../doc/effective_go.html">Effective Go</a></li>
    <li><a href="../../../../../doc/go_faq.html">FAQ</a></li>
    <li><a href="../../../../../doc/go_lang_faq.html">Language Design FAQ</a></li>
    <li><a href="http://www.youtube.com/watch?v=rKnDgT73v8s">Tech talk (1 hour)</a> (<a href="../../../../../doc/go_talk-20091030.pdf">PDF</a>)</li>
    <li><a href="../../../../../doc/go_spec.html">Language Specification</a></li>
    <li><a href="../../../../../doc/go_mem.html">Memory Model</a></li>
    <li><a href="../../../../../doc/go_for_cpp_programmers.html">Go for C++ Programmers</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">How To</li>
    <li><a href="../../../../../doc/install.html">Install Go</a></li>
    <li><a href="../../../../../doc/contribute.html">Contribute code</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Programming</li>
    <li><a href="../../../../../cmd/index.html">Command documentation</a></li>
    <li><a href="../../../../../pkg/index.html">Package documentation</a></li>
    <li><a href="../../../../index.html">Source files</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Help</li>
    <li>#go-nuts on irc.freenode.net</li>
    <li><a href="http://groups.google.com/group/golang-nuts">Go Nuts mailing list</a></li>
    <li><a href="http://code.google.com/p/go/issues/list">Issue tracker</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Go code search</li>
    <form method="GET" action="http://golang.org/search" class="search">
    <input type="search" name="q" value="" size="25" style="width:80%; max-width:200px" />
    <input type="submit" value="Go" />
    </form>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Last update</li>
	<li>Thu Nov 12 15:48:37 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/exp/nacl/av/av.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// Native Client audio/video</span>

<a id="L7"></a><span class="comment">// Package av implements audio and video access for Native Client</span>
<a id="L8"></a><span class="comment">// binaries running standalone or embedded in a web browser window.</span>
<a id="L9"></a><span class="comment">//</span>
<a id="L10"></a><span class="comment">// The C version of the API is documented at</span>
<a id="L11"></a><span class="comment">// http://nativeclient.googlecode.com/svn/data/docs_tarball/nacl/googleclient/native_client/scons-out/doc/html/group__audio__video.html</span>
<a id="L12"></a>package av

<a id="L14"></a>import (
    <a id="L15"></a>&#34;bytes&#34;;
    <a id="L16"></a>&#34;exp/draw&#34;;
    <a id="L17"></a>&#34;exp/nacl/srpc&#34;;
    <a id="L18"></a>&#34;log&#34;;
    <a id="L19"></a>&#34;os&#34;;
    <a id="L20"></a>&#34;syscall&#34;;
    <a id="L21"></a>&#34;unsafe&#34;;
<a id="L22"></a>)

<a id="L24"></a>var srpcEnabled = srpc.Enabled()

<a id="L26"></a><span class="comment">// native_client/src/trusted/service_runtime/include/sys/audio_video.h</span>

<a id="L28"></a><span class="comment">// Subsystem values for Init.</span>
<a id="L29"></a>const (
    <a id="L30"></a>SubsystemVideo = 1 &lt;&lt; iota;
    <a id="L31"></a>SubsystemAudio;
    <a id="L32"></a>SubsystemEmbed;
<a id="L33"></a>)
<a id="L34"></a><span class="comment">//	SubsystemRawEvents;</span>

<a id="L36"></a><span class="comment">// Audio formats.</span>
<a id="L37"></a>const (
    <a id="L38"></a>AudioFormatStereo44K = iota;
    <a id="L39"></a>AudioFormatStereo48K;
<a id="L40"></a>)

<a id="L42"></a><span class="comment">// A Window represents a connection to the Native Client window.</span>
<a id="L43"></a><span class="comment">// It implements draw.Context.</span>
<a id="L44"></a>type Window struct {
    <a id="L45"></a>Embedded bool; <span class="comment">// running as part of a web page?</span>
    <a id="L46"></a>*Image;        <span class="comment">// screen image</span>

    <a id="L48"></a>mousec  chan draw.Mouse;
    <a id="L49"></a>kbdc    chan int;
    <a id="L50"></a>quitc   chan bool;
    <a id="L51"></a>resizec chan bool;
<a id="L52"></a>}

<a id="L54"></a><span class="comment">// *Window implements draw.Context</span>
<a id="L55"></a>var _ draw.Context = (*Window)(nil)

<a id="L57"></a>func (w *Window) KeyboardChan() &lt;-chan int { return w.kbdc }

<a id="L59"></a>func (w *Window) MouseChan() &lt;-chan draw.Mouse {
    <a id="L60"></a>return w.mousec
<a id="L61"></a>}

<a id="L63"></a>func (w *Window) QuitChan() &lt;-chan bool { return w.quitc }

<a id="L65"></a>func (w *Window) ResizeChan() &lt;-chan bool { return w.resizec }

<a id="L67"></a>func (w *Window) Screen() draw.Image { return w.Image }

<a id="L69"></a><span class="comment">// Init initializes the Native Client subsystems specified by subsys.</span>
<a id="L70"></a><span class="comment">// Init must be called before using any of the other functions</span>
<a id="L71"></a><span class="comment">// in this package, and it must be called only once.</span>
<a id="L72"></a><span class="comment">//</span>
<a id="L73"></a><span class="comment">// If the SubsystemVideo flag is set, Init requests a window of size dx×dy.</span>
<a id="L74"></a><span class="comment">// When embedded in a web page, the web page&#39;s window specification</span>
<a id="L75"></a><span class="comment">// overrides the parameters to Init, so the returned Window may have</span>
<a id="L76"></a><span class="comment">// a different size than requested.</span>
<a id="L77"></a><span class="comment">//</span>
<a id="L78"></a><span class="comment">// If the SubsystemAudio flag is set, Init requests a connection to the</span>
<a id="L79"></a><span class="comment">// audio device carrying 44 kHz 16-bit stereo PCM audio samples.</span>
<a id="L80"></a>func Init(subsys int, dx, dy int) (*Window, os.Error) {
    <a id="L81"></a>xsubsys := subsys;
    <a id="L82"></a>if srpcEnabled {
        <a id="L83"></a>waitBridge();
        <a id="L84"></a>xsubsys &amp;^= SubsystemVideo | SubsystemEmbed;
    <a id="L85"></a>}

    <a id="L87"></a>if xsubsys&amp;SubsystemEmbed != 0 {
        <a id="L88"></a>return nil, os.NewError(&#34;not embedded&#34;)
    <a id="L89"></a>}

    <a id="L91"></a>w := new(Window);
    <a id="L92"></a>err := multimediaInit(xsubsys);
    <a id="L93"></a>if err != nil {
        <a id="L94"></a>return nil, err
    <a id="L95"></a>}

    <a id="L97"></a>if subsys&amp;SubsystemVideo != 0 {
        <a id="L98"></a>if dx, dy, err = videoInit(dx, dy); err != nil {
            <a id="L99"></a>return nil, err
        <a id="L100"></a>}
        <a id="L101"></a>w.Image = newImage(dx, dy, bridge.pixel);
        <a id="L102"></a>w.resizec = make(chan bool, 64);
        <a id="L103"></a>w.kbdc = make(chan int, 64);
        <a id="L104"></a>w.mousec = make(chan draw.Mouse, 64);
        <a id="L105"></a>w.quitc = make(chan bool);
    <a id="L106"></a>}

    <a id="L108"></a>if subsys&amp;SubsystemAudio != 0 {
        <a id="L109"></a>var n int;
        <a id="L110"></a>if n, err = audioInit(AudioFormatStereo44K, 2048); err != nil {
            <a id="L111"></a>return nil, err
        <a id="L112"></a>}
        <a id="L113"></a>println(&#34;audio&#34;, n);
    <a id="L114"></a>}

    <a id="L116"></a>if subsys&amp;SubsystemVideo != 0 {
        <a id="L117"></a>go w.readEvents()
    <a id="L118"></a>}

    <a id="L120"></a>return w, nil;
<a id="L121"></a>}

<a id="L123"></a>func (w *Window) FlushImage() {
    <a id="L124"></a>if w.Image == nil {
        <a id="L125"></a>return
    <a id="L126"></a>}
    <a id="L127"></a>videoUpdate(w.Image.Linear);
<a id="L128"></a>}

<a id="L130"></a>func multimediaInit(subsys int) (err os.Error) {
    <a id="L131"></a>return os.NewSyscallError(&#34;multimedia_init&#34;, syscall.MultimediaInit(subsys))
<a id="L132"></a>}

<a id="L134"></a>func videoInit(dx, dy int) (ndx, ndy int, err os.Error) {
    <a id="L135"></a>if srpcEnabled {
        <a id="L136"></a>bridge.share.ready = 1;
        <a id="L137"></a>return int(bridge.share.width), int(bridge.share.height), nil;
    <a id="L138"></a>}
    <a id="L139"></a>if e := syscall.VideoInit(dx, dy); e != 0 {
        <a id="L140"></a>return 0, 0, os.NewSyscallError(&#34;video_init&#34;, int(e))
    <a id="L141"></a>}
    <a id="L142"></a>return dx, dy, nil;
<a id="L143"></a>}

<a id="L145"></a>func videoUpdate(data []Color) (err os.Error) {
    <a id="L146"></a>if srpcEnabled {
        <a id="L147"></a>bridge.flushRPC.Call(&#34;upcall&#34;, nil);
        <a id="L148"></a>return;
    <a id="L149"></a>}
    <a id="L150"></a>return os.NewSyscallError(&#34;video_update&#34;, syscall.VideoUpdate((*uint32)(&amp;data[0])));
<a id="L151"></a>}

<a id="L153"></a>var noEvents = os.NewError(&#34;no events&#34;)

<a id="L155"></a>func videoPollEvent(ev []byte) (err os.Error) {
    <a id="L156"></a>if srpcEnabled {
        <a id="L157"></a>r := bridge.share.eq.ri;
        <a id="L158"></a>if r == bridge.share.eq.wi {
            <a id="L159"></a>return noEvents
        <a id="L160"></a>}
        <a id="L161"></a>bytes.Copy(ev, &amp;bridge.share.eq.event[r]);
        <a id="L162"></a>bridge.share.eq.ri = (r + 1) % eqsize;
        <a id="L163"></a>return nil;
    <a id="L164"></a>}
    <a id="L165"></a>return os.NewSyscallError(&#34;video_poll_event&#34;, syscall.VideoPollEvent(&amp;ev[0]));
<a id="L166"></a>}

<a id="L168"></a>func audioInit(fmt int, want int) (got int, err os.Error) {
    <a id="L169"></a>var x int;
    <a id="L170"></a>e := syscall.AudioInit(fmt, want, &amp;x);
    <a id="L171"></a>if e == 0 {
        <a id="L172"></a>return x, nil
    <a id="L173"></a>}
    <a id="L174"></a>return 0, os.NewSyscallError(&#34;audio_init&#34;, e);
<a id="L175"></a>}

<a id="L177"></a>var audioSize uintptr

<a id="L179"></a><span class="comment">// AudioStream provides access to the audio device.</span>
<a id="L180"></a><span class="comment">// Each call to AudioStream writes the given data,</span>
<a id="L181"></a><span class="comment">// which should be a slice of 16-bit stereo PCM audio samples,</span>
<a id="L182"></a><span class="comment">// and returns the number of samples required by the next</span>
<a id="L183"></a><span class="comment">// call to AudioStream.</span>
<a id="L184"></a><span class="comment">//</span>
<a id="L185"></a><span class="comment">// To find out the initial number of samples to write, call AudioStream(nil).</span>
<a id="L186"></a><span class="comment">//</span>
<a id="L187"></a>func AudioStream(data []uint16) (nextSize int, err os.Error) {
    <a id="L188"></a>if audioSize == 0 {
        <a id="L189"></a>e := os.NewSyscallError(&#34;audio_stream&#34;, syscall.AudioStream(nil, &amp;audioSize));
        <a id="L190"></a>return int(audioSize), e;
    <a id="L191"></a>}
    <a id="L192"></a>if data == nil {
        <a id="L193"></a>return int(audioSize), nil
    <a id="L194"></a>}
    <a id="L195"></a>if uintptr(len(data))*2 != audioSize {
        <a id="L196"></a>log.Stdoutf(&#34;invalid audio size want %d got %d&#34;, audioSize, len(data))
    <a id="L197"></a>}
    <a id="L198"></a>e := os.NewSyscallError(&#34;audio_stream&#34;, syscall.AudioStream(&amp;data[0], &amp;audioSize));
    <a id="L199"></a>return int(audioSize), e;
<a id="L200"></a>}

<a id="L202"></a><span class="comment">// Synchronization structure to wait for bridge to become ready.</span>
<a id="L203"></a>var bridge struct {
    <a id="L204"></a>c         chan bool;
    <a id="L205"></a>displayFd int;
    <a id="L206"></a>rpcFd     int;
    <a id="L207"></a>share     *videoShare;
    <a id="L208"></a>pixel     []Color;
    <a id="L209"></a>client    *srpc.Client;
    <a id="L210"></a>flushRPC  *srpc.RPC;
<a id="L211"></a>}

<a id="L213"></a><span class="comment">// Wait for bridge to become ready.</span>
<a id="L214"></a><span class="comment">// When chan is first created, there is nothing in it,</span>
<a id="L215"></a><span class="comment">// so this blocks.  Once the bridge is ready, multimediaBridge.Run</span>
<a id="L216"></a><span class="comment">// will drop a value into the channel.  Then any calls</span>
<a id="L217"></a><span class="comment">// to waitBridge will finish, taking the value out and immediately putting it back.</span>
<a id="L218"></a>func waitBridge() { bridge.c &lt;- &lt;-bridge.c }

<a id="L220"></a>const eqsize = 64

<a id="L222"></a><span class="comment">// Data structure shared with host via mmap.</span>
<a id="L223"></a>type videoShare struct {
    <a id="L224"></a>revision int32; <span class="comment">// definition below is rev 100 unless noted</span>
    <a id="L225"></a>mapSize  int32;

    <a id="L227"></a><span class="comment">// event queue</span>
    <a id="L228"></a>eq  struct {
        <a id="L229"></a>ri    uint32; <span class="comment">// read index [0,eqsize)</span>
        <a id="L230"></a>wi    uint32; <span class="comment">// write index [0,eqsize)</span>
        <a id="L231"></a>eof   int32;
        <a id="L232"></a>event [eqsize][64]byte;
    <a id="L233"></a>};

    <a id="L235"></a><span class="comment">// now unused</span>
    <a id="L236"></a>_, _, _, _ int32;

    <a id="L238"></a><span class="comment">// video backing store information</span>
    <a id="L239"></a>width, height, _, size int32;
    <a id="L240"></a>ready                  int32; <span class="comment">// rev 0x101</span>
<a id="L241"></a>}

<a id="L243"></a><span class="comment">// The frame buffer data is videoShareSize bytes after</span>
<a id="L244"></a><span class="comment">// the videoShare begins.</span>
<a id="L245"></a>const videoShareSize = 16 * 1024

<a id="L247"></a>type multimediaBridge struct{}

<a id="L249"></a><span class="comment">// If using SRPC, the runtime will call this method to pass in two file descriptors,</span>
<a id="L250"></a><span class="comment">// one to mmap to get the display memory, and another to use for SRPCs back</span>
<a id="L251"></a><span class="comment">// to the main process.</span>
<a id="L252"></a>func (multimediaBridge) Run(arg, ret []interface{}, size []int) srpc.Errno {
    <a id="L253"></a>bridge.displayFd = arg[0].(int);
    <a id="L254"></a>bridge.rpcFd = arg[1].(int);

    <a id="L256"></a>var st syscall.Stat_t;
    <a id="L257"></a>if errno := syscall.Fstat(bridge.displayFd, &amp;st); errno != 0 {
        <a id="L258"></a>log.Exitf(&#34;mmbridge stat display: %s&#34;, os.Errno(errno))
    <a id="L259"></a>}

    <a id="L261"></a>addr, _, errno := syscall.Syscall6(syscall.SYS_MMAP,
        <a id="L262"></a>0,
        <a id="L263"></a>uintptr(st.Size),
        <a id="L264"></a>syscall.PROT_READ|syscall.PROT_WRITE,
        <a id="L265"></a>syscall.MAP_SHARED,
        <a id="L266"></a>uintptr(bridge.displayFd),
        <a id="L267"></a>0);
    <a id="L268"></a>if errno != 0 {
        <a id="L269"></a>log.Exitf(&#34;mmap display: %s&#34;, os.Errno(errno))
    <a id="L270"></a>}

    <a id="L272"></a>bridge.share = (*videoShare)(unsafe.Pointer(addr));

    <a id="L274"></a><span class="comment">// Overestimate frame buffer size</span>
    <a id="L275"></a><span class="comment">// (must use a compile-time constant)</span>
    <a id="L276"></a><span class="comment">// and then reslice.  256 megapixels (1 GB) should be enough.</span>
    <a id="L277"></a>fb := (*[256 * 1024 * 1024]Color)(unsafe.Pointer(addr + videoShareSize));
    <a id="L278"></a>bridge.pixel = fb[0 : (st.Size-videoShareSize)/4];

    <a id="L280"></a><span class="comment">// Configure RPC connection back to client.</span>
    <a id="L281"></a>var err os.Error;
    <a id="L282"></a>bridge.client, err = srpc.NewClient(bridge.rpcFd);
    <a id="L283"></a>if err != nil {
        <a id="L284"></a>log.Exitf(&#34;NewClient: %s&#34;, err)
    <a id="L285"></a>}
    <a id="L286"></a>bridge.flushRPC = bridge.client.NewRPC(nil);

    <a id="L288"></a><span class="comment">// Notify waiters that the bridge is ready.</span>
    <a id="L289"></a>println(&#34;bridged&#34;, bridge.share.revision);
    <a id="L290"></a>bridge.c &lt;- true;

    <a id="L292"></a>return srpc.OK;
<a id="L293"></a>}

<a id="L295"></a>func init() {
    <a id="L296"></a>bridge.c = make(chan bool, 1);
    <a id="L297"></a>if srpcEnabled {
        <a id="L298"></a>srpc.Add(&#34;nacl_multimedia_bridge&#34;, &#34;hh:&#34;, multimediaBridge{})
    <a id="L299"></a>}
<a id="L300"></a>}
</pre>

</div>

<div id="footer">
<p>Except as noted, this content is
   licensed under <a href="http://creativecommons.org/licenses/by/3.0/">
   Creative Commons Attribution 3.0</a>.
</div>

<script type="text/javascript">
var gaJsHost = (("https:" == document.location.protocol) ? "https://ssl." : "http://www.");
document.write(unescape("%3Cscript src='" + gaJsHost + "google-analytics.com/ga.js' type='text/javascript'%3E%3C/script%3E"));
</script>
<script type="text/javascript">
var pageTracker = _gat._getTracker("UA-11222381-2");
pageTracker._trackPageview();
</script>
</body>
</html>
<!-- generated at Thu Nov 12 15:42:51 PST 2009 -->
