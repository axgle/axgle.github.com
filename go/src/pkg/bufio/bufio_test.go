<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/bufio/bufio_test.go</title>

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
	<li>Thu Nov 12 16:00:00 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/bufio/bufio_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package bufio

<a id="L7"></a>import (
    <a id="L8"></a>&#34;bytes&#34;;
    <a id="L9"></a>&#34;fmt&#34;;
    <a id="L10"></a>&#34;io&#34;;
    <a id="L11"></a>&#34;os&#34;;
    <a id="L12"></a>&#34;strings&#34;;
    <a id="L13"></a>&#34;testing&#34;;
    <a id="L14"></a>&#34;testing/iotest&#34;;
<a id="L15"></a>)

<a id="L17"></a><span class="comment">// Reads from a reader and rot13s the result.</span>
<a id="L18"></a>type rot13Reader struct {
    <a id="L19"></a>r io.Reader;
<a id="L20"></a>}

<a id="L22"></a>func newRot13Reader(r io.Reader) *rot13Reader {
    <a id="L23"></a>r13 := new(rot13Reader);
    <a id="L24"></a>r13.r = r;
    <a id="L25"></a>return r13;
<a id="L26"></a>}

<a id="L28"></a>func (r13 *rot13Reader) Read(p []byte) (int, os.Error) {
    <a id="L29"></a>n, e := r13.r.Read(p);
    <a id="L30"></a>if e != nil {
        <a id="L31"></a>return n, e
    <a id="L32"></a>}
    <a id="L33"></a>for i := 0; i &lt; n; i++ {
        <a id="L34"></a>c := p[i] | 0x20; <span class="comment">// lowercase byte</span>
        <a id="L35"></a>if &#39;a&#39; &lt;= c &amp;&amp; c &lt;= &#39;m&#39; {
            <a id="L36"></a>p[i] += 13
        <a id="L37"></a>} else if &#39;n&#39; &lt;= c &amp;&amp; c &lt;= &#39;z&#39; {
            <a id="L38"></a>p[i] -= 13
        <a id="L39"></a>}
    <a id="L40"></a>}
    <a id="L41"></a>return n, nil;
<a id="L42"></a>}

<a id="L44"></a><span class="comment">// Call ReadByte to accumulate the text of a file</span>
<a id="L45"></a>func readBytes(buf *Reader) string {
    <a id="L46"></a>var b [1000]byte;
    <a id="L47"></a>nb := 0;
    <a id="L48"></a>for {
        <a id="L49"></a>c, e := buf.ReadByte();
        <a id="L50"></a>if e == os.EOF {
            <a id="L51"></a>break
        <a id="L52"></a>}
        <a id="L53"></a>if e != nil {
            <a id="L54"></a>panic(&#34;Data: &#34; + e.String())
        <a id="L55"></a>}
        <a id="L56"></a>b[nb] = c;
        <a id="L57"></a>nb++;
    <a id="L58"></a>}
    <a id="L59"></a>return string(b[0:nb]);
<a id="L60"></a>}

<a id="L62"></a>func TestReaderSimple(t *testing.T) {
    <a id="L63"></a>data := &#34;hello world&#34;;
    <a id="L64"></a>b := NewReader(bytes.NewBufferString(data));
    <a id="L65"></a>if s := readBytes(b); s != &#34;hello world&#34; {
        <a id="L66"></a>t.Errorf(&#34;simple hello world test failed: got %q&#34;, s)
    <a id="L67"></a>}

    <a id="L69"></a>b = NewReader(newRot13Reader(bytes.NewBufferString(data)));
    <a id="L70"></a>if s := readBytes(b); s != &#34;uryyb jbeyq&#34; {
        <a id="L71"></a>t.Error(&#34;rot13 hello world test failed: got %q&#34;, s)
    <a id="L72"></a>}
<a id="L73"></a>}


<a id="L76"></a>type readMaker struct {
    <a id="L77"></a>name string;
    <a id="L78"></a>fn   func(io.Reader) io.Reader;
<a id="L79"></a>}

<a id="L81"></a>var readMakers = []readMaker{
    <a id="L82"></a>readMaker{&#34;full&#34;, func(r io.Reader) io.Reader { return r }},
    <a id="L83"></a>readMaker{&#34;byte&#34;, iotest.OneByteReader},
    <a id="L84"></a>readMaker{&#34;half&#34;, iotest.HalfReader},
    <a id="L85"></a>readMaker{&#34;data+err&#34;, iotest.DataErrReader},
<a id="L86"></a>}

<a id="L88"></a><span class="comment">// Call ReadString (which ends up calling everything else)</span>
<a id="L89"></a><span class="comment">// to accumulate the text of a file.</span>
<a id="L90"></a>func readLines(b *Reader) string {
    <a id="L91"></a>s := &#34;&#34;;
    <a id="L92"></a>for {
        <a id="L93"></a>s1, e := b.ReadString(&#39;\n&#39;);
        <a id="L94"></a>if e == os.EOF {
            <a id="L95"></a>break
        <a id="L96"></a>}
        <a id="L97"></a>if e != nil {
            <a id="L98"></a>panic(&#34;GetLines: &#34; + e.String())
        <a id="L99"></a>}
        <a id="L100"></a>s += s1;
    <a id="L101"></a>}
    <a id="L102"></a>return s;
<a id="L103"></a>}

<a id="L105"></a><span class="comment">// Call Read to accumulate the text of a file</span>
<a id="L106"></a>func reads(buf *Reader, m int) string {
    <a id="L107"></a>var b [1000]byte;
    <a id="L108"></a>nb := 0;
    <a id="L109"></a>for {
        <a id="L110"></a>n, e := buf.Read(b[nb : nb+m]);
        <a id="L111"></a>nb += n;
        <a id="L112"></a>if e == os.EOF {
            <a id="L113"></a>break
        <a id="L114"></a>}
    <a id="L115"></a>}
    <a id="L116"></a>return string(b[0:nb]);
<a id="L117"></a>}

<a id="L119"></a>type bufReader struct {
    <a id="L120"></a>name string;
    <a id="L121"></a>fn   func(*Reader) string;
<a id="L122"></a>}

<a id="L124"></a>var bufreaders = []bufReader{
    <a id="L125"></a>bufReader{&#34;1&#34;, func(b *Reader) string { return reads(b, 1) }},
    <a id="L126"></a>bufReader{&#34;2&#34;, func(b *Reader) string { return reads(b, 2) }},
    <a id="L127"></a>bufReader{&#34;3&#34;, func(b *Reader) string { return reads(b, 3) }},
    <a id="L128"></a>bufReader{&#34;4&#34;, func(b *Reader) string { return reads(b, 4) }},
    <a id="L129"></a>bufReader{&#34;5&#34;, func(b *Reader) string { return reads(b, 5) }},
    <a id="L130"></a>bufReader{&#34;7&#34;, func(b *Reader) string { return reads(b, 7) }},
    <a id="L131"></a>bufReader{&#34;bytes&#34;, readBytes},
    <a id="L132"></a>bufReader{&#34;lines&#34;, readLines},
<a id="L133"></a>}

<a id="L135"></a>var bufsizes = []int{
    <a id="L136"></a>1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
    <a id="L137"></a>23, 32, 46, 64, 93, 128, 1024, 4096,
<a id="L138"></a>}

<a id="L140"></a>func TestReader(t *testing.T) {
    <a id="L141"></a>var texts [31]string;
    <a id="L142"></a>str := &#34;&#34;;
    <a id="L143"></a>all := &#34;&#34;;
    <a id="L144"></a>for i := 0; i &lt; len(texts)-1; i++ {
        <a id="L145"></a>texts[i] = str + &#34;\n&#34;;
        <a id="L146"></a>all += texts[i];
        <a id="L147"></a>str += string(i%26 + &#39;a&#39;);
    <a id="L148"></a>}
    <a id="L149"></a>texts[len(texts)-1] = all;

    <a id="L151"></a>for h := 0; h &lt; len(texts); h++ {
        <a id="L152"></a>text := texts[h];
        <a id="L153"></a>for i := 0; i &lt; len(readMakers); i++ {
            <a id="L154"></a>for j := 0; j &lt; len(bufreaders); j++ {
                <a id="L155"></a>for k := 0; k &lt; len(bufsizes); k++ {
                    <a id="L156"></a>readmaker := readMakers[i];
                    <a id="L157"></a>bufreader := bufreaders[j];
                    <a id="L158"></a>bufsize := bufsizes[k];
                    <a id="L159"></a>read := readmaker.fn(bytes.NewBufferString(text));
                    <a id="L160"></a>buf, _ := NewReaderSize(read, bufsize);
                    <a id="L161"></a>s := bufreader.fn(buf);
                    <a id="L162"></a>if s != text {
                        <a id="L163"></a>t.Errorf(&#34;reader=%s fn=%s bufsize=%d want=%q got=%q&#34;,
                            <a id="L164"></a>readmaker.name, bufreader.name, bufsize, text, s)
                    <a id="L165"></a>}
                <a id="L166"></a>}
            <a id="L167"></a>}
        <a id="L168"></a>}
    <a id="L169"></a>}
<a id="L170"></a>}

<a id="L172"></a><span class="comment">// A StringReader delivers its data one string segment at a time via Read.</span>
<a id="L173"></a>type StringReader struct {
    <a id="L174"></a>data []string;
    <a id="L175"></a>step int;
<a id="L176"></a>}

<a id="L178"></a>func (r *StringReader) Read(p []byte) (n int, err os.Error) {
    <a id="L179"></a>if r.step &lt; len(r.data) {
        <a id="L180"></a>s := r.data[r.step];
        <a id="L181"></a>for i := 0; i &lt; len(s); i++ {
            <a id="L182"></a>p[i] = s[i]
        <a id="L183"></a>}
        <a id="L184"></a>n = len(s);
        <a id="L185"></a>r.step++;
    <a id="L186"></a>} else {
        <a id="L187"></a>err = os.EOF
    <a id="L188"></a>}
    <a id="L189"></a>return;
<a id="L190"></a>}

<a id="L192"></a>func readRuneSegments(t *testing.T, segments []string) {
    <a id="L193"></a>got := &#34;&#34;;
    <a id="L194"></a>want := strings.Join(segments, &#34;&#34;);
    <a id="L195"></a>r := NewReader(&amp;StringReader{data: segments});
    <a id="L196"></a>for {
        <a id="L197"></a>rune, _, err := r.ReadRune();
        <a id="L198"></a>if err != nil {
            <a id="L199"></a>if err != os.EOF {
                <a id="L200"></a>return
            <a id="L201"></a>}
            <a id="L202"></a>break;
        <a id="L203"></a>}
        <a id="L204"></a>got += string(rune);
    <a id="L205"></a>}
    <a id="L206"></a>if got != want {
        <a id="L207"></a>t.Errorf(&#34;segments=%v got=%s want=%s&#34;, segments, got, want)
    <a id="L208"></a>}
<a id="L209"></a>}

<a id="L211"></a>var segmentList = [][]string{
    <a id="L212"></a>[]string{},
    <a id="L213"></a>[]string{&#34;&#34;},
    <a id="L214"></a>[]string{&#34;日&#34;, &#34;本語&#34;},
    <a id="L215"></a>[]string{&#34;\u65e5&#34;, &#34;\u672c&#34;, &#34;\u8a9e&#34;},
    <a id="L216"></a>[]string{&#34;\U000065e5&#34;, &#34;\U0000672c&#34;, &#34;\U00008a9e&#34;},
    <a id="L217"></a>[]string{&#34;\xe6&#34;, &#34;\x97\xa5\xe6&#34;, &#34;\x9c\xac\xe8\xaa\x9e&#34;},
    <a id="L218"></a>[]string{&#34;Hello&#34;, &#34;, &#34;, &#34;World&#34;, &#34;!&#34;},
    <a id="L219"></a>[]string{&#34;Hello&#34;, &#34;, &#34;, &#34;&#34;, &#34;World&#34;, &#34;!&#34;},
<a id="L220"></a>}

<a id="L222"></a>func TestReadRune(t *testing.T) {
    <a id="L223"></a>for _, s := range segmentList {
        <a id="L224"></a>readRuneSegments(t, s)
    <a id="L225"></a>}
<a id="L226"></a>}

<a id="L228"></a>func TestWriter(t *testing.T) {
    <a id="L229"></a>var data [8192]byte;

    <a id="L231"></a>for i := 0; i &lt; len(data); i++ {
        <a id="L232"></a>data[i] = byte(&#39; &#39; + i%(&#39;~&#39;-&#39; &#39;))
    <a id="L233"></a>}
    <a id="L234"></a>w := new(bytes.Buffer);
    <a id="L235"></a>for i := 0; i &lt; len(bufsizes); i++ {
        <a id="L236"></a>for j := 0; j &lt; len(bufsizes); j++ {
            <a id="L237"></a>nwrite := bufsizes[i];
            <a id="L238"></a>bs := bufsizes[j];

            <a id="L240"></a><span class="comment">// Write nwrite bytes using buffer size bs.</span>
            <a id="L241"></a><span class="comment">// Check that the right amount makes it out</span>
            <a id="L242"></a><span class="comment">// and that the data is correct.</span>

            <a id="L244"></a>w.Reset();
            <a id="L245"></a>buf, e := NewWriterSize(w, bs);
            <a id="L246"></a>context := fmt.Sprintf(&#34;nwrite=%d bufsize=%d&#34;, nwrite, bs);
            <a id="L247"></a>if e != nil {
                <a id="L248"></a>t.Errorf(&#34;%s: NewWriterSize %d: %v&#34;, context, bs, e);
                <a id="L249"></a>continue;
            <a id="L250"></a>}
            <a id="L251"></a>n, e1 := buf.Write(data[0:nwrite]);
            <a id="L252"></a>if e1 != nil || n != nwrite {
                <a id="L253"></a>t.Errorf(&#34;%s: buf.Write %d = %d, %v&#34;, context, nwrite, n, e1);
                <a id="L254"></a>continue;
            <a id="L255"></a>}
            <a id="L256"></a>if e = buf.Flush(); e != nil {
                <a id="L257"></a>t.Errorf(&#34;%s: buf.Flush = %v&#34;, context, e)
            <a id="L258"></a>}

            <a id="L260"></a>written := w.Bytes();
            <a id="L261"></a>if len(written) != nwrite {
                <a id="L262"></a>t.Errorf(&#34;%s: %d bytes written&#34;, context, len(written))
            <a id="L263"></a>}
            <a id="L264"></a>for l := 0; l &lt; len(written); l++ {
                <a id="L265"></a>if written[i] != data[i] {
                    <a id="L266"></a>t.Errorf(&#34;%s: wrong bytes written&#34;);
                    <a id="L267"></a>t.Errorf(&#34;want=%s&#34;, data[0:len(written)]);
                    <a id="L268"></a>t.Errorf(&#34;have=%s&#34;, written);
                <a id="L269"></a>}
            <a id="L270"></a>}
        <a id="L271"></a>}
    <a id="L272"></a>}
<a id="L273"></a>}

<a id="L275"></a><span class="comment">// Check that write errors are returned properly.</span>

<a id="L277"></a>type errorWriterTest struct {
    <a id="L278"></a>n, m   int;
    <a id="L279"></a>err    os.Error;
    <a id="L280"></a>expect os.Error;
<a id="L281"></a>}

<a id="L283"></a>func (w errorWriterTest) Write(p []byte) (int, os.Error) {
    <a id="L284"></a>return len(p) * w.n / w.m, w.err
<a id="L285"></a>}

<a id="L287"></a>var errorWriterTests = []errorWriterTest{
    <a id="L288"></a>errorWriterTest{0, 1, nil, io.ErrShortWrite},
    <a id="L289"></a>errorWriterTest{1, 2, nil, io.ErrShortWrite},
    <a id="L290"></a>errorWriterTest{1, 1, nil, nil},
    <a id="L291"></a>errorWriterTest{0, 1, os.EPIPE, os.EPIPE},
    <a id="L292"></a>errorWriterTest{1, 2, os.EPIPE, os.EPIPE},
    <a id="L293"></a>errorWriterTest{1, 1, os.EPIPE, os.EPIPE},
<a id="L294"></a>}

<a id="L296"></a>func TestWriteErrors(t *testing.T) {
    <a id="L297"></a>for _, w := range errorWriterTests {
        <a id="L298"></a>buf := NewWriter(w);
        <a id="L299"></a>_, e := buf.Write(strings.Bytes(&#34;hello world&#34;));
        <a id="L300"></a>if e != nil {
            <a id="L301"></a>t.Errorf(&#34;Write hello to %v: %v&#34;, w, e);
            <a id="L302"></a>continue;
        <a id="L303"></a>}
        <a id="L304"></a>e = buf.Flush();
        <a id="L305"></a>if e != w.expect {
            <a id="L306"></a>t.Errorf(&#34;Flush %v: got %v, wanted %v&#34;, w, e, w.expect)
        <a id="L307"></a>}
    <a id="L308"></a>}
<a id="L309"></a>}

<a id="L311"></a>func TestNewReaderSizeIdempotent(t *testing.T) {
    <a id="L312"></a>const BufSize = 1000;
    <a id="L313"></a>b, err := NewReaderSize(bytes.NewBufferString(&#34;hello world&#34;), BufSize);
    <a id="L314"></a>if err != nil {
        <a id="L315"></a>t.Error(&#34;NewReaderSize create fail&#34;, err)
    <a id="L316"></a>}
    <a id="L317"></a><span class="comment">// Does it recognize itself?</span>
    <a id="L318"></a>b1, err2 := NewReaderSize(b, BufSize);
    <a id="L319"></a>if err2 != nil {
        <a id="L320"></a>t.Error(&#34;NewReaderSize #2 create fail&#34;, err2)
    <a id="L321"></a>}
    <a id="L322"></a>if b1 != b {
        <a id="L323"></a>t.Error(&#34;NewReaderSize did not detect underlying Reader&#34;)
    <a id="L324"></a>}
    <a id="L325"></a><span class="comment">// Does it wrap if existing buffer is too small?</span>
    <a id="L326"></a>b2, err3 := NewReaderSize(b, 2*BufSize);
    <a id="L327"></a>if err3 != nil {
        <a id="L328"></a>t.Error(&#34;NewReaderSize #3 create fail&#34;, err3)
    <a id="L329"></a>}
    <a id="L330"></a>if b2 == b {
        <a id="L331"></a>t.Error(&#34;NewReaderSize did not enlarge buffer&#34;)
    <a id="L332"></a>}
<a id="L333"></a>}

<a id="L335"></a>func TestNewWriterSizeIdempotent(t *testing.T) {
    <a id="L336"></a>const BufSize = 1000;
    <a id="L337"></a>b, err := NewWriterSize(new(bytes.Buffer), BufSize);
    <a id="L338"></a>if err != nil {
        <a id="L339"></a>t.Error(&#34;NewWriterSize create fail&#34;, err)
    <a id="L340"></a>}
    <a id="L341"></a><span class="comment">// Does it recognize itself?</span>
    <a id="L342"></a>b1, err2 := NewWriterSize(b, BufSize);
    <a id="L343"></a>if err2 != nil {
        <a id="L344"></a>t.Error(&#34;NewWriterSize #2 create fail&#34;, err2)
    <a id="L345"></a>}
    <a id="L346"></a>if b1 != b {
        <a id="L347"></a>t.Error(&#34;NewWriterSize did not detect underlying Writer&#34;)
    <a id="L348"></a>}
    <a id="L349"></a><span class="comment">// Does it wrap if existing buffer is too small?</span>
    <a id="L350"></a>b2, err3 := NewWriterSize(b, 2*BufSize);
    <a id="L351"></a>if err3 != nil {
        <a id="L352"></a>t.Error(&#34;NewWriterSize #3 create fail&#34;, err3)
    <a id="L353"></a>}
    <a id="L354"></a>if b2 == b {
        <a id="L355"></a>t.Error(&#34;NewWriterSize did not enlarge buffer&#34;)
    <a id="L356"></a>}
<a id="L357"></a>}

<a id="L359"></a>func TestWriteString(t *testing.T) {
    <a id="L360"></a>const BufSize = 8;
    <a id="L361"></a>buf := new(bytes.Buffer);
    <a id="L362"></a>b, err := NewWriterSize(buf, BufSize);
    <a id="L363"></a>if err != nil {
        <a id="L364"></a>t.Error(&#34;NewWriterSize create fail&#34;, err)
    <a id="L365"></a>}
    <a id="L366"></a>b.WriteString(&#34;0&#34;);                         <span class="comment">// easy</span>
    <a id="L367"></a>b.WriteString(&#34;123456&#34;);                    <span class="comment">// still easy</span>
    <a id="L368"></a>b.WriteString(&#34;7890&#34;);                      <span class="comment">// easy after flush</span>
    <a id="L369"></a>b.WriteString(&#34;abcdefghijklmnopqrstuvwxy&#34;); <span class="comment">// hard</span>
    <a id="L370"></a>b.WriteString(&#34;z&#34;);
    <a id="L371"></a>b.Flush();
    <a id="L372"></a>if b.err != nil {
        <a id="L373"></a>t.Error(&#34;WriteString&#34;, b.err)
    <a id="L374"></a>}
    <a id="L375"></a>s := &#34;01234567890abcdefghijklmnopqrstuvwxyz&#34;;
    <a id="L376"></a>if string(buf.Bytes()) != s {
        <a id="L377"></a>t.Errorf(&#34;WriteString wants %q gets %q&#34;, s, string(buf.Bytes()))
    <a id="L378"></a>}
<a id="L379"></a>}
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
