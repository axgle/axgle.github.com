<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/io/pipe_test.go</title>

  <link rel="stylesheet" type="text/css" href="../../../doc/style.css">
  <script type="text/javascript" src="../../../doc/godocs.js"></script>

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
        <a href="../../../index.html"><img src="../../../doc/logo-153x55.png" height="55" width="153" alt="Go Home Page" style="border:0" /></a>
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
    <li class="navhead"><a href="../../../index.html">Home</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Documents</li>
    <li><a href="../../../doc/go_tutorial.html">Tutorial</a></li>
    <li><a href="../../../doc/effective_go.html">Effective Go</a></li>
    <li><a href="../../../doc/go_faq.html">FAQ</a></li>
    <li><a href="../../../doc/go_lang_faq.html">Language Design FAQ</a></li>
    <li><a href="http://www.youtube.com/watch?v=rKnDgT73v8s">Tech talk (1 hour)</a> (<a href="../../../doc/go_talk-20091030.pdf">PDF</a>)</li>
    <li><a href="../../../doc/go_spec.html">Language Specification</a></li>
    <li><a href="../../../doc/go_mem.html">Memory Model</a></li>
    <li><a href="../../../doc/go_for_cpp_programmers.html">Go for C++ Programmers</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">How To</li>
    <li><a href="../../../doc/install.html">Install Go</a></li>
    <li><a href="../../../doc/contribute.html">Contribute code</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Programming</li>
    <li><a href="../../../cmd/index.html">Command documentation</a></li>
    <li><a href="../../../pkg/index.html">Package documentation</a></li>
    <li><a href="../../index.html">Source files</a></li>

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
	<li>Sun Nov 15 20:28:13 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/io/pipe_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package io_test

<a id="L7"></a>import (
    <a id="L8"></a>&#34;fmt&#34;;
    <a id="L9"></a>. &#34;io&#34;;
    <a id="L10"></a>&#34;os&#34;;
    <a id="L11"></a>&#34;strings&#34;;
    <a id="L12"></a>&#34;testing&#34;;
    <a id="L13"></a>&#34;time&#34;;
<a id="L14"></a>)

<a id="L16"></a>func checkWrite(t *testing.T, w Writer, data []byte, c chan int) {
    <a id="L17"></a>n, err := w.Write(data);
    <a id="L18"></a>if err != nil {
        <a id="L19"></a>t.Errorf(&#34;write: %v&#34;, err)
    <a id="L20"></a>}
    <a id="L21"></a>if n != len(data) {
        <a id="L22"></a>t.Errorf(&#34;short write: %d != %d&#34;, n, len(data))
    <a id="L23"></a>}
    <a id="L24"></a>c &lt;- 0;
<a id="L25"></a>}

<a id="L27"></a><span class="comment">// Test a single read/write pair.</span>
<a id="L28"></a>func TestPipe1(t *testing.T) {
    <a id="L29"></a>c := make(chan int);
    <a id="L30"></a>r, w := Pipe();
    <a id="L31"></a>var buf = make([]byte, 64);
    <a id="L32"></a>go checkWrite(t, w, strings.Bytes(&#34;hello, world&#34;), c);
    <a id="L33"></a>n, err := r.Read(buf);
    <a id="L34"></a>if err != nil {
        <a id="L35"></a>t.Errorf(&#34;read: %v&#34;, err)
    <a id="L36"></a>} else if n != 12 || string(buf[0:12]) != &#34;hello, world&#34; {
        <a id="L37"></a>t.Errorf(&#34;bad read: got %q&#34;, buf[0:n])
    <a id="L38"></a>}
    <a id="L39"></a>&lt;-c;
    <a id="L40"></a>r.Close();
    <a id="L41"></a>w.Close();
<a id="L42"></a>}

<a id="L44"></a>func reader(t *testing.T, r Reader, c chan int) {
    <a id="L45"></a>var buf = make([]byte, 64);
    <a id="L46"></a>for {
        <a id="L47"></a>n, err := r.Read(buf);
        <a id="L48"></a>if err == os.EOF {
            <a id="L49"></a>c &lt;- 0;
            <a id="L50"></a>break;
        <a id="L51"></a>}
        <a id="L52"></a>if err != nil {
            <a id="L53"></a>t.Errorf(&#34;read: %v&#34;, err)
        <a id="L54"></a>}
        <a id="L55"></a>c &lt;- n;
    <a id="L56"></a>}
<a id="L57"></a>}

<a id="L59"></a><span class="comment">// Test a sequence of read/write pairs.</span>
<a id="L60"></a>func TestPipe2(t *testing.T) {
    <a id="L61"></a>c := make(chan int);
    <a id="L62"></a>r, w := Pipe();
    <a id="L63"></a>go reader(t, r, c);
    <a id="L64"></a>var buf = make([]byte, 64);
    <a id="L65"></a>for i := 0; i &lt; 5; i++ {
        <a id="L66"></a>p := buf[0 : 5+i*10];
        <a id="L67"></a>n, err := w.Write(p);
        <a id="L68"></a>if n != len(p) {
            <a id="L69"></a>t.Errorf(&#34;wrote %d, got %d&#34;, len(p), n)
        <a id="L70"></a>}
        <a id="L71"></a>if err != nil {
            <a id="L72"></a>t.Errorf(&#34;write: %v&#34;, err)
        <a id="L73"></a>}
        <a id="L74"></a>nn := &lt;-c;
        <a id="L75"></a>if nn != n {
            <a id="L76"></a>t.Errorf(&#34;wrote %d, read got %d&#34;, n, nn)
        <a id="L77"></a>}
    <a id="L78"></a>}
    <a id="L79"></a>w.Close();
    <a id="L80"></a>nn := &lt;-c;
    <a id="L81"></a>if nn != 0 {
        <a id="L82"></a>t.Errorf(&#34;final read got %d&#34;, nn)
    <a id="L83"></a>}
<a id="L84"></a>}

<a id="L86"></a>type pipeReturn struct {
    <a id="L87"></a>n   int;
    <a id="L88"></a>err os.Error;
<a id="L89"></a>}

<a id="L91"></a><span class="comment">// Test a large write that requires multiple reads to satisfy.</span>
<a id="L92"></a>func writer(w WriteCloser, buf []byte, c chan pipeReturn) {
    <a id="L93"></a>n, err := w.Write(buf);
    <a id="L94"></a>w.Close();
    <a id="L95"></a>c &lt;- pipeReturn{n, err};
<a id="L96"></a>}

<a id="L98"></a>func TestPipe3(t *testing.T) {
    <a id="L99"></a>c := make(chan pipeReturn);
    <a id="L100"></a>r, w := Pipe();
    <a id="L101"></a>var wdat = make([]byte, 128);
    <a id="L102"></a>for i := 0; i &lt; len(wdat); i++ {
        <a id="L103"></a>wdat[i] = byte(i)
    <a id="L104"></a>}
    <a id="L105"></a>go writer(w, wdat, c);
    <a id="L106"></a>var rdat = make([]byte, 1024);
    <a id="L107"></a>tot := 0;
    <a id="L108"></a>for n := 1; n &lt;= 256; n *= 2 {
        <a id="L109"></a>nn, err := r.Read(rdat[tot : tot+n]);
        <a id="L110"></a>if err != nil &amp;&amp; err != os.EOF {
            <a id="L111"></a>t.Fatalf(&#34;read: %v&#34;, err)
        <a id="L112"></a>}

        <a id="L114"></a><span class="comment">// only final two reads should be short - 1 byte, then 0</span>
        <a id="L115"></a>expect := n;
        <a id="L116"></a>if n == 128 {
            <a id="L117"></a>expect = 1
        <a id="L118"></a>} else if n == 256 {
            <a id="L119"></a>expect = 0;
            <a id="L120"></a>if err != os.EOF {
                <a id="L121"></a>t.Fatalf(&#34;read at end: %v&#34;, err)
            <a id="L122"></a>}
        <a id="L123"></a>}
        <a id="L124"></a>if nn != expect {
            <a id="L125"></a>t.Fatalf(&#34;read %d, expected %d, got %d&#34;, n, expect, nn)
        <a id="L126"></a>}
        <a id="L127"></a>tot += nn;
    <a id="L128"></a>}
    <a id="L129"></a>pr := &lt;-c;
    <a id="L130"></a>if pr.n != 128 || pr.err != nil {
        <a id="L131"></a>t.Fatalf(&#34;write 128: %d, %v&#34;, pr.n, pr.err)
    <a id="L132"></a>}
    <a id="L133"></a>if tot != 128 {
        <a id="L134"></a>t.Fatalf(&#34;total read %d != 128&#34;, tot)
    <a id="L135"></a>}
    <a id="L136"></a>for i := 0; i &lt; 128; i++ {
        <a id="L137"></a>if rdat[i] != byte(i) {
            <a id="L138"></a>t.Fatalf(&#34;rdat[%d] = %d&#34;, i, rdat[i])
        <a id="L139"></a>}
    <a id="L140"></a>}
<a id="L141"></a>}

<a id="L143"></a><span class="comment">// Test read after/before writer close.</span>

<a id="L145"></a>type closer interface {
    <a id="L146"></a>CloseWithError(os.Error) os.Error;
    <a id="L147"></a>Close() os.Error;
<a id="L148"></a>}

<a id="L150"></a>type pipeTest struct {
    <a id="L151"></a>async          bool;
    <a id="L152"></a>err            os.Error;
    <a id="L153"></a>closeWithError bool;
<a id="L154"></a>}

<a id="L156"></a>func (p pipeTest) String() string {
    <a id="L157"></a>return fmt.Sprintf(&#34;async=%v err=%v closeWithError=%v&#34;, p.async, p.err, p.closeWithError)
<a id="L158"></a>}

<a id="L160"></a>var pipeTests = []pipeTest{
    <a id="L161"></a>pipeTest{true, nil, false},
    <a id="L162"></a>pipeTest{true, nil, true},
    <a id="L163"></a>pipeTest{true, ErrShortWrite, true},
    <a id="L164"></a>pipeTest{false, nil, false},
    <a id="L165"></a>pipeTest{false, nil, true},
    <a id="L166"></a>pipeTest{false, ErrShortWrite, true},
<a id="L167"></a>}

<a id="L169"></a>func delayClose(t *testing.T, cl closer, ch chan int, tt pipeTest) {
    <a id="L170"></a>time.Sleep(1e6); <span class="comment">// 1 ms</span>
    <a id="L171"></a>var err os.Error;
    <a id="L172"></a>if tt.closeWithError {
        <a id="L173"></a>err = cl.CloseWithError(tt.err)
    <a id="L174"></a>} else {
        <a id="L175"></a>err = cl.Close()
    <a id="L176"></a>}
    <a id="L177"></a>if err != nil {
        <a id="L178"></a>t.Errorf(&#34;delayClose: %v&#34;, err)
    <a id="L179"></a>}
    <a id="L180"></a>ch &lt;- 0;
<a id="L181"></a>}

<a id="L183"></a>func TestPipeReadClose(t *testing.T) {
    <a id="L184"></a>for _, tt := range pipeTests {
        <a id="L185"></a>c := make(chan int, 1);
        <a id="L186"></a>r, w := Pipe();
        <a id="L187"></a>if tt.async {
            <a id="L188"></a>go delayClose(t, w, c, tt)
        <a id="L189"></a>} else {
            <a id="L190"></a>delayClose(t, w, c, tt)
        <a id="L191"></a>}
        <a id="L192"></a>var buf = make([]byte, 64);
        <a id="L193"></a>n, err := r.Read(buf);
        <a id="L194"></a>&lt;-c;
        <a id="L195"></a>want := tt.err;
        <a id="L196"></a>if want == nil {
            <a id="L197"></a>want = os.EOF
        <a id="L198"></a>}
        <a id="L199"></a>if err != want {
            <a id="L200"></a>t.Errorf(&#34;read from closed pipe: %v want %v&#34;, err, want)
        <a id="L201"></a>}
        <a id="L202"></a>if n != 0 {
            <a id="L203"></a>t.Errorf(&#34;read on closed pipe returned %d&#34;, n)
        <a id="L204"></a>}
        <a id="L205"></a>if err = r.Close(); err != nil {
            <a id="L206"></a>t.Errorf(&#34;r.Close: %v&#34;, err)
        <a id="L207"></a>}
    <a id="L208"></a>}
<a id="L209"></a>}

<a id="L211"></a><span class="comment">// Test write after/before reader close.</span>

<a id="L213"></a>func TestPipeWriteClose(t *testing.T) {
    <a id="L214"></a>for _, tt := range pipeTests {
        <a id="L215"></a>c := make(chan int, 1);
        <a id="L216"></a>r, w := Pipe();
        <a id="L217"></a>if tt.async {
            <a id="L218"></a>go delayClose(t, r, c, tt)
        <a id="L219"></a>} else {
            <a id="L220"></a>delayClose(t, r, c, tt)
        <a id="L221"></a>}
        <a id="L222"></a>n, err := WriteString(w, &#34;hello, world&#34;);
        <a id="L223"></a>&lt;-c;
        <a id="L224"></a>expect := tt.err;
        <a id="L225"></a>if expect == nil {
            <a id="L226"></a>expect = os.EPIPE
        <a id="L227"></a>}
        <a id="L228"></a>if err != expect {
            <a id="L229"></a>t.Errorf(&#34;write on closed pipe: %v want %v&#34;, err, expect)
        <a id="L230"></a>}
        <a id="L231"></a>if n != 0 {
            <a id="L232"></a>t.Errorf(&#34;write on closed pipe returned %d&#34;, n)
        <a id="L233"></a>}
        <a id="L234"></a>if err = w.Close(); err != nil {
            <a id="L235"></a>t.Errorf(&#34;w.Close: %v&#34;, err)
        <a id="L236"></a>}
    <a id="L237"></a>}
<a id="L238"></a>}
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
