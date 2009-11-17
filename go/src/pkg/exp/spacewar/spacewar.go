<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/exp/spacewar/spacewar.go</title>

  <link rel="stylesheet" type="text/css" href="../../../../doc/style.css">
  <script type="text/javascript" src="../../../../doc/godocs.js"></script>

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
        <a href="../../../../index.html"><img src="../../../../doc/logo-153x55.png" height="55" width="153" alt="Go Home Page" style="border:0" /></a>
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
    <li class="navhead"><a href="../../../../index.html">Home</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Documents</li>
    <li><a href="../../../../doc/go_tutorial.html">Tutorial</a></li>
    <li><a href="../../../../doc/effective_go.html">Effective Go</a></li>
    <li><a href="../../../../doc/go_faq.html">FAQ</a></li>
    <li><a href="../../../../doc/go_lang_faq.html">Language Design FAQ</a></li>
    <li><a href="http://www.youtube.com/watch?v=rKnDgT73v8s">Tech talk (1 hour)</a> (<a href="../../../../doc/go_talk-20091030.pdf">PDF</a>)</li>
    <li><a href="../../../../doc/go_spec.html">Language Specification</a></li>
    <li><a href="../../../../doc/go_mem.html">Memory Model</a></li>
    <li><a href="../../../../doc/go_for_cpp_programmers.html">Go for C++ Programmers</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">How To</li>
    <li><a href="../../../../doc/install.html">Install Go</a></li>
    <li><a href="../../../../doc/contribute.html">Contribute code</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Programming</li>
    <li><a href="../../../../cmd/index.html">Command documentation</a></li>
    <li><a href="../../../../pkg/index.html">Package documentation</a></li>
    <li><a href="../../../index.html">Source files</a></li>

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
	<li>Thu Nov 12 15:57:42 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/exp/spacewar/spacewar.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright (c) 1996 Barry Silverman, Brian Silverman, Vadim Gerasimov.</span>
<a id="L2"></a><span class="comment">// Portions Copyright (c) 2009 The Go Authors.</span>
<a id="L3"></a><span class="comment">//</span>
<a id="L4"></a><span class="comment">// Permission is hereby granted, free of charge, to any person obtaining a copy</span>
<a id="L5"></a><span class="comment">// of this software and associated documentation files (the &#34;Software&#34;), to deal</span>
<a id="L6"></a><span class="comment">// in the Software without restriction, including without limitation the rights</span>
<a id="L7"></a><span class="comment">// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell</span>
<a id="L8"></a><span class="comment">// copies of the Software, and to permit persons to whom the Software is</span>
<a id="L9"></a><span class="comment">// furnished to do so, subject to the following conditions:</span>
<a id="L10"></a><span class="comment">//</span>
<a id="L11"></a><span class="comment">// The above copyright notice and this permission notice shall be included in</span>
<a id="L12"></a><span class="comment">// all copies or substantial portions of the Software.</span>
<a id="L13"></a><span class="comment">//</span>
<a id="L14"></a><span class="comment">// THE SOFTWARE IS PROVIDED &#34;AS IS&#34;, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR</span>
<a id="L15"></a><span class="comment">// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,</span>
<a id="L16"></a><span class="comment">// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE</span>
<a id="L17"></a><span class="comment">// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER</span>
<a id="L18"></a><span class="comment">// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,</span>
<a id="L19"></a><span class="comment">// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN</span>
<a id="L20"></a><span class="comment">// THE SOFTWARE.</span>

<a id="L22"></a>package main

<a id="L24"></a>import (
    <a id="L25"></a>&#34;bytes&#34;;
    <a id="L26"></a>&#34;exp/draw&#34;;
    <a id="L27"></a>&#34;exp/nacl/av&#34;;
    <a id="L28"></a>&#34;exp/nacl/srpc&#34;;
    <a id="L29"></a>&#34;image&#34;;
    <a id="L30"></a>&#34;log&#34;;
    <a id="L31"></a>&#34;os&#34;;
    <a id="L32"></a>&#34;runtime&#34;;
    <a id="L33"></a>&#34;strings&#34;;
    <a id="L34"></a>&#34;time&#34;;
    <a id="L35"></a>&#34;./pdp1&#34;;
<a id="L36"></a>)

<a id="L38"></a>func main() {
    <a id="L39"></a>runtime.LockOSThread();
    <a id="L40"></a>if srpc.Enabled() {
        <a id="L41"></a>go srpc.ServeRuntime()
    <a id="L42"></a>}

    <a id="L44"></a>w, err := av.Init(av.SubsystemVideo, 512, 512);
    <a id="L45"></a>if err != nil {
        <a id="L46"></a>log.Exitf(&#34;av.Init: %s&#34;, err)
    <a id="L47"></a>}

    <a id="L49"></a>go quitter(w.QuitChan());

    <a id="L51"></a>var m SpacewarPDP1;
    <a id="L52"></a>m.Init(w);
    <a id="L53"></a>m.PC = 4;
    <a id="L54"></a>f := bytes.NewBuffer(strings.Bytes(spacewarCode));
    <a id="L55"></a>if err = m.Load(f); err != nil {
        <a id="L56"></a>log.Exitf(&#34;loading %s: %s&#34;, &#34;spacewar.lst&#34;, err)
    <a id="L57"></a>}
    <a id="L58"></a>for err == nil {
        <a id="L59"></a><span class="comment">//fmt.Printf(&#34;step PC=%06o &#34;, m.PC);</span>
        <a id="L60"></a><span class="comment">//fmt.Printf(&#34;inst=%06o AC=%06o IO=%06o OV=%o\n&#34;,</span>
        <a id="L61"></a><span class="comment">//	m.Mem[m.PC], m.AC, m.IO, m.OV);</span>
        <a id="L62"></a>err = m.Step()
    <a id="L63"></a>}
    <a id="L64"></a>log.Exitf(&#34;step: %s&#34;, err);
<a id="L65"></a>}

<a id="L67"></a>func quitter(c &lt;-chan bool) {
    <a id="L68"></a>&lt;-c;
    <a id="L69"></a>os.Exit(0);
<a id="L70"></a>}

<a id="L72"></a><span class="comment">// A SpacewarPDP1 is a PDP-1 machine configured to run Spacewar!</span>
<a id="L73"></a><span class="comment">// It responds to traps by drawing on the display, and it flushes the</span>
<a id="L74"></a><span class="comment">// display and pauses every second time the program counter reaches</span>
<a id="L75"></a><span class="comment">// instruction 02051.</span>
<a id="L76"></a>type SpacewarPDP1 struct {
    <a id="L77"></a>pdp1.M;
    <a id="L78"></a>nframe     int;
    <a id="L79"></a>frameTime  int64;
    <a id="L80"></a>ctxt       draw.Context;
    <a id="L81"></a>dx, dy     int;
    <a id="L82"></a>screen     draw.Image;
    <a id="L83"></a>ctl        pdp1.Word;
    <a id="L84"></a>kc         &lt;-chan int;
    <a id="L85"></a>colorModel image.ColorModel;
    <a id="L86"></a>cmap       []image.Color;
    <a id="L87"></a>pix        [][]uint8;
<a id="L88"></a>}

<a id="L90"></a>func min(a, b int) int {
    <a id="L91"></a>if a &lt; b {
        <a id="L92"></a>return a
    <a id="L93"></a>}
    <a id="L94"></a>return b;
<a id="L95"></a>}

<a id="L97"></a>func (m *SpacewarPDP1) Init(ctxt draw.Context) {
    <a id="L98"></a>m.ctxt = ctxt;
    <a id="L99"></a>m.kc = ctxt.KeyboardChan();
    <a id="L100"></a>m.screen = ctxt.Screen();
    <a id="L101"></a>m.dx = m.screen.Width();
    <a id="L102"></a>m.dy = m.screen.Height();
    <a id="L103"></a>m.colorModel = m.screen.ColorModel();
    <a id="L104"></a>m.pix = make([][]uint8, m.dy);
    <a id="L105"></a>for i := range m.pix {
        <a id="L106"></a>m.pix[i] = make([]uint8, m.dx)
    <a id="L107"></a>}
    <a id="L108"></a>m.cmap = make([]image.Color, 256);
    <a id="L109"></a>for i := range m.cmap {
        <a id="L110"></a>var r, g, b uint8;
        <a id="L111"></a>r = uint8(min(0, 255));
        <a id="L112"></a>g = uint8(min(i*2, 255));
        <a id="L113"></a>b = uint8(min(0, 255));
        <a id="L114"></a>m.cmap[i] = m.colorModel.Convert(image.RGBAColor{r, g, b, 0xff});
    <a id="L115"></a>}
<a id="L116"></a>}

<a id="L118"></a>const (
    <a id="L119"></a>frameDelay = 56 * 1e6; <span class="comment">// 56 ms</span>
<a id="L120"></a>)

<a id="L122"></a>var ctlBits = [...]pdp1.Word{
    <a id="L123"></a>&#39;f&#39;: 0000001,
    <a id="L124"></a>&#39;d&#39;: 0000002,
    <a id="L125"></a>&#39;a&#39;: 0000004,
    <a id="L126"></a>&#39;s&#39;: 0000010,
    <a id="L127"></a>&#39;\&#39;&#39;: 0040000,
    <a id="L128"></a>&#39;;&#39;: 0100000,
    <a id="L129"></a>&#39;k&#39;: 0200000,
    <a id="L130"></a>&#39;l&#39;: 0400000,
<a id="L131"></a>}

<a id="L133"></a>func (m *SpacewarPDP1) Step() os.Error {
    <a id="L134"></a>if m.PC == 02051 {
        <a id="L135"></a>m.pollInput();
        <a id="L136"></a>m.nframe++;
        <a id="L137"></a>if m.nframe&amp;1 == 0 {
            <a id="L138"></a>m.flush();
            <a id="L139"></a>t := time.Nanoseconds();
            <a id="L140"></a>if t &gt;= m.frameTime+3*frameDelay {
                <a id="L141"></a>m.frameTime = t
            <a id="L142"></a>} else {
                <a id="L143"></a>m.frameTime += frameDelay;
                <a id="L144"></a>for t &lt; m.frameTime {
                    <a id="L145"></a>time.Sleep(m.frameTime - t);
                    <a id="L146"></a>t = time.Nanoseconds();
                <a id="L147"></a>}
            <a id="L148"></a>}
        <a id="L149"></a>}
    <a id="L150"></a>}
    <a id="L151"></a>return m.M.Step(m);
<a id="L152"></a>}

<a id="L154"></a>func (m *SpacewarPDP1) Trap(y pdp1.Word) {
    <a id="L155"></a>switch y &amp; 077 {
    <a id="L156"></a>case 7:
        <a id="L157"></a>x := int(m.AC+0400000) &amp; 0777777;
        <a id="L158"></a>y := int(m.IO+0400000) &amp; 0777777;
        <a id="L159"></a>x = x * m.dx / 0777777;
        <a id="L160"></a>y = y * m.dy / 0777777;
        <a id="L161"></a>if 0 &lt;= x &amp;&amp; x &lt; m.dx &amp;&amp; 0 &lt;= y &amp;&amp; y &lt; m.dy {
            <a id="L162"></a>n := uint8(min(int(m.pix[y][x])+128, 255));
            <a id="L163"></a>m.pix[y][x] = n;
        <a id="L164"></a>}
    <a id="L165"></a>case 011:
        <a id="L166"></a>m.IO = m.ctl
    <a id="L167"></a>}
<a id="L168"></a>}

<a id="L170"></a>func (m *SpacewarPDP1) flush() {
    <a id="L171"></a><span class="comment">// Update screen image; simulate phosphor decay.</span>
    <a id="L172"></a>for y := 0; y &lt; m.dy; y++ {
        <a id="L173"></a>for x := 0; x &lt; m.dx; x++ {
            <a id="L174"></a>m.screen.Set(x, y, m.cmap[m.pix[y][x]]);
            <a id="L175"></a>m.pix[y][x] &gt;&gt;= 1;
        <a id="L176"></a>}
    <a id="L177"></a>}
    <a id="L178"></a>m.ctxt.FlushImage();
<a id="L179"></a>}

<a id="L181"></a>func (m *SpacewarPDP1) pollInput() {
    <a id="L182"></a>for {
        <a id="L183"></a>select {
        <a id="L184"></a>case ch := &lt;-m.kc:
            <a id="L185"></a>if 0 &lt;= ch &amp;&amp; ch &lt; len(ctlBits) {
                <a id="L186"></a>m.ctl |= ctlBits[ch]
            <a id="L187"></a>}
            <a id="L188"></a>if 0 &lt;= -ch &amp;&amp; -ch &lt; len(ctlBits) {
                <a id="L189"></a>m.ctl &amp;^= ctlBits[-ch]
            <a id="L190"></a>}
        <a id="L191"></a>default:
            <a id="L192"></a>return
        <a id="L193"></a>}
    <a id="L194"></a>}
<a id="L195"></a>}
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
